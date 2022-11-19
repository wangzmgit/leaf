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
		viper.GetString("mysql.hostname"),
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
}
