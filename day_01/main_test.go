package main

import (
	"testing"
)

func TestPart1(t *testing.T) {
	result := countIncreases("test.txt", 1)
	if result != 7 {
		t.Fatalf("part 1 sanity test did not pass, should have gotten 7 but got %d", result)
	}
	t.Logf("%d", countIncreases("input.txt", 1))
}

func TestPart2(t *testing.T) {
	result := countIncreases("test.txt", 3)
	if result != 5 {
		t.Fatalf("part 2 sanity test did not pass, should have gotten 5 but got %d", result)
	}
	t.Logf("%d", countIncreases("input.txt", 3))
}
