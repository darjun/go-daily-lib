package main

import (
	"fmt"

	"github.com/shirou/gopsutil/process"
)

func main() {
	var rootProcess *process.Process
	processes, _ := process.Processes()
	for _, p := range processes {
		if p.Pid == 0 {
			rootProcess = p
			break
		}
	}

	fmt.Println(rootProcess)

	fmt.Println("children:")
	children, _ := rootProcess.Children()
	for _, p := range children {
		fmt.Println(p)
	}
}
