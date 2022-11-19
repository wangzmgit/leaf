package routes

import (
	"github.com/gin-gonic/gin"
	"kuukaa.fun/leaf/logger"
)

func InitRouter() {
	// gin 模式
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	r.Use(logger.GinLogger, logger.GinRecovery(true))
	
	// 设置信任网络 []string nil 为不计算，避免性能消耗，上线应当设置
	_ = r.SetTrustedProxies(nil)

	// 收集添加路由
	CollectRoutes(r)

	// 运行
	r.Run(":8081")
}

func CollectRoutes(r *gin.Engine) *gin.Engine {

	v1 := r.Group("/api/v1")
	{
		// 用户相关路由
		CollectUserRoutes(v1)
	}

	return r
}
