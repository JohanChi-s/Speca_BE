package model

import (
	"time"
)

type Profile struct {
	ID        string `gorm:"primaryKey"`
	UserID    string `gorm:"unique"`
	FullName  string
	Age       int
	Address   string
	AvatarURL string
	CreatedAt time.Time `gorm:"default:now()"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
	User      User      `gorm:"foreignKey:UserID"`
}

func (m *Profile) TableName() string {
	return "profile"
}
