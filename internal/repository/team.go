package repository

import (
	"context"
	"core_service/internal/model"
)

type TeamRepository interface {
	GetTeam(ctx context.Context, id string) (*model.Team, error)
	CreateTeam(ctx context.Context, team *model.Team) error
	UpdateTeam(ctx context.Context, team *model.Team) error
	DeleteTeam(ctx context.Context, id string) error
	GetAllTeams(ctx context.Context) ([]*model.Team, error)
	GetTeamsByUserID(ctx context.Context, userID string) ([]*model.Team, error)
	GetTeamsByWorkspaceID(ctx context.Context, workspaceID string) ([]*model.Team, error)
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

// CreateTeam creates a new team record in the database
func (r *teamRepository) CreateTeam(ctx context.Context, team *model.Team) error {
	err := r.db.Create(team).Error
	return err
}

// UpdateTeam updates an existing team record in the database
func (r *teamRepository) UpdateTeam(ctx context.Context, team *model.Team) error {
	err := r.db.Save(team).Error
	return err
}

// DeleteTeam deletes a team record from the database by ID
func (r *teamRepository) DeleteTeam(ctx context.Context, id string) error {
	err := r.db.Where("id = ?", id).Delete(&model.Team{}).Error
	return err
}

// GetAllTeams returns all team records from the database
func (r *teamRepository) GetAllTeams(ctx context.Context) ([]*model.Team, error) {
	var teams []*model.Team
	err := r.db.Find(&teams).Error
	return teams, err
}

// GetTeasmByUserID returns teams record from the database by user ID
func (r *teamRepository) GetTeamsByUserID(ctx context.Context, userID string) ([]*model.Team, error) {
	var teams []*model.Team
	err := r.db.Where("user_id = ?", userID).Find(&teams).Error
	return teams, err
}

// GetTeamsByWorkspaceID returns teams record from the database by workspace ID
func (r *teamRepository) GetTeamsByWorkspaceID(ctx context.Context, workspaceID string) ([]*model.Team, error) {
	var teams []*model.Team
	err := r.db.Where("workspace_id = ?", workspaceID).Find(&teams).Error
	return teams, err
}
