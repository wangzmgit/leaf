package model

import "gorm.io/gorm"

type Partition struct {
	gorm.Model
	Content  string `gorm:"varchar(20);comment:'分区名称'"`
	ParentId uint   `gorm:"default:0;comment:'所属分区ID'"`
}

func (table *Partition) TableName() string {
	return "partition"
}
