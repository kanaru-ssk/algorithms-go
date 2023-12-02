package main

import (
	"testing"
)

func TestKnapsack(t *testing.T) {
	if Knapsack(
		500,
		[]Item{
			{Id: 1, Cost: 150, Value: 4},
			{Id: 2, Cost: 200, Value: 7},
			{Id: 3, Cost: 250, Value: 8},
			{Id: 4, Cost: 200, Value: 5},
			{Id: 5, Cost: 50, Value: 1}}) != 16 {
		t.Fatalf("error")
	}

	if Knapsack(
		720,
		[]Item{
			{Id: 1, Cost: 150, Value: 4},
			{Id: 2, Cost: 200, Value: 7},
			{Id: 3, Cost: 250, Value: 8},
			{Id: 4, Cost: 200, Value: 5},
			{Id: 5, Cost: 50, Value: 1}}) != 21 {
		t.Fatalf("error")
	}
}
