package mok

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_endpointStorage_ServerGet(t *testing.T) {
	ctx := context.Background()
	sampleEndpoint := NewEndpoint("sample")

	testCases := []struct {
		name        string
		endpointURL string
		storage     *endpointStorage
		expErrorMsg string
	}{
		{
			name:        "ok",
			endpointURL: sampleEndpoint.URL(),
			storage:     testEndpointStorage(10, sampleEndpoint),
		},
		{
			name:        "error: not found",
			endpointURL: sampleEndpoint.URL(),
			storage:     testEndpointStorage(10),
			expErrorMsg: ErrNotFound.Error(),
		},
	}

	t.Parallel()

	for _, tc := range testCases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			ept, err := tc.storage.EndpointGet(ctx, tc.endpointURL)
			if len(tc.expErrorMsg) > 0 {
				require.Error(t, err)
				assert.ErrorContains(t, err, tc.expErrorMsg)
			} else {
				require.NoError(t, err)
				require.NotNil(t, ept)
				require.Equal(t, tc.endpointURL, ept.URL())
			}
		})
	}
}

func Test_endpointStorage_ServerCreate(t *testing.T) {
	ctx := context.Background()
	sampleEndpoint := NewEndpoint("sample")

	testCases := []struct {
		name        string
		endpoint    *endpoint
		storage     *endpointStorage
		expErrorMsg string
	}{
		{
			name:     "ok",
			endpoint: sampleEndpoint,
			storage:  testEndpointStorage(10),
		},
		{
			name:        "fail: already exist",
			endpoint:    sampleEndpoint,
			storage:     testEndpointStorage(10, sampleEndpoint),
			expErrorMsg: ErrAlreadyExist.Error(),
		},
	}

	t.Parallel()

	for _, tc := range testCases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			countServersBefore := len(tc.storage.storage)

			err := tc.storage.EndpointCreate(ctx, tc.endpoint)
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

func Test_endpointStorage_ServerDelete(t *testing.T) {
	ctx := context.Background()
	sampleEndpoint := NewEndpoint("sample")

	testCases := []struct {
		name        string
		endpointURL string
		storage     *endpointStorage
		expErrorMsg string
	}{
		{
			name:        "ok",
			endpointURL: sampleEndpoint.URL(),
			storage:     testEndpointStorage(10, sampleEndpoint),
		},
		{
			name:        "fail: not found",
			endpointURL: sampleEndpoint.URL(),
			storage:     testEndpointStorage(10),
			expErrorMsg: ErrNotFound.Error(),
		},
	}

	t.Parallel()

	for _, tc := range testCases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			countServersBefore := len(tc.storage.storage)

			err := tc.storage.EndpointDelete(ctx, tc.endpointURL)
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

func testEndpointStorage(n int, endpoints ...Endpoint) *endpointStorage {
	s := NewEndpointStorage(endpoints...)

	for i := 0; i < n; i++ {
		url := fmt.Sprint(i)
		s.storage[url] = NewEndpoint(url)
	}

	return s
}
