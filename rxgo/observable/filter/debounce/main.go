package main

import (
	"fmt"
	"time"

	"github.com/reactivex/rxgo/v2"
)

func main() {
	ch := make(chan rxgo.Item)

	go func() {
		ch <- rxgo.Of(1)
		time.Sleep(2 * time.Second)
		ch <- rxgo.Of(2)
		ch <- rxgo.Of(3)
		time.Sleep(2 * time.Second)
		close(ch)
	}()

	observable := rxgo.FromChannel(ch).Debounce(rxgo.WithDuration(1 * time.Second))
	for item := range observable.Observe() {
		fmt.Println(item.V)
	}
}
