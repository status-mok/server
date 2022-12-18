package expectation

import (
	"context"
	"testing"

	uuid "github.com/satori/go.uuid"
	"github.com/status-mok/server/internal/mok"
	"github.com/status-mok/server/internal/mok/mocks"
	expectationAPI "github.com/status-mok/server/pkg/expectation-api"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func Test_expectationService_Create(t *testing.T) {
	ctx := context.Background()
	id := uuid.NewV4().String()
	sampleReq := &expectationAPI.CreateRequest{ServerName: "sample", RouteUrl: "/sample-url", Id: &id}

	type testCase struct {
		name        string
		req         *expectationAPI.CreateRequest
		setupMocks  func(t *testing.T, tc *testCase) *mocks.ServerStorageMock
		expErrorMsg string
	}

	testCases := []testCase{
		{
			name: "ok",
			req:  sampleReq,
			setupMocks: func(t *testing.T, tc *testCase) *mocks.ServerStorageMock {
				serverStorageMock, serverMock, routeMock := mocks.NewServerStorageMock(t), mocks.NewServerMock(t), mocks.NewRouteMock(t)

				serverStorageMock.
					On("ServerGet", mock.Anything, tc.req.GetServerName()).
					Return(serverMock, nil).
					Once()

				serverMock.
					On("RouteGet", mock.Anything, tc.req.GetRouteUrl()).
					Return(routeMock, nil).
					Once()

				exp := mok.NewExpectation(tc.req.GetId())

				routeMock.
					On("ExpectationCreate", mock.Anything, exp).
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
					On("ServerGet", mock.Anything, tc.req.GetServerName()).
					Return(nil, mok.ErrNotFound).
					Once()

				return serverStorageMock
			},
			expErrorMsg: mok.ErrNotFound.Error(),
		},
		{
			name: "error: route not found",
			req:  sampleReq,
			setupMocks: func(t *testing.T, tc *testCase) *mocks.ServerStorageMock {
				serverStorageMock, serverMock := mocks.NewServerStorageMock(t), mocks.NewServerMock(t)

				serverStorageMock.
					On("ServerGet", mock.Anything, tc.req.GetServerName()).
					Return(serverMock, nil).
					Once()

				serverMock.
					On("RouteGet", mock.Anything, tc.req.GetRouteUrl()).
					Return(nil, mok.ErrNotFound).
					Once()

				return serverStorageMock
			},
			expErrorMsg: mok.ErrNotFound.Error(),
		},
		{
			name: "error: expectation already exist",
			req:  sampleReq,
			setupMocks: func(t *testing.T, tc *testCase) *mocks.ServerStorageMock {
				serverStorageMock, serverMock, routeMock := mocks.NewServerStorageMock(t), mocks.NewServerMock(t), mocks.NewRouteMock(t)

				serverStorageMock.
					On("ServerGet", mock.Anything, tc.req.GetServerName()).
					Return(serverMock, nil).
					Once()

				serverMock.
					On("RouteGet", mock.Anything, tc.req.GetRouteUrl()).
					Return(routeMock, nil).
					Once()

				exp := mok.NewExpectation(tc.req.GetId())

				routeMock.
					On("ExpectationCreate", mock.Anything, exp).
					Return(mok.ErrAlreadyExist).
					Once()

				return serverStorageMock
			},
			expErrorMsg: mok.ErrAlreadyExist.Error(),
		},
	}

	t.Parallel()

	for _, tc := range testCases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			serverStorageMock := tc.setupMocks(t, &tc)

			svc := NewExpectationService(serverStorageMock)

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

func Test_expectationService_Delete(t *testing.T) {
	ctx := context.Background()
	sampleReq := &expectationAPI.DeleteRequest{ServerName: "sample", RouteUrl: "/sample-url", Id: uuid.NewV4().String()}

	type testCase struct {
		name        string
		req         *expectationAPI.DeleteRequest
		setupMocks  func(t *testing.T, tc *testCase) *mocks.ServerStorageMock
		expErrorMsg string
	}

	testCases := []testCase{
		{
			name: "ok",
			req:  sampleReq,
			setupMocks: func(t *testing.T, tc *testCase) *mocks.ServerStorageMock {
				serverStorageMock, serverMock, routeMock := mocks.NewServerStorageMock(t), mocks.NewServerMock(t), mocks.NewRouteMock(t)

				serverStorageMock.
					On("ServerGet", mock.Anything, tc.req.GetServerName()).
					Return(serverMock, nil).
					Once()

				serverMock.
					On("RouteGet", mock.Anything, tc.req.GetRouteUrl()).
					Return(routeMock, nil).
					Once()

				routeMock.
					On("ExpectationDelete", mock.Anything, tc.req.GetId()).
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
					On("ServerGet", mock.Anything, tc.req.GetServerName()).
					Return(nil, mok.ErrNotFound).
					Once()

				return serverStorageMock
			},
			expErrorMsg: mok.ErrNotFound.Error(),
		},
		{
			name: "error: route not found",
			req:  sampleReq,
			setupMocks: func(t *testing.T, tc *testCase) *mocks.ServerStorageMock {
				serverStorageMock, serverMock := mocks.NewServerStorageMock(t), mocks.NewServerMock(t)

				serverStorageMock.
					On("ServerGet", mock.Anything, tc.req.GetServerName()).
					Return(serverMock, nil).
					Once()

				serverMock.
					On("RouteGet", mock.Anything, tc.req.GetRouteUrl()).
					Return(nil, mok.ErrNotFound).
					Once()

				return serverStorageMock
			},
			expErrorMsg: mok.ErrNotFound.Error(),
		},
		{
			name: "error: expectation not found",
			req:  sampleReq,
			setupMocks: func(t *testing.T, tc *testCase) *mocks.ServerStorageMock {
				serverStorageMock, serverMock, routeMock := mocks.NewServerStorageMock(t), mocks.NewServerMock(t), mocks.NewRouteMock(t)

				serverStorageMock.
					On("ServerGet", mock.Anything, tc.req.GetServerName()).
					Return(serverMock, nil).
					Once()

				serverMock.
					On("RouteGet", mock.Anything, tc.req.GetRouteUrl()).
					Return(routeMock, nil).
					Once()

				routeMock.
					On("ExpectationDelete", mock.Anything, tc.req.GetId()).
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

			svc := NewExpectationService(serverStorageMock)

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
