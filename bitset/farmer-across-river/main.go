package main

import (
	"fmt"

	"github.com/bits-and-blooms/bitset"
)

const (
	FARMER = iota
	WOLF
	SHEEP
	CABBAGE
)

func IsStateValid(state *bitset.BitSet) bool {
	if state.Test(WOLF) == state.Test(SHEEP) &&
		state.Test(WOLF) != state.Test(FARMER) {
		// 狼和羊在同一边，并且不和农夫在同一边
		// 狼会吃掉羊，非法
		return false
	}

	if state.Test(SHEEP) == state.Test(CABBAGE) &&
		state.Test(SHEEP) != state.Test(FARMER) {
		// 羊和白菜在同一边，并且不和农夫在同一边
		// 羊会吃掉白菜，非法
		return false
	}

	return true
}

func search(b *bitset.BitSet, visited map[string]struct{}) bool {
	if !IsStateValid(b) {
		return false
	}

	if _, exist := visited[b.String()]; exist {
		// 状态已遍历
		return false
	}

	if b.Count() == 4 {
		return true
	}

	visited[b.String()] = struct{}{}
	for index := uint(FARMER); index <= CABBAGE; index++ {
		if b.Test(index) != b.Test(FARMER) {
			// 与农夫不在一边，不能带上船
			continue
		}

		// 带到对岸去
		b.Flip(index)
		if index != FARMER {
			// 如果 index 为 FARMER，表示不带任何东西
			b.Flip(FARMER)
		}

		if search(b, visited) {
			return true
		}

		// 状态恢复
		b.Flip(index)
		if index != FARMER {
			b.Flip(FARMER)
		}
	}

	return false
}

func main() {
	b := bitset.New(4)

	visited := make(map[string]struct{})
	fmt.Println(search(b, visited))
}
