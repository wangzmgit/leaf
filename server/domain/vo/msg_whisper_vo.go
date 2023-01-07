package vo

import (
	"time"

	"kuukaa.fun/leaf/domain/model"
)

// 消息列表
type WhisperGroupVO struct {
	CreatedAt time.Time  `json:"created_at"`
	User      BaseUserVO `json:"user"`
	Status    bool       `json:"status"` //已读状态
}

// 消息详情
type WhisperDetailsVO struct {
	Fid       uint      `json:"fid"`
	FromId    uint      `json:"from_id"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

// 单条消息
type WhisperVO struct {
	Fid     uint   `json:"fid"`
	Content string `json:"content"`
}

func ToWhisperGroupVoList(messages []model.Whisper, users []model.User) []WhisperGroupVO {
	length := len(messages)
	newWhisperGroup := make([]WhisperGroupVO, length)
	for i := 0; i < length; i++ {
		newWhisperGroup[i].CreatedAt = messages[i].CreatedAt
		newWhisperGroup[i].Status = messages[i].Status
		newWhisperGroup[i].User = ToBaseUserVO(users[i])
	}
	return newWhisperGroup
}

func ToMsgDetailsVoList(messages []model.Whisper) []WhisperDetailsVO {
	length := len(messages)
	newWhisperDetails := make([]WhisperDetailsVO, length)
	for i := 0; i < length; i++ {
		newWhisperDetails[i].Fid = messages[i].Fid
		newWhisperDetails[i].FromId = messages[i].FromId
		newWhisperDetails[i].Content = messages[i].Content
		newWhisperDetails[i].CreatedAt = messages[i].CreatedAt
	}
	return newWhisperDetails
}
