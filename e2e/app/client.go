//go:build e2e

package app

//go:generate mkdir -p http-client

//go:generate swagger generate client -c client/server_api -q -f ../../pkg/server-api/server_api.swagger.json -t http-client

//go:generate swagger generate client -c client/route_api -q -f ../../pkg/route-api/route_api.swagger.json -t http-client

//go:generate swagger generate client -c client/expectation_api -q -f ../../pkg/expectation-api/expectation_api.swagger.json -t http-client

import (
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	expectationHTTPapi "github.com/status-mok/server/e2e/app/http-client/client/expectation_api/expectation_service"
	routeHTTPapi "github.com/status-mok/server/e2e/app/http-client/client/route_api/route_service"
	serverHTTPapi "github.com/status-mok/server/e2e/app/http-client/client/server_api/server_service"
	expectationAPI "github.com/status-mok/server/pkg/expectation-api"
	routeAPI "github.com/status-mok/server/pkg/route-api"
	serverAPI "github.com/status-mok/server/pkg/server-api"
	"google.golang.org/grpc"
)

type grpcClient struct {
	serverService      serverAPI.ServerServiceClient
	routeService       routeAPI.RouteServiceClient
	expectationService expectationAPI.ExpectationServiceClient
}

func NewGRPCClient(grpcConn *grpc.ClientConn) *grpcClient {
	return &grpcClient{
		serverService:      serverAPI.NewServerServiceClient(grpcConn),
		routeService:       routeAPI.NewRouteServiceClient(grpcConn),
		expectationService: expectationAPI.NewExpectationServiceClient(grpcConn),
	}
}

func (c *grpcClient) ServerService() serverAPI.ServerServiceClient {
	return c.serverService
}

func (c *grpcClient) RouteService() routeAPI.RouteServiceClient {
	return c.routeService
}

func (c *grpcClient) ExpectationService() expectationAPI.ExpectationServiceClient {
	return c.expectationService
}

type httpClient struct {
	serverService      serverHTTPapi.ClientService
	routeService       routeHTTPapi.ClientService
	expectationService expectationHTTPapi.ClientService
}

func NewHTTPClient(httpTransport *httptransport.Runtime) *httpClient {
	return &httpClient{
		serverService:      serverHTTPapi.New(httpTransport, strfmt.Default),
		routeService:       routeHTTPapi.New(httpTransport, strfmt.Default),
		expectationService: expectationHTTPapi.New(httpTransport, strfmt.Default),
	}
}

func (c *httpClient) ServerService() serverHTTPapi.ClientService {
	return c.serverService
}

func (c *httpClient) RouteService() routeHTTPapi.ClientService {
	return c.routeService
}

func (c *httpClient) ExpectationService() expectationHTTPapi.ClientService {
	return c.expectationService
}
