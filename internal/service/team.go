package service

import (
	"context"
	"core_service/internal/model"
	"core_service/internal/repository"
)

type TeamService interface {
	GetTeam(ctx context.Context, id string) (*model.Team, error)
}

func NewTeamService(service *Service, teamRepository repository.TeamRepository) TeamService {
	return &teamService{
		Service:        service,
		teamRepository: teamRepository,
	}
}

type teamService struct {
	*Service
	teamRepository repository.TeamRepository
}

func (s *teamService) GetTeam(ctx context.Context, id string) (*model.Team, error) {
	return s.teamRepository.GetTeam(ctx, id)
}
