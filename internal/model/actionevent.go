package model

import (
	"time"
)

type ActionEvent struct {
	ID           string `gorm:"primaryKey"`
	Action       string
	CreatedAt    time.Time
	Actor        string
	Assignee     string
	Assigner     string
	DocumentID   string
	CollectionID string
}

func (m *ActionEvent) TableName() string {
	return "action_event"
}
