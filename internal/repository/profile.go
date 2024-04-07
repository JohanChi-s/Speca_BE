package repository

import (
	"context"
	"core_service/internal/model"
)

type ProfileRepository interface {
	GetProfile(ctx context.Context, id int64) (*model.Profile, error)
}

func NewProfileRepository(
	repository *Repository,
) ProfileRepository {
	return &profileRepository{
		Repository: repository,
	}
}

type profileRepository struct {
	*Repository
}

func (r *profileRepository) GetProfile(ctx context.Context, id int64) (*model.Profile, error) {
	var profile model.Profile

	return &profile, nil
}
