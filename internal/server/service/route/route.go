package route

import (
	"context"

	"github.com/status-mok/server/internal/mok"
	"github.com/status-mok/server/internal/pkg/errors"
	routeAPI "github.com/status-mok/server/pkg/route-api"
)

type routeService struct {
	storage mok.ServerStorage

	routeAPI.UnimplementedRouteServiceServer
}

func NewRouteService(storage mok.ServerStorage) *routeService {
	return &routeService{
		storage: storage,
	}
}

func (s routeService) Create(ctx context.Context, req *routeAPI.CreateRequest) (*routeAPI.CreateResponse, error) {
	routeType := mok.RouteType(req.GetType().Number())
	if err := routeType.Validate(); err != nil {
		return nil, err
	}

	srv, err := s.storage.ServerGet(ctx, req.GetServerName())
	if err != nil {
		return nil, err
	}

	ept := mok.NewRoute(req.GetUrl(), routeType)

	if err := srv.RouteCreate(ctx, ept); err != nil {
		return nil, errors.Wrapf(err, "failed to create route '%s' for server '%s'", req.GetUrl(), req.GetServerName())
	}

	return &routeAPI.CreateResponse{Success: true}, nil
}

func (s routeService) Delete(ctx context.Context, req *routeAPI.DeleteRequest) (*routeAPI.DeleteResponse, error) {
	srv, err := s.storage.ServerGet(ctx, req.GetServerName())
	if err != nil {
		return nil, err
	}

	if err := srv.RouteDelete(ctx, req.GetUrl()); err != nil {
		return nil, errors.Wrapf(err, "failed to delete route '%s' from server '%s'", req.GetUrl(), req.GetServerName())
	}

	return &routeAPI.DeleteResponse{Success: true}, nil
}
