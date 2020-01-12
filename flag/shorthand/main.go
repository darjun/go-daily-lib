package main

import (
	"flag"
	"fmt"
)

var (
    logLevel string
)

func init() {
	const (
		defaultLogLevel = "DEBUG"
		usage = "set log level value"
	)
	
	flag.StringVar(&logLevel, "log_type", defaultLogLevel, usage)
	flag.StringVar(&logLevel, "l", defaultLogLevel, usage + "(shorthand)")
}

func main() {
    flag.Parse()
    
	fmt.Println("log level:", logLevel)
}