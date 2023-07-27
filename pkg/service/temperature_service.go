package service

import (
	"context"
	"iot_devices/pkg/models"
	"iot_devices/pkg/repository"
	"time"
)

type TemperatureService struct {
	repo *repository.TemperatureRepository
}

func NewTemperatureService(repo *repository.TemperatureRepository) *TemperatureService {
	return &TemperatureService{repo: repo}
}

func (s *TemperatureService) CreateTemperatureData(ctx context.Context, temp models.Temperature) (models.Temperature, error) {
	saved, err := s.repo.AddTemperatureData(ctx, temp)
	if err != nil {
		return models.Temperature{}, err
	}
	newTemp := models.Temperature(*saved)

	return newTemp, nil
}

func (s *TemperatureService) ListTemperatureData(ctx context.Context, start time.Time, end time.Time, limit int) ([]models.Temperature, error) {
	tempList, err := s.repo.ListTemperatureData(ctx, start, end, limit)
	if err != nil {
		return []models.Temperature{}, err
	}

	return tempList, nil
}
