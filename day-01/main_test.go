package main

import (
	"testing"
)

func TestPart1(t *testing.T) {
	c := countIncreases("test.txt", 1)
	if c != 7 {
		t.Fatalf("part 1 sanity test did not pass, should have gotten 7 but got %d", c)
	}
	t.Logf("%d", countIncreases("input.txt", 1))
}

func TestPart2(t *testing.T) {
	if countIncreases("test.txt", 3) != 5 {
		t.Fatalf("part 1 sanity test did not pass, should have gotten 5")
	}
	t.Logf("%d", countIncreases("input.txt", 3))
}
