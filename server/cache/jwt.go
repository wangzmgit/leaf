package cache

import (
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

func GetRefreshToken(id uint) string {
	return Get(REFRESH_TOKEN_KEY + convert.UintToString(id))
}

func SetRefreshToken(id uint, token string) {
	Set(REFRESH_TOKEN_KEY+convert.UintToString(id), token,
		time.Hour*REFRESH_TOKEN_EXPRIRATION_TIME)
}

func DelRefreshToken(id uint) {
	Del(REFRESH_TOKEN_KEY + convert.UintToString(id))
}
