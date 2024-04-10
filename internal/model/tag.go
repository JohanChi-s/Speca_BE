package model

import "gorm.io/gorm"

type Tag struct {
	gorm.Model
	Name        string `gorm:"unique"`
	Description string
	Documents   []Document   `gorm:"many2many:document_tags;"`   // Relationship with Document model
	Collections []Collection `gorm:"many2many:collection_tags;"` // Relationship with Collection model
}

func (m *Tag) TableName() string {
	return "tag"
}
