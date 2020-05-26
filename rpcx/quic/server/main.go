package main

import (
	"context"
	"crypto/tls"
	"errors"

	"github.com/smallnest/rpcx/server"
)

type Args struct {
	A, B int
}

type Quotient struct {
	Quo, Rem int
}

type Arith int

func (t *Arith) Mul(cxt context.Context, args *Args, reply *int) error {
	*reply = args.A * args.B
	return nil
}

func (t *Arith) Div(cxt context.Context, args *Args, quo *Quotient) error {
	if args.B == 0 {
		return errors.New("divide by 0")
	}

	quo.Quo = args.A / args.B
	quo.Rem = args.A % args.B
	return nil
}

func main() {
	cert, _ := tls.LoadX509KeyPair("server.pem", "server.key")
	config := &tls.Config{Certificates: []tls.Certificate{cert}}

	s := server.NewServer(server.WithTLSConfig(config))
	s.RegisterName("Arith", new(Arith), "")
	s.Serve("quic", "localhost:8972")
}
