package main

import (
	"fmt"
	"strings"
	"time"
)

// 解法一：边界DFS（推荐解法）
// 时间复杂度：O(m×n)，空间复杂度：O(m×n)
func solve(board [][]byte) {
	if len(board) == 0 || len(board[0]) == 0 {
		return
	}

	m, n := len(board), len(board[0])

	// 从边界开始DFS，标记所有连通到边界的'O'
	// 上边界和下边界
	for j := 0; j < n; j++ {
		if board[0][j] == 'O' {
			dfs(board, 0, j)
		}
		if board[m-1][j] == 'O' {
			dfs(board, m-1, j)
		}
	}

	// 左边界和右边界
	for i := 0; i < m; i++ {
		if board[i][0] == 'O' {
			dfs(board, i, 0)
		}
		if board[i][n-1] == 'O' {
			dfs(board, i, n-1)
		}
	}

	// 最终处理：'O' -> 'X'（被围绕），'#' -> 'O'（不被围绕）
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			switch board[i][j] {
			case 'O':
				board[i][j] = 'X' // 被围绕的区域
			case '#':
				board[i][j] = 'O' // 连通到边界的区域
			}
		}
	}
}

// DFS递归函数
func dfs(board [][]byte, i, j int) {
	if i < 0 || i >= len(board) || j < 0 || j >= len(board[0]) || board[i][j] != 'O' {
		return
	}

	// 标记为临时字符，表示连通到边界
	board[i][j] = '#'

	// 四个方向递归
	dfs(board, i-1, j) // 上
	dfs(board, i+1, j) // 下
	dfs(board, i, j-1) // 左
	dfs(board, i, j+1) // 右
}

// 解法二：边界BFS
// 时间复杂度：O(m×n)，空间复杂度：O(m×n)
func solveBFS(board [][]byte) {
	if len(board) == 0 || len(board[0]) == 0 {
		return
	}

	m, n := len(board), len(board[0])
	queue := [][]int{}

	// 将边界的'O'加入队列
	for j := 0; j < n; j++ {
		if board[0][j] == 'O' {
			board[0][j] = '#'
			queue = append(queue, []int{0, j})
		}
		if board[m-1][j] == 'O' {
			board[m-1][j] = '#'
			queue = append(queue, []int{m - 1, j})
		}
	}

	for i := 0; i < m; i++ {
		if board[i][0] == 'O' {
			board[i][0] = '#'
			queue = append(queue, []int{i, 0})
		}
		if board[i][n-1] == 'O' {
			board[i][n-1] = '#'
			queue = append(queue, []int{i, n - 1})
		}
	}

	// BFS遍历
	directions := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

	for len(queue) > 0 {
		cell := queue[0]
		queue = queue[1:]
		x, y := cell[0], cell[1]

		for _, dir := range directions {
			nx, ny := x+dir[0], y+dir[1]

			if nx >= 0 && nx < m && ny >= 0 && ny < n && board[nx][ny] == 'O' {
				board[nx][ny] = '#'
				queue = append(queue, []int{nx, ny})
			}
		}
	}

	// 最终处理
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == 'O' {
				board[i][j] = 'X'
			} else if board[i][j] == '#' {
				board[i][j] = 'O'
			}
		}
	}
}

// 解法三：并查集
// 时间复杂度：O(m×n×α(m×n))，空间复杂度：O(m×n)
type UnionFind struct {
	parent []int
	rank   []int
}

func newUnionFind(n int) *UnionFind {
	parent := make([]int, n)
	rank := make([]int, n)
	for i := range parent {
		parent[i] = i
	}
	return &UnionFind{parent, rank}
}

func (uf *UnionFind) find(x int) int {
	if uf.parent[x] != x {
		uf.parent[x] = uf.find(uf.parent[x]) // 路径压缩
	}
	return uf.parent[x]
}

func (uf *UnionFind) union(x, y int) {
	px, py := uf.find(x), uf.find(y)
	if px == py {
		return
	}

	// 按秩合并
	if uf.rank[px] < uf.rank[py] {
		uf.parent[px] = py
	} else if uf.rank[px] > uf.rank[py] {
		uf.parent[py] = px
	} else {
		uf.parent[py] = px
		uf.rank[px]++
	}
}

func (uf *UnionFind) connected(x, y int) bool {
	return uf.find(x) == uf.find(y)
}

