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

func search(b *bitset.BitSet, visited map[string]struct{}, path *[]*bitset.BitSet) bool {
	if !IsStateValid(b) {
		return false
	}

	if _, exist := visited[b.String()]; exist {
		// 状态已遍历
		return false
	}

	*path = append(*path, b.Clone())
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

		if search(b, visited, path) {
			return true
		}

		// 状态恢复
		b.Flip(index)
		if index != FARMER {
			b.Flip(FARMER)
		}
	}
	*path = (*path)[:len(*path)-1]

	return false
}

var names = []string{"农夫", "狼", "羊", "白菜"}

func PrintState(b *bitset.BitSet) {
	fmt.Println("=======================")
	fmt.Println("河左岸：")
	for index := uint(FARMER); index <= CABBAGE; index++ {
		if !b.Test(index) {
			fmt.Println(names[index])
		}
	}

	fmt.Println("河右岸：")
	for index := uint(FARMER); index <= CABBAGE; index++ {
		if b.Test(index) {
			fmt.Println(names[index])
		}
	}
	fmt.Println("=======================")
}

func PrintMove(cur, next *bitset.BitSet) {
	for index := uint(WOLF); index <= CABBAGE; index++ {
		if cur.Test(index) != next.Test(index) {
			if !cur.Test(FARMER) {
				fmt.Printf("农夫将【%s】从河左岸带到河右岸\n", names[index])
			} else {
				fmt.Printf("农夫将【%s】从河右岸带到河左岸\n", names[index])

			}
			return
		}
	}

	if !cur.Test(FARMER) {
		fmt.Println("农夫独自从河左岸到河右岸")
	} else {
		fmt.Println("农夫独自从河右岸到河左岸")
	}
}

func PrintPath(path []*bitset.BitSet) {
	cur := path[0]
	PrintState(cur)

	for i := 1; i < len(path); i++ {
		next := path[i]
		PrintMove(cur, next)
		PrintState(next)
		cur = next
	}
}

func main() {
	b := bitset.New(4)

	visited := make(map[string]struct{})
	var path []*bitset.BitSet
	if search(b, visited, &path) {
		PrintPath(path)
	}
}
