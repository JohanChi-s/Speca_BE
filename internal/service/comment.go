package service

import (
	"context"
	"core_service/internal/model"
	"core_service/internal/repository"
)

type CommentService interface {
	GetComment(ctx context.Context, id string) (*model.Comment, error)
}

func NewCommentService(service *Service, commentRepository repository.CommentRepository) CommentService {
	return &commentService{
		Service:           service,
		commentRepository: commentRepository,
	}
}

type commentService struct {
	*Service
	commentRepository repository.CommentRepository
}

func (s *commentService) GetComment(ctx context.Context, id string) (*model.Comment, error) {
	return s.commentRepository.GetComment(ctx, id)
}
