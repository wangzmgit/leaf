package initialize

import (
	"github.com/spf13/viper"
	"github.com/wangzmgit/unioss"
	"go.uber.org/zap"
)

func Oss() {
	config := unioss.Config{
		KeyID:     viper.GetString("oss.key_id"),
		KeySecret: viper.GetString("oss.key_secret"),
		Bucket:    viper.GetString("oss.bucket"),
		Endpoint:  viper.GetString("oss.endpoint"),
		AppID:     viper.GetString("oss.app_id"),
		Region:    viper.GetString("oss.region"),
		Domain:    viper.GetString("oss.domain"),
		Private:   viper.GetBool("oss.private"),
	}

	var err error
	switch viper.GetString("oss.type") {
	case unioss.ALIYUN:
		err = unioss.NewStorage(unioss.ALIYUN, config)
	case unioss.QINIU:
		err = unioss.NewStorage(unioss.QINIU, config)
	case unioss.TENCENT:
		err = unioss.NewStorage(unioss.TENCENT, config)
	}
	if err != nil {
		zap.L().Error("初始化OSS-" + viper.GetString("oss.type") + "失败，err: " + err.Error())
	}
}
