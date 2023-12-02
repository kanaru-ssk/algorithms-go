package main

func Fib1(n uint32) uint32 {
	if n <= 2 {
		return 1
	}

	return Fib1(n-1) + Fib1(n-2)
}

func Fib2(n uint32) uint32 {
	if n == 0 || n == 1 {
		return n
	}

	a := make([]uint32, n)

	a[0] = 1
	a[1] = 1

	for i := uint32(2); i < n; i++ {
		a[i] = a[i-1] + a[i-2]
	}

	return a[n-1]
}
