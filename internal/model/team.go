package model

import (
	"gorm.io/gorm"
)

type Team struct {
	gorm.Model
	Name            string
	AvatarURL       string
	SubDomain       string
	Theme           string
	CanComment      bool
	CanShare        bool
	InviteRequired  bool
	DefaultUserRole string
	Documents       []Document
	WorkspaceID     uint
	Workspace       Workspace `gorm:"foreignKey:WorkspaceID"`
}

func (m *Team) TableName() string {
	return "team"
}
