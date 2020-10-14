package main

import (
	"fmt"

	"github.com/reactivex/rxgo/v2"
)

func main() {
	count := 3
	observable := rxgo.Range(0, 10).GroupBy(count, func(item rxgo.Item) int {
		return item.V.(int) % count
	}, rxgo.WithBufferedChannel(10))

	for subObservable := range observable.Observe() {
		fmt.Println("New observable:")

		for item := range subObservable.V.(rxgo.Observable).Observe() {
			fmt.Printf("item: %v\n", item.V)
		}
	}
}
