package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/panjf2000/ants/v2"
)

func wrapper(i int, wg *sync.WaitGroup) func() {
	return func() {
		fmt.Printf("hello from task:%d\n", i)
		time.Sleep(1 * time.Second)
		wg.Done()
	}
}

func main() {
	defer ants.Release()

	var wg sync.WaitGroup
	wg.Add(2)
	for i := 1; i <= 2; i++ {
		ants.Submit(wrapper(i, &wg))
	}
	wg.Wait()
}
