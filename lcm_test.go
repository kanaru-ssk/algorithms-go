package main

import (
	"testing"
)

func TestLcm(t *testing.T) {
	if Lcm(18, 12) != 36 {
		t.Fatalf("error")
	}

	if Lcm(12, 18) != 36 {
		t.Fatalf("error")
	}

	if Lcm(1, 1) != 1 {
		t.Fatalf("error")
	}

	if Lcm(10, 0) != 0 {
		t.Fatalf("error")
	}

	if Lcm(0, 10) != 0 {
		t.Fatalf("error")
	}

	if Lcm(0, 0) != 0 {
		t.Fatalf("error")
	}
}
