package api

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"kuukaa.fun/leaf/domain/dto"
	"kuukaa.fun/leaf/domain/model"
	"kuukaa.fun/leaf/domain/resp"
	"kuukaa.fun/leaf/domain/vo"
	"kuukaa.fun/leaf/service"
	"kuukaa.fun/leaf/util/convert"
)

// 关注
func Follow(ctx *gin.Context) {
	var idDTO dto.IdDTO
	if err := ctx.Bind(&idDTO); err != nil {
		resp.Response(ctx, resp.RequestParamError, "", nil)
		zap.L().Error("请求参数有误")
		return
	}

	//判断关注的是否为自己
	userId := ctx.GetUint("userId")
	if idDTO.ID == userId {
		resp.Response(ctx, resp.FollowYourselfError, "", nil)
		zap.L().Error("不能关注自己 ")
		return
	}

	// 是否已经关注
	if !service.IsFollow(userId, idDTO.ID) {
		// 存入数据库
		service.InsertFollow(model.Follow{Uid: userId, Fid: idDTO.ID})
	}

	// 返回给前端
	resp.OK(ctx, "ok", nil)
}

// 取消关注
func UnFollow(ctx *gin.Context) {
	var idDTO dto.IdDTO
	if err := ctx.Bind(&idDTO); err != nil {
		resp.Response(ctx, resp.RequestParamError, "", nil)
		zap.L().Error("请求参数有误")
		return
	}

	// 删除关注记录
	userId := ctx.GetUint("userId")
	service.DeleteFollow(userId, idDTO.ID)

	// 返回给前端
	resp.OK(ctx, "ok", nil)
}

// 获取关注状态
func GetFollowStatus(ctx *gin.Context) {
	fid := convert.StringToUint(ctx.DefaultQuery("fid", "0"))

	userId := ctx.GetUint("userId")
	follow := service.IsFollow(userId, fid)

	// 返回给前端
	resp.OK(ctx, "ok", gin.H{"follow": follow})
}

// 通过UID获取关注列表
func GetFollowings(ctx *gin.Context) {
	uid := convert.StringToUint(ctx.DefaultQuery("uid", "0"))
	page := convert.StringToInt(ctx.DefaultQuery("page", "1"))
	pageSize := convert.StringToInt(ctx.DefaultQuery("page_size", "10"))

	if pageSize > 30 {
		resp.Response(ctx, resp.TooManyRequestsError, "", nil)
		zap.L().Error("请求数量过多 ")
		return
	}

	users := service.SelectFollowingUser(uid, page, pageSize)

	// 返回给前端
	resp.OK(ctx, "ok", gin.H{"users": vo.ToBaseUserVoList(users)})
}

// 通过UID获取粉丝列表
func GetFollowers(ctx *gin.Context) {
	uid := convert.StringToUint(ctx.DefaultQuery("uid", "0"))
	page := convert.StringToInt(ctx.DefaultQuery("page", "1"))
	pageSize := convert.StringToInt(ctx.DefaultQuery("page_size", "10"))

	if pageSize > 30 {
		resp.Response(ctx, resp.TooManyRequestsError, "", nil)
		zap.L().Error("请求数量过多 ")
		return
	}

	users := service.SelectFollowerUser(uid, page, pageSize)

	// 返回给前端
	resp.OK(ctx, "ok", gin.H{"users": vo.ToBaseUserVoList(users)})
}

// 通过UID获取粉丝数
func GetFollowCount(ctx *gin.Context) {
	uid := convert.StringToUint(ctx.DefaultQuery("uid", "0"))

	following, follower := service.GetFollowCount(uid)
	// 返回给前端
	resp.OK(ctx, "ok", gin.H{"following": following, "follower": follower})
}
