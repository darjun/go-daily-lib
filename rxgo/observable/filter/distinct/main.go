package main

import (
	"context"
	"fmt"

	"github.com/reactivex/rxgo/v2"
)

func main() {
	observable := rxgo.Just(1, 2, 2, 3, 3, 4, 4)().
		Distinct(func(_ context.Context, i interface{}) (interface{}, error) {
			return i, nil
		})
	for item := range observable.Observe() {
		fmt.Println(item.V)
	}
}
