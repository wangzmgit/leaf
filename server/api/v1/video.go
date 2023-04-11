package api

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"kuukaa.fun/leaf/cache"
	"kuukaa.fun/leaf/common"
	"kuukaa.fun/leaf/domain/dto"
	"kuukaa.fun/leaf/domain/model"
	"kuukaa.fun/leaf/domain/resp"
	"kuukaa.fun/leaf/domain/valid"
	"kuukaa.fun/leaf/domain/vo"
	"kuukaa.fun/leaf/service"
	"kuukaa.fun/leaf/util/convert"
	"kuukaa.fun/leaf/ws"
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
	vid, err := service.InsertVideo(video)
	if err != nil {
		resp.Response(ctx, resp.Error, "创建视频失败", nil)
		zap.L().Error("创建视频失败 " + err.Error())
		return
	}

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
	oldVideoInfo := service.GetVideoInfo(modifyVideoDTO.VID)
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
	video := service.GetVideoInfo(videoId)

	resources := service.SelectResourceByVideo(videoId, false)

	// 返回给前端
	resp.OK(ctx, "ok", gin.H{"video": vo.ToVideoStatusVO(video, resources)})
}

// 获取视频信息
func GetVideoByID(ctx *gin.Context) {
	ios := convert.StringToUint(ctx.DefaultQuery("ios", "0"))
	vid := convert.StringToUint(ctx.DefaultQuery("vid", "0"))

	video := service.GetVideoInfo(vid)
	if video.ID == 0 || video.Status != common.AUDIT_APPROVED {
		resp.Response(ctx, resp.VideoNotExistError, "", nil)
		zap.L().Error("视频不存在")
		return
	}

	//获取作者信息
	video.Author = service.GetUserInfo(video.Uid)

	//获取视频资源
	resources := service.SelectResourceByVideo(video.ID, true)

	//增加播放量(一个ip在同一个视频下，每30分钟可重新增加1播放量)
	service.AddVideoClicks(video.ID, ctx.ClientIP())

	// 获取播放量
	video.Clicks = service.GetVideoClicks(video.ID)

	// 返回给前端
	resp.OK(ctx, "ok", gin.H{"video": vo.ToVideoVO(video, resources, ios)})
}

// 提交审核
func SubmitReview(ctx *gin.Context) {
	//获取参数
	var idDTO dto.IdDTO
	if err := ctx.Bind(&idDTO); err != nil {
		resp.Response(ctx, resp.RequestParamError, "", nil)
		zap.L().Error("请求参数有误")
		return
	}

	if service.SelectResourceCountByVid(idDTO.ID) == 0 {
		resp.Response(ctx, resp.ResourceNotExistError, "", nil)
		zap.L().Error("资源不存在")
		return
	}

	// 更新视频状态
	service.UpadteVideoStatus(idDTO.ID, common.WAITING_REVIEW)

	// 返回给前端
	resp.OK(ctx, "ok", nil)
}

// 获取收藏视频列表
func GetCollectVideo(ctx *gin.Context) {
	id := convert.StringToUint(ctx.DefaultQuery("id", "0"))

	userId := ctx.GetUint("userId")
	page := convert.StringToInt(ctx.DefaultQuery("page", "1"))
	pageSize := convert.StringToInt(ctx.DefaultQuery("page_size", "10"))

	if pageSize > 30 {
		resp.Response(ctx, resp.TooManyRequestsError, "", nil)
		zap.L().Error("请求数量过多 ")
		return
	}

	collection := service.SelectCollectionByID(id)
	if !collection.Open && collection.Uid != userId {
		resp.Response(ctx, resp.CollectionNotExistError, "", nil)
		zap.L().Error("收藏夹不存在")
		return
	}

	videos, total, err := service.SelectCollectVideo(id, page, pageSize)
	if err != nil {
		resp.Response(ctx, resp.Error, "", nil)
		zap.L().Error("获取收藏视频失败" + err.Error())
		return
	}

	// 返回给前端
	resp.OK(ctx, "ok", gin.H{"total": total, "videos": vo.ToBaseVideoVoList(videos)})
}

