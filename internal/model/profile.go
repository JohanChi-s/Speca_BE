package model

import "gorm.io/gorm"

type Profile struct {
	gorm.Model
}

func (m *Profile) TableName() string {
    return "profile"
}
