package repository

import (
	"context"
	"core_service/internal/model"
)

type CollectionRepository interface {
	GetCollection(ctx context.Context, id string) (*model.Collection, error)
}

func NewCollectionRepository(
	repository *Repository,
) CollectionRepository {
	return &collectionRepository{
		Repository: repository,
	}
}

type collectionRepository struct {
	*Repository
}

func (r *collectionRepository) GetCollection(ctx context.Context, id string) (*model.Collection, error) {
	var collection model.Collection

	return &collection, nil
}
