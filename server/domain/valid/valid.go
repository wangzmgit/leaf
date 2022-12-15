package valid

import (
	"regexp"
	"strconv"

	"go.uber.org/zap"
)

func Email(email string) bool {
	pattern := `\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*` //匹配电子邮箱
	reg := regexp.MustCompile(pattern)
	return reg.MatchString(email)
}

// 从query中获取的id，可能为负数
func ID(id int) bool {
	return id > 0
}

func Title(title string) bool {
	return len(title) > 0
}

func Content(content string) bool {
	return len(content) > 0
}

func Name(name string) bool {
	return len(name) > 0
}

func Password(password string) bool {
	return len(password) >= 6
}

func EmailCode(code string) bool {
	return len(code) == 4
}

func ImgCode(code string) bool {
	return len(code) == 4
}

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
