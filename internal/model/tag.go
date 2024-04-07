package model

type Tag struct {
	ID          string `gorm:"primaryKey"`
	Name        string `gorm:"unique;not null"`
	Description string
}

func (m *Tag) TableName() string {
	return "tag"
}
