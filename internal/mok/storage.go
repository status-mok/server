package mok

import (
	"context"
	"errors"
	"sync"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	ErrNotFound     = status.Error(codes.NotFound, "not found")
	ErrAlreadyExist = status.Error(codes.AlreadyExists, "already exist")
)

type Storage interface {
	ServerGet(ctx context.Context, name string) (Server, error)
	ServerCreate(ctx context.Context, srv Server) error
	ServerDelete(ctx context.Context, name string) error
}

type storage struct {
	servers map[string]Server

	mu sync.RWMutex
}

func NewStorage() *storage {
	return &storage{
		servers: make(map[string]Server),
	}
}

func (m *storage) ServerGet(_ context.Context, name string) (Server, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	s, ok := m.servers[name]
	if !ok {
		return nil, ErrNotFound
	}

	return s, nil
}

func (m *storage) ServerCreate(ctx context.Context, srv Server) error {
	if s, err := m.ServerGet(ctx, srv.Name()); err != nil {
		if !errors.Is(err, ErrNotFound) {
			return err
		}
	} else if s != nil {
		return ErrAlreadyExist
	}

	m.mu.Lock()
	defer m.mu.Unlock()

	m.servers[srv.Name()] = srv

	return nil
}

func (m *storage) ServerDelete(ctx context.Context, name string) error {
	if _, err := m.ServerGet(ctx, name); err != nil {
		return err
	}

	m.mu.Lock()
	defer m.mu.Unlock()

	delete(m.servers, name)

	return nil
}
