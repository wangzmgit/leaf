package service

import (
	"kuukaa.fun/leaf/domain/model"
)

func InsertAnnounce(announce model.Announce) error {
	return mysqlClient.Create(&announce).Error
}

// 删除公告
func DeleteAnnounce(id uint) {
	mysqlClient.Where("id = ?", id).Delete(&model.Announce{})
}

// 查询公告
func SelectAnnounce(page, pageSize int) (total int64, announces []model.Announce) {
	mysqlClient.Model(&model.Announce{}).Count(&total)
	mysqlClient.Limit(pageSize).Offset((page - 1) * pageSize).Find(&announces)
	return
}
