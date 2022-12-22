//go:build e2e

package app

import (
	"context"
	"fmt"

	httptransport "github.com/go-openapi/runtime/client"
	"github.com/status-mok/server/internal/pkg/tester"
	"github.com/status-mok/server/internal/server/app"
	"github.com/status-mok/server/internal/server/config"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type application interface {
	Run(ctx context.Context) error
}

type TestAppServer struct {
	app  application
	conf *config.AppConfig

	grpcClient *grpcClient
	httpClient *httpClient

	ctxCancel func()
}

func NewAppServer() *TestAppServer {
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

	return &TestAppServer{
		app:  a,
		conf: conf,

		grpcClient: NewGRPCClient(conn),
		httpClient: NewHTTPClient(httptransport.New(conf.AdminHTTPAddress(), "", []string{"http"})),

		ctxCancel: cancel,
	}
}

func (s *TestAppServer) Close() {
	s.ctxCancel()
	return
}

func (s *TestAppServer) GRPCClient() *grpcClient {
	return s.grpcClient
}

func (s *TestAppServer) HTTPClient() *httpClient {
	return s.httpClient
}
