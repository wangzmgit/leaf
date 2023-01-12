package cache

import (
	"strconv"
	"time"

	"go.uber.org/zap"
)

// 0 未验证  1验证成功

func GetCaptchaStatus(email string) int {
	s := Get(CAPTCHA_STATUS_KEY + email)
	status, err := strconv.Atoi(s)
	if err != nil {
		zap.L().Error("人机验证状态转换为int类型失败")
	}
	return status

}

func SetCaptchaStatus(email string, status int) {
	Set(CAPTCHA_STATUS_KEY+email, status, time.Minute*CAPTCHA_STATUS_EXPRIRATION_TIME)
}

func DelCaptchaStatus(email string) {
	Del(CAPTCHA_STATUS_KEY + email)
}
