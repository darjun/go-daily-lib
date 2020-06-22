package main

import (
	"fmt"
	"math/rand"
	"net/http"

	"github.com/urfave/negroni"
)

func RandomMiddleware(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	if rand.Int31n(100) <= 50 {
		fmt.Fprintf(w, "hello from RandomMiddleware")
	} else {
		next(w, r)
	}
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World!")
	})

	n := negroni.New()
	n.Use(negroni.HandlerFunc(RandomMiddleware))
	n.UseHandler(mux)

	http.ListenAndServe(":3000", n)
}
