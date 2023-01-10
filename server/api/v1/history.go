package api

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"kuukaa.fun/leaf/domain/dto"
	"kuukaa.fun/leaf/domain/resp"
	"kuukaa.fun/leaf/domain/vo"
	"kuukaa.fun/leaf/service"
	"kuukaa.fun/leaf/util/convert"
)

// 历史记录
func AddHistory(ctx *gin.Context) {
	var historyDTO dto.HistoryDTO
	if err := ctx.Bind(&historyDTO); err != nil {
		resp.Response(ctx, resp.RequestParamError, "", nil)
		zap.L().Error("请求参数有误")
		return
	}

	// 数据校验
	video := service.GetVideoInfo(historyDTO.Vid)
	if video.ID == 0 {
		resp.Response(ctx, resp.VideoNotExistError, "", nil)
		zap.L().Error("视频不存在")
		return
	}

	if historyDTO.Part == 0 {
		historyDTO.Part = 1
	}

	// 保存到数据库
	userId := ctx.GetUint("userId")
	history := dto.HistoryDtoToHistory(historyDTO, userId)
	service.InsertOrUpdateHistory(history)

	// 返回给前端
	resp.OK(ctx, "ok", nil)
}

// 获取播放进度
func GetHistoryProgress(ctx *gin.Context) {
	videoId := convert.StringToUint(ctx.Query("vid"))

	userId := ctx.GetUint("userId")

	history := service.SelectHistory(videoId, userId)

	// 返回给前端
	resp.OK(ctx, "ok", gin.H{"progress": vo.ToHistoryProgressVO(history)})
}

// 获取播放进度
func GetHistoryVideo(ctx *gin.Context) {
	page := convert.StringToInt(ctx.Query("page"))
	pageSize := convert.StringToInt(ctx.DefaultQuery("page_size", "15"))

	if pageSize > 30 {
		resp.Response(ctx, resp.TooManyRequestsError, "", nil)
		zap.L().Error("请求数量过多 ")
		return
	}

	userId := ctx.GetUint("userId")
	history := service.SelectHistoryVideo(userId, page, pageSize)

	// 更新播放量数据和作者信息
	for i := 0; i < len(history); i++ {
		history[i].Video = service.GetVideoInfo(history[i].Vid)
	}

	// 返回给前端
	resp.OK(ctx, "ok", gin.H{"history": vo.ToHistoryVideoVoList(history)})
}
