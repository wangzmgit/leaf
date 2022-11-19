package cache

import "time"

func GetEmailCode(email string) string {
	return Get(EMAIL_CODE_KEY + email)
}
func SetEmailCode(email string, code string) {
	Set(EMAIL_CODE_KEY+email, code, time.Minute*time.Duration(EMIAL_CODE_EXPIRATION_TIME))
}
func DelEmailCode(email string) {
	Del(EMAIL_CODE_KEY + email)
}
