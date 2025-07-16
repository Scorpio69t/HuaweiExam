package main

import (
	"fmt"
)

// 解法1: 深度优先搜索 (DFS)
func numIslandsDFS(grid [][]byte) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	
	rows, cols := len(grid), len(grid[0])
	count := 0
	
	// DFS函数：将连通的陆地标记为已访问
	var dfs func(i, j int)
	dfs = func(i, j int) {
		// 边界检查和水域检查
		if i < 0 || i >= rows || j < 0 || j >= cols || grid[i][j] == '0' {
			return
		}
		
		// 标记当前位置为已访问
		grid[i][j] = '0'
		
		// 递归访问四个方向
		dfs(i-1, j) // 上
		dfs(i+1, j) // 下
		dfs(i, j-1) // 左
		dfs(i, j+1) // 右
	}
	
	// 遍历整个网格
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if grid[i][j] == '1' {
				count++      // 发现新岛屿
				dfs(i, j)    // 将整个岛屿标记为已访问
			}
		}
	}
	
	return count
}

// 解法2: 广度优先搜索 (BFS)
func numIslandsBFS(grid [][]byte) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	
	rows, cols := len(grid), len(grid[0])
	count := 0
	directions := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if grid[i][j] == '1' {
				count++
				
				// BFS队列
				queue := [][]int{{i, j}}
				grid[i][j] = '0' // 标记为已访问
				
				for len(queue) > 0 {
					cur := queue[0]
					queue = queue[1:]
					
					// 检查四个方向
					for _, dir := range directions {
						ni, nj := cur[0]+dir[0], cur[1]+dir[1]
						
						if ni >= 0 && ni < rows && nj >= 0 && nj < cols && grid[ni][nj] == '1' {
							grid[ni][nj] = '0' // 标记为已访问
							queue = append(queue, []int{ni, nj})
						}
					}
				}
			}
		}
	}
	
	return count
}

// 解法3: 并查集 (Union-Find)
type UnionFind struct {
	parent []int
	rank   []int
	count  int
}

func NewUnionFind(n int) *UnionFind {
	parent := make([]int, n)
	rank := make([]int, n)
	for i := 0; i < n; i++ {
		parent[i] = i
		rank[i] = 0
	}
	return &UnionFind{parent: parent, rank: rank, count: 0}
}

func (uf *UnionFind) find(x int) int {
	if uf.parent[x] != x {
		uf.parent[x] = uf.find(uf.parent[x]) // 路径压缩
	}
	return uf.parent[x]
}

func (uf *UnionFind) union(x, y int) {
	rootX, rootY := uf.find(x), uf.find(y)
	if rootX != rootY {
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
}

func numIslandsUnionFind(grid [][]byte) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	
	rows, cols := len(grid), len(grid[0])
	uf := NewUnionFind(rows * cols)
	directions := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	
	// 统计陆地数量并初始化并查集
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if grid[i][j] == '1' {
				uf.count++
			}
		}
	}
	
	// 合并相邻的陆地
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if grid[i][j] == '1' {
				for _, dir := range directions {
					ni, nj := i+dir[0], j+dir[1]
					if ni >= 0 && ni < rows && nj >= 0 && nj < cols && grid[ni][nj] == '1' {
						uf.union(i*cols+j, ni*cols+nj)
					}
				}
			}
		}
	}
	
	return uf.count
}

// 解法4: DFS（不修改原数组）
func numIslandsDFSNoModify(grid [][]byte) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	
	rows, cols := len(grid), len(grid[0])
	visited := make([][]bool, rows)
	for i := range visited {
		visited[i] = make([]bool, cols)
	}
	
	count := 0
	
	var dfs func(i, j int)
	dfs = func(i, j int) {
		if i < 0 || i >= rows || j < 0 || j >= cols || 
		   grid[i][j] == '0' || visited[i][j] {
			return
		}
		
		visited[i][j] = true
		dfs(i-1, j)
		dfs(i+1, j)
		dfs(i, j-1)
		dfs(i, j+1)
	}
	
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if grid[i][j] == '1' && !visited[i][j] {
				count++
				dfs(i, j)
			}
		}
	}
	
	return count
}

// 测试函数
func runTests() {
	fmt.Println("=== 岛屿数量算法测试 ===")
	
	// 测试用例1
	grid1 := [][]byte{
		{'1', '1', '1', '1', '0'},
		{'1', '1', '0', '1', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '0', '0', '0'},
	}
	
	// 测试用例2  
	grid2 := [][]byte{
		{'1', '1', '0', '0', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '1', '0', '0'},
		{'0', '0', '0', '1', '1'},
	}
	
	// 复制网格用于不同算法测试
	copyGrid := func(original [][]byte) [][]byte {
		if len(original) == 0 {
			return [][]byte{}
		}
		rows, cols := len(original), len(original[0])
		copy := make([][]byte, rows)
		for i := range copy {
			copy[i] = make([]byte, cols)
			for j := range copy[i] {
				copy[i][j] = original[i][j]
			}
		}
		return copy
	}
	
	fmt.Println("\n测试用例1:")
	printGrid(grid1)
	fmt.Printf("DFS解法: %d\n", numIslandsDFS(copyGrid(grid1)))
	fmt.Printf("BFS解法: %d\n", numIslandsBFS(copyGrid(grid1)))
	fmt.Printf("并查集解法: %d\n", numIslandsUnionFind(copyGrid(grid1)))
	fmt.Printf("DFS(不修改原数组): %d\n", numIslandsDFSNoModify(grid1))
	
	fmt.Println("\n测试用例2:")
	printGrid(grid2)
	fmt.Printf("DFS解法: %d\n", numIslandsDFS(copyGrid(grid2)))
	fmt.Printf("BFS解法: %d\n", numIslandsBFS(copyGrid(grid2)))
	fmt.Printf("并查集解法: %d\n", numIslandsUnionFind(copyGrid(grid2)))
	fmt.Printf("DFS(不修改原数组): %d\n", numIslandsDFSNoModify(grid2))
	
	// 边界测试
	fmt.Println("\n边界测试:")
	emptyGrid := [][]byte{}
	fmt.Printf("空网格: %d\n", numIslandsDFSNoModify(emptyGrid))
	
	singleLand := [][]byte{{'1'}}
	fmt.Printf("单个陆地: %d\n", numIslandsDFSNoModify(singleLand))
	
	singleWater := [][]byte{{'0'}}
	fmt.Printf("单个水域: %d\n", numIslandsDFSNoModify(singleWater))
	
	fmt.Println("\n所有测试完成!")
}

// 打印网格的辅助函数
func printGrid(grid [][]byte) {
	for _, row := range grid {
		for _, cell := range row {
			fmt.Printf("%c ", cell)
		}
		fmt.Println()
	}
}

func main() {
	runTests()
} 