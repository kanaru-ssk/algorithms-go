package main

import (
	"testing"
)

func TestHelloName(t *testing.T) {
	result := Factorial(3)

	if result != 6 {
		t.Fatalf("error")
	}
}
