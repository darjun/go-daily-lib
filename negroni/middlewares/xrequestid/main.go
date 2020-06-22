package main

import (
	"net/http"
	"fmt"

	"github.com/pilu/xrequestid"
	"github.com/urfave/negroni"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "X-Request-Id is `%s`", r.Header.Get("X-Request-Id"))
	})

	n := negroni.New()
	n.Use(xrequestid.New(16))
	n.UseHandler(mux)
	n.Run(":3000")
}