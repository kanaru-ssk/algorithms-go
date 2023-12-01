package main

import (
	"testing"
)

func TestFactorial(t *testing.T) {
	if Factorial(3) != 6 {
		t.Fatalf("error")
	}
}
