package service

import (
	"context"
	"iot_devices/pkg/models"
	"iot_devices/pkg/repository"
	"time"
)

type VelocityService struct {
	repo *repository.VelocityRepository
}

func NewVelocityService(repo *repository.VelocityRepository) *VelocityService {
	return &VelocityService{repo: repo}
}

func (s *VelocityService) CreateVelocityData(ctx context.Context, vel models.Velocity) (models.Velocity, error) {
	saved, err := s.repo.AddVelocityData(ctx, vel)
	if err != nil {
		return models.Velocity{}, err
	}
	newTemp := models.Velocity(*saved)

	return newTemp, nil
}

func (s *VelocityService) ListVelocityData(ctx context.Context, start time.Time, end time.Time, limit int) ([]models.Velocity, error) {
	velList, err := s.repo.ListVelocityData(ctx, start, end, limit)
	if err != nil {
		return []models.Velocity{}, err
	}

	return velList, nil
}
