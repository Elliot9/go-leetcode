package main

import (
	"sort"
	"testing"
)

func TestRemoveSubfolders(t *testing.T) {
	testCases := []struct {
		name     string
		folder   []string
		expected []string
	}{
		{
			name:     "Example 1",
			folder:   []string{"/a", "/a/b", "/c/d", "/c/d/e", "/c/f"},
			expected: []string{"/a", "/c/d", "/c/f"},
		},
		{
			name:     "Example 2",
			folder:   []string{"/a", "/a/b/c", "/a/b/d"},
			expected: []string{"/a"},
		},
		{
			name:     "Example 3",
			folder:   []string{"/a/b/c", "/a/b/ca", "/a/b/d"},
			expected: []string{"/a/b/c", "/a/b/ca", "/a/b/d"},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := removeSubfolders(tc.folder)
			if !compareFolders(result, tc.expected) {
				t.Errorf("removeSubfolders(%v) = %v; want %v", tc.folder, result, tc.expected)
			}
		})
	}
}

func compareFolders(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	sort.Strings(a)
	sort.Strings(b)
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
