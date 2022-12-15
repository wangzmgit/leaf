package service

import "kuukaa.fun/leaf/domain/model"

func InsertResource(resource model.Resource) uint {
	mysqlClient.Create(&resource)
	return resource.ID
}

func SelectResourceByID(resourceId uint) (resource model.Resource) {
	mysqlClient.First(&resource, resourceId)
	return
}

// 获取资源数量
func SelectResourceCount(resourceId uint, status int) (count int64) {
	mysqlClient.Model(&model.Resource{}).Where("id = ? and status = ?", resourceId, status).Count(&count)
	return
}

// 通过视频ID和资源状态获取资源
func SelectResourceByVideo(videoId uint) (resources []model.Resource) {
	mysqlClient.Model(&model.Resource{}).Where("vid = ?", videoId).Find(&resources)
	return
}

// 更新资源状态
func UpadteResourceStatus(resourceId uint, status int) error {
	err := mysqlClient.Model(&model.Resource{}).Where("id = ?", resourceId).Update("status", status).Error
	if err != nil {
		return err
	}
	return nil
}
