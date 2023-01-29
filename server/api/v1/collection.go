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
	"kuukaa.fun/leaf/util/convert"
)

// 添加收藏夹
func CreateCollection(ctx *gin.Context) {
	var collectionDTO dto.CollectionDTO
	if err := ctx.Bind(&collectionDTO); err != nil {
		resp.Response(ctx, resp.RequestParamError, "", nil)
		zap.L().Error("请求参数有误")
		return
	}

	// 参数校验
	if !valid.Name(collectionDTO.Name) {
		resp.Response(ctx, resp.RequestParamError, valid.COLLECTION_NAME_ERROR, nil)
		zap.L().Error(valid.COLLECTION_NAME_ERROR)
		return
	}

	userId := ctx.GetUint("userId")
	collection := dto.CollectionDtoToCollection(userId, collectionDTO)

	// 插入数据库
	collectionId, _ := service.InsertCollection(collection)

	// 返回给前端
	resp.OK(ctx, "ok", gin.H{"id": collectionId})
}

// 修改收藏夹
func ModifyCollection(ctx *gin.Context) {
	var modifyCollectionDTO dto.ModifyCollectionDTO
	if err := ctx.Bind(&modifyCollectionDTO); err != nil {
		resp.Response(ctx, resp.RequestParamError, "", nil)
		zap.L().Error("请求参数有误")
		return
	}

	if !valid.Name(modifyCollectionDTO.Name) {
		resp.Response(ctx, resp.RequestParamError, valid.COLLECTION_NAME_ERROR, nil)
		zap.L().Error(valid.COLLECTION_NAME_ERROR)
		return
	}

	userId := ctx.GetUint("userId")
	if modifyCollectionDTO.Cover != "" && cache.GetUploadImage(modifyCollectionDTO.Cover) != userId {
		resp.Response(ctx, resp.InvalidLinkError, "", nil)
		zap.L().Error("文件链接无效")
		return
	}

	if !service.IsCollectionBelongUser(modifyCollectionDTO.ID, userId) {
		resp.Response(ctx, resp.CollectionNotExistError, "", nil)
		zap.L().Error("收藏夹不存在")
		return
	}

	service.ModifyCollection(modifyCollectionDTO)

	// 返回给前端
	resp.OK(ctx, "ok", nil)
}

// 删除收藏夹
func DeleteCollection(ctx *gin.Context) {
	var idDTO dto.IdDTO
	if err := ctx.Bind(&idDTO); err != nil {
		resp.Response(ctx, resp.RequestParamError, "", nil)
		zap.L().Error("请求参数有误")
		return
	}

	userId := ctx.GetUint("userId")
	if !service.IsCollectionBelongUser(idDTO.ID, userId) {
		resp.Response(ctx, resp.CollectionNotExistError, "", nil)
		zap.L().Error("收藏夹不存在")
		return
	}

	service.DeleteCollection(idDTO.ID)

	// 返回给前端
	resp.OK(ctx, "ok", nil)
}

// 获取收藏夹列表
func GetCollectionList(ctx *gin.Context) {
	userId := ctx.GetUint("userId")
	collections := service.SelectCollectionListByUid(userId)

	// 返回给前端
	resp.OK(ctx, "ok", gin.H{"collections": vo.CollectionListToVoList(collections)})
}

// 获取收藏夹信息
func GetCollectionInfo(ctx *gin.Context) {
	id := convert.StringToUint(ctx.DefaultQuery("id", "0"))

	userId := ctx.GetUint("userId")
	collection := service.SelectCollectionByID(id)
	if !collection.Open && collection.Uid != userId {
		resp.Response(ctx, resp.CollectionNotExistError, "", nil)
		zap.L().Error("收藏夹不存在")
		return
	}

	user := service.GetUserInfo(collection.Uid)

	// 返回给前端
	resp.OK(ctx, "ok", gin.H{"collection": vo.CollectionToVo(collection), "author": vo.ToBaseUserVO(user)})
}
