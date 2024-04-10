package model

import (
	"time"
)

type ActionEvent struct {
	ID           string `gorm:"primaryKey"`
	Action       string
	CreatedAt    time.Time `gorm:"default:now()"`
	Actor        string
	AssigneeID   string // Foreign key column
	Assignee     User   `gorm:"foreignKey:AssigneeID"` // Foreign key relationship
	AssignerID   string // Foreign key column
	Assigner     User
	DocumentID   string
	Document     Document `gorm:"foreignKey:DocumentID"`
	CollectionID string
	Collection   Collection `gorm:"foreignKey:CollectionID"`
}
