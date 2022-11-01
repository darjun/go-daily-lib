package main

import (
	"bytes"
	"log"
	"os"
	"os/exec"
)

func main() {
	cmd := exec.Command("cat")
	cmd.Stdin = bytes.NewBufferString("hello\nworld")
	cmd.Stdout = os.Stdout
	err := cmd.Run()
	if err != nil {
		log.Fatalf("cmd.Run() failed: %v\n", err)
	}
}
