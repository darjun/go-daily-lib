//+build wireinject

package main

import "github.com/google/wire"

func InitPlayer() Player {
	wire.Build(NewMission, wire.FieldsOf(new(Mission), "Player"))
	return Player{}
}

func InitMonster() Monster {
	wire.Build(NewMission, wire.FieldsOf(new(Mission), "Monster"))
	return Monster{}
}
