package mok

import (
	"sync"
)

type Endpoint interface {
	ExpectationStorage

	URL() string

	// FindExpectation(ctx context.Context, req *http.Request) (err error)
}

type endpoint struct {
	*expectationStorage

	url   string       `mapstructure:"url"`
	_type EndpointType `mapstructure:"type"`

	isDisabled bool

	mu sync.Mutex
}

func NewEndpoint(url string, _type EndpointType) *endpoint {
	return &endpoint{
		expectationStorage: NewExpectationStorage(),

		url:   url,
		_type: _type,
	}
}

func (e *endpoint) URL() string {
	return e.url
}
