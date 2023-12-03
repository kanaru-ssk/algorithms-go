package main

import (
	"testing"
)

func TestLcm(t *testing.T) {
	samples := []struct {
		Input    []uint32
		Expected uint32
	}{
		{Input: []uint32{18, 12}, Expected: 36},
		{Input: []uint32{12, 18}, Expected: 36},
		{Input: []uint32{1, 1}, Expected: 1},
		{Input: []uint32{10, 0}, Expected: 0},
		{Input: []uint32{0, 10}, Expected: 0},
		{Input: []uint32{0, 0}, Expected: 0},
	}

	for _, s := range samples {
		if result := Lcm(s.Input[0], s.Input[1]); result != s.Expected {
			t.Errorf("Lcm(%v, %v) = %v, but expected %v", s.Input[0], s.Input[1], result, s.Expected)
		}
	}
}

func TestLcmVec(t *testing.T) {
	samples := []struct {
		Input    []uint32
		Expected uint32
	}{
		{Input: []uint32{18, 12}, Expected: 36},
		{Input: []uint32{36, 12, 48, 120}, Expected: 720},
		{Input: []uint32{0, 0, 0}, Expected: 0},
		{Input: []uint32{3}, Expected: 3},
		{Input: []uint32{0}, Expected: 0},
		{Input: []uint32{}, Expected: 0},
	}

	for _, s := range samples {
		if result := LcmVec(s.Input); result != s.Expected {
			t.Errorf("LcmVec(%v) = %v, but expected %v", s.Input, result, s.Expected)
		}
	}
}
