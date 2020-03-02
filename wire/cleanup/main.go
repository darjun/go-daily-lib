package main

import (
	"errors"
	"fmt"
	"log"
	"time"
)

type Monster struct {
	Name string
}

func NewMonster() Monster {
	return Monster{Name: "kitty"}
}

type Player struct {
	Name string
}

func NewPlayer(name string) (Player, func(), error) {
	cleanup := func() {
		fmt.Println("cleanup!")
	}
	if time.Now().Unix()%2 == 0 {
		return Player{}, cleanup, errors.New("player dead")
	}
	return Player{Name: name}, cleanup, nil
}

type Mission struct {
	Player  Player
	Monster Monster
}

func NewMission(p Player, m Monster) Mission {
	return Mission{p, m}
}

func (m Mission) Start() {
	fmt.Printf("%s defeats %s, world peace!\n", m.Player.Name, m.Monster.Name)
}

func main() {
	mission, cleanup, err := InitMission("dj")
	if err != nil {
		log.Fatal(err)
	}

	mission.Start()
	cleanup()
}
