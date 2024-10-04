package main

import "testing"

func TestContainsNearbyAlmostDuplicate(t *testing.T) {
	tests := []struct {
		nums      []int
		indexDiff int
		valueDiff int
		want      bool
	}{
		{[]int{1, 2, 3, 1}, 3, 0, true},
		{[]int{1, 5, 9, 1, 5, 9}, 2, 3, false},
		{[]int{1, 0, 1, 1}, 1, 2, true},
		{[]int{1, 4, 9, 1, 4, 9}, 3, 1, true},
	}

	for _, tt := range tests {
		got := containsNearbyAlmostDuplicate(tt.nums, tt.indexDiff, tt.valueDiff)
		if got != tt.want {
			t.Errorf("containsNearbyAlmostDuplicate(%v, %d, %d) = %v; want %v",
				tt.nums, tt.indexDiff, tt.valueDiff, got, tt.want)
		}
	}
}
