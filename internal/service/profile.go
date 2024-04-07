package service

import (
	"context"
	"core_service/internal/model"
	"core_service/internal/repository"
)

type ProfileService interface {
	GetProfile(ctx context.Context, id string) (*model.Profile, error)
}

func NewProfileService(service *Service, profileRepository repository.ProfileRepository) ProfileService {
	return &profileService{
		Service:           service,
		profileRepository: profileRepository,
	}
}

type profileService struct {
	*Service
	profileRepository repository.ProfileRepository
}

func (s *profileService) GetProfile(ctx context.Context, id string) (*model.Profile, error) {
	return s.profileRepository.GetProfile(ctx, id)
}
