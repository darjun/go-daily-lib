package main

import (
	"fmt"

	"github.com/RoaringBitmap/roaring"
)

func main() {
	bm := roaring.BitmapOf(1, 2, 3, 4, 5, 100, 1000)

	i := bm.Iterator()
	for i.HasNext() {
		fmt.Println(i.Next())
	}
}
