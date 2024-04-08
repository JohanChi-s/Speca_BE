package repository

import (
	"context"
	"core_service/internal/model"
)

type CollectionRepository interface {
	GetCollection(ctx context.Context, id string) (*model.Collection, error)
	CreateCollection(ctx context.Context, collection *model.Collection) error
	UpdateCollection(ctx context.Context, collection *model.Collection) error
	DeleteCollection(ctx context.Context, id string) error
	GetAllCollections(ctx context.Context) ([]*model.Collection, error)
	GetCollectionsByTeamID(ctx context.Context, teamID string) ([]*model.Collection, error)
	GetCollectionsByWorkspaceID(ctx context.Context, workspaceID string) ([]*model.Collection, error)
	FilterCollections(ctx context.Context, filters map[string]interface{}) ([]*model.Collection, error)
}

func NewCollectionRepository(
	repository *Repository,
) CollectionRepository {
	return &collectionRepository{
		Repository: repository,
	}
}

type collectionRepository struct {
	*Repository
}

func (r *collectionRepository) GetCollection(ctx context.Context, id string) (*model.Collection, error) {
	var collection model.Collection

	return &collection, nil
}

// CreateCollection creates a new collection
func (r *collectionRepository) CreateCollection(ctx context.Context, collection *model.Collection) error {
	return r.db.Create(collection).Error
}

// UpdateCollection updates an existing collection
func (r *collectionRepository) UpdateCollection(ctx context.Context, collection *model.Collection) error {
	return r.db.Save(collection).Error
}

// DeleteCollection deletes a collection by ID
func (r *collectionRepository) DeleteCollection(ctx context.Context, id string) error {
	return r.db.Where("id = ?", id).Delete(&model.Collection{}).Error
}

// GetAllCollections returns all collections
func (r *collectionRepository) GetAllCollections(ctx context.Context) ([]*model.Collection, error) {
	var collections []*model.Collection
	err := r.db.Find(&collections).Error
	return collections, err
}

// GetCollectionsByTeamID returns collections by team ID
func (r *collectionRepository) GetCollectionsByTeamID(ctx context.Context, teamID string) ([]*model.Collection, error) {
	var collections []*model.Collection
	err := r.db.Where("team_id = ?", teamID).Find(&collections).Error
	return collections, err
}

// GetCollectionsByWorkspaceID returns collections by workspace ID
func (r *collectionRepository) GetCollectionsByWorkspaceID(ctx context.Context, workspaceID string) ([]*model.Collection, error) {
	var collections []*model.Collection
	err := r.db.Where("workspace_id = ?", workspaceID).Find(&collections).Error
	return collections, err
}

// FilterCollections filters collections based on the provided filters
func (r *collectionRepository) FilterCollections(ctx context.Context, filters map[string]interface{}) ([]*model.Collection, error) {
	var collections []*model.Collection
	query := r.db
	for key, value := range filters {
		query = query.Where(key, value)
	}
	err := query.Find(&collections).Error
	return collections, err
}
