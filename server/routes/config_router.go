package routes

import (
	"github.com/gin-gonic/gin"
	"kuukaa.fun/leaf/api/v1"
	"kuukaa.fun/leaf/middleware"
)

func CollectConfigRoutes(r *gin.RouterGroup) {
	config := r.Group("config")
	config.Use(middleware.Auth())
	{
		// 获取邮箱配置
		config.GET("email/get", api.GetEmailConfig)
		// 修改邮箱配置
		config.POST("email/set", api.SetEmailConfig)
		// 获取存储配置
		config.GET("storage/get", api.GetStorageConfig)
		// 修改存储配置
		config.POST("storage/set", api.SetStorageConfig)
		// 获取其他配置
		config.GET("other/get", api.GetOtherConfig)
		// 修改其他配置
		config.POST("other/set", api.SetOtherConfig)
	}

}
