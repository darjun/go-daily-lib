package main

import (
	"fmt"

	"github.com/shirou/gopsutil/host"
)

func main() {
	version, _ := host.KernelVersion()
	fmt.Println(version)

	platform, family, version, _ := host.PlatformInformation()
	fmt.Println("platform:", platform)
	fmt.Println("family:", family)
	fmt.Println("version:", version)
}
