package main

import (
	"fmt"
)

// 方法一：深度优先搜索 (DFS)
// 时间复杂度：O(n²)，空间复杂度：O(n)
func findCircleNumDFS(isConnected [][]int) int {
	n := len(isConnected)
	if n == 0 {
		return 0
	}

	// 访问数组，记录每个城市是否被访问过
	visited := make([]bool, n)
	count := 0

	// 从每个未访问的城市开始DFS
	for i := 0; i < n; i++ {
		if !visited[i] {
			dfs(isConnected, visited, i)
			count++
		}
	}

	return count
}

// DFS辅助函数
func dfs(isConnected [][]int, visited []bool, city int) {
	visited[city] = true
	
	// 遍历所有与当前城市相连的城市
	for nextCity := 0; nextCity < len(isConnected); nextCity++ {
		if isConnected[city][nextCity] == 1 && !visited[nextCity] {
			dfs(isConnected, visited, nextCity)
		}
	}
}

// 方法二：广度优先搜索 (BFS)
// 时间复杂度：O(n²)，空间复杂度：O(n)
func findCircleNumBFS(isConnected [][]int) int {
	n := len(isConnected)
	if n == 0 {
		return 0
	}

	visited := make([]bool, n)
	count := 0

	// 从每个未访问的城市开始BFS
	for i := 0; i < n; i++ {
		if !visited[i] {
			bfs(isConnected, visited, i)
			count++
		}
	}

	return count
}

// BFS辅助函数
func bfs(isConnected [][]int, visited []bool, startCity int) {
	queue := []int{startCity}
	visited[startCity] = true

	for len(queue) > 0 {
		city := queue[0]
		queue = queue[1:]

		// 遍历所有与当前城市相连的城市
		for nextCity := 0; nextCity < len(isConnected); nextCity++ {
			if isConnected[city][nextCity] == 1 && !visited[nextCity] {
				visited[nextCity] = true
				queue = append(queue, nextCity)
			}
		}
	}
}

// 方法三：并查集 (Union-Find)
// 时间复杂度：O(n² × α(n))，空间复杂度：O(n)
func findCircleNumUnionFind(isConnected [][]int) int {
	n := len(isConnected)
	if n == 0 {
		return 0
	}

	// 初始化并查集
	uf := NewUnionFind(n)

	// 遍历邻接矩阵，合并相连的城市
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ { // 利用对称性，只遍历上三角
			if isConnected[i][j] == 1 {
				uf.Union(i, j)
			}
		}
	}

	return uf.Count()
}

// 并查集结构
type UnionFind struct {
	parent []int
	rank   []int
	count  int
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
		count:  n,
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
	
	uf.count--
}

// 返回集合数量
func (uf *UnionFind) Count() int {
	return uf.count
}

// 方法四：优化的DFS（使用栈避免递归）
// 时间复杂度：O(n²)，空间复杂度：O(n)
func findCircleNumDFSIterative(isConnected [][]int) int {
	n := len(isConnected)
	if n == 0 {
		return 0
	}

	visited := make([]bool, n)
	count := 0

	for i := 0; i < n; i++ {
		if !visited[i] {
			dfsIterative(isConnected, visited, i)
			count++
		}
	}

	return count
}

// 迭代式DFS
func dfsIterative(isConnected [][]int, visited []bool, startCity int) {
	stack := []int{startCity}
	visited[startCity] = true

	for len(stack) > 0 {
		city := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		for nextCity := 0; nextCity < len(isConnected); nextCity++ {
			if isConnected[city][nextCity] == 1 && !visited[nextCity] {
				visited[nextCity] = true
				stack = append(stack, nextCity)
			}
		}
	}
}

