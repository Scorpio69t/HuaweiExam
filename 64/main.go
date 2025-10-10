package main

import (
	"fmt"
)

// ===== 方法一：二维动态规划（经典解法） =====
// dp[i][j] 表示到达位置 (i,j) 的最小路径和
// 状态转移：dp[i][j] = grid[i][j] + min(dp[i-1][j], dp[i][j-1])
// 时间 O(m*n)，空间 O(m*n)
func minPathSumDP2D(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}

	m, n := len(grid), len(grid[0])
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}

	// 初始化起点
	dp[0][0] = grid[0][0]

	// 初始化第一行（只能从左边来）
	for j := 1; j < n; j++ {
		dp[0][j] = dp[0][j-1] + grid[0][j]
	}

	// 初始化第一列（只能从上边来）
	for i := 1; i < m; i++ {
		dp[i][0] = dp[i-1][0] + grid[i][0]
	}

	// 填充DP表
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = grid[i][j] + min(dp[i-1][j], dp[i][j-1])
		}
	}

	return dp[m-1][n-1]
}

// ===== 方法二：一维动态规划（空间优化） =====
// 滚动数组优化：只保存一行的状态
// 时间 O(m*n)，空间 O(n)
func minPathSumDP1D(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}

	m, n := len(grid), len(grid[0])
	dp := make([]int, n)

	// 初始化第一行
	dp[0] = grid[0][0]
	for j := 1; j < n; j++ {
		dp[j] = dp[j-1] + grid[0][j]
	}

	// 逐行更新
	for i := 1; i < m; i++ {
		// 第一列只能从上方来
		dp[0] += grid[i][0]

		// 其他列取上方和左方的最小值
		for j := 1; j < n; j++ {
			dp[j] = grid[i][j] + min(dp[j], dp[j-1])
		}
	}

	return dp[n-1]
}

// ===== 方法三：原地修改（极致空间优化） =====
// 直接在 grid 上修改，空间 O(1)
// 注意：会修改原数组
func minPathSumInPlace(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}

	m, n := len(grid), len(grid[0])

	// 初始化第一行
	for j := 1; j < n; j++ {
		grid[0][j] += grid[0][j-1]
	}

	// 初始化第一列
	for i := 1; i < m; i++ {
		grid[i][0] += grid[i-1][0]
	}

	// 填充其余位置
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			grid[i][j] += min(grid[i-1][j], grid[i][j-1])
		}
	}

	return grid[m-1][n-1]
}

// ===== 方法四：记忆化递归（自顶向下） =====
// 从终点递归到起点，记忆化避免重复计算
// 时间 O(m*n)，空间 O(m*n)
func minPathSumMemo(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}

	m, n := len(grid), len(grid[0])
	memo := make([][]int, m)
	for i := range memo {
		memo[i] = make([]int, n)
		for j := range memo[i] {
			memo[i][j] = -1
		}
	}

	var dfs func(i, j int) int
	dfs = func(i, j int) int {
		// 边界：到达起点
		if i == 0 && j == 0 {
			return grid[0][0]
		}

		// 越界处理
		if i < 0 || j < 0 {
			return 1<<31 - 1 // 返回极大值
		}

		// 已计算过
		if memo[i][j] != -1 {
			return memo[i][j]
		}

		// 递归计算
		memo[i][j] = grid[i][j] + min(dfs(i-1, j), dfs(i, j-1))
		return memo[i][j]
	}

	return dfs(m-1, n-1)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func copyGrid(grid [][]int) [][]int {
	m, n := len(grid), len(grid[0])
	newGrid := make([][]int, m)
	for i := range newGrid {
		newGrid[i] = make([]int, n)
		copy(newGrid[i], grid[i])
	}
	return newGrid
}

func main() {
	tests := []struct {
		grid   [][]int
		expect int
	}{
		{
			grid:   [][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}},
			expect: 7,
		},
		{
			grid:   [][]int{{1, 2, 3}, {4, 5, 6}},
			expect: 12,
		},
		{
			grid:   [][]int{{1}},
			expect: 1,
		},
		{
			grid:   [][]int{{1, 2}, {1, 1}},
			expect: 3,
		},
		{
			grid:   [][]int{{0, 0, 0}, {0, 0, 0}},
			expect: 0,
		},
	}

	methods := []struct {
		name string
		fn   func([][]int) int
	}{
		{"二维DP", minPathSumDP2D},
		{"一维DP", minPathSumDP1D},
		{"原地修改", minPathSumInPlace},
		{"记忆化递归", minPathSumMemo},
	}

	fmt.Println("64. 最小路径和 - 多解法对比")
	for idx, tc := range tests {
		fmt.Printf("用例%d: grid=%v\n", idx+1, tc.grid)
		for _, m := range methods {
			// 原地修改会改变数组，需要拷贝
			gridCopy := copyGrid(tc.grid)
			got := m.fn(gridCopy)
			status := "✅"
			if got != tc.expect {
				status = "❌"
			}
			fmt.Printf("  %-8s => %d %s\n", m.name, got, status)
		}
		fmt.Printf("  期望 => %d\n", tc.expect)
		fmt.Println("------------------------------")
	}
}
