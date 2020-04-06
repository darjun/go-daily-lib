package main

import (
	"fmt"
	"time"

	"github.com/shirou/gopsutil/host"
)

func main() {
	timestamp, _ := host.BootTime()
	t := time.Unix(int64(timestamp), 0)
	fmt.Println(t.Local().Format("2006-01-02 15:04:05"))
}
