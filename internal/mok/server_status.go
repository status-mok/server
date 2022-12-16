package mok

import (
	"github.com/status-mok/server/internal/pkg/errors"
)

var ErrServerStatusUnknown = errors.New("unknown server status")

type ServerStatus int32

const (
	ServerStatusStopped = iota
	ServerStatusStarting
	ServerStatusRunning
)

var (
	ServerStatusAllowed = map[ServerStatus]struct{}{
		ServerStatusStopped:  {},
		ServerStatusStarting: {},
		ServerStatusRunning:  {},
	}
)

func (t ServerStatus) Validate() error {
	if _, isAllowed := ServerStatusAllowed[t]; !isAllowed {
		return errors.Wrapf(ErrServerStatusUnknown, "failed to validate server status %d", t)
	}

	return nil
}
