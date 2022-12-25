package mok

import (
	"github.com/status-mok/server/internal/pkg/errors"
)

var ErrRouteTypeUnknown = errors.New("unknown route type")

type RouteType int32

const (
	RouteTypeUnspecified = 0
	RouteTypeReqResp     = 1
	RouteTypeWebSocket   = 2

	RouteTypeReqRespString   = "req-resp"
	RouteTypeWebSocketString = "websocket"
)

var (
	RouteTypeAllowed = map[RouteType]struct{}{
		RouteTypeReqResp:   {},
		RouteTypeWebSocket: {},
	}

	StringToRouteType = map[string]RouteType{
		RouteTypeReqRespString:   RouteTypeReqResp,
		RouteTypeWebSocketString: RouteTypeWebSocket,
	}

	RouteTypeToString = map[RouteType]string{
		RouteTypeReqResp:   RouteTypeReqRespString,
		RouteTypeWebSocket: RouteTypeWebSocketString,
	}
)

func (t RouteType) Validate() error {
	if _, isAllowed := RouteTypeAllowed[t]; !isAllowed {
		return errors.Wrapf(ErrRouteTypeUnknown, "failed to validate route type %d", RouteTypeAllowed[t])
	}

	return nil
}
