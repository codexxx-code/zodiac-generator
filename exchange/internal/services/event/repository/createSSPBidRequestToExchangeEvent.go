package repository

import (
	"context"

	"exchange/internal/config"
	"pkg/openrtb"
)

func (r *EventRepository) CreateSSPBidRequestToExchangeEvent(ctx context.Context, req openrtb.BidRequest) error {
	return r.writeToTopic(ctx, req, config.Load().Queue.Topic.Event.SSPBidRequestsToExchange)
}
