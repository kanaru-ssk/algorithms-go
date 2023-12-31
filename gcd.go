package main

func Gcd(a, b uint32) uint32 {
	if b == 0 {
		return a
	}

	// 割り切れるまで再帰的に呼び出す
	return Gcd(b, a%b)
}

func GcdVec(vec []uint32) uint32 {
	result := uint32(0)

	for _, num := range vec {
		result = Gcd(result, num)
	}

	return result
}
