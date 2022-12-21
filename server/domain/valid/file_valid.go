package valid

import (
	"regexp"
	"strconv"

	"go.uber.org/zap"
)

func FileType(suffix string, isImg bool) bool {
	pattern := `.(png|jpeg|jpg)`
	if !isImg {
		pattern = `.(mp4)`
	}
	reg := regexp.MustCompile(pattern)
	return reg.MatchString(suffix)
}

func FileSize(contentLength string, targetSize int64) bool {
	contentSize, err := strconv.ParseInt(contentLength, 10, 64)
	if err != nil {
		zap.L().Error("请求大小转换失败: " + err.Error())
		return false
	}
	//'限制整体大小为 目标大小 + 1 MB
	if contentSize > (targetSize+1)*1024*1024 {
		return false
	}
	return true
}
