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

type ServerStorage interface {
	ServerGet(ctx context.Context, name string) (Server, error)
	ServerCreate(ctx context.Context, srv Server) error
	ServerDelete(ctx context.Context, name string) error
}

type serverStorage struct {
	storage map[string]Server

	mu sync.RWMutex
}

func NewServerStorage(items ...Server) *serverStorage {
	srvStorage := &serverStorage{
		storage: make(map[string]Server, len(items)),
	}

	for _, item := range items {
		srvStorage.storage[item.Name()] = item
	}

	return srvStorage
}

func (m *serverStorage) ServerGet(_ context.Context, name string) (Server, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	s, ok := m.storage[name]
	if !ok {
		return nil, ErrNotFound
	}

	return s, nil
}

func (m *serverStorage) ServerCreate(ctx context.Context, srv Server) error {
	if s, err := m.ServerGet(ctx, srv.Name()); err != nil {
		if !errors.Is(err, ErrNotFound) {
			return err
		}
	} else if s != nil {
		return ErrAlreadyExist
	}

	m.mu.Lock()
	defer m.mu.Unlock()

	m.storage[srv.Name()] = srv

	return nil
}

func (m *serverStorage) ServerDelete(ctx context.Context, name string) error {
	if _, err := m.ServerGet(ctx, name); err != nil {
		return err
	}

	m.mu.Lock()
	defer m.mu.Unlock()

	delete(m.storage, name)

	return nil
}
