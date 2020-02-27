package main

type Player struct {
	level uint32
}

func NewPlayer() *Player {
	return &Player{}
}

func (p *Player) LevelUp() {
	oldLevel := p.level
	newLevel := p.level+1
	p.level++

	bus.Publish("UserLevelUp", oldLevel, newLevel)
}