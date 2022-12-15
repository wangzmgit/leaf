package cache

import (
	"time"

	"kuukaa.fun/leaf/util/convert"
)

func GetClicks(videoId uint, ip string) string {
	s := Get(VIDEO_CLICKS_KEY + convert.UintToString(videoId) + ":" + ip)

	return s
}

func SetClicks(videoId uint, ip string) {
	Set(VIDEO_CLICKS_KEY+convert.UintToString(videoId)+":"+ip, 1,
		time.Minute*VIDEO_CLICKS_EXPRIRATION_TIME)
}
