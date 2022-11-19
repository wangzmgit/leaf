package cache

import (
	"strconv"
	"time"

	"go.uber.org/zap"
)

func GetSliderX(email string) int {
	s := Get(SLIDER_X_KEY + email)
	x, err := strconv.Atoi(s)
	if err != nil {
		zap.L().Error("缓存中滑块x坐标转换为int类型失败")
	}
	return x
}

func SetSliderX(email string, x int) {
	Set(SLIDER_X_KEY+email, x, time.Minute*SLIDER_X_EXPRIRATION_TIME)
}

func DelSlider(email string) {
	Del(SLIDER_X_KEY + email)
}
