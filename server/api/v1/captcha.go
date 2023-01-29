package api

import (
	"github.com/gin-gonic/gin"
	"github.com/wangzmgit/jigsaw"
	"go.uber.org/zap"
	"kuukaa.fun/leaf/cache"
	"kuukaa.fun/leaf/domain/dto"
	"kuukaa.fun/leaf/domain/resp"
	"kuukaa.fun/leaf/domain/valid"
	"kuukaa.fun/leaf/domain/vo"
)

// 获取滑块验证码
func GetSliderCaptcha(ctx *gin.Context) {
	//获取参数
	email := ctx.DefaultQuery("email", "")

	// 参数校验
	if !valid.Email(email) { // 邮箱格式验证
		resp.Response(ctx, resp.RequestParamError, valid.EMAIL_ERROR, nil)
		zap.L().Error(valid.EMAIL_ERROR)
		return
	}

	// 生成滑块验证
	slider, bg, x, y, err := jigsaw.Create()
	if err != nil {
		zap.L().Error("滑块验证资源生成失败")
		resp.Response(ctx, resp.Error, "", nil)
	}
	// 保存x坐标到缓存
	cache.SetSliderX(email, x)
	resp.OK(ctx, "ok", gin.H{"slider_captcha": vo.CaptchaVO{
		BgImg:     bg,
		SliderImg: slider,
		Y:         y,
	}})
}

// 验证滑块
func ValidateSlider(ctx *gin.Context) {
	// 获取参数
	var validateSliderDTO dto.ValidateSliderDTO
	if err := ctx.Bind(&validateSliderDTO); err != nil {
		resp.Response(ctx, resp.RequestParamError, "", nil)
		zap.L().Error("请求参数有误")
		return
	}

	//获取缓存中的x并删除缓存
	x := cache.GetSliderX(validateSliderDTO.Email)
	cache.DelSlider(validateSliderDTO.Email)
	//与用户拖动位置对比
	if validateSliderDTO.X > x-3 && validateSliderDTO.X < x+3 {
		cache.SetCaptchaStatus(validateSliderDTO.Email, 1)
		resp.OK(ctx, "ok", nil)
		return
	}

	zap.L().Error("滑块验证失败")
	resp.Response(ctx, resp.RequestParamError, "", nil)
}
