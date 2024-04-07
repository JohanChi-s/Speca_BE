package repository

import (
	"context"
	"core_service/internal/model"
)

type CommentRepository interface {
	GetComment(ctx context.Context, id string) (*model.Comment, error)
}

func NewCommentRepository(
	repository *Repository,
) CommentRepository {
	return &commentRepository{
		Repository: repository,
	}
}

type commentRepository struct {
	*Repository
}

func (r *commentRepository) GetComment(ctx context.Context, id string) (*model.Comment, error) {
	var comment model.Comment

	return &comment, nil
}
