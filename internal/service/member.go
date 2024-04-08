package service

import (
    "context"
	"core_service/internal/model"
	"core_service/internal/repository"
)

type MemberService interface {
	GetMember(ctx context.Context, id int64) (*model.Member, error)
}

func NewMemberService(service *Service, memberRepository repository.MemberRepository) MemberService {
	return &memberService{
		Service:        service,
		memberRepository: memberRepository,
	}
}

type memberService struct {
	*Service
	memberRepository repository.MemberRepository
}

func (s *memberService) GetMember(ctx context.Context, id int64) (*model.Member, error) {
	return s.memberRepository.GetMember(ctx, id)
}
