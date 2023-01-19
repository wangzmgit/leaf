package valid

const (
	EMAIL_ERROR      = "邮箱格式错误"
	PASSWORD_ERROR   = "密码长度至少6位"
	EMAIL_CODE_ERROR = "邮箱验证码错误"
	NAME_ERROR       = "用户名不能为空"

	ID_ERROR = "ID错误"

	TITLE_ERROR   = "标题不能为空"
	CONTENT_ERROR = "内容不能为空"

	COLLECTION_NAME_ERROR = "收藏夹名不能为空"

	FILE_TYPE_ERROR = "文件类型不符合要求"
	FILE_SIZE_ERROR = "文件大小不符合要求"

	// 用户
	ROLE_ERROR = "无效的角色"

	// 视频
	REVIEW_STATUS_ERROR = "无效的视频状态"

	// 评论校验
	COMMENT_CONTENT_ERROR = "评论或回复内容不能为空"

	// 公告校验
	ANNOUNCE_TITLE_ERROR   = "公告标题不符合长度要求"
	ANNOUNCE_CONTENT_ERROR = "公告内容不符合长度要求"
	ANNOUNCE_URL_ERROR     = "公告链接不符合长度要求"

	// 消息校验
	MESSAGE_SEND_ERROR    = "消息发送失败"
	MESSAGE_CONTENT_ERROR = "消息内容不能为空"
	SEND_YOURSELF_ERROR   = "不能发送给自己"

	// 弹幕
	DANMAKU_TEXT_ERROR = "弹幕内容不能为空"
)
