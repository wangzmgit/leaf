package routes

import (
	"github.com/gin-gonic/gin"
	"kuukaa.fun/leaf/api/v1"
	"kuukaa.fun/leaf/middleware"
)

func CollectHistoryRoutes(r *gin.RouterGroup) {
	history := r.Group("history")
	{
		//需要用户登录
		auth := history.Group("")
		auth.Use(middleware.Auth())
		{
			// 记录历史记录
			auth.POST("add", api.AddHistory)
			// 获取历史记录
			auth.GET("video/get", api.GetHistoryVideo)
			// 获取播放进度
			auth.GET("progress/get", api.GetHistoryProgress)
		}
	}
}
