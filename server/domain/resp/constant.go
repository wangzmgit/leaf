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
	Captcha = R{httpStatus: http.StatusOK, code: -1, msg: "需要人机验证"}

	// 10** 通用错误
	CreateError             = R{httpStatus: http.StatusOK, code: 1000, msg: "创建失败"}
	SelectError             = R{httpStatus: http.StatusOK, code: 1000, msg: "查询失败"}
	UpdateError             = R{httpStatus: http.StatusOK, code: 1000, msg: "更新失败"}
	DeleteError             = R{httpStatus: http.StatusOK, code: 1000, msg: "删除失败"}
	SliderVerificationError = R{httpStatus: http.StatusOK, code: 1010, msg: "滑块验证失败"}
	InvalidLinkError        = R{httpStatus: http.StatusOK, code: 1020, msg: "链接已经失效"}

	// 30** 认证授权相关错误
	TokenExpriedError = R{httpStatus: http.StatusOK, code: 3000, msg: "token已过期"}

	EmailCodeError = R{httpStatus: http.StatusOK, code: 3010, msg: "邮箱验证码错误"}

	UsernamePasswordNotMatchError = R{httpStatus: http.StatusOK, code: 3020, msg: "用户名或密码错误"}

	UnauthorizedError = R{httpStatus: http.StatusOK, code: 3030, msg: "用户未授权"}

	// 40** 请求相关错误
	RequestParamError = R{httpStatus: http.StatusOK, code: 4010, msg: "请求参数有误"}

	FileCheckError = R{httpStatus: http.StatusOK, code: 4020, msg: "文件不符合要求"}

	FileUploadError = R{httpStatus: http.StatusOK, code: 4030, msg: "文件上传失败"}

	PartitionError          = R{httpStatus: http.StatusOK, code: 4040, msg: "分区不存在"}
	ParentPartitionError    = R{httpStatus: http.StatusOK, code: 4040, msg: "所属分区不存在"}
	UserNotExistError       = R{httpStatus: http.StatusOK, code: 4040, msg: "用户不存在"}
	VideoNotExistError      = R{httpStatus: http.StatusOK, code: 4040, msg: "视频不存在"}
	ResourceNotExistError   = R{httpStatus: http.StatusOK, code: 4040, msg: "资源不存在"}
	CollectionNotExistError = R{httpStatus: http.StatusOK, code: 4040, msg: "收藏夹不存在"}
	CommentNotExistError    = R{httpStatus: http.StatusOK, code: 4040, msg: "评论或回复不存在"}

	TooManyRequestsError = R{httpStatus: http.StatusOK, code: 4050, msg: "请求数量过多"}

	FollowYourselfError = R{httpStatus: http.StatusOK, code: 4060, msg: "不能关注自己"}

	// 50** 服务器相关错误

	// 60** 用户相关错误
	NameExistError  = R{httpStatus: http.StatusOK, code: 6000, msg: "用户名已存在"}
	EmailExistError = R{httpStatus: http.StatusOK, code: 6000, msg: "邮箱已存在"}

	// 90** 第三方服务错误
	SendMailError = R{httpStatus: http.StatusOK, code: 9010, msg: "邮件发送失败"}
	OssError      = R{httpStatus: http.StatusOK, code: 9020, msg: "文件存储错误"}
)
