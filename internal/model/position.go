package model

import (
	"time"
)

type Position struct {
	ID        string `gorm:"primaryKey"`
	Line      int
	Col       int
	ToLine    string
	ToCol     string
	Content   string
	CreatedAt time.Time `gorm:"default:now()"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
	CommentID string
	Comment   Comment `gorm:"foreignKey:CommentID"`
}

func (m *Position) TableName() string {
	return "position"
}
