package main

func Gcd(a, b uint32) uint32 {
	// 割り切れるまで再帰的に呼び出す
	if b != 0 {
		return Gcd(b, a%b)
	}

	return a
}
