package mok

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_expectationStorage_EndpointGet(t *testing.T) {
	ctx := context.Background()
	sampleExpectation := NewExpectation("sample")

	testCases := []struct {
		name          string
		expectationID string
		storage       *expectationStorage
		expErrorMsg   string
	}{
		{
			name:          "ok",
			expectationID: sampleExpectation.ID(),
			storage:       testExpectationStorage(10, sampleExpectation),
		},
		{
			name:          "error: not found",
			expectationID: sampleExpectation.ID(),
			storage:       testExpectationStorage(10),
			expErrorMsg:   ErrNotFound.Error(),
		},
	}

	t.Parallel()

	for _, tc := range testCases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			exp, err := tc.storage.ExpectationGet(ctx, tc.expectationID)
			if len(tc.expErrorMsg) > 0 {
				require.Error(t, err)
				assert.ErrorContains(t, err, tc.expErrorMsg)
			} else {
				require.NoError(t, err)
				require.NotNil(t, exp)
				require.Equal(t, tc.expectationID, exp.ID())
			}
		})
	}
}

func Test_expectationStorage_EndpointCreate(t *testing.T) {
	ctx := context.Background()
	sampleExpectation := NewExpectation("sample")

	testCases := []struct {
		name        string
		expectation *expectation
		storage     *expectationStorage
		expErrorMsg string
	}{
		{
			name:        "ok",
			expectation: sampleExpectation,
			storage:     testExpectationStorage(10),
		},
		{
			name:        "fail: already exist",
			expectation: sampleExpectation,
			storage:     testExpectationStorage(10, sampleExpectation),
			expErrorMsg: ErrAlreadyExist.Error(),
		},
	}

	t.Parallel()

	for _, tc := range testCases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			countExpectationsBefore := len(tc.storage.storage)

			err := tc.storage.ExpectationCreate(ctx, tc.expectation)
			if len(tc.expErrorMsg) > 0 {
				require.Error(t, err)
				assert.ErrorContains(t, err, tc.expErrorMsg)
				require.Equal(t, countExpectationsBefore, len(tc.storage.storage))
			} else {
				require.NoError(t, err)
				require.Equal(t, countExpectationsBefore+1, len(tc.storage.storage))
			}
		})
	}
}

func Test_expectationStorage_EndpointDelete(t *testing.T) {
	ctx := context.Background()
	sampleExpectation := NewExpectation("sample")

	testCases := []struct {
		name          string
		expectationID string
		storage       *expectationStorage
		expErrorMsg   string
	}{
		{
			name:          "ok",
			expectationID: sampleExpectation.ID(),
			storage:       testExpectationStorage(10, sampleExpectation),
		},
		{
			name:          "fail: not found",
			expectationID: sampleExpectation.ID(),
			storage:       testExpectationStorage(10),
			expErrorMsg:   ErrNotFound.Error(),
		},
	}

	t.Parallel()

	for _, tc := range testCases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			countExpectationsBefore := len(tc.storage.storage)

			err := tc.storage.ExpectationDelete(ctx, tc.expectationID)
			if len(tc.expErrorMsg) > 0 {
				require.Error(t, err)
				assert.ErrorContains(t, err, tc.expErrorMsg)
				require.Equal(t, countExpectationsBefore, len(tc.storage.storage))
			} else {
				require.NoError(t, err)
				require.Equal(t, countExpectationsBefore-1, len(tc.storage.storage))
			}
		})
	}
}

func Test_expectationStorage_ExpectationFindMatch(t *testing.T) {
	// TODO: add test cases
}

func testExpectationStorage(n int, expectations ...Expectation) *expectationStorage {
	s := NewExpectationStorage(expectations...)

	for i := 0; i < n; i++ {
		id := fmt.Sprint(i)
		s.storage[id] = NewExpectation(id)
	}

	return s
}
