package model

import "gorm.io/gorm"

//收藏夹
type Collection struct {
	gorm.Model
	Uid   uint   `gorm:"comment:'所属用户';not null"`
	Name  string `gorm:"type:varchar(20);comment:'收藏夹名称'"`
	Desc  string `gorm:"type:varchar(150);comment:'简介'"`
	Cover string `gorm:"size:255;comment:'封面'"`
	Open  bool   `gorm:"comment:'是否公开';default:false"`
}

func (table *Collection) TableName() string {
	return "collection"
}
