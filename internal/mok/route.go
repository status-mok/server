package mok

import (
	"sync"
)

type Route interface {
	ExpectationStorage

	URL() string
	Type() RouteType
	Disabled() bool

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

func (e *route) Type() RouteType {
	return e._type
}

func (e *route) Disabled() bool {
	return e.isDisabled
}
