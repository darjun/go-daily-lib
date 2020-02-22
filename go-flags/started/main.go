package main

import (
	"fmt"

	"github.com/jessevdk/go-flags"
)

type Option struct {
	Verbose []bool `short:"v" long:"verbose" description:"Show verbose debug message"`
}

func main() {
	var opt Option
	flags.Parse(&opt)

	fmt.Println(opt.Verbose)
}
