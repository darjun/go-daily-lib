package main

import (
	"encoding/json"
	"fmt"

	"github.com/shirou/gopsutil/mem"
)

func main() {
	swapMemory, _ := mem.SwapMemory()
	data, _ := json.MarshalIndent(swapMemory, "", " ")
	fmt.Println(string(data))
}
