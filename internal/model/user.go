package model

import (
	"time"
)

type User struct {
	ID           string `gorm:"primaryKey"`
	Email        string `gorm:"unique"`
	Username     string `gorm:"unique"`
	Password     string
	FirstName    string
	LastName     string
	IsAdmin      bool `gorm:"default:false"`
	IsActive     bool
	IsViewer     bool
	Language     string
	LastActiveAt time.Time `gorm:"autoUpdateTime"`
	Roles        []string  `gorm:"type:text[]"`
	CreatedAt    time.Time `gorm:"default:now()"`
	UpdatedAt    time.Time `gorm:"autoUpdateTime"`
}

func (u *User) TableName() string {
	return "user"
}
