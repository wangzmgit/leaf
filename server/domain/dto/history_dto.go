package dto

import (
	"kuukaa.fun/leaf/domain/model"
)

type HistoryDTO struct {
	Vid  uint
	Part uint
	Time float64
}

/**
 * historyDTO结构体转化为History结构体
 * param: historyDTO historyDTO
 * param: userID 用户ID
 * return: History结构体
 */
func HistoryDtoToHistory(historyDTO HistoryDTO, userId uint) model.History {
	return model.History{
		Vid:  historyDTO.Vid,
		Uid:  userId,
		Part: historyDTO.Part,
		Time: historyDTO.Time,
	}
}
