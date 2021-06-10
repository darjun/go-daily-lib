package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"

	"github.com/Jeffail/tunny"
)

const (
	DataSize    = 10000
	DataPerTask = 100
)

func main() {
	numCPUs := runtime.NumCPU()

	p := tunny.NewFunc(numCPUs, func(payload interface{}) interface{} {
		var sum int
		for _, n := range payload.([]int) {
			sum += n
		}
		return sum
	})
	defer p.Close()

	nums := make([]int, DataSize)
	for i := range nums {
		nums[i] = rand.Intn(1000)
	}

	var wg sync.WaitGroup
	wg.Add(DataSize / DataPerTask)
	partialSums := make([]int, DataSize/DataPerTask)
	for i := 0; i < DataSize/DataPerTask; i++ {
		go func(i int) {
			partialSums[i] = p.Process(nums[i*DataPerTask : (i+1)*DataPerTask]).(int)
			wg.Done()
		}(i)
	}

	wg.Wait()

	var sum int
	for _, s := range partialSums {
		sum += s
	}

	var expect int
	for _, num := range nums {
		expect += num
	}

	fmt.Printf("finish all tasks, result is %d expect:%d\n", sum, expect)
}
