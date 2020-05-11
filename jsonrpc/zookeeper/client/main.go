package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"
)

var (
	zookeeperAddr *string
)

func init() {
	zookeeperAddr = flag.String("addr", ":2181", "zookeeper address")
}

type Args struct {
	A, B int
}

func main() {
	flag.Parse()

	fmt.Println(*zookeeperAddr)
	p := NewProxy(*zookeeperAddr)
	p.Connect()

	go p.Run()

	for i := 0; i < 10; i++ {
		var reply int
		args := &Args{rand.Intn(1000), rand.Intn(1000)}
		p.Call("Arith.Multiply", args, &reply)
		fmt.Printf("%d*%d=%d\n", args.A, args.B, reply)
	}

	time.Sleep(1 * time.Minute)

	for i := 0; i < 100; i++ {
		var reply int
		args := &Args{rand.Intn(1000), rand.Intn(1000)}
		p.Call("Arith.Multiply", args, &reply)
		fmt.Printf("%d*%d=%d\n", args.A, args.B, reply)
	}
}
