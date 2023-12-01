package main

func Gcd(a, b uint32) uint32 {
	if b == 0 {
		return a
	} else {
		return Gcd(b, a%b)
	}
}
