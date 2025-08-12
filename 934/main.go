package main

import (
	"fmt"
	"math"
	"strings"
	"time"
)

// 方向数组：上下左右
var directions = [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

// ========== 方法1：经典DFS+BFS解法 ==========
func shortestBridge1(grid [][]int) int {
	n := len(grid)

	// 第一步：找到并标记第一个岛屿
	found := false
	for i := 0; i < n && !found; i++ {
		for j := 0; j < n && !found; j++ {
			if grid[i][j] == 1 {
				markIsland(grid, i, j, 2) // 标记为2
				found = true
			}
		}
	}

	// 第二步：收集第一个岛屿的边界点
	boundary := collectBoundary(grid, 2)

	// 第三步：多源BFS寻找最短路径
	return multiSourceBFS(grid, boundary)
}

// DFS标记岛屿
func markIsland(grid [][]int, i, j, islandId int) {
	if i < 0 || i >= len(grid) || j < 0 || j >= len(grid[0]) || grid[i][j] != 1 {
		return
	}
	grid[i][j] = islandId
	for _, dir := range directions {
		markIsland(grid, i+dir[0], j+dir[1], islandId)
	}
}

// 收集岛屿边界点(与水相邻的陆地)
func collectBoundary(grid [][]int, islandId int) [][]int {
	var boundary [][]int
	n := len(grid)

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == islandId {
				// 检查是否为边界点
				for _, dir := range directions {
					ni, nj := i+dir[0], j+dir[1]
					if ni >= 0 && ni < n && nj >= 0 && nj < n && grid[ni][nj] == 0 {
						boundary = append(boundary, []int{i, j})
						break
					}
				}
			}
		}
	}
	return boundary
}

// 多源BFS搜索
func multiSourceBFS(grid [][]int, boundary [][]int) int {
	n := len(grid)
	queue := make([][]int, 0)

	// 初始化队列和距离数组
	dist := make([][]int, n)
	for i := range dist {
		dist[i] = make([]int, n)
		for j := range dist[i] {
			dist[i][j] = -1
		}
	}

	// 将边界点加入队列
	for _, point := range boundary {
		queue = append(queue, []int{point[0], point[1], 0})
		dist[point[0]][point[1]] = 0
	}

	// BFS搜索
	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]
		x, y, d := curr[0], curr[1], curr[2]

		for _, dir := range directions {
			nx, ny := x+dir[0], y+dir[1]

			if nx >= 0 && nx < n && ny >= 0 && ny < n {
				if grid[nx][ny] == 1 { // 找到第二个岛屿
					return d
				}
				if grid[nx][ny] == 0 && dist[nx][ny] == -1 { // 未访问的水域
					dist[nx][ny] = d + 1
					queue = append(queue, []int{nx, ny, d + 1})
				}
			}
		}
	}

	return -1
}

// ========== 方法2：双向BFS优化解法 ==========
func shortestBridge2(grid [][]int) int {
	n := len(grid)

	// 标记两个岛屿
	islandId := 2
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				markIslandBidirectional(grid, i, j, islandId)
				islandId++
				if islandId > 3 { // 已找到两个岛屿
					break
				}
			}
		}
		if islandId > 3 {
			break
		}
	}

	// 收集两个岛屿的边界点
	boundary1 := collectBoundary(grid, 2)
	boundary2 := collectBoundary(grid, 3)

	// 双向BFS
	return bidirectionalBFS(grid, boundary1, boundary2)
}

// 双向BFS标记岛屿
func markIslandBidirectional(grid [][]int, i, j, islandId int) {
	if i < 0 || i >= len(grid) || j < 0 || j >= len(grid[0]) || grid[i][j] != 1 {
		return
	}
	grid[i][j] = islandId
	for _, dir := range directions {
		markIslandBidirectional(grid, i+dir[0], j+dir[1], islandId)
	}
}

