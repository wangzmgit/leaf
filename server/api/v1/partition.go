package api

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"kuukaa.fun/leaf/cache"
	"kuukaa.fun/leaf/domain/dto"
	"kuukaa.fun/leaf/domain/resp"
	"kuukaa.fun/leaf/domain/valid"
	"kuukaa.fun/leaf/domain/vo"
	"kuukaa.fun/leaf/service"
)

// 获取分区
func GetPartitionList(ctx *gin.Context) {
	var partitions []vo.PartitionVo

	partitions = cache.GetPartition()
	if len(partitions) == 0 {
		partitions = vo.ToPartitionVo(service.SelectPartition())

		// 存入缓存
		cache.SetPartition(partitions)
	}

	// 返回给前端
	resp.OK(ctx, "ok", gin.H{"partitions": partitions})
}

// 添加分区
func AddPartition(ctx *gin.Context) {
	//获取参数
	var partitionDTO dto.PartitionDTO
	if err := ctx.Bind(&partitionDTO); err != nil {
		resp.Response(ctx, resp.RequestParamError, "", nil)
		zap.L().Error("请求参数有误")
		return
	}

	// 参数校验
	if !valid.Content(partitionDTO.Content) {
		resp.Response(ctx, resp.RequestParamError, valid.CONTENT_ERROR, nil)
		zap.L().Error(valid.CONTENT_ERROR)
		return
	}

	if partitionDTO.ParentId != 0 && !service.IsParentPartitionExist(partitionDTO.ParentId) {
		resp.Response(ctx, resp.ParentPartitionError, "", nil)
		zap.L().Error("所属分区不存在")
		return
	}

	partition := dto.PartitionDtoToPartition(partitionDTO)
	service.InsertPartition(partition)

	// 存入缓存
	cache.SetPartition(vo.ToPartitionVo(service.SelectPartition()))

	// 返回给前端
	resp.OK(ctx, "ok", nil)
}

// 删除分区
func DeletePartition(ctx *gin.Context) {
	var idDTO dto.IdDTO
	if err := ctx.Bind(&idDTO); err != nil {
		resp.Response(ctx, resp.RequestParamError, "", nil)
		zap.L().Error("请求参数有误")
		return
	}

	// 从数据库中删除
	service.DelPartitionByID(idDTO.ID)

	// 删除缓存
	cache.DelPartition()

	// 返回给前端
	resp.OK(ctx, "ok", nil)
}
