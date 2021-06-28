package main

import (
	"fmt"
	"log"

	"github.com/go-resty/resty/v2"
)

func main() {
	client := resty.New()

	resp, err := client.R().
		EnableTrace().
		Get("https://httpbin.org/get")

	if err != nil {
		log.Fatal(err)
	}

	ti := resp.Request.TraceInfo()
	fmt.Println("Request Trace Info:")
	fmt.Println("DNSLookup:", ti.DNSLookup)
	fmt.Println("ConnTime:", ti.ConnTime)
	fmt.Println("TCPConnTime:", ti.TCPConnTime)
	fmt.Println("TLSHandshake:", ti.TLSHandshake)
	fmt.Println("ServerTime:", ti.ServerTime)
	fmt.Println("ResponseTime:", ti.ResponseTime)
	fmt.Println("TotalTime:", ti.TotalTime)
	fmt.Println("IsConnReused:", ti.IsConnReused)
	fmt.Println("IsConnWasIdle:", ti.IsConnWasIdle)
	fmt.Println("ConnIdleTime:", ti.ConnIdleTime)
	fmt.Println("RequestAttempt:", ti.RequestAttempt)
	fmt.Println("RemoteAddr:", ti.RemoteAddr.String())
}
