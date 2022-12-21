package dto

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"kuukaa.fun/leaf/domain/model"
)

type CommentDTO struct {
	Vid      uint
	Content  string
	ParentID primitive.ObjectID
}

type DeleteReplyDTO struct {
	CommentID primitive.ObjectID
	ReplyID   primitive.ObjectID
}

/**
 * 评论DTO结构体转化为Comment结构体
 * param: commentDTO 评论DTO结构体
 * return: comment结构体
 */
func CommentDtoToComment(commentDTO CommentDTO, userId uint) model.Comment {
	return model.Comment{
		ID:        primitive.NewObjectID(),
		Vid:       commentDTO.Vid,
		CreatedAt: time.Now().UnixMilli(),
		Content:   commentDTO.Content,
		Uid:       userId,
		Reply:     []model.Reply{},
		IsDelete:  false,
	}
}

/**
 * 评论DTO结构体转化为Reply结构体
 * param: commentDTO 评论DTO结构体
 * return: reply结构体
 */
func CommentDtoToReply(commentDTO CommentDTO, userId uint) model.Reply {
	return model.Reply{
		ID:        primitive.NewObjectID(),
		CreatedAt: time.Now().UnixMilli(),
		Content:   commentDTO.Content,
		Uid:       userId,
		IsDelete:  false,
	}
}
