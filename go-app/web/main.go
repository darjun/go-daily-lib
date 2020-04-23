package main

import (
	"log"
	"net/http"

	"github.com/maxence-charriere/go-app/v6/pkg/app"
)

func main() {
	h := &app.Handler{
		Title:  "Go-App",
		Author: "dj",
	}

	if err := http.ListenAndServe(":8080", h); err != nil {
		log.Fatal(err)
	}
}
