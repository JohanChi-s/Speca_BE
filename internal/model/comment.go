package model

import (
	"time"
)

type Comment struct {
	ID              string `gorm:"primaryKey"`
	Content         string
	CreatedAt       time.Time
	UpdatedAt       time.Time `gorm:"autoUpdateTime"`
	UserID          string
	DocumentID      string
	ParentCommentID string
}

func (m *Comment) TableName() string {
	return "comment"
}
