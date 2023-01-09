package service

import "kuukaa.fun/leaf/domain/model"

func InsertManyWhisper(messages []model.Whisper) error {
	return mysqlClient.Create(&messages).Error
}

// 查询消息列表
func SelectWhisper(userId, fid uint, page, pageSize int) (messages []model.Whisper) {
	offset := (page - 1) * pageSize
	mysqlClient.Where("uid = ? and fid = ?", userId, fid).Limit(pageSize).Offset(offset).Order("id desc").Find(&messages)
	return
}

// 查询消息(分组)列表
func SelectWhisperGroup(userId uint) (messages []model.Whisper) {
	messageIds := mysqlClient.Model(&model.Whisper{}).Select("max(id)").Where("uid = ?", userId).Group("fid")
	mysqlClient.Where("id in (?)", messageIds).Order("created_at desc").Find(&messages)

	return
}

// 更新消息已读状态
func UpdateWhisperStatus(userId, fid uint) error {
	return mysqlClient.Model(&model.Whisper{}).Where("uid = ? and fid = ?", userId, fid).Update("status", true).Error
}
