package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	ptTemplate.ExecuteTemplate(w, "home.tpl", nil)
}

func main() {
	r := mux.NewRouter()
	r.Use(PanicRecover, WithLogger, Metric)
	r.HandleFunc("/", HomeHandler)
	InitBooksRouter(r)
	InitMoviesRouter(r)
	InitLoginRouter(r)

	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
