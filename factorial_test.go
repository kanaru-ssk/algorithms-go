package main

import (
	"testing"
)

func TestFactorial(t *testing.T) {
	samples := []struct {
		Input    uint32
		Expected uint32
	}{
		{Input: 0, Expected: 1},
		{Input: 1, Expected: 1},
		{Input: 2, Expected: 2},
		{Input: 3, Expected: 6},
		{Input: 4, Expected: 24},
		{Input: 5, Expected: 120},
		{Input: 10, Expected: 3628800},
	}

	for _, s := range samples {
		if result := Factorial(s.Input); result != s.Expected {
			t.Errorf("Factorial(%v) = %v, but expected %v", s.Input, result, s.Expected)
		}
	}
}
