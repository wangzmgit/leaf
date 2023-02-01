package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username   string    `gorm:"type:varchar(30);comment:'用户名';not null"`
	Email      string    `gorm:"type:varchar(30);comment:'邮箱';not null"`
	Password   string    `gorm:"type:varchar(128);comment:'密码';not null"`
	Avatar     string    `gorm:"type:varchar(255);comment:'头像'"`
	SpaceCover string    `gorm:"type:varchar(255);comment:'空间封面'"`
	Gender     int       `gorm:"size:1;default:0;comment:'性别:0未知、1男、3女'"`
	Birthday   time.Time `gorm:"default:'1970-01-01';comment:'生日'"`
	Sign       string    `gorm:"type:varchar(50);comment:'个性签名';default:'这个人很懒，什么都没有留下'"`
	ClientIp   string    `gorm:"type:varchar(20);comment:'客户端IP'"`
	Status     string    `gorm:"type:char(1);default:'0';commment:'账号状态:0正常、1停用、2删除'"`
	Role       int       `gorm:"size:1;default:0;commment:'角色身份:0用户、1审核、2管理、3超管'"`
}

func (table *User) TableName() string {
	return "user"
}
