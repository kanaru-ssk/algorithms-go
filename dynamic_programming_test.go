package main

import (
	"testing"
)

func TestBinPacking(t *testing.T) {
	items := []Item{
		{Id: 1, Cost: 150, Value: 150},
		{Id: 2, Cost: 200, Value: 200},
		{Id: 3, Cost: 250, Value: 250},
		{Id: 4, Cost: 200, Value: 200},
		{Id: 5, Cost: 50, Value: 50}}

	capacity1 := uint32(500)
	result1 := DynamicProgramming(capacity1, items)
	if result1 != 500 {
		t.Fatalf("error")
	}
	if result1 == capacity1 != true {
		t.Fatalf("error")
	}

	capacity2 := uint32(720)
	result2 := DynamicProgramming(capacity2, items)
	if result2 != 700 {
		t.Fatalf("error")
	}
	if result2 == capacity2 != false {
		t.Fatalf("error")
	}
}

func TestKnapsack(t *testing.T) {
	items := []Item{
		{Id: 1, Cost: 150, Value: 4},
		{Id: 2, Cost: 200, Value: 7},
		{Id: 3, Cost: 250, Value: 8},
		{Id: 4, Cost: 200, Value: 5},
		{Id: 5, Cost: 50, Value: 1}}

	capacity1 := uint32(500)
	if DynamicProgramming(capacity1, items) != 16 {
		t.Fatalf("error")
	}

	capacity2 := uint32(720)
	if DynamicProgramming(capacity2, items) != 21 {
		t.Fatalf("error")
	}
}
