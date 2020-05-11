package main

import (
	"flag"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

var (
	addr *string
)

type Args struct {
	A, B int
}

type Arith int

func (t *Arith) Multiply(args *Args, reply *int) error {
	*reply = args.A * args.B
	return nil
}

func init() {
	addr = flag.String("addr", ":1111", "addr to listen")
}

func main() {
	flag.Parse()

	l, err := net.Listen("tcp", *addr)
	if err != nil {
		log.Fatal("listen error:", err)
	}

	arith := new(Arith)
	rpc.Register(arith)

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Fatal("accept error:", err)
		}

		go rpc.ServeCodec(jsonrpc.NewServerCodec(conn))
	}
}
