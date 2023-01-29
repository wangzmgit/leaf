package desensitization

import "regexp"

// 邮箱脱敏
func HideEmail(email string) (result string) {
	pattern := `(\w?)(\w+)(\w)(@\w+\.[a-z]+(\.[a-z]+)?)`
	reg := regexp.MustCompile(pattern)
	return reg.ReplaceAllString(email, "$1****$3$4")
}

// 手机号脱敏
func HidePhoneNumber(phone string) string {
	pattern := `(\d{3})(\d*)(\d{4})`
	reg := regexp.MustCompile(pattern)
	return reg.ReplaceAllString(phone, "$1****$3")
}
