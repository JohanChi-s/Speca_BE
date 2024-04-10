package model

import (
	"gorm.io/gorm"
)

type ActionEvent struct {
	gorm.Model
	Action       string
	Actor        string
	AssigneeID   uint // Foreign key column
	Assignee     User `gorm:"foreignKey:AssigneeID"` // Foreign key relationship
	AssignerID   uint // Foreign key column
	Assigner     User
	DocumentID   uint
	Document     Document `gorm:"foreignKey:DocumentID"`
	CollectionID uint
	Collection   Collection `gorm:"foreignKey:CollectionID"`
}
