package service

import (
	"context"
	"core_service/internal/model"
	"core_service/internal/repository"
)

type DocumentService interface {
	GetDocument(ctx context.Context, id string) (*model.Document, error)
}

func NewDocumentService(service *Service, documentRepository repository.DocumentRepository) DocumentService {
	return &documentService{
		Service:            service,
		documentRepository: documentRepository,
	}
}

type documentService struct {
	*Service
	documentRepository repository.DocumentRepository
}

func (s *documentService) GetDocument(ctx context.Context, id string) (*model.Document, error) {
	return s.documentRepository.GetDocument(ctx, id)
}
