package service

import (
	"kuukaa.fun/leaf/cache"
)

// 获取播放量
func GetVideoClicks(videoId uint) int64 {
	clicks, err := cache.GetClicks(videoId)
	if err != nil {
		// 从数据库中查询并写入缓存
		clicks = SelectVideoClicks(videoId)
		cache.SetClicks(videoId, clicks)
	}

	return clicks
}

// 增加播放量
func AddVideoClicks(videoId uint, ip string) {
	// 一个ip在同一个视频下，每30分钟可重新增加1播放量
	if cache.GetClicksLimit(videoId, ip) == "" {
		cache.AddClicks(videoId)
		cache.SetClicksLimit(videoId, ip)
	}
}
