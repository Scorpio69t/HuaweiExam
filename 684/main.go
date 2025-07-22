package main

import (
	"fmt"
)

// 方法一：并查集 (Union-Find) - 推荐解法
// 时间复杂度：O(n × α(n))，空间复杂度：O(n)
func findRedundantConnection(edges [][]int) []int {
	n := len(edges)
	uf := NewUnionFind(n + 1) // 节点编号从1开始
	
	for _, edge := range edges {
		u, v := edge[0], edge[1]
		if uf.Find(u) == uf.Find(v) {
			// 这条边会形成环，返回这条边
			return edge
		}
		uf.Union(u, v)
	}
	
	return nil
}

// 并查集结构
type UnionFind struct {
	parent []int
	rank   []int
}

// 创建新的并查集
func NewUnionFind(n int) *UnionFind {
	parent := make([]int, n)
	rank := make([]int, n)
	
	for i := 0; i < n; i++ {
		parent[i] = i
		rank[i] = 1
	}
	
	return &UnionFind{
		parent: parent,
		rank:   rank,
	}
}

// 查找根节点（路径压缩）
func (uf *UnionFind) Find(x int) int {
	if uf.parent[x] != x {
		uf.parent[x] = uf.Find(uf.parent[x]) // 路径压缩
	}
	return uf.parent[x]
}

// 合并两个集合（按秩合并）
func (uf *UnionFind) Union(x, y int) {
	rootX := uf.Find(x)
	rootY := uf.Find(y)
	
	if rootX == rootY {
		return
	}
	
	// 按秩合并
	if uf.rank[rootX] < uf.rank[rootY] {
		uf.parent[rootX] = rootY
	} else if uf.rank[rootX] > uf.rank[rootY] {
		uf.parent[rootY] = rootX
	} else {
		uf.parent[rootY] = rootX
		uf.rank[rootX]++
	}
}

// 方法二：深度优先搜索 (DFS)
// 时间复杂度：O(n²)，空间复杂度：O(n)
func findRedundantConnectionDFS(edges [][]int) []int {
	n := len(edges)
	
	// 从最后一条边开始尝试删除
	for i := n - 1; i >= 0; i-- {
		// 构建删除边i后的图
		graph := make(map[int][]int)
		for j := 0; j < n; j++ {
			if j != i {
				u, v := edges[j][0], edges[j][1]
				graph[u] = append(graph[u], v)
				graph[v] = append(graph[v], u)
			}
		}
		
		// 检查是否有环
		if !hasCycle(graph, n) {
			return edges[i]
		}
	}
	
	return nil
}

// 检查图中是否有环
func hasCycle(graph map[int][]int, n int) bool {
	visited := make([]bool, n+1)
	
	for i := 1; i <= n; i++ {
		if !visited[i] {
			if dfsHasCycle(graph, visited, i, -1) {
				return true
			}
		}
	}
	
	return false
}

// DFS检测环
func dfsHasCycle(graph map[int][]int, visited []bool, node, parent int) bool {
	visited[node] = true
	
	for _, neighbor := range graph[node] {
		if !visited[neighbor] {
			if dfsHasCycle(graph, visited, neighbor, node) {
				return true
			}
		} else if neighbor != parent {
			// 访问到已访问的节点且不是父节点，说明有环
			return true
		}
	}
	
	return false
}

// 方法三：优化的并查集（简化版）
// 时间复杂度：O(n × α(n))，空间复杂度：O(n)
func findRedundantConnectionOptimized(edges [][]int) []int {
	n := len(edges)
	parent := make([]int, n+1)
	
	// 初始化并查集
	for i := 1; i <= n; i++ {
		parent[i] = i
	}
	
	// 查找函数（带路径压缩）
	var find func(x int) int
	find = func(x int) int {
		if parent[x] != x {
			parent[x] = find(parent[x])
		}
		return parent[x]
	}
	
	// 合并函数
	union := func(x, y int) bool {
		rootX := find(x)
		rootY := find(y)
		if rootX == rootY {
			return false // 已经在同一集合中
		}
		parent[rootX] = rootY
		return true
	}
	
	for _, edge := range edges {
		u, v := edge[0], edge[1]
		if !union(u, v) {
			// 无法合并，说明会形成环
			return edge
		}
	}
	
	return nil
}

// 方法四：广度优先搜索 (BFS)
// 时间复杂度：O(n²)，空间复杂度：O(n)
func findRedundantConnectionBFS(edges [][]int) []int {
	n := len(edges)
	
	// 从最后一条边开始尝试删除
	for i := n - 1; i >= 0; i-- {
		// 构建删除边i后的图
		graph := make(map[int][]int)
		for j := 0; j < n; j++ {
			if j != i {
				u, v := edges[j][0], edges[j][1]
				graph[u] = append(graph[u], v)
				graph[v] = append(graph[v], u)
			}
		}
		
		// 检查是否有环
		if !hasCycleBFS(graph, n) {
			return edges[i]
		}
	}
	
	return nil
}

