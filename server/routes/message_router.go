package routes

import (
	"github.com/gin-gonic/gin"
	"kuukaa.fun/leaf/api/v1"
	"kuukaa.fun/leaf/middleware"
)

func CollectMessageRoutes(r *gin.RouterGroup) {
	message := r.Group("message")
	{
		// 公告
		announce := message.Group("announce")
		{
			// 获取公告
			announce.GET("get", api.GetAnnounce)
			// 获取最新重要公告
			announce.GET("important/get", api.GetImportantAnnounce)

			//需要用户登录
			announceAuth := announce.Group("")
			announceAuth.Use(middleware.Auth())
			{
				// 添加公告
				announceAuth.POST("add", api.AddAnnounce)
				// 删除公告
				announceAuth.POST("delete", api.DeleteAnnounce)
			}
		}

		// 点赞
		likeAuth := message.Group("like")
		likeAuth.Use(middleware.Auth())
		{
			likeAuth.GET("get", api.GetLikeMessage)
		}

		// 点赞
		atAuth := message.Group("at")
		atAuth.Use(middleware.Auth())
		{
			atAuth.GET("get", api.GetAtMessage)
		}

		// 回复
		replyAuth := message.Group("reply")
		replyAuth.Use(middleware.Auth())
		{
			replyAuth.GET("get", api.GetReplyMessage)
		}

		// 私信
		whisper := message.Group("whisper")
		{
			// websocket
			wsAuth := whisper.Group("")
			wsAuth.Use(middleware.WsAuth())
			{
				// 连接消息ws
				wsAuth.GET("/ws", api.GetWhisperConnect)
			}

			whisperAuth := whisper.Group("")
			whisperAuth.Use(middleware.Auth())
			{
				// 获取消息列表
				whisperAuth.GET("/list", api.GetWhisperList)
				// 获取消息详情
				whisperAuth.GET("/details", api.GetMessageDetails)
				// 发送消息
				whisperAuth.POST("/send", api.SendWhisper)
				// 已读消息
				whisperAuth.POST("/read", api.ReadWhisper)
			}
		}

	}
}