// 获取自己的视频
func GetUploadVideoList(ctx *gin.Context) {
	userId := ctx.GetUint("userId")
	page := convert.StringToInt(ctx.Query("page"))
	pageSize := convert.StringToInt(ctx.Query("page_size"))

	if pageSize > 30 {
		resp.Response(ctx, resp.TooManyRequestsError, "", nil)
		zap.L().Error("请求数量过多 ")
		return
	}

	total, videos := service.SelectUploadVideo(userId, page, pageSize)
	// 更新播放量数据
	for i := 0; i < len(videos); i++ {
		videos[i].Clicks = service.GetVideoClicks(videos[i].ID)
	}

	// 返回给前端
	resp.OK(ctx, "ok", gin.H{"total": total, "videos": vo.ToUserUploadVideoVoList(videos)})
}

// 通过用户ID获取视频列表
func GetVideoListByUid(ctx *gin.Context) {
	uid := convert.StringToUint(ctx.Query("uid"))
	page := convert.StringToInt(ctx.Query("page"))
	pageSize := convert.StringToInt(ctx.Query("page_size"))

	if pageSize > 30 {
		resp.Response(ctx, resp.TooManyRequestsError, "", nil)
		zap.L().Error("请求数量过多 ")
		return
	}

	total, videos := service.SelectVideoByUserId(uid, page, pageSize)
	// 更新播放量数据
	for i := 0; i < len(videos); i++ {
		videos[i].Clicks = service.GetVideoClicks(videos[i].ID)
	}

	// 返回给前端
	resp.OK(ctx, "ok", gin.H{"total": total, "videos": vo.ToBaseVideoVoList(videos)})
}

// 删除视频
func DeleteVideo(ctx *gin.Context) {
	var idDTO dto.IdDTO
	if err := ctx.Bind(&idDTO); err != nil {
		resp.Response(ctx, resp.RequestParamError, "", nil)
		zap.L().Error("请求参数有误")
		return
	}

	userId := ctx.GetUint("userId")
	if !service.IsVideoBelongUser(idDTO.ID, userId) {
		resp.Response(ctx, resp.VideoNotExistError, "", nil)
		zap.L().Error("视频不存在")
		return
	}

	service.DeleteVideo(idDTO.ID)

	// 返回给前端
	resp.OK(ctx, "ok", nil)
}

// 获取视频列表
func GetVideoList(ctx *gin.Context) {
	page := convert.StringToInt(ctx.Query("page"))
	pageSize := convert.StringToInt(ctx.Query("page_size"))
	partitionId := convert.StringToUint(ctx.DefaultQuery("partition", "0")) //分区

	if pageSize > 30 {
		resp.Response(ctx, resp.TooManyRequestsError, "", nil)
		zap.L().Error("请求数量过多 ")
		return
	}

	var videos []model.Video
	if partitionId == 0 {
		//不传分区参数默认查询全部
		_, videos = service.SelectVideoListByStatus(page, pageSize, common.AUDIT_APPROVED)
	} else if service.IsSubpartition(partitionId) {
		// 如果为子分区，查询分区下的视频
		_, videos = service.SelectVideoListBySubpartition(partitionId, page, pageSize)
	} else {
		// 获取该分区下的视频
		_, videos = service.SelectVideoListByPartition(partitionId, page, pageSize)
	}

	// 更新播放量数据和作者信息
	for i := 0; i < len(videos); i++ {
		videos[i].Clicks = service.GetVideoClicks(videos[i].ID)
		videos[i].Author = service.GetUserInfo(videos[i].Uid)
	}

	// 返回给前端
	resp.OK(ctx, "ok", gin.H{"videos": vo.ToSearchVideoVoList(videos)})
}