// 双向BFS搜索
func bidirectionalBFS(grid [][]int, boundary1, boundary2 [][]int) int {
	n := len(grid)

	// 初始化两个距离数组
	dist1 := make([][]int, n)
	dist2 := make([][]int, n)
	for i := range dist1 {
		dist1[i] = make([]int, n)
		dist2[i] = make([]int, n)
		for j := range dist1[i] {
			dist1[i][j] = -1
			dist2[i][j] = -1
		}
	}

	// 初始化两个队列
	queue1 := make([][]int, 0)
	queue2 := make([][]int, 0)

	for _, point := range boundary1 {
		queue1 = append(queue1, []int{point[0], point[1], 0})
		dist1[point[0]][point[1]] = 0
	}

	for _, point := range boundary2 {
		queue2 = append(queue2, []int{point[0], point[1], 0})
		dist2[point[0]][point[1]] = 0
	}

	// 交替扩展两个队列
	for len(queue1) > 0 || len(queue2) > 0 {
		// 选择较小的队列进行扩展
		if len(queue2) == 0 || (len(queue1) > 0 && len(queue1) <= len(queue2)) {
			result := expandQueue(grid, &queue1, dist1, dist2)
			if result != -1 {
				return result
			}
		} else {
			result := expandQueue(grid, &queue2, dist2, dist1)
			if result != -1 {
				return result
			}
		}
	}

	return -1
}

// 扩展队列
func expandQueue(grid [][]int, queue *[][]int, myDist, otherDist [][]int) int {
	if len(*queue) == 0 {
		return -1
	}

	nextQueue := make([][]int, 0)

	for _, curr := range *queue {
		x, y, d := curr[0], curr[1], curr[2]

		for _, dir := range directions {
			nx, ny := x+dir[0], y+dir[1]

			if nx >= 0 && nx < len(grid) && ny >= 0 && ny < len(grid[0]) {
				if otherDist[nx][ny] != -1 { // 遇到另一侧的搜索
					return d + otherDist[nx][ny]
				}
				if grid[nx][ny] == 0 && myDist[nx][ny] == -1 { // 未访问的水域
					myDist[nx][ny] = d + 1
					nextQueue = append(nextQueue, []int{nx, ny, d + 1})
				}
			}
		}
	}

	*queue = nextQueue
	return -1
}

// ========== 方法3：A*搜索算法 ==========
type AStarNode struct {
	x, y  int
	gCost int // 从起点到当前点的实际距离
	hCost int // 从当前点到终点的启发式距离
	fCost int // gCost + hCost
}

func shortestBridge3(grid [][]int) int {
	// 找到两个岛屿
	islands := findTwoIslands(grid)
	if len(islands) != 2 {
		return -1
	}

	// 使用A*算法搜索
	return aStarSearch(grid, islands[0], islands[1])
}

// 找到两个岛屿的所有点
func findTwoIslands(grid [][]int) [][][]int {
	n := len(grid)
	islands := make([][][]int, 0)
	visited := make([][]bool, n)
	for i := range visited {
		visited[i] = make([]bool, n)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 && !visited[i][j] {
				island := make([][]int, 0)
				dfsCollectIsland(grid, visited, i, j, &island)
				islands = append(islands, island)
			}
		}
	}

	return islands
}

// DFS收集岛屿所有点
func dfsCollectIsland(grid [][]int, visited [][]bool, i, j int, island *[][]int) {
	if i < 0 || i >= len(grid) || j < 0 || j >= len(grid[0]) ||
		visited[i][j] || grid[i][j] == 0 {
		return
	}

	visited[i][j] = true
	*island = append(*island, []int{i, j})

	for _, dir := range directions {
		dfsCollectIsland(grid, visited, i+dir[0], j+dir[1], island)
	}
}

