package main

import (
	"fmt"
	"sync"

	messagebus "github.com/vardius/message-bus"
)

func main() {
	queueSize := 100
	bus := messagebus.New(queueSize)

	var wg sync.WaitGroup
	wg.Add(2)

	_ = bus.Subscribe("topic", func(v bool) {
		defer wg.Done()
		fmt.Println(v)
	})

	_ = bus.Subscribe("topic", func(v bool) {
		defer wg.Done()
		fmt.Println(v)
	})

	bus.Publish("topic", true)
	wg.Wait()
}
