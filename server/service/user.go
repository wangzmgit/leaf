package service

import (
	"kuukaa.fun/leaf/cache"
	"kuukaa.fun/leaf/domain/model"
)

func InsertUser(user model.User) {
	mysqlClient.Create(&user)
}

func SelectUserByEmail(email string) (user model.User) {
	mysqlClient.Where(model.User{Email: email}).First(&user)
	// 存入缓存
	cache.SetUser(user)
	return
}
