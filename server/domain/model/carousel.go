package model

import "gorm.io/gorm"

type Carousel struct {
	gorm.Model
	Img   string `gorm:"type:varchar(255);comment:'图片链接'"`
	Title string `gorm:"type:varchar(20);comment:'标题'"`
	Url   string `gorm:"type:varchar(100);comment:'指向的链接'"`
	Color string `gorm:"type:varchar(20);comment:'主题色'"`
}

func (table *Carousel) TableName() string {
	return "carousel"
}
