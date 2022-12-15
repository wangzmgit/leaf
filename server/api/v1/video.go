package api

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"kuukaa.fun/leaf/cache"
	"kuukaa.fun/leaf/common"
	"kuukaa.fun/leaf/domain/dto"
	"kuukaa.fun/leaf/domain/resp"
	"kuukaa.fun/leaf/domain/valid"
	"kuukaa.fun/leaf/domain/vo"
	"kuukaa.fun/leaf/service"
	"kuukaa.fun/leaf/util/convert"
)

// 上传视频信息
func UploadVideoInfo(ctx *gin.Context) {
	var uploadVideoDTO dto.UploadVideoDTO
	if err := ctx.Bind(&uploadVideoDTO); err != nil {
		resp.Response(ctx, resp.RequestParamError, "", nil)
		zap.L().Error("请求参数有误")
		return
	}

	// 参数校验
	if !valid.Title(uploadVideoDTO.Title) {
		resp.Response(ctx, resp.RequestParamError, valid.TITLE_ERROR, nil)
		zap.L().Error(valid.TITLE_ERROR)
		return
	}

	userId := ctx.GetUint("userId")
	if cache.GetUploadImage(uploadVideoDTO.Cover) != userId {
		resp.Response(ctx, resp.InvalidLinkError, "", nil)
		zap.L().Error("文件链接无效")
		return
	}

	if !service.IsSubpartition(uploadVideoDTO.Partition) {
		resp.Response(ctx, resp.PartitionError, "", nil)
		zap.L().Error("分区不存在")
		return
	}

	video := dto.UploadVideoDtoToVideo(userId, uploadVideoDTO)
	vid := service.InsertVideo(video)

	// 返回给前端
	resp.OK(ctx, "ok", gin.H{"vid": vid})
}

// 修改视频信息
func ModifyVideoInfo(ctx *gin.Context) {
	var modifyVideoDTO dto.ModifyVideoDTO
	if err := ctx.Bind(&modifyVideoDTO); err != nil {
		resp.Response(ctx, resp.RequestParamError, "", nil)
		zap.L().Error("请求参数有误")
		return
	}

	// 参数校验
	if !valid.Title(modifyVideoDTO.Title) {
		resp.Response(ctx, resp.RequestParamError, valid.TITLE_ERROR, nil)
		zap.L().Error(valid.TITLE_ERROR)
		return
	}

	// 校验用户是否为视频作者
	userId := ctx.GetUint("userId")
	oldVideoInfo := service.SelectVideoByID(modifyVideoDTO.VID)
	if oldVideoInfo.Uid != userId {
		if modifyVideoDTO.Cover != oldVideoInfo.Cover && cache.GetUploadImage(modifyVideoDTO.Cover) != userId {
			resp.Response(ctx, resp.VideoNotExistError, "", nil)
			zap.L().Error("视频不存在")
			return
		}
	}

	// 校验封面图文件是否有效
	if modifyVideoDTO.Cover != oldVideoInfo.Cover && cache.GetUploadImage(modifyVideoDTO.Cover) != userId {
		resp.Response(ctx, resp.InvalidLinkError, "", nil)
		zap.L().Error("文件链接无效")
		return
	}

	// 保存到数据库
	service.UpdateVideoInfo(modifyVideoDTO)
	// 返回给前端
	resp.OK(ctx, "ok", nil)
}

// 获取视频状态
func GetVideoStatus(ctx *gin.Context) {
	videoId := convert.StringToUint(ctx.DefaultQuery("vid", "0"))
	video := service.SelectVideoByID(videoId)

	// 返回给前端
	resp.OK(ctx, "ok", gin.H{"video": vo.ToVideoStatusVO(video)})
}

// 获取视频信息
func GetVideoByID(ctx *gin.Context) {
	vid := convert.StringToUint(ctx.DefaultQuery("vid", "0"))

	video := service.SelectVideoByID(vid)
	if video.ID == 0 || video.Status != common.AUDIT_APPROVED {
		resp.Response(ctx, resp.VideoNotExistError, "", nil)
		zap.L().Error("视频不存在")
		return
	}

	//获取作者信息
	author := service.GetUserInfo(video.Uid)

	//获取视频资源
	resources := service.SelectResourceByVideo(video.ID)

	//获取视频统计数据
	// TODO: 获取点赞和收藏数据

	//增加播放量(一个ip在同一个视频下，每30分钟可重新增加1播放量)
	if cache.GetClicks(video.ID, ctx.ClientIP()) == "" {
		service.AddClicks(video.ID)
		cache.SetClicks(video.ID, ctx.ClientIP())
	}

	// 返回给前端
	resp.OK(ctx, "ok", gin.H{"video": vo.ToVideoVO(video, author, resources)})
}
