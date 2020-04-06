package main

import (
	"encoding/json"
	"fmt"

	"github.com/shirou/gopsutil/disk"
)

func main() {
	info, _ := disk.Usage("D:/code/golang")
	data, _ := json.MarshalIndent(info, "", "  ")
	fmt.Println(string(data))
}
