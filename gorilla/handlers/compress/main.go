package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello World")
}

type greeting string

func (g greeting) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome, %s", g)
}

func main() {
	r := mux.NewRouter()
	r.Use(handlers.CompressHandler)
	r.HandleFunc("/", index)
	r.Handle("/greeting/", greeting("dj"))

	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
