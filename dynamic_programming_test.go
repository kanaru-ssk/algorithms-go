package main

import (
	"testing"
)

func TestProgrammingKnapsack(t *testing.T) {
	type Item struct {
		Id, Cost, Value uint32
	}
	items := []Item{
		{Id: 1, Cost: 150, Value: 4},
		{Id: 2, Cost: 200, Value: 7},
		{Id: 3, Cost: 250, Value: 8},
		{Id: 4, Cost: 200, Value: 5},
		{Id: 5, Cost: 50, Value: 1}}
	calcCell := func(y, x uint32, table [][]uint32, items []Item) uint32 {
		if y == 0 {
			return 0
		}
		if x < items[y-1].Cost {
			return table[y-1][x]
		}
		return Max(table[y-1][x-items[y-1].Cost]+items[y-1].Value, table[y-1][x])
	}

	capacity1 := uint32(500)
	if DynamicProgramming(items, capacity1, calcCell) != 16 {
		t.Fatalf("error")
	}

	capacity2 := uint32(720)
	if DynamicProgramming(items, capacity2, calcCell) != 21 {
		t.Fatalf("error")
	}
}

func TestBinPacking(t *testing.T) {

	capacity := uint32(10)
	items := []uint32{4, 7, 8, 5, 1}
	calcCell := func(y, x uint32, table [][]bool, items []uint32) bool {
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

	if DynamicProgramming(items, capacity, calcCell) != true {
		t.Fatalf("error")
	}
}

func TestFibonacci(t *testing.T) {
	var items []uint32
	calcCell := func(y, x uint32, table [][]uint32, items []uint32) uint32 {
		if x <= 0 {
			return 0
		}
		if x <= 2 {
			return 1
		}
		return table[y][x-1] + table[y][x-2]
	}

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

		if actual := DynamicProgramming(items, s.Input, calcCell); actual != s.Expected {
			t.Errorf("Abs(%v) = %v, want %v", s.Input, actual, s.Expected)
		}
	}
}
