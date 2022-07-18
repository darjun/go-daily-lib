package main

import (
	"fmt"

	"github.com/RoaringBitmap/roaring"
)

func main() {
	bm1 := roaring.BitmapOf(1, 2, 3, 4, 5, 100, 1000)
	bm2 := roaring.BitmapOf(1, 100, 500)
	bm3 := roaring.BitmapOf(1, 10, 1000)

	bmAnd := roaring.ParAnd(4, bm1, bm2, bm3)
	fmt.Println(bmAnd.String())         // {1}
	fmt.Println(bmAnd.GetCardinality()) // 1
	fmt.Println(bmAnd.Contains(1))      // true
	fmt.Println(bmAnd.Contains(100))    // false

	bmOr := roaring.ParOr(4, bm1, bm2, bm3)
	fmt.Println(bmOr.String())         // {1,2,3,4,5,10,100,500,1000}
	fmt.Println(bmOr.GetCardinality()) // 9
	fmt.Println(bmOr.Contains(10))     // true
}
