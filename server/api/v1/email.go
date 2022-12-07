package api

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"kuukaa.fun/leaf/cache"
	"kuukaa.fun/leaf/domain/dto"
	"kuukaa.fun/leaf/domain/resp"
	"kuukaa.fun/leaf/util/random"
)

func SendEmailCode(ctx *gin.Context) {
	// 获取参数
	var codeDTO dto.EmailDTO
	if err := ctx.Bind(&codeDTO); err != nil {
		resp.Response(ctx, resp.RequestParamError, "", nil)
		zap.L().Error("请求参数有误")
		return
	}
	email := codeDTO.Email

	zap.L().Debug(email)

	// 生成code
	code := random.GenerateNumberCode(4)

	zap.L().Debug(code)

	// 发送code
	// if err := mail.SendCaptcha(email, code); err != nil {
	// 	resp.Response(ctx, resp.SendMailError, "邮箱验证码发送失败", nil)
	// 	zap.L().Error("邮箱验证码发送失败")
	// 	return
	// }

	// code放入缓存
	cache.SetEmailCode(email, code)

	resp.OK(ctx, "验证码已发送到您的邮箱", nil)
}