// 测试函数
func main() {
	// 测试用例1：示例1
	isConnected1 := [][]int{
		{1, 1, 0},
		{1, 1, 0},
		{0, 0, 1},
	}
	fmt.Println("测试用例1:")
	fmt.Printf("输入: %v\n", isConnected1)
	fmt.Printf("DFS结果: %d\n", findCircleNumDFS(isConnected1))
	fmt.Printf("BFS结果: %d\n", findCircleNumBFS(isConnected1))
	fmt.Printf("并查集结果: %d\n", findCircleNumUnionFind(isConnected1))
	fmt.Printf("迭代DFS结果: %d\n", findCircleNumDFSIterative(isConnected1))
	fmt.Println("期望结果: 2")
	fmt.Println()

	// 测试用例2：示例2
	isConnected2 := [][]int{
		{1, 0, 0},
		{0, 1, 0},
		{0, 0, 1},
	}
	fmt.Println("测试用例2:")
	fmt.Printf("输入: %v\n", isConnected2)
	fmt.Printf("DFS结果: %d\n", findCircleNumDFS(isConnected2))
	fmt.Printf("BFS结果: %d\n", findCircleNumBFS(isConnected2))
	fmt.Printf("并查集结果: %d\n", findCircleNumUnionFind(isConnected2))
	fmt.Printf("迭代DFS结果: %d\n", findCircleNumDFSIterative(isConnected2))
	fmt.Println("期望结果: 3")
	fmt.Println()

	// 测试用例3：所有城市相连
	isConnected3 := [][]int{
		{1, 1, 1},
		{1, 1, 1},
		{1, 1, 1},
	}
	fmt.Println("测试用例3 (所有城市相连):")
	fmt.Printf("输入: %v\n", isConnected3)
	fmt.Printf("DFS结果: %d\n", findCircleNumDFS(isConnected3))
	fmt.Printf("BFS结果: %d\n", findCircleNumBFS(isConnected3))
	fmt.Printf("并查集结果: %d\n", findCircleNumUnionFind(isConnected3))
	fmt.Printf("迭代DFS结果: %d\n", findCircleNumDFSIterative(isConnected3))
	fmt.Println("期望结果: 1")
	fmt.Println()

	// 测试用例4：单个城市
	isConnected4 := [][]int{{1}}
	fmt.Println("测试用例4 (单个城市):")
	fmt.Printf("输入: %v\n", isConnected4)
	fmt.Printf("DFS结果: %d\n", findCircleNumDFS(isConnected4))
	fmt.Printf("BFS结果: %d\n", findCircleNumBFS(isConnected4))
	fmt.Printf("并查集结果: %d\n", findCircleNumUnionFind(isConnected4))
	fmt.Printf("迭代DFS结果: %d\n", findCircleNumDFSIterative(isConnected4))
	fmt.Println("期望结果: 1")
	fmt.Println()

	// 测试用例5：空矩阵
	var isConnected5 [][]int
	fmt.Println("测试用例5 (空矩阵):")
	fmt.Printf("输入: %v\n", isConnected5)
	fmt.Printf("DFS结果: %d\n", findCircleNumDFS(isConnected5))
	fmt.Printf("BFS结果: %d\n", findCircleNumBFS(isConnected5))
	fmt.Printf("并查集结果: %d\n", findCircleNumUnionFind(isConnected5))
	fmt.Printf("迭代DFS结果: %d\n", findCircleNumDFSIterative(isConnected5))
	fmt.Println("期望结果: 0")
	fmt.Println()

	// 测试用例6：复杂情况
	isConnected6 := [][]int{
		{1, 0, 0, 1},
		{0, 1, 1, 0},
		{0, 1, 1, 1},
		{1, 0, 1, 1},
	}
	fmt.Println("测试用例6 (复杂情况):")
	fmt.Printf("输入: %v\n", isConnected6)
	fmt.Printf("DFS结果: %d\n", findCircleNumDFS(isConnected6))
	fmt.Printf("BFS结果: %d\n", findCircleNumBFS(isConnected6))
	fmt.Printf("并查集结果: %d\n", findCircleNumUnionFind(isConnected6))
	fmt.Printf("迭代DFS结果: %d\n", findCircleNumDFSIterative(isConnected6))
	fmt.Println("期望结果: 1")
}
