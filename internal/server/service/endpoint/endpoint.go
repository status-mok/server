package endpoint

import (
	"context"

	"github.com/status-mok/server/internal/mok"
	"github.com/status-mok/server/internal/pkg/errors"
	endpointAPI "github.com/status-mok/server/pkg/endpoint-api"
)

type endpointService struct {
	storage mok.ServerStorage

	endpointAPI.UnimplementedEndpointServiceServer
}

func NewEndpointService(storage mok.ServerStorage) *endpointService {
	return &endpointService{
		storage: storage,
	}
}

func (s endpointService) Create(ctx context.Context, req *endpointAPI.CreateRequest) (*endpointAPI.CreateResponse, error) {
	endpointType := mok.EndpointType(req.GetType().Number())
	if err := endpointType.Validate(); err != nil {
		return nil, err
	}

	srv, err := s.storage.ServerGet(ctx, req.GetServerName())
	if err != nil {
		return nil, err
	}

	ept := mok.NewEndpoint(req.GetUrl(), endpointType)

	if err := srv.EndpointCreate(ctx, ept); err != nil {
		return nil, errors.Wrapf(err, "failed to create endpoint '%s' for server '%s'", req.GetUrl(), req.GetServerName())
	}

	return &endpointAPI.CreateResponse{Success: true}, nil
}

func (s endpointService) Delete(ctx context.Context, req *endpointAPI.DeleteRequest) (*endpointAPI.DeleteResponse, error) {
	srv, err := s.storage.ServerGet(ctx, req.GetServerName())
	if err != nil {
		return nil, err
	}

	if err := srv.EndpointDelete(ctx, req.GetUrl()); err != nil {
		return nil, errors.Wrapf(err, "failed to delete endpoint '%s' from server '%s'", req.GetUrl(), req.GetServerName())
	}

	return &endpointAPI.DeleteResponse{Success: true}, nil
}
