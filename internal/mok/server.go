package mok

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"sync"

	"github.com/status-mok/server/internal/pkg/errors"
	"github.com/status-mok/server/internal/pkg/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	ErrAlreadyRunning = status.Error(codes.FailedPrecondition, "already running")
	ErrAlreadyStopped = status.Error(codes.FailedPrecondition, "already stopped")
)

type mokServer struct {
	name  string     `mapstructure:"name"`
	ip    string     `mapstructure:"ip"`
	port  uint32     `mapstructure:"port"`
	_type ServerType `mapstructure:"type"`

	status ServerStatus

	listener   net.Listener
	httpServer *http.Server

	mu sync.Mutex
}

func NewServer(name, ip string, port uint32, serverType ServerType) *mokServer {
	return &mokServer{
		name:  name,
		ip:    ip,
		port:  port,
		_type: serverType,

		status: ServerStatusStopped,
	}
}

func (s *mokServer) Addr() string {
	if s.listener != nil {
		return s.listener.Addr().String()
	}

	return fmt.Sprintf("%s:%d", s.ip, s.port)
}

func (s *mokServer) Name() string {
	return s.name
}

func (s *mokServer) Status() ServerStatus {
	return s.status
}

func (s *mokServer) Start(ctx context.Context) (err error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if s.status == ServerStatusRunning {
		return ErrAlreadyRunning
	}

	defer func() {
		if err != nil {
			s.status = ServerStatusStopped
		} else {
			s.status = ServerStatusRunning
		}
	}()

	if s._type == ServerTypeHTTP {
		s.httpServer = &http.Server{
			Handler:  s.httpHandler(ctx),
			ErrorLog: log.StdLogger(ctx),
		}

		s.listener, err = net.Listen("tcp", s.Addr())
		if err != nil {
			return errors.Wrapf(err, "failed to start tcp listening to '%s'", s.Addr())
		}

		go func() {
			s.mu.Lock()
			s.status = ServerStatusRunning
			s.mu.Unlock()

			if errS := s.httpServer.Serve(s.listener); err != nil && err != http.ErrServerClosed {
				log.L(ctx).With("error", errS).Errorf("http server at '%s' stopped with error: '%s'", s.Addr(), errS)
			}

			s.mu.Lock()
			s.status = ServerStatusStopped
			s.mu.Unlock()
		}()
	} else {

	}

	return nil
}

func (s *mokServer) Stop(ctx context.Context) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if s.status == ServerStatusStopped {
		return ErrAlreadyStopped
	}

	defer func() {
		s.status = ServerStatusStopped
	}()

	if s._type == ServerTypeHTTP {
		return s.httpServer.Shutdown(ctx)
	} else {

	}

	return nil
}
