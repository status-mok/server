package mok

import (
	"github.com/status-mok/server/internal/pkg/errors"
)

var ErrEndpointTypeUnknown = errors.New("unknown endpoint type")

type EndpointType int32

const (
	EndpointTypeUnspecified = 0
	EndpointTypeReqResp     = 1
	EndpointTypeWebSocket   = 2
)

var (
	EndpointTypeAllowed = map[EndpointType]struct{}{
		EndpointTypeReqResp:   {},
		EndpointTypeWebSocket: {},
	}
)

func (t EndpointType) Validate() error {
	if _, isAllowed := EndpointTypeAllowed[t]; !isAllowed {
		return errors.Wrapf(ErrEndpointTypeUnknown, "failed to validate endpoint type %d", EndpointTypeAllowed[t])
	}

	return nil
}
