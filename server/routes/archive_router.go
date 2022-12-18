package routes

import (
	"github.com/gin-gonic/gin"
	"kuukaa.fun/leaf/api/v1"
	"kuukaa.fun/leaf/middleware"
)

func CollectArchiveRoutes(r *gin.RouterGroup) {
	archive := r.Group("archive")
	{
		auth := archive.Group("")
		auth.Use(middleware.Auth())
		{
			// 点赞
			auth.POST("like", api.Like)
			// 取消赞
			auth.POST("cancel/like", api.CancelLike)
			// 是否点赞
			auth.GET("has/like", api.HasLike)
		}
	}
}
