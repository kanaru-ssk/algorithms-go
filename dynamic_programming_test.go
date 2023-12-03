package main

import (
	"testing"
)

func TestFibonacci(t *testing.T) {
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

	fib := func(y, x uint32, table [][]uint32, items []uint32) uint32 {
		if x == 0 {
			return 0
		}
		if x <= 2 {
			return 1
		}
		return table[y][x-1] + table[y][x-2]
	}

	for _, s := range samples {
		if actual := DynamicProgramming([]uint32{}, s.Input, fib); actual != s.Expected {
			t.Errorf("DynamicProgramming([], %v, fib) = %v, but expected %v", s.Input, actual, s.Expected)
		}
	}
}

func TestBinPacking(t *testing.T) {
	items := []uint32{4, 7, 8, 5, 1}
	type Input struct {
		Items    []uint32
		Capacity uint32
	}
	samples := []struct {
		Input    Input
		Expected bool
	}{
		{Input: Input{Items: items, Capacity: 10}, Expected: true},
		{Input: Input{Items: items, Capacity: 22}, Expected: false},
	}

	binPacking := func(y, x uint32, table [][]bool, items []uint32) bool {
		if y == 0 {
			return x == 0
		}
		if table[y-1][x] {
			return true
		}
		if x < items[y-1] {
			return false
		}
		return table[y-1][x-items[y-1]]
	}

	for _, s := range samples {
		if result := DynamicProgramming(s.Input.Items, s.Input.Capacity, binPacking); result != s.Expected {
			t.Errorf("DynamicProgramming(%v, %v, binPacking) = %v, but expected %v", s.Input.Items, s.Input.Capacity, result, s.Expected)
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
	knapsack := func(y, x uint32, table [][]uint32, items []Item) uint32 {
		if y == 0 {
			return 0
		}
		if x < items[y-1].Cost {
			return table[y-1][x]
		}
		return Max(table[y-1][x-items[y-1].Cost]+items[y-1].Value, table[y-1][x])
	}

	for _, s := range samples {
		if result := DynamicProgramming(s.Input.Items, s.Input.Capacity, knapsack); result != s.Expected {
			t.Errorf("DynamicProgramming(%v, %v, knapsack) = %v, but expected %v", s.Input.Items, s.Input.Capacity, result, s.Expected)
		}
	}
}
