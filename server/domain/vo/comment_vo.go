package vo

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"kuukaa.fun/leaf/domain/model"
)

type CommentVo struct {
	ID        primitive.ObjectID `json:"id"`
	Content   string             `json:"content"`
	Author    UserVO             `json:"author"`
	Reply     []ReplyVo          `json:"reply"`
	CreatedAt int64              `json:"created_at"`
}

type ReplyVo struct {
	ID        primitive.ObjectID `json:"id"`
	Content   string             `json:"content"`
	Author    UserVO             `json:"author"`
	CreatedAt int64              `json:"created_at"`
}

func ToCommentVO(comments []model.Comment) []CommentVo {
	length := len(comments)
	newComments := make([]CommentVo, length)
	for i := 0; i < length; i++ {
		newComments[i].ID = comments[i].ID
		newComments[i].Content = comments[i].Content
		newComments[i].CreatedAt = comments[i].CreatedAt
		newComments[i].Author = ToUserVO(comments[i].Author)
		newComments[i].Reply = ToReplyVO(comments[i].Reply)
	}
	return newComments
}

func ToReplyVO(replies []model.Reply) []ReplyVo {
	length := len(replies)
	newReplies := make([]ReplyVo, length)
	for i := 0; i < length; i++ {
		newReplies[i].ID = replies[i].ID
		newReplies[i].Content = replies[i].Content
		newReplies[i].CreatedAt = replies[i].CreatedAt
		newReplies[i].Author = ToUserVO(replies[i].Author)

	}
	return newReplies
}
