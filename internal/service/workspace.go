package service

import (
	"context"
	"core_service/internal/model"
	"core_service/internal/repository"
)

type WorkspaceService interface {
	GetWorkspace(ctx context.Context, id string) (*model.Workspace, error)
}

func NewWorkspaceService(service *Service, workspaceRepository repository.WorkspaceRepository) WorkspaceService {
	return &workspaceService{
		Service:             service,
		workspaceRepository: workspaceRepository,
	}
}

type workspaceService struct {
	*Service
	workspaceRepository repository.WorkspaceRepository
}

func (s *workspaceService) GetWorkspace(ctx context.Context, id string) (*model.Workspace, error) {
	return s.workspaceRepository.GetWorkspace(ctx, id)
}