func solveUnionFind(board [][]byte) {
	if len(board) == 0 || len(board[0]) == 0 {
		return
	}

	m, n := len(board), len(board[0])

	// 创建并查集，额外添加一个虚拟边界节点
	uf := newUnionFind(m*n + 1)
	boundaryNode := m * n

	// 将坐标转换为一维索引
	getIndex := func(i, j int) int {
		return i*n + j
	}

	// 遍历矩阵，构建并查集
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == 'O' {
				index := getIndex(i, j)

				// 如果在边界，与虚拟边界节点合并
				if i == 0 || i == m-1 || j == 0 || j == n-1 {
					uf.union(index, boundaryNode)
				}

				// 与相邻的'O'合并
				directions := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
				for _, dir := range directions {
					ni, nj := i+dir[0], j+dir[1]
					if ni >= 0 && ni < m && nj >= 0 && nj < n && board[ni][nj] == 'O' {
						uf.union(index, getIndex(ni, nj))
					}
				}
			}
		}
	}

	// 最终处理：不与边界连通的'O'改为'X'
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == 'O' && !uf.connected(getIndex(i, j), boundaryNode) {
				board[i][j] = 'X'
			}
		}
	}
}

// 解法四：迭代DFS（避免栈溢出）
// 时间复杂度：O(m×n)，空间复杂度：O(m×n)
func solveIterative(board [][]byte) {
	if len(board) == 0 || len(board[0]) == 0 {
		return
	}

	m, n := len(board), len(board[0])
	stack := [][]int{}

	// 将边界的'O'压入栈
	for j := 0; j < n; j++ {
		if board[0][j] == 'O' {
			stack = append(stack, []int{0, j})
		}
		if board[m-1][j] == 'O' {
			stack = append(stack, []int{m - 1, j})
		}
	}

	for i := 0; i < m; i++ {
		if board[i][0] == 'O' {
			stack = append(stack, []int{i, 0})
		}
		if board[i][n-1] == 'O' {
			stack = append(stack, []int{i, n - 1})
		}
	}

	// 迭代DFS
	directions := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

	for len(stack) > 0 {
		cell := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		x, y := cell[0], cell[1]

		if x < 0 || x >= m || y < 0 || y >= n || board[x][y] != 'O' {
			continue
		}

		board[x][y] = '#'

		for _, dir := range directions {
			stack = append(stack, []int{x + dir[0], y + dir[1]})
		}
	}

	// 最终处理
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == 'O' {
				board[i][j] = 'X'
			} else if board[i][j] == '#' {
				board[i][j] = 'O'
			}
		}
	}
}

// 辅助函数：复制矩阵
func copyBoard(board [][]byte) [][]byte {
	result := make([][]byte, len(board))
	for i := range board {
		result[i] = make([]byte, len(board[i]))
		copy(result[i], board[i])
	}
	return result
}

// 辅助函数：打印矩阵
func printBoard(board [][]byte) {
	for _, row := range board {
		fmt.Printf("[")
		for j, cell := range row {
			if j > 0 {
				fmt.Print(",")
			}
			fmt.Printf("\"%c\"", cell)
		}
		fmt.Printf("]\n")
	}
}

// 辅助函数：比较两个矩阵是否相等
func equalBoards(a, b [][]byte) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if len(a[i]) != len(b[i]) {
			return false
		}
		for j := range a[i] {
			if a[i][j] != b[i][j] {
				return false
			}
		}
	}
	return true
}

// 测试函数
func testSolve() {
	testCases := []struct {
		board    [][]byte
		expected [][]byte
		desc     string
	}{
		{
			[][]byte{
				{'X', 'X', 'X', 'X'},
				{'X', 'O', 'O', 'X'},
				{'X', 'X', 'O', 'X'},
				{'X', 'O', 'X', 'X'},
			},
			[][]byte{
				{'X', 'X', 'X', 'X'},
				{'X', 'X', 'X', 'X'},
				{'X', 'X', 'X', 'X'},
				{'X', 'O', 'X', 'X'},
			},
			"示例1：标准被围绕区域",
		},
		{
			[][]byte{{'X'}},
			[][]byte{{'X'}},
			"示例2：单元素矩阵",
		},
		{
			[][]byte{
				{'O', 'O', 'O'},
				{'O', 'O', 'O'},
				{'O', 'O', 'O'},
			},
			[][]byte{
				{'O', 'O', 'O'},
				{'O', 'O', 'O'},
				{'O', 'O', 'O'},
			},
			"全是O的矩阵",
		},
		{
			[][]byte{
				{'X', 'X', 'X'},
				{'X', 'X', 'X'},
				{'X', 'X', 'X'},
			},
			[][]byte{
				{'X', 'X', 'X'},
				{'X', 'X', 'X'},
				{'X', 'X', 'X'},
			},
			"全是X的矩阵",
		},
		{
			[][]byte{
				{'O', 'X', 'O'},
				{'X', 'O', 'X'},
				{'O', 'X', 'O'},
			},
			[][]byte{
				{'O', 'X', 'O'},
				{'X', 'X', 'X'},
				{'O', 'X', 'O'},
			},
			"中心被围绕",
		},
		{
			[][]byte{{'O'}},
			[][]byte{{'O'}},
			"单个O在边界",
		},
		{
			[][]byte{
				{'O', 'O'},
				{'O', 'O'},
			},
			[][]byte{
				{'O', 'O'},
				{'O', 'O'},
			},
			"小矩阵边界连通",
		},
		{
			[][]byte{
				{'X', 'O', 'X'},
				{'O', 'X', 'O'},
				{'X', 'O', 'X'},
			},
			[][]byte{
				{'X', 'O', 'X'},
				{'O', 'X', 'O'},
				{'X', 'O', 'X'},
			},
			"复杂边界连通",
		},
	}

	fmt.Println("=== 被围绕的区域测试 ===")
	fmt.Println()

	for i, tc := range testCases {
		// 测试不同算法
		board1 := copyBoard(tc.board)
		board2 := copyBoard(tc.board)
		board3 := copyBoard(tc.board)

		solve(board1)
		solveBFS(board2)
		solveUnionFind(board3)

		status := "✅"
		if !equalBoards(board1, tc.expected) {
			status = "❌"
		}

		fmt.Printf("测试 %d: %s\n", i+1, tc.desc)
		fmt.Print("输入: ")
		printBoard(tc.board)
		fmt.Print("期望: ")
		printBoard(tc.expected)
		fmt.Print("DFS法: ")
		printBoard(board1)
		fmt.Printf("BFS一致: %t, 并查集一致: %t\n",
			equalBoards(board1, board2),
			equalBoards(board1, board3))
		fmt.Printf("结果: %s\n", status)
		fmt.Println(strings.Repeat("-", 50))
	}
}

