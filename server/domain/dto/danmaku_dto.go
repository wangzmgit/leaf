package dto

import "kuukaa.fun/leaf/domain/model"

type DanmakuDTO struct {
	Vid   uint
	Part  uint
	Time  uint //时间
	Type  int  //类型0滚动;1顶部;2底部
	Color string
	Text  string
}

/**
 * DanmakuDTO结构体转化为Danmaku结构体
 * param: danmakuDTO DanmakuDTO
 * param: userID 用户ID
 * return: Danmaku结构体
 */
func DanmakuDtoToDanmaku(danmakuDTO DanmakuDTO, userId uint) model.Danmaku {
	return model.Danmaku{
		Vid:   danmakuDTO.Vid,
		Uid:   userId,
		Part:  danmakuDTO.Part,
		Time:  danmakuDTO.Time,
		Type:  danmakuDTO.Type,
		Color: danmakuDTO.Color,
		Text:  danmakuDTO.Text,
	}
}
