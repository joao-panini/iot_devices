package service

import (
	"context"
	"iot_devices/pkg/models"
	"iot_devices/pkg/repository"
	"time"
)

type TrafficService struct {
	repo *repository.TrafficRepository
}

func NewTrafficService(repo *repository.TrafficRepository) *TrafficService {
	return &TrafficService{repo: repo}
}

func (s *TrafficService) CreateTrafficData(ctx context.Context, traffic models.Traffic) (models.Traffic, error) {
	saved, err := s.repo.AddTrafficData(ctx, traffic)
	if err != nil {
		return models.Traffic{}, err
	}
	newTemp := models.Traffic(*saved)

	return newTemp, nil
}

func (s *TrafficService) ListTrafficData(ctx context.Context, start time.Time, end time.Time, limit int) ([]models.Traffic, error) {
	trafficList, err := s.repo.ListTrafficData(ctx, start, end, limit)
	if err != nil {
		return []models.Traffic{}, err
	}

	return trafficList, nil
}
