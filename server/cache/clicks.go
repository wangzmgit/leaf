package cache

import (
	"strconv"
	"time"

	"go.uber.org/zap"
	"kuukaa.fun/leaf/util/convert"
)

func GetClicksLimit(videoId uint, ip string) string {
	s := Get(VIDEO_CLICKS_LIMIT_KEY + convert.UintToString(videoId) + ":" + ip)

	return s
}

func SetClicksLimit(videoId uint, ip string) {
	Set(VIDEO_CLICKS_LIMIT_KEY+convert.UintToString(videoId)+":"+ip, 1,
		time.Minute*VIDEO_CLICKS_LIMIT_EXPRIRATION_TIME)
}

func GetClicks(videoId uint) (int64, error) {
	s := Get(VIDEO_CLICKS_KEY + convert.UintToString(videoId))
	count, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		zap.L().Error("缓存中播放量转换为int64类型失败")
	}
	return count, err
}

func SetClicks(videoId uint, count int64) {
	Set(VIDEO_CLICKS_KEY+convert.UintToString(videoId), count,
		time.Hour*VIDEO_CLICKS_EXPRIRATION_TIME)
}

// 删除播放量
func DelClicks(videoId uint) {
	Del(VIDEO_CLICKS_KEY + convert.UintToString(videoId))
}

func AddClicks(videoId uint) {
	Incr(VIDEO_CLICKS_KEY + convert.UintToString(videoId))
}

func GetClicksKeys() []string {
	return Keys(VIDEO_CLICKS_KEY + "*")
}

// 获取点击量过期时间
func ClickTTL(videoId uint) time.Duration {
	return TTL(VIDEO_CLICKS_KEY + convert.UintToString(videoId))
}
