package resp

import "net/http"

type R struct {
	httpStatus int
	code       int
	msg        string
}

var (
	Success = R{httpStatus: http.StatusOK, code: 200, msg: "OK"}
	Error   = R{httpStatus: http.StatusInternalServerError, code: 500, msg: "服务器异常"}

	// 10** 通用错误

	// 30** 认证授权相关错误
	EmailCodeError = R{httpStatus: http.StatusOK, code: 3010, msg: "邮箱验证码错误"}

	UsernamePasswordNotMatchError = R{httpStatus: http.StatusOK, code: 3020, msg: "用户名或密码错误"}

	// 40** 请求相关错误
	RequestParamError = R{httpStatus: http.StatusOK, code: 4010, msg: "请求参数有误"}

	UnauthorizedError = R{httpStatus: http.StatusOK, code: 4020, msg: "用户未授权"}

	// 50** 服务器相关错误

	// 60** 用户相关错误

	// 90** 第三方服务错误
	SendMailError = R{httpStatus: http.StatusOK, code: 9010, msg: "邮件发送失败"}
)
