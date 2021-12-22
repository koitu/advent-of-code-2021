package main

import (
	"testing"
)

func TestPart1(t *testing.T) {
	result := sevenSegmentMatches("test.txt", false)
	if result != 26 {
		t.Fatalf("part 1 sanity test did not pass, should have gotten 26 but got %d", result)
	}
	t.Logf("%d", sevenSegmentMatches("input.txt", false))
}

func TestPart2(t *testing.T) {
	result := sevenSegmentMatches("test.txt", true)
	if result != 61229 {
		t.Fatalf("part 2 sanity test did not pass, should have gotten 61229 but got %d", result)
	}
	t.Logf("%d", sevenSegmentMatches("input.txt", true))
}
