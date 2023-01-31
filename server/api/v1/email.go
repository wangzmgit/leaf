package api

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"kuukaa.fun/leaf/cache"
	"kuukaa.fun/leaf/domain/dto"
	"kuukaa.fun/leaf/domain/resp"
	"kuukaa.fun/leaf/domain/valid"
	"kuukaa.fun/leaf/util/mail"
	"kuukaa.fun/leaf/util/random"
)

func SendRegisterEmailCode(ctx *gin.Context) {
	// 获取参数
	var codeDTO dto.EmailDTO
	if err := ctx.Bind(&codeDTO); err != nil {
		resp.Response(ctx, resp.RequestParamError, "", nil)
		zap.L().Error("请求参数有误")
		return
	}

	// 参数校验
	if !valid.Email(codeDTO.Email) { // 邮箱格式验证
		resp.Response(ctx, resp.RequestParamError, valid.EMAIL_ERROR, nil)
		zap.L().Error(valid.EMAIL_ERROR)
		return
	}

	// 如果未进行人机验证
	if cache.GetCaptchaStatus(codeDTO.Email) == 0 {
		resp.Response(ctx, resp.Captcha, "", nil)
		zap.L().Info("需要人机验证")
		return
	}

	// 生成code
	code := random.GenerateNumberCode(4)

	// 发送code
	if viper.GetBool("mail.debug") {
		// 验证码debug模式不发送邮件
		zap.L().Debug("邮箱:" + codeDTO.Email + ",验证码:" + code)
	} else {
		if err := mail.SendCaptcha(codeDTO.Email, code); err != nil {
			resp.Response(ctx, resp.SendMailError, "邮箱验证码发送失败", nil)
			zap.L().Error("邮箱验证码发送失败" + err.Error())
			return
		}
	}

	// code放入缓存
	cache.SetEmailCode(codeDTO.Email, code)
	// 删除人机验证状态
	cache.DelCaptchaStatus(codeDTO.Email)

	resp.OK(ctx, "ok", nil)
}
