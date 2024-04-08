package model

import (
	"time"
)

type Member struct {
	ID           string `gorm:"primaryKey"`
	UserID       string
	CollectionID string
	DocumentID   string
	WorkspaceID  string
	Role         string
	CreatedAt    time.Time
	UpdateAt     time.Time
}

func (m *Member) TableName() string {
	return "member"
}
