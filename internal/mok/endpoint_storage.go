package mok

import (
	"context"
	"sync"

	"github.com/status-mok/server/internal/pkg/errors"
)

type EndpointStorage interface {
	EndpointGet(ctx context.Context, url string) (Endpoint, error)
	EndpointCreate(ctx context.Context, ept Endpoint) error
	EndpointDelete(ctx context.Context, url string) error
}

type endpointStorage struct {
	storage map[string]Endpoint

	mu sync.RWMutex
}

func NewEndpointStorage(items ...Endpoint) *endpointStorage {
	eptStorage := &endpointStorage{
		storage: make(map[string]Endpoint, len(items)),
	}

	for _, item := range items {
		eptStorage.storage[item.URL()] = item
	}

	return eptStorage
}

func (e *endpointStorage) EndpointGet(_ context.Context, url string) (Endpoint, error) {
	e.mu.RLock()
	defer e.mu.RUnlock()

	ept, ok := e.storage[url]
	if !ok {
		return nil, ErrNotFound
	}

	return ept, nil
}

func (e *endpointStorage) EndpointCreate(ctx context.Context, ept Endpoint) error {
	if s, err := e.EndpointGet(ctx, ept.URL()); err != nil {
		if !errors.Is(err, ErrNotFound) {
			return err
		}
	} else if s != nil {
		return ErrAlreadyExist
	}

	e.mu.Lock()
	defer e.mu.Unlock()

	e.storage[ept.URL()] = ept

	return nil
}

func (e *endpointStorage) EndpointDelete(ctx context.Context, url string) error {
	if _, err := e.EndpointGet(ctx, url); err != nil {
		return err
	}

	e.mu.Lock()
	defer e.mu.Unlock()

	delete(e.storage, url)

	return nil
}
