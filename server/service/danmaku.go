package service

import "kuukaa.fun/leaf/domain/model"

func InsertDanmaku(danmaku model.Danmaku) error {
	return mysqlClient.Create(&danmaku).Error
}

func SelectDanmakuByVidAndPart(videoId, part int) (danmaku []model.Danmaku) {
	mysqlClient.Where("vid = ? and part = ?", videoId, part).Order("time").Find(&danmaku)
	return
}
