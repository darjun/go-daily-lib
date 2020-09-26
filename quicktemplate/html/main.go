package main

import (
	"log"
	"net/http"

	"github.com/darjun/go-daily-lib/quicktemplate/html/templates"
)

func index(w http.ResponseWriter, r *http.Request) {
	templates.WriteIndex(w, r.FormValue("name"))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", index)

	server := &http.Server{
		Handler: mux,
		Addr:    ":8080",
	}

	log.Fatal(server.ListenAndServe())
}
