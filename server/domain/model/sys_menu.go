package model

import (
	"gorm.io/gorm"
)

type SysMenu struct {
	gorm.Model
	Name   string `gorm:"type:varchar(10);comment:'菜单名称'"`
	Router string `gorm:"type:varchar(128);comment:'菜单路径'"`
}

func (table *SysMenu) TableName() string {
	return "sys_menu"
}
