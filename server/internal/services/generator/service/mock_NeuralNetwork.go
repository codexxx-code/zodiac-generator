// Code generated by mockery v2.46.2. DO NOT EDIT.

package service

import (
	context "context"
	model "server/internal/services/generator/model"

	mock "github.com/stretchr/testify/mock"
)

// MockNeuralNetwork is an autogenerated mock type for the NeuralNetwork type
type MockNeuralNetwork struct {
	mock.Mock
}

// Generate provides a mock function with given fields: ctx, req
func (_m *MockNeuralNetwork) Generate(ctx context.Context, req model.GenerateReq) (model.GenerateRes, error) {
	ret := _m.Called(ctx, req)

	if len(ret) == 0 {
		panic("no return value specified for Generate")
	}

	var r0 model.GenerateRes
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, model.GenerateReq) (model.GenerateRes, error)); ok {
		return rf(ctx, req)
	}
	if rf, ok := ret.Get(0).(func(context.Context, model.GenerateReq) model.GenerateRes); ok {
		r0 = rf(ctx, req)
	} else {
		r0 = ret.Get(0).(model.GenerateRes)
	}

	if rf, ok := ret.Get(1).(func(context.Context, model.GenerateReq) error); ok {
		r1 = rf(ctx, req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewMockNeuralNetwork creates a new instance of MockNeuralNetwork. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockNeuralNetwork(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockNeuralNetwork {
	mock := &MockNeuralNetwork{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
