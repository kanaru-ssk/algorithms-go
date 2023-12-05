package main

import (
	"testing"
)

func TestFibonacci(t *testing.T) {
	type Input struct {
		n uint32
	}
	samples := []struct {
		Input    Input
		Expected uint32
	}{
		{Input: Input{n: 0}, Expected: 0},
		{Input: Input{n: 1}, Expected: 1},
		{Input: Input{n: 2}, Expected: 1},
		{Input: Input{n: 3}, Expected: 2},
		{Input: Input{n: 4}, Expected: 3},
		{Input: Input{n: 5}, Expected: 5},
		{Input: Input{n: 6}, Expected: 8},
		{Input: Input{n: 20}, Expected: 6765},
	}

	fib := func(table []uint32, iterators []int, elements Input) uint32 {
		i := iterators[0]
		if i <= 0 {
			return 0
		}
		if i <= 2 {
			return 1
		}
		return table[i-1] + table[i-2]
	}

	for _, s := range samples {
		if actual := DynamicProgramming(s.Input, fib); actual != s.Expected {
			t.Errorf("DynamicProgramming([], %v, fib) = %v, but expected %v", s.Input, actual, s.Expected)
		}
	}
}

func TestBinPacking(t *testing.T) {
	type Input struct {
		Items    []uint32
		Capacity uint32
	}
	items := []uint32{4, 7, 8, 5, 1}
	samples := []struct {
		Input    Input
		Expected bool
	}{
		{Input: Input{Items: items, Capacity: 10}, Expected: true},
		{Input: Input{Items: items, Capacity: 22}, Expected: false},
	}

	binPacking := func(table []bool, iterators []int, elements Input) bool {
		y := uint32(iterators[0])
		x := uint32(iterators[1])
		h := elements.Capacity + 1
		if y == 0 {
			return x == 0
		}
		if table[(y-1)*h+x] {
			return true
		}
		if x < items[y-1] {
			return false
		}
		return table[(y-1)*h+x-items[y-1]]
	}

	for _, s := range samples {
		if result := DynamicProgramming(s.Input, binPacking); result != s.Expected {
			t.Errorf("DynamicProgramming(%v, binPacking) = %v, but expected %v", s.Input, result, s.Expected)
		}
	}
}

func TestKnapsack(t *testing.T) {
	type Item struct {
		Id, Cost, Value uint32
	}
	items := []Item{
		{Id: 1, Cost: 150, Value: 4},
		{Id: 2, Cost: 200, Value: 7},
		{Id: 3, Cost: 250, Value: 8},
		{Id: 4, Cost: 200, Value: 5},
		{Id: 5, Cost: 50, Value: 1}}
	type Input struct {
		Items    []Item
		Capacity uint32
	}
	samples := []struct {
		Input    Input
		Expected uint32
	}{
		{Input: Input{Items: items, Capacity: 500}, Expected: 16},
		{Input: Input{Items: items, Capacity: 720}, Expected: 21},
	}
	knapsack := func(table []uint32, iterators []int, elements Input) uint32 {
		y := uint32(iterators[0])
		x := uint32(iterators[1])
		h := elements.Capacity + 1
		if y == 0 {
			return 0
		}
		if x < items[y-1].Cost {
			return table[(y-1)*h+x]
		}
		return Max(table[(y-1)*h+x-items[y-1].Cost]+items[y-1].Value, table[(y-1)*h+x])
	}

	for _, s := range samples {
		if result := DynamicProgramming(s.Input, knapsack); result != s.Expected {
			t.Errorf("DynamicProgramming(%v, knapsack) = %v, but expected %v", s.Input, result, s.Expected)
		}
	}
}
