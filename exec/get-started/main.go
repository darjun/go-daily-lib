package main

import (
	"log"
	"os/exec"
)

func main() {
	cmd := exec.Command("cal")
	err := cmd.Run()
	if err != nil {
		log.Fatalf("cmd.Run() failed: %v\n", err)
	}
}
