package main

import (
	"encoding/json"
	"fmt"

	"github.com/shirou/gopsutil/disk"
)

func main() {
	mapStat, _ := disk.IOCounters()
	for name, stat := range mapStat {
		fmt.Println(name)
		data, _ := json.MarshalIndent(stat, "", "  ")
		fmt.Println(string(data))
	}
}
