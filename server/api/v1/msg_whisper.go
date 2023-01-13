package api

import (
	"encoding/json"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"kuukaa.fun/leaf/domain/dto"
	"kuukaa.fun/leaf/domain/model"
	"kuukaa.fun/leaf/domain/resp"
	"kuukaa.fun/leaf/domain/valid"
	"kuukaa.fun/leaf/domain/vo"
	"kuukaa.fun/leaf/service"
	"kuukaa.fun/leaf/util/convert"
	"kuukaa.fun/leaf/ws"
)

// 获取消息Websocket连接
func GetWhisperConnect(ctx *gin.Context) {
	userId := ctx.GetUint("userId")
	// 升级为websocket长链接
	ws.MsgWsHandler(ctx.Writer, ctx.Request, userId)
}

// 发送私信
func SendWhisper(ctx *gin.Context) {
	var messageDTO dto.WhisperDTO
	if err := ctx.Bind(&messageDTO); err != nil {
		resp.Response(ctx, resp.RequestParamError, "", nil)
		zap.L().Error("请求参数有误")
		return
	}

	userId := ctx.GetUint("userId")

	//验证数据
	if messageDTO.Fid == 0 {
		resp.Response(ctx, resp.RequestParamError, valid.MESSAGE_SEND_ERROR, nil)
		zap.L().Error(valid.MESSAGE_SEND_ERROR)
		return
	}

	if messageDTO.Fid == userId {
		resp.Response(ctx, resp.RequestParamError, valid.SEND_YOURSELF_ERROR, nil)
		zap.L().Error(valid.SEND_YOURSELF_ERROR)
		return
	}

	if !valid.MessageContent(messageDTO.Content) {
		resp.Response(ctx, resp.RequestParamError, valid.MESSAGE_CONTENT_ERROR, nil)
		zap.L().Error(valid.MESSAGE_CONTENT_ERROR)
		return
	}

	// 存入数据库
	messages := dto.WhisperDtoToWhisper(messageDTO, userId)
	service.InsertManyWhisper(messages)

	//推送消息给接收者
	data, _ := json.Marshal(&vo.WhisperVO{
		Fid:     userId,
		Content: messageDTO.Content,
	})
	ws.SendMsg(messageDTO.Fid, data)

	// 返回给前端
	resp.OK(ctx, "ok", nil)
}

// 获取消息列表
func GetWhisperList(ctx *gin.Context) {
	userId := ctx.GetUint("userId")
	messages := service.SelectWhisperGroup(userId)

	users := make([]model.User, len(messages))
	for i := 0; i < len(messages); i++ {
		users[i] = service.GetUserInfo(messages[i].Fid)
	}

	// 返回给前端
	resp.OK(ctx, "ok", gin.H{"messages": vo.ToWhisperGroupVoList(messages, users)})
}

// 获取消息详细信息
func GetMessageDetails(ctx *gin.Context) {
	userId := ctx.GetUint("userId")
	fid := convert.StringToUint(ctx.Query("fid"))
	page := convert.StringToInt(ctx.Query("page"))
	pageSize := convert.StringToInt(ctx.Query("page_size"))

	if pageSize > 30 {
		resp.Response(ctx, resp.TooManyRequestsError, "", nil)
		zap.L().Error("请求数量过多 ")
		return
	}

	messages := service.SelectWhisper(userId, fid, page, pageSize)

	// 此时查询到的消息为为倒叙，需要进行反转
	for i, j := 0, len(messages)-1; i < j; i, j = i+1, j-1 {
		messages[i], messages[j] = messages[j], messages[i]
	}

	// 返回给前端
	resp.OK(ctx, "ok", gin.H{"messages": vo.ToMsgDetailsVoList(messages)})
}

// 已读消息
func ReadWhisper(ctx *gin.Context) {
	var idDTO dto.IdDTO
	if err := ctx.Bind(&idDTO); err != nil {
		resp.Response(ctx, resp.RequestParamError, "", nil)
		zap.L().Error("请求参数有误")
		return
	}

	userId := ctx.GetUint("userId")
	service.UpdateWhisperStatus(userId, idDTO.ID)

	// 返回给前端
	resp.OK(ctx, "ok", nil)
}
