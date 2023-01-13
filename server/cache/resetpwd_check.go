package cache

import (
	"strconv"
	"time"

	"go.uber.org/zap"
)

// 状态1: 验证成功

func GetResetPwdCheckStatus(email string) int {
	s := Get(RESET_PWD_CHECK_KEY + email)
	status, err := strconv.Atoi(s)
	if err != nil {
		zap.L().Error("重置密码状态转换为int类型失败")
	}
	return status

}

func SetResetPwdCheckStatus(email string, status int) {
	Set(RESET_PWD_CHECK_KEY+email, status, time.Minute*RESET_PWD_CHECK_EXPRIRATION_TIME)
}

func DelResetPwdCheckStatus(email string) {
	Del(RESET_PWD_CHECK_KEY + email)
}
