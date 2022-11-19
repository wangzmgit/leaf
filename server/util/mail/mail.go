package mail

import (
	"github.com/spf13/viper"
	"gopkg.in/gomail.v2"
)

/**
 * 发送邮箱验证码
 * param: email 目标邮箱
 * param: code 邮箱验证码
 * return: 发送失败时的错误信息
 */
func SendCaptcha(email string, code string) error {
	// 定义收件人
	mailTo := []string{email}
	// 邮件主题
	subject := "leaf弹幕网站"
	// 邮件正文
	body := "<h3>尊敬的用户：</h3><p>您好! 您的验证码是 <span style='color:red'> " + code + "</span>，五分钟内有效，祝您生活愉快！</p>"
	return Send(mailTo, subject, body)
}

/**
 * 发送电子邮件
 * param: emailList 目标邮箱数组
 * param: subject 邮件主题
 * param: body 邮件内容
 * return: 发送失败时的错误信息
 */
func Send(emailList []string, subject string, body string) error {
	user, pass, host, port, addresser :=
		viper.GetString("mail.user"),
		viper.GetString("mail.pass"),
		viper.GetString("mail.host"),
		viper.GetInt("mail.port"),
		viper.GetString("mail.addresser")

	m := gomail.NewMessage()
	m.SetHeader("From", addresser+"<"+user+">") //添加别名
	m.SetHeader("To", emailList...)             //发送给多个用户
	m.SetHeader("Subject", subject)             //设置邮件主题
	m.SetBody("text/html", body)                //设置邮件正文

	d := gomail.NewDialer(host, port, user, pass)
	err := d.DialAndSend(m)
	return err
}
