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
	e.Cc = []string{"xxx@126.com"}
	e.Subject = "Go 每日一库"
	e.HTML = []byte(`
	<ul>
<li><a "https://darjun.github.io/2020/01/10/godailylib/flag/">Go 每日一库之 flag</a></li>
<li><a "https://darjun.github.io/2020/01/10/godailylib/go-flags/">Go 每日一库之 go-flags</a></li>
<li><a "https://darjun.github.io/2020/01/14/godailylib/go-homedir/">Go 每日一库之 go-homedir</a></li>
<li><a "https://darjun.github.io/2020/01/15/godailylib/go-ini/">Go 每日一库之 go-ini</a></li>
<li><a "https://darjun.github.io/2020/01/17/godailylib/cobra/">Go 每日一库之 cobra</a></li>
<li><a "https://darjun.github.io/2020/01/18/godailylib/viper/">Go 每日一库之 viper</a></li>
<li><a "https://darjun.github.io/2020/01/19/godailylib/fsnotify/">Go 每日一库之 fsnotify</a></li>
<li><a "https://darjun.github.io/2020/01/20/godailylib/cast/">Go 每日一库之 cast</a></li>
</ul>
	`)
	err := e.Send("smtp.126.com:25", smtp.PlainAuth("", "xxx@126.com", "yyy", "smtp.126.com"))
	if err != nil {
		log.Fatal("failed to send email:", err)
	}
}
