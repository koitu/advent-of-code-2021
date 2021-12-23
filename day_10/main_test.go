package main

import (
	"testing"
)

func TestPart1(t *testing.T) {
	result := scoreSyntax("test.txt", false)
	if result != 26397 {
		t.Fatalf("part 1 sanity test did not pass, should have gotten 26397 but got %d", result)
	}
	t.Logf("%d", scoreSyntax("input.txt", false))
}

func TestPart2(t *testing.T) {
	result := scoreSyntax("test.txt", true)
	if result != 288957 {
		t.Fatalf("part 2 sanity test did not pass, should have gotten 288957 but got %d", result)
	}
	t.Logf("%d", scoreSyntax("input.txt", true))
}
