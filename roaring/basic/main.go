package main

import (
	"fmt"

	"github.com/RoaringBitmap/roaring"
)

func main() {
	bm1 := roaring.BitmapOf(1, 2, 3, 4, 5, 100, 1000)
	fmt.Println(bm1.String())         // {1,2,3,4,5,100,1000}
	fmt.Println(bm1.GetCardinality()) // 7
	fmt.Println(bm1.Contains(3))      // true

	bm2 := roaring.BitmapOf(1, 100, 500)
	fmt.Println(bm2.String())         // {1,100,500}
	fmt.Println(bm2.GetCardinality()) // 3
	fmt.Println(bm2.Contains(300))    // false

	bm3 := roaring.New()
	bm3.Add(1)
	bm3.Add(11)
	bm3.Add(111)
	fmt.Println(bm3.String())         // {1,11,111}
	fmt.Println(bm3.GetCardinality()) // 3
	fmt.Println(bm3.Contains(11))     // true

	bm1.Or(bm2)                       // 执行并集
	fmt.Println(bm1.String())         // {1,2,3,4,5,100,500,1000}
	fmt.Println(bm1.GetCardinality()) // 8
	fmt.Println(bm1.Contains(500))    // true

	bm2.And(bm3)                      // 执行交集
	fmt.Println(bm2.String())         // {1}
	fmt.Println(bm2.GetCardinality()) // 1
	fmt.Println(bm2.Contains(1))      // true
}
