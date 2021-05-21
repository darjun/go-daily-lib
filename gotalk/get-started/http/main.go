package main

import (
	"fmt"
	"log"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello world")
}

func main() {
	http.HandleFunc("/", index)

	if err := http.ListenAndServe(":8888", nil); err != nil {
		log.Fatal(err)
	}
}
