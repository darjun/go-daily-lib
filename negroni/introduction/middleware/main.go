package main

import (
	"fmt"
	"net/http"
	"time"
)

func elasped(h func(w http.ResponseWriter, r *http.Request)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		start := time.Now()
		h(w, r)
		fmt.Printf("path:%s elasped:%fs\n", path, time.Since(start).Seconds())
	}
}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "home page")
}

func greeting(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	if name == "" {
		name = "world"
	}

	fmt.Fprintf(w, "hello %s", name)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", elasped(index))
	mux.HandleFunc("/greeting", elasped(greeting))

	http.ListenAndServe(":8000", mux)
}
