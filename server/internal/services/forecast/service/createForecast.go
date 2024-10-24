package service

import (
	"context"

	"server/internal/services/forecast/model"
)

func (s *ForecastService) CreateForecast(ctx context.Context, req model.CreateForecastReq) (uint32, error) {
	return s.forecastRepository.CreateForecast(ctx, req)
}
