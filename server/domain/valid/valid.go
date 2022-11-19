package valid

import "regexp"

func Email(email string) bool {
	pattern := `\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*` //匹配电子邮箱
	reg := regexp.MustCompile(pattern)
	return reg.MatchString(email)
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
