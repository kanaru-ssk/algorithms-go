package main

import (
	"testing"
)

func TestFib1(t *testing.T) {
	if Fib1(1) != 1 {
		t.Fatalf("error")
	}

	if Fib1(6) != 8 {
		t.Fatalf("error")
	}

	if Fib1(0) != 1 {
		t.Fatalf("error")
	}
}

func TestFib2(t *testing.T) {
	if Fib2(1) != 1 {
		t.Fatalf("error")
	}

	if Fib2(6) != 8 {
		t.Fatalf("error")
	}

	if Fib2(0) != 0 {
		t.Fatalf("error")
	}
}
