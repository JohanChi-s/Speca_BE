package repository

import (
	"context"
	"core_service/internal/model"
)

type ActionEventRepository interface {
	GetActionEvent(ctx context.Context, id string) (*model.ActionEvent, error)
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
