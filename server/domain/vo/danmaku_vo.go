package vo

import "kuukaa.fun/leaf/domain/model"

type DanmakuVO struct {
	Time  uint   `json:"time"`
	Type  int    `json:"type"`
	Color string `json:"color"`
	Text  string `json:"text"`
}

func ToDanmakuVoList(danmaku []model.Danmaku) []DanmakuVO {
	length := len(danmaku)
	newDanmaku := make([]DanmakuVO, length)
	for i := 0; i < length; i++ {
		newDanmaku[i].Time = danmaku[i].Time
		newDanmaku[i].Type = danmaku[i].Type
		newDanmaku[i].Color = danmaku[i].Color
		newDanmaku[i].Text = danmaku[i].Text
	}

	return newDanmaku
}
