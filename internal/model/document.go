package model

import (
	"gorm.io/gorm"
)

type Document struct {
	gorm.Model
	Title        string
	Text         string
	Emoji        string
	IsPublic     bool
	IsFullWidth  bool
	AuthorID     *uint `gorm:"nullable" default:"NULL"`
	TeamID       *uint `gorm:"nullable"`
	WorkspaceID  *uint `gorm:"nullable"`
	CollectionID *uint `gorm:"nullable"`
	Tags         []Tag `gorm:"many2many:document_tags;"` // Relationship with Tag model
}

func (m *Document) TableName() string {
	return "document"
}
