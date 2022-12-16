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

type Storage struct {
	servers map[string]*mokServer

	mu sync.RWMutex
}

func NewStorage() *Storage {
	return &Storage{
		servers: make(map[string]*mokServer),
	}
}

func (m *Storage) ServerGet(_ context.Context, name string) (*mokServer, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	s, ok := m.servers[name]
	if !ok {
		return nil, ErrNotFound
	}

	return s, nil
}

func (m *Storage) ServerCreate(ctx context.Context, srv *mokServer) error {
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

func (m *Storage) ServerDelete(ctx context.Context, name string) error {
	if _, err := m.ServerGet(ctx, name); err != nil {
		return err
	}

	m.mu.Lock()
	defer m.mu.Unlock()

	delete(m.servers, name)

	return nil
}
