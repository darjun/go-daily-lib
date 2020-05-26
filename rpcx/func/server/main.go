package main

import (
	"context"
	"errors"

	"github.com/smallnest/rpcx/server"
)

type Args struct {
	A, B int
}

type Quotient struct {
	Quo, Rem int
}


func Mul(cxt context.Context, args *Args, reply *int) error {
	*reply = args.A * args.B
	return nil
}

func Div(cxt context.Context, args *Args, quo *Quotient) error {
	if args.B == 0 {
		return errors.New("divide by 0")
	}

	quo.Quo = args.A / args.B
	quo.Rem = args.A % args.B
	return nil
}

func main() {
	s := server.NewServer()
	s.RegisterFunction("function", Mul, "")
	s.RegisterFunction("function", Div, "")
	s.Serve("tcp", ":8972")
}
