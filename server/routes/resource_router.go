package routes

import (
	"github.com/gin-gonic/gin"
	"kuukaa.fun/leaf/api/v1"
	"kuukaa.fun/leaf/middleware"
)

func CollectResourceRoutes(r *gin.RouterGroup) {
	resource := r.Group("resource")
	{
		//需要用户登录
		auth := resource.Group("")
		auth.Use(middleware.Auth())
		{
			auth.POST("title/modify", api.ModifyResourceTitle)
			auth.POST("delete", api.DeleteResource)
		}
	}
}
