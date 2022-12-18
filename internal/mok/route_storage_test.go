package mok

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_routeStorage_ServerGet(t *testing.T) {
	ctx := context.Background()
	sampleRoute := NewRoute("sample", RouteTypeReqResp)

	testCases := []struct {
		name        string
		routeURL    string
		storage     *routeStorage
		expErrorMsg string
	}{
		{
			name:     "ok",
			routeURL: sampleRoute.URL(),
			storage:  testRouteStorage(10, sampleRoute),
		},
		{
			name:        "error: not found",
			routeURL:    sampleRoute.URL(),
			storage:     testRouteStorage(10),
			expErrorMsg: ErrNotFound.Error(),
		},
	}

	t.Parallel()

	for _, tc := range testCases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			ept, err := tc.storage.RouteGet(ctx, tc.routeURL)
			if len(tc.expErrorMsg) > 0 {
				require.Error(t, err)
				assert.ErrorContains(t, err, tc.expErrorMsg)
			} else {
				require.NoError(t, err)
				require.NotNil(t, ept)
				require.Equal(t, tc.routeURL, ept.URL())
			}
		})
	}
}

func Test_routeStorage_ServerCreate(t *testing.T) {
	ctx := context.Background()
	sampleRoute := NewRoute("sample", RouteTypeReqResp)

	testCases := []struct {
		name        string
		route       *route
		storage     *routeStorage
		expErrorMsg string
	}{
		{
			name:    "ok",
			route:   sampleRoute,
			storage: testRouteStorage(10),
		},
		{
			name:        "fail: already exist",
			route:       sampleRoute,
			storage:     testRouteStorage(10, sampleRoute),
			expErrorMsg: ErrAlreadyExist.Error(),
		},
	}

	t.Parallel()

	for _, tc := range testCases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			countServersBefore := len(tc.storage.storage)

			err := tc.storage.RouteCreate(ctx, tc.route)
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

func Test_routeStorage_ServerDelete(t *testing.T) {
	ctx := context.Background()
	sampleRoute := NewRoute("sample", RouteTypeReqResp)

	testCases := []struct {
		name        string
		routeURL    string
		storage     *routeStorage
		expErrorMsg string
	}{
		{
			name:     "ok",
			routeURL: sampleRoute.URL(),
			storage:  testRouteStorage(10, sampleRoute),
		},
		{
			name:        "fail: not found",
			routeURL:    sampleRoute.URL(),
			storage:     testRouteStorage(10),
			expErrorMsg: ErrNotFound.Error(),
		},
	}

	t.Parallel()

	for _, tc := range testCases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			countServersBefore := len(tc.storage.storage)

			err := tc.storage.RouteDelete(ctx, tc.routeURL)
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

func testRouteStorage(n int, routes ...Route) *routeStorage {
	s := NewRouteStorage(routes...)

	for i := 0; i < n; i++ {
		url := fmt.Sprint(i)
		s.storage[url] = NewRoute(url, RouteTypeReqResp)
	}

	return s
}
