package repository

import (
	"context"
	"core_service/internal/model"
)

type WorkspaceRepository interface {
	GetWorkspace(ctx context.Context, id string) (*model.Workspace, error)
}

func NewWorkspaceRepository(
	repository *Repository,
) WorkspaceRepository {
	return &workspaceRepository{
		Repository: repository,
	}
}

type workspaceRepository struct {
	*Repository
}

func (r *workspaceRepository) GetWorkspace(ctx context.Context, id string) (*model.Workspace, error) {
	var workspace model.Workspace

	return &workspace, nil
}
