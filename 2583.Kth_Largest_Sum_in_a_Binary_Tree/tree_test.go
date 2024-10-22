package main

import (
	"testing"
)

func TestKthLargestLevelSum(t *testing.T) {
	testCases := []struct {
		name     string
		root     *TreeNode
		k        int
		expected int64
	}{
		{
			name: "Example 1",
			root: &TreeNode{
				Val: 5,
				Left: &TreeNode{
					Val: 8,
					Left: &TreeNode{
						Val: 2,
					},
					Right: &TreeNode{
						Val: 1,
					},
				},
				Right: &TreeNode{
					Val: 9,
					Left: &TreeNode{
						Val: 3,
					},
					Right: &TreeNode{
						Val: 7,
					},
				},
			},
			k:        2,
			expected: 13,
		},
		{
			name: "Example 2",
			root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 3,
					},
				},
			},
			k:        1,
			expected: 3,
		},
		{
			name: "K larger than number of levels",
			root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
				},
				Right: &TreeNode{
					Val: 3,
				},
			},
			k:        3,
			expected: -1,
		},
		{
			name: "New test case",
			root: &TreeNode{
				Val: 806819,
				Left: &TreeNode{
					Val: 136611,
					Left: &TreeNode{
						Val: 985,
					},
					Right: &TreeNode{
						Val: 637251,
					},
				},
				Right: &TreeNode{
					Val: 911588,
					Left: &TreeNode{
						Val: 243509,
					},
					Right: &TreeNode{
						Val: 678170,
						Left: &TreeNode{
							Val: 31783,
						},
					},
				},
			},
			k:        2,
			expected: 1048199,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := kthLargestLevelSum(tc.root, tc.k)
			if result != tc.expected {
				t.Errorf("Expected %d, but got %d", tc.expected, result)
			}
		})
	}
}
