package main

import (
	"testing"
)

func TestPart1(t *testing.T) {
	result := passagePaths("test1.txt", false)
	if result != 10 {
		t.Fatalf("part 1 sanity test 1 did not pass, should have gotten 10 but got %d", result)
	}
	result = passagePaths("test2.txt", false)
	if result != 19 {
		t.Fatalf("part 1 sanity test 2 did not pass, should have gotten 19 but got %d", result)
	}
	result = passagePaths("test3.txt", false)
	if result != 226 {
		t.Fatalf("part 1 sanity test 3 did not pass, should have gotten 226 but got %d", result)
	}
	t.Logf("%d", passagePaths("input.txt", false))
}

func TestPart2(t *testing.T) {
	result := passagePaths("test1.txt", true)
	if result != 36 {
		t.Fatalf("part 2 sanity test 1 did not pass, should have gotten 36 but got %d", result)
	}
	result = passagePaths("test2.txt", true)
	if result != 103 {
		t.Fatalf("part 2 sanity test 2 did not pass, should have gotten 103 but got %d", result)
	}
	result = passagePaths("test3.txt", true)
	if result != 3509 {
		t.Fatalf("part 2 sanity test 3 did not pass, should have gotten 3509 but got %d", result)
	}
	t.Logf("%d", passagePaths("input.txt", true))
}
