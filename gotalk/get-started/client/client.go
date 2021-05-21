package main

import (
	"fmt"
	"log"

	"github.com/rsms/gotalk"
)

func main() {
	s, err := gotalk.Connect("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}

	for i := 0; i < 10; i++ {
		var echo string
		if err := s.Request("echo", "hello", &echo); err != nil {
			log.Fatal(err)
		}

		fmt.Println(echo)
	}

	s.Close()
}
