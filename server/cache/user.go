package cache

import (
	"encoding/json"
	"time"

	"go.uber.org/zap"
	"kuukaa.fun/leaf/domain/model"
	"kuukaa.fun/leaf/util/convert"
)

func GetUser(id uint) (user model.User) {
	jsonStr := Get(USER_KEY + convert.UintToString(id))
	// 反序列化
	if err := json.Unmarshal([]byte(jsonStr), &user); err != nil {
		zap.L().Error("用户信息反序列化失败: " + err.Error())
	}
	return
}

func SetUser(user model.User) {
	//先序列化user
	ub, err := json.Marshal(user)
	if err != nil {
		zap.L().Error("用户信息序列化失败: " + err.Error())
		return
	}
	Set(USER_KEY+convert.UintToString(user.ID), ub,
		time.Hour*USER_EXPIRATION_TIME)
}

func DelUser(id uint) {
	Del(USER_KEY + convert.UintToString(id))
}
