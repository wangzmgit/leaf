package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"kuukaa.fun/leaf/domain/dto"
	"kuukaa.fun/leaf/domain/resp"
	"kuukaa.fun/leaf/domain/vo"
	"kuukaa.fun/leaf/initialize"
)

// 获取邮箱配置信息
func GetEmailConfig(ctx *gin.Context) {
	resp.OK(ctx, "ok", gin.H{"config": vo.EmailConfigVO{
		User:      viper.GetString("mail.user"),
		Pass:      viper.GetString("mail.pass"),
		Host:      viper.GetString("mail.host"),
		Port:      viper.GetInt("mail.port"),
		Addresser: viper.GetString("mail.addresser"),
	}})
}

// 配置邮箱
func SetEmailConfig(ctx *gin.Context) {
	var emailConfigDTO dto.EmailConfigDTO
	if err := ctx.Bind(&emailConfigDTO); err != nil {
		resp.Response(ctx, resp.RequestParamError, "", nil)
		zap.L().Error("请求参数有误")
		return
	}

	viper.Set("mail.user", emailConfigDTO.User)
	viper.Set("mail.host", emailConfigDTO.Host)
	viper.Set("mail.port", emailConfigDTO.Port)
	viper.Set("mail.addresser", emailConfigDTO.Addresser)

	if len(emailConfigDTO.Pass) != 0 {
		viper.Set("mail.pass", emailConfigDTO.Pass)
	}

	viper.WriteConfig()

	// 返回给前端
	resp.OK(ctx, "ok", nil)
}

// 获取存储配置信息
func GetStorageConfig(ctx *gin.Context) {
	resp.OK(ctx, "ok", gin.H{"config": vo.StorageConfigVO{
		MaxImgSize:         viper.GetInt("file.max_img_size"),
		MaxVideoSize:       viper.GetInt("file.max_video_size"),
		VideoAdaptationIos: viper.GetBool("file.video_adaptation_ios"),

		Type:      viper.GetString("oss.type"),
		KeyID:     viper.GetString("oss.key_id"),
		KeySecret: viper.GetString("oss.key_secret"),
		Bucket:    viper.GetString("oss.bucket"),
		Endpoint:  viper.GetString("oss.endpoint"),
		AppID:     viper.GetString("oss.app_id"),
		Region:    viper.GetString("oss.region"),
		Domain:    viper.GetString("oss.domain"),
		Private:   viper.GetBool("oss.private"),
	}})
}

// 修改存储配置
func SetStorageConfig(ctx *gin.Context) {
	var storageConfigDTO dto.StorageConfigDTO
	if err := ctx.Bind(&storageConfigDTO); err != nil {
		resp.Response(ctx, resp.RequestParamError, "", nil)
		zap.L().Error("请求参数有误")
		return
	}

	fmt.Println(storageConfigDTO)

	viper.Set("file.max_img_size", storageConfigDTO.MaxImgSize)
	viper.Set("file.max_video_size", storageConfigDTO.MaxVideoSize)
	viper.Set("file.video_adaptation_ios", storageConfigDTO.VideoAdaptationIos)

	viper.Set("oss.type", storageConfigDTO.Type)
	viper.Set("oss.key_id", storageConfigDTO.KeyID)
	viper.Set("oss.bucket", storageConfigDTO.Bucket)
	viper.Set("oss.endpoint", storageConfigDTO.Endpoint)
	viper.Set("oss.app_id", storageConfigDTO.AppID)
	viper.Set("oss.region", storageConfigDTO.Region)
	viper.Set("oss.domain", storageConfigDTO.Domain)
	viper.Set("oss.private", storageConfigDTO.Private)

	if len(storageConfigDTO.KeySecret) != 0 {
		viper.Set("oss.key_secret", storageConfigDTO.KeySecret)
	}

	viper.WriteConfig()

	// 重新初始化OSS
	if storageConfigDTO.Type != "local" {
		initialize.Oss()
	}

	// 返回给前端
	resp.OK(ctx, "ok", nil)
}

// 获取其他配置信息
func GetOtherConfig(ctx *gin.Context) {
	resp.OK(ctx, "ok", gin.H{"config": vo.OtherConfigVO{
		AllowOrigin: viper.GetString("cors.allow_origin"),
		Prefix:      viper.GetString("user.prefix"),
	}})
}

// 修改其他配置信息
func SetOtherConfig(ctx *gin.Context) {
	var otherConfigDTO dto.OtherConfigDTO
	if err := ctx.Bind(&otherConfigDTO); err != nil {
		resp.Response(ctx, resp.RequestParamError, "", nil)
		zap.L().Error("请求参数有误")
		return
	}

	viper.Set("cors.allow_origin", otherConfigDTO.AllowOrigin)
	viper.Set("user.prefix", otherConfigDTO.Prefix)

	viper.WriteConfig()

	// 返回给前端
	resp.OK(ctx, "ok", nil)
}
