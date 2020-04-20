package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/roylee0704/gron"
	"github.com/roylee0704/gron/xtime"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)

	c := gron.New()
	c.AddFunc(gron.Every(1*time.Second), func() {
		fmt.Println("runs every second.")
	})
	c.AddFunc(gron.Every(1*time.Minute), func() {
		fmt.Println("runs every minute.")
	})
	c.AddFunc(gron.Every(1*time.Hour), func() {
		fmt.Println("runs every hour.")
	})
	c.AddFunc(gron.Every(1*xtime.Day), func() {
		fmt.Println("runs every day.")
	})
	c.AddFunc(gron.Every(1*xtime.Week), func() {
		fmt.Println("runs every week.")
	})
	t, _ := time.ParseDuration("4m10s")
	c.AddFunc(gron.Every(t), func() {
		fmt.Println("runs every 4 minutes 10 seconds.")
	})
	c.Start()

	wg.Wait()
}