// A*搜索算法
func aStarSearch(grid [][]int, island1, island2 [][]int) int {
	n := len(grid)

	// 计算两个岛屿之间的最小曼哈顿距离作为启发
	minDist := math.MaxInt32
	for _, p1 := range island1 {
		for _, p2 := range island2 {
			dist := abs(p1[0]-p2[0]) + abs(p1[1]-p2[1])
			if dist < minDist {
				minDist = dist
			}
		}
	}

	// 从第一个岛屿的边界开始A*搜索
	boundary1 := filterBoundaryPoints(grid, island1)
	target := make(map[string]bool)
	for _, p := range island2 {
		target[fmt.Sprintf("%d,%d", p[0], p[1])] = true
	}

	// A*主搜索循环
	dist := make([][]int, n)
	for i := range dist {
		dist[i] = make([]int, n)
		for j := range dist[i] {
			dist[i][j] = -1
		}
	}

	openSet := make([]*AStarNode, 0)
	for _, p := range boundary1 {
		node := &AStarNode{
			x: p[0], y: p[1],
			gCost: 0,
			hCost: manhattanDistance(p[0], p[1], island2),
		}
		node.fCost = node.gCost + node.hCost
		openSet = append(openSet, node)
		dist[p[0]][p[1]] = 0
	}

	for len(openSet) > 0 {
		// 找到fCost最小的节点
		minIdx := 0
		for i := 1; i < len(openSet); i++ {
			if openSet[i].fCost < openSet[minIdx].fCost {
				minIdx = i
			}
		}

		current := openSet[minIdx]
		openSet = append(openSet[:minIdx], openSet[minIdx+1:]...)

		// 检查是否到达目标
		key := fmt.Sprintf("%d,%d", current.x, current.y)
		if target[key] {
			return current.gCost - 1 // -1因为不计算到达岛屿的步数
		}

		// 扩展邻居节点
		for _, dir := range directions {
			nx, ny := current.x+dir[0], current.y+dir[1]

			if nx >= 0 && nx < n && ny >= 0 && ny < n {
				newGCost := current.gCost + 1

				if target[fmt.Sprintf("%d,%d", nx, ny)] {
					return newGCost - 1
				}

				if grid[nx][ny] == 0 && (dist[nx][ny] == -1 || newGCost < dist[nx][ny]) {
					dist[nx][ny] = newGCost
					node := &AStarNode{
						x: nx, y: ny,
						gCost: newGCost,
						hCost: manhattanDistance(nx, ny, island2),
					}
					node.fCost = node.gCost + node.hCost
					openSet = append(openSet, node)
				}
			}
		}
	}

	return -1
}

// 过滤边界点
func filterBoundaryPoints(grid [][]int, island [][]int) [][]int {
	n := len(grid)
	boundary := make([][]int, 0)

	for _, point := range island {
		i, j := point[0], point[1]
		isBoundary := false

		for _, dir := range directions {
			ni, nj := i+dir[0], j+dir[1]
			if ni >= 0 && ni < n && nj >= 0 && nj < n && grid[ni][nj] == 0 {
				isBoundary = true
				break
			}
		}

		if isBoundary {
			boundary = append(boundary, point)
		}
	}

	return boundary
}

// 计算到岛屿的曼哈顿距离
func manhattanDistance(x, y int, island [][]int) int {
	minDist := math.MaxInt32
	for _, point := range island {
		dist := abs(x-point[0]) + abs(y-point[1])
		if dist < minDist {
			minDist = dist
		}
	}
	return minDist
}

// ========== 方法4：并查集+BFS混合解法 ==========
type UnionFind struct {
	parent []int
	rank   []int
}

func NewUnionFind(n int) *UnionFind {
	parent := make([]int, n)
	rank := make([]int, n)
	for i := range parent {
		parent[i] = i
	}
	return &UnionFind{parent, rank}
}

func (uf *UnionFind) Find(x int) int {
	if uf.parent[x] != x {
		uf.parent[x] = uf.Find(uf.parent[x])
	}
	return uf.parent[x]
}

func (uf *UnionFind) Union(x, y int) {
	px, py := uf.Find(x), uf.Find(y)
	if px == py {
		return
	}

	if uf.rank[px] < uf.rank[py] {
		uf.parent[px] = py
	} else if uf.rank[px] > uf.rank[py] {
		uf.parent[py] = px
	} else {
		uf.parent[py] = px
		uf.rank[px]++
	}
}

