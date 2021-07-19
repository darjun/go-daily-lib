package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	InitBooksRouter(r)
	InitMoviesRouter(r)

	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
