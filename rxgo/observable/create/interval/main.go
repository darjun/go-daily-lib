package main

import (
	"fmt"
	"time"

	"github.com/reactivex/rxgo/v2"
)

func main() {
	observable := rxgo.Interval(rxgo.WithDuration(5 * time.Second))
	for item := range observable.Observe() {
		fmt.Println(item.V)
	}
}
