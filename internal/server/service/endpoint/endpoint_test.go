package endpoint

import (
	"context"
	"testing"

	"github.com/status-mok/server/internal/mok"
	"github.com/status-mok/server/internal/mok/mocks"
	endpointAPI "github.com/status-mok/server/pkg/endpoint-api"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func Test_endpointService_Create(t *testing.T) {
	ctx := context.Background()
	sampleReq := &endpointAPI.CreateRequest{
		ServerName: "sample",
		Type:       endpointAPI.EndpointType_ENDPOINT_TYPE_REQ_RESP,
	}

	type testCase struct {
		name        string
		req         *endpointAPI.CreateRequest
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

				ept := mok.NewEndpoint(tc.req.GetUrl(), mok.EndpointType(tc.req.GetType().Number()))

				serverMock.
					On("EndpointCreate", mock.Anything, ept).
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
			name: "error: endpoint already exist",
			req:  sampleReq,
			setupMocks: func(t *testing.T, tc *testCase) *mocks.ServerStorageMock {
				serverStorageMock, serverMock := mocks.NewServerStorageMock(t), mocks.NewServerMock(t)

				serverStorageMock.
					On("ServerGet", mock.Anything, tc.req.GetServerName()).
					Return(serverMock, nil).
					Once()

				ept := mok.NewEndpoint(tc.req.GetUrl(), mok.EndpointType(tc.req.GetType().Number()))

				serverMock.
					On("EndpointCreate", mock.Anything, ept).
					Return(mok.ErrAlreadyExist).
					Once()

				return serverStorageMock
			},
			expErrorMsg: mok.ErrAlreadyExist.Error(),
		},
		{
			name: "error: invalid entrypoint type",
			req: &endpointAPI.CreateRequest{
				ServerName: "sample",
				Type:       endpointAPI.EndpointType_ENDPOINT_TYPE_UNSPECIFIED,
			},
			setupMocks: func(t *testing.T, tc *testCase) *mocks.ServerStorageMock {
				return mocks.NewServerStorageMock(t)
			},
			expErrorMsg: mok.ErrEndpointTypeUnknown.Error(),
		},
	}

	t.Parallel()

	for _, tc := range testCases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			serverStorageMock := tc.setupMocks(t, &tc)

			svc := NewEndpointService(serverStorageMock)

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

func Test_endpointService_Delete(t *testing.T) {
	ctx := context.Background()
	sampleReq := &endpointAPI.DeleteRequest{ServerName: "sample", Url: "/sample-url"}

	type testCase struct {
		name        string
		req         *endpointAPI.DeleteRequest
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
					On("EndpointDelete", mock.Anything, tc.req.GetUrl()).
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
			name: "error: endpoint not found",
			req:  sampleReq,
			setupMocks: func(t *testing.T, tc *testCase) *mocks.ServerStorageMock {
				serverStorageMock, serverMock := mocks.NewServerStorageMock(t), mocks.NewServerMock(t)

				serverStorageMock.
					On("ServerGet", mock.Anything, tc.req.GetServerName()).
					Return(serverMock, nil).
					Once()

				serverMock.
					On("EndpointDelete", mock.Anything, tc.req.GetUrl()).
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

			svc := NewEndpointService(serverStorageMock)

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
