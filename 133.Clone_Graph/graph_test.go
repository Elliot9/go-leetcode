package main

import (
	"testing"
)

func TestCloneGraph(t *testing.T) {
	testCases := []struct {
		name     string
		graph    *Node
		expected *Node
	}{
		{
			name:     "Empty graph",
			graph:    nil,
			expected: nil,
		},
		{
			name:     "Single node",
			graph:    &Node{Val: 1},
			expected: &Node{Val: 1},
		},
		{
			name: "Two connected nodes",
			graph: func() *Node {
				node1 := &Node{Val: 1}
				node2 := &Node{Val: 2}
				node1.Neighbors = []*Node{node2}
				node2.Neighbors = []*Node{node1}
				return node1
			}(),
			expected: func() *Node {
				node1 := &Node{Val: 1}
				node2 := &Node{Val: 2}
				node1.Neighbors = []*Node{node2}
				node2.Neighbors = []*Node{node1}
				return node1
			}(),
		},
		{
			name: "Four connected nodes",
			graph: func() *Node {
				node1 := &Node{Val: 1}
				node2 := &Node{Val: 2}
				node3 := &Node{Val: 3}
				node4 := &Node{Val: 4}
				node1.Neighbors = []*Node{node2, node4}
				node2.Neighbors = []*Node{node1, node3}
				node3.Neighbors = []*Node{node2, node4}
				node4.Neighbors = []*Node{node1, node3}
				return node1
			}(),
			expected: func() *Node {
				node1 := &Node{Val: 1}
				node2 := &Node{Val: 2}
				node3 := &Node{Val: 3}
				node4 := &Node{Val: 4}
				node1.Neighbors = []*Node{node2, node4}
				node2.Neighbors = []*Node{node1, node3}
				node3.Neighbors = []*Node{node2, node4}
				node4.Neighbors = []*Node{node1, node3}
				return node1
			}(),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := cloneGraph(tc.graph)
			if !isGraphEqual(result, tc.expected) {
				t.Errorf("Expected %v, but got %v", tc.expected, result)
			}
		})
	}
}

func isGraphEqual(g1, g2 *Node) bool {
	if g1 == nil && g2 == nil {
		return true
	}
	if g1 == nil || g2 == nil {
		return false
	}
	if g1.Val != g2.Val {
		return false
	}
	if len(g1.Neighbors) != len(g2.Neighbors) {
		return false
	}

	visited := make(map[int]bool)
	return dfsCompare(g1, g2, visited)
}

func dfsCompare(n1, n2 *Node, visited map[int]bool) bool {
	if visited[n1.Val] {
		return true
	}
	visited[n1.Val] = true

	if len(n1.Neighbors) != len(n2.Neighbors) {
		return false
	}

	for i := range n1.Neighbors {
		if n1.Neighbors[i].Val != n2.Neighbors[i].Val {
			return false
		}
		if !dfsCompare(n1.Neighbors[i], n2.Neighbors[i], visited) {
			return false
		}
	}

	return true
}
