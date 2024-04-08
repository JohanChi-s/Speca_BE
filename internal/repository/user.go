package repository

import (
	"context"
	v1 "core_service/api/v1"
	"core_service/internal/model"
	"errors"

	"gorm.io/gorm"
)

type UserRepository interface {
	Create(ctx context.Context, user *model.User) error
	Update(ctx context.Context, user *model.User) error
	GetByID(ctx context.Context, id string) (*model.User, error)
	GetByEmail(ctx context.Context, email string) (*model.User, error)
	CreateUser(ctx context.Context, user *model.User) error
	UpdateUser(ctx context.Context, user *model.User) error
	DeleteUser(ctx context.Context, id string) error
	GetAllUsers(ctx context.Context) ([]*model.User, error)
	GetUserByWorkspaceID(ctx context.Context, workspaceID string) (*model.User, error)
	GetUserByUserID(ctx context.Context, userID string) (*model.User, error)
	GetUserByEmail(ctx context.Context, email string) (*model.User, error)
	GetUserByRefreshToken(ctx context.Context, refreshToken string) (*model.User, error)
	GetUserByUsername(ctx context.Context, username string) (*model.User, error)
}

func NewUserRepository(r *Repository) UserRepository {
	return &userRepository{
		Repository: r,
	}
}

type userRepository struct {
	*Repository
}

func (r *userRepository) Create(ctx context.Context, user *model.User) error {
	if err := r.DB(ctx).Create(user).Error; err != nil {
		return err
	}
	return nil
}

func (r *userRepository) Update(ctx context.Context, user *model.User) error {
	if err := r.DB(ctx).Save(user).Error; err != nil {
		return err
	}
	return nil
}

func (r *userRepository) GetByID(ctx context.Context, userId string) (*model.User, error) {
	var user model.User
	if err := r.DB(ctx).Where("user_id = ?", userId).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, v1.ErrNotFound
		}
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) GetByEmail(ctx context.Context, email string) (*model.User, error) {
	var user model.User
	if err := r.DB(ctx).Where("email = ?", email).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

// CreateUser creates a new user record in the database
func (r *userRepository) CreateUser(ctx context.Context, user *model.User) error {
	err := r.db.Create(user).Error
	return err
}

// UpdateUser updates an existing user record in the database
func (r *userRepository) UpdateUser(ctx context.Context, user *model.User) error {
	err := r.db.Save(user).Error
	return err
}

// DeleteUser deletes a user record from the database by ID
func (r *userRepository) DeleteUser(ctx context.Context, id string) error {
	err := r.db.Where("id = ?", id).Delete(&model.User{}).Error
	return err
}

// GetAllUsers returns all user records from the database
func (r *userRepository) GetAllUsers(ctx context.Context) ([]*model.User, error) {
	var users []*model.User
	err := r.db.Find(&users).Error
	return users, err
}

// GetUserByWorkspaceID returns a user record from the database by workspace ID
func (r *userRepository) GetUserByWorkspaceID(ctx context.Context, workspaceID string) (*model.User, error) {
	var user model.User
	err := r.db.Where("workspace_id = ?", workspaceID).First(&user).Error
	return &user, err
}

// GetUserByUserID returns a user record from the database by user ID
func (r *userRepository) GetUserByUserID(ctx context.Context, userID string) (*model.User, error) {
	var user model.User
	err := r.db.Where("user_id = ?", userID).First(&user).Error
	return &user, err
}

// GetUserByEmail returns a user record from the database by email
func (r *userRepository) GetUserByEmail(ctx context.Context, email string) (*model.User, error) {
	var user model.User
	err := r.db.Where("email = ?", email).First(&user).Error
	return &user, err
}

// GetUserByRefreshToken returns a user record from the database by refresh token
func (r *userRepository) GetUserByRefreshToken(ctx context.Context, refreshToken string) (*model.User, error) {
	var user model.User
	err := r.db.Where("refresh_token = ?", refreshToken).First(&user).Error
	return &user, err
}

// GetUserByUsername returns a user record from the database by username
func (r *userRepository) GetUserByUsername(ctx context.Context, username string) (*model.User, error) {
	var user model.User
	err := r.db.Where("username = ?", username).First(&user).Error
	return &user, err
}
