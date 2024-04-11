package model

import (
	"gorm.io/gorm"
)

type Profile struct {
	gorm.Model
	UserID    uint `gorm:"unique"`
	FullName  string
	Age       int
	Address   string
	AvatarURL string
	User      User `gorm:"foreignKey:UserID"`
}

func (m *Profile) TableName() string {
	return "profile"
}