// 获取推荐视频
func GetRecommendedVideo(ctx *gin.Context) {
	pageSize := convert.StringToInt(ctx.DefaultQuery("page_size", "15"))

	if pageSize > 30 {
		resp.Response(ctx, resp.TooManyRequestsError, "", nil)
		zap.L().Error("请求数量过多 ")
		return
	}

	// 没有推荐功能，按点击量查询视频（点击量不是实时数据）
	videos := service.SelectVideoListByClicks(pageSize)

	// 更新播放量数据和作者信息
	for i := 0; i < len(videos); i++ {
		videos[i].Clicks = service.GetVideoClicks(videos[i].ID)
		videos[i].Author = service.GetUserInfo(videos[i].Uid)
	}

	// 返回给前端
	resp.OK(ctx, "ok", gin.H{"videos": vo.ToSearchVideoVoList(videos)})
}

// 搜索视频
func SearchVideo(ctx *gin.Context) {
	keywords := ctx.Query("keywords")
	page := convert.StringToInt(ctx.Query("page"))
	pageSize := convert.StringToInt(ctx.DefaultQuery("page_size", "15"))

	if pageSize > 30 {
		resp.Response(ctx, resp.TooManyRequestsError, "", nil)
		zap.L().Error("请求数量过多 ")
		return
	}

	var videos []model.Video
	if len(keywords) == 0 {
		_, videos = service.SelectVideoListByStatus(page, pageSize, common.AUDIT_APPROVED)
	} else {
		// 直接用mysql模糊查询，之后可能会更换为es
		videos = service.SelectVideoListByKeywords(keywords, page, pageSize)
	}

	// 更新播放量数据和作者信息
	for i := 0; i < len(videos); i++ {
		videos[i].Clicks = service.GetVideoClicks(videos[i].ID)
		videos[i].Author = service.GetUserInfo(videos[i].Uid)
	}

	// 返回给前端
	resp.OK(ctx, "ok", gin.H{"videos": vo.ToSearchVideoVoList(videos)})
}

// 获取待审核视频列表
func GetReviewVideoList(ctx *gin.Context) {
	//获取参数
	page := convert.StringToInt(ctx.Query("page"))
	pageSize := convert.StringToInt(ctx.Query("page_size"))

	if pageSize > 30 {
		resp.Response(ctx, resp.TooManyRequestsError, "", nil)
		zap.L().Error("请求数量过多 ")
		return
	}

	total, videos := service.SelectVideoListByStatus(page, pageSize, common.WAITING_REVIEW)

	// 返回给前端
	resp.OK(ctx, "ok", gin.H{"total": total, "videos": vo.ToSearchVideoVoList(videos)})
}

// 审核视频
func ReviewVideo(ctx *gin.Context) {
	var reviewDTO dto.ReviewDTO
	if err := ctx.Bind(&reviewDTO); err != nil {
		resp.Response(ctx, resp.RequestParamError, "", nil)
		zap.L().Error("请求参数有误")
		return
	}

	if !valid.ReviewStatus(reviewDTO.Status) {
		resp.Response(ctx, resp.RequestParamError, valid.REVIEW_STATUS_ERROR, nil)
		zap.L().Error(valid.REVIEW_STATUS_ERROR)
		return
	}

	if reviewDTO.Status == common.AUDIT_APPROVED {
		if service.SelectResourceCountByStatus(reviewDTO.ID, common.AUDIT_APPROVED) == 0 {
			resp.Response(ctx, resp.ResourceNotExistError, "", nil)
			zap.L().Error("资源不存在")
			return
		}
	}

	service.UpadteVideoStatus(reviewDTO.ID, reviewDTO.Status)

	// 返回给前端
	resp.OK(ctx, "ok", nil)
}

