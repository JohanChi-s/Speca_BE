package repository

import (
	"context"
	"core_service/internal/model"
)

type TeamRepository interface {
	GetTeam(ctx context.Context, id string) (*model.Team, error)
}

func NewTeamRepository(
	repository *Repository,
) TeamRepository {
	return &teamRepository{
		Repository: repository,
	}
}

type teamRepository struct {
	*Repository
}

func (r *teamRepository) GetTeam(ctx context.Context, id string) (*model.Team, error) {
	var team model.Team

	return &team, nil
}
