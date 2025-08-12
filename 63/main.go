package main

import (
	"fmt"
	"strings"
	"time"
)

// 解法一：二维动态规划（直观）
// 状态定义：
//
//	dp[i][j] 表示到达单元格 (i, j) 的不同路径数
//
// 状态转移：
//
//	若 obstacleGrid[i][j] 为障碍(1)，则 dp[i][j] = 0
//	否则 dp[i][j] = dp[i-1][j] + dp[i][j-1]（来自上方和左方）
//
// 初始条件：
//
//	dp[0][0] = 1（前提是 obstacleGrid[0][0] == 0）
//
// 答案：
//
//	dp[m-1][n-1]
func uniquePathsWithObstaclesDP2D(obstacleGrid [][]int) int {
	m := len(obstacleGrid)
	if m == 0 {
		return 0
	}
	n := len(obstacleGrid[0])
	if n == 0 {
		return 0
	}
	if obstacleGrid[0][0] == 1 || obstacleGrid[m-1][n-1] == 1 {
		return 0
	}

	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}
	dp[0][0] = 1

	// 初始化首行
	for j := 1; j < n; j++ {
		if obstacleGrid[0][j] == 1 {
			dp[0][j] = 0
		} else {
			dp[0][j] = dp[0][j-1]
		}
	}
	// 初始化首列
	for i := 1; i < m; i++ {
		if obstacleGrid[i][0] == 1 {
			dp[i][0] = 0
		} else {
			dp[i][0] = dp[i-1][0]
		}
	}
	// 填表
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if obstacleGrid[i][j] == 1 {
				dp[i][j] = 0
				continue
			}
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}
	return dp[m-1][n-1]
}

// 解法二：一维动态规划（空间优化到 O(n)）
// 复用一行 dp：dp[j] 表示当前行列 j 的到达路径数
// 当遇到障碍时将 dp[j] 置 0；否则 dp[j] += dp[j-1]
func uniquePathsWithObstaclesDP1D(obstacleGrid [][]int) int {
	m := len(obstacleGrid)
	if m == 0 {
		return 0
	}
	n := len(obstacleGrid[0])
	if n == 0 {
		return 0
	}
	if obstacleGrid[0][0] == 1 || obstacleGrid[m-1][n-1] == 1 {
		return 0
	}

	dp := make([]int, n)
	dp[0] = 1
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if obstacleGrid[i][j] == 1 {
				dp[j] = 0
			} else if j > 0 {
				dp[j] += dp[j-1]
			}
		}
	}
	return dp[n-1]
}

// 解法三：记忆化搜索（自顶向下）
// 注意：在最坏情况下与 DP 等价，但实现上更贴近转移关系
func uniquePathsWithObstaclesMemo(obstacleGrid [][]int) int {
	m := len(obstacleGrid)
	if m == 0 {
		return 0
	}
	n := len(obstacleGrid[0])
	if n == 0 {
		return 0
	}
	if obstacleGrid[0][0] == 1 || obstacleGrid[m-1][n-1] == 1 {
		return 0
	}
	memo := make([][]int, m)
	for i := 0; i < m; i++ {
		memo[i] = make([]int, n)
		for j := 0; j < n; j++ {
			memo[i][j] = -1
		}
	}
	var dfs func(i, j int) int
	dfs = func(i, j int) int {
		if i < 0 || j < 0 {
			return 0
		}
		if obstacleGrid[i][j] == 1 {
			return 0
		}
		if i == 0 && j == 0 {
			return 1
		}
		if memo[i][j] != -1 {
			return memo[i][j]
		}
		memo[i][j] = dfs(i-1, j) + dfs(i, j-1)
		return memo[i][j]
	}
	return dfs(m-1, n-1)
}

// 辅助：对比多个算法的结果，确保一致性
func runTestCases() {
	type testCase struct {
		grid     [][]int
		expected int
		desc     string
	}
	tests := []testCase{
		{[][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}}, 2, "示例1：中间有障碍"},
		{[][]int{{0, 1}, {0, 0}}, 1, "示例2：两行两列"},
		{[][]int{{0}}, 1, "1x1 无障碍"},
		{[][]int{{1}}, 0, "1x1 有障碍"},
		{[][]int{{0, 0, 0, 0}}, 1, "单行无障碍"},
		{[][]int{{0, 1, 0, 0}}, 0, "单行有障碍阻断"},
		{[][]int{{0}, {0}, {0}}, 1, "单列无障碍"},
		{[][]int{{0}, {1}, {0}}, 0, "单列有障碍阻断"},
		{[][]int{{0, 0}, {1, 0}}, 1, "简单障碍-右下可达"},
		{[][]int{{0, 0}, {0, 1}}, 0, "终点为障碍"},
		{[][]int{{1, 0}, {0, 0}}, 0, "起点为障碍"},
	}

	fmt.Println("=== 63. 不同路径 II - 测试 ===")
	for i, tc := range tests {
		r1 := uniquePathsWithObstaclesDP2D(cloneGrid(tc.grid))
		r2 := uniquePathsWithObstaclesDP1D(cloneGrid(tc.grid))
		r3 := uniquePathsWithObstaclesMemo(cloneGrid(tc.grid))
		ok := (r1 == tc.expected) && (r2 == tc.expected) && (r3 == tc.expected)
		status := "✅"
		if !ok {
			status = "❌"
		}
		fmt.Printf("用例 %d: %s\n", i+1, tc.desc)
		fmt.Printf("输入: %v\n", tc.grid)
		fmt.Printf("期望: %d\n", tc.expected)
		fmt.Printf("二维DP: %d, 一维DP: %d, 记忆化: %d\n", r1, r2, r3)
		fmt.Printf("结果: %s\n", status)
		fmt.Println(strings.Repeat("-", 40))
	}
}

// 简单性能比较
func benchmark() {
	fmt.Println("\n=== 性能对比（粗略） ===")
	big := make([][]int, 100)
	for i := range big {
		big[i] = make([]int, 100)
	}
	start := time.Now()
	_ = uniquePathsWithObstaclesDP2D(big)
	d1 := time.Since(start)

	start = time.Now()
	_ = uniquePathsWithObstaclesDP1D(big)
	d2 := time.Since(start)

	start = time.Now()
	_ = uniquePathsWithObstaclesMemo(big)
	d3 := time.Since(start)

	fmt.Printf("二维DP: %v\n", d1)
	fmt.Printf("一维DP: %v\n", d2)
	fmt.Printf("记忆化: %v\n", d3)
}

func cloneGrid(g [][]int) [][]int {
	m := len(g)
	res := make([][]int, m)
	for i := 0; i < m; i++ {
		res[i] = append([]int(nil), g[i]...)
	}
	return res
}

func main() {
	fmt.Println("63. 不同路径 II")
	fmt.Println(strings.Repeat("=", 40))
	runTestCases()
	benchmark()
}
