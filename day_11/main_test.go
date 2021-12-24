package main

import (
	"testing"
)

func TestPart1(t *testing.T) {
	result := dumboOctoSim("test.txt", false)
	if result != 1656 {
		t.Fatalf("part 1 sanity test did not pass, should have gotten 1656 but got %d", result)
	}
	t.Logf("%d", dumboOctoSim("input.txt", false))
}

func TestPart2(t *testing.T) {
	result := dumboOctoSim("test.txt", true)
	if result != 195 {
		t.Fatalf("part 2 sanity test did not pass, should have gotten 195 but got %d", result)
	}
	t.Logf("%d", dumboOctoSim("input.txt", true))
}
