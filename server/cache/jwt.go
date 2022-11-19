package cache

import (
	"strconv"
	"time"
)

func GetAccessToken(id uint) string {
	return Get(ACCESS_TOKEN_KEY + strconv.FormatUint(uint64(id), 10))
}
func SetAccessToken(id uint, token string) {
	Set(ACCESS_TOKEN_KEY+strconv.FormatUint(uint64(id), 10), token,
		time.Minute*ACCESS_TOKEN_EXPRIRATION_TIME)
}
func DelAccessToken(id uint) {
	Del(ACCESS_TOKEN_KEY + strconv.FormatUint(uint64(id), 10))
}

func GetRefreshToken(id uint) string {
	return Get(REFRESH_TOKEN_KEY + strconv.FormatUint(uint64(id), 10))
}
func SetRefreshToken(id uint, token string) {
	Set(REFRESH_TOKEN_KEY+strconv.FormatUint(uint64(id), 10), token,
		time.Hour*REFRESH_TOKEN_EXPRIRATION_TIME)
}
func DelRefreshToken(id uint) {
	Del(REFRESH_TOKEN_KEY + strconv.FormatUint(uint64(id), 10))
}
