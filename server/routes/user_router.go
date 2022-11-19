package routes

import (
	"github.com/gin-gonic/gin"
	"kuukaa.fun/leaf/api/v1"
)

func CollectUserRoutes(r *gin.RouterGroup) {
	user := r.Group("user")
	{
		// 用户注册
		user.POST("register", api.Register)
		// 用户登录
		user.POST("login", api.Login)
		// 获取邮箱验证码
		user.POST("emailcode", api.SendEmailCode)

		// //需要用户登录
		// userAuth := user.Group("")
		// userAuth.Use(middleware.Auth())
		// {
		// 	userAuth.GET("/info/get", controller.UserInfo)       //用户获取个人信息
		// 	userAuth.POST("/info/modify", controller.ModifyInfo) //用户修改个人信息
		// }

	}

}
