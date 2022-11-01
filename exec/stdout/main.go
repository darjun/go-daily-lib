package main

import (
	"log"
	"os"
	"os/exec"
)

func main() {
	cmd := exec.Command("cal")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		log.Fatalf("cmd.Run() failed: %v\n", err)
	}
}
