package model

import (
	"gorm.io/gorm"
)

type Workspace struct {
	gorm.Model
	Name       string
	Domain     string
	IsPublic   bool `gorm:"default:false"`
	Teams      []Team
	Members    []Member
	Collection []Collection
	Document   []Document
}

func (m *Workspace) TableName() string {
	return "workspace"
}
