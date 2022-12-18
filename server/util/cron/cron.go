package cron

import (
	"strings"
	"time"

	"github.com/jasonlvhit/gocron"
	"go.uber.org/zap"
	"kuukaa.fun/leaf/cache"
	"kuukaa.fun/leaf/service"
	"kuukaa.fun/leaf/util/convert"
)

func Init() {
	c := gocron.NewScheduler()

	// 每天凌晨2点同步播放量数据
	c.Every(1).Days().At("2:00").Do(syncClicks)

	<-c.Start()
}

// 同步播放量到数据库
func syncClicks() {
	start := time.Now() // 获取当前时间
	zap.L().Info("开始同步播放量")
	keys := cache.GetClicksKeys()
	for _, v := range keys {
		// 获取视频ID和播放量
		videoId := convert.StringToUint(v[strings.LastIndex(v, ":")+1:])
		clicks, err := cache.GetClicks(videoId)
		if err != nil {
			cache.DelClicks(videoId)
			continue
		}
		// 写入数据库
		service.UpdateClicks(videoId, clicks)

		// 如果缓存过期时间小于 22 小时, 删除缓存
		if cache.ClickTTL(videoId) < time.Hour*22 {
			cache.DelClicks(videoId)
		}
	}
	zap.L().Info("播放量同步完成，耗时 " + time.Since(start).String())
}
