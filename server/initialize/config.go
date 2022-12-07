package initialize

import (
	"github.com/spf13/viper"
)

func ConfigFiles() {
	viper.SetConfigName("application")
	viper.AddConfigPath("./config")
	err := viper.ReadInConfig()
	if err != nil {
		panic("配置文件读取失败")
	}
}
