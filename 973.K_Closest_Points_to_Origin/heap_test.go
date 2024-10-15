package main

import (
	"sort"
	"testing"
)

func TestKClosest(t *testing.T) {
	testCases := []struct {
		name   string
		points [][]int
		k      int
		want   [][]int
	}{
		{
			name:   "Example 1",
			points: [][]int{{1, 3}, {-2, 2}},
			k:      1,
			want:   [][]int{{-2, 2}},
		},
		{
			name:   "Example 2",
			points: [][]int{{3, 3}, {5, -1}, {-2, 4}},
			k:      2,
			want:   [][]int{{3, 3}, {-2, 4}},
		},
		{
			name:   "Custom Example",
			points: [][]int{{-95, 76}, {17, 7}, {-55, -58}, {53, 20}, {-69, -8}, {-57, 87}, {-2, -42}, {-10, -87}, {-36, -57}, {97, -39}, {97, 49}},
			k:      5,
			want:   [][]int{{17, 7}, {-2, -42}, {53, 20}, {-36, -57}, {-69, -8}},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := kClosest(tc.points, tc.k)
			if !comparePoints(got, tc.want) {
				t.Errorf("kClosest(%v, %d) = %v; want %v", tc.points, tc.k, got, tc.want)
			}
		})
	}
}

// comparePoints 比較兩個點集合是否相等（忽略順序）
func comparePoints(a, b [][]int) bool {
	if len(a) != len(b) {
		return false
	}

	sortPoints := func(points [][]int) {
		sort.Slice(points, func(i, j int) bool {
			if points[i][0] != points[j][0] {
				return points[i][0] < points[j][0]
			}
			return points[i][1] < points[j][1]
		})
	}

	sortPoints(a)
	sortPoints(b)

	for i := range a {
		if !pointsEqual(a[i], b[i]) {
			return false
		}
	}

	return true
}

// pointsEqual 比較兩個點是否相等
func pointsEqual(a, b []int) bool {
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
