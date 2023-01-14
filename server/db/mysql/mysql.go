package mysql

import (
	"fmt"

	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"kuukaa.fun/leaf/domain/model"
)

var mysqlClient *gorm.DB

func Init() {
	var err error
	dns := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?%s",
		viper.GetString("mysql.username"),
		viper.GetString("mysql.password"),
		viper.GetString("mysql.host"),
		viper.GetString("mysql.port"),
		viper.GetString("mysql.datasource"),
		viper.GetString("mysql.param"))
	mysqlClient, err = gorm.Open(mysql.Open(dns), &gorm.Config{})
	if err != nil {
		// TODO 数据库连接失败处理
		zap.L().Error("mysql连接失败")
		return
	}
	zap.L().Info("mysql连接成功")
}

func GetMysqlClient() *gorm.DB {
	return mysqlClient
}

func InitTables() {
	mysqlClient.AutoMigrate(&model.User{})
	mysqlClient.AutoMigrate(&model.Partition{})
	mysqlClient.AutoMigrate(&model.Video{})
	mysqlClient.AutoMigrate(&model.Resource{})
	mysqlClient.AutoMigrate(&model.Collection{})
	mysqlClient.AutoMigrate(&model.Follow{})
	mysqlClient.AutoMigrate(&model.Announce{})
	mysqlClient.AutoMigrate(&model.Whisper{})
	mysqlClient.AutoMigrate(&model.AtMessage{})
	mysqlClient.AutoMigrate(&model.LikeMessage{})
	mysqlClient.AutoMigrate(&model.ReplyMessage{})
	mysqlClient.AutoMigrate(&model.History{})
	mysqlClient.AutoMigrate(&model.Danmaku{})
	mysqlClient.AutoMigrate(&model.Carousel{})
}
