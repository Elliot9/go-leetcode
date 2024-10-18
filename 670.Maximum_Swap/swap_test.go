package main

import (
	"testing"
)

func TestMaximumSwap(t *testing.T) {
	testCases := []struct {
		input    int
		expected int
	}{
		{2736, 7236},
		{9973, 9973},
		{1993, 9913},
		{98368, 98863},
		{1, 1},
		{10, 10},
		{100, 100},
		{12345, 52341},
		{99901, 99910},
	}

	for _, tc := range testCases {
		result := maximumSwap(tc.input)
		if result != tc.expected {
			t.Errorf("For input %d, expected %d, but got %d", tc.input, tc.expected, result)
		}
	}
}

func TestIntToArray(t *testing.T) {
	testCases := []struct {
		input    int
		expected []int
	}{
		{2736, []int{2, 7, 3, 6}},
		{9973, []int{9, 9, 7, 3}},
		{0, []int{0}},
		{1, []int{1}},
		{10, []int{1, 0}},
		{100, []int{1, 0, 0}},
	}

	for _, tc := range testCases {
		result := intToArray(tc.input)
		if !compareSlices(result, tc.expected) {
			t.Errorf("For input %d, expected %v, but got %v", tc.input, tc.expected, result)
		}
	}
}

func TestArrayToInt(t *testing.T) {
	testCases := []struct {
		input    []int
		expected int
	}{
		{[]int{2, 7, 3, 6}, 2736},
		{[]int{9, 9, 7, 3}, 9973},
		{[]int{0}, 0},
		{[]int{1}, 1},
		{[]int{1, 0}, 10},
		{[]int{1, 0, 0}, 100},
	}

	for _, tc := range testCases {
		result := arrayToInt(tc.input)
		if result != tc.expected {
			t.Errorf("For input %v, expected %d, but got %d", tc.input, tc.expected, result)
		}
	}
}

func compareSlices(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
