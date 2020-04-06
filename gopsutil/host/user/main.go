package main

import (
	"encoding/json"
	"fmt"

	"github.com/shirou/gopsutil/host"
)

func main() {
	users, _ := host.Users()
	for _, user := range users {
		data, _ := json.MarshalIndent(user, "", " ")
		fmt.Println(string(data))
	}
}
