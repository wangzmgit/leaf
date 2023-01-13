package routes

import (
	"github.com/gin-gonic/gin"
	"kuukaa.fun/leaf/api/v1"
	"kuukaa.fun/leaf/middleware"
)

func CollectUploadRoutes(r *gin.RouterGroup) {
	upload := r.Group("upload")
	{
		auth := upload.Group("")
		auth.Use(middleware.Auth())
		{
			auth.POST("image", api.UploadImg)
			auth.POST("video/:vid", api.UploadVideo)
		}
	}
}
