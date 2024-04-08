package repository

import (
	"context"
	"core_service/internal/model"
)

type WorkspaceRepository interface {
	GetWorkspace(ctx context.Context, id string) (*model.Workspace, error)
	CreateWorkspace(ctx context.Context, workspace *model.Workspace) error
	UpdateWorkspace(ctx context.Context, workspace *model.Workspace) error
	DeleteWorkspace(ctx context.Context, id string) error
	GetAllWorkspaces(ctx context.Context) ([]*model.Workspace, error)
	GetWorkspaceByUserID(ctx context.Context, userID string) (*model.Workspace, error)
	GetWorkspaceByWorkspaceID(ctx context.Context, workspaceID string) (*model.Workspace, error)
	GetWorkspaceByWorkspaceIDAndUserID(ctx context.Context, workspaceID string, userID string) (*model.Workspace, error)
	UpsertWorkspace(ctx context.Context, workspace *model.Workspace) error
	AddUserToWorkspace(ctx context.Context, workspaceID string, userID string) error
	RemoveUserFromWorkspace(ctx context.Context, workspaceID string, userID string) error
	AddTeamToWorkspace(ctx context.Context, workspaceID string, teamID string) error
	RemoveTeamFromWorkspace(ctx context.Context, workspaceID string, teamID string) error
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

// CreateWorkspace creates a new workspace record in the database
func (r *workspaceRepository) CreateWorkspace(ctx context.Context, workspace *model.Workspace) error {
	err := r.db.Create(workspace).Error
	return err
}

// UpdateWorkspace updates an existing workspace record in the database
func (r *workspaceRepository) UpdateWorkspace(ctx context.Context, workspace *model.Workspace) error {
	err := r.db.Save(workspace).Error
	return err
}

// DeleteWorkspace deletes a workspace record from the database by ID
func (r *workspaceRepository) DeleteWorkspace(ctx context.Context, id string) error {
	err := r.db.Where("id = ?", id).Delete(&model.Workspace{}).Error
	return err
}

// GetAllWorkspaces returns all workspace records from the database
func (r *workspaceRepository) GetAllWorkspaces(ctx context.Context) ([]*model.Workspace, error) {
	var workspaces []*model.Workspace
	err := r.db.Find(&workspaces).Error
	return workspaces, err
}

// GetWorkspaceByUserID returns a workspace record from the database by user ID
func (r *workspaceRepository) GetWorkspaceByUserID(ctx context.Context, userID string) (*model.Workspace, error) {
	var workspace model.Workspace
	err := r.db.Where("user_id = ?", userID).First(&workspace).Error
	return &workspace, err
}

// GetWorkspaceByWorkspaceID returns a workspace record from the database by workspace ID
func (r *workspaceRepository) GetWorkspaceByWorkspaceID(ctx context.Context, workspaceID string) (*model.Workspace, error) {
	var workspace model.Workspace
	err := r.db.Where("workspace_id = ?", workspaceID).First(&workspace).Error
	return &workspace, err
}

// GetWorkspaceByWorkspaceIDAndUserID returns a workspace record from the database by workspace ID and user ID
func (r *workspaceRepository) GetWorkspaceByWorkspaceIDAndUserID(ctx context.Context, workspaceID string, userID string) (*model.Workspace, error) {
	var workspace model.Workspace
	err := r.db.Where("workspace_id = ? AND user_id = ?", workspaceID, userID).First(&workspace).Error
	return &workspace, err
}

// UpsertWorkspace upserts a workspace record in the database
func (r *workspaceRepository) UpsertWorkspace(ctx context.Context, workspace *model.Workspace) error {
	err := r.db.Save(workspace).Error
	return err
}

// AddUserToWorkspace adds a user to a workspace
func (r *workspaceRepository) AddUserToWorkspace(ctx context.Context, workspaceID string, userID string) error {
	err := r.db.Exec("INSERT INTO workspace_users (workspace_id, user_id) VALUES (?, ?)", workspaceID, userID).Error
	return err
}

// RemoveUserFromWorkspace removes a user from a workspace
func (r *workspaceRepository) RemoveUserFromWorkspace(ctx context.Context, workspaceID string, userID string) error {
	err := r.db.Exec("DELETE FROM workspace_users WHERE workspace_id = ? AND user_id = ?", workspaceID, userID).Error
	return err
}

// AddTeamToWorkspace adds a team to a workspace
func (r *workspaceRepository) AddTeamToWorkspace(ctx context.Context, workspaceID string, teamID string) error {
	err := r.db.Exec("INSERT INTO workspace_teams (workspace_id, team_id) VALUES (?, ?)", workspaceID, teamID).Error
	return err
}

// RemoveTeamFromWorkspace removes a team from a workspace
func (r *workspaceRepository) RemoveTeamFromWorkspace(ctx context.Context, workspaceID string, teamID string) error {
	err := r.db.Exec("DELETE FROM workspace_teams WHERE workspace_id = ? AND team_id = ?", workspaceID, teamID).Error
	return err
}
