package model

import (
	"time"
)

type Workspace struct {
	ID        string `gorm:"primaryKey"`
	Name      string
	Domain    string
	IsPublic  bool
	CreatedAt time.Time
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
	URL       string
}

type WorkspaceDocument struct {
	WorkspaceID string `gorm:"primaryKey"`
	DocumentID  string `gorm:"primaryKey"`
}

func (m *Workspace) TableName() string {
	return "workspace"
}
