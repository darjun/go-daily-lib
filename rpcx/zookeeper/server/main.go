package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"log"
	"time"

	metrics "github.com/rcrowley/go-metrics"
	"github.com/smallnest/rpcx/server"
	"github.com/smallnest/rpcx/serverplugin"
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

type Arith int

func (t *Arith) Mul(cxt context.Context, args *Args, reply *int) error {
	fmt.Println("Mul on", *addr)
	*reply = args.A * args.B
	return nil
}

func (t *Arith) Div(cxt context.Context, args *Args, quo *Quotient) error {
	fmt.Println("Div on", *addr)
	if args.B == 0 {
		return errors.New("divide by 0")
	}

	quo.Quo = args.A / args.B
	quo.Rem = args.A % args.B
	return nil
}

func main() {
	flag.Parse()

	p := &serverplugin.ZooKeeperRegisterPlugin{
		ServiceAddress:   "tcp@" + *addr,
		ZooKeeperServers: []string{*zkAddr},
		BasePath:         *basePath,
		Metrics:          metrics.NewRegistry(),
		UpdateInterval:   time.Minute,
	}
	if err := p.Start(); err != nil {
		log.Fatal(err)
	}

	s := server.NewServer()
	s.Plugins.Add(p)

	s.RegisterName("Arith", new(Arith), "")
	s.Serve("tcp", *addr)
}
