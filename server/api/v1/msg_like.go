package api

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"kuukaa.fun/leaf/domain/resp"
	"kuukaa.fun/leaf/domain/vo"
	"kuukaa.fun/leaf/service"
	"kuukaa.fun/leaf/util/convert"
)

func GetLikeMessage(ctx *gin.Context) {
	page := convert.StringToInt(ctx.Query("page"))
	pageSize := convert.StringToInt(ctx.Query("page_size"))

	if pageSize > 30 {
		resp.Response(ctx, resp.TooManyRequestsError, "", nil)
		zap.L().Error("请求数量过多 ")
		return
	}

	userId := ctx.GetUint("userId")

	messages := service.SelectLikeMessage(userId, page, pageSize)
	// 查询对应的用户和视频
	for i := 0; i < len(messages); i++ {
		messages[i].User = service.GetUserInfo(messages[i].Fid)
		messages[i].Video = service.GetVideoInfo(messages[i].Vid)
	}

	// 返回给前端
	resp.OK(ctx, "ok", gin.H{"messages": vo.ToLikeMessageVoList(messages)})
}
