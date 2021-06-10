package main

import (
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
		time.Sleep(990 * time.Millisecond)
		return result
	})
	defer p.Close()

	var wg sync.WaitGroup
	wg.Add(2 * numCPUs)
	for i := 0; i < 2*numCPUs; i++ {
		go func(i int) {
			n := rand.Intn(30)
			result, err := p.ProcessTimed(n, time.Second)
			nowStr := time.Now().Format("2006-01-02 15:04:05")
			if err != nil {
				fmt.Printf("[%s]task(%d) failed:%v\n", nowStr, i, err)
			} else {
				fmt.Printf("[%s]fib(%d) = %d\n", nowStr, n, result)
			}
			wg.Done()
		}(i)
	}

	wg.Wait()
}
