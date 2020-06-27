package main

import (
	"fmt"
	"time"

	"github.com/robfig/cron/v3"
)

func main() {
	parser := cron.NewParser(
		cron.Second | cron.Minute | cron.Hour | cron.Dom | cron.Month | cron.Dow | cron.Descriptor,
	)
	c := cron.New(cron.WithParser(parser))
	c.AddFunc("1 * * * * *", func () {
		fmt.Println("every 1 second")
	})
	c.Start()

	time.Sleep(5 * time.Second)
}
