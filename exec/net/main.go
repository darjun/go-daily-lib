package main

import (
	"log"
	"net/http"
	"os/exec"
)

func cal(w http.ResponseWriter, r *http.Request) {
	year := r.URL.Query().Get("year")
	month := r.URL.Query().Get("month")

	cmd := exec.Command("cal", month, year)
	cmd.Stdout = w
	cmd.Stderr = w

	err := cmd.Run()
	if err != nil {
		log.Fatalf("cmd.Run() failed: %v\n", err)
	}
}

func main() {
	http.HandleFunc("/cal", cal)
	http.ListenAndServe(":8080", nil)
}
