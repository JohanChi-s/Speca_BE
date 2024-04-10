package model

import (
	"time"
)

type Workspace struct {
	ID        string `gorm:"primaryKey"`
	Name      string
	Domain    string
	IsPublic  bool      `gorm:"default:false"`
	CreatedAt time.Time `gorm:"default:now()"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
	Teams     []Team    `gorm:"foreignKey:WorkspaceID"`
	Members   []Member  `gorm:"foreignKey:WorkspaceID"`
	Document  []Document
}

func (m *Workspace) TableName() string {
	return "workspace"
}
