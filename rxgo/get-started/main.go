package main

import (
	"fmt"

	"github.com/reactivex/rxgo/v2"
)

func main() {
	observable := rxgo.Just(1, 2, 3, 4, 5)()
	ch := observable.Observe()
	for item := range ch {
		fmt.Println(item.V)
	}
}
