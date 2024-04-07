package repository

import (
	"context"
	"core_service/internal/model"
)

type PositionRepository interface {
	GetPosition(ctx context.Context, id string) (*model.Position, error)
}

func NewPositionRepository(
	repository *Repository,
) PositionRepository {
	return &positionRepository{
		Repository: repository,
	}
}

type positionRepository struct {
	*Repository
}

func (r *positionRepository) GetPosition(ctx context.Context, id string) (*model.Position, error) {
	var position model.Position

	return &position, nil
}
