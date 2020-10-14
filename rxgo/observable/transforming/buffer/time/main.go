package main

import (
	"fmt"
	"time"

	"github.com/reactivex/rxgo/v2"
)

func main() {
	ch := make(chan rxgo.Item, 1)

	go func() {
		i := 0
		for range time.Tick(time.Second) {
			ch <- rxgo.Of(i)
			i++
		}
	}()

	observable := rxgo.FromChannel(ch).BufferWithTime(rxgo.WithDuration(3 * time.Second))

	for item := range observable.Observe() {
		fmt.Println(item.V)
	}
}
