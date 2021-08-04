package fib

func Fib1(n int) int {
	if n <= 1 {
		return n
	}

	return Fib1(n-1) + Fib1(n-2)
}

func fibHelper(n int, m map[int]int) int {
	if n <= 1 {
		return n
	}

	if v, ok := m[n]; ok {
		return v
	}

	v := fibHelper(n-2, m) + fibHelper(n-1, m)
	m[n] = v
	return v
}

func Fib2(n int) int {
	m := make(map[int]int)
	return fibHelper(n, m)
}

func Fib3(n int) int {
	if n <= 1 {
		return n
	}

	f1, f2 := 0, 1
	for i := 2; i <= n; i++ {
		f1, f2 = f2, f1+f2
	}

	return f2
}
