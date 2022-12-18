package mok

import (
	"context"
	"sync"

	"github.com/status-mok/server/internal/pkg/errors"
)

type RouteStorage interface {
	RouteGet(ctx context.Context, url string) (Route, error)
	RouteCreate(ctx context.Context, ept Route) error
	RouteDelete(ctx context.Context, url string) error
}

type routeStorage struct {
	storage map[string]Route

	mu sync.RWMutex
}

func NewRouteStorage(items ...Route) *routeStorage {
	rtStorage := &routeStorage{
		storage: make(map[string]Route, len(items)),
	}

	for _, item := range items {
		rtStorage.storage[item.URL()] = item
	}

	return rtStorage
}

func (e *routeStorage) RouteGet(_ context.Context, url string) (Route, error) {
	e.mu.RLock()
	defer e.mu.RUnlock()

	ept, ok := e.storage[url]
	if !ok {
		return nil, ErrNotFound
	}

	return ept, nil
}

func (e *routeStorage) RouteCreate(ctx context.Context, ept Route) error {
	if s, err := e.RouteGet(ctx, ept.URL()); err != nil {
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

func (e *routeStorage) RouteDelete(ctx context.Context, url string) error {
	if _, err := e.RouteGet(ctx, url); err != nil {
		return err
	}

	e.mu.Lock()
	defer e.mu.Unlock()

	delete(e.storage, url)

	return nil
}
