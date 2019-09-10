package service

import (
	"fmt"
	"strconv"

	"gopkg.in/gomail.v2"
)

//EmailHelp 邮箱发送助手
func EmailHelp() *gomail.Dialer {
	host := Config("MAIL_HOST", "localhost")
	p := Config("MAIL_PORT", "2525")
	user := Config("MAIL_USERNAME", "root")
	pass := Config("MAIL_PASSWORD", "root")

	port, _ := strconv.Atoi(p)
	return gomail.NewDialer(host, port, user, pass)
}

//RegisterTemplate 注册模板
func RegisterTemplate(email, name string) {
	help := EmailHelp()
	m := gomail.NewMessage()
	user := Config("MAIL_USERNAME", "root")
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
