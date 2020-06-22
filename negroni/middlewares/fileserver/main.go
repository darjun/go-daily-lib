package main

import (
	"fmt"
	"net/http"

	"github.com/urfave/negroni"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "hello world")
	})

	n := negroni.New()
	n.UseHandler(http.FileServer(http.Dir("./public")))
	n.UseHandler(mux)

	http.ListenAndServe(":3000", n)
}
