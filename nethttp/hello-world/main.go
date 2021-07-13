package main

import (
	"fmt"
	"net/http"
)

type greeting string

func (g greeting) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, g)
}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello World")
}

func main() {
	http.HandleFunc("/", index)
	http.Handle("/greeting", greeting("Welcome, dj"))
	http.ListenAndServe(":8080", nil)
}
