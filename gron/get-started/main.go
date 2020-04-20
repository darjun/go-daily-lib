package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/roylee0704/gron"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)

	c := gron.New()
	c.AddFunc(gron.Every(5*time.Second), func() {
		fmt.Println("runs every 5 seconds.")
	})
	c.Start()

	wg.Wait()
}
