package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

func main() {
	cmd := exec.Command("bash", "-c", "./test.sh")

	nameEnv := "NAME=darjun"
	ageEnv := "AGE=18"

	newEnv := append(os.Environ(), nameEnv, ageEnv)
	cmd.Env = newEnv

	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("cmd.Run() failed: %v\n", err)
	}

	fmt.Println(string(out))
}
