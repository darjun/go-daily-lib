package main

import (
	"fmt"

	"github.com/shirou/gopsutil/winservices"
)

func main() {
	services, _ := winservices.ListServices()

	for _, service := range services {
		newservice, _ := winservices.NewService(service.Name)
		newservice.GetServiceDetail()
		fmt.Println("Name:", newservice.Name, "Binary Path:", newservice.Config.BinaryPathName, "State: ", newservice.Status.State)
	}
}
