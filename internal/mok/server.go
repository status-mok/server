package mok

import (
	"context"
	"fmt"
	"net"
	"net/http"

	"github.com/status-mok/server/internal/pkg/errors"
	"github.com/status-mok/server/internal/pkg/log"
)

type mokServer struct {
	name  string     `mapstructure:"name"`
	ip    string     `mapstructure:"ip"`
	port  uint32     `mapstructure:"port"`
	_type ServerType `mapstructure:"type"`

	listener   net.Listener
	httpServer *http.Server

	status ServerStatus
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
			Handler:  s.httpHandler(ctx),
			ErrorLog: log.StdLogger(ctx),
		}

		l, err := net.Listen("tcp", s.Addr())
		if err != nil {
			return errors.Wrapf(err, "failed to start listening to '%s'", s.Addr())
		}

		go func() {
			if errS := s.httpServer.Serve(l); err != nil && err != http.ErrServerClosed {
				log.L(ctx).With("error", errS).Errorf("http server at '%s' stopped with error: '%s'", s.Addr(), errS)
			}
		}()

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
