package service

import (
	"context"
	v1 "core_service/api/v1"
	"core_service/internal/model"
	"core_service/internal/repository"
	"errors"
	"time"

	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	Register(ctx context.Context, req *v1.RegisterRequest) error
	Login(ctx context.Context, req *v1.LoginRequest) (string, error)
	GetProfile(ctx context.Context, userId string) (*v1.GetProfileResponseData, error)
	UpdateProfile(ctx context.Context, userId string, req *v1.UpdateProfileRequest) error
	GetUser(ctx context.Context, user_id string) (*model.User, error)
	DeleteUserByID(ctx context.Context, user_id string) error
}

func NewUserService(service *Service, userRepo repository.UserRepository, profileRepo repository.ProfileRepository) UserService {
	return &userService{
		userRepo:    userRepo,
		profileRepo: profileRepo,
		Service:     service,
	}
}

type userService struct {
	userRepo repository.UserRepository
	// workspaceRepo repository.WorkspaceRepository
	profileRepo repository.ProfileRepository
	// docRepo       repository.DocumentRepository
	*Service
}

// Get User
func (s *userService) GetUser(ctx context.Context, user_id string) (*model.User, error) {
	user, err := s.userRepo.GetByID(ctx, user_id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *userService) Register(ctx context.Context, req *v1.RegisterRequest) error {
	// check username
	s.logger.Info("Register", zap.String("email", req.Email))
	user, err := s.userRepo.GetByEmail(ctx, req.Email)
	if err != nil {
		return v1.ErrInternalServerError
	}
	if err == nil && user != nil {
		return v1.ErrEmailAlreadyUse
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	// Generate user ID
	userId, err := s.sid.GenUUID()
	if err != nil {
		return err
	}
	user = &model.User{
		ID:       userId,
		Email:    req.Email,
		Password: string(hashedPassword),
	}
	// Transaction demo
	err = s.tm.Transaction(ctx, func(ctx context.Context) error {
		// create Profile
		profile := &model.Profile{
			UserID: userId,
		}
		if s.profileRepo == nil {
			return errors.New("profileRepo is nil")
		}

		if err = s.profileRepo.CreateProfile(ctx, profile); err != nil {
			return err
		}

		// Create a user
		if err = s.userRepo.Create(ctx, user); err != nil {
			return err
		}
		// todo: other repo
		return nil
	})
	return err
}

func (s *userService) Login(ctx context.Context, req *v1.LoginRequest) (string, error) {
	user, err := s.userRepo.GetByEmail(ctx, req.Email)
	if err != nil || user == nil {
		return "", v1.ErrUnauthorized
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return "", err
	}
	token, err := s.jwt.GenToken(user.ID, time.Now().Add(time.Hour*24*90))
	if err != nil {
		return "", err
	}

	return token, nil
}

func (s *userService) GetProfile(ctx context.Context, userId string) (*v1.GetProfileResponseData, error) {
	user, err := s.userRepo.GetByID(ctx, userId)
	if err != nil {
		return nil, err
	}

	return &v1.GetProfileResponseData{
		UserId:   user.ID,
		Username: user.Username,
	}, nil
}

func (s *userService) UpdateProfile(ctx context.Context, userId string, req *v1.UpdateProfileRequest) error {
	user, err := s.userRepo.GetByID(ctx, userId)
	if err != nil {
		return err
	}

	user.Email = req.Email
	user.Username = req.Username

	if err = s.userRepo.Update(ctx, user); err != nil {
		return err
	}

	return nil
}

func (s *userService) DeleteUserByID(ctx context.Context, user_id string) error {
	return s.userRepo.DeleteUser(ctx, user_id)
}
