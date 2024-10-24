// Code generated by mockery v2.46.2. DO NOT EDIT.

package endpoint

import (
	context "context"
	model "server/internal/services/forecast/model"

	mock "github.com/stretchr/testify/mock"
)

// mockForecastService is an autogenerated mock type for the forecastService type
type mockForecastService struct {
	mock.Mock
}

// GetForecasts provides a mock function with given fields: _a0, _a1
func (_m *mockForecastService) GetForecasts(_a0 context.Context, _a1 model.GetForecastsReq) ([]model.Forecast, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for GetForecasts")
	}

	var r0 []model.Forecast
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, model.GetForecastsReq) ([]model.Forecast, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, model.GetForecastsReq) []model.Forecast); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.Forecast)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, model.GetForecastsReq) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// newMockForecastService creates a new instance of mockForecastService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func newMockForecastService(t interface {
	mock.TestingT
	Cleanup(func())
}) *mockForecastService {
	mock := &mockForecastService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
