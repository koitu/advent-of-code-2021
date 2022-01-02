package main

import (
	"os"
	"testing"
)

func TestPart1(t *testing.T) {
	b, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	tests := []struct {
		name   string
		input  string
		expect int
	}{
		{
			name:   "test_1",
			input:  "8A004A801A8002F478",
			expect: 16,
		},
		{
			name:   "test_2",
			input:  "620080001611562C8802118E34",
			expect: 12,
		},
		{
			name:   "test_3",
			input:  "C0015000016115A2E0802F182340",
			expect: 23,
		},
		{
			name:   "test_4",
			input:  "A0016C880162017C3686B18A3D4780",
			expect: 31,
		},
		{
			name:   "input",
			input:  string(b),
			expect: 821,
		},
	}

	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if val := packetDecode(tc.input, false); val != tc.expect {
				t.Errorf("part 1 sanity test %d did not pass, expected %d but got %d", i+1, tc.expect, val)
			}
		})
	}
}

func TestPart2(t *testing.T) {
	b, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	tests := []struct {
		name   string
		input  string
		expect int
	}{
		{
			name:   "test_1",
			input:  "C200B40A82",
			expect: 3,
		},
		{
			name:   "test_2",
			input:  "04005AC33890",
			expect: 54,
		},
		{
			name:   "test_3",
			input:  "880086C3E88112",
			expect: 7,
		},
		{
			name:   "test_4",
			input:  "CE00C43D881120",
			expect: 9,
		},
		{
			name:   "test_5",
			input:  "D8005AC2A8F0",
			expect: 1,
		},
		{
			name:   "test_6",
			input:  "F600BC2D8F",
			expect: 0,
		},
		{
			name:   "test_7",
			input:  "9C005AC2F8F0",
			expect: 0,
		},
		{
			name:   "test_8",
			input:  "9C0141080250320F1802104A08",
			expect: 1,
		},
		{
			name:   "input",
			input:  string(b),
			expect: 2056021084691,
		},
	}
	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if val := packetDecode(tc.input, true); val != tc.expect {
				t.Fatalf("part 2 sanity test %d did not pass, expected %d but got %d", i+1, tc.expect, val)
			}
		})
	}
}
