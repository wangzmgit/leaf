package api

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"kuukaa.fun/leaf/domain/resp"
	"kuukaa.fun/leaf/domain/vo"
	"kuukaa.fun/leaf/service"
	"kuukaa.fun/leaf/util/convert"
)

// 获取卡片数据
func GetCardData(ctx *gin.Context) {
	userCount := service.SelectUserCount()
	videoCount := service.SelectVideoCount()
	reviewVideoCount := service.SelectReviewVideoCount()
	newUserCount, _ := service.SelectNewUserCount(0)
	newVideoCount, _ := service.SelectNewVideoCount(0)

	// 返回给前端
	resp.OK(ctx, "ok", gin.H{"data": vo.CardDataVO{
		UserCount:        userCount,
		VideoCount:       videoCount,
		ReviewVideoCount: reviewVideoCount,
		NewUserCount:     newUserCount,
		NewVideoCount:    newVideoCount,
	}})
}

// 获取网站近期(5天)数据
func GetTrendData(ctx *gin.Context) {
	trend := make([]vo.OneDayTrendVO, 5)
	for i := 0; i < 5; i++ {
		trend[i].User, trend[i].Date = service.SelectNewUserCount(i)
		trend[i].Video, _ = service.SelectNewVideoCount(i)
	}

	// 返回给前端
	resp.OK(ctx, "ok", gin.H{"data": trend})
}

// 获取网站分区视频数
func GetPartitionData(ctx *gin.Context) {
	partitionId := convert.StringToUint(ctx.DefaultQuery("partition_id", "0")) //分区
	if partitionId == 0 {
		resp.Response(ctx, resp.RequestParamError, "分区Id不能为0", nil)
		zap.L().Error("分区Id不能为0 ")
		return
	}

	partitionVideoCount := service.SelectPartitionVideoCount(partitionId)

	// 返回给前端
	resp.OK(ctx, "ok", gin.H{"data": partitionVideoCount})
}
