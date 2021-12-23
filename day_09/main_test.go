package main

import (
	"testing"
)

func TestPart1(t *testing.T) {
	result := sumOfLowPoints("test.txt", false)
	if result != 15 {
		t.Fatalf("part 1 sanity test did not pass, should have gotten 15 but got %d", result)
	}
	t.Logf("%d", sumOfLowPoints("input.txt", false))
}

func TestPart2(t *testing.T) {
	result := sumOfLowPoints("test.txt", true)
	if result != 1134 {
		t.Fatalf("part 2 sanity test did not pass, should have gotten 1134 but got %d", result)
	}
	t.Logf("%d", sumOfLowPoints("input.txt", true))
}
