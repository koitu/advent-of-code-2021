package main

import (
	"testing"
)

func TestPart1(t *testing.T) {
	result := polymerization("test.txt", 10)
	if result != 1588 {
		t.Fatalf("part 1 sanity test did not pass, should have gotten 1588 but got %d", result)
	}
	t.Logf("%d", polymerization("input.txt", 10))
}

func TestPart2(t *testing.T) {
	result := polymerization("test.txt", 40)
	if result != 2188189693529 {
		t.Fatalf("part 2 sanity test did not pass, should have gotten 2188189693529 but got %d", result)
	}
	t.Logf("%d", polymerization("input.txt", 40))
}
