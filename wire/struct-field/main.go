package main

import (
	"fmt"
)

type Monster struct {
	Name string
}

type Player struct {
	Name string
}

type Mission struct {
	Player  Player
	Monster Monster
}

func NewMission() Mission {
	p := Player{Name: "dj"}
	m := Monster{Name: "kitty"}

	return Mission{p, m}
}

func (m Mission) Start() {
	fmt.Printf("%s defeats %s, world peace!\n", m.Player.Name, m.Monster.Name)
}

func main() {
	p := InitPlayer()

	fmt.Println(p.Name)
}
