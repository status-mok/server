package mok

import (
	"context"
	"sync"

	uuid "github.com/satori/go.uuid"
)

type Expectation interface {
	ID() string

	Match(ctx context.Context, routeType RouteType, req any) bool
}

type expectation struct {
	id string

	isDisabled bool

	mu sync.Mutex
}

func NewExpectation(id string) *expectation {
	if len(id) == 0 {
		id = uuid.NewV4().String()
	}

	return &expectation{
		id: id,
	}
}

func (e *expectation) ID() string {
	return e.id
}

func (e *expectation) Match(_ context.Context, _ RouteType, _ any) bool {
	return true
}
