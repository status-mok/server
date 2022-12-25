package mok

import (
	"github.com/status-mok/server/internal/pkg/errors"
)

var ErrServerTypeUnknown = errors.New("unknown server type")

type ServerType int32

const (
	ServerTypeUnspecified = 0
	ServerTypeHTTP        = 1
	ServerTypeGRPC        = 2
	ServerTypeThrift      = 3
	ServerTypeTCP         = 4
	ServerTypeUDP         = 5

	ServerTypeHTTPString   = "http"
	ServerTypeGRPCString   = "grpc"
	ServerTypeThriftString = "thrift"
	ServerTypeTCPString    = "tcp"
	ServerTypeUDPString    = "udp"
)

var (
	ServerTypesAllowed = map[ServerType]struct{}{
		ServerTypeHTTP:   {},
		ServerTypeGRPC:   {},
		ServerTypeThrift: {},
		ServerTypeTCP:    {},
		ServerTypeUDP:    {},
	}

	StringToServerType = map[string]ServerType{
		ServerTypeHTTPString:   ServerTypeHTTP,
		ServerTypeGRPCString:   ServerTypeGRPC,
		ServerTypeThriftString: ServerTypeThrift,
		ServerTypeTCPString:    ServerTypeTCP,
		ServerTypeUDPString:    ServerTypeUDP,
	}

	ServerTypeToString = map[ServerType]string{
		ServerTypeHTTP:   ServerTypeHTTPString,
		ServerTypeGRPC:   ServerTypeGRPCString,
		ServerTypeThrift: ServerTypeThriftString,
		ServerTypeTCP:    ServerTypeTCPString,
		ServerTypeUDP:    ServerTypeUDPString,
	}
)

func (t ServerType) Validate() error {
	if _, isAllowed := ServerTypesAllowed[t]; !isAllowed {
		return errors.Wrapf(ErrServerTypeUnknown, "failed to validate server type %d", t)
	}

	return nil
}
