package main

func PrimeFactorize(n uint32) map[uint32]uint32 {
	result := make(map[uint32]uint32)

	if n == 0 {
		return result
	}

	for i := uint32(2); i*i <= n; i++ {
		if n%i != 0 {
			continue
		}

		var e uint32 = 0
		for n%i == 0 {
			e++
			n /= i
		}

		result[i] = e
	}

	if n != 1 {
		result[n] = 1
	}

	return result
}
