package service

import (
	"context"
	v1 "core_service/api/v1"
	"core_service/internal/model"
	"core_service/internal/repository"
	"log"
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserService interface {
	// Auth
	Register(ctx context.Context, req *v1.RegisterRequest) error
	Login(ctx context.Context, req *v1.LoginRequest) (string, error)

	//User
	GetProfile(ctx context.Context, userId uint) (*v1.GetProfileResponseData, error)
	UpdateProfile(ctx context.Context, userId uint, req *v1.UpdateProfileRequest) error
	GetUser(ctx context.Context, user_id uint) (*model.User, error)
	DeleteUserByID(ctx context.Context, user_id uint) error
}

func NewUserService(
	service *Service,
	userRepo repository.UserRepository,
	profileRepo repository.ProfileRepository,
	workspaceRepo repository.WorkspaceRepository,
	memberRepo repository.MemberRepository) UserService {
	return &userService{
		userRepo:      userRepo,
		profileRepo:   profileRepo,
		workspaceRepo: workspaceRepo,
		memberRepo:    memberRepo,
		Service:       service,
	}
}

type userService struct {
	userRepo      repository.UserRepository
	profileRepo   repository.ProfileRepository
	workspaceRepo repository.WorkspaceRepository
	memberRepo    repository.MemberRepository
	*Service
}

// Get User
func (s *userService) GetUser(ctx context.Context, user_id uint) (*model.User, error) {
	user, err := s.userRepo.GetByID(ctx, user_id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *userService) Register(ctx context.Context, req *v1.RegisterRequest) error {
	// check user exist
	log.Println("Registering email", req.Email)
	checkUser, err := s.userRepo.GetUserByEmail(ctx, req.Email)
	if err != nil && err != gorm.ErrRecordNotFound {
		return err
	}
	if checkUser != nil {
		log.Println("User already exists", checkUser.Email)
		return v1.ErrBadRequest
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user := &model.User{
		Email:    req.Email,
		Username: req.Email,
		Password: string(hashedPassword),
	}
	// Transaction demo
	err = s.tm.Transaction(ctx, func(ctx context.Context) error {
		if err := s.userRepo.CreateUser(ctx, user); err != nil {
			return err
		}

		profile := &model.Profile{
			UserID:    user.ID,
			FullName:  user.Email, // Set default full name or prompt user to enter it
			AvatarURL: "",         // Set default avatar URL or leave it empty
		}

		if err := s.profileRepo.CreateProfile(ctx, profile); err != nil {
			return err
		}

		// Create default workspace for the new user
		defaultWorkspace := &model.Workspace{
			Name:     "My Workspace",
			Domain:   "my_workspace",
			IsPublic: false, // Adjust visibility as needed
		}
		if err := s.workspaceRepo.CreateWorkspace(ctx, defaultWorkspace); err != nil {
			return err
		}

		// Create member with role owner and assign it to the workspace
		ownerMember := &model.Member{
			UserID:      user.ID,
			WorkspaceID: &defaultWorkspace.ID,
			DocumentID:  nil,
			Role:        v1.RoleToString(v1.OwnerRole),
		}
		if err := s.memberRepo.CreateMember(ctx, ownerMember); err != nil {
			return err
		}

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

func (s *userService) GetProfile(ctx context.Context, userId uint) (*v1.GetProfileResponseData, error) {
	user, err := s.userRepo.GetByID(ctx, userId)
	if err != nil {
		return nil, err
	}

	return &v1.GetProfileResponseData{
		UserId:   user.ID,
		Username: user.Username,
	}, nil
}

func (s *userService) UpdateProfile(ctx context.Context, userId uint, req *v1.UpdateProfileRequest) error {
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

func (s *userService) DeleteUserByID(ctx context.Context, user_id uint) error {
	return s.userRepo.DeleteUser(ctx, user_id)
}
