package cache

import (
	"fmt"
	"time"

	"kuukaa.fun/leaf/util/convert"
)

func IsTokenExist(userId uint, token string) bool {
	return ZScore(TOKEN_KEY+convert.UintToString(userId), token) != 0
}

func SetToken(id uint, token string) {
	key := TOKEN_KEY + convert.UintToString(id)
	fmt.Println(ZCard(key))
	if ZCard(key) >= MAX_LOGIN_LIMIT {
		// 保留MAX_LOGIN_LIMIT - 1个token
		ZRemRangeByRank(key, 0, 1-MAX_LOGIN_LIMIT)
	}

	ZAdd(key, float64(time.Now().Add(TOKEN_EXPRIRATION_TIME*time.Hour).Unix()), token)
}

func DelToken(id uint) {
	Del(TOKEN_KEY + convert.UintToString(id))
}
