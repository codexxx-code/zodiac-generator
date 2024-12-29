// Code generated by mockery v2.46.2. DO NOT EDIT.

package service

import (
	context "context"
	model "exchange/internal/services/fraudScore/model"

	mock "github.com/stretchr/testify/mock"
)

// MockFraudScoreService is an autogenerated mock type for the FraudScoreService type
type MockFraudScoreService struct {
	mock.Mock
}

// IsFraud provides a mock function with given fields: ctx, req
func (_m *MockFraudScoreService) IsFraud(ctx context.Context, req model.IsFraudReq) (bool, error) {
	ret := _m.Called(ctx, req)

	if len(ret) == 0 {
		panic("no return value specified for IsFraud")
	}

	var r0 bool
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, model.IsFraudReq) (bool, error)); ok {
		return rf(ctx, req)
	}
	if rf, ok := ret.Get(0).(func(context.Context, model.IsFraudReq) bool); ok {
		r0 = rf(ctx, req)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(context.Context, model.IsFraudReq) error); ok {
		r1 = rf(ctx, req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewMockFraudScoreService creates a new instance of MockFraudScoreService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockFraudScoreService(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockFraudScoreService {
	mock := &MockFraudScoreService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}