package api

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"kuukaa.fun/leaf/domain/dto"
	"kuukaa.fun/leaf/domain/resp"
	"kuukaa.fun/leaf/service"
	"kuukaa.fun/leaf/util/convert"
)

// 添加收藏
func Collect(ctx *gin.Context) {
	var addCollectDTO dto.AddCollectDTO
	if err := ctx.Bind(&addCollectDTO); err != nil {
		resp.Response(ctx, resp.RequestParamError, "", nil)
		zap.L().Error("请求参数有误")
		return
	}

	video := service.GetVideoInfo(addCollectDTO.Vid)
	if video.ID == 0 {
		resp.Response(ctx, resp.VideoNotExistError, "", nil)
		zap.L().Error("视频不存在")
		return
	}

	userId := ctx.GetUint("userId")
	//处理收藏部分
	service.Collect(addCollectDTO.Vid, userId, addCollectDTO.AddList)
	//处理取消收藏部分
	service.CancelCollect(addCollectDTO.Vid, userId, addCollectDTO.CancelList)

	// 返回给前端
	resp.OK(ctx, "ok", nil)
}

// 是否收藏
func HasCollect(ctx *gin.Context) {
	videoId := convert.StringToUint(ctx.DefaultQuery("vid", "0"))

	userId := ctx.GetUint("userId")
	isCollect, err := service.IsCollect(videoId, userId)
	if err != nil {
		resp.Response(ctx, resp.Error, "", nil)
		zap.L().Error("获取收藏状态失败 " + err.Error())
		return
	}

	// 返回给前端
	resp.OK(ctx, "ok", gin.H{"collect": isCollect})
}

// 获取已收藏的收藏夹
func GetCollectedInfo(ctx *gin.Context) {

	videoId := convert.StringToUint(ctx.DefaultQuery("vid", "0"))

	userId := ctx.GetUint("userId")
	collectionIds, err := service.SelectCollectedInfo(videoId, userId)
	if err != nil {
		resp.Response(ctx, resp.Error, "", nil)
		zap.L().Error("获取收藏状态失败 " + err.Error())
		return
	}

	// 返回给前端
	resp.OK(ctx, "ok", gin.H{"collectionIds": collectionIds})
}
