package main

import (
	"testing"
)

func TestGcd(t *testing.T) {
	if Gcd(18, 12) != 6 {
		t.Fatalf("error")
	}

	if Gcd(12, 18) != 6 {
		t.Fatalf("error")
	}

	if Gcd(1, 1) != 1 {
		t.Fatalf("error")
	}

	if Gcd(10, 0) != 10 {
		t.Fatalf("error")
	}

	if Gcd(0, 10) != 10 {
		t.Fatalf("error")
	}

	if Gcd(0, 0) != 0 {
		t.Fatalf("error")
	}
}
