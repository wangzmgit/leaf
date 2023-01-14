package api

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"kuukaa.fun/leaf/cache"
	"kuukaa.fun/leaf/domain/dto"
	"kuukaa.fun/leaf/domain/resp"
	"kuukaa.fun/leaf/domain/vo"
	"kuukaa.fun/leaf/service"
)

// 获取轮播图
func GetCarousel(ctx *gin.Context) {
	carousels := service.SelectCarousel()

	// 返回给前端
	resp.OK(ctx, "ok", gin.H{"carousels": vo.ToCarouselVoList(carousels)})
}

// 上传轮播图信息
func AddCarousel(ctx *gin.Context) {
	// 获取参数
	var carouselDTO dto.CarouselDTO
	if err := ctx.Bind(&carouselDTO); err != nil {
		resp.Response(ctx, resp.RequestParamError, "", nil)
		zap.L().Error("请求参数有误")
		return
	}

	// 参数校验
	userId := ctx.GetUint("userId")
	if carouselDTO.Img == "" || cache.GetUploadImage(carouselDTO.Img) != userId {
		resp.Response(ctx, resp.InvalidLinkError, "", nil)
		zap.L().Error("文件链接无效")
		return
	}

	// 存入数据库
	carousel := dto.CarouselDtoToCarousel(carouselDTO)
	service.InsertCarousel(carousel)

	// 返回给前端
	resp.OK(ctx, "ok", nil)
}

// 删除轮播图
func DeleteCarousel(ctx *gin.Context) {
	var idDTO dto.IdDTO
	if err := ctx.Bind(&idDTO); err != nil {
		resp.Response(ctx, resp.RequestParamError, "", nil)
		zap.L().Error("请求参数有误")
		return
	}

	service.DeleteCarousel(idDTO.ID)

	// 返回给前端
	resp.OK(ctx, "ok", nil)
}
