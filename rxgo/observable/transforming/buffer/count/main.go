package main

import (
	"fmt"

	"github.com/reactivex/rxgo/v2"
)

func main() {
	observable := rxgo.Just(1, 2, 3, 4)()

	observable = observable.BufferWithCount(3)

	for item := range observable.Observe() {
		fmt.Println(item.V)
	}
}
