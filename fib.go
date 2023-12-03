package main

func Fib1(n uint32) uint32 {
	if n == 0 {
		return 0
	}
	if n <= 2 {
		return 1
	}

	return Fib1(n-1) + Fib1(n-2)
}

func Fib2(n uint32) uint32 {
	calcCell := func(y, x uint32, table [][]uint32, items []uint32) uint32 {
		if x == 0 {
			return 0
		}
		if x <= 2 {
			return 1
		}
		return table[y][x-1] + table[y][x-2]
	}

	return DynamicProgramming([]uint32{}, n, calcCell)
}
