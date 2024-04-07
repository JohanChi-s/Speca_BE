package service

import (
	"context"
	"core_service/internal/model"
	"core_service/internal/repository"
)

type CollectionService interface {
	GetCollection(ctx context.Context, id string) (*model.Collection, error)
}

func NewCollectionService(service *Service, collectionRepository repository.CollectionRepository) CollectionService {
	return &collectionService{
		Service:              service,
		collectionRepository: collectionRepository,
	}
}

type collectionService struct {
	*Service
	collectionRepository repository.CollectionRepository
}

func (s *collectionService) GetCollection(ctx context.Context, id string) (*model.Collection, error) {
	return s.collectionRepository.GetCollection(ctx, id)
}
