package cache

import (
	"strconv"
	"time"

	"go.uber.org/zap"
)

func GetUploadImage(url string) uint {
	s := Get(UPLOAD_IMAGE_KEY + url)
	userId, err := strconv.Atoi(s)
	if err != nil {
		zap.L().Error("缓存中用户id转换为uint类型失败")
	}
	return uint(userId)
}

func SetUploadImage(url string, userID uint) {
	Set(UPLOAD_IMAGE_KEY+url, userID, time.Minute*UPLOAD_IMAGE_EXPRIRATION_TIME)
}

func DelUploadImage(email string) {
	Del(UPLOAD_IMAGE_KEY + email)
}
