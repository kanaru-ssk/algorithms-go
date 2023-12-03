package main

import (
	"reflect"
	"testing"
)

func TestPrimeFactorize(t *testing.T) {
	samples := []struct {
		Input    uint32
		Expected map[uint32]uint32
	}{
		{Input: 1, Expected: map[uint32]uint32{}},
		{Input: 2, Expected: map[uint32]uint32{2: 1}},
		{Input: 12, Expected: map[uint32]uint32{2: 2, 3: 1}},
		{Input: 53, Expected: map[uint32]uint32{53: 1}},
		{Input: 4876, Expected: map[uint32]uint32{2: 2, 23: 1, 53: 1}},
		{Input: 0, Expected: map[uint32]uint32{}},
	}

	for _, s := range samples {
		if result := PrimeFactorize(s.Input); !reflect.DeepEqual(result, s.Expected) {
			t.Errorf("PrimeFactorize(%v) = %v, but expected %v", s.Input, result, s.Expected)
		}
	}
}
