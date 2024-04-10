package model

import (
	"time"
)

type Team struct {
	ID              string `gorm:"primaryKey"`
	Name            string
	AvatarURL       string
	SubDomain       string
	Theme           string
	CanComment      bool
	CanShare        bool
	InviteRequired  bool
	DefaultUserRole string
	CreatedAt       time.Time `gorm:"default:now()"`
	UpdatedAt       time.Time `gorm:"autoUpdateTime"`
	Documents       []Document
	WorkspaceID     string
	Workspace       Workspace `gorm:"foreignKey:WorkspaceID"`
}

func (m *Team) TableName() string {
	return "team"
}
