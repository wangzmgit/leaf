package service

import (
	"gorm.io/gorm"
	"kuukaa.fun/leaf/common"
	"kuukaa.fun/leaf/domain/dto"
	"kuukaa.fun/leaf/domain/model"
)

func InsertVideo(video model.Video) uint {
	mysqlClient.Create(&video)
	return video.ID
}

func SelectVideoByID(videoId uint) (video model.Video) {
	mysqlClient.First(&video, videoId)
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

func UpadteVideoStatus(videoId uint, status int) error {
	err := mysqlClient.Model(&model.Video{}).Where("id = ?", videoId).Update("status", status).Error
	if err != nil {
		return err
	}
	return nil
}

func AddClicks(videoId uint) {
	mysqlClient.Model(&model.Video{}).Where("id = ?", videoId).UpdateColumn("clicks", gorm.Expr("clicks + 1"))
}
