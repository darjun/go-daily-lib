package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello world")
}

func main() {
	r := mux.NewRouter()
	r.Use(handlers.CanonicalHost("http://www.gorillatoolkit.org", 302))
	r.HandleFunc("/", index)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
