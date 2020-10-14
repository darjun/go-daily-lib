package main

import (
	"context"
	"fmt"

	"github.com/reactivex/rxgo/v2"
)

func main() {
	observable := rxgo.Defer([]rxgo.Producer{func(_ context.Context, ch chan<- rxgo.Item) {
		for i := 0; i < 3; i++ {
			ch <- rxgo.Of(i)
		}
	}})

	for item := range observable.Observe() {
		fmt.Println(item.V)
	}

	for item := range observable.Observe() {
		fmt.Println(item.V)
	}
}
