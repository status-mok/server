package app

import (
	"context"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/status-mok/server/internal/mok"
	"github.com/status-mok/server/internal/pkg/docs"
	"github.com/status-mok/server/internal/pkg/errors"
	"github.com/status-mok/server/internal/pkg/log"
	grpcMiddleware "github.com/status-mok/server/internal/pkg/middleware/grpc"
	"github.com/status-mok/server/internal/server/config"
	"github.com/status-mok/server/internal/server/service/expectation"
	"github.com/status-mok/server/internal/server/service/route"
	"github.com/status-mok/server/internal/server/service/server"
	expectationAPI "github.com/status-mok/server/pkg/expectation-api"
	routeAPI "github.com/status-mok/server/pkg/route-api"
	serverAPI "github.com/status-mok/server/pkg/server-api"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type app struct {
	serverService      serverAPI.ServerServiceServer
	routeService       routeAPI.RouteServiceServer
	expectationService expectationAPI.ExpectationServiceServer

	conf *config.AppConfig
}

type grpcGatewayRegister func(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error

func NewApp(ctx context.Context, configPath string) (*app, error) {
	conf, err := config.NewAppConfig(ctx, configPath)
	if err != nil {
		return nil, err
	}

	log.SetLogger(log.New(conf.LogLevel(), os.Stdout))

	return &app{conf: conf}, nil
}

func (a *app) Run(ctx context.Context) error {
	a.initServices(ctx)

	errGroup, errCtx := errgroup.WithContext(ctx)

	errGroup.Go(func() error {
		errG := a.startGRPCServer(errCtx)
		if errG != nil {
			return errors.Wrap(errG, "failed to start grpc admin API server")
		}

		return nil
	})

	errGroup.Go(func() error {
		errH := a.startHTTPServer(errCtx)
		if errH != nil {
			return errors.Wrap(errH, "failed to start http admin API server")
		}

		return nil
	})

	return errGroup.Wait()
}

func (a *app) initServices(_ context.Context) {
	storage := mok.NewServerStorage()

	a.serverService = server.NewServerService(storage)
	a.routeService = route.NewRouteService(storage)
	a.expectationService = expectation.NewExpectationService(storage)
}

func (a *app) startHTTPServer(ctx context.Context) error {
	mux := chi.NewMux()

	mux.Use(middleware.RealIP)
	middleware.DefaultLogger = middleware.RequestLogger(&middleware.DefaultLogFormatter{Logger: log.StdLogger(ctx), NoColor: true})
	mux.Use(middleware.Logger)
	mux.Use(middleware.Recoverer)

	mux.Mount("/metrics", promhttp.Handler())

	serviceDocs := []docs.ServiceDoc{
		{"server-api", serverAPI.SwaggerJSON},
		{"route-api", routeAPI.SwaggerJSON},
		{"expectation-api", expectationAPI.SwaggerJSON},
	}
	mux.Mount("/docs", docs.NewServiceDocsHandler(serviceDocs...))

	gwMux := runtime.NewServeMux()
	mux.Mount("/", gwMux)

	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	serviceRegisters := []grpcGatewayRegister{
		serverAPI.RegisterServerServiceHandlerFromEndpoint,
		routeAPI.RegisterRouteServiceHandlerFromEndpoint,
		expectationAPI.RegisterExpectationServiceHandlerFromEndpoint,
	}
	for _, registerFn := range serviceRegisters {
		if err := registerFn(ctx, gwMux, a.conf.AdminGRPCAddress(), opts); err != nil {
			return err
		}
	}

	srv := http.Server{
		Addr:     a.conf.AdminHTTPAddress(),
		Handler:  mux,
		ErrorLog: log.StdLogger(ctx),
	}

	go func() {
		blockUntilExitSignalOrCtxTermination(ctx)

		timeoutCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		if err := srv.Shutdown(timeoutCtx); err != nil {
			log.L(ctx).Errorf("failed to stop http server: '%v'", err)
		}
		log.L(ctx).Info("admin api http server stopped")
	}()

	log.L(ctx).Infof("admin api http server listening at '%v'", a.conf.AdminHTTPAddress())
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		return err
	}

	return nil
}

func (a *app) startGRPCServer(ctx context.Context) error {
	listener, err := net.Listen("tcp", a.conf.AdminGRPCAddress())
	if err != nil {
		log.L(ctx).Fatalf("failed to init tcp listener: '%v'", err)
	}

	opts := []grpc.ServerOption{
		grpc.ChainUnaryInterceptor(
			grpcMiddleware.NewErrorUnaryInterceptor(),
			grpcMiddleware.NewMessageValidatorInterceptor(),
		),
	}
	grpcServer := grpc.NewServer(opts...)
	serverAPI.RegisterServerServiceServer(grpcServer, a.serverService)
	routeAPI.RegisterRouteServiceServer(grpcServer, a.routeService)
	expectationAPI.RegisterExpectationServiceServer(grpcServer, a.expectationService)

	go func() {
		blockUntilExitSignalOrCtxTermination(ctx)

		grpcServer.GracefulStop()
		log.L(ctx).Info("admin api grpc server stopped")
	}()

	log.L(ctx).Infof("admin api grpc server listening on '%v'", a.conf.AdminGRPCAddress())
	if err = grpcServer.Serve(listener); err != nil {
		return err
	}

	return nil
}

func blockUntilExitSignalOrCtxTermination(ctx context.Context) {
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	for {
		select {
		case <-done:
			return
		case <-ctx.Done():
			return
		}
	}
}
