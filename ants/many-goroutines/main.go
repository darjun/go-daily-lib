package main

import (
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(10000000)
	for i := 0; i < 10000000; i++ {
		go func() {
			time.Sleep(1 * time.Minute)
		}()
	}
	wg.Wait()
}
