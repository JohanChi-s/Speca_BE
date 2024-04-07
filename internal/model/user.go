package model

import (
	"time"
)

type User struct {
	ID           string `gorm:"primaryKey"`
	Email        string `gorm:"unique;not null"`
	Username     string `gorm:"unique;not null"`
	Password     string `gorm:"not null"`
	FirstName    string
	LastName     string
	IsAdmin      bool
	IsActive     bool
	IsViewer     bool
	Language     string
	LastActiveAt time.Time
	Roles        []byte
	CreatedAt    time.Time
	UpdatedAt    time.Time `gorm:"autoUpdateTime"`
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
	return "users"
}
