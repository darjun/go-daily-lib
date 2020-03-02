//+build wireinject

package main

import "github.com/google/wire"

func InitMission(name string) (Mission, func(), error) {
	wire.Build(NewMonster, NewPlayer, NewMission)
	return Mission{}, nil, nil
}
