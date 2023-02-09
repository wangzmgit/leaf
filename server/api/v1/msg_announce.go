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

// 获取公告
func GetAnnounce(ctx *gin.Context) {
	page := convert.StringToInt(ctx.Query("page"))
	pageSize := convert.StringToInt(ctx.Query("page_size"))

	if pageSize > 30 {
		resp.Response(ctx, resp.TooManyRequestsError, "", nil)
		zap.L().Error("请求数量过多 ")
		return
	}

	total, announces := service.SelectAnnounce(page, pageSize)

	// 返回给前端
	resp.OK(ctx, "ok", gin.H{"total": total, "announces": vo.ToAnnounceVoList(announces)})
}

// 获取重要公告
func GetImportantAnnounce(ctx *gin.Context) {
	announce := service.SelectImportantAnnounce()

	// 返回给前端
	resp.OK(ctx, "ok", gin.H{"announce": vo.ToAnnounceVO(announce)})
}

// 添加公告
func AddAnnounce(ctx *gin.Context) {
	//获取参数
	var announceDTO dto.AnnounceDTO
	if err := ctx.Bind(&announceDTO); err != nil {
		resp.Response(ctx, resp.RequestParamError, "", nil)
		zap.L().Error("请求参数有误")
		return
	}

	// 参数校验
	if !valid.AnnounceTitle(announceDTO.Title) {
		resp.Response(ctx, resp.RequestParamError, valid.ANNOUNCE_TITLE_ERROR, nil)
		zap.L().Error(valid.ANNOUNCE_TITLE_ERROR)
		return
	}

	if !valid.AnnounceContent(announceDTO.Content) {
		resp.Response(ctx, resp.RequestParamError, valid.ANNOUNCE_CONTENT_ERROR, nil)
		zap.L().Error(valid.ANNOUNCE_CONTENT_ERROR)
		return
	}

	if !valid.AnnounceUrl(announceDTO.Url) {
		resp.Response(ctx, resp.RequestParamError, valid.ANNOUNCE_URL_ERROR, nil)
		zap.L().Error(valid.ANNOUNCE_URL_ERROR)
		return
	}

	// 保存到数据库
	announce := dto.AnnounceDtoToAnnounce(announceDTO)
	service.InsertAnnounce(announce)

	// 返回给前端
	resp.OK(ctx, "ok", nil)
}

// 删除公告
func DeleteAnnounce(ctx *gin.Context) {
	var idDTO dto.IdDTO
	if err := ctx.Bind(&idDTO); err != nil {
		resp.Response(ctx, resp.RequestParamError, "", nil)
		zap.L().Error("请求参数有误")
		return
	}

	service.DeleteAnnounce(idDTO.ID)

	// 返回给前端
	resp.OK(ctx, "ok", nil)
}
