package api

import (
	"path"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"kuukaa.fun/leaf/cache"
	"kuukaa.fun/leaf/domain/resp"
	"kuukaa.fun/leaf/domain/valid"
	"kuukaa.fun/leaf/service"
	"kuukaa.fun/leaf/util/convert"
)

func UploadImg(ctx *gin.Context) {
	img, err := ctx.FormFile("image")
	if err != nil {
		resp.Response(ctx, resp.FileUploadError, "", nil)
		zap.L().Error("文件上传失败")
		return
	}

	suffix := path.Ext(img.Filename)
	fileName := service.GenerateUniqueFilename("img_", suffix)
	// 参数校验
	if !valid.FileType(suffix, true) { // 文件后缀
		resp.Response(ctx, resp.FileCheckError, valid.FILE_TYPE_ERROR, nil)
		zap.L().Error(valid.FILE_TYPE_ERROR)
		return
	}

	//文件大小限制
	if !valid.FileSize(ctx.GetHeader("Content-Length"), viper.GetInt64("file.max_img_size")) {
		resp.Response(ctx, resp.FileCheckError, valid.FILE_SIZE_ERROR, nil)
		zap.L().Error(valid.FILE_SIZE_ERROR)
		return
	}

	//保存文件
	if err := ctx.SaveUploadedFile(img, "./upload/image/"+fileName); err != nil {
		resp.Response(ctx, resp.Error, "文件保存失败", nil)
		zap.L().Error("文件保存失败")
		return
	}

	url, err := service.UploadImgService("img", fileName)
	if err != nil {
		resp.Response(ctx, resp.OssError, "", nil)
		zap.L().Error("上传OSS错误:" + err.Error())
		return
	}

	userId := ctx.GetUint("userId")

	// 缓存url
	cache.SetUploadImage(url, userId)

	// 记录日志
	zap.L().Info("用户上传图片:" + fileName + ",用户ID:" + convert.UintToString(userId))

	// 返回给前端URL
	resp.OK(ctx, "", gin.H{"url": url})
}
