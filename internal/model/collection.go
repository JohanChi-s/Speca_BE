package model

import (
	"gorm.io/gorm"
)

type Collection struct {
	gorm.Model
	Name               string
	Description        string
	Icon               string
	IsSaving           bool
	CanShare           bool
	DownloadPermission []string `gorm:"type:text[]"`
	WorkspaceID        uint
	Workspace          Workspace
	ParentCollectionID uint
	UserID             uint
	User               User  `gorm:"foreignKey:UserID"`
	Tags               []Tag `gorm:"many2many:collection_tags;"` // Relationship with Tag model
}

func (m *Collection) TableName() string {
	return "collections"
}
