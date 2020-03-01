package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/ThreeDotsLabs/watermill/pubsub/gochannel"
)

var (
	logger = watermill.NewStdLogger(false, false)
)

func main() {
	router, err := message.NewRouter(message.RouterConfig{}, logger)
	if err != nil {
		panic(err)
	}

	pubSub := gochannel.NewGoChannel(gochannel.Config{}, logger)
	go publishMessages(pubSub)

	router.AddMiddleware(myMiddleware{Name: "dj"}.Middleware)

	router.AddHandler("myhandler", "in_topic", pubSub, "out_topic", pubSub, myHandler{}.Handler)
	router.AddNoPublisherHandler("print_in_messages", "in_topic", pubSub, printMessages)
	router.AddNoPublisherHandler("print_out_messages", "out_topic", pubSub, printMessages)

	ctx := context.Background()
	if err := router.Run(ctx); err != nil {
		panic(err)
	}
}

func publishMessages(publisher message.Publisher) {
	for {
		msg := message.NewMessage(watermill.NewUUID(), []byte("Hello, world!"))
		if err := publisher.Publish("in_topic", msg); err != nil {
			panic(err)
		}

		time.Sleep(time.Second)
	}
}

func printMessages(msg *message.Message) error {
	fmt.Printf("\n> Received message: %s\n> %s\n>\n", msg.UUID, string(msg.Payload))
	return nil
}

type myHandler struct {
}

func (m myHandler) Handler(msg *message.Message) ([]*message.Message, error) {
	log.Println("myHandler received message", msg.UUID)

	msg = message.NewMessage(watermill.NewUUID(), []byte("message produced by myHandler"))
	return message.Messages{msg}, nil
}

type myMiddleware struct {
	Name string
}

func (m myMiddleware) Middleware(h message.HandlerFunc) message.HandlerFunc {
	return func(message *message.Message) ([]*message.Message, error) {
		fields := watermill.LogFields{"name": m.Name}
		logger.Info("myMiddleware before", fields)
		ms, err := h(message)
		logger.Info("myMiddleware after", fields)
		return ms, err
	}
}
