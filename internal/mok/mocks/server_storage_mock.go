// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	context "context"

	mok "github.com/status-mok/server/internal/mok"
	mock "github.com/stretchr/testify/mock"
)

// ServerStorageMock is an autogenerated mock type for the ServerStorageMock type
type ServerStorageMock struct {
	mock.Mock
}

// ServerCreate provides a mock function with given fields: ctx, srv
func (_m *ServerStorageMock) ServerCreate(ctx context.Context, srv mok.Server) error {
	ret := _m.Called(ctx, srv)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, mok.Server) error); ok {
		r0 = rf(ctx, srv)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ServerDelete provides a mock function with given fields: ctx, name
func (_m *ServerStorageMock) ServerDelete(ctx context.Context, name string) error {
	ret := _m.Called(ctx, name)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, name)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ServerGet provides a mock function with given fields: ctx, name
func (_m *ServerStorageMock) ServerGet(ctx context.Context, name string) (mok.Server, error) {
	ret := _m.Called(ctx, name)

	var r0 mok.Server
	if rf, ok := ret.Get(0).(func(context.Context, string) mok.Server); ok {
		r0 = rf(ctx, name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(mok.Server)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewServerStorageMock interface {
	mock.TestingT
	Cleanup(func())
}

// NewServerStorageMock creates a new instance of ServerStorageMock. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewServerStorageMock(t mockConstructorTestingTNewServerStorageMock) *ServerStorageMock {
	mock := &ServerStorageMock{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}