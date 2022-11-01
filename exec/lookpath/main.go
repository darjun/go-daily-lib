package main

import (
	"fmt"
	"os/exec"
)

func main() {
	path, err := exec.LookPath("ls")
	if err != nil {
		fmt.Printf("no cmd ls: %v\n", err)
	} else {
		fmt.Printf("find ls in path:%s\n", path)
	}

	path, err = exec.LookPath("not-exist")
	if err != nil {
		fmt.Printf("no cmd not-exist: %v\n", err)
	} else {
		fmt.Printf("find not-exist in path:%s\n", path)
	}
}
