package model

import (
	"gorm.io/gorm"
)

type Position struct {
	gorm.Model
	Line      int
	Col       int
	ToLine    string
	ToCol     string
	Content   string
	CommentID uint
	Comment   Comment `gorm:"foreignKey:CommentID"`
}

func (m *Position) TableName() string {
	return "position"
}
