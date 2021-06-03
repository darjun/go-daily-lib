package main

import (
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/panjf2000/ants/v2"
)

func wrapper(i int, wg *sync.WaitGroup) func() {
	return func() {
		fmt.Printf("hello from task:%d\n", i)
		if i%2 == 0 {
			panic(fmt.Sprintf("panic from task:%d", i))
		}
		wg.Done()
	}
}

func panicHandler(err interface{}) {
	fmt.Fprintln(os.Stderr, err)
}

func main() {
	p, _ := ants.NewPool(2, ants.WithPanicHandler(panicHandler))
	defer p.Release()

	var wg sync.WaitGroup
	wg.Add(3)
	for i := 1; i <= 2; i++ {
		p.Submit(wrapper(i, &wg))
	}

	time.Sleep(1 * time.Second)
	p.Submit(wrapper(3, &wg))
	p.Submit(wrapper(5, &wg))
	wg.Wait()
}
