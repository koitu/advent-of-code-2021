package main

import (
	"testing"
)

func TestPart1(t *testing.T) {
	result := beaconScanner("test.txt", false)
	if result != 79 {
		t.Fatalf("part 1 sanity test did not pass, should have gotten 79 but got %d", result)
	}
	// 10 secs to get 396 jeez
	t.Logf("%d", beaconScanner("input.txt", false))
}

func TestPart2(t *testing.T) {
	result := beaconScanner("test.txt", true)
	if result != 3621 {
		t.Fatalf("part 2 sanity test did not pass, should have gotten 3621 but got %d", result)
	}
	// 12 secs to get 11828
	t.Logf("%d", beaconScanner("input.txt", true))
}
