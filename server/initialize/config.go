package initialize

import (
	"time"

	"github.com/spf13/viper"
	"kuukaa.fun/leaf/util/random"
)

func ConfigFiles() {
	viper.SetConfigName("application")
	viper.AddConfigPath("./config")
	err := viper.ReadInConfig()
	if err != nil {
		panic("配置文件读取失败")
	}

	if viper.GetString("security.access_jwt_secret") == "" {
		viper.Set("server.access_jwt_secret", random.GenerateNumberCode(16))
	}

	time.Sleep(time.Millisecond)

	if viper.GetString("security.refresh_jwt_secret") == "" {
		viper.Set("server.refresh_jwt_secret", random.GenerateNumberCode(16))
	}

	viper.WriteConfig()
}
