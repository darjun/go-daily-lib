package main

import (
	"log"
	"net/http"
	"net/rpc"
)

type Args struct {
	A, B int
}

type Arith int

func (t *Arith) Multiply(args *Args, reply *int) error {
	*reply = args.A * args.B
	return nil
}

func main() {
	arith := new(Arith)
	server := rpc.NewServer()
	server.RegisterName("math", arith)
	server.HandleHTTP(rpc.DefaultRPCPath, rpc.DefaultDebugPath)

	if err := http.ListenAndServe(":1234", nil); err != nil {
		log.Fatal("serve error:", err)
	}
}
