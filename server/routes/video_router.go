package routes

import (
	"github.com/gin-gonic/gin"
	"kuukaa.fun/leaf/api/v1"
	"kuukaa.fun/leaf/middleware"
)

func CollectVideoRoutes(r *gin.RouterGroup) {
	video := r.Group("video")
	{
		// 获取视频详情
		video.GET("get", api.GetVideoByID)
		// 获取视频列表
		video.GET("list", api.GetVideoList)
		// 获取推荐视频列表
		video.GET("recommended", api.GetRecommendedVideo)
		// 搜索视频
		video.GET("search", api.SearchVideo)
		// 获取用户视频
		video.GET("user/get", api.GetVideoListByUid)
		//在线人数连接
		video.GET("online/ws", api.GetRoomConnect)

		//需要用户登录
		auth := video.Group("")
		auth.Use(middleware.Auth())
		{
			// 上传视频信息
			auth.POST("info/upload", api.UploadVideoInfo)
			// 修改视频信息
			auth.POST("info/modify", api.ModifyVideoInfo)
			// 获取视频状态
			auth.GET("status", api.GetVideoStatus)
			// 提交审核
			auth.POST("review/submit", api.SubmitReview)
			// 删除视频
			auth.POST("delete", api.DeleteVideo)
			// 获取收藏视频
			auth.GET("collect", api.GetCollectVideo)
			// 获取上传的视频
			auth.GET("upload/get", api.GetUploadVideoList)
		}

		manage := video.Group("manage")
		manage.Use(middleware.Auth())
		{
			// 管理员获取视频列表
			manage.GET("list", api.AdminGetVideoList)
			// 管理员搜索视频
			manage.GET("search", api.AdminSearchVideo)
			// 管理员删除视频
			manage.POST("delete", api.AdminDeleteVideo)
			// 获取审核列表
			manage.GET("review/list", api.GetReviewVideoList)
			// 获取待审核视频资源
			manage.GET("review/resource/list", api.GetReviewVideoByID)
			// 审核视频
			manage.POST("review/video", api.ReviewVideo)
			// 审核资源
			manage.POST("review/resource", api.ReviewResource)
		}
	}
}
