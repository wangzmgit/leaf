package service

import (
	"kuukaa.fun/leaf/common"
	"kuukaa.fun/leaf/domain/dto"
	"kuukaa.fun/leaf/domain/model"
)

func InsertVideo(video model.Video) (uint, error) {
	err := mysqlClient.Create(&video).Error
	return video.ID, err
}

func SelectVideoByID(videoId uint) (video model.Video) {
	mysqlClient.First(&video, videoId)
	return
}

func SelectVideoClicks(videoId uint) (clicks int64) {
	mysqlClient.Model(model.Video{}).Where("id = ?", videoId).Pluck("clicks", &clicks)
	return
}

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
	return nil
}

func UpdateClicks(videoId uint, clicks int64) error {
	err := mysqlClient.Model(&model.Video{}).Where("id = ?", videoId).Update("clicks", clicks).Error
	if err != nil {
		return err
	}
	return nil
}

func UpadteVideoStatus(videoId uint, status int) error {
	err := mysqlClient.Model(&model.Video{}).Where("id = ?", videoId).Update("status", status).Error
	if err != nil {
		return err
	}
	return nil
}

// 删除视频
func DeleteVideo(id uint) {
	mysqlClient.Where("id = ?", id).Delete(&model.Video{})
}
