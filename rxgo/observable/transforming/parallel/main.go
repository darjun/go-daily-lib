package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"github.com/reactivex/rxgo/v2"
)

func main() {
	observable := rxgo.Range(1, 100)

	observable = observable.Map(func(_ context.Context, i interface{}) (interface{}, error) {
		time.Sleep(time.Duration(rand.Int31()))
		return i.(int)*2 + 1, nil
	}, rxgo.WithCPUPool())

	for item := range observable.Observe() {
		fmt.Println(item.V)
	}
}
