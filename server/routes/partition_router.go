package routes

import (
	"github.com/gin-gonic/gin"
	"kuukaa.fun/leaf/api/v1"
	"kuukaa.fun/leaf/middleware"
)

func CollectPartitionRoutes(route *gin.RouterGroup) {
	partition := route.Group("/partition")
	{
		//获取分区列表
		partition.GET("/get", api.GetPartitionList)

		auth := partition.Group("")
		auth.Use(middleware.Auth())
		{
			//添加分区
			auth.POST("/add", api.AddPartition)
			//删除分区
			auth.POST("/delete", api.DeletePartition)
		}
	}
}
