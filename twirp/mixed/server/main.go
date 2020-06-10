package main

import (
	"context"
	"log"
	"net/http"

	"github.com/darjun/go-daily-lib/twirp/get-started/proto"
)

type Server struct{}

func (s *Server) Say(ctx context.Context, request *proto.Request) (*proto.Response, error) {
	return &proto.Response{Text: request.GetText()}, nil
}

func greeting(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	if name == "" {
		name = "world"
	}

	w.Write([]byte("hi," + name))
}

func main() {
	server := &Server{}
	twirpHandler := proto.NewEchoServer(server, nil)

	mux := http.NewServeMux()
	mux.Handle(twirpHandler.PathPrefix(), twirpHandler)
	mux.HandleFunc("/greeting", greeting)

	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal(err)
	}
}
