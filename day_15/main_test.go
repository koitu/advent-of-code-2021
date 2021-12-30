package main

import (
	"testing"
)

func TestPart1(t *testing.T) {
	result := traverseCave("test.txt", false)
	if result != 40 {
		t.Fatalf("part 1 sanity test did not pass, should have gotten 40 but got %d", result)
	}
	t.Logf("%d", traverseCave("input.txt", false))
}

func TestPart2(t *testing.T) {
	result := traverseCave("test.txt", true)
	if result != 315 {
		t.Fatalf("part 2 sanity test did not pass, should have gotten 315 but got %d", result)
	}
	t.Logf("%d", traverseCave("input.txt", true))
}
