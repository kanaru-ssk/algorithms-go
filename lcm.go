package main

func Lcm(a, b uint32) uint32 {
	if a == 0 && b == 0 {
		return 0
	}

	return a / Gcd(a, b) * b
}
