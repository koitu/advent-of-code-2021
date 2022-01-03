package main

import (
	"testing"
)

func TestPart1(t *testing.T) {
	result := snailFish("test.txt", false)
	if result != 4140 {
		t.Fatalf("part 1 sanity test did not pass, should have gotten 4140 but got %d", result)
	}
	t.Logf("%d", snailFish("input.txt", false))
}

func TestPart2(t *testing.T) {
	result := snailFish("test.txt", true)
	if result != 3993 {
		t.Fatalf("part 2 sanity test did not pass, should have gotten 3993 but got %d", result)
	}
	t.Logf("%d", snailFish("input.txt", true))
}
