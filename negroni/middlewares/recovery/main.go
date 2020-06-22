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
	n.Use(negroni.NewRecovery())
	n.UseHandler(mux)

	http.ListenAndServe(":3000", n)
}
