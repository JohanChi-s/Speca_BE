package service

import (
	"context"
	"core_service/internal/model"
	"core_service/internal/repository"
)

type PositionService interface {
	GetPosition(ctx context.Context, id string) (*model.Position, error)
}

func NewPositionService(service *Service, positionRepository repository.PositionRepository) PositionService {
	return &positionService{
		Service:            service,
		positionRepository: positionRepository,
	}
}

type positionService struct {
	*Service
	positionRepository repository.PositionRepository
}

func (s *positionService) GetPosition(ctx context.Context, id string) (*model.Position, error) {
	return s.positionRepository.GetPosition(ctx, id)
}
