package cache

import (
	"encoding/json"
	"strconv"
	"time"

	"go.uber.org/zap"
	"kuukaa.fun/leaf/domain/model"
)

func GetUser(id uint) (user model.User) {
	jsonStr := Get(USER_KEY + strconv.FormatUint(uint64(id), 10))
	// 反序列化
	if err := json.Unmarshal([]byte(jsonStr), &user); err != nil {
		zap.L().Error("用户信息反序列化失败")
	}
	return
}
func SetUser(user model.User) {
	Set(USER_KEY+strconv.FormatUint(uint64(user.ID), 10), user,
		time.Hour*USER_EXPIRATION_TIME)
}
