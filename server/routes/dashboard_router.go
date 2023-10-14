package routes

import (
	"github.com/gin-gonic/gin"
	"kuukaa.fun/leaf/api/v1"
	"kuukaa.fun/leaf/middleware"
)

func CollectDashboardRoutes(r *gin.RouterGroup) {
	dashboard := r.Group("dashboard")
	{

		//需要用户登录
		auth := dashboard.Group("")
		auth.Use(middleware.Auth())
		{
			// 获取卡片数据
			auth.GET("card/data", api.GetCardData)
			// 获取趋势数据
			auth.GET("trend", api.GetTrendData)
			// 获取视频分区数据
			auth.GET("partition", api.GetPartitionData)
		}
	}
}
