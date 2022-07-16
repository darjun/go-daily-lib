package main

import (
	"fmt"
	"math/rand"

	"github.com/bits-and-blooms/bitset"
)

type Player struct {
	sign *bitset.BitSet
}

func NewPlayer(sign uint) *Player {
	return &Player{
		sign: bitset.From([]uint64{uint64(sign)}),
	}
}

func (this *Player) Sign(day uint) {
	this.sign.Set(day)
}

func (this *Player) IsSigned(day uint) bool {
	return this.sign.Test(day)
}

func main() {
	player := NewPlayer(1) // 第一天已经签到
	for day := uint(2); day <= 7; day++ {
		if rand.Intn(100)&1 == 0 {
			player.Sign(day - 1)
		}
	}

	for day := uint(1); day <= 7; day++ {
		if player.IsSigned(day - 1) {
			fmt.Printf("day:%d signed\n", day)
		}
	}
}
