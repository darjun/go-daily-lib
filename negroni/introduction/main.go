package main

import (
	"fmt"
	"net/http"
	"time"
)

func index(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	fmt.Fprintf(w, "home page")
	fmt.Printf("index elasped:%fs\n", time.Since(start).Seconds())
}

func greeting(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	name := r.FormValue("name")
	if name == "" {
		name = "world"
	}

	fmt.Fprintf(w, "hello %s", name)
	fmt.Printf("greeting elasped:%fs\n", time.Since(start).Seconds())
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", index)
	mux.HandleFunc("/greeting", greeting)

	http.ListenAndServe(":8000", mux)
}
