package repository

import (
	"context"
	"core_service/internal/model"
)

type TagRepository interface {
	GetTag(ctx context.Context, id string) (*model.Tag, error)
}

func NewTagRepository(
	repository *Repository,
) TagRepository {
	return &tagRepository{
		Repository: repository,
	}
}

type tagRepository struct {
	*Repository
}

func (r *tagRepository) GetTag(ctx context.Context, id string) (*model.Tag, error) {
	var tag model.Tag

	return &tag, nil
}
