package service

import (
	"gorm.io/gorm"
	"kuukaa.fun/leaf/db/mysql"
)

var mysqlClient *gorm.DB

func InitMysqlClient() {
	mysqlClient = mysql.GetMysqlClient()
}
