package main

import (
	"fmt"
	"sync"

	"github.com/roylee0704/gron"
	"github.com/roylee0704/gron/xtime"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)

	c := gron.New()
	c.AddFunc(gron.Every(1*xtime.Day).At("22:00"), func() {
		fmt.Println("runs every second.")
	})
	c.Start()

	wg.Wait()
}
