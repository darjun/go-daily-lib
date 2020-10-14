package main

import (
	"fmt"

	"github.com/reactivex/rxgo/v2"
)

func main() {
	ch := make(chan rxgo.Item)
	go func() {
		for i := 1; i <= 5; i++ {
			ch <- rxgo.Of(i)
		}
		close(ch)
	}()

	observable := rxgo.FromChannel(ch)
	for item := range observable.Observe() {
		fmt.Println(item.V)
	}
}
