package service

import (
	"fmt"
	"strconv"

	"github.com/mojiajuzi/forum/config"

	"gopkg.in/gomail.v2"
)

//EmailHelp 邮箱发送助手
func EmailHelp() *gomail.Dialer {
	host := config.Config("MAIL_HOST", "localhost")
	p := config.Config("MAIL_PORT", "2525")
	user := config.Config("MAIL_USERNAME", "root")
	pass := config.Config("MAIL_PASSWORD", "root")

	port, _ := strconv.Atoi(p)
	return gomail.NewDialer(host, port, user, pass)
}

//RegisterTemplate 注册模板
func RegisterTemplate(email, name string) {
	help := EmailHelp()
	m := gomail.NewMessage()
	user := config.Config("MAIL_USERNAME", "root")
	m.SetHeader("From", user)
	m.SetHeader("To", email)
	m.SetHeader("Subject", "welcome register")
	content := fmt.Sprintf("Hello <b>%s</b>, welcome to register this forum!", name)
	m.SetBody("text/html", content)
	if err := help.DialAndSend(m); err != nil {
		//TODO 添加错误日志
		fmt.Println(err)
		return
	}
	fmt.Println("success")
}
