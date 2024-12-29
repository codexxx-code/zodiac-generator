// Code generated by mockery v2.46.2. DO NOT EDIT.

package beforeResponseToSSP

import (
	context "context"
	analyticWritermodel "exchange/internal/services/analyticWriter/model"

	mock "github.com/stretchr/testify/mock"

	model "exchange/internal/services/exchange/model"
)

// MockExchangeService is an autogenerated mock type for the ExchangeService type
type MockExchangeService struct {
	mock.Mock
}

// CreateAnalyticEventResponseToSSP provides a mock function with given fields: ctx, dto, exchangeBidID, sspResponseBid
func (_m *MockExchangeService) CreateAnalyticEventResponseToSSP(ctx context.Context, dto *model.AuctionDTO, exchangeBidID string, sspResponseBid analyticWritermodel.AnalyticResponseBidModel) error {
	ret := _m.Called(ctx, dto, exchangeBidID, sspResponseBid)

	if len(ret) == 0 {
		panic("no return value specified for CreateAnalyticEventResponseToSSP")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *model.AuctionDTO, string, analyticWritermodel.AnalyticResponseBidModel) error); ok {
		r0 = rf(ctx, dto, exchangeBidID, sspResponseBid)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CreateAnalyticEventResponseToSSPForClickunder provides a mock function with given fields: ctx, dto, sspResponseBid
func (_m *MockExchangeService) CreateAnalyticEventResponseToSSPForClickunder(ctx context.Context, dto *model.AuctionDTO, sspResponseBid analyticWritermodel.AnalyticResponseBidModel) error {
	ret := _m.Called(ctx, dto, sspResponseBid)

	if len(ret) == 0 {
		panic("no return value specified for CreateAnalyticEventResponseToSSPForClickunder")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *model.AuctionDTO, analyticWritermodel.AnalyticResponseBidModel) error); ok {
		r0 = rf(ctx, dto, sspResponseBid)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewMockExchangeService creates a new instance of MockExchangeService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockExchangeService(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockExchangeService {
	mock := &MockExchangeService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
