package main

import (
	"testing"
)

func TestFib(t *testing.T) {
	samples := []struct {
		Input    uint32
		Expected uint32
	}{
		{Input: 0, Expected: 0},
		{Input: 1, Expected: 1},
		{Input: 2, Expected: 1},
		{Input: 3, Expected: 2},
		{Input: 4, Expected: 3},
		{Input: 5, Expected: 5},
		{Input: 6, Expected: 8},
		{Input: 20, Expected: 6765},
	}

	for _, s := range samples {
		if result := Fib1(s.Input); result != s.Expected {
			t.Errorf("Fib1(%v) = %v, but expected %v", s.Input, result, s.Expected)
		}
	}

	for _, s := range samples {
		if result := Fib2(s.Input); result != s.Expected {
			t.Errorf("Fib2(%v) = %v, but expected %v", s.Input, result, s.Expected)
		}
	}
}
