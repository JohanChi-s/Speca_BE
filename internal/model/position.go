package model

import (
	"time"
)

type Position struct {
	ID         string `gorm:"primaryKey"`
	Line       int
	Col        int
	ToLine     string
	ToCol      string
	Content    string
	CreatedAt  time.Time
	UpdatedAt  time.Time `gorm:"autoUpdateTime"`
	CommentsID string
}

func (m *Position) TableName() string {
	return "position"
}
