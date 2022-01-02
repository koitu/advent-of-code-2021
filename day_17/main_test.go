package main

import (
	"testing"
)

func TestPart1(t *testing.T) {
	result := trickShot("test.txt", false)
	if result != 45 {
		t.Fatalf("part 1 sanity test did not pass, should have gotten 45 but got %d", result)
	}
	t.Logf("%d", trickShot("input.txt", false))
}

func TestPart2(t *testing.T) {
	result := trickShot("test.txt", true)
	if result != 112 {
		t.Fatalf("part 2 sanity test did not pass, should have gotten 112 but got %d", result)
	}
	t.Logf("%d", trickShot("input.txt", true))
}
