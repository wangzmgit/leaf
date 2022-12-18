package service

import (
	"gorm.io/gorm"
	"kuukaa.fun/leaf/db/mongodb"
	"kuukaa.fun/leaf/db/mysql"
)

var mysqlClient *gorm.DB
var mongoClient *mongodb.MongoClient

func InitMysqlClient() {
	mysqlClient = mysql.GetMysqlClient()
}

func InitMongoClient() {
	mongoClient = mongodb.GetMongoClient()
}
