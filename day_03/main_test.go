package main

import (
	"testing"
)

func TestPart1(t *testing.T) {
	result := binaryDiagnostic("test.txt", false)
	if result != 198 {
		t.Fatalf("part 1 sanity test did not pass, should have gotten 198 but got %d", result)
	}
	t.Logf("%d", binaryDiagnostic("input.txt", false))
}

func TestPart2(t *testing.T) {
	result := binaryDiagnostic("test.txt", true)
	if result != 230 {
		t.Fatalf("part 2 sanity test did not pass, should have gotten 230 but got %d", result)
	}
	t.Logf("%d", binaryDiagnostic("input.txt", true))
}
