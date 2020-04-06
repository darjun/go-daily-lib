package main

import (
	"encoding/json"
	"fmt"

	"github.com/shirou/gopsutil/disk"
)

func main() {
	infos, _ := disk.Partitions(true)
	for _, info := range infos {
		data, _ := json.MarshalIndent(info, "", "  ")
		fmt.Println(string(data))
	}
}
