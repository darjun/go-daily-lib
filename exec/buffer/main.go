package main

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
)

func main() {
	buf := bytes.NewBuffer(nil)
	cmd := exec.Command("cal")
	cmd.Stdout = buf
	cmd.Stderr = buf
	err := cmd.Run()
	if err != nil {
		log.Fatalf("cmd.Run() failed: %v\n", err)
	}

	fmt.Println(buf.String())
}
