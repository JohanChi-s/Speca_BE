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
