package service

import (
	"kuukaa.fun/leaf/cache"
	"kuukaa.fun/leaf/common"
	"kuukaa.fun/leaf/domain/dto"
	"kuukaa.fun/leaf/domain/model"
)

func InsertVideo(video model.Video) (uint, error) {

	if err := mysqlClient.Create(&video).Error; err != nil {
		return 0, err
	}

	// 创建点赞数据
	if err := InsertLike(video.ID); err != nil {
		// 删除MySQL中的收藏夹
		DeleteVideo(video.ID)
		return 0, err
	}

	return video.ID, nil
}

func SelectVideoByID(videoId uint) (video model.Video) {
	mysqlClient.First(&video, videoId)
	// 存入缓存
	cache.SetVideo(video)
	return
}

// 获取视频信息 (通过ID，先查缓存)
func GetVideoInfo(videoId uint) (video model.Video) {
	video = cache.GetVideo(videoId)
	if video.ID == 0 {
		video = SelectVideoByID(videoId)
	}

	return
}

// 通过分区查询视频列表(通过审核)
func SelectVideoListByPartition(partitionId uint, page, pageSize int) (total int64, videos []model.Video) {
	partitionIds := mysqlClient.Model(&model.Partition{}).Select("id").Where("parent_id = ?", partitionId)
	client := mysqlClient.Where("status = ? and partition_id in (?)", common.AUDIT_APPROVED, partitionIds)
	client.Model(&model.Video{}).Count(&total)
	client.Limit(pageSize).Offset((page - 1) * pageSize).Find(&videos)
	return
}

// 通过子分区查询视频列表(通过审核)
func SelectVideoListBySubpartition(partitionId uint, page, pageSize int) (total int64, videos []model.Video) {
	client := mysqlClient.Where("status = ? and partition_id = ?", common.AUDIT_APPROVED, partitionId)
	client.Model(&model.Video{}).Count(&total)
	client.Limit(pageSize).Offset((page - 1) * pageSize).Find(&videos)
	return
}

// 通过关键词查询视频
func SelectVideoListByKeywords(keywords string, page, pageSize int) (videos []model.Video) {
	mysqlClient.Where("status = ? and title like ?", common.AUDIT_APPROVED, "%"+keywords+"%").
		Limit(pageSize).Offset((page - 1) * pageSize).Find(&videos)

	return
}

// 通过视频状态查询视频列表
func SelectVideoListByStatus(page, pageSize, status int) (total int64, videos []model.Video) {
	client := mysqlClient.Where("status = ?", status)
	client.Model(&model.Video{}).Count(&total)
	client.Limit(pageSize).Offset((page - 1) * pageSize).Find(&videos)
	return
}

// 管理员通过关键词查询视频
func AdminSelectVideoListByKeywords(keywords string, page, pageSize int) (total int64, videos []model.Video) {
	search := mysqlClient.Where("status = ? and title like ?", common.AUDIT_APPROVED, "%"+keywords+"%")
	search.Model(&model.Video{}).Count(&total)
	search.Limit(pageSize).Offset((page - 1) * pageSize).Find(&videos)

	return
}

// 查询点击量高的视频
func SelectVideoListByClicks(pageSize int) (videos []model.Video) {
	mysqlClient.Debug().Where("status = ?", common.AUDIT_APPROVED).Limit(pageSize).Order("clicks desc").Find(&videos)
	return
}

// 查询视频播放量
func SelectVideoClicks(videoId uint) (clicks int64) {
	mysqlClient.Model(model.Video{}).Where("id = ?", videoId).Pluck("clicks", &clicks)
	return
}

// 查询视频作者
func SelectVideoAuthorId(videoId uint) (userId uint) {
	mysqlClient.Model(model.Video{}).Where("id = ?", videoId).Pluck("uid", &userId)
	return
}

// 查询用户视频
func SelectVideoByUserId(userId uint, page, pageSize int) (total int64, videos []model.Video) {
	offset := (page - 1) * pageSize
	mysqlClient.Model(&model.Video{}).Where("uid = ? and status = ?", userId, common.AUDIT_APPROVED).Count(&total)
	mysqlClient.Where("uid = ? and status = ?", userId, common.AUDIT_APPROVED).Limit(pageSize).Offset(offset).Find(&videos)
	return
}

// 查询上传的视频
func SelectUploadVideo(userId uint, page, pageSize int) (total int64, videos []model.Video) {
	offset := (page - 1) * pageSize
	mysqlClient.Model(&model.Video{}).Where("uid = ?", userId).Count(&total)
	mysqlClient.Where("uid = ?", userId).Limit(pageSize).Offset(offset).Find(&videos)
	return
}

// 更新视频信息
func UpdateVideoInfo(modifyDTO dto.ModifyVideoDTO) error {
	if err := mysqlClient.Model(&model.Video{}).Where("id = ?", modifyDTO.VID).Updates(
		map[string]interface{}{
			"title":     modifyDTO.Title,
			"cover":     modifyDTO.Cover,
			"desc":      modifyDTO.Desc,
			"copyright": modifyDTO.Copyright,
			"status":    common.WAITING_REVIEW,
		},
	).Error; err != nil {
		return err
	}

	// 移除缓存
	cache.DelVideo(modifyDTO.VID)

	return nil
}

// 更新播放量
func UpdateClicks(videoId uint, clicks int64) error {
	err := mysqlClient.Model(&model.Video{}).Where("id = ?", videoId).Update("clicks", clicks).Error
	if err != nil {
		return err
	}
	return nil
}

// 更新视频状态
func UpadteVideoStatus(videoId uint, status int) error {
	err := mysqlClient.Model(&model.Video{}).Where("id = ?", videoId).Update("status", status).Error
	if err != nil {
		return err
	}

	// 移除缓存
	cache.DelVideo(videoId)

	return nil
}

// 删除视频
func DeleteVideo(id uint) {
	cache.DelVideo(id)
	mysqlClient.Where("id = ?", id).Delete(&model.Video{})
}

// 视频是否属于用户
func IsVideoBelongUser(id, userId uint) bool {
	var video model.Video
	mysqlClient.Where("id = ? and uid = ?", id, userId).First(&video)

	return video.ID != 0
}
