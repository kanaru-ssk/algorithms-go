package main

import (
	"testing"
)

func TestLcs(t *testing.T) {
	samples := []struct {
		Input    [2]string
		Expected uint32
	}{
		{Input: [2]string{"ACCGGTCGAGTGCGCGGAAGCCGGCCGAA", "GTCGTTCGGAATGCCGTTGCTCTGTAAA"}, Expected: 20},
		{Input: [2]string{"A00B0000C", "XXAXBXCXX"}, Expected: 3},
	}

	for _, s := range samples {
		if result := Lcs(s.Input[0], s.Input[1]); result != s.Expected {
			t.Errorf("Max(%v, %v) = %v, but expected %v", s.Input[0], s.Input[1], result, s.Expected)
		}
	}
}
