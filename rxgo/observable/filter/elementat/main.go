package main

import (
	"fmt"

	"github.com/reactivex/rxgo/v2"
)

func main() {
	observable := rxgo.Just(0, 1, 2, 3, 4)().ElementAt(2)

	for item := range observable.Observe() {
		fmt.Println(item.V)
	}
}
