package main

func Lcm(a, b uint32) uint32 {
	if a == 0 && b == 0 {
		return 0
	}

	return a / Gcd(a, b) * b
}

func LcmVec(vec []uint32) uint32 {
	if len(vec) == 0 {
		return 0
	}

	var result = vec[0]

	for _, num := range vec {
		result = Lcm(result, num)
	}

	return result
}
