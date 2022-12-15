package routes

import (
	"github.com/gin-gonic/gin"
	"kuukaa.fun/leaf/api/v1"
	"kuukaa.fun/leaf/middleware"
)

func CollectVideoRoutes(r *gin.RouterGroup) {
	video := r.Group("video")
	{
		//获取视频详情
		video.GET("/get", api.GetVideoByID)

		//需要用户登录
		auth := video.Group("")
		auth.Use(middleware.Auth())
		{
			// 上传视频信息
			auth.POST("/info/upload", api.UploadVideoInfo)
			// 修改视频信息
			auth.POST("/info/modify", api.ModifyVideoInfo)
			// 获取视频状态
			auth.GET("/status", api.GetVideoStatus)
		}
	}
}
