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
	"github.com/status-mok/server/internal/pkg/errors"
	"github.com/status-mok/server/internal/pkg/log"
	"github.com/status-mok/server/internal/server/config"
	"github.com/status-mok/server/internal/server/docs"
	serverAPI "github.com/status-mok/server/pkg/server-api"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type app struct {
	serverService serverAPI.ServerServiceServer

	conf *config.AppConfig
}

func NewApp() *app {
	return &app{}
}

func (app *app) Start(ctx context.Context, configPath string) error {
	var err error
	app.conf, err = config.NewAppConfig(ctx, configPath)
	if err != nil {
		return err
	}

	log.SetLogger(log.New(app.conf.LogLevel(), os.Stdout))

	app.initServices(ctx)

	errGroup, errCtx := errgroup.WithContext(ctx)

	errGroup.Go(func() error {
		errG := app.startGRPCServer(errCtx)
		if errG != nil {
			return errors.Wrap(errG, "failed to start grpc admin API server")
		}

		return nil
	})

	errGroup.Go(func() error {
		errH := app.startHTTPServer(errCtx)
		if errH != nil {
			return errors.Wrap(errH, "failed to start http admin API server")
		}

		return nil
	})

	return errGroup.Wait()
}

func (app *app) initServices(_ context.Context) {
	app.serverService = new(serverAPI.UnimplementedServerServiceServer)
}

func (app *app) startHTTPServer(ctx context.Context) error {
	mux := chi.NewMux()

	mux.Use(middleware.RealIP)
	middleware.DefaultLogger = middleware.RequestLogger(&middleware.DefaultLogFormatter{Logger: log.StdLogger(ctx), NoColor: true})
	mux.Use(middleware.Logger)
	mux.Use(middleware.Recoverer)

	mux.Mount("/metrics", promhttp.Handler())
	mux.Mount("/docs", docs.Handler())

	gwMux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	if err := serverAPI.RegisterServerServiceHandlerFromEndpoint(ctx, gwMux, app.conf.AdminGRPCAddress(), opts); err != nil {
		return err
	}
	mux.Handle("/", gwMux)

	srv := http.Server{
		Addr:     app.conf.AdminHTTPAddress(),
		Handler:  mux,
		ErrorLog: log.StdLogger(ctx),
	}

	go func() {
		blockUntilExitSignalOrCtxTermination(ctx)

		timeoutCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		if err := srv.Shutdown(timeoutCtx); err != nil {
			log.Logger(ctx).Errorf("failed to stop http server: '%v'", err)
		}
		log.Logger(ctx).Info("admin api http server stopped")
	}()

	log.Logger(ctx).Infof("admin api http server listening at '%v'", app.conf.AdminHTTPAddress())
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		return err
	}

	return nil
}

func (app *app) startGRPCServer(ctx context.Context) error {
	listener, err := net.Listen("tcp", app.conf.AdminGRPCAddress())
	if err != nil {
		log.Logger(ctx).Fatalf("failed to init tcp listener: '%v'", err)
	}

	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	serverAPI.RegisterServerServiceServer(grpcServer, app.serverService)

	go func() {
		blockUntilExitSignalOrCtxTermination(ctx)

		grpcServer.GracefulStop()
		log.Logger(ctx).Info("admin api grpc server stopped")
	}()

	log.Logger(ctx).Infof("admin api grpc server listening on '%v'", app.conf.AdminGRPCAddress())
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
