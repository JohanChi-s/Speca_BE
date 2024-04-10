package model

import (
	"time"
)

type Profile struct {
	ID        string    `gorm:"primaryKey;column:id"`
	UserID    string    `gorm:"uniqueIndex;column:userId"`
	FullName  string    `gorm:"column:fullName"`
	Age       int       `gorm:"column:age"`
	Address   string    `gorm:"column:address"`
	AvatarURL string    `gorm:"column:avatarUrl"`
	CreatedAt time.Time `gorm:"column:createdAt;autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime;column:updatedAt"`
}

func (m *Profile) TableName() string {
	return "Profile"
}
