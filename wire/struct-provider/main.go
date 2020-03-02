package main

import "fmt"

type Monster struct {
	Name string
}

func NewMonster() Monster {
	return Monster{Name: "kitty"}
}

type Player struct {
	Name string
}

func NewPlayer(name string) Player {
	return Player{Name: name}
}

type EndingA struct {
	Player  Player
	Monster Monster
}

func (p EndingA) Appear() {
	fmt.Printf("%s defeats %s, world peace!\n", p.Player.Name, p.Monster.Name)
}

type EndingB struct {
	Player  Player
	Monster Monster
}

func (p EndingB) Appear() {
	fmt.Printf("%s defeats %s, but become monster, world darker!\n", p.Player.Name, p.Monster.Name)
}

func main() {
	endingA := InitEndingA("dj")
	endingB := InitEndingB("xj")

	endingA.Appear()
	endingB.Appear()
}
