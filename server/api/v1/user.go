package api

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"kuukaa.fun/leaf/cache"
	"kuukaa.fun/leaf/domain/dto"
	"kuukaa.fun/leaf/domain/resp"
	"kuukaa.fun/leaf/domain/valid"
	"kuukaa.fun/leaf/domain/vo"
	"kuukaa.fun/leaf/service"
	"kuukaa.fun/leaf/util/convert"
)

func GetUserInfo(ctx *gin.Context) {
	userId := ctx.GetUint("userId")
	userInfo := service.GetUserInfo(userId)

	resp.OK(ctx, "", gin.H{"user_info": vo.ToUserVO(userInfo)})
}

func ModifyUserInfo(ctx *gin.Context) {
	var modifyDTO dto.ModifyUserInfoDTO
	if err := ctx.Bind(&modifyDTO); err != nil {
		resp.Response(ctx, resp.RequestParamError, "", nil)
		zap.L().Error("请求参数有误")
		return
	}

	// 参数校验
	if !valid.Name(modifyDTO.Name) {
		resp.Response(ctx, resp.RequestParamError, valid.NAME_ERROR, nil)
		zap.L().Error(valid.NAME_ERROR)
		return
	}

	userId := ctx.GetUint("userId")
	if exist, uid := service.IsNameExist(modifyDTO.Name); exist && uid != userId {
		resp.Response(ctx, resp.NameExistError, "", nil)
		zap.L().Error("用户名已存在")
		return
	}

	// 检测文件是否有效
	oldUserInfo := service.SelectUserByID(userId)
	if modifyDTO.Avatar != oldUserInfo.Avatar && cache.GetUploadImage(modifyDTO.Avatar) != userId {
		resp.Response(ctx, resp.InvalidLinkError, "", nil)
		zap.L().Error("文件链接无效")
		return
	}

	// 保存到数据库
	service.UpdateUserInfo(userId, modifyDTO)
	// 返回
	resp.OK(ctx, "ok", nil)
}

func ModifySpaceCover(ctx *gin.Context) {
	var modifyDTO dto.ModifySpaceCoverDTO
	if err := ctx.Bind(&modifyDTO); err != nil {
		resp.Response(ctx, resp.RequestParamError, "", nil)
		zap.L().Error("请求参数有误")
		return
	}

	// 参数校验
	userId := ctx.GetUint("userId")
	if cache.GetUploadImage(modifyDTO.SpaceCover) != userId {
		resp.Response(ctx, resp.InvalidLinkError, "", nil)
		zap.L().Error("文件链接无效")
		return
	}

	// 保存到数据库
	service.UpdateUserSpaceCover(userId, modifyDTO.SpaceCover)
	// 返回
	resp.OK(ctx, "ok", nil)
}

// 通过用户ID获取用户信息
func GetUserInfoByID(ctx *gin.Context) {
	uid := convert.StringToUint(ctx.DefaultQuery("uid", "0"))

	user := service.GetUserInfo(uid)

	// 返回给前端
	resp.OK(ctx, "ok", gin.H{"user": vo.ToBaseUserVO(user)})
}

// 通过用户名获取用户ID
func GetUserIdByName(ctx *gin.Context) {
	userName := ctx.Query("name")

	userId := service.SelectUserIdByName(userName)

	// 返回给前端
	resp.OK(ctx, "ok", gin.H{"uid": userId})
}

// 重置密码检查
func ResetPwdCheck(ctx *gin.Context) {
	// 获取参数
	email := ctx.Query("email")

	// 参数校验
	if !valid.Email(email) { // 邮箱格式验证
		resp.Response(ctx, resp.RequestParamError, valid.EMAIL_ERROR, nil)
		zap.L().Error(valid.EMAIL_ERROR)
		return
	}

	// 如果未进行人机验证
	if cache.GetCaptchaStatus(email) == 0 {
		resp.Response(ctx, resp.Captcha, "", nil)
		zap.L().Info("需要人机验证")
		return
	}

	// 删除人机验证状态
	cache.DelCaptchaStatus(email)

	if service.SelectUserByEmail(email).ID == 0 {
		resp.Response(ctx, resp.UserNotExistError, "", nil)
		zap.L().Error("用户不存在")
		return
	}

	// 更新验证状态
	cache.SetResetPwdCheckStatus(email, 1)

	// 返回给前端
	resp.OK(ctx, "", nil)
}

// 修改密码
func ModifyPwd(ctx *gin.Context) {
	var modifyDTO dto.ModifyPwdDTO
	if err := ctx.Bind(&modifyDTO); err != nil {
		resp.Response(ctx, resp.RequestParamError, "", nil)
		zap.L().Error("请求参数有误")
		return
	}

	// 参数校验
	if !valid.Email(modifyDTO.Email) { // 邮箱格式验证
		resp.Response(ctx, resp.RequestParamError, valid.EMAIL_ERROR, nil)
		zap.L().Error(valid.EMAIL_ERROR)
		return
	}

	if !valid.Password(modifyDTO.Password) { // 密码格式验证
		resp.Response(ctx, resp.RequestParamError, valid.PASSWORD_ERROR, nil)
		zap.L().Error(valid.PASSWORD_ERROR)
		return
	}

	if !valid.EmailCode(modifyDTO.Code) { //邮箱验证码格式验证
		resp.Response(ctx, resp.RequestParamError, valid.EMAIL_CODE_ERROR, nil)
		zap.L().Error(valid.EMAIL_CODE_ERROR)
		return
	}

	// 如果用户未验证

	if cache.GetEmailCode(modifyDTO.Email) != modifyDTO.Code { // 验证邮箱验证码
		resp.Response(ctx, resp.EmailCodeError, "", nil)
		zap.L().Error("邮箱验证错误")
		return
	}

	service.UpdateUserPwd(modifyDTO)

	// 删除验证状态
	cache.DelResetPwdCheckStatus(modifyDTO.Email)

	// 返回给前端
	resp.OK(ctx, "", nil)
}
