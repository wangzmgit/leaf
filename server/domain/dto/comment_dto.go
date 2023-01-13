package dto

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"kuukaa.fun/leaf/domain/model"
)

type CommentDTO struct {
	Vid     uint
	Content string
	At      []string
}

type ReplyDTO struct {
	Vid          uint
	Content      string
	ParentID     primitive.ObjectID
	At           []string
	ReplyUserID  uint
	ReplyContent string
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
func CommentDtoToComment(commentDTO CommentDTO, userId uint, atUsers []uint) model.Comment {
	return model.Comment{
		ID:        primitive.NewObjectID(),
		Vid:       commentDTO.Vid,
		CreatedAt: time.Now().UnixMilli(),
		Content:   commentDTO.Content,
		Uid:       userId,
		Reply:     []model.Reply{},
		At:        atUsers,
		IsDelete:  false,
	}
}

/**
 * 回复DTO结构体转化为Reply结构体
 * param: replyDTO 回复DTO结构体
 * return: reply结构体
 */
func ReplyDtoToReply(replyDTO ReplyDTO, userId uint, atUsers []uint) model.Reply {
	if replyDTO.ReplyUserID != 0 {
		atUsers = append(atUsers, replyDTO.ReplyUserID)
	}
	return model.Reply{
		ID:        primitive.NewObjectID(),
		CreatedAt: time.Now().UnixMilli(),
		Content:   replyDTO.Content,
		Uid:       userId,
		At:        atUsers,
		IsDelete:  false,
	}
}

/**
 * 评论DTO结构体转化为ReplyMessage结构体
 * param: commentDTO 评论DTO结构体
 * param: commentId 评论DTO结构体
 * param: userId 接收者ID
 * param: fid 发送者ID
 * return: replyMessage结构体
 */
func CommentDtoToReplyMessage(commentDTO CommentDTO, commentId primitive.ObjectID, userId, fid uint) model.ReplyMessage {
	return model.ReplyMessage{
		Vid:       commentDTO.Vid,
		Uid:       userId,
		Fid:       fid,
		Content:   commentDTO.Content,
		CommentId: commentId.String(),
	}
}

/**
 * 回复DTO结构体转化为ReplyMessage结构体
 * param: replyDTO 回复DTO结构体
 * param: userId 接收者ID
 * param: fid 发送者ID
 * param: rootContent 根评论内容
 * return: replyMessage结构体
 */
func ReplyDtoToReplyMessage(replyDTO ReplyDTO, userId, fid uint, rootContent string) model.ReplyMessage {
	return model.ReplyMessage{
		Vid:                replyDTO.Vid,
		Uid:                userId,
		Fid:                fid,
		Content:            replyDTO.Content,
		RootContent:        rootContent,
		TargetReplyContent: replyDTO.ReplyContent,
		CommentId:          replyDTO.ParentID.String(),
	}
}
