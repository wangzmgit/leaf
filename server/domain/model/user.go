package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username   string    `gorm:"type:varchar(30);comment:'用户名';not null;unique"`
	Email      string    `gorm:"type:varchar(30);comment:'邮箱';not null;unique"`
	Password   string    `gorm:"type:varchar(128);comment:'密码'not null"`
	ClientIp   string    `gorm:"type:varchar(20);comment:'客户端IP'"`
	LoginTime  time.Time `gorm:"type:datetime(3);comment:'登录时间'"`
	LogoutTime time.Time `gorm:"type:datetime(3);comment:'退出时间'"`
	IsLogout   bool      `gorm:"type:tinyint(1);comment:'是否退出'"`
	Status     string    `gorm:"type:char(1);;commment:'账号状态:0正常、1停用、2删除';default:'0'"`
}

func (table *User) TableName() string {
	return "user"
}
