package repository

import (
	"context"
	"core_service/internal/model"
)

type DocumentRepository interface {
	GetDocument(ctx context.Context, id string) (*model.Document, error)
}

func NewDocumentRepository(
	repository *Repository,
) DocumentRepository {
	return &documentRepository{
		Repository: repository,
	}
}

type documentRepository struct {
	*Repository
}

func (r *documentRepository) GetDocument(ctx context.Context, id string) (*model.Document, error) {
	var document model.Document

	return &document, nil
}

// CreateDocument creates a new document record in the database
func (r *documentRepository) CreateDocument(ctx context.Context, document *model.Document) error {
	err := r.db.Create(document).Error
	return err
}

// UpdateDocument updates an existing document record in the database
func (r *documentRepository) UpdateDocument(ctx context.Context, document *model.Document) error {
	err := r.db.Save(document).Error
	return err
}

// DeleteDocument deletes a document record from the database by ID
func (r *documentRepository) DeleteDocument(ctx context.Context, id string) error {
	err := r.db.Where("id = ?", id).Delete(&model.Document{}).Error
	return err
}

// GetAllDocuments returns all document records from the database
func (r *documentRepository) GetAllDocuments(ctx context.Context) ([]*model.Document, error) {
	var documents []*model.Document
	err := r.db.Find(&documents).Error
	return documents, err
}

// GetDocumentsByUserID returns documents record from the database by user ID
func (r *documentRepository) GetDocumentsByUserID(ctx context.Context, userID string) ([]*model.Document, error) {
	var documents []*model.Document
	err := r.db.Where("user_id = ?", userID).Find(&documents).Error
	return documents, err
}

// GetDocumentsByWorkspaceID returns documents record from the database by workspace ID
func (r *documentRepository) GetDocumentsByWorkspaceID(ctx context.Context, workspaceID string) ([]*model.Document, error) {
	var documents []*model.Document
	err := r.db.Where("workspace_id = ?", workspaceID).Find(&documents).Error
	return documents, err
}

// GetDocumentsByTeamID returns documents record from the database by team ID
func (r *documentRepository) GetDocumentsByTeamID(ctx context.Context, teamID string) ([]*model.Document, error) {
	var documents []*model.Document
	err := r.db.Where("team_id = ?", teamID).Find(&documents).Error
	return documents, err
}

// GetDocumentsByWorkspaceIDAndTeamID returns documents record from the database by workspace ID and team ID
func (r *documentRepository) GetDocumentsByWorkspaceIDAndTeamID(ctx context.Context, workspaceID string, teamID string) ([]*model.Document, error) {
	var documents []*model.Document
	err := r.db.Where("workspace_id = ? AND team_id = ?", workspaceID, teamID).Find(&documents).Error
	return documents, err
}

// GetDocumentsByCollectionID returns documents record from the database by collection ID
func (r *documentRepository) GetDocumentsByCollectionID(ctx context.Context, collectionID string) ([]*model.Document, error) {
	var documents []*model.Document
	err := r.db.Where("collection_id = ?", collectionID).Find(&documents).Error
	return documents, err
}
