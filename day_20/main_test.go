package main

import (
	"testing"
)

func TestPart1(t *testing.T) {
	result := mapEnhance("test.txt", 2)
	if result != 40 {
		t.Fatalf("part 1 sanity test did not pass, should have gotten 40 but got %d", result)
	}
	t.Logf("%d", mapEnhance("input.txt", 2))
}

func TestPart2(t *testing.T) {
	result := mapEnhance("test.txt", 2)
	if result != 315 {
		t.Fatalf("part 2 sanity test did not pass, should have gotten 315 but got %d", result)
	}
	t.Logf("%d", mapEnhance("input.txt", 2))
}
