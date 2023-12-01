package main

func Factorial(n uint32) uint32 {
	var result uint32 = 1
	for i := n; i > 1; i-- {
		result *= i
	}

	return result
}
