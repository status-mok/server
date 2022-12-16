package server

import (
	"context"

	"github.com/status-mok/server/internal/mok"
	"github.com/status-mok/server/internal/pkg/errors"
	serverAPI "github.com/status-mok/server/pkg/server-api"
)

type serverService struct {
	storage *mok.Storage

	serverAPI.UnimplementedServerServiceServer
}

func NewServerService(storage *mok.Storage) *serverService {
	return &serverService{
		storage: storage,
	}
}

func (s serverService) Create(ctx context.Context, req *serverAPI.CreateRequest) (*serverAPI.CreateResponse, error) {
	serverType := mok.ServerType(req.GetType().Number())
	if err := serverType.Validate(); err != nil {
		return nil, err
	}

	srv := mok.NewServer(
		req.GetName(),
		req.GetIp(),
		req.GetPort(),
		serverType,
	)

	if err := s.storage.ServerCreate(ctx, srv); err != nil {
		return nil, errors.Wrapf(err, "failed to create server '%s'", req.GetName())
	}

	return &serverAPI.CreateResponse{Success: true}, nil
}

func (s serverService) Delete(ctx context.Context, req *serverAPI.DeleteRequest) (*serverAPI.DeleteResponse, error) {
	if err := s.storage.ServerDelete(ctx, req.GetName()); err != nil {
		return nil, errors.Wrapf(err, "failed to delete server '%s'", req.GetName())
	}

	return &serverAPI.DeleteResponse{Success: true}, nil
}

func (s serverService) Start(ctx context.Context, req *serverAPI.StartRequest) (*serverAPI.StartResponse, error) {
	srv, err := s.storage.ServerGet(ctx, req.GetName())
	if err != nil {
		return nil, errors.Wrapf(err, "failed to get server '%s'", req.GetName())
	}

	if err = srv.Start(ctx); err != nil {
		return nil, errors.Wrapf(err, "failed to start server '%s'", req.GetName())
	}

	return &serverAPI.StartResponse{Success: true}, nil
}

func (s serverService) Stop(ctx context.Context, req *serverAPI.StopRequest) (*serverAPI.StopResponse, error) {
	srv, err := s.storage.ServerGet(ctx, req.GetName())
	if err != nil {
		return nil, errors.Wrapf(err, "failed to get server '%s'", req.GetName())
	}

	if err = srv.Stop(ctx); err != nil {
		return nil, errors.Wrapf(err, "failed to stop server '%s'", req.GetName())
	}

	return &serverAPI.StopResponse{Success: true}, nil
}
