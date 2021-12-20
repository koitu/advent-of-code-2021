package main

import (
	"testing"
)

func TestPart1(t *testing.T) {
	result := bingoSubsystem("test.txt", false)
	if result != 4512 {
		t.Fatalf("part 1 sanity test did not pass, should have gotten 4512 but got %d", result)
	}
	t.Logf("%d", bingoSubsystem("input.txt", false))
}

func TestPart2(t *testing.T) {
	result := bingoSubsystem("test.txt", true)
	if result != 1924 {
		t.Fatalf("part 2 sanity test did not pass, should have gotten 1924 but got %d", result)
	}
	t.Logf("%d", bingoSubsystem("input.txt", true))
}
