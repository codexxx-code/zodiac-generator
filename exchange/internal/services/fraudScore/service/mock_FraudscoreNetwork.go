// Code generated by mockery v2.46.2. DO NOT EDIT.

package service

import (
	context "context"
	model "exchange/internal/services/fraudScore/model"

	mock "github.com/stretchr/testify/mock"
)

// MockFraudscoreNetwork is an autogenerated mock type for the FraudscoreNetwork type
type MockFraudscoreNetwork struct {
	mock.Mock
}

// CheckFraudScore provides a mock function with given fields: _a0, req
func (_m *MockFraudscoreNetwork) CheckFraudScore(_a0 context.Context, req model.IsFraudReq) (bool, error) {
	ret := _m.Called(_a0, req)

	if len(ret) == 0 {
		panic("no return value specified for CheckFraudScore")
	}

	var r0 bool
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, model.IsFraudReq) (bool, error)); ok {
		return rf(_a0, req)
	}
	if rf, ok := ret.Get(0).(func(context.Context, model.IsFraudReq) bool); ok {
		r0 = rf(_a0, req)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(context.Context, model.IsFraudReq) error); ok {
		r1 = rf(_a0, req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewMockFraudscoreNetwork creates a new instance of MockFraudscoreNetwork. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockFraudscoreNetwork(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockFraudscoreNetwork {
	mock := &MockFraudscoreNetwork{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}