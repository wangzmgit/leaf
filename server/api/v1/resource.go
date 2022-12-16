package api

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"kuukaa.fun/leaf/domain/dto"
	"kuukaa.fun/leaf/domain/resp"
	"kuukaa.fun/leaf/domain/valid"
	"kuukaa.fun/leaf/service"
)

// 修改资源标题
func ModifyResourceTitle(ctx *gin.Context) {
	var modifyResourceTitleDTO dto.ModifyResourceTitleDTO
	if err := ctx.Bind(&modifyResourceTitleDTO); err != nil {
		resp.Response(ctx, resp.RequestParamError, "", nil)
		zap.L().Error("请求参数有误")
		return
	}

	// 参数校验
	if !valid.Title(modifyResourceTitleDTO.Title) {
		resp.Response(ctx, resp.RequestParamError, valid.TITLE_ERROR, nil)
		zap.L().Error(valid.TITLE_ERROR)
		return
	}

	// 是否为资源作者
	userId := ctx.GetUint("userId")
	resource := service.SelectResourceByID(modifyResourceTitleDTO.ID)
	if resource.ID == 0 || resource.Uid != userId {
		resp.Response(ctx, resp.ResourceNotExistError, "", nil)
		zap.L().Error("资源不存在")
		return
	}

	// 更新数据库
	service.ModifyResourceTitleService(modifyResourceTitleDTO)

	// 返回给前端
	resp.OK(ctx, "ok", nil)
}

// 删除资源
func DeleteResource(ctx *gin.Context) {
	//获取参数
	var idDTO dto.IdDTO
	if err := ctx.Bind(&idDTO); err != nil {
		resp.Response(ctx, resp.RequestParamError, "", nil)
		zap.L().Error("请求参数有误")
		return
	}

	// 是否为资源作者
	userId := ctx.GetUint("userId")
	resource := service.SelectResourceByID(idDTO.ID)
	if resource.ID == 0 || resource.Uid != userId {
		resp.Response(ctx, resp.ResourceNotExistError, "", nil)
		zap.L().Error("资源不存在")
		return
	}

	service.DeleteResource(idDTO.ID)

	// 返回给前端
	resp.OK(ctx, "ok", nil)
}