func shortestBridge4(grid [][]int) int {
	n := len(grid)

	// 使用并查集构建岛屿连通性
	uf := NewUnionFind(n * n)
	islands := make([][]int, 0)

	// 建立岛屿内部连通性
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				islands = append(islands, []int{i, j})
				for _, dir := range directions {
					ni, nj := i+dir[0], j+dir[1]
					if ni >= 0 && ni < n && nj >= 0 && nj < n && grid[ni][nj] == 1 {
						uf.Union(i*n+j, ni*n+nj)
					}
				}
			}
		}
	}

	// 分别收集两个岛屿的点
	islandGroups := make(map[int][][]int)
	for _, point := range islands {
		i, j := point[0], point[1]
		root := uf.Find(i*n + j)
		islandGroups[root] = append(islandGroups[root], point)
	}

	// 提取两个岛屿
	groups := make([][][]int, 0)
	for _, group := range islandGroups {
		groups = append(groups, group)
	}

	if len(groups) != 2 {
		return -1
	}

	// 使用BFS计算最短距离
	return unionFindBFS(grid, groups[0], groups[1])
}

// 并查集辅助的BFS搜索
func unionFindBFS(grid [][]int, island1, island2 [][]int) int {
	n := len(grid)

	// 标记两个岛屿
	for _, point := range island1 {
		grid[point[0]][point[1]] = 2
	}
	for _, point := range island2 {
		grid[point[0]][point[1]] = 3
	}

	// 收集第一个岛屿的边界并开始BFS
	boundary := filterBoundaryPointsUnionFind(grid, island1, 2)

	queue := make([][]int, 0)
	dist := make([][]int, n)
	for i := range dist {
		dist[i] = make([]int, n)
		for j := range dist[i] {
			dist[i][j] = -1
		}
	}

	for _, point := range boundary {
		queue = append(queue, []int{point[0], point[1], 0})
		dist[point[0]][point[1]] = 0
	}

	// BFS搜索
	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]
		x, y, d := curr[0], curr[1], curr[2]

		for _, dir := range directions {
			nx, ny := x+dir[0], y+dir[1]

			if nx >= 0 && nx < n && ny >= 0 && ny < n {
				if grid[nx][ny] == 3 { // 找到第二个岛屿
					return d
				}
				if grid[nx][ny] == 0 && dist[nx][ny] == -1 { // 未访问的水域
					dist[nx][ny] = d + 1
					queue = append(queue, []int{nx, ny, d + 1})
				}
			}
		}
	}

	return -1
}

// 并查集版本的边界点过滤
func filterBoundaryPointsUnionFind(grid [][]int, island [][]int, islandId int) [][]int {
	n := len(grid)
	boundary := make([][]int, 0)

	for _, point := range island {
		i, j := point[0], point[1]
		isBoundary := false

		for _, dir := range directions {
			ni, nj := i+dir[0], j+dir[1]
			if ni >= 0 && ni < n && nj >= 0 && nj < n && grid[ni][nj] == 0 {
				isBoundary = true
				break
			}
		}

		if isBoundary {
			boundary = append(boundary, point)
		}
	}

	return boundary
}

// ========== 工具函数 ==========
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func copyGrid(grid [][]int) [][]int {
	n := len(grid)
	newGrid := make([][]int, n)
	for i := range newGrid {
		newGrid[i] = make([]int, len(grid[i]))
		copy(newGrid[i], grid[i])
	}
	return newGrid
}

func printGrid(grid [][]int) {
	for _, row := range grid {
		fmt.Println(row)
	}
	fmt.Println()
}

