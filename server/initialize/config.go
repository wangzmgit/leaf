package initialize

import (
	"github.com/spf13/viper"
	"github.com/wangzmgit/jigsaw"
)

func ConfigFiles() {
	viper.SetConfigName("application")
	viper.AddConfigPath("./config")
	err := viper.ReadInConfig()
	if err != nil {
		panic("配置文件读取失败")
	}

	j := jigsaw.New()
	j.SetBgDir("./static/images/bg/")
	j.SetMaskPath("./static/images/mask.png")
}
