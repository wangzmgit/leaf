package api

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
	"kuukaa.fun/leaf/cache"
	"kuukaa.fun/leaf/domain/dto"
	"kuukaa.fun/leaf/domain/resp"
	"kuukaa.fun/leaf/domain/valid"
	"kuukaa.fun/leaf/service"
	"kuukaa.fun/leaf/util/jwt"
)

func Register(ctx *gin.Context) {
	// 获取参数
	var registerDTO dto.RegisterDTO
	if err := ctx.Bind(&registerDTO); err != nil {
		resp.Response(ctx, resp.RequestParamError, "", nil)
		zap.L().Error("请求参数有误")
		return
	}

	// 参数校验
	if !valid.Email(registerDTO.Email) { // 邮箱格式验证
		resp.Response(ctx, resp.RequestParamError, valid.EMAIL_ERROR, nil)
		zap.L().Error(valid.EMAIL_ERROR)
		return
	}

	if !valid.Password(registerDTO.Password) { // 密码格式验证
		resp.Response(ctx, resp.RequestParamError, valid.PASSWORD_ERROR, nil)
		zap.L().Error(valid.PASSWORD_ERROR)
		return
	}

	if !valid.EmailCode(registerDTO.Code) { //邮箱验证码格式验证
		resp.Response(ctx, resp.RequestParamError, valid.EMAIL_CODE_ERROR, nil)
		zap.L().Error(valid.EMAIL_CODE_ERROR)
		return
	}

	if service.SelectUserByEmail(registerDTO.Email).ID != 0 {
		resp.Response(ctx, resp.EmailExistError, "", nil)
		zap.L().Error("邮箱已存在")
		return
	}

	if cache.GetEmailCode(registerDTO.Email) != registerDTO.Code { // 验证邮箱验证码
		resp.Response(ctx, resp.EmailCodeError, "", nil)
		zap.L().Error("邮箱验证错误")
		return
	}

	// 保存到数据库
	user := dto.RegisterToUser(registerDTO)
	service.InsertUser(user)

	// 删除邮箱验证码
	cache.DelEmailCode(user.Email)

	// 返回
	resp.OK(ctx, "ok", nil)
}

func Login(ctx *gin.Context) {
	// 获取参数
	var loginDTO dto.LoginDTO
	if err := ctx.Bind(&loginDTO); err != nil {
		resp.Response(ctx, resp.RequestParamError, "", nil)
		zap.L().Error("请求参数有误")
		return
	}

	// 参数校验
	if !valid.Email(loginDTO.Email) { // 邮箱格式验证
		resp.Response(ctx, resp.RequestParamError, valid.EMAIL_ERROR, nil)
		zap.L().Error(valid.EMAIL_ERROR)
		return
	}

	if !valid.Password(loginDTO.Password) { // 密码格式验证
		resp.Response(ctx, resp.RequestParamError, valid.PASSWORD_ERROR, nil)
		zap.L().Error(valid.PASSWORD_ERROR)
		return
	}

	// 如果人机验证通过，删除登录尝试次数
	if cache.GetCaptchaStatus(loginDTO.Email) == 1 {
		// 删除登录尝试次数
		cache.DelLoginTryCount(loginDTO.Email)
		// 删除人机验证状态
		cache.DelCaptchaStatus(loginDTO.Email)
	}

	// 读取登录尝试次数，超过3次进行滑块验证
	loginTryCount := cache.GetLoginTryCount(loginDTO.Email)
	if loginTryCount >= 3 {
		resp.Response(ctx, resp.Captcha, "", nil)
		zap.L().Info("需要人机验证")
		return
	}

	// 读取数据库
	user := service.SelectUserByEmail(loginDTO.Email)
	// 验证账号密码
	passwordError := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginDTO.Password))
	if passwordError != nil {
		resp.Response(ctx, resp.UsernamePasswordNotMatchError, "", nil)
		zap.L().Info("用户名密码不匹配")
		// 记录登录尝试次数
		loginTryCount++
		cache.SetLoginTryCount(loginDTO.Email, loginTryCount)
		return
	}

	// 生成token
	var err error
	var token string
	if token, err = jwt.GenerateToken(user.ID); err != nil {
		resp.Response(ctx, resp.Error, "token生成失败", nil)
		zap.L().Error("token生成失败")
		return
	}

	// 存入缓存
	cache.SetToken(user.ID, token)

	// 返回给前端
	resp.OK(ctx, "ok", gin.H{"token": token})
}

// 邮箱登录
func EmailLogin(ctx *gin.Context) {
	// 获取参数
	var loginDTO dto.EmailLoginDTO
	if err := ctx.Bind(&loginDTO); err != nil {
		resp.Response(ctx, resp.RequestParamError, "", nil)
		zap.L().Error("请求参数有误")
		return
	}

	// 参数校验
	if !valid.Email(loginDTO.Email) { // 邮箱格式验证
		resp.Response(ctx, resp.RequestParamError, valid.EMAIL_ERROR, nil)
		zap.L().Error(valid.EMAIL_ERROR)
		return
	}

	// 如果人机验证通过，删除登录尝试次数
	if cache.GetCaptchaStatus(loginDTO.Email) == 1 {
		// 删除登录尝试次数
		cache.DelLoginTryCount(loginDTO.Email)
		// 删除人机验证状态
		cache.DelCaptchaStatus(loginDTO.Email)
	}

	loginTryCount := cache.GetLoginTryCount(loginDTO.Email)
	if loginTryCount >= 3 {
		resp.Response(ctx, resp.Captcha, "", nil)
		zap.L().Info("需要人机验证")
		return
	}

	// 验证邮箱验证码
	if cache.GetEmailCode(loginDTO.Email) != loginDTO.Code {
		resp.Response(ctx, resp.EmailCodeError, "", nil)
		zap.L().Error("邮箱验证错误")
		return
	}

	// 读取数据库
	user := service.SelectUserByEmail(loginDTO.Email)
	if user.ID == 0 {
		resp.Response(ctx, resp.UserNotExistError, "", nil)
		zap.L().Error("用户不存在")
		return
	}

	// 生成token
	var err error
	var token string
	if token, err = jwt.GenerateToken(user.ID); err != nil {
		resp.Response(ctx, resp.Error, "token生成失败", nil)
		zap.L().Error("token生成失败")
		return
	}

	// 存入缓存
	cache.SetToken(user.ID, token)

	// 返回给前端
	resp.OK(ctx, "ok", gin.H{"token": token})
}
