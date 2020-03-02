//+build wireinject

package main

import "github.com/google/wire"

func InitMission(m MonsterParam, p PlayerParam) Mission {
	wire.Build(NewMonster, NewPlayer, NewMission)
	return Mission{}
}
