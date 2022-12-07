package routes

import (
	"github.com/gin-gonic/gin"
	"kuukaa.fun/leaf/api/v1"
)

func CollectCaptchaRoutes(r *gin.RouterGroup) {
	captcha := r.Group("captcha")
	{
		// 获取滑块验证
		captcha.GET("get", api.GetSliderCaptcha)
		// 验证滑块
		captcha.POST("validate", api.ValidateSlider)
	}

}