// 性能测试
func benchmarkSolve() {
	fmt.Println()
	fmt.Println("=== 性能测试 ===")
	fmt.Println()

	// 构造测试数据
	testData := []struct {
		board [][]byte
		desc  string
	}{
		{generateBoard(10, 10, 0.3), "10×10矩阵"},
		{generateBoard(50, 50, 0.3), "50×50矩阵"},
		{generateBoard(100, 100, 0.3), "100×100矩阵"},
		{generateBoard(200, 200, 0.3), "200×200矩阵"},
	}

	algorithms := []struct {
		name string
		fn   func([][]byte)
	}{
		{"边界DFS", solve},
		{"边界BFS", solveBFS},
		{"并查集", solveUnionFind},
		{"迭代DFS", solveIterative},
	}

	for _, data := range testData {
		fmt.Printf("%s:\n", data.desc)

		for _, algo := range algorithms {
			board := copyBoard(data.board)
			start := time.Now()
			algo.fn(board)
			duration := time.Since(start)

			fmt.Printf("  %s: 耗时 %v\n", algo.name, duration)
		}
		fmt.Println()
	}
}

// 生成测试矩阵
func generateBoard(m, n int, oRatio float64) [][]byte {
	board := make([][]byte, m)
	for i := range board {
		board[i] = make([]byte, n)
		for j := range board[i] {
			// 根据比例随机生成'O'和'X'
			if float64((i*n+j)%100)/100 < oRatio {
				board[i][j] = 'O'
			} else {
				board[i][j] = 'X'
			}
		}
	}
	return board
}

// 演示DFS过程
func demonstrateDFS() {
	fmt.Println()
	fmt.Println("=== DFS过程演示 ===")

	board := [][]byte{
		{'X', 'X', 'X', 'X'},
		{'X', 'O', 'O', 'X'},
		{'X', 'X', 'O', 'X'},
		{'X', 'O', 'X', 'X'},
	}

	fmt.Println("原始矩阵:")
	printBoard(board)

	fmt.Println("从边界(3,1)的'O'开始DFS标记:")

	// 手动演示一步步标记过程
	demo := copyBoard(board)
	fmt.Println("步骤1: 标记(3,1) -> '#'")
	demo[3][1] = '#'
	printBoard(demo)

	fmt.Println("步骤2: 完成所有连通标记")
	solve(board)
	printBoard(board)

	fmt.Println("说明: 底部的'O'连通到边界，保持为'O'")
	fmt.Println("     内部的'O'被完全围绕，变为'X'")
}

func main() {
	fmt.Println("130. 被围绕的区域 - 多种解法实现")
	fmt.Println("====================================")

	// 基础功能测试
	testSolve()

	// 性能对比测试
	benchmarkSolve()

	// DFS过程演示
	demonstrateDFS()

	// 展示算法特点
	fmt.Println()
	fmt.Println("=== 算法特点分析 ===")
	fmt.Println("1. 边界DFS：从边界开始递归标记，简单直观")
	fmt.Println("2. 边界BFS：层次遍历标记，避免栈溢出")
	fmt.Println("3. 并查集：维护连通性，适合复杂查询")
	fmt.Println("4. 迭代DFS：模拟递归栈，控制空间使用")

	fmt.Println()
	fmt.Println("=== 关键技巧总结 ===")
	fmt.Println("• 逆向思维：标记不被围绕的区域而非被围绕的")
	fmt.Println("• 边界优先：从矩阵边界开始搜索")
	fmt.Println("• 临时标记：使用临时字符区分不同状态")
	fmt.Println("• 原地修改：节省空间复杂度")
	fmt.Println("• 连通性：理解图的连通分量概念")
}
