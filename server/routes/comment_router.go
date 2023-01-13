package routes

import (
	"github.com/gin-gonic/gin"
	"kuukaa.fun/leaf/api/v1"
	"kuukaa.fun/leaf/middleware"
)

func CollectCommentRoutes(r *gin.RouterGroup) {
	comment := r.Group("comment")
	{
		// 获取评论
		comment.GET("get", api.GetComment)
		// 获取回复
		comment.GET("reply/get", api.GetReply)

		auth := comment.Group("")
		auth.Use(middleware.Auth())
		{
			// 添加评论回复
			auth.POST("add", api.Comment)
			// 添加评论回复
			auth.POST("reply/add", api.Reply)
			// 删除评论
			auth.POST("delete", api.DeleteComment)
			// 删除回复
			auth.POST("reply/delete", api.DeleteReply)
		}
	}
}
