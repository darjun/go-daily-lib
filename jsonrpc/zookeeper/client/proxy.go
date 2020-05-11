package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/rpc"
	"net/rpc/jsonrpc"
	"strings"
	"sync"
	"time"

	"github.com/samuel/go-zookeeper/zk"
)

type Proxy struct {
	zookeeper     string
	clients       map[string]*rpc.Client
	events        <-chan zk.Event
	zookeeperConn *zk.Conn
	mutex         sync.Mutex
}

func NewProxy(addr string) *Proxy {
	return &Proxy{
		zookeeper: addr,
		clients:   make(map[string]*rpc.Client),
	}
}

func (p *Proxy) Call(method string, args interface{}, reply interface{}) error {
	var client *rpc.Client
	var addr string
	idx := rand.Int31n(int32(len(p.clients)))
	var i int32
	p.mutex.Lock()
	for a, c := range p.clients {
		if i == idx {
			client = c
			addr = a
			break
		}
		i++
	}
	p.mutex.Unlock()

	fmt.Println("use", addr)
	return client.Call(method, args, reply)
}

func (p *Proxy) Connect() {
	c, _, err := zk.Connect([]string{p.zookeeper}, time.Second) //*10)
	if err != nil {
		panic(err)
	}

	data, _, event, err := c.GetW("/rpcserver")
	if err != nil {
		panic(err)
	}

	p.events = event
	p.zookeeperConn = c

	p.CreateClients(string(data))
}

func (p *Proxy) CreateClients(server string) {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	addrs := strings.Split(server, ",")
	allAddr := make(map[string]struct{})
	for _, addr := range addrs {
		allAddr[addr] = struct{}{}
		if _, exist := p.clients[addr]; exist {
			continue
		}

		client, err := jsonrpc.Dial("tcp", addr)
		if err != nil {
			log.Println("jsonrpc Dial error:", err)
			continue
		}

		p.clients[addr] = client
		log.Println("new addr:", addr)
	}

	for addr, oldClient := range p.clients {
		if _, exist := allAddr[addr]; !exist {
			oldClient.Close()

			delete(p.clients, addr)
			log.Println("delete addr", addr)
		}
	}
}

func (p *Proxy) Run() {
	for {
		select {
		case event := <-p.events:
			if event.Type == zk.EventNodeDataChanged {
				data, _, err := p.zookeeperConn.Get("/rpcserver")
				if err != nil {
					log.Println("get zookeeper data failed:", err)
					continue
				}

				p.CreateClients(string(data))
			}
		}
	}
}
