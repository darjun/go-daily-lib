package main

import "fmt"

type Achievement struct {
	// ...
}

func NewAchievement() *Achievement {
	a := &Achievement{}
	bus.Subscribe("UserLevelUp", a.OnUserLevelUp)
	return a
}

func (a *Achievement) OnUserLevelUp(oldLevel, newLevel uint32) {
	fmt.Printf("daily mission old level:%d new level:%d\n", oldLevel, newLevel)
}