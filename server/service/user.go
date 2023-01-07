package service

import (
	"time"

	"go.uber.org/zap"
	"kuukaa.fun/leaf/cache"
	"kuukaa.fun/leaf/domain/dto"
	"kuukaa.fun/leaf/domain/model"
)

func InsertUser(user model.User) {
	mysqlClient.Create(&user)
}

// 获取用户信息 (通过邮箱，只查数据库)
func SelectUserByEmail(email string) (user model.User) {
	mysqlClient.Where(model.User{Email: email}).First(&user)
	// 存入缓存
	cache.SetUser(user)
	return
}

// 获取用户信息 (通过ID，只查数据库)
func SelectUserByID(userId uint) (user model.User) {
	mysqlClient.First(&user, userId)
	// 存入缓存
	cache.SetUser(user)
	return
}

// 获取用户信息 (通过ID，先查缓存)
func GetUserInfo(userId uint) (user model.User) {
	user = cache.GetUser(userId)
	if user.ID == 0 {
		user = SelectUserByID(userId)
	}

	return
}

// 查询用户名是否存在
func IsNameExist(name string) (bool, uint) {
	var user model.User
	mysqlClient.Where("username = ?", name).First(&user)
	if user.ID != 0 {
		return true, user.ID
	}
	return false, uint(0)
}

// 更新用户信息
func UpdateUserInfo(userId uint, modifyDTO dto.ModifyUserInfoDTO) error {
	birthday, err := time.Parse("2006-01-02", modifyDTO.Birthday)
	if err != nil {
		zap.L().Error("日期转换失败，原始值: " + modifyDTO.Birthday + "，err: " + err.Error())
		birthday = time.Unix(0, 0)
	}
	//移除缓存
	cache.DelUser(userId)
	if err := mysqlClient.Model(&model.User{}).Where("id = ?", userId).Updates(
		map[string]interface{}{
			"avatar":   modifyDTO.Avatar,
			"username": modifyDTO.Name,
			"gender":   modifyDTO.Gender,
			"birthday": birthday,
			"sign":     modifyDTO.Sign,
		},
	).Error; err != nil {
		return err
	}
	return nil
}

// 更新用户空间封面
func UpdateUserSpaceCover(userId uint, url string) error {
	//移除缓存
	cache.DelUser(userId)
	err := mysqlClient.Model(&model.User{}).Where("id = ?", userId).Update("space_cover", url).Error
	if err != nil {
		return err
	}
	return nil
}

// 通过用户名查询多个用户
func SelectUserIdsByName(names []string) (ids []uint) {
	mysqlClient.Model(model.User{}).Where("username in (?)", names).Pluck("id", &ids)
	return
}

// 通过用户名查询一个用户
func SelectUserIdByName(name string) uint {
	var user model.User
	mysqlClient.Where("username = ?", name).First(&user)

	return user.ID
}
