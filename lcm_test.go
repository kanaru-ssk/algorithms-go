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

func TestLcmVec(t *testing.T) {
	if LcmVec([]uint32{18, 12}) != 36 {
		t.Fatalf("error")
	}

	if LcmVec([]uint32{36, 12, 48, 120}) != 720 {
		t.Fatalf("error")
	}

	if LcmVec([]uint32{36, 12, 48, 0}) != 0 {
		t.Fatalf("error")
	}

	if LcmVec([]uint32{0, 0, 0}) != 0 {
		t.Fatalf("error")
	}

	if LcmVec([]uint32{3}) != 3 {
		t.Fatalf("error")
	}

	if LcmVec([]uint32{}) != 0 {
		t.Fatalf("error")
	}
}
