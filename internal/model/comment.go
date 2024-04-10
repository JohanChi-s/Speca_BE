package model

import (
	"time"
)

type Comment struct {
	ID              string `gorm:"primaryKey"`
	Content         string
	CreatedAt       time.Time `gorm:"default:now()"`
	UpdatedAt       time.Time `gorm:"autoUpdateTime"`
	UserID          string
	User            User `gorm:"foreignKey:UserID"`
	DocumentID      string
	Document        Document `gorm:"foreignKey:DocumentID"`
	ParentCommentID string
}

func (m *Comment) TableName() string {
	return "comment"
}
