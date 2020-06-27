package main

import (
	"fmt"
	"time"

	"github.com/robfig/cron/v3"
)

func main() {
	c := cron.New()

	c.AddFunc("30 * * * *", func() {
		fmt.Println("Every hour on the half hour")
	})

	c.AddFunc("30 3-6,20-23 * * *", func() {
		fmt.Println("On the half hour of 3-6am, 8-11pm")
	})

	c.AddFunc("0 0 1 1 *", func() {
		fmt.Println("Jun 1 every year")
	})

	c.Start()

	for {
		time.Sleep(time.Second)
	}
}
