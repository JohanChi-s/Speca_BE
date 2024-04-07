package service

import (
	"context"
	"core_service/internal/model"
	"core_service/internal/repository"
)

type TagService interface {
	GetTag(ctx context.Context, id string) (*model.Tag, error)
}

func NewTagService(service *Service, tagRepository repository.TagRepository) TagService {
	return &tagService{
		Service:       service,
		tagRepository: tagRepository,
	}
}

type tagService struct {
	*Service
	tagRepository repository.TagRepository
}

func (s *tagService) GetTag(ctx context.Context, id string) (*model.Tag, error) {
	return s.tagRepository.GetTag(ctx, id)
}
