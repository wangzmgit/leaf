package cache

import (
	"fmt"
	"time"

	"kuukaa.fun/leaf/util/convert"
)

func GetAccessToken(id uint) string {
	return Get(ACCESS_TOKEN_KEY + convert.UintToString(id))
}

func SetAccessToken(id uint, token string) {
	Set(ACCESS_TOKEN_KEY+convert.UintToString(id), token,
		time.Minute*ACCESS_TOKEN_EXPRIRATION_TIME)
}

func DelAccessToken(id uint) {
	Del(ACCESS_TOKEN_KEY + convert.UintToString(id))
}

func IsRefreshTokenExist(userId uint, token string) bool {
	return ZScore(REFRESH_TOKEN_KEY+convert.UintToString(userId), token) != 0
}

func SetRefreshToken(id uint, token string) {
	key := REFRESH_TOKEN_KEY + convert.UintToString(id)
	fmt.Println(ZCard(key))
	if ZCard(key) >= MAX_LOGIN_LIMIT {
		// 保留MAX_LOGIN_LIMIT - 1个token
		ZRemRangeByRank(key, 0, 1-MAX_LOGIN_LIMIT)
	}

	ZAdd(key, float64(time.Now().Add(REFRESH_TOKEN_EXPRIRATION_TIME*time.Hour).Unix()), token)
}

func DelRefreshToken(id uint) {
	Del(REFRESH_TOKEN_KEY + convert.UintToString(id))
}
