package main

import (
	"testing"
)

func TestPart1(t *testing.T) {
	result := countOverlaps("test.txt", false)
	if result != 5 {
		t.Fatalf("part 1 sanity test did not pass, should have gotten 5 but got %d", result)
	}
	t.Logf("%d", countOverlaps("input.txt", false))
}

func TestPart2(t *testing.T) {
	result := countOverlaps("test.txt", true)
	if result != 12 {
		t.Fatalf("part 2 sanity test did not pass, should have gotten 12 but got %d", result)
	}
	t.Logf("%d", countOverlaps("input.txt", true))
}
