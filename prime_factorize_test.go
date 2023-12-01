package main

import (
	"testing"
)

func TestPrimeFactorize(t *testing.T) {
	result1 := PrimeFactorize(0)
	if len(result1) != 0 {
		t.Fatalf("error")
	}

	result2 := PrimeFactorize(1)
	if len(result2) != 0 {
		t.Fatalf("error")
	}

	result3 := PrimeFactorize(2)
	if len(result3) != 1 || result3[2] != 1 {
		t.Fatalf("error")
	}

	result4 := PrimeFactorize(12)
	if len(result4) != 2 || result4[2] != 2 || result4[3] != 1 {
		t.Fatalf("error")
	}

	result5 := PrimeFactorize(53)
	if len(result5) != 1 || result5[53] != 1 {
		t.Fatalf("error")
	}

	result6 := PrimeFactorize(4876)
	if len(result6) != 3 || result6[2] != 2 || result6[23] != 1 || result6[53] != 1 {
		t.Fatalf("error")
	}
}
