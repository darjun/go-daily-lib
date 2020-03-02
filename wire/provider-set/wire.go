//+build wireinject

package main

import "github.com/google/wire"

var monsterPlayerSet = wire.NewSet(NewMonster, NewPlayer)

func InitEndingA(name string) EndingA {
	wire.Build(monsterPlayerSet, NewEndingA)
	return EndingA{}
}

func InitEndingB(name string) EndingB {
	wire.Build(monsterPlayerSet, NewEndingB)
	return EndingB{}
}
