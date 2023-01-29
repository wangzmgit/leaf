package common

//视频审核状态
const (
	// 审核通过
	AUDIT_APPROVED = 0

	// 成功创建视频
	CREATED_VIDEO = 100
	// 视频转码中
	VIDEO_PROCESSING = 200
	// 提交审核
	SUBMIT_REVIEW = 300
	// 等待审核
	WAITING_REVIEW = 500

	// 审核不通过
	// 视频信息存在问题
	WRONG_VIDEO_INFO = 2100
	// 视频内容存在问题
	WRONG_VIDEO_CONTENT = 2200
	// 视频处理失败
	PROCESSING_FAIL = 2300
)
