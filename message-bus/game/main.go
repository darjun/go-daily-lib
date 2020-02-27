package main

import "time"

func main() {
	p := NewPlayer()
	NewDailyMission()
	NewAchievement()

	p.LevelUp()
	p.LevelUp()
	p.LevelUp()

	time.Sleep(1000)
}