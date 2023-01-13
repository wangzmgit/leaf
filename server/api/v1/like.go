package api

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"kuukaa.fun/leaf/domain/dto"
	"kuukaa.fun/leaf/domain/resp"
	"kuukaa.fun/leaf/service"
	"kuukaa.fun/leaf/util/convert"
)

func Like(ctx *gin.Context) {
	var idDTO dto.IdDTO
	if err := ctx.Bind(&idDTO); err != nil {
		resp.Response(ctx, resp.RequestParamError, "", nil)
		zap.L().Error("请求参数有误")
		return
	}

	userId := ctx.GetUint("userId")
	isLike, err := service.IsLike(idDTO.ID, userId)
	if err != nil {
		resp.Response(ctx, resp.Error, "", nil)
		zap.L().Error("获取点赞状态失败 " + err.Error())
		return
	}

	if isLike {
		resp.Response(ctx, resp.Error, "已经点赞", nil)
		zap.L().Error("已经点赞 ")
		return
	}

	// 查询视频作者
	authorId := service.SelectVideoAuthorId(idDTO.ID)

	// 记录点赞内容
	service.Like(idDTO.ID, userId)

	// 添加点赞通知
	msg := dto.ToLikeMessage(idDTO.ID, authorId, userId)
	service.InsertLikeMessage(msg)

	// 返回给前端
	resp.OK(ctx, "ok", nil)
}

func CancelLike(ctx *gin.Context) {
	var idDTO dto.IdDTO
	if err := ctx.Bind(&idDTO); err != nil {
		resp.Response(ctx, resp.RequestParamError, "", nil)
		zap.L().Error("请求参数有误")
		return
	}

	userId := ctx.GetUint("userId")
	isLike, err := service.IsLike(idDTO.ID, userId)
	if err != nil {
		resp.Response(ctx, resp.Error, "", nil)
		zap.L().Error("获取点赞状态失败 " + err.Error())
		return
	}

	if !isLike {
		resp.Response(ctx, resp.Error, "未点赞", nil)
		zap.L().Error("未点赞 ")
		return
	}

	service.CancelLike(idDTO.ID, userId)

	// 删除点赞通知
	service.DeleteLikeMessage(idDTO.ID, userId)

	// 返回给前端
	resp.OK(ctx, "ok", nil)
}

func HasLike(ctx *gin.Context) {
	videoId := convert.StringToUint(ctx.DefaultQuery("vid", "0"))

	userId := ctx.GetUint("userId")
	isLike, err := service.IsLike(videoId, userId)
	if err != nil {
		resp.Response(ctx, resp.Error, "", nil)
		zap.L().Error("获取点赞状态失败 " + err.Error())
		return
	}

	// 返回给前端
	resp.OK(ctx, "ok", gin.H{"like": isLike})
}