// ========== 测试和性能评估 ==========
func main() {
	// 测试用例
	testCases := []struct {
		name     string
		grid     [][]int
		expected int
	}{
		{
			name: "示例1: 2x2简单情况",
			grid: [][]int{
				{0, 1},
				{1, 0},
			},
			expected: 1,
		},
		{
			name: "示例2: 3x3需要搭桥",
			grid: [][]int{
				{0, 1, 0},
				{0, 0, 0},
				{0, 0, 1},
			},
			expected: 2,
		},
		{
			name: "示例3: 5x5复杂情况",
			grid: [][]int{
				{1, 1, 1, 1, 1},
				{1, 0, 0, 0, 1},
				{1, 0, 1, 0, 1},
				{1, 0, 0, 0, 1},
				{1, 1, 1, 1, 1},
			},
			expected: 1,
		},
		{
			name: "测试4: 4x4对角分布",
			grid: [][]int{
				{1, 0, 0, 0},
				{0, 0, 0, 0},
				{0, 0, 0, 0},
				{0, 0, 0, 1},
			},
			expected: 5,
		},
		{
			name: "测试5: 6x6大岛屿",
			grid: [][]int{
				{1, 1, 0, 0, 0, 0},
				{1, 1, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 1, 1},
				{0, 0, 0, 0, 1, 1},
				{0, 0, 0, 0, 1, 1},
			},
			expected: 4,
		},
		{
			name: "测试6: 长条形岛屿",
			grid: [][]int{
				{1, 1, 1, 0, 0},
				{0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0},
				{0, 0, 0, 1, 1},
				{0, 0, 0, 0, 0},
			},
			expected: 3,
		},
		{
			name: "测试7: L形岛屿",
			grid: [][]int{
				{1, 0, 0, 0, 0},
				{1, 0, 0, 0, 0},
				{1, 1, 0, 1, 0},
				{0, 0, 0, 1, 0},
				{0, 0, 0, 1, 1},
			},
			expected: 1,
		},
		{
			name: "测试8: 相邻岛屿",
			grid: [][]int{
				{1, 1, 0, 1},
				{1, 0, 0, 1},
				{0, 0, 0, 0},
				{0, 0, 0, 0},
			},
			expected: 1,
		},
		{
			name: "测试9: 最小网格",
			grid: [][]int{
				{1, 0},
				{0, 1},
			},
			expected: 1,
		},
		{
			name: "测试10: 复杂形状",
			grid: [][]int{
				{1, 1, 0, 0, 1},
				{1, 0, 0, 0, 1},
				{0, 0, 0, 0, 1},
				{0, 0, 1, 1, 0},
				{0, 0, 1, 0, 0},
			},
			expected: 2,
		},
	}

	// 算法方法
	methods := []struct {
		name string
		fn   func([][]int) int
	}{
		{"DFS+BFS标准解法", shortestBridge1},
		{"双向BFS优化", shortestBridge2},
		{"A*启发式搜索", shortestBridge3},
		{"并查集+BFS", shortestBridge4},
	}

	fmt.Println("=== LeetCode 934. 最短的桥 - 测试结果 ===")
	fmt.Println()

	// 运行测试
	for _, tc := range testCases {
		fmt.Printf("测试用例: %s\n", tc.name)
		fmt.Println("输入网格:")
		printGrid(tc.grid)

		allPassed := true
		var results []int
		var times []time.Duration

		for _, method := range methods {
			// 复制网格以避免修改原数据
			gridCopy := copyGrid(tc.grid)

			start := time.Now()
			result := method.fn(gridCopy)
			elapsed := time.Since(start)

			results = append(results, result)
			times = append(times, elapsed)

			status := "✅"
			if result != tc.expected {
				status = "❌"
				allPassed = false
			}

			fmt.Printf("  %s: %d %s (耗时: %v)\n", method.name, result, status, elapsed)
		}

		fmt.Printf("期望结果: %d\n", tc.expected)
		if allPassed {
			fmt.Println("✅ 所有方法均通过")
		} else {
			fmt.Println("❌ 存在失败的方法")
		}
		fmt.Println(strings.Repeat("-", 50))
	}

	// 性能对比测试
	fmt.Println("\n=== 性能对比测试 ===")
	performanceTest()

	// 算法特性总结
	fmt.Println("\n=== 算法特性总结 ===")
	fmt.Println("1. DFS+BFS标准解法:")
	fmt.Println("   - 时间复杂度: O(n²)")
	fmt.Println("   - 空间复杂度: O(n²)")
	fmt.Println("   - 特点: 经典解法，易理解，性能稳定")
	fmt.Println()
	fmt.Println("2. 双向BFS优化:")
	fmt.Println("   - 时间复杂度: O(n²)")
	fmt.Println("   - 空间复杂度: O(n²)")
	fmt.Println("   - 特点: 理论上更快，两端同时搜索")
	fmt.Println()
	fmt.Println("3. A*启发式搜索:")
	fmt.Println("   - 时间复杂度: O(n²logn)")
	fmt.Println("   - 空间复杂度: O(n²)")
	fmt.Println("   - 特点: 智能搜索，适合复杂场景")
	fmt.Println()
	fmt.Println("4. 并查集+BFS:")
	fmt.Println("   - 时间复杂度: O(n²α(n))")
	fmt.Println("   - 空间复杂度: O(n²)")
	fmt.Println("   - 特点: 动态连通性，适合变化场景")
}

