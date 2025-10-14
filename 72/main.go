package main

import (
	"fmt"
)

// =========================== 方法一：二维DP（标准解法） ===========================

func minDistance(word1 string, word2 string) int {
	m, n := len(word1), len(word2)
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	// 初始化边界
	for i := 0; i <= m; i++ {
		dp[i][0] = i
	}
	for j := 0; j <= n; j++ {
		dp[0][j] = j
	}

	// 填表
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = min(dp[i-1][j], dp[i][j-1], dp[i-1][j-1]) + 1
			}
		}
	}

	return dp[m][n]
}

// =========================== 方法二：一维DP（空间优化） ===========================

func minDistance2(word1 string, word2 string) int {
	m, n := len(word1), len(word2)
	if m < n {
		word1, word2 = word2, word1
		m, n = n, m
	}

	dp := make([]int, n+1)
	for j := 0; j <= n; j++ {
		dp[j] = j
	}

	for i := 1; i <= m; i++ {
		prev := dp[0]
		dp[0] = i
		for j := 1; j <= n; j++ {
			temp := dp[j]
			if word1[i-1] == word2[j-1] {
				dp[j] = prev
			} else {
				dp[j] = min(dp[j], dp[j-1], prev) + 1
			}
			prev = temp
		}
	}

	return dp[n]
}

// =========================== 方法三：递归+记忆化 ===========================

func minDistance3(word1 string, word2 string) int {
	m, n := len(word1), len(word2)
	memo := make([][]int, m)
	for i := range memo {
		memo[i] = make([]int, n)
		for j := range memo[i] {
			memo[i][j] = -1
		}
	}
	return dfs(word1, word2, m-1, n-1, memo)
}

func dfs(word1, word2 string, i, j int, memo [][]int) int {
	if i < 0 {
		return j + 1
	}
	if j < 0 {
		return i + 1
	}
	if memo[i][j] != -1 {
		return memo[i][j]
	}

	if word1[i] == word2[j] {
		memo[i][j] = dfs(word1, word2, i-1, j-1, memo)
	} else {
		memo[i][j] = min(
			dfs(word1, word2, i-1, j, memo),
			dfs(word1, word2, i, j-1, memo),
			dfs(word1, word2, i-1, j-1, memo),
		) + 1
	}

	return memo[i][j]
}

// =========================== 方法四：递归暴力（会超时） ===========================

func minDistance4(word1 string, word2 string) int {
	return bruteForce(word1, word2, len(word1)-1, len(word2)-1)
}

func bruteForce(word1, word2 string, i, j int) int {
	if i < 0 {
		return j + 1
	}
	if j < 0 {
		return i + 1
	}

	if word1[i] == word2[j] {
		return bruteForce(word1, word2, i-1, j-1)
	}

	return min(
		bruteForce(word1, word2, i-1, j),
		bruteForce(word1, word2, i, j-1),
		bruteForce(word1, word2, i-1, j-1),
	) + 1
}

func min(a, b, c int) int {
	if a < b {
		if a < c {
			return a
		}
		return c
	}
	if b < c {
		return b
	}
	return c
}

// =========================== 测试代码 ===========================

func main() {
	fmt.Println("=== LeetCode 72: 编辑距离 ===\n")

	testCases := []struct {
		word1  string
		word2  string
		expect int
	}{
		{"horse", "ros", 3},
		{"intention", "execution", 5},
		{"", "", 0},
		{"a", "", 1},
		{"", "a", 1},
		{"abc", "abc", 0},
		{"abc", "def", 3},
	}

	fmt.Println("方法一：二维DP")
	runTests(testCases, minDistance)

	fmt.Println("\n方法二：一维DP")
	runTests(testCases, minDistance2)

	fmt.Println("\n方法三：递归+记忆化")
	runTests(testCases, minDistance3)

	fmt.Println("\n方法四：递归暴力（仅短字符串）")
	runTests(testCases[:5], minDistance4) // 只测试前5个
}

func runTests(testCases []struct {
	word1  string
	word2  string
	expect int
}, fn func(string, string) int) {
	passCount := 0
	for i, tc := range testCases {
		result := fn(tc.word1, tc.word2)
		status := "✅"
		if result != tc.expect {
			status = "❌"
		} else {
			passCount++
		}
		fmt.Printf("  测试%d: %s \"%s\" → \"%s\" = %d\n", i+1, status, tc.word1, tc.word2, result)
	}
	fmt.Printf("  通过: %d/%d\n", passCount, len(testCases))
}
