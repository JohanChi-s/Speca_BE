package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
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
}

func (u *User) TableName() string {
	return "user"
}