// 审核视频资源
func ReviewResource(ctx *gin.Context) {
	var reviewDTO dto.ReviewDTO
	if err := ctx.Bind(&reviewDTO); err != nil {
		resp.Response(ctx, resp.RequestParamError, "", nil)
		zap.L().Error("请求参数有误")
		return
	}

	if !valid.ReviewStatus(reviewDTO.Status) {
		resp.Response(ctx, resp.RequestParamError, valid.REVIEW_STATUS_ERROR, nil)
		zap.L().Error(valid.REVIEW_STATUS_ERROR)
		return
	}

	service.UpadteResourceStatus(reviewDTO.ID, reviewDTO.Status)

	// 返回给前端
	resp.OK(ctx, "ok", nil)
}

// 通过视频ID获取待审核视频资源
func GetReviewVideoByID(ctx *gin.Context) {
	vid := convert.StringToUint(ctx.Query("vid"))

	resources := service.SelectResourceByVideo(vid, false)

	// 返回给前端
	resp.OK(ctx, "ok", gin.H{"resources": vo.ToResourceVoList(resources)})
}

// 管理员获取视频列表
func AdminGetVideoList(ctx *gin.Context) {
	page := convert.StringToInt(ctx.Query("page"))
	pageSize := convert.StringToInt(ctx.Query("page_size"))
	partitionId := convert.StringToUint(ctx.DefaultQuery("partition", "0")) //分区

	if pageSize > 30 {
		resp.Response(ctx, resp.TooManyRequestsError, "", nil)
		zap.L().Error("请求数量过多 ")
		return
	}

	var total int64
	var videos []model.Video
	if partitionId == 0 {
		//不传分区参数默认查询全部
		total, videos = service.SelectVideoListByStatus(page, pageSize, common.AUDIT_APPROVED)
	} else if service.IsSubpartition(partitionId) {
		// 如果为子分区，查询分区下的视频
		total, videos = service.SelectVideoListBySubpartition(partitionId, page, pageSize)
	} else {
		// 获取该分区下的视频
		total, videos = service.SelectVideoListByPartition(partitionId, page, pageSize)
	}

	// 更新播放量数据和作者信息
	for i := 0; i < len(videos); i++ {
		videos[i].Clicks = service.GetVideoClicks(videos[i].ID)
		videos[i].Author = service.GetUserInfo(videos[i].Uid)
	}

	// 返回给前端
	resp.OK(ctx, "ok", gin.H{"total": total, "videos": vo.ToSearchVideoVoList(videos)})
}

// 管理员搜索视频
func AdminSearchVideo(ctx *gin.Context) {
	keywords := ctx.Query("keywords")
	page := convert.StringToInt(ctx.Query("page"))
	pageSize := convert.StringToInt(ctx.Query("page_size"))

	if pageSize > 30 {
		resp.Response(ctx, resp.TooManyRequestsError, "", nil)
		zap.L().Error("请求数量过多 ")
		return
	}

	total, videos := service.AdminSelectVideoListByKeywords(keywords, page, pageSize)

	// 更新播放量数据和作者信息
	for i := 0; i < len(videos); i++ {
		videos[i].Clicks = service.GetVideoClicks(videos[i].ID)
		videos[i].Author = service.GetUserInfo(videos[i].Uid)
	}

	// 返回给前端
	resp.OK(ctx, "ok", gin.H{"total": total, "videos": vo.ToSearchVideoVoList(videos)})
}

// 删除视频
func AdminDeleteVideo(ctx *gin.Context) {
	var idDTO dto.IdDTO
	if err := ctx.Bind(&idDTO); err != nil {
		resp.Response(ctx, resp.RequestParamError, "", nil)
		zap.L().Error("请求参数有误")
		return
	}

	service.DeleteVideo(idDTO.ID)

	// 返回给前端
	resp.OK(ctx, "ok", nil)
}

// 视频Websocket连接(统计在线人数)
func GetRoomConnect(ctx *gin.Context) {
	vid := convert.StringToUint(ctx.Query("vid"))
	clientId := ctx.Query("client_id")
	if vid == 0 {
		return
	}

	// 升级为websocket长链接
	ws.RoomWsHandler(ctx.Writer, ctx.Request, vid, clientId)
}
