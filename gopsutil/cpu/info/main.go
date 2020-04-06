package main

import (
	"encoding/json"
	"fmt"

	"github.com/shirou/gopsutil/cpu"
)

func main() {
	infos, _ := cpu.Info()
	for _, info := range infos {
		data, _ := json.MarshalIndent(info, "", " ")
		fmt.Print(string(data))
	}
}
