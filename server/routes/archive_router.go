package routes

import (
	"github.com/gin-gonic/gin"
	"kuukaa.fun/leaf/api/v1"
	"kuukaa.fun/leaf/middleware"
)

func CollectArchiveRoutes(r *gin.RouterGroup) {
	archive := r.Group("archive")
	{
		// 获取点赞收藏数据
		archive.GET("stat", api.GetArchiveStat)

		auth := archive.Group("")
		auth.Use(middleware.Auth())
		{
			// 点赞
			auth.POST("like", api.Like)
			// 取消赞
			auth.POST("cancel/like", api.CancelLike)
			// 是否点赞
			auth.GET("has/like", api.HasLike)

			// 收藏
			auth.POST("collect", api.Collect)
			// 已收藏的文件夹
			auth.GET("collect/collected", api.GetCollectedInfo)
			// 是否收藏
			auth.GET("has/collect", api.HasCollect)
		}
	}
}
