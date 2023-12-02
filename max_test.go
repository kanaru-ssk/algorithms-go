package main

import (
	"testing"
)

func TestMax(t *testing.T) {
	if Max(1, 2) != 2 {
		t.Fatalf("error")
	}

	if Max(2, 2) != 2 {
		t.Fatalf("error")
	}

	if Max(3, 2) != 3 {
		t.Fatalf("error")
	}

	if Max(0, 0) != 0 {
		t.Fatalf("error")
	}
}
