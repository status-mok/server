package route

import (
	"context"
	"testing"

	"github.com/status-mok/server/internal/mok"
	"github.com/status-mok/server/internal/mok/mocks"
	routeAPI "github.com/status-mok/server/pkg/route-api"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func Test_routeService_Create(t *testing.T) {
	ctx := context.Background()
	sampleReq := &routeAPI.CreateRequest{
		ServerName: "sample",
		Url:        "/sample-url",
		Type:       routeAPI.RouteType_ROUTE_TYPE_REQ_RESP,
	}

	type testCase struct {
		name        string
		req         *routeAPI.CreateRequest
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
					On("ServerGet", mock.Anything, tc.req.GetServerName()).
					Return(serverMock, nil).
					Once()

				ept := mok.NewRoute(tc.req.GetUrl(), mok.RouteType(tc.req.GetType().Number()))

				serverMock.
					On("RouteCreate", mock.Anything, ept).
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
			name: "error: route already exist",
			req:  sampleReq,
			setupMocks: func(t *testing.T, tc *testCase) *mocks.ServerStorageMock {
				serverStorageMock, serverMock := mocks.NewServerStorageMock(t), mocks.NewServerMock(t)

				serverStorageMock.
					On("ServerGet", mock.Anything, tc.req.GetServerName()).
					Return(serverMock, nil).
					Once()

				ept := mok.NewRoute(tc.req.GetUrl(), mok.RouteType(tc.req.GetType().Number()))

				serverMock.
					On("RouteCreate", mock.Anything, ept).
					Return(mok.ErrAlreadyExist).
					Once()

				return serverStorageMock
			},
			expErrorMsg: mok.ErrAlreadyExist.Error(),
		},
		{
			name: "error: invalid route type",
			req: &routeAPI.CreateRequest{
				ServerName: "sample",
				Url:        "/sample-url",
				Type:       routeAPI.RouteType_ROUTE_TYPE_UNSPECIFIED,
			},
			setupMocks: func(t *testing.T, tc *testCase) *mocks.ServerStorageMock {
				return mocks.NewServerStorageMock(t)
			},
			expErrorMsg: mok.ErrRouteTypeUnknown.Error(),
		},
	}

	t.Parallel()

	for _, tc := range testCases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			serverStorageMock := tc.setupMocks(t, &tc)

			svc := NewRouteService(serverStorageMock)

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

func Test_routeService_Delete(t *testing.T) {
	ctx := context.Background()
	sampleReq := &routeAPI.DeleteRequest{ServerName: "sample", Url: "/sample-url"}

	type testCase struct {
		name        string
		req         *routeAPI.DeleteRequest
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
					On("ServerGet", mock.Anything, tc.req.GetServerName()).
					Return(serverMock, nil).
					Once()

				serverMock.
					On("RouteDelete", mock.Anything, tc.req.GetUrl()).
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
					On("RouteDelete", mock.Anything, tc.req.GetUrl()).
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

			svc := NewRouteService(serverStorageMock)

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
