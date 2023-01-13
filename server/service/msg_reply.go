package service

import "kuukaa.fun/leaf/domain/model"

func InsertReplyMessage(replyMessage model.ReplyMessage) error {
	return mysqlClient.Create(&replyMessage).Error
}

func SelectReplyMessage(userId uint, page, pageSize int) (messages []model.ReplyMessage) {
	mysqlClient.Where("uid = ?", userId).Limit(pageSize).Offset((page - 1) * pageSize).Find(&messages)
	return
}
