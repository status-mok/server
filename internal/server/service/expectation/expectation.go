package expectation

import (
	"context"

	"github.com/status-mok/server/internal/mok"
	"github.com/status-mok/server/internal/pkg/errors"
	expectationAPI "github.com/status-mok/server/pkg/expectation-api"
)

type expectationService struct {
	storage mok.ServerStorage

	expectationAPI.UnimplementedExpectationServiceServer
}

func NewExpectationService(storage mok.ServerStorage) *expectationService {
	return &expectationService{
		storage: storage,
	}
}

func (s expectationService) Create(ctx context.Context, req *expectationAPI.CreateRequest) (*expectationAPI.CreateResponse, error) {
	srv, err := s.storage.ServerGet(ctx, req.GetServerName())
	if err != nil {
		return nil, err
	}

	rt, err := srv.RouteGet(ctx, req.GetRouteUrl())
	if err != nil {
		return nil, err
	}

	exp := mok.NewExpectation(req.GetId())

	if err := rt.ExpectationCreate(ctx, exp); err != nil {
		return nil, errors.Wrapf(err, "failed to create expectation '%s' for route '%s' on server '%s'", req.GetId(), req.GetRouteUrl(), req.GetServerName())
	}

	return &expectationAPI.CreateResponse{Success: true}, nil
}

func (s expectationService) Delete(ctx context.Context, req *expectationAPI.DeleteRequest) (*expectationAPI.DeleteResponse, error) {
	srv, err := s.storage.ServerGet(ctx, req.GetServerName())
	if err != nil {
		return nil, err
	}

	rt, err := srv.RouteGet(ctx, req.GetRouteUrl())
	if err != nil {
		return nil, err
	}

	if err := rt.ExpectationDelete(ctx, req.GetId()); err != nil {
		return nil, errors.Wrapf(err, "failed to delete expectation '%s' from route '%s' on server '%s'", req.GetId(), req.GetRouteUrl(), req.GetServerName())
	}

	return &expectationAPI.DeleteResponse{Success: true}, nil
}
