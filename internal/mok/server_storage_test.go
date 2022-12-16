package mok

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_serverStorage_ServerGet(t *testing.T) {
	ctx := context.Background()
	sampleServer := NewServer("sample", "", 0, ServerTypeHTTP)

	testCases := []struct {
		name        string
		serverName  string
		storage     *serverStorage
		expErrorMsg string
	}{
		{
			name:       "ok",
			serverName: sampleServer.Name(),
			storage:    testServerStorage(10, sampleServer),
		},
		{
			name:        "error: not found",
			serverName:  sampleServer.Name(),
			storage:     testServerStorage(10),
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

func Test_serverStorage_ServerCreate(t *testing.T) {
	ctx := context.Background()
	sampleServer := NewServer("sample", "", 0, ServerTypeHTTP)

	testCases := []struct {
		name        string
		server      *server
		storage     *serverStorage
		expErrorMsg string
	}{
		{
			name:    "ok",
			server:  sampleServer,
			storage: testServerStorage(10),
		},
		{
			name:        "fail: already exist",
			server:      sampleServer,
			storage:     testServerStorage(10, sampleServer),
			expErrorMsg: ErrAlreadyExist.Error(),
		},
	}

	t.Parallel()

	for _, tc := range testCases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			countServersBefore := len(tc.storage.storage)

			err := tc.storage.ServerCreate(ctx, tc.server)
			if len(tc.expErrorMsg) > 0 {
				require.Error(t, err)
				assert.ErrorContains(t, err, tc.expErrorMsg)
				require.Equal(t, countServersBefore, len(tc.storage.storage))
			} else {
				require.NoError(t, err)
				require.Equal(t, countServersBefore+1, len(tc.storage.storage))
			}
		})
	}
}

func Test_serverStorage_ServerDelete(t *testing.T) {
	ctx := context.Background()
	sampleServer := NewServer("sample", "", 0, ServerTypeHTTP)

	testCases := []struct {
		name        string
		serverName  string
		storage     *serverStorage
		expErrorMsg string
	}{
		{
			name:       "ok",
			serverName: sampleServer.Name(),
			storage:    testServerStorage(10, sampleServer),
		},
		{
			name:        "fail: not found",
			serverName:  sampleServer.Name(),
			storage:     testServerStorage(10),
			expErrorMsg: ErrNotFound.Error(),
		},
	}

	t.Parallel()

	for _, tc := range testCases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			countServersBefore := len(tc.storage.storage)

			err := tc.storage.ServerDelete(ctx, tc.serverName)
			if len(tc.expErrorMsg) > 0 {
				require.Error(t, err)
				assert.ErrorContains(t, err, tc.expErrorMsg)
				require.Equal(t, countServersBefore, len(tc.storage.storage))
			} else {
				require.NoError(t, err)
				require.Equal(t, countServersBefore-1, len(tc.storage.storage))
			}
		})
	}
}

func testServerStorage(n int, servers ...Server) *serverStorage {
	s := NewServerStorage(servers...)

	for i := 0; i < n; i++ {
		name := fmt.Sprint(i)
		s.storage[name] = NewServer(name, "", 0, ServerTypeHTTP)
	}

	return s
}
