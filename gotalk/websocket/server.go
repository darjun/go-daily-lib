package main

import (
	"log"
	"net/http"

	"github.com/rsms/gotalk"
)

func main() {
	gotalk.Handle("echo", func(in string) (string, error) {
		return in, nil
	})

	http.Handle("/gotalk/", gotalk.WebSocketHandler())
	http.Handle("/", http.FileServer(http.Dir(".")))
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
