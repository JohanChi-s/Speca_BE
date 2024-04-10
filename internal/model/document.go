package model

import (
	"time"
)

type Document struct {
	ID           string `gorm:"primaryKey"`
	Title        string
	Text         string
	Emoji        string
	IsPublic     bool
	IsFullWidth  bool
	CreatedAt    time.Time `gorm:"default:now()"`
	UpdatedAt    time.Time `gorm:"autoUpdateTime"`
	AuthorID     string
	Author       User `gorm:"foreignKey:AuthorID"`
	TeamID       string
	Team         Team `gorm:"foreignKey:TeamID"`
	WorkspaceID  string
	Workspace    Workspace `gorm:"foreignKey:WorkspaceID"`
	CollectionID string
	Collection   Collection `gorm:"foreignKey:CollectionID"`
	Tags         []Tag      `gorm:"many2many:document_tags;"` // Relationship with Tag model
}

func (m *Document) TableName() string {
	return "document"
}
