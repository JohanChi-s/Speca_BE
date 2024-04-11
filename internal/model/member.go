package model

import (
	"gorm.io/gorm"
)

type Member struct {
	gorm.Model
	UserID       uint  `gorm:"unique"`
	User         User  `gorm:"foreignKey:UserID"`
	CollectionID *uint `gorm:"nullable"`
	DocumentID   *uint `gorm:"nullable"`
	WorkspaceID  *uint `gorm:"nullable"`
	TeamID       *uint `gorm:"nullable"`
	Role         string
}

func (m *Member) TableName() string {
	return "member"
}
