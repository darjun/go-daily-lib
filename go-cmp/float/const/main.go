package main

import "fmt"

const (
	_  = 1 << (10 * iota)
	KB // 1024
	MB // 1048576
	GB // 1073741824
	TB // ‭1099511627776‬
	PB // ‭1125899906842624‬
	EB // ‭1152921504606846976‬
	ZB // ‭1180591620717411303424‬
	YB // ‭1208925819614629174706176‬
)

func main() {
	// constant ‭1180591620717411303424‬ overflows int
	// fmt.Println(ZB)

	// constant 1208925819614629174706176 overflows uint64
	// var mem uint64 = YB

	fmt.Println(YB / ZB)
}
