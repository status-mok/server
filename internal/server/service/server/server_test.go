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

	type testCase struct {
		name        string
		req         *serverAPI.CreateRequest
		setupMocks  func(t *testing.T, tc *testCase) *mocks.ServerStorageMock
		expErrorMsg string
	}

	testCases := []testCase{
		{
			name: "ok",
			req: &serverAPI.CreateRequest{
				Name: "sample",
				Type: 1,
			},
			setupMocks: func(t *testing.T, tc *testCase) *mocks.ServerStorageMock {
				storageMock := mocks.NewServerStorageMock(t)

				storageMock.
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

				return storageMock
			},
		},
		{
			name: "error",
			req: &serverAPI.CreateRequest{
				Name: "sample",
				Type: 1,
			},
			setupMocks: func(t *testing.T, tc *testCase) *mocks.ServerStorageMock {
				storageMock := mocks.NewServerStorageMock(t)

				storageMock.
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

				return storageMock
			},
			expErrorMsg: mok.ErrAlreadyExist.Error(),
		},
	}

	t.Parallel()

	for _, tc := range testCases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			storageMock := tc.setupMocks(t, &tc)

			svc := NewServerService(storageMock)

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

	type testCase struct {
		name        string
		req         *serverAPI.DeleteRequest
		setupMocks  func(t *testing.T, tc *testCase) *mocks.ServerStorageMock
		expErrorMsg string
	}

	testCases := []testCase{
		{
			name: "ok",
			req:  &serverAPI.DeleteRequest{Name: "sample"},
			setupMocks: func(t *testing.T, tc *testCase) *mocks.ServerStorageMock {
				storageMock := mocks.NewServerStorageMock(t)

				storageMock.
					On("ServerDelete", mock.Anything, tc.req.GetName()).
					Return(nil).
					Once()

				return storageMock
			},
		},
		{
			name: "error",
			req:  &serverAPI.DeleteRequest{Name: "sample"},
			setupMocks: func(t *testing.T, tc *testCase) *mocks.ServerStorageMock {
				storageMock := mocks.NewServerStorageMock(t)

				storageMock.
					On("ServerDelete", mock.Anything, tc.req.GetName()).
					Return(mok.ErrNotFound).
					Once()

				return storageMock
			},
			expErrorMsg: mok.ErrNotFound.Error(),
		},
	}

	t.Parallel()

	for _, tc := range testCases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			storageMock := tc.setupMocks(t, &tc)

			svc := NewServerService(storageMock)

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
			req:  &serverAPI.StartRequest{Name: "sample"},
			setupMocks: func(t *testing.T, tc *testCase) *mocks.ServerStorageMock {
				storageMock, serverMock := mocks.NewServerStorageMock(t), mocks.NewServerMock(t)

				storageMock.
					On("ServerGet", mock.Anything, tc.req.GetName()).
					Return(serverMock, nil).
					Once()

				serverMock.On("Start", mock.Anything).Return(nil).Once()
				serverMock.On("Addr", mock.Anything).Return(sampleAddr).Once()

				return storageMock
			},
		},
		{
			name: "error: server not found",
			req:  &serverAPI.StartRequest{Name: "sample"},
			setupMocks: func(t *testing.T, tc *testCase) *mocks.ServerStorageMock {
				storageMock := mocks.NewServerStorageMock(t)

				storageMock.
					On("ServerGet", mock.Anything, tc.req.GetName()).
					Return(nil, mok.ErrNotFound).
					Once()

				return storageMock
			},
			expErrorMsg: mok.ErrNotFound.Error(),
		},
		{
			name: "error: server start error",
			req:  &serverAPI.StartRequest{Name: "sample"},
			setupMocks: func(t *testing.T, tc *testCase) *mocks.ServerStorageMock {
				storageMock, serverMock := mocks.NewServerStorageMock(t), mocks.NewServerMock(t)

				storageMock.
					On("ServerGet", mock.Anything, tc.req.GetName()).
					Return(serverMock, nil).
					Once()

				serverMock.On("Start", mock.Anything).Return(mok.ErrAlreadyRunning).Once()

				return storageMock
			},
			expErrorMsg: mok.ErrAlreadyRunning.Error(),
		},
	}

	t.Parallel()

	for _, tc := range testCases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			storageMock := tc.setupMocks(t, &tc)

			svc := NewServerService(storageMock)

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

	type testCase struct {
		name        string
		req         *serverAPI.StopRequest
		setupMocks  func(t *testing.T, tc *testCase) *mocks.ServerStorageMock
		expErrorMsg string
	}

	testCases := []testCase{
		{
			name: "ok",
			req:  &serverAPI.StopRequest{Name: "sample"},
			setupMocks: func(t *testing.T, tc *testCase) *mocks.ServerStorageMock {
				storageMock, serverMock := mocks.NewServerStorageMock(t), mocks.NewServerMock(t)

				storageMock.
					On("ServerGet", mock.Anything, tc.req.GetName()).
					Return(serverMock, nil).
					Once()

				serverMock.On("Stop", mock.Anything).Return(nil).Once()

				return storageMock
			},
		},
		{
			name: "error: server not found",
			req:  &serverAPI.StopRequest{Name: "sample"},
			setupMocks: func(t *testing.T, tc *testCase) *mocks.ServerStorageMock {
				storageMock := mocks.NewServerStorageMock(t)

				storageMock.
					On("ServerGet", mock.Anything, tc.req.GetName()).
					Return(nil, mok.ErrNotFound).
					Once()

				return storageMock
			},
			expErrorMsg: mok.ErrNotFound.Error(),
		},
		{
			name: "error: server stop error",
			req:  &serverAPI.StopRequest{Name: "sample"},
			setupMocks: func(t *testing.T, tc *testCase) *mocks.ServerStorageMock {
				storageMock, serverMock := mocks.NewServerStorageMock(t), mocks.NewServerMock(t)

				storageMock.
					On("ServerGet", mock.Anything, tc.req.GetName()).
					Return(serverMock, nil).
					Once()

				serverMock.On("Stop", mock.Anything).Return(mok.ErrAlreadyStopped).Once()

				return storageMock
			},
			expErrorMsg: mok.ErrAlreadyStopped.Error(),
		},
	}

	t.Parallel()

	for _, tc := range testCases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			storageMock := tc.setupMocks(t, &tc)

			svc := NewServerService(storageMock)

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
