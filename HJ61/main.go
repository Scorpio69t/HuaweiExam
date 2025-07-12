package main

import "fmt"

// 解决放苹果问题的动态规划函数
func countWays(m, n int) int {
	// 创建二维DP表
	// dp[i][j] 表示将i个苹果放入j个盘子的方案数
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	// 边界条件初始化
	// 0个苹果只有一种放法（都不放）
	for j := 1; j <= n; j++ {
		dp[0][j] = 1
	}

	// 只有一个盘子，只有一种放法（全放一个盘子）
	for i := 0; i <= m; i++ {
		dp[i][1] = 1
	}

	// 填充DP表
	for i := 1; i <= m; i++ {
		for j := 2; j <= n; j++ {
			if i < j {
				// 苹果数小于盘子数，必有空盘子
				// 等价于将i个苹果放入i个盘子
				dp[i][j] = dp[i][i]
			} else {
				// 苹果数大于等于盘子数，分两种情况：
				// 1. 有空盘子：dp[i][j-1]
				// 2. 无空盘子（每个盘子至少放一个）：dp[i-j][j]
				dp[i][j] = dp[i][j-1] + dp[i-j][j]
			}
		}
	}

	return dp[m][n]
}

// 递归解法（供参考，理解递推关系）
func countWaysRecursive(m, n int) int {
	// 边界条件
	if m == 0 || n == 1 {
		return 1
	}

	// 苹果数小于盘子数，等价于m个苹果放m个盘子
	if m < n {
		return countWaysRecursive(m, m)
	}

	// 递推关系：有空盘子的情况 + 无空盘子的情况
	return countWaysRecursive(m, n-1) + countWaysRecursive(m-n, n)
}

func main() {
	var m, n int

	// 读取输入
	fmt.Scan(&m, &n)

	// 计算方案数
	result := countWaysRecursive(m, n)

	// 输出结果
	fmt.Println(result)
}
