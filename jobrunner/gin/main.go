package main

import (
	"fmt"
	"net/smtp"
	"time"

	"github.com/bamzi/jobrunner"
	"github.com/gin-gonic/gin"
	"github.com/jordan-wright/email"
)

type EmailJob struct {
	Name  string
	Email string
}

type User struct {
	Name  string `form:"name"`
	Email string `form:"email"`
}

func (j EmailJob) Run() {
	e := email.NewEmail()
	e.From = "leedarjun@126.com"
	e.To = []string{j.Email}
	e.Cc = []string{"leedarjun@126.com"}
	e.Subject = "Welcome To Awesome-Web"
	e.Text = []byte(fmt.Sprintf(`
	Hello, %s
	Welcome Back
	`, j.Name))

	err := e.Send("smtp.126.com:25", smtp.PlainAuth("", "leedarjun@126.com", "LFODPEEAMKAIIKMB", "smtp.126.com"))
	if err != nil {
		fmt.Printf("failed to send email to %s, err:%v", j.Name, err)
	}
}

func login(c *gin.Context) {
	var u User
	if c.ShouldBind(&u) == nil {
		c.String(200, "login success")

		jobrunner.In(5*time.Second, EmailJob{Name: u.Name, Email: u.Email})
	} else {
		c.String(404, "login failed")
	}
}

func main() {
	r := gin.Default()
	r.GET("/login", login)
	r.Run(":8888")
}
