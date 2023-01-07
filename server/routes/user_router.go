package routes

import (
	"github.com/gin-gonic/gin"
	"kuukaa.fun/leaf/api/v1"
	"kuukaa.fun/leaf/middleware"
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
		// 通过用户ID获取用户信息
		user.GET("info/other", api.GetUserInfoByID)
		// 通过用户名获取用户ID
		user.GET("uid", api.GetUserIdByName)

		//需要用户登录
		auth := user.Group("")
		auth.Use(middleware.Auth())
		{
			//刷新token
			auth.GET("token/refresh", api.RefreshAccessToken)
			//用户获取个人信息
			auth.GET("/info/get", api.GetUserInfo)
			//用户修改个人信息
			auth.POST("/info/modify", api.ModifyUserInfo)
			//用户修改空间封面图
			auth.POST("/cover/modify", api.ModifySpaceCover)
		}

	}

}
