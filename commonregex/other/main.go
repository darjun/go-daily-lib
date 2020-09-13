package main

import (
	"fmt"

	"github.com/mingrammer/commonregex"
)

func main() {
	text := `mac address: ac:de:48:00:11:22, ip: 192.168.3.20, md5: fdbf72fdabb67ea6ef7ff5155a44def4`

	macList := commonregex.MACAddresses(text)
	ipList := commonregex.IPs(text)
	md5List := commonregex.MD5Hexes(text)

	fmt.Println("mac list:", macList)
	fmt.Println("ip list:", ipList)
	fmt.Println("md5 list:", md5List)
}
