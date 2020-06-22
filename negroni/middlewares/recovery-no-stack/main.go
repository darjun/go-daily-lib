package main

import (
	"net/http"

	"github.com/urfave/negroni"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		panic("internal server error")
	})

	n := negroni.New()
	r := negroni.NewRecovery()
	r.PrintStack = false
	n.Use(r)
	n.UseHandler(mux)

	http.ListenAndServe(":3000", n)
}
