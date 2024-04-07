package repository

import (
	"context"
	"core_service/internal/model"
)

type ProfileRepository interface {
	GetProfile(ctx context.Context, id string) (*model.Profile, error)
	CreateProfile(ctx context.Context, profile *model.Profile) error
	UpdateProfile(ctx context.Context, profile *model.Profile) error
	DeleteProfile(ctx context.Context, id int64) error
	GetProfileByUserID(ctx context.Context, userID string) (*model.Profile, error)
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

func (r *profileRepository) GetProfile(ctx context.Context, id string) (*model.Profile, error) {
	var profile model.Profile

	return &profile, nil
}

// CreateProfile creates a new profile record in the database
func (r *profileRepository) CreateProfile(ctx context.Context, profile *model.Profile) error {
	err := r.db.Create(profile).Error
	return err
}

// UpdateProfile updates an existing profile record in the database
func (r *profileRepository) UpdateProfile(ctx context.Context, profile *model.Profile) error {
	err := r.db.Save(profile).Error
	return err
}

// DeleteProfile deletes a profile record from the database by ID
func (r *profileRepository) DeleteProfile(ctx context.Context, id int64) error {
	err := r.db.Where("id = ?", id).Delete(&model.Profile{}).Error
	return err
}

// GetProfileByUserID returns profile record from the database by user ID
func (r *profileRepository) GetProfileByUserID(ctx context.Context, userID string) (*model.Profile, error) {
	var profile model.Profile
	err := r.db.Where("user_id = ?", userID).First(&profile).Error
	return &profile, err
}
