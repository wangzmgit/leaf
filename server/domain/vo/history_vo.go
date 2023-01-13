package vo

import (
	"time"

	"kuukaa.fun/leaf/domain/model"
)

type HistoryProgressVO struct {
	Time float64 `json:"time"`
	Part uint    `json:"part"`
}

type HistoryVideoVO struct {
	ID        uint        `json:"id"`
	Time      float64     `json:"time"`
	Part      uint        `json:"part"`
	Video     BaseVideoVO `json:"video"`
	CreatedAt time.Time   `json:"created_at"`
}

func ToHistoryProgressVO(history model.History) HistoryProgressVO {
	return HistoryProgressVO{
		Time: history.Time,
		Part: history.Part,
	}
}

func ToHistoryVideoVoList(historyList []model.History) []HistoryVideoVO {
	length := len(historyList)
	newHistoryList := make([]HistoryVideoVO, length)
	for i := 0; i < length; i++ {
		newHistoryList[i].ID = historyList[i].ID
		newHistoryList[i].Part = historyList[i].Part
		newHistoryList[i].Time = historyList[i].Time
		newHistoryList[i].Video = ToBaseVideoVO(historyList[i].Video)
		newHistoryList[i].CreatedAt = historyList[i].UpdatedAt
	}

	return newHistoryList
}
