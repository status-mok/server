//go:build e2e

package server

import (
	"context"
	"fmt"

	"github.com/status-mok/server/internal/pkg/tester"
	"github.com/status-mok/server/internal/server/app"
	"github.com/status-mok/server/internal/server/config"
	expectationAPI "github.com/status-mok/server/pkg/expectation-api"
	routeAPI "github.com/status-mok/server/pkg/route-api"
	serverAPI "github.com/status-mok/server/pkg/server-api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type application interface {
	Run(ctx context.Context) error
}

type TestServer struct {
	app  application
	conf *config.AppConfig

	grpcConn  *grpc.ClientConn
	ctxCancel func()
}

func NewServer() *TestServer {
	ctx, cancel := context.WithCancel(context.Background())

	a, err := app.NewApp(ctx, "")
	if err != nil {
		panic(err)
	}

	conf := new(config.AppConfig)
	conf.AdminAPI.GRPC.Port = fmt.Sprint(tester.GetFreePort())
	conf.AdminAPI.HTTP.Port = fmt.Sprint(tester.GetFreePort())
	a.SetConfig(conf)

	go func() {
		if err = a.Run(ctx); err != nil {
			panic(err)
		}
	}()

	opts := []grpc.DialOption{
		grpc.WithBlock(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}
	conn, err := grpc.Dial(conf.AdminGRPCAddress(), opts...)
	if err != nil {
		panic(err)
	}

	return &TestServer{
		app:  a,
		conf: conf,

		grpcConn:  conn,
		ctxCancel: cancel,
	}
}

func (s *TestServer) Close() {
	_ = s.grpcConn.Close()
	s.ctxCancel()
	return
}

func (s *TestServer) ServerGRPCClient() serverAPI.ServerServiceClient {
	return serverAPI.NewServerServiceClient(s.grpcConn)
}

func (s *TestServer) RouteGRPCClient() routeAPI.RouteServiceClient {
	return routeAPI.NewRouteServiceClient(s.grpcConn)
}

func (s *TestServer) ExpectationGRPCClient() expectationAPI.ExpectationServiceClient {
	return expectationAPI.NewExpectationServiceClient(s.grpcConn)
}
