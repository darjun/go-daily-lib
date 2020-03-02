//+build wireinject

package main

import "github.com/google/wire"

func InitMission(name string) (Mission, error) {
	wire.Build(NewMonster, NewPlayer, NewMission)
	return Mission{}, nil
}
