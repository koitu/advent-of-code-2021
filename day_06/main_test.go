package main

import (
	"testing"
)

func TestPart1(t *testing.T) {
	result := lanternFish("test.txt", 80)
	if result != 5934 {
		t.Fatalf("part 1 sanity test did not pass, should have gotten 5934 but got %d", result)
	}
	t.Logf("%d", lanternFish("input.txt", 80))
}

func TestPart2(t *testing.T) {
	result := lanternFish("test.txt", 256)
	if result != 26984457539 {
		t.Fatalf("part 2 sanity test did not pass, should have gotten 26984457539 but got %d", result)
	}
	t.Logf("%d", lanternFish("input.txt", 256))
}
