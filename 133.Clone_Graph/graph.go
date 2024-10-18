package main

type Node struct {
	Val       int
	Neighbors []*Node
}

func cloneGraph(node *Node) *Node {

	// 記錄每個節點的拷貝，避免重複
	visited := make(map[int]*Node)

	var dfs func(*Node, map[int]*Node) *Node
	dfs = func(n *Node, m map[int]*Node) *Node {
		if n == nil {
			return nil
		}

		// 如果這個節點已經被拷貝，返回其拷貝
		if node, exists := m[n.Val]; exists {
			return node
		}

		// 創建新的節點拷貝
		newNode := &Node{
			Val: n.Val,
		}
		// 保存拷貝到映射中
		visited[n.Val] = newNode

		// 複製所有鄰居節點
		for _, neighbor := range n.Neighbors {
			newNode.Neighbors = append(newNode.Neighbors, dfs(neighbor, visited))
		}

		// 拷貝圖的過程中，你的目標是建立一個完整的深拷貝，而且每個節點只應該被拷貝一次
		// 不需要刪除的原因是 保證每個節點只會被複製一次
		// 如果刪除了 系統會重新進行一次拷貝操作，即重新創建一個 Node 給它，這樣雖然它的 Val 相同，但是記憶體位置不同，因此這兩個 Node 被視為不同的物件
		// delete(visted, n.Val)

		return newNode
	}

	return dfs(node, visited)
}
