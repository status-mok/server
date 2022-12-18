package mok

import (
	"sync"
)

type Route interface {
	ExpectationStorage

	URL() string

	// FindExpectation(ctx context.Context, req *http.Request) (err error)
}

type route struct {
	*expectationStorage

	url   string    `mapstructure:"url"`
	_type RouteType `mapstructure:"type"`

	isDisabled bool

	mu sync.Mutex
}

func NewRoute(url string, _type RouteType) *route {
	return &route{
		expectationStorage: NewExpectationStorage(),

		url:   url,
		_type: _type,
	}
}

func (e *route) URL() string {
	return e.url
}
