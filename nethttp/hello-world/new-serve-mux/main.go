package main

import (
	"fmt"
	"net/http"
	"time"
)

type greeting string

func (g greeting) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, g)
}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello World")
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", index)
	mux.Handle("/greeting", greeting("Welcome, dj"))

	server := &http.Server{
		Addr:         ":8080",
		Handler:      mux,
		ReadTimeout:  20 * time.Second,
		WriteTimeout: 20 * time.Second,
	}
	server.ListenAndServe()
}
