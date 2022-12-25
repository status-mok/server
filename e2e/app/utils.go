//go:build e2e

package app

import (
	"context"

	"github.com/status-mok/server/internal/pkg/errors"
	serverAPI "github.com/status-mok/server/pkg/server-api"
)

func (s *TestAppServer) CreateRunningServers(ctx context.Context, names ...string) (map[string]string, error) {
	return s.createServers(ctx, true, names...)
}

func (s *TestAppServer) CreateStoppedServers(ctx context.Context, names ...string) error {
	_, err := s.createServers(ctx, false, names...)
	return err
}

func (s *TestAppServer) createServers(ctx context.Context, start bool, names ...string) (addrMap map[string]string, err error) {
	addrMap = make(map[string]string)

	for _, name := range names {
		_, err := s.GRPCClient().ServerService().Create(ctx, &serverAPI.CreateRequest{
			Name: name,
			Type: serverAPI.ServerType_SERVER_TYPE_HTTP,
		})
		if err != nil {
			return nil, errors.Wrapf(err, "failed to create server '%s'", name)
		}

		if start {
			resp, err := s.GRPCClient().ServerService().Start(ctx, &serverAPI.StartRequest{
				Name: name,
			})
			if err != nil {
				return nil, errors.Wrapf(err, "failed to start server '%s'", name)
			}
			addrMap[name] = resp.GetAddress()
		}
	}

	return addrMap, nil
}
