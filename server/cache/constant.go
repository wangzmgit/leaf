package cache

/**
 * 缓存常量
 * author: hsx
 * date: 2022.11.10
 */

// 邮箱验证码缓存标识符
const EMAIL_CODE_KEY = "email_code_key:"

// 邮箱验证码过期时间 n 分钟
const EMIAL_CODE_EXPIRATION_TIME int = 5

// 图片验证码缓存标识符
const IMG_CODE_KEY = "img_code_key:"

// 图片验证码过期时间 n 分钟
const IMG_CODE_EXPIRATION_TIME = 3

// 用户信息缓存标识符
const USER_KEY = "user_key:"

// 用户信息过期时间 n 小时
const USER_EXPIRATION_TIME = 24

// 视频信息缓存标识符
const VIDEO_KEY = "video_key:"

// 视频信息过期时间 n 小时
const VIDEO_EXPIRATION_TIME = 24

// 最大登录数量
const MAX_LOGIN_LIMIT = 3

// 验证token缓存标识符
const ACCESS_TOKEN_KEY = "access_token_key:"

// 验证token过期时间 n 分钟
const ACCESS_TOKEN_EXPRIRATION_TIME = 60

// 刷新token缓存标识符
const REFRESH_TOKEN_KEY = "refresh_token_key:"

// 刷新token过期时间 n 小时
const REFRESH_TOKEN_EXPRIRATION_TIME = 336 // 14 * 24

// 登录尝试次数缓存标识符
const LOGIN_TRY_COUNT_KEY = "login_try_count_key:"

// 登录尝试次数过期时间 n 分钟
const LOGIN_TRY_COUNT_EXPRIRATION_TIME = 30

// 滑块x坐标缓存标识符
const SLIDER_X_KEY = "slider_x_key:"

// 滑块x坐标过期时间 n 分钟
const SLIDER_X_EXPRIRATION_TIME = 5

// 人机验证状态
const CAPTCHA_STATUS_KEY = "captcha_status_key:"

// 人机验证状态过期时间 n 分钟
const CAPTCHA_STATUS_EXPRIRATION_TIME = 30

// 上传文件缓存标识符
const UPLOAD_IMAGE_KEY = "upload_image_key:"

// 上传文件过期时间 n 分钟
const UPLOAD_IMAGE_EXPRIRATION_TIME = 20

// 分区缓存标识符
const PARTITION_KEY = "partition_key"

// 分区过期时间 不过期
const PARTITION_EXPRIRATION_TIME = 0

// 视频点击量限制标识符
const VIDEO_CLICKS_KEY = "video_clicks_key:"

// 视频点击量限制过期时间  n 分钟
const VIDEO_CLICKS_EXPRIRATION_TIME = 30

// 视频点击量标识符
const VIDEO_CLICKS_LIMIT_KEY = "video_clicks_limit_key:"

// 点击量过期时间  n 小时
const VIDEO_CLICKS_LIMIT_EXPRIRATION_TIME = 72

// 重置密码验证状态
const RESET_PWD_CHECK_KEY = "reset_pwd_check_key:"

// 重置密码验证状态过期时间 n 分钟
const RESET_PWD_CHECK_EXPRIRATION_TIME = 30
