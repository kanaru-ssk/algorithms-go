package main

import (
	"testing"
)

func TestMax(t *testing.T) {
	samples := []struct {
		Input    []uint32
		Expected uint32
	}{
		{Input: []uint32{1, 2}, Expected: 2},
		{Input: []uint32{2, 2}, Expected: 2},
		{Input: []uint32{3, 2}, Expected: 3},
		{Input: []uint32{0, 0}, Expected: 0},
	}

	for _, s := range samples {
		if result := Max(s.Input[0], s.Input[1]); result != s.Expected {
			t.Errorf("Max(%v, %v) = %v, but expected %v", s.Input[0], s.Input[1], result, s.Expected)
		}
	}
}
