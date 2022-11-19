package cache

import (
	"strconv"
	"time"

	"go.uber.org/zap"
)

func GetLoginTryCount(email string) int {
	s := Get(LOGIN_TRY_COUNT_KEY + email)
	count, err := strconv.Atoi(s)
	if err != nil {
		zap.L().Error("缓存中用户登录次数转换为int类型失败")
	}
	return count
}

func SetLoginTryCount(email string, count int) {
	Set(LOGIN_TRY_COUNT_KEY+email, count, time.Minute*LOGIN_TRY_COUNT_EXPRIRATION_TIME)
}

func DelLoginTryCount(email string) {
	Del(LOGIN_TRY_COUNT_KEY + email)
}

