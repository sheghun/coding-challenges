package main

import (
	"fmt"
	"testing"
)

// Test case struct.
type testCase struct {
	input    []int
	expected int
}

// Helper function to compare two arrays for equality.
func assertEqual(t *testing.T, got, expected int) {
	if got != expected {
		t.Errorf("got %d, expected %d", got, expected)
	}
}

func TestLongestPeak(t *testing.T) {
	// Define test cases
	tests := []testCase{
		{input: []int{1, 2, 3, 4, 0, 10, 6, 5, -1, -3, 2, 3}, expected: 6},
		{input: []int{1, 3, 2}, expected: 3},
		{input: []int{1, 3, 2, 5, 4, 6, 7, 8, 3, 2, 1, 9, 8, 7}, expected: 7},
		{input: []int{1, 2, 3, 4, 5}, expected: 0},
		{input: []int{5, 4, 3, 2, 1}, expected: 0},
		{input: []int{1, 2, 3, 3, 2, 1}, expected: 0},
		{input: []int{1, 3, 2, 1, 3, 2, 1}, expected: 4},
		{input: []int{1, 2}, expected: 0},
		{input: []int{3}, expected: 0},
		{input: []int{2, 2, 2, 2, 2}, expected: 0},
		{input: []int{9, 8, 7, 6, 1, 2, 3, 4, 5}, expected: 0},
		{input: []int{1, 2, 3, 6, 7, 8, 9, 8, 7}, expected: 9},
		{input: []int{5, 10, 5, 2, 1, 4, 6, 3, 2, 0, 2, 1}, expected: 6},
		{input: []int{0, 0, 1, 2, 3, 2, 1, 0, 0}, expected: 7},
	}

	// Run tests
	for _, tc := range tests {
		t.Run(fmt.Sprintf("input=%v", tc.input), func(t *testing.T) {
			got := longestPeak(tc.input)
			assertEqual(t, got, tc.expected)
		})
	}
}
