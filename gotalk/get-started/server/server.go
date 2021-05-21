package main

import (
	"log"

	"github.com/rsms/gotalk"
)

func main() {
	gotalk.Handle("echo", func(in string) (string, error) {
		return in, nil
	})
	if err := gotalk.Serve("tcp", ":8080", nil); err != nil {
		log.Fatal(err)
	}
}
