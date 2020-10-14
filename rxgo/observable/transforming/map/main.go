package main

import (
	"context"
	"fmt"

	"github.com/reactivex/rxgo/v2"
)

func main() {
	observable := rxgo.Just(1, 2, 3)()

	observable = observable.Map(func(_ context.Context, i interface{}) (interface{}, error) {
		return i.(int)*2 + 1, nil
	}).Map(func(_ context.Context, i interface{}) (interface{}, error) {
		return i.(int)*3 + 2, nil
	})

	for item := range observable.Observe() {
		fmt.Println(item.V)
	}
}
