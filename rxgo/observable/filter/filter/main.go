package main

import (
	"fmt"

	"github.com/reactivex/rxgo/v2"
)

func main() {
	observable := rxgo.Range(1, 10)

	observable = observable.Filter(func(i interface{}) bool {
		return i.(int)%2 == 0
	})

	for item := range observable.Observe() {
		fmt.Println(item.V)
	}
}
