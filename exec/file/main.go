package main

import (
	"log"
	"os"
	"os/exec"
)

func main() {
	f, err := os.OpenFile("out.txt", os.O_WRONLY|os.O_CREATE, os.ModePerm)
	if err != nil {
		log.Fatalf("os.OpenFile() failed: %v\n", err)
	}

	cmd := exec.Command("cal")
	cmd.Stdout = f
	cmd.Stderr = f
	err = cmd.Run()
	if err != nil {
		log.Fatalf("cmd.Run() failed: %v\n", err)
	}
}
