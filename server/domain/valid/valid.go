package valid

import (
	"regexp"
)

func Email(email string) bool {
	pattern := `\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*` //匹配电子邮箱
	reg := regexp.MustCompile(pattern)
	return reg.MatchString(email)
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
