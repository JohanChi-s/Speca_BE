package model

import (
	"time"
)

type Collection struct {
	ID                 string `gorm:"primaryKey"`
	Name               string
	Description        string
	Icon               string
	IsSaving           bool
	CanShare           bool
	ChildCollectionIDs string
	DownloadPermission []byte
	CreatedAt          time.Time
	UpdatedAt          time.Time `gorm:"autoUpdateTime"`
	WorkspaceID        string
	ParentCollectionID string
	OwnerUserID        string
}

type CollectionMember struct {
	CollectionID string `gorm:"primaryKey"`
	MemberID     string `gorm:"primaryKey"`
}

func (m *Collection) TableName() string {
	return "collection"
}
