package resp

import (
	"github.com/gin-gonic/gin"
)

const CODE_TAG, DATA_TAG, MSG_TAG = "code", "data", "msg"

func Response(ctx *gin.Context, r R, msg string, data gin.H) {
	// 没传入错误信息，用默认错误信息
	if msg == "" {
		msg = r.msg
	}
	ctx.JSON(r.httpStatus, gin.H{CODE_TAG: r.code, MSG_TAG: msg, DATA_TAG: data})
}

func OK(ctx *gin.Context, msg string, data gin.H) {
	Response(ctx, Success, msg, data)
}

func Fail(ctx *gin.Context, msg string, data gin.H) {
	Response(ctx, Error, msg, data)
}
