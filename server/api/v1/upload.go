package api

import (
	"os"
	"path"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"kuukaa.fun/leaf/cache"
	"kuukaa.fun/leaf/domain/dto"
	"kuukaa.fun/leaf/domain/resp"
	"kuukaa.fun/leaf/domain/valid"
	"kuukaa.fun/leaf/service"
	"kuukaa.fun/leaf/util/convert"
)

func UploadImg(ctx *gin.Context) {
	img, err := ctx.FormFile("image")
	if err != nil {
		resp.Response(ctx, resp.FileUploadError, "", nil)
		zap.L().Error("文件上传失败" + err.Error())
		return
	}

	suffix := path.Ext(img.Filename)
	fileName := service.GenerateImgFilename(suffix)
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
		zap.L().Error("文件保存失败" + err.Error())
		return
	}

	url, err := service.GenerateFileUrl("image/" + fileName)
	if err != nil {
		resp.Response(ctx, resp.Error, "文件保存失败", nil)
		zap.L().Error("生成url失败" + err.Error())
		return
	}

	if viper.GetString("oss.type") != "local" {
		if err := service.UploadImgToOss(fileName); err != nil {
			resp.Response(ctx, resp.OssError, "", nil)
			zap.L().Error("上传OSS错误:" + err.Error())
			return
		}
	}

	// 缓存url
	userId := ctx.GetUint("userId")
	cache.SetUploadImage(url, userId)

	// 记录日志
	zap.L().Info("用户上传图片:" + fileName + ",用户ID:" + convert.UintToString(userId))

	// 返回给前端URL
	resp.OK(ctx, "ok", gin.H{"url": url})
}

func UploadVideo(ctx *gin.Context) {
	// 获取视频ID
	vid := convert.StringToUint(ctx.Param("vid"))
	if vid == 0 {
		resp.Response(ctx, resp.RequestParamError, "", nil)
		zap.L().Error("请求参数有误")
		return
	}

	img, err := ctx.FormFile("video")
	if err != nil {
		resp.Response(ctx, resp.FileUploadError, "", nil)
		zap.L().Error("文件上传失败" + err.Error())
		return
	}

	suffix := path.Ext(img.Filename)
	filenameOnly := service.GenerateVideoFilename()
	// 参数校验
	userId := ctx.GetUint("userId")
	videoInfo := service.GetVideoInfo(vid)
	if videoInfo.ID == 0 || videoInfo.Uid != userId {
		resp.Response(ctx, resp.VideoNotExistError, "", nil)
		zap.L().Error("视频不存在")
		return
	}

	if !valid.FileType(suffix, false) { // 文件后缀
		resp.Response(ctx, resp.FileCheckError, valid.FILE_TYPE_ERROR, nil)
		zap.L().Error(valid.FILE_TYPE_ERROR)
		return
	}

	//文件大小限制
	if !valid.FileSize(ctx.GetHeader("Content-Length"), viper.GetInt64("file.max_video_size")) {
		resp.Response(ctx, resp.FileCheckError, valid.FILE_SIZE_ERROR, nil)
		zap.L().Error(valid.FILE_SIZE_ERROR)
		return
	}

	if err := os.Mkdir("./upload/video/"+filenameOnly, os.ModePerm); err != nil {
		resp.Response(ctx, resp.Error, "文件保存失败", nil)
		zap.L().Error("创建视频文件夹失败" + err.Error())
		return
	}

	//保存文件
	uploadVideoPath := "./upload/video/" + filenameOnly + "/upload.mp4"
	if err := ctx.SaveUploadedFile(img, uploadVideoPath); err != nil {
		resp.Response(ctx, resp.Error, "文件保存失败", nil)
		zap.L().Error("文件保存失败" + err.Error())
		return
	}

	quality, duration, err := service.PreTreatmentVideo(uploadVideoPath)
	if err != nil {
		resp.Response(ctx, resp.Error, "处理视频失败", nil)
		zap.L().Error("预处理视频失败" + err.Error())
		return
	}

	// 生成url
	url, err := service.GenerateFileUrl("video/" + filenameOnly + "/index.mpd")
	if err != nil {
		resp.Response(ctx, resp.Error, "文件保存失败", nil)
		zap.L().Error("生成url失败" + err.Error())
		return
	}

	// 生成原始文件的url
	originalUrl := "video/" + filenameOnly + "/upload.mp4"
	if viper.GetBool("file.video_adaptation_ios") {
		originalUrl, _ = service.GenerateFileUrl(originalUrl)
	}

	// 存入数据库
	resource := dto.ResourceDtoToResource(vid, userId, quality, duration, url, originalUrl)
	rid := service.InsertResource(resource)

	// 启动转码服务
	go service.VideoTransCoding(rid, quality, filenameOnly)

	// 记录日志
	zap.L().Info("用户上传视频:" + filenameOnly + ",用户ID:" + convert.UintToString(userId))

	// 返回给前端
	resp.OK(ctx, "ok", nil)
}
