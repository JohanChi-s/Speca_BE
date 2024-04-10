package model

import (
	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	Content         string
	UserID          uint
	User            User `gorm:"foreignKey:UserID"`
	DocumentID      uint
	Document        Document `gorm:"foreignKey:DocumentID"`
	ParentCommentID uint
}

func (m *Comment) TableName() string {
	return "comment"
}
