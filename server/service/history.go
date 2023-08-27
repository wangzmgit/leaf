package service

import (
	"kuukaa.fun/leaf/domain/model"
)

func InsertHistory(history model.History) error {
	return mysqlClient.Create(&history).Error
}

func InsertOrUpdateHistory(history model.History) error {
	id := SelectHistory(history.Vid, history.Uid).ID
	if id == 0 {
		return InsertHistory(history)
	}

	// 更新
	return mysqlClient.Model(&model.History{}).Where("uid = ? and vid = ?", history.Uid, history.Vid).Updates(
		map[string]interface{}{
			"time": history.Time,
			"part": history.Part,
		},
	).Error
}

func SelectHistory(videoId, userId uint) (history model.History) {
	mysqlClient.Where("uid = ? and vid = ?", userId, videoId).First(&history)
	return
}

func SelectHistoryVideo(userId uint, page, pageSize int) (history []model.History) {
	mysqlClient.Where("uid = ?", userId).Order("updated_at desc").
		Limit(pageSize).Offset((page - 1) * pageSize).Find(&history)
	return
}
