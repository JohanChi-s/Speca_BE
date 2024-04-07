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
	CreatedAt       time.Time
	UpdatedAt       time.Time `gorm:"autoUpdateTime"`
	URL             string
}

type TeamWorkspace struct {
	TeamID      string `gorm:"primaryKey"`
	WorkspaceID string `gorm:"primaryKey"`
}

type TeamDocument struct {
	TeamID     string `gorm:"primaryKey"`
	DocumentID string `gorm:"primaryKey"`
}

func (m *Team) TableName() string {
	return "team"
}
