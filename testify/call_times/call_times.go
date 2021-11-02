package call_times

import "fmt"

type IExample interface {
	Hello(n int) int
}

type Example struct {
}

func (e *Example) Hello(n int) int {
	fmt.Printf("Hello with %d\n", n)
	return n
}

func ExampleFunc(e IExample) {
	for n := 1; n <= 3; n++ {
		for i := 0; i <= n; i++ {
			e.Hello(n)
		}
	}
}
