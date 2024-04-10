package model

import (
	"time"
)

type User struct {
	ID           string    `gorm:"primaryKey column:id"`
	Email        string    `gorm:"unique column:email"`
	Username     string    `gorm:"unique column:username"`
	Password     string    `gorm:"column:password"`
	FirstName    string    `gorm:"column:firstName"`
	LastName     string    `gorm:"column:lastName"`
	IsAdmin      bool      `gorm:"column:isAdmin"`
	IsActive     bool      `gorm:"column:isActive"`
	IsViewer     bool      `gorm:"column:isViewer"`
	Language     string    `gorm:"column:language"`
	LastActiveAt time.Time `gorm:"column:lastActiveAt;default:now()"`
	Roles        []string  `gorm:"type:text[]"`
	CreatedAt    time.Time `gorm:"column:createdAt;default:now()"`
	UpdatedAt    time.Time `gorm:"column:updatedAt;default:now()"`
}

type UserTeam struct {
	UserID string `gorm:"primaryKey"`
	TeamID string `gorm:"primaryKey"`
}

type UserWorkspace struct {
	UserID      string `gorm:"primaryKey"`
	WorkspaceID string `gorm:"primaryKey"`
}

type UserDocument struct {
	UserID     string `gorm:"primaryKey"`
	DocumentID string `gorm:"primaryKey"`
}

func (u *User) TableName() string {
	return "User"
}
