// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	context "context"

	mok "github.com/status-mok/server/internal/mok"
	mock "github.com/stretchr/testify/mock"
)

// RouteMock is an autogenerated mock type for the RouteMock type
type RouteMock struct {
	mock.Mock
}

// ExpectationCreate provides a mock function with given fields: ctx, exp
func (_m *RouteMock) ExpectationCreate(ctx context.Context, exp mok.Expectation) error {
	ret := _m.Called(ctx, exp)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, mok.Expectation) error); ok {
		r0 = rf(ctx, exp)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ExpectationDelete provides a mock function with given fields: ctx, id
func (_m *RouteMock) ExpectationDelete(ctx context.Context, id string) error {
	ret := _m.Called(ctx, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ExpectationFindMatch provides a mock function with given fields: ctx, routeType, req
func (_m *RouteMock) ExpectationFindMatch(ctx context.Context, routeType mok.RouteType, req interface{}) (mok.Expectation, error) {
	ret := _m.Called(ctx, routeType, req)

	var r0 mok.Expectation
	if rf, ok := ret.Get(0).(func(context.Context, mok.RouteType, interface{}) mok.Expectation); ok {
		r0 = rf(ctx, routeType, req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(mok.Expectation)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, mok.RouteType, interface{}) error); ok {
		r1 = rf(ctx, routeType, req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ExpectationGet provides a mock function with given fields: ctx, id
func (_m *RouteMock) ExpectationGet(ctx context.Context, id string) (mok.Expectation, error) {
	ret := _m.Called(ctx, id)

	var r0 mok.Expectation
	if rf, ok := ret.Get(0).(func(context.Context, string) mok.Expectation); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(mok.Expectation)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// URL provides a mock function with given fields:
func (_m *RouteMock) URL() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

type mockConstructorTestingTNewRouteMock interface {
	mock.TestingT
	Cleanup(func())
}

// NewRouteMock creates a new instance of RouteMock. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewRouteMock(t mockConstructorTestingTNewRouteMock) *RouteMock {
	mock := &RouteMock{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}