package service

import (
	"kuukaa.fun/leaf/common"
	"kuukaa.fun/leaf/domain/dto"
	"kuukaa.fun/leaf/domain/model"
)

func InsertResource(resource model.Resource) uint {
	mysqlClient.Create(&resource)
	return resource.ID
}

func SelectResourceByID(resourceId uint) (resource model.Resource) {
	mysqlClient.First(&resource, resourceId)
	return
}

// 获取资源数量(按视频)
func SelectResourceCountByVid(videoId uint) (count int64) {
	mysqlClient.Model(&model.Resource{}).Where("vid = ? ", videoId).Count(&count)
	return
}

// 获取资源数量(按状态)
func SelectResourceCountByStatus(videoId uint, status int) (count int64) {
	mysqlClient.Model(&model.Resource{}).Where("vid = ? and status = ?", videoId, status).Count(&count)
	return
}

// 通过视频ID和资源状态获取资源
func SelectResourceByVideo(videoId uint, pass bool) (resources []model.Resource) {
	db := mysqlClient.Model(&model.Resource{}).Where("vid = ?", videoId)
	if pass {
		db.Where("status = ?", common.AUDIT_APPROVED)
	}

	db.Find(&resources)
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

// 修改资源标题
func ModifyResourceTitleService(modifyDTO dto.ModifyResourceTitleDTO) error {
	err := mysqlClient.Model(&model.Resource{}).Where("id = ?", modifyDTO.ID).Update("title", modifyDTO.Title).Error
	if err != nil {
		return err
	}
	return nil
}

// 删除资源
func DeleteResource(id uint) {
	mysqlClient.Where("id = ?", id).Delete(&model.Resource{})
}
