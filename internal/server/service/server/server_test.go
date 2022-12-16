package server

import (
	"context"
	"testing"

	"github.com/status-mok/server/internal/mok"
	"github.com/status-mok/server/internal/mok/mocks"
	serverAPI "github.com/status-mok/server/pkg/server-api"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func Test_serverService_Create(t *testing.T) {
	ctx := context.Background()
	sampleReq := &serverAPI.CreateRequest{
		Name: "sample",
		Type: serverAPI.ServerType_SERVER_TYPE_HTTP,
	}

	type testCase struct {
		name        string
		req         *serverAPI.CreateRequest
		setupMocks  func(t *testing.T, tc *testCase) *mocks.ServerStorageMock
		expErrorMsg string
	}

	testCases := []testCase{
		{
			name: "ok",
			req:  sampleReq,
			setupMocks: func(t *testing.T, tc *testCase) *mocks.ServerStorageMock {
				serverStorageMock := mocks.NewServerStorageMock(t)

				serverStorageMock.
					On(
						"ServerCreate",
						mock.Anything,
						mok.NewServer(
							tc.req.GetName(),
							tc.req.GetIp(),
							tc.req.GetPort(),
							mok.ServerType(tc.req.GetType().Number()),
						),
					).
					Return(nil).
					Once()

				return serverStorageMock
			},
		},
		{
			name: "error: server already exist",
			req:  sampleReq,
			setupMocks: func(t *testing.T, tc *testCase) *mocks.ServerStorageMock {
				serverStorageMock := mocks.NewServerStorageMock(t)

				serverStorageMock.
					On(
						"ServerCreate",
						mock.Anything,
						mok.NewServer(
							tc.req.GetName(),
							tc.req.GetIp(),
							tc.req.GetPort(),
							mok.ServerType(tc.req.GetType().Number()),
						),
					).
					Return(mok.ErrAlreadyExist).
					Once()

				return serverStorageMock
			},
			expErrorMsg: mok.ErrAlreadyExist.Error(),
		},
		{
			name: "error: invalid server type",
			req: &serverAPI.CreateRequest{
				Name: "sample",
				Type: serverAPI.ServerType_SERVER_TYPE_UNSPECIFIED,
			},
			setupMocks: func(t *testing.T, tc *testCase) *mocks.ServerStorageMock {
				return mocks.NewServerStorageMock(t)
			},
			expErrorMsg: mok.ErrServerTypeUnknown.Error(),
		},
	}

	t.Parallel()

	for _, tc := range testCases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			serverStorageMock := tc.setupMocks(t, &tc)

			svc := NewServerService(serverStorageMock)

			resp, err := svc.Create(ctx, tc.req)
			if len(tc.expErrorMsg) > 0 {
				require.Error(t, err)
				assert.ErrorContains(t, err, tc.expErrorMsg)
			} else {
				require.NoError(t, err)
				require.NotNil(t, resp)
				require.True(t, resp.GetSuccess())
			}
		})
	}
}

func Test_serverService_Delete(t *testing.T) {
	ctx := context.Background()
	sampleReq := &serverAPI.DeleteRequest{Name: "sample"}

	type testCase struct {
		name        string
		req         *serverAPI.DeleteRequest
		setupMocks  func(t *testing.T, tc *testCase) *mocks.ServerStorageMock
		expErrorMsg string
	}

	testCases := []testCase{
		{
			name: "ok",
			req:  sampleReq,
			setupMocks: func(t *testing.T, tc *testCase) *mocks.ServerStorageMock {
				serverStorageMock := mocks.NewServerStorageMock(t)

				serverStorageMock.
					On("ServerDelete", mock.Anything, tc.req.GetName()).
					Return(nil).
					Once()

				return serverStorageMock
			},
		},
		{
			name: "error: server not found",
			req:  sampleReq,
			setupMocks: func(t *testing.T, tc *testCase) *mocks.ServerStorageMock {
				serverStorageMock := mocks.NewServerStorageMock(t)

				serverStorageMock.
					On("ServerDelete", mock.Anything, tc.req.GetName()).
					Return(mok.ErrNotFound).
					Once()

				return serverStorageMock
			},
			expErrorMsg: mok.ErrNotFound.Error(),
		},
	}

	t.Parallel()

	for _, tc := range testCases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			serverStorageMock := tc.setupMocks(t, &tc)

			svc := NewServerService(serverStorageMock)

			resp, err := svc.Delete(ctx, tc.req)
			if len(tc.expErrorMsg) > 0 {
				require.Error(t, err)
				assert.ErrorContains(t, err, tc.expErrorMsg)
			} else {
				require.NoError(t, err)
				require.NotNil(t, resp)
				require.True(t, resp.GetSuccess())
			}
		})
	}
}

