package service

import "kuukaa.fun/leaf/domain/model"

func InsertLikeMessage(likeMessage model.LikeMessage) error {
	return mysqlClient.Create(&likeMessage).Error
}

func DeleteLikeMessage(videoId, userId uint) {
	mysqlClient.Where("vid = ? and fid = ?", videoId, userId).Delete(&model.LikeMessage{})
}

func SelectLikeMessage(userId uint, page, pageSize int) (messages []model.LikeMessage) {
	mysqlClient.Where("uid = ?", userId).Limit(pageSize).Offset((page - 1) * pageSize).Find(&messages)
	return
}
