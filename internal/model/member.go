package model

import (
	"time"
)

type Member struct {
	ID           string `gorm:"primaryKey"`
	UserID       string `gorm:"unique"`
	User         User   `gorm:"foreignKey:UserID"`
	CollectionID string
	Collection   Collection `gorm:"foreignKey:CollectionID"`
	DocumentID   string
	Document     Document `gorm:"foreignKey:DocumentID"`
	WorkspaceID  string
	Workspace    Workspace `gorm:"foreignKey:WorkspaceID"`
	TeamID       string
	Team         Team `gorm:"foreignKey:TeamID"`
	Role         string
	CreatedAt    time.Time `gorm:"default:now()"`
	UpdatedAt    time.Time `gorm:"autoUpdateTime"`
}

func (m *Member) TableName() string {
	return "member"
}
