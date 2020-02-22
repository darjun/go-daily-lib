package main

import (
	"fmt"
	"net/http"

	"go.uber.org/dig"
)

type Handler struct {
	Greeting string
	Path     string
}

func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s from %s", h.Greeting, h.Path)
}

func NewHello1Handler() HandlerResult {
	return HandlerResult{
		Handler: Handler{
			Path:     "/hello1",
			Greeting: "welcome",
		},
	}
}

func NewHello2Handler() HandlerResult {
	return HandlerResult{
		Handler: Handler{
			Path:     "/hello2",
			Greeting: "😄",
		},
	}
}

type HandlerResult struct {
	dig.Out

	Handler Handler `group:"server"`
}

type HandlerParams struct {
	dig.In

	Handlers []Handler `group:"server"`
}

func RunServer(params HandlerParams) error {
	mux := http.NewServeMux()
	for _, h := range params.Handlers {
		mux.Handle(h.Path, h)
	}

	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
	if err := server.ListenAndServe(); err != nil {
		return err
	}

	return nil
}

func main() {
	container := dig.New()

	container.Provide(NewHello1Handler)
	container.Provide(NewHello2Handler)

	container.Invoke(RunServer)
}
