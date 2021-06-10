package main

import (
	"context"
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"

	"github.com/Jeffail/tunny"
)

func fib(n int) int {
	if n <= 1 {
		return 1
	}

	return fib(n-1) + fib(n-2)
}

func main() {
	numCPUs := runtime.NumCPU()
	fmt.Printf("cpu num:%d\n", numCPUs)

	p := tunny.NewFunc(numCPUs, func(payload interface{}) interface{} {
		n := payload.(int)
		result := fib(n)
		time.Sleep(1 * time.Second)
		return result
	})
	defer p.Close()

	var wg sync.WaitGroup
	wg.Add(numCPUs)
	for i := 0; i < numCPUs; i++ {
		go func(i int) {
			n := rand.Intn(30)
			ctx, cancel := context.WithCancel(context.Background())
			if i%2 == 0 {
				go func() {
					time.Sleep(500 * time.Millisecond)
					cancel()
				}()
			}

			result, err := p.ProcessCtx(ctx, n)
			if err != nil {
				fmt.Printf("task(%d) failed:%v\n", i, err)
			} else {
				fmt.Printf("fib(%d) = %d\n", n, result)
			}
			wg.Done()
		}(i)
	}

	wg.Wait()
}
