package vo

import (
	"time"

	"kuukaa.fun/leaf/domain/model"
)

type ReplyMessageVO struct {
	CreatedAt          time.Time   `json:"created_at"`
	Content            string      `json:"content"`
	TargetReplyContent string      `json:"target_reply_content"`
	RootContent        string      `json:"root_content"`
	CommentId          string      `json:"comment_id"`
	User               BaseUserVO  `json:"user"`
	Video              BaseVideoVO `json:"video"`
}

func ToReplyMessageVoList(messages []model.ReplyMessage) []ReplyMessageVO {
	length := len(messages)
	newReplyMessages := make([]ReplyMessageVO, length)

	for i := 0; i < length; i++ {
		newReplyMessages[i].Content = messages[i].Content
		newReplyMessages[i].TargetReplyContent = messages[i].TargetReplyContent
		newReplyMessages[i].RootContent = messages[i].RootContent
		newReplyMessages[i].CommentId = messages[i].CommentId
		newReplyMessages[i].User = ToBaseUserVO(messages[i].User)
		newReplyMessages[i].Video = ToBaseVideoVO(messages[i].Video)
		newReplyMessages[i].CreatedAt = messages[i].CreatedAt
	}

	return newReplyMessages
}
