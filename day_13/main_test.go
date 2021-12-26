package main

import (
	"testing"
)

func TestPart1(t *testing.T) {
	result := foldPaper("test.txt", false)
	if result != 17 {
		t.Fatalf("part 1 sanity test did not pass, should have gotten 17 but got %d", result)
	}
	t.Logf("%d", foldPaper("input.txt", false))
}

func TestPart2(t *testing.T) {
	t.Logf("%d", foldPaper("input.txt", true))
}
