package mok

import (
	"context"
	"sync"

	"github.com/status-mok/server/internal/pkg/errors"
)

var (
	ErrNoMatchFound = errors.New("no match found")
)

type ExpectationStorage interface {
	ExpectationGet(ctx context.Context, id string) (Expectation, error)
	ExpectationCreate(ctx context.Context, exp Expectation) error
	ExpectationDelete(ctx context.Context, id string) error
	ExpectationFindMatch(ctx context.Context, endpointType EndpointType, req any) (Expectation, error)
}

type expectationStorage struct {
	storage map[string]Expectation

	mu sync.RWMutex
}

func NewExpectationStorage(items ...Expectation) *expectationStorage {
	expStorage := &expectationStorage{
		storage: make(map[string]Expectation, len(items)),
	}

	for _, item := range items {
		expStorage.storage[item.ID()] = item
	}

	return expStorage
}

func (e *expectationStorage) ExpectationGet(_ context.Context, id string) (Expectation, error) {
	e.mu.RLock()
	defer e.mu.RUnlock()

	exp, ok := e.storage[id]
	if !ok {
		return nil, ErrNotFound
	}

	return exp, nil
}

func (e *expectationStorage) ExpectationCreate(ctx context.Context, exp Expectation) error {
	if s, err := e.ExpectationGet(ctx, exp.ID()); err != nil {
		if !errors.Is(err, ErrNotFound) {
			return err
		}
	} else if s != nil {
		return ErrAlreadyExist
	}

	e.mu.Lock()
	defer e.mu.Unlock()

	e.storage[exp.ID()] = exp

	return nil
}

func (e *expectationStorage) ExpectationDelete(ctx context.Context, id string) error {
	if _, err := e.ExpectationGet(ctx, id); err != nil {
		return err
	}

	e.mu.Lock()
	defer e.mu.Unlock()

	delete(e.storage, id)

	return nil
}

func (e *expectationStorage) ExpectationFindMatch(ctx context.Context, endpointType EndpointType, req any) (Expectation, error) {
	e.mu.RLock()
	defer e.mu.RUnlock()

	for _, exp := range e.storage {
		if exp.Match(ctx, endpointType, req) {
			return exp, nil
		}
	}

	return nil, ErrNoMatchFound
}
