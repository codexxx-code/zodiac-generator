// Code generated by mockery v2.46.2. DO NOT EDIT.

package service

import (
	context "context"
	model "exchange/internal/services/dsp/model"

	mock "github.com/stretchr/testify/mock"
)

// MockDSPRepository is an autogenerated mock type for the DSPRepository type
type MockDSPRepository struct {
	mock.Mock
}

// CreateDSP provides a mock function with given fields: _a0, _a1
func (_m *MockDSPRepository) CreateDSP(_a0 context.Context, _a1 model.CreateDSPReq) error {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for CreateDSP")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, model.CreateDSPReq) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteDSP provides a mock function with given fields: _a0, _a1
func (_m *MockDSPRepository) DeleteDSP(_a0 context.Context, _a1 model.DeleteDSPReq) error {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for DeleteDSP")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, model.DeleteDSPReq) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FindDSPs provides a mock function with given fields: _a0, _a1
func (_m *MockDSPRepository) FindDSPs(_a0 context.Context, _a1 model.FindDSPsReq) ([]model.DSP, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for FindDSPs")
	}

	var r0 []model.DSP
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, model.FindDSPsReq) ([]model.DSP, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, model.FindDSPsReq) []model.DSP); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.DSP)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, model.FindDSPsReq) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetDSPs provides a mock function with given fields: _a0, _a1
func (_m *MockDSPRepository) GetDSPs(_a0 context.Context, _a1 model.GetDSPsReq) ([]model.DSP, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for GetDSPs")
	}

	var r0 []model.DSP
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, model.GetDSPsReq) ([]model.DSP, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, model.GetDSPsReq) []model.DSP); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.DSP)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, model.GetDSPsReq) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetDSPsCount provides a mock function with given fields: _a0, _a1
func (_m *MockDSPRepository) GetDSPsCount(_a0 context.Context, _a1 model.FindDSPsReq) (int, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for GetDSPsCount")
	}

	var r0 int
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, model.FindDSPsReq) (int, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, model.FindDSPsReq) int); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Get(0).(int)
	}

	if rf, ok := ret.Get(1).(func(context.Context, model.FindDSPsReq) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateDSP provides a mock function with given fields: _a0, _a1
func (_m *MockDSPRepository) UpdateDSP(_a0 context.Context, _a1 model.UpdateDSPReq) error {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for UpdateDSP")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, model.UpdateDSPReq) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewMockDSPRepository creates a new instance of MockDSPRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockDSPRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockDSPRepository {
	mock := &MockDSPRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
