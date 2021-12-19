package main

import (
	"testing"
)

func TestPart1(t *testing.T) {
	result := getSubPosition("test.txt", false)
	if result != 150 {
		t.Fatalf("part 1 sanity test did not pass, should have gotten 150 but got %d", result)
	}
	t.Logf("%d", getSubPosition("input.txt", false))
}

func TestPart2(t *testing.T) {
	result := getSubPosition("test.txt", true)
	if result != 900 {
		t.Fatalf("part 2 sanity test did not pass, should have gotten 900 but got %d", result)
	}
	t.Logf("%d", getSubPosition("input.txt", true))
}
