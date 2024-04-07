package model

import (
	"time"
)

type Profile struct {
	ID        string `gorm:"primaryKey"`
	UserID    string `gorm:"uniqueIndex"`
	FullName  string
	Age       int
	Address   string
	AvatarURL string
	CreatedAt time.Time
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

func (m *Profile) TableName() string {
	return "profile"
}
