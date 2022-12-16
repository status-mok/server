package mok

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestStorage_ServerGet(t *testing.T) {
	ctx := context.Background()

	sampleServer := NewServer("sample", "", 0, ServerTypeHTTP)
	nonEmptyStorage := func() *Storage {
		s := NewStorage()
		s.servers[sampleServer.Name()] = sampleServer
		return s
	}

	testCases := []struct {
		name        string
		serverName  string
		storage     *Storage
		expErrorMsg string
	}{
		{
			name:       "ok",
			serverName: sampleServer.Name(),
			storage:    nonEmptyStorage(),
		},
		{
			name:        "error: not found",
			serverName:  "qwerty",
			storage:     nonEmptyStorage(),
			expErrorMsg: ErrNotFound.Error(),
		},
		{
			name:        "error: not found",
			serverName:  "qwerty",
			storage:     NewStorage(),
			expErrorMsg: ErrNotFound.Error(),
		},
	}

	t.Parallel()

	for _, tc := range testCases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			server, err := tc.storage.ServerGet(ctx, tc.serverName)
			if len(tc.expErrorMsg) > 0 {
				require.Error(t, err)
				assert.ErrorContains(t, err, tc.expErrorMsg)
			} else {
				require.NoError(t, err)
				require.NotNil(t, server)
				require.Equal(t, tc.serverName, server.Name())
			}
		})
	}
}

func TestStorage_ServerCreate(t *testing.T) {
	ctx := context.Background()

	sampleServer := NewServer("sample", "", 0, ServerTypeHTTP)
	nonEmptyStorage := func(n int) *Storage {
		s := NewStorage()
		s.servers[sampleServer.Name()] = sampleServer
		for i := 1; i < n; i++ {
			s.servers[sampleServer.Name()+fmt.Sprint()] = NewServer("sample", "", 0, ServerTypeHTTP)
		}
		return s
	}

	testCases := []struct {
		name        string
		server      *mokServer
		storage     *Storage
		expErrorMsg string
	}{
		{
			name:    "ok",
			server:  sampleServer,
			storage: NewStorage(),
		},
		{
			name:        "fail: already exist",
			server:      sampleServer,
			storage:     nonEmptyStorage(10),
			expErrorMsg: ErrAlreadyExist.Error(),
		},
	}

	t.Parallel()

	for _, tc := range testCases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			countServersBefore := len(tc.storage.servers)

			err := tc.storage.ServerCreate(ctx, tc.server)
			if len(tc.expErrorMsg) > 0 {
				require.Error(t, err)
				assert.ErrorContains(t, err, tc.expErrorMsg)
				require.Equal(t, countServersBefore, len(tc.storage.servers))
			} else {
				require.NoError(t, err)
				require.Equal(t, countServersBefore+1, len(tc.storage.servers))
			}
		})
	}
}

func TestStorage_ServerDelete(t *testing.T) {
	ctx := context.Background()

	sampleServer := NewServer("sample", "", 0, ServerTypeHTTP)
	nonEmptyStorage := func(n int) *Storage {
		s := NewStorage()
		s.servers[sampleServer.Name()] = sampleServer
		for i := 1; i < n; i++ {
			s.servers[sampleServer.Name()+fmt.Sprint()] = NewServer("sample", "", 0, ServerTypeHTTP)
		}
		return s
	}

	testCases := []struct {
		name        string
		serverName  string
		storage     *Storage
		expErrorMsg string
	}{
		{
			name:       "ok",
			serverName: sampleServer.Name(),
			storage:    nonEmptyStorage(10),
		},
		{
			name:        "fail: not found",
			serverName:  sampleServer.Name(),
			storage:     NewStorage(),
			expErrorMsg: ErrNotFound.Error(),
		},
	}

	t.Parallel()

	for _, tc := range testCases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			countServersBefore := len(tc.storage.servers)

			err := tc.storage.ServerDelete(ctx, tc.serverName)
			if len(tc.expErrorMsg) > 0 {
				require.Error(t, err)
				assert.ErrorContains(t, err, tc.expErrorMsg)
				require.Equal(t, countServersBefore, len(tc.storage.servers))
			} else {
				require.NoError(t, err)
				require.Equal(t, countServersBefore-1, len(tc.storage.servers))
			}
		})
	}
}
