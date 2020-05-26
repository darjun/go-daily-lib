package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	"github.com/smallnest/rpcx/client"
)

type Args struct {
	A, B int
}

type Quotient struct {
	Quo, Rem int
}

var (
	addr     = flag.String("addr", ":8972", "service address")
	zkAddr   = flag.String("zkAddr", "127.0.0.1:2181", "zookeeper address")
	basePath = flag.String("basePath", "/services/math", "service base path")
)

func main() {
	flag.Parse()

	d := client.NewZookeeperDiscovery(*basePath, "Arith", []string{*zkAddr}, nil)
	xclient := client.NewXClient("Arith", client.Failtry, client.RandomSelect, d, client.DefaultOption)
	defer xclient.Close()

	args := &Args{A: 10, B: 20}
	var reply int

	err := xclient.Call(context.Background(), "Mul", args, &reply)
	if err != nil {
		log.Fatalf("failed to call: %v", err)
	}

	fmt.Printf("%d * %d = %d\n", args.A, args.B, reply)

	args = &Args{50, 20}
	var quo Quotient
	err = xclient.Call(context.Background(), "Div", args, &quo)
	if err != nil {
		log.Fatalf("failed to call: %v", err)
	}

	fmt.Printf("%d * %d = %d...%d\n", args.A, args.B, quo.Quo, quo.Rem)
}
