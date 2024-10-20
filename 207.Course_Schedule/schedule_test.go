package main

import (
	"testing"
)

func TestCanFinish(t *testing.T) {
	testCases := []struct {
		name          string
		numCourses    int
		prerequisites [][]int
		expected      bool
	}{
		{
			name:          "Example 1",
			numCourses:    2,
			prerequisites: [][]int{{1, 0}},
			expected:      true,
		},
		{
			name:          "Example 2",
			numCourses:    2,
			prerequisites: [][]int{{1, 0}, {0, 1}},
			expected:      false,
		},
		{
			name:          "No prerequisites",
			numCourses:    3,
			prerequisites: [][]int{},
			expected:      true,
		},
		{
			name:          "Complex case",
			numCourses:    4,
			prerequisites: [][]int{{1, 0}, {2, 1}, {3, 2}},
			expected:      true,
		},
		{
			name:          "Cycle case",
			numCourses:    3,
			prerequisites: [][]int{{0, 1}, {1, 2}, {2, 0}},
			expected:      false,
		},
		{
			name:          "Single course",
			numCourses:    1,
			prerequisites: [][]int{},
			expected:      true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := canFinish(tc.numCourses, tc.prerequisites)
			if result != tc.expected {
				t.Errorf("Expected %v, but got %v", tc.expected, result)
			}
		})
	}
}
