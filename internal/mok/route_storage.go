package mok

import (
	"context"
	"fmt"
	"net/http"
	"sync"

	"github.com/go-chi/chi"
	"github.com/status-mok/server/internal/pkg/errors"
)

type RouteStorage interface {
	RouteGet(ctx context.Context, url string) (Route, error)
	RouteCreate(ctx context.Context, rt Route) error
	RouteDelete(ctx context.Context, url string) error
}

type routeStorage struct {
	storage      []*routeItem
	storageIndex map[string]int

	httpMux chi.Router

	mu sync.RWMutex
}

type routeItem struct {
	route   Route
	handler http.Handler
}

func NewRouteStorage(items ...Route) (*routeStorage, error) {
	ctx := context.Background()

	rtStorage := &routeStorage{
		storage:      make([]*routeItem, 0, len(items)),
		storageIndex: make(map[string]int, len(items)),
		httpMux:      chi.NewMux(),
	}

	for _, item := range items {
		if err := rtStorage.RouteCreate(ctx, item); err != nil {
			return nil, err
		}
	}

	return rtStorage, nil
}

func (rs *routeStorage) RouteGet(_ context.Context, url string) (Route, error) {
	rs.mu.RLock()
	defer rs.mu.RUnlock()

	idx, ok := rs.storageIndex[url]
	if !ok {
		return nil, ErrNotFound
	}

	return rs.storage[idx].route, nil
}

func (rs *routeStorage) RouteCreate(ctx context.Context, rt Route) error {
	if s, err := rs.RouteGet(ctx, rt.URL()); err != nil {
		if !errors.Is(err, ErrNotFound) {
			return err
		}
	} else if s != nil {
		return ErrAlreadyExist
	}

	rth := &routeItem{
		route:   rt,
		handler: routeHTTPHandler(rt),
	}

	rs.mu.Lock()
	defer rs.mu.Unlock()

	if err := rs.addToMux(rth); err != nil {
		return err
	}

	rs.storage = append(rs.storage, rth)
	rs.storageIndex[rt.URL()] = len(rs.storage) - 1

	return nil
}

func (rs *routeStorage) RouteDelete(ctx context.Context, url string) error {
	if _, err := rs.RouteGet(ctx, url); err != nil {
		return err
	}

	rs.mu.Lock()
	defer rs.mu.Unlock()

	idx, ok := rs.storageIndex[url]
	if !ok {
		return ErrNotFound
	}

	rs.storage[idx] = nil
	delete(rs.storageIndex, url)

	if err := rs.recreateMux(); err != nil {
		return err
	}

	return nil
}

func (rs *routeStorage) addToMux(rth *routeItem) (err error) {
	defer func() {
		if e := recover(); e != nil {
			switch v := e.(type) {
			case string:
				err = errors.New(v)
			case error:
				err = v
			default:
				err = fmt.Errorf("'%v'", e)
			}
		}
	}()

	rs.httpMux.Handle(rth.route.URL(), rth.handler)

	return nil
}

func (rs *routeStorage) recreateMux() error {
	newMux := chi.NewMux()

	for _, item := range rs.storage {
		if item == nil {
			continue
		}

		newMux.Handle(item.route.URL(), item.handler)
	}

	rs.httpMux = newMux

	return nil
}
