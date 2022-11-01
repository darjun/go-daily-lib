package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
)

func cal(w http.ResponseWriter, r *http.Request) {
	year := r.URL.Query().Get("year")
	month := r.URL.Query().Get("month")

	f, _ := os.OpenFile("out.txt", os.O_CREATE|os.O_WRONLY, os.ModePerm)
	buf := bytes.NewBuffer(nil)
	mw := io.MultiWriter(w, f, buf)

	cmd := exec.Command("cal", month, year)
	cmd.Stdout = mw
	cmd.Stderr = mw

	err := cmd.Run()
	if err != nil {
		log.Fatalf("cmd.Run() failed: %v\n", err)
	}

	fmt.Println(buf.String())
}

func main() {
	http.HandleFunc("/cal", cal)
	http.ListenAndServe(":8080", nil)
}
