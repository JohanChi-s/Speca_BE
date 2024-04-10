package model

import (
	"time"
)

type Collection struct {
	ID                 string `gorm:"primaryKey"`
	Name               string
	Description        string
	Icon               string
	IsSaving           bool
	CanShare           bool
	DownloadPermission []string  `gorm:"type:text[]"`
	CreatedAt          time.Time `gorm:"default:now()"`
	UpdatedAt          time.Time `gorm:"autoUpdateTime"`
	WorkspaceID        string
	Workspace          Workspace `gorm:"foreignKey:WorkspaceID"`
	ParentCollectionID string
	OwnerID            string
	OwnerUser          User  `gorm:"foreignKey:OwnerID"`
	Tags               []Tag `gorm:"many2many:collection_tags;"` // Relationship with Tag model
}

func (m *Collection) TableName() string {
	return "collection"
}
