package main

import (
	"testing"
)

func TestFlipEquiv(t *testing.T) {
	testCases := []struct {
		name     string
		root1    *TreeNode
		root2    *TreeNode
		expected bool
	}{
		{
			name:     "Example 1",
			root1:    buildTree(1, 2, 3, 4, 5, 6, nil, nil, nil, 7, 8),
			root2:    buildTree(1, 3, 2, nil, 6, 4, 5, nil, nil, nil, nil, 8, 7),
			expected: true,
		},
		{
			name:     "Example 2",
			root1:    buildTree(1, 2, 3, 4, 5, 6, nil, nil, nil, 7, 8),
			root2:    buildTree(1, 2, 3, 4, 5, 6, nil, nil, nil, 8, 7),
			expected: true,
		},
		{
			name:     "Both Null",
			root1:    nil,
			root2:    nil,
			expected: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := flipEquiv(tc.root1, tc.root2)
			if result != tc.expected {
				t.Errorf("flipEquiv(%v, %v) = %v; want %v", tc.root1, tc.root2, result, tc.expected)
			}
		})
	}
}

func buildTree(values ...interface{}) *TreeNode {
	if len(values) == 0 || values[0] == nil {
		return nil
	}

	root := &TreeNode{Val: values[0].(int)}
	queue := []*TreeNode{root}
	i := 1

	for len(queue) > 0 && i < len(values) {
		node := queue[0]
		queue = queue[1:]

		if i < len(values) && values[i] != nil {
			node.Left = &TreeNode{Val: values[i].(int)}
			queue = append(queue, node.Left)
		}
		i++

		if i < len(values) && values[i] != nil {
			node.Right = &TreeNode{Val: values[i].(int)}
			queue = append(queue, node.Right)
		}
		i++
	}

	return root
}
