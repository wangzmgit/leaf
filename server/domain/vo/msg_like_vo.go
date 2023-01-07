package vo

import (
	"time"

	"kuukaa.fun/leaf/domain/model"
)

type LikeMessageVO struct {
	CreatedAt time.Time   `json:"created_at"`
	User      BaseUserVO  `json:"user"`
	Video     BaseVideoVO `json:"video"`
}

func ToLikeMessageVoList(messages []model.LikeMessage) []LikeMessageVO {
	length := len(messages)
	newLikeMessages := make([]LikeMessageVO, length)

	for i := 0; i < length; i++ {
		newLikeMessages[i].User = ToBaseUserVO(messages[i].User)
		newLikeMessages[i].Video = ToBaseVideoVO(messages[i].Video)
		newLikeMessages[i].CreatedAt = messages[i].CreatedAt
	}

	return newLikeMessages
}
