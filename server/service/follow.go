package service

import (
	"kuukaa.fun/leaf/domain/model"
)

// 关注
func InsertFollow(follow model.Follow) error {
	return mysqlClient.Create(&follow).Error
}

// 取消关注
func DeleteFollow(uid uint, fid uint) error {
	return mysqlClient.Where("uid = ? and fid = ?", uid, fid).Delete(&model.Follow{}).Error
}

// 获取关注状态
func IsFollow(uid uint, fid uint) bool {
	var follow model.Follow
	mysqlClient.Where("uid = ? and fid = ?", uid, fid).First(&follow)

	return follow.ID != 0
}

// 通过UID获取关注列表
func SelectFollowingUser(userId uint, page, pageSize int) (users []model.User) {
	offset := (page - 1) * pageSize
	userIds := mysqlClient.Model(&model.Follow{}).Select("fid").Where("uid = ?", userId)
	mysqlClient.Where("id in (?)", userIds).Limit(pageSize).Offset(offset).Find(&users)
	return
}

// 通过UID获取粉丝列表
func SelectFollowerUser(userId uint, page int, pageSize int) (users []model.User) {
	offset := (page - 1) * pageSize
	userIds := mysqlClient.Model(&model.Follow{}).Select("uid").Where("fid = ?", userId)
	mysqlClient.Where("id in (?)", userIds).Limit(pageSize).Offset(offset).Find(&users)
	return
}

// 通过UID获取粉丝数
func GetFollowCount(uid uint) (following, follower int64) {
	mysqlClient.Model(&model.Follow{}).Where("uid = ?", uid).Count(&following)
	mysqlClient.Model(&model.Follow{}).Where("fid = ?", uid).Count(&follower)
	return
}
