package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/darjun/go-daily-lib/twirp/get-started/proto"
	"github.com/twitchtv/twirp"
)

type Server struct{}

func (s *Server) Say(ctx context.Context, request *proto.Request) (*proto.Response, error) {
	token := ctx.Value("token").(string)
	fmt.Println("token:", token)

	err := twirp.SetHTTPResponseHeader(ctx, "Token-Lifecycle", "60")
	if err != nil {
		return nil, twirp.InternalErrorWith(err)
	}
	return &proto.Response{Text: request.GetText()}, nil
}

func WithTwirpToken(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		token := r.Header.Get("Twirp-Token")
		ctx = context.WithValue(ctx, "token", token)
		r = r.WithContext(ctx)

		h.ServeHTTP(w, r)
	})
}

func main() {
	server := &Server{}
	twirpHandler := proto.NewEchoServer(server, nil)
	wrapped := WithTwirpToken(twirpHandler)

	http.ListenAndServe(":8080", wrapped)
}
