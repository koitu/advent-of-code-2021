package main

import (
	"testing"
)

func TestPart1(t *testing.T) {
	result := alignCrabCost("test.txt", false)
	if result != 37 {
		t.Fatalf("part 1 sanity test did not pass, should have gotten 37 but got %d", result)
	}
	t.Logf("%d", alignCrabCost("input.txt", false))
}

func TestPart2(t *testing.T) {
	result := alignCrabCost("test.txt", true)
	if result != 168 {
		t.Fatalf("part 2 sanity test did not pass, should have gotten 168 but got %d", result)
	}
	t.Logf("%d", alignCrabCost("input.txt", true))
}
