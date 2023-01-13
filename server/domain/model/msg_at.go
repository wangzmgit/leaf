package model

import "gorm.io/gorm"

type AtMessage struct {
	gorm.Model
	Vid uint `gorm:"comment:'所在视频id';not null"`
	Uid uint `gorm:"comment:'所属用户ID';not null"`
	Fid uint `gorm:"comment:'关联用户ID';not null"`

	User  User  `gorm:"-"` // 用户
	Video Video `gorm:"-"` // 视频
}

func (table *AtMessage) TableName() string {
	return "msg_at"
}
