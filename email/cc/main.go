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
	e.Cc = []string{"test1@126.com", "test2@126.com"}
	e.Bcc = []string{"secret@126.com"}
	e.Subject = "Awesome web"
	e.Text = []byte("Text Body is, of course, supported!")
	err := e.Send("smtp.126.com:25", smtp.PlainAuth("", "xxx@126.com", "yyy", "smtp.126.com"))
	if err != nil {
		log.Fatal(err)
	}
}
