package main

import (
	"bytes"
	"fmt"

	"github.com/RoaringBitmap/roaring"
)

func main() {
	bm := roaring.BitmapOf(1, 3, 5, 7, 100, 300, 500, 700)

	buf := &bytes.Buffer{}
	bm.WriteTo(buf)

	newBm := roaring.New()
	newBm.ReadFrom(buf)
	if bm.Equals(newBm) {
		fmt.Println("write and read back ok.")
	}
}
