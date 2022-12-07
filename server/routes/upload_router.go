package routes

import (
	"github.com/gin-gonic/gin"
	"kuukaa.fun/leaf/api/v1"
	"kuukaa.fun/leaf/middleware"
)

func CollectUploadRoutes(r *gin.RouterGroup) {
	upload := r.Group("upload")
	{
		// 用户注册
		uploadAuth := upload.Group("")
		uploadAuth.Use(middleware.Auth())
		{
			uploadAuth.POST("image", api.UploadImg)
		}
	}
}
