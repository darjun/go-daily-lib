// Multi-room chat app implemented in gotalk
package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/rsms/gotalk"
)

type Room struct {
	Name     string `json:"name"`
	mu       sync.RWMutex
	messages []*Message
}

func (room *Room) appendMessage(m *Message) {
	room.mu.Lock()
	defer room.mu.Unlock()
	room.messages = append(room.messages, m)
}

type Message struct {
	Author string `json:"author"`
	Body   string `json:"body"`
}

type NewMessage struct {
	Room    string  `json:"room"`
	Message Message `json:"message"`
}

type RoomMap map[string]*Room

var (
	rooms   RoomMap
	roomsmu sync.RWMutex
	socks   map[*gotalk.WebSocket]int
	socksmu sync.RWMutex
)

func onConnect(s *gotalk.WebSocket) {
	// Keep track of connected sockets
	socksmu.Lock()
	defer socksmu.Unlock()
	socks[s] = 1

	// When the connection closes, remove the socket from our lists of connected peers
	s.CloseHandler = func(s *gotalk.WebSocket, _ int) {
		fmt.Printf("Peer %s diconnected\n", s)
		socksmu.Lock()
		defer socksmu.Unlock()
		delete(socks, s)
	}

	// log a message when a peer connects
	fmt.Printf("Peer %s connected on %s\n", s, s.Conn().LocalAddr())

	// Send list of rooms
	roomsmu.RLock()
	defer roomsmu.RUnlock()
	s.Notify("rooms", rooms)

	// Assign the socket a random username
	username := randomName()
	s.UserData = username
	s.Notify("username", username)
}

func broadcast(name string, in interface{}) {
	socksmu.RLock()
	defer socksmu.RUnlock()
	for s := range socks {
		s.Notify(name, in)
	}
}

func findRoom(name string) *Room {
	roomsmu.RLock()
	defer roomsmu.RUnlock()
	return rooms[name]
}

func createRoom(name string) *Room {
	roomsmu.Lock()
	defer roomsmu.Unlock()
	room := rooms[name]
	if room == nil {
		room = &Room{Name: name}
		rooms[name] = room
		broadcast("rooms", rooms)
	}
	return room
}

// Instead of asking the user for her/his name, we randomly assign one
var names struct {
	First []string
	Last  []string
}

func randomName() string {
	first := names.First[rand.Intn(len(names.First))]
	return first
	// last := names.Last[rand.Intn(len(names.Last))][:1]
	// return first + " " + last
}

func main() {
	socks = make(map[*gotalk.WebSocket]int)
	rooms = make(RoomMap)

	// Load names data
	if namesjson, err := ioutil.ReadFile("names.json"); err != nil {
		panic("failed to read names.json: " + err.Error())
	} else if err := json.Unmarshal(namesjson, &names); err != nil {
		panic("failed to read names.json: " + err.Error())
	}
	rand.Seed(time.Now().UTC().UnixNano())

	// Add some sample rooms and messages
	createRoom("animals").appendMessage(
		&Message{randomName(), "I like cats"})
	createRoom("jokes").appendMessage(
		&Message{randomName(), "Two tomatoes walked across the street ..."})
	createRoom("golang").appendMessage(
		&Message{randomName(), "func(func(func(func())func()))func()"})

	// Register our handlers
	gotalk.Handle("list-messages", func(roomName string) ([]*Message, error) {
		room := findRoom(roomName)
		if room == nil {
			return nil, errors.New("no such room")
		}
		return room.messages, nil
	})

	gotalk.Handle("send-message", func(s *gotalk.Sock, r NewMessage) error {
		if len(r.Message.Body) == 0 {
			return errors.New("empty message")
		}
		username, _ := s.UserData.(string)
		room := findRoom(r.Room)
		room.appendMessage(&Message{username, r.Message.Body})
		r.Message.Author = username
		broadcast("newmsg", &r)
		return nil
	})

	gotalk.Handle("create-room", func(name string) (*Room, error) {
		if len(name) == 0 {
			return nil, errors.New("empty name")
		}
		return createRoom(name), nil
	})

	// Serve gotalk at "/gotalk/"
	gh := gotalk.WebSocketHandler()
	gh.OnConnect = onConnect
	routes := &http.ServeMux{}
	server := &http.Server{Addr: "localhost:1235", Handler: routes}
	routes.Handle("/gotalk/", gh)

	// Handle static files
	routes.Handle("/", http.FileServer(http.Dir(".")))

	// Enable SIGINT (^C) to perform a graceful shut down
	done := enableGracefulShutdown(server, 5*time.Second)

	// Start server
	fmt.Printf("Listening on http://%s/\n", server.Addr)
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		panic(err)
	}

	// Wait for shutdown to complete
	<-done
}

func enableGracefulShutdown(server *http.Server, timeout time.Duration) chan struct{} {
	server.RegisterOnShutdown(func() {
		// close all connected sockets
		fmt.Printf("graceful shutdown: closing sockets\n")
		socksmu.RLock()
		defer socksmu.RUnlock()
		for s := range socks {
			s.CloseHandler = nil // avoid deadlock on socksmu (also not needed)
			s.Close()
		}
	})
	done := make(chan struct{})
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT)
	go func() {
		<-quit // wait for signal

		fmt.Printf("graceful shutdown initiated\n")
		ctx, cancel := context.WithTimeout(context.Background(), timeout)
		defer cancel()

		server.SetKeepAlivesEnabled(false)
		if err := server.Shutdown(ctx); err != nil {
			fmt.Printf("server.Shutdown error: %s\n", err)
		}

		fmt.Printf("graceful shutdown complete\n")
		close(done)
	}()
	return done
}
