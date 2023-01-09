package api

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"kuukaa.fun/leaf/domain/dto"
	"kuukaa.fun/leaf/domain/resp"
	"kuukaa.fun/leaf/domain/valid"
	"kuukaa.fun/leaf/domain/vo"
	"kuukaa.fun/leaf/service"
	"kuukaa.fun/leaf/util/convert"
	"kuukaa.fun/leaf/util/number"
)

// 发布评论
func Comment(ctx *gin.Context) {
	var commentDTO dto.CommentDTO
	if err := ctx.Bind(&commentDTO); err != nil {
		resp.Response(ctx, resp.RequestParamError, "", nil)
		zap.L().Error("请求参数有误")
		return
	}

	// 参数校验
	if !valid.CommentContent(commentDTO.Content) {
		resp.Response(ctx, resp.RequestParamError, valid.COMMENT_CONTENT_ERROR, nil)
		zap.L().Error(valid.COMMENT_CONTENT_ERROR)
		return
	}

	userId := ctx.GetUint("userId")

	// 处理@的用户
	atUserIds := service.SelectUserIdsByName(commentDTO.At)

	// 将DTO转为model并存入数据库
	comment := dto.CommentDtoToComment(commentDTO, userId, atUserIds)
	id, err := service.InsertComment(comment)
	if err != nil {
		resp.Response(ctx, resp.Error, "发布失败", nil)
		zap.L().Error("评论发布失败" + err.Error())
		return
	}

	// 获取评论的视频信息，并将回复通知写入数据库
	authorId := service.SelectVideoAuthorId(commentDTO.Vid)
	replyMsg := dto.CommentDtoToReplyMessage(commentDTO, id, authorId, userId)
	service.InsertReplyMessage(replyMsg)

	// 处理@通知并写入数据库
	atMsg := dto.UserIdsToAtMessage(atUserIds, commentDTO.Vid, userId)
	service.InsertManyAt(atMsg)

	// 返回给前端
	resp.OK(ctx, "ok", gin.H{"id": id})
}

// 发布回复
func Reply(ctx *gin.Context) {
	var commentDTO dto.ReplyDTO
	if err := ctx.Bind(&commentDTO); err != nil {
		resp.Response(ctx, resp.RequestParamError, "", nil)
		zap.L().Error("请求参数有误")
		return
	}

	// 参数校验
	if !valid.CommentContent(commentDTO.Content) {
		resp.Response(ctx, resp.RequestParamError, valid.COMMENT_CONTENT_ERROR, nil)
		zap.L().Error(valid.COMMENT_CONTENT_ERROR)
		return
	}

	userId := ctx.GetUint("userId")

	// 将DTO转为model
	atUserIds := service.SelectUserIdsByName(commentDTO.At)
	reply := dto.ReplyDtoToReply(commentDTO, userId, atUserIds)

	// 存入数据库
	id, err := service.InsertReply(commentDTO.ParentID, reply)
	if err != nil {
		resp.Response(ctx, resp.Error, "发布失败", nil)
		zap.L().Error("回复发布失败" + err.Error())
		return
	}

	// 获取回复的评论信息
	comment, _ := service.SelectCommentByID(commentDTO.ParentID)
	// 如果回复的用户ID不存在则通知评论作者
	notifyUserId := number.UintMax(commentDTO.ReplyUserID, comment.Uid)
	if notifyUserId != userId {
		replyMsg := dto.ReplyDtoToReplyMessage(commentDTO, notifyUserId, userId, comment.Content)
		// 回复通知
		service.InsertReplyMessage(replyMsg)
	}

	// 处理@通知并写入数据库
	atMsg := dto.UserIdsToAtMessage(atUserIds, commentDTO.Vid, userId)
	service.InsertManyAt(atMsg)

	// 返回给前端
	resp.OK(ctx, "ok", gin.H{"id": id})
}

// 获取评论
func GetComment(ctx *gin.Context) {
	vid := convert.StringToUint(ctx.DefaultQuery("vid", "0"))
	page := convert.StringToInt(ctx.DefaultQuery("page", "1"))
	pageSize := convert.StringToInt(ctx.DefaultQuery("page_size", "10"))

	if pageSize > 30 {
		resp.Response(ctx, resp.TooManyRequestsError, "", nil)
		zap.L().Error("请求数量过多")
		return
	}

	comment, err := service.SelectCommentList(vid, page, pageSize)
	if err != nil {
		resp.Response(ctx, resp.Error, "", nil)
		zap.L().Error("查询评论失败" + err.Error())
		return
	}

	// 添加用户信息
	for i := 0; i < len(comment); i++ {
		comment[i].Author = service.GetUserInfo(comment[i].Uid)
		for j := 0; j < len(comment[i].Reply); j++ {
			comment[i].Reply[j].Author = service.GetUserInfo(comment[i].Reply[j].Uid)
		}
	}

	// 返回给前端
	resp.OK(ctx, "ok", gin.H{"comments": vo.ToCommentVO(comment)})
}

// 获取回复
func GetReply(ctx *gin.Context) {
	commentId := ctx.Query("cid")
	page := convert.StringToInt(ctx.DefaultQuery("page", "1"))
	pageSize := convert.StringToInt(ctx.DefaultQuery("page_size", "10"))

	if pageSize > 30 {
		resp.Response(ctx, resp.TooManyRequestsError, "", nil)
		zap.L().Error("请求数量过多 ")
		return
	}

	reply, err := service.SelectReplyList(commentId, page, pageSize)
	if err != nil {
		resp.Response(ctx, resp.Error, "", nil)
		zap.L().Error("查询回复失败" + err.Error())
		return
	}

	// 添加用户信息
	for i := 0; i < len(reply); i++ {
		reply[i].Author = service.GetUserInfo(reply[i].Uid)
	}

	// 返回给前端
	resp.OK(ctx, "ok", gin.H{"replies": vo.ToReplyVO(reply)})
}

// 删除评论
func DeleteComment(ctx *gin.Context) {
	var idDTO dto.ObjectIdDTO
	if err := ctx.Bind(&idDTO); err != nil {
		resp.Response(ctx, resp.RequestParamError, "", nil)
		zap.L().Error("请求参数有误")
		return
	}

	userId := ctx.GetUint("userId")
	if comment, _ := service.SelectCommentByID(idDTO.ID); comment.Uid != userId {
		resp.Response(ctx, resp.CommentNotExistError, "", nil)
		zap.L().Error("评论不存在")
		return
	}

	service.DeleteComment(idDTO.ID)

	// 返回给前端
	resp.OK(ctx, "ok", nil)
}

// 删除回复
func DeleteReply(ctx *gin.Context) {
	var deleteReplyDTO dto.DeleteReplyDTO
	if err := ctx.Bind(&deleteReplyDTO); err != nil {
		resp.Response(ctx, resp.RequestParamError, "", nil)
		zap.L().Error("请求参数有误")
		return
	}

	userId := ctx.GetUint("userId")
	reply, _ := service.SelectReplyByID(deleteReplyDTO.CommentID, deleteReplyDTO.ReplyID)
	if reply.Uid != userId {
		resp.Response(ctx, resp.CommentNotExistError, "", nil)
		zap.L().Error("回复不存在")
		return
	}

	service.DeleteReply(deleteReplyDTO.CommentID, deleteReplyDTO.ReplyID)

	// 返回给前端
	resp.OK(ctx, "ok", nil)
}
