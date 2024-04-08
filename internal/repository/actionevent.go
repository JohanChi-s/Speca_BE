package repository

import (
	"context"
	"core_service/internal/model"
)

type ActionEventRepository interface {
	GetActionEvent(ctx context.Context, id string) (*model.ActionEvent, error)
	CreateActionEvent(ctx context.Context, actionEvent *model.ActionEvent) error
	UpdateActionEvent(ctx context.Context, actionEvent *model.ActionEvent) error
	DeleteActionEvent(ctx context.Context, id string) error
	GetAllActionEvents(ctx context.Context) ([]*model.ActionEvent, error)
	GetActionEventsByDocumentID(ctx context.Context, documentID string) ([]*model.ActionEvent, error)
	GetActionEventsByCollectionID(ctx context.Context, collectionID string) ([]*model.ActionEvent, error)
}

func NewActionEventRepository(
	repository *Repository,
) ActionEventRepository {
	return &actionEventRepository{
		Repository: repository,
	}
}

type actionEventRepository struct {
	*Repository
}

func (r *actionEventRepository) GetActionEvent(ctx context.Context, id string) (*model.ActionEvent, error) {
	var actionEvent model.ActionEvent

	return &actionEvent, nil
}

// CreateActionEvent creates a new action event
func (r *actionEventRepository) CreateActionEvent(ctx context.Context, actionEvent *model.ActionEvent) error {
	return r.db.Create(actionEvent).Error
}

// UpdateActionEvent updates an existing action event
func (r *actionEventRepository) UpdateActionEvent(ctx context.Context, actionEvent *model.ActionEvent) error {
	return r.db.Save(actionEvent).Error
}

// DeleteActionEvent deletes an action event by ID
func (r *actionEventRepository) DeleteActionEvent(ctx context.Context, id string) error {
	return r.db.Where("id = ?", id).Delete(&model.ActionEvent{}).Error
}

// GetAllActionEvents returns all action events
func (r *actionEventRepository) GetAllActionEvents(ctx context.Context) ([]*model.ActionEvent, error) {
	var actionEvents []*model.ActionEvent
	err := r.db.Find(&actionEvents).Error
	return actionEvents, err
}

// GetActionEventsByDocumentID returns action events by document ID
func (r *actionEventRepository) GetActionEventsByDocumentID(ctx context.Context, documentID string) ([]*model.ActionEvent, error) {
	var actionEvents []*model.ActionEvent
	err := r.db.Where("document_id = ?", documentID).Find(&actionEvents).Error
	return actionEvents, err
}

// GetActionEventsByCollectionID returns action events by collection ID
func (r *actionEventRepository) GetActionEventsByCollectionID(ctx context.Context, collectionID string) ([]*model.ActionEvent, error) {
	var actionEvents []*model.ActionEvent
	err := r.db.Where("collection_id = ?", collectionID).Find(&actionEvents).Error
	return actionEvents, err
}
