package api

import (
	"github.com/gin-gonic/gin"
	"kuukaa.fun/leaf/domain/resp"
	"kuukaa.fun/leaf/domain/vo"
	"kuukaa.fun/leaf/service"
	"kuukaa.fun/leaf/util/convert"
)

// 通过视频ID获取点赞收藏数据
func GetArchiveStat(ctx *gin.Context) {
	videoId := convert.StringToUint(ctx.DefaultQuery("vid", "0"))

	likeCount, _ := service.SelectLikeCount(videoId)
	collectCount, _ := service.SelectCollectCount(videoId)

	// 返回给前端
	resp.OK(ctx, "ok", gin.H{"stat": vo.ArchiveStatVO{
		Like:    int64(likeCount),
		Collect: int64(collectCount),
	}})
}
