package main

import "fmt"

type DailyMission struct {
	// ...
}

func NewDailyMission() *DailyMission {
	d := &DailyMission{}
	bus.Subscribe("UserLevelUp", d.OnUserLevelUp)
	return d
}

func (d *DailyMission) OnUserLevelUp(oldLevel, newLevel uint32) {
	fmt.Printf("daily mission old level:%d new level:%d\n", oldLevel, newLevel)
}