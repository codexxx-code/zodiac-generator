// Code generated by mockery v2.46.2. DO NOT EDIT.

package service

import (
	context "context"
	model "server/internal/services/forecast/model"

	mock "github.com/stretchr/testify/mock"
)

// MockForecastService is an autogenerated mock type for the ForecastService type
type MockForecastService struct {
	mock.Mock
}

// CreateForecast provides a mock function with given fields: ctx, req
func (_m *MockForecastService) CreateForecast(ctx context.Context, req model.CreateForecastReq) (uint32, error) {
	ret := _m.Called(ctx, req)

	if len(ret) == 0 {
		panic("no return value specified for CreateForecast")
	}

	var r0 uint32
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, model.CreateForecastReq) (uint32, error)); ok {
		return rf(ctx, req)
	}
	if rf, ok := ret.Get(0).(func(context.Context, model.CreateForecastReq) uint32); ok {
		r0 = rf(ctx, req)
	} else {
		r0 = ret.Get(0).(uint32)
	}

	if rf, ok := ret.Get(1).(func(context.Context, model.CreateForecastReq) error); ok {
		r1 = rf(ctx, req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewMockForecastService creates a new instance of MockForecastService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockForecastService(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockForecastService {
	mock := &MockForecastService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
