package main

import (
	"fmt"
)

func main() {
	var s, t string
	fmt.Scanln(&s)
	fmt.Scanln(&t)

	result := editDistance(s, t)
	fmt.Println(result)
}

// editDistance 计算两个字符串的编辑距离
// 使用动态规划算法，时间复杂度O(mn)，空间复杂度O(mn)
func editDistance(s, t string) int {
	m, n := len(s), len(t)

	// dp[i][j] 表示 s[0:i] 和 t[0:j] 的编辑距离
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	// 初始化边界条件
	// s为空字符串时，需要插入t的所有字符
	for j := 0; j <= n; j++ {
		dp[0][j] = j
	}
	// t为空字符串时，需要删除s的所有字符
	for i := 0; i <= m; i++ {
		dp[i][0] = i
	}

	// 动态规划填表
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if s[i-1] == t[j-1] {
				// 字符相同，不需要操作
				dp[i][j] = dp[i-1][j-1]
			} else {
				// 字符不同，取三种操作的最小值
				dp[i][j] = min(
					dp[i-1][j]+1,   // 删除s[i-1]
					dp[i][j-1]+1,   // 插入t[j-1]
					dp[i-1][j-1]+1, // 替换s[i-1]为t[j-1]
				)
			}
		}
	}

	return dp[m][n]
}

// 空间优化版本：只使用O(min(m,n))空间
func editDistanceOptimized(s, t string) int {
	m, n := len(s), len(t)

	// 确保s是较短的字符串，优化空间使用
	if m > n {
		s, t = t, s
		m, n = n, m
	}

	// 只需要两行：当前行和上一行
	prev := make([]int, m+1)
	curr := make([]int, m+1)

	// 初始化第一行
	for i := 0; i <= m; i++ {
		prev[i] = i
	}

	// 逐行计算
	for j := 1; j <= n; j++ {
		curr[0] = j // 边界条件

		for i := 1; i <= m; i++ {
			if s[i-1] == t[j-1] {
				curr[i] = prev[i-1]
			} else {
				curr[i] = min(
					prev[i]+1,   // 删除
					curr[i-1]+1, // 插入
					prev[i-1]+1, // 替换
				)
			}
		}

		// 交换prev和curr
		prev, curr = curr, prev
	}

	return prev[m]
}

// 递归+记忆化版本（展示不同实现思路）
func editDistanceMemo(s, t string) int {
	m, n := len(s), len(t)
	memo := make([][]int, m+1)
	for i := range memo {
		memo[i] = make([]int, n+1)
		for j := range memo[i] {
			memo[i][j] = -1
		}
	}

	var dfs func(i, j int) int
	dfs = func(i, j int) int {
		// 边界条件
		if i == 0 {
			return j
		}
		if j == 0 {
			return i
		}

		// 记忆化
		if memo[i][j] != -1 {
			return memo[i][j]
		}

		if s[i-1] == t[j-1] {
			memo[i][j] = dfs(i-1, j-1)
		} else {
			memo[i][j] = min(
				dfs(i-1, j)+1,   // 删除
				dfs(i, j-1)+1,   // 插入
				dfs(i-1, j-1)+1, // 替换
			)
		}

		return memo[i][j]
	}

	return dfs(m, n)
}

// min 返回三个数中的最小值
func min(a, b, c int) int {
	if a <= b && a <= c {
		return a
	}
	if b <= c {
		return b
	}
	return c
}

// 用于两个数的min函数
func min2(a, b int) int {
	if a < b {
		return a
	}
	return b
}