// BFS检测环
func hasCycleBFS(graph map[int][]int, n int) bool {
	visited := make([]bool, n+1)
	
	for i := 1; i <= n; i++ {
		if !visited[i] {
			if bfsHasCycle(graph, visited, i) {
				return true
			}
		}
	}
	
	return false
}

// BFS检测环的具体实现
func bfsHasCycle(graph map[int][]int, visited []bool, start int) bool {
	queue := [][]int{{start, -1}} // [节点, 父节点]
	visited[start] = true
	
	for len(queue) > 0 {
		node, parent := queue[0][0], queue[0][1]
		queue = queue[1:]
		
		for _, neighbor := range graph[node] {
			if !visited[neighbor] {
				visited[neighbor] = true
				queue = append(queue, []int{neighbor, node})
			} else if neighbor != parent {
				// 访问到已访问的节点且不是父节点，说明有环
				return true
			}
		}
	}
	
	return false
}

// 测试函数
func main() {
	// 测试用例1：示例1
	edges1 := [][]int{{1, 2}, {1, 3}, {2, 3}}
	fmt.Println("测试用例1:")
	fmt.Printf("输入: %v\n", edges1)
	fmt.Printf("并查集结果: %v\n", findRedundantConnection(edges1))
	fmt.Printf("DFS结果: %v\n", findRedundantConnectionDFS(edges1))
	fmt.Printf("优化并查集结果: %v\n", findRedundantConnectionOptimized(edges1))
	fmt.Printf("BFS结果: %v\n", findRedundantConnectionBFS(edges1))
	fmt.Println("期望结果: [2 3]")
	fmt.Println()

	// 测试用例2：示例2
	edges2 := [][]int{{1, 2}, {2, 3}, {3, 4}, {1, 4}, {1, 5}}
	fmt.Println("测试用例2:")
	fmt.Printf("输入: %v\n", edges2)
	fmt.Printf("并查集结果: %v\n", findRedundantConnection(edges2))
	fmt.Printf("DFS结果: %v\n", findRedundantConnectionDFS(edges2))
	fmt.Printf("优化并查集结果: %v\n", findRedundantConnectionOptimized(edges2))
	fmt.Printf("BFS结果: %v\n", findRedundantConnectionBFS(edges2))
	fmt.Println("期望结果: [1 4]")
	fmt.Println()

	// 测试用例3：边界情况 - 三角形环
	edges3 := [][]int{{1, 2}, {2, 3}, {3, 1}}
	fmt.Println("测试用例3 (三角形环):")
	fmt.Printf("输入: %v\n", edges3)
	fmt.Printf("并查集结果: %v\n", findRedundantConnection(edges3))
	fmt.Printf("DFS结果: %v\n", findRedundantConnectionDFS(edges3))
	fmt.Printf("优化并查集结果: %v\n", findRedundantConnectionOptimized(edges3))
	fmt.Printf("BFS结果: %v\n", findRedundantConnectionBFS(edges3))
	fmt.Println("期望结果: [3 1]")
	fmt.Println()

	// 测试用例4：复杂情况 - 大环
	edges4 := [][]int{{1, 2}, {2, 3}, {3, 4}, {4, 5}, {5, 6}, {6, 1}}
	fmt.Println("测试用例4 (大环):")
	fmt.Printf("输入: %v\n", edges4)
	fmt.Printf("并查集结果: %v\n", findRedundantConnection(edges4))
	fmt.Printf("DFS结果: %v\n", findRedundantConnectionDFS(edges4))
	fmt.Printf("优化并查集结果: %v\n", findRedundantConnectionOptimized(edges4))
	fmt.Printf("BFS结果: %v\n", findRedundantConnectionBFS(edges4))
	fmt.Println("期望结果: [6 1]")
	fmt.Println()

	// 测试用例5：多条冗余边的情况
	edges5 := [][]int{{1, 2}, {2, 3}, {3, 4}, {4, 1}, {1, 3}}
	fmt.Println("测试用例5 (多条冗余边):")
	fmt.Printf("输入: %v\n", edges5)
	fmt.Printf("并查集结果: %v\n", findRedundantConnection(edges5))
	fmt.Printf("DFS结果: %v\n", findRedundantConnectionDFS(edges5))
	fmt.Printf("优化并查集结果: %v\n", findRedundantConnectionOptimized(edges5))
	fmt.Printf("BFS结果: %v\n", findRedundantConnectionBFS(edges5))
	fmt.Println("期望结果: [1 3] (返回最后出现的冗余边)")
	fmt.Println()

	// 测试用例6：最小情况
	edges6 := [][]int{{1, 2}, {2, 3}, {3, 1}}
	fmt.Println("测试用例6 (最小情况):")
	fmt.Printf("输入: %v\n", edges6)
	fmt.Printf("并查集结果: %v\n", findRedundantConnection(edges6))
	fmt.Printf("DFS结果: %v\n", findRedundantConnectionDFS(edges6))
	fmt.Printf("优化并查集结果: %v\n", findRedundantConnectionOptimized(edges6))
	fmt.Printf("BFS结果: %v\n", findRedundantConnectionBFS(edges6))
	fmt.Println("期望结果: [3 1]")
}
