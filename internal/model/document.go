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
	CreatedAt    time.Time
	UpdatedAt    time.Time `gorm:"autoUpdateTime"`
	AuthorID     string
	TeamID       string
	WorkspaceID  string
	CollectionID string
}

type DocumentCollection struct {
	DocumentID   string `gorm:"primaryKey"`
	CollectionID string `gorm:"primaryKey"`
}

func (m *Document) TableName() string {
	return "document"
}
