package main

import (
	"bufio"
	"encoding/json"
	"io"
	"log"
	"net"
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

type JsonServerCodec struct {
	rwc    io.ReadWriteCloser
	dec    *json.Decoder
	enc    *json.Encoder
	encBuf *bufio.Writer
	closed bool
}

func NewJsonServerCodec(conn io.ReadWriteCloser) *JsonServerCodec {
	buf := bufio.NewWriter(conn)
	return &JsonServerCodec{conn, json.NewDecoder(conn), json.NewEncoder(buf), buf, false}
}

func (c *JsonServerCodec) ReadRequestHeader(r *rpc.Request) error {
	return c.dec.Decode(r)
}

func (c *JsonServerCodec) ReadRequestBody(body interface{}) error {
	return c.dec.Decode(body)
}

func (c *JsonServerCodec) WriteResponse(r *rpc.Response, body interface{}) (err error) {
	if err = c.enc.Encode(r); err != nil {
		if c.encBuf.Flush() == nil {
			log.Println("rpc: json error encoding response:", err)
			c.Close()
		}
		return
	}
	if err = c.enc.Encode(body); err != nil {
		if c.encBuf.Flush() == nil {
			log.Println("rpc: json error encoding body:", err)
			c.Close()
		}
		return
	}
	return c.encBuf.Flush()
}

func (c *JsonServerCodec) Close() error {
	if c.closed {
		return nil
	}
	c.closed = true
	return c.rwc.Close()
}

func main() {
	l, err := net.Listen("tcp", ":1234")
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

		go rpc.ServeCodec(NewJsonServerCodec(conn))
	}
}
