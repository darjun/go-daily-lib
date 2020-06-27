package main

import (
	"fmt"
	"time"

	"github.com/robfig/cron/v3"
)

func main() {
	nyc, _ := time.LoadLocation("America/New_York")
	c := cron.New(cron.WithLocation(nyc))
	c.AddFunc("0 6 * * ?", func() {
		fmt.Println("Every 6 o'clock at New York")
	})

	c.AddFunc("CRON_TZ=Asia/Tokyo 0 6 * * ?", func() {
		fmt.Println("Every 6 o'clock at Tokyo")
	})

	c.Start()

	for {
		time.Sleep(time.Second)
	}
}
