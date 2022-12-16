package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"kuukaa.fun/leaf/logger"
	"kuukaa.fun/leaf/middleware"
)

func InitRouter() {
	// gin 模式
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	r.Use(logger.GinLogger, logger.GinRecovery(true))

	// 设置信任网络 []string nil 为不计算，避免性能消耗，上线应当设置
	_ = r.SetTrustedProxies(nil)

	//跨域中间件
	r.Use(middleware.CORS())

	// 收集添加路由
	CollectRoutes(r)

	// 运行
	r.Run(":9000")
}

func CollectRoutes(r *gin.Engine) *gin.Engine {

	v1 := r.Group("/api/v1")
	{
		// 用户相关路由
		CollectUserRoutes(v1)
		// 人机验证相关路由
		CollectCaptchaRoutes(v1)
		// 文件上传相关路由
		CollectUploadRoutes(v1)
		// 分区相关路由
		CollectPartitionRoutes(v1)
		// 视频相关路由
		CollectVideoRoutes(v1)
		// 视频资源相关录音
		CollectResourceRoutes(v1)
	}

	//获取静态文件
	r.StaticFS("/api/img", http.Dir("./upload/image"))
	r.StaticFS("/api/video", http.Dir("./upload/video"))

	return r
}
