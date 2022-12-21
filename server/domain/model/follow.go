package model

import "gorm.io/gorm"

type Follow struct {
	gorm.Model
	Uid uint `gorm:"comment:'用户ID';not null;index"`
	Fid uint `gorm:"comment:'关注的用户ID';not null;index"`
}

func (table *Follow) TableName() string {
	return "follow"
}
