package vo

import (
	"time"

	"kuukaa.fun/leaf/domain/model"
)

type AtMessageVO struct {
	CreatedAt time.Time   `json:"created_at"`
	User      BaseUserVO  `json:"user"`
	Video     BaseVideoVO `json:"video"`
}

func ToAtMessageVoList(messages []model.AtMessage) []AtMessageVO {
	length := len(messages)
	newAtMessages := make([]AtMessageVO, length)

	for i := 0; i < length; i++ {
		newAtMessages[i].User = ToBaseUserVO(messages[i].User)
		newAtMessages[i].Video = ToBaseVideoVO(messages[i].Video)
		newAtMessages[i].CreatedAt = messages[i].CreatedAt
	}

	return newAtMessages
}
