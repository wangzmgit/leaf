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
		// 用户登录(密码)
		user.POST("login", api.Login)
		// 用户登录(邮箱)
		user.POST("login/email", api.EmailLogin)
		// 获取邮箱验证码
		user.POST("code/email", api.SendRegisterEmailCode)
		// 通过用户ID获取用户信息
		user.GET("info/other", api.GetUserInfoByID)
		// 通过用户名获取用户ID
		user.GET("uid", api.GetUserIdByName)
		// 验证修改密码的用户
		user.GET("resetpwd/check", api.ResetPwdCheck)
		// 修改密码
		user.POST("pwd/modify", api.ModifyPwd)

		//需要用户登录
		auth := user.Group("")
		auth.Use(middleware.Auth())
		{
			//用户获取个人信息
			auth.GET("/info/get", api.GetUserInfo)
			//用户修改个人信息
			auth.POST("/info/modify", api.ModifyUserInfo)
			//用户修改空间封面图
			auth.POST("/cover/modify", api.ModifySpaceCover)
		}

		manage := auth.Group("manage")
		{
			// 管理员获取用户列表
			manage.GET("list", api.AdminGetUserList)
			// 管理员搜索用户信息
			manage.GET("search", api.AdminSearchUserInfo)
			// 管理员更新用户信息
			manage.POST("modify", api.AdminModifyUserInfo)
			// 管理员更新用户权限信息
			manage.POST("role/modify", api.AdminModifyUserRole)
			// 管理员删除用户
			manage.POST("delete", api.AdminDeleteUser)
		}

	}

}
