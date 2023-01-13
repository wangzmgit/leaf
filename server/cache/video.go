package cache

import (
	"encoding/json"
	"time"

	"go.uber.org/zap"
	"kuukaa.fun/leaf/domain/model"
	"kuukaa.fun/leaf/util/convert"
)

func GetVideo(id uint) (video model.Video) {
	jsonStr := Get(VIDEO_KEY + convert.UintToString(id))
	// 反序列化
	if err := json.Unmarshal([]byte(jsonStr), &video); err != nil {
		zap.L().Error("视频信息反序列化失败: " + err.Error())
	}
	return
}

func SetVideo(video model.Video) {
	//先序列化video
	ub, err := json.Marshal(video)
	if err != nil {
		zap.L().Error("视频信息序列化失败: " + err.Error())
		return
	}
	Set(VIDEO_KEY+convert.UintToString(video.ID), ub,
		time.Hour*VIDEO_EXPIRATION_TIME)
}

func DelVideo(id uint) {
	Del(VIDEO_KEY + convert.UintToString(id))
}