func performanceTest() {
	// 生成大规模测试用例
	sizes := []int{10, 20, 30, 50}

	for _, size := range sizes {
		fmt.Printf("\n性能测试 - 网格大小: %dx%d\n", size, size)

		// 生成测试网格
		grid := generateTestGrid(size)

		methods := []struct {
			name string
			fn   func([][]int) int
		}{
			{"DFS+BFS", shortestBridge1},
			{"双向BFS", shortestBridge2},
			{"A*搜索", shortestBridge3},
			{"并查集+BFS", shortestBridge4},
		}

		for _, method := range methods {
			gridCopy := copyGrid(grid)

			start := time.Now()
			result := method.fn(gridCopy)
			elapsed := time.Since(start)

			fmt.Printf("  %s: 结果=%d, 耗时=%v\n", method.name, result, elapsed)
		}
	}
}

func generateTestGrid(size int) [][]int {
	grid := make([][]int, size)
	for i := range grid {
		grid[i] = make([]int, size)
	}

	// 创建两个岛屿
	// 第一个岛屿在左上角
	for i := 0; i < size/4; i++ {
		for j := 0; j < size/4; j++ {
			grid[i][j] = 1
		}
	}

	// 第二个岛屿在右下角
	for i := 3 * size / 4; i < size; i++ {
		for j := 3 * size / 4; j < size; j++ {
			grid[i][j] = 1
		}
	}

	return grid
}

// 桥梁建设可视化演示
func demonstrateBridgeBuilding() {
	fmt.Println("\n=== 桥梁建设可视化演示 ===")

	grid := [][]int{
		{0, 1, 0, 0, 0},
		{0, 1, 0, 0, 0},
		{0, 0, 0, 0, 0},
		{0, 0, 0, 1, 0},
		{0, 0, 0, 1, 1},
	}

	fmt.Println("原始网格 (1=陆地, 0=水域):")
	printGrid(grid)

	fmt.Println("建设最短桥梁后:")
	// 这里可以添加可视化桥梁建设过程的代码
	result := shortestBridge1(copyGrid(grid))
	fmt.Printf("需要填充 %d 个水域格子来连接两个岛屿\n", result)
}

// 实际应用场景模拟
func realWorldApplications() {
	fmt.Println("\n=== 实际应用场景模拟 ===")

	scenarios := []struct {
		name        string
		description string
		grid        [][]int
	}{
		{
			name:        "海岛桥梁规划",
			description: "规划两个海岛之间的最短桥梁",
			grid: [][]int{
				{1, 1, 0, 0, 0, 0},
				{1, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 1, 1},
				{0, 0, 0, 0, 1, 0},
				{0, 0, 0, 0, 0, 0},
			},
		},
		{
			name:        "网络节点连接",
			description: "连接两个网络集群的最短路径",
			grid: [][]int{
				{1, 0, 0, 0, 1},
				{0, 0, 0, 0, 1},
				{0, 0, 0, 0, 0},
				{1, 0, 0, 0, 0},
				{1, 1, 0, 0, 0},
			},
		},
	}

	for _, scenario := range scenarios {
		fmt.Printf("场景: %s\n", scenario.name)
		fmt.Printf("描述: %s\n", scenario.description)
		fmt.Println("网格:")
		printGrid(scenario.grid)

		result := shortestBridge1(copyGrid(scenario.grid))
		fmt.Printf("最短连接距离: %d\n", result)
		fmt.Println(strings.Repeat("-", 40))
	}
}
