package mok

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"testing"

	"github.com/status-mok/server/internal/pkg/tester"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_server_Start(t *testing.T) {
	ctx := context.Background()
	portInUse := tester.GetFreePort()
	_, err := net.Listen("tcp", fmt.Sprintf(":%d", portInUse))
	require.NoError(t, err)

	type testCase struct {
		name        string
		server      *server
		beforeTest  func(t *testing.T, tc *testCase)
		expErrorMsg string
	}

	testCases := []testCase{
		{
			name:   "ok: specified port",
			server: NewServer("", "", tester.GetFreePort(), ServerTypeHTTP),
		},
		{
			name:   "ok: random port",
			server: NewServer("", "", 0, ServerTypeHTTP),
		},
		{
			name:        "error: port in use",
			server:      NewServer("", "", portInUse, ServerTypeHTTP),
			expErrorMsg: "address already in use",
		},
		{
			name:   "error: already running",
			server: NewServer("", "", 0, ServerTypeHTTP),
			beforeTest: func(t *testing.T, tc *testCase) {
				require.NoError(t, tc.server.Start(ctx))
			},
			expErrorMsg: ErrAlreadyRunning.Error(),
		},
	}

	t.Parallel()

	for _, tc := range testCases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			defer tc.server.Stop(ctx)

			if tc.beforeTest != nil {
				tc.beforeTest(t, &tc)
			}

			err := tc.server.Start(ctx)
			if len(tc.expErrorMsg) > 0 {
				require.Error(t, err)
				assert.ErrorContains(t, err, tc.expErrorMsg)
			} else {
				require.NoError(t, err)

				serverResp, err := http.Get("http://" + tc.server.Addr())
				require.NoError(t, err)

				assert.Equal(t, http.StatusOK, serverResp.StatusCode)
			}
		})
	}
}

func Test_server_Stop(t *testing.T) {
	ctx := context.Background()

	type testCase struct {
		name        string
		server      *server
		beforeTest  func(t *testing.T, tc *testCase)
		expErrorMsg string
	}

	testCases := []testCase{
		{
			name:   "ok",
			server: NewServer("", "", 0, ServerTypeHTTP),
			beforeTest: func(t *testing.T, tc *testCase) {
				require.NoError(t, tc.server.Start(ctx))
			},
		},
		{
			name:        "error: initially stopped server",
			server:      NewServer("", "", tester.GetFreePort(), ServerTypeHTTP),
			expErrorMsg: ErrAlreadyStopped.Error(),
		},
		{
			name:   "error: stopped server",
			server: NewServer("", "", 0, ServerTypeHTTP),
			beforeTest: func(t *testing.T, tc *testCase) {
				require.NoError(t, tc.server.Start(ctx))
				require.NoError(t, tc.server.Stop(ctx))
			},
			expErrorMsg: ErrAlreadyStopped.Error(),
		},
	}

	t.Parallel()

	for _, tc := range testCases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			if tc.beforeTest != nil {
				tc.beforeTest(t, &tc)
			}

			err := tc.server.Stop(ctx)
			if len(tc.expErrorMsg) > 0 {
				require.Error(t, err)
				assert.ErrorContains(t, err, tc.expErrorMsg)
			} else {
				require.NoError(t, err)
			}

			_, err = http.Get("http://" + tc.server.Addr())
			require.Error(t, err)
			require.ErrorContains(t, err, "connection refused")
		})
	}
}
