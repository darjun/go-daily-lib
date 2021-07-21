package main

import (
	"log"
	"net/http"

	"errors"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func PANIC(w http.ResponseWriter, r *http.Request) {
	panic(errors.New("unexpected error"))
}

func main() {
	r := mux.NewRouter()
	r.Use(handlers.RecoveryHandler(handlers.PrintRecoveryStack(true)))
	r.HandleFunc("/", PANIC)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
