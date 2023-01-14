package routes

import (
	"github.com/gin-gonic/gin"
	"kuukaa.fun/leaf/api/v1"
	"kuukaa.fun/leaf/middleware"
)

func CollectCarouselRoutes(route *gin.RouterGroup) {
	carousel := route.Group("/carousel")
	{
		//获取轮播图
		carousel.GET("/get", api.GetCarousel)

		auth := carousel.Group("")
		auth.Use(middleware.Auth())
		{
			//添加轮播图
			auth.POST("/add", api.AddCarousel)
			//删除轮播图
			auth.POST("/delete", api.DeleteCarousel)
		}
	}
}
