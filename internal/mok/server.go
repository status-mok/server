package mok

import (
	"context"
	"fmt"
	"net"
	"net/http"

	"github.com/status-mok/server/internal/pkg/log"
)

type mokServer struct {
	name  string     `mapstructure:"name"`
	ip    string     `mapstructure:"ip"`
	port  uint32     `mapstructure:"port"`
	_type ServerType `mapstructure:"type"`

	listener   net.Listener
	httpServer *http.Server
}

func NewServer(name, ip string, port uint32, serverType ServerType) *mokServer {
	return &mokServer{
		name:  name,
		ip:    ip,
		port:  port,
		_type: serverType,
	}
}

func (s *mokServer) Addr() string {
	return fmt.Sprintf("%s:%d", s.ip, s.port)
}

func (s *mokServer) Name() string {
	return s.name
}

func (s *mokServer) Start(ctx context.Context) error {
	if s._type == ServerTypeHTTP {
		s.httpServer = &http.Server{
			Addr:     s.Addr(),
			Handler:  s.httpHandler(),
			ErrorLog: log.StdLogger(ctx),
		}

		if err := s.httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			return err
		}

		return nil
	} else {

	}

	return nil
}

func (s *mokServer) Stop(ctx context.Context) error {
	if s._type == ServerTypeHTTP {
		return s.httpServer.Shutdown(ctx)
	} else {

	}

	return nil
}
