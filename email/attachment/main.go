package main

import (
	"log"
	"net/smtp"

	"github.com/jordan-wright/email"
)

func main() {
	e := email.NewEmail()
	e.From = "dj <xxx@126.com>"
	e.To = []string{"935653229@qq.com"}
	e.Subject = "Go 每日一库"
	e.Text = []byte("请看附件")
	e.AttachFile("test.txt")
	err := e.Send("smtp.126.com:25", smtp.PlainAuth("", "xxx@126.com", "yyy", "smtp.126.com"))
	if err != nil {
		log.Fatal("failed to send email:", err)
	}
}