func Test_serverService_Start(t *testing.T) {
	ctx := context.Background()
	sampleReq := &serverAPI.StartRequest{Name: "sample"}
	sampleAddr := "0.0.0.0:8080"

	type testCase struct {
		name        string
		req         *serverAPI.StartRequest
		setupMocks  func(t *testing.T, tc *testCase) *mocks.ServerStorageMock
		expErrorMsg string
	}

	testCases := []testCase{
		{
			name: "ok",
			req:  sampleReq,
			setupMocks: func(t *testing.T, tc *testCase) *mocks.ServerStorageMock {
				serverStorageMock, serverMock := mocks.NewServerStorageMock(t), mocks.NewServerMock(t)

				serverStorageMock.
					On("ServerGet", mock.Anything, tc.req.GetName()).
					Return(serverMock, nil).
					Once()

				serverMock.On("Start", mock.Anything).Return(nil).Once()
				serverMock.On("Addr", mock.Anything).Return(sampleAddr).Once()

				return serverStorageMock
			},
		},
		{
			name: "error: server not found",
			req:  sampleReq,
			setupMocks: func(t *testing.T, tc *testCase) *mocks.ServerStorageMock {
				serverStorageMock := mocks.NewServerStorageMock(t)

				serverStorageMock.
					On("ServerGet", mock.Anything, tc.req.GetName()).
					Return(nil, mok.ErrNotFound).
					Once()

				return serverStorageMock
			},
			expErrorMsg: mok.ErrNotFound.Error(),
		},
		{
			name: "error: server already running",
			req:  sampleReq,
			setupMocks: func(t *testing.T, tc *testCase) *mocks.ServerStorageMock {
				serverStorageMock, serverMock := mocks.NewServerStorageMock(t), mocks.NewServerMock(t)

				serverStorageMock.
					On("ServerGet", mock.Anything, tc.req.GetName()).
					Return(serverMock, nil).
					Once()

				serverMock.On("Start", mock.Anything).Return(mok.ErrAlreadyRunning).Once()

				return serverStorageMock
			},
			expErrorMsg: mok.ErrAlreadyRunning.Error(),
		},
	}

	t.Parallel()

	for _, tc := range testCases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			serverStorageMock := tc.setupMocks(t, &tc)

			svc := NewServerService(serverStorageMock)

			resp, err := svc.Start(ctx, tc.req)
			if len(tc.expErrorMsg) > 0 {
				require.Error(t, err)
				assert.ErrorContains(t, err, tc.expErrorMsg)
			} else {
				require.NoError(t, err)
				require.NotNil(t, resp)
				require.True(t, resp.GetSuccess())
				require.Equal(t, sampleAddr, resp.GetAddress())
			}
		})
	}
}

func Test_serverService_Stop(t *testing.T) {
	ctx := context.Background()
	sampleReq := &serverAPI.StopRequest{Name: "sample"}

	type testCase struct {
		name        string
		req         *serverAPI.StopRequest
		setupMocks  func(t *testing.T, tc *testCase) *mocks.ServerStorageMock
		expErrorMsg string
	}

	testCases := []testCase{
		{
			name: "ok",
			req:  sampleReq,
			setupMocks: func(t *testing.T, tc *testCase) *mocks.ServerStorageMock {
				serverStorageMock, serverMock := mocks.NewServerStorageMock(t), mocks.NewServerMock(t)

				serverStorageMock.
					On("ServerGet", mock.Anything, tc.req.GetName()).
					Return(serverMock, nil).
					Once()

				serverMock.On("Stop", mock.Anything).Return(nil).Once()

				return serverStorageMock
			},
		},
		{
			name: "error: server not found",
			req:  sampleReq,
			setupMocks: func(t *testing.T, tc *testCase) *mocks.ServerStorageMock {
				serverStorageMock := mocks.NewServerStorageMock(t)

				serverStorageMock.
					On("ServerGet", mock.Anything, tc.req.GetName()).
					Return(nil, mok.ErrNotFound).
					Once()

				return serverStorageMock
			},
			expErrorMsg: mok.ErrNotFound.Error(),
		},
		{
			name: "error: server already stopped",
			req:  sampleReq,
			setupMocks: func(t *testing.T, tc *testCase) *mocks.ServerStorageMock {
				serverStorageMock, serverMock := mocks.NewServerStorageMock(t), mocks.NewServerMock(t)

				serverStorageMock.
					On("ServerGet", mock.Anything, tc.req.GetName()).
					Return(serverMock, nil).
					Once()

				serverMock.On("Stop", mock.Anything).Return(mok.ErrAlreadyStopped).Once()

				return serverStorageMock
			},
			expErrorMsg: mok.ErrAlreadyStopped.Error(),
		},
	}

	t.Parallel()

	for _, tc := range testCases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			serverStorageMock := tc.setupMocks(t, &tc)

			svc := NewServerService(serverStorageMock)

			resp, err := svc.Stop(ctx, tc.req)
			if len(tc.expErrorMsg) > 0 {
				require.Error(t, err)
				assert.ErrorContains(t, err, tc.expErrorMsg)
			} else {
				require.NoError(t, err)
				require.NotNil(t, resp)
				require.True(t, resp.GetSuccess())
			}
		})
	}
}
