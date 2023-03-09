package main

import (
	"kuukaa.fun/leaf/cache"
	"kuukaa.fun/leaf/db/mongodb"
	"kuukaa.fun/leaf/db/mysql"
	"kuukaa.fun/leaf/initialize"
	"kuukaa.fun/leaf/logger"
	"kuukaa.fun/leaf/routes"
	"kuukaa.fun/leaf/service"
	"kuukaa.fun/leaf/util/authentication"
	"kuukaa.fun/leaf/util/cron"
)

func main() {
	// 初始化配置文件
	initialize.ConfigFiles()
	// 初始化日志
	logger.InitLogger()
	// 初始化滑块验证码生成
	initialize.Jigsaw()
	// 初始化OSS
	initialize.Oss()
	// 初始化casbin
	authentication.InitCasbin()
	// 初始化mysql
	mysql.Init()
	// 初始化数据库表
	mysql.InitTables()
	// 初始化MongoDB
	mongodb.Init()
	// 初始化缓存
	cache.Init()
	// 初始化mysql客户端
	service.InitMysqlClient()
	// 初始化mongodb客户端
	service.InitMongoClient()
	// 开启定时任务
	go cron.Init()

	// fmt.Println(jwt.GenerateAccessToken(1))

	// 初始化路由
	routes.InitRouter()
}
