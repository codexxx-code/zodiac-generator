// Code generated by mockery v2.46.2. DO NOT EDIT.

package endpoint

import (
	context "context"
	model "partners/internal/services/ebay/model"

	mock "github.com/stretchr/testify/mock"
)

// MockEbayService is an autogenerated mock type for the EbayService type
type MockEbayService struct {
	mock.Mock
}

// GetBreadcrumbs provides a mock function with given fields: _a0, _a1
func (_m *MockEbayService) GetBreadcrumbs(_a0 context.Context, _a1 model.GetBreadcrumbsReq) (model.Category, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for GetBreadcrumbs")
	}

	var r0 model.Category
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, model.GetBreadcrumbsReq) (model.Category, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, model.GetBreadcrumbsReq) model.Category); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Get(0).(model.Category)
	}

	if rf, ok := ret.Get(1).(func(context.Context, model.GetBreadcrumbsReq) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetCategories provides a mock function with given fields: _a0, _a1
func (_m *MockEbayService) GetCategories(_a0 context.Context, _a1 model.GetCategoriesReq) ([]model.Category, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for GetCategories")
	}

	var r0 []model.Category
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, model.GetCategoriesReq) ([]model.Category, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, model.GetCategoriesReq) []model.Category); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.Category)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, model.GetCategoriesReq) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetItems provides a mock function with given fields: _a0, _a1
func (_m *MockEbayService) GetItems(_a0 context.Context, _a1 model.GetItemsReq) ([]model.Item, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for GetItems")
	}

	var r0 []model.Item
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, model.GetItemsReq) ([]model.Item, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, model.GetItemsReq) []model.Item); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.Item)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, model.GetItemsReq) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewMockEbayService creates a new instance of MockEbayService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockEbayService(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockEbayService {
	mock := &MockEbayService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}