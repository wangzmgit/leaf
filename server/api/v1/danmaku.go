package api

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"kuukaa.fun/leaf/domain/dto"
	"kuukaa.fun/leaf/domain/resp"
	"kuukaa.fun/leaf/domain/valid"
	"kuukaa.fun/leaf/domain/vo"
	"kuukaa.fun/leaf/service"
	"kuukaa.fun/leaf/util/convert"
)

// 获取弹幕
func GetDanmaku(ctx *gin.Context) {
	vid := convert.StringToInt(ctx.Query("vid"))
	part := convert.StringToInt(ctx.Query("part"))

	danmaku := service.SelectDanmakuByVidAndPart(vid, part)

	// 返回给前端
	resp.OK(ctx, "ok", gin.H{"danmaku": vo.ToDanmakuVoList(danmaku)})
}

// 发送弹幕
func SendDanmaku(ctx *gin.Context) {
	// 获取参数
	var danmakuDTO dto.DanmakuDTO
	if err := ctx.Bind(&danmakuDTO); err != nil {
		resp.Response(ctx, resp.RequestParamError, "", nil)
		zap.L().Error("请求参数有误")
		return
	}

	// 参数校验
	if !valid.DanmakuText(danmakuDTO.Text) { //弹幕内容验证
		resp.Response(ctx, resp.RequestParamError, valid.EMAIL_ERROR, nil)
		zap.L().Error(valid.EMAIL_ERROR)
		return
	}

	video := service.GetVideoInfo(danmakuDTO.Vid)
	if video.ID == 0 { // 验证视频是否存在
		resp.Response(ctx, resp.VideoNotExistError, valid.DANMAKU_TEXT_ERROR, nil)
		zap.L().Error(valid.DANMAKU_TEXT_ERROR)
		return
	}

	if danmakuDTO.Part == 0 {
		danmakuDTO.Part = 1
	}

	// 存入数据库
	userId := ctx.GetUint("userId")
	danmaku := dto.DanmakuDtoToDanmaku(danmakuDTO, userId)
	service.InsertDanmaku(danmaku)

	// 返回给前端
	resp.OK(ctx, "ok", nil)
}
