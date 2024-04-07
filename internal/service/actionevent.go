package service

import (
	"context"
	"core_service/internal/model"
	"core_service/internal/repository"
)

type ActionEventService interface {
	GetActionEvent(ctx context.Context, id string) (*model.ActionEvent, error)
}

func NewActionEventService(service *Service, actionEventRepository repository.ActionEventRepository) ActionEventService {
	return &actionEventService{
		Service:               service,
		actionEventRepository: actionEventRepository,
	}
}

type actionEventService struct {
	*Service
	actionEventRepository repository.ActionEventRepository
}

func (s *actionEventService) GetActionEvent(ctx context.Context, id string) (*model.ActionEvent, error) {
	return s.actionEventRepository.GetActionEvent(ctx, id)
}
