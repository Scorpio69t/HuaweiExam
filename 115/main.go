package main

import (
	"fmt"
)

// =========================== 方法一：二维动态规划（最优解法） ===========================
func numDistinct1(s string, t string) int {
	n, m := len(s), len(t)
	if n < m {
		return 0
	}

	// dp[i][j]表示s的前i个字符中t的前j个字符出现的次数
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, m+1)
	}

	// 边界条件：空字符串是任何字符串的子序列
	for i := 0; i <= n; i++ {
		dp[i][0] = 1
	}

	// 状态转移
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			if s[i-1] == t[j-1] {
				// 可以选择匹配或不匹配
				dp[i][j] = dp[i-1][j-1] + dp[i-1][j]
			} else {
				// 只能不匹配
				dp[i][j] = dp[i-1][j]
			}
		}
	}

	return dp[n][m]
}

// =========================== 方法二：滚动数组优化 ===========================
func numDistinct2(s string, t string) int {
	n, m := len(s), len(t)
	if n < m {
		return 0
	}

	// 使用一维数组，从后往前更新
	dp := make([]int, m+1)
	dp[0] = 1

	for i := 1; i <= n; i++ {
		// 从后往前更新，避免覆盖
		for j := m; j >= 1; j-- {
			if s[i-1] == t[j-1] {
				dp[j] += dp[j-1]
			}
		}
	}

	return dp[m]
}

// =========================== 方法三：递归 + 记忆化 ===========================
func numDistinct3(s string, t string) int {
	memo := make(map[string]int)
	return helper(s, t, 0, 0, memo)
}

func helper(s, t string, i, j int, memo map[string]int) int {
	// 如果t已经匹配完，返回1
	if j == len(t) {
		return 1
	}
	// 如果s已经用完但t还没匹配完，返回0
	if i == len(s) {
		return 0
	}

	key := fmt.Sprintf("%d,%d", i, j)
	if val, ok := memo[key]; ok {
		return val
	}

	result := 0
	// 如果当前字符匹配，可以选择匹配
	if s[i] == t[j] {
		result += helper(s, t, i+1, j+1, memo)
	}
	// 无论是否匹配，都可以选择不匹配
	result += helper(s, t, i+1, j, memo)

	memo[key] = result
	return result
}

// =========================== 方法四：优化的动态规划 ===========================
func numDistinct4(s string, t string) int {
	n, m := len(s), len(t)
	if n < m {
		return 0
	}

	// 提前检查：如果t中某个字符在s中不存在，返回0
	tChars := make(map[byte]bool)
	for i := 0; i < m; i++ {
		tChars[t[i]] = true
	}
	sChars := make(map[byte]bool)
	for i := 0; i < n; i++ {
		sChars[s[i]] = true
	}
	for ch := range tChars {
		if !sChars[ch] {
			return 0
		}
	}

	// 使用滚动数组
	dp := make([]int, m+1)
	dp[0] = 1

	for i := 1; i <= n; i++ {
		for j := m; j >= 1; j-- {
			if s[i-1] == t[j-1] {
				dp[j] += dp[j-1]
			}
		}
	}

	return dp[m]
}

// =========================== 测试 ===========================
func main() {
	fmt.Println("=== LeetCode 115: 不同的子序列 ===\n")

	testCases := []struct {
		name     string
		s        string
		t        string
		expected int
	}{
		{
			name:     "例1: s=\"rabbbit\", t=\"rabbit\"",
			s:        "rabbbit",
			t:        "rabbit",
			expected: 3,
		},
		{
			name:     "例2: s=\"babgbag\", t=\"bag\"",
			s:        "babgbag",
			t:        "bag",
			expected: 5,
		},
		{
			name:     "空字符串: s=\"\", t=\"\"",
			s:        "",
			t:        "",
			expected: 1,
		},
		{
			name:     "s为空: s=\"\", t=\"a\"",
			s:        "",
			t:        "a",
			expected: 0,
		},
		{
			name:     "t为空: s=\"abc\", t=\"\"",
			s:        "abc",
			t:        "",
			expected: 1,
		},
		{
			name:     "完全相同: s=\"abc\", t=\"abc\"",
			s:        "abc",
			t:        "abc",
			expected: 1,
		},
		{
			name:     "无匹配: s=\"abc\", t=\"def\"",
			s:        "abc",
			t:        "def",
			expected: 0,
		},
		{
			name:     "重复字符: s=\"aaa\", t=\"aa\"",
			s:        "aaa",
			t:        "aa",
			expected: 3,
		},
		{
			name:     "单字符: s=\"abc\", t=\"a\"",
			s:        "abc",
			t:        "a",
			expected: 1,
		},
		{
			name:     "复杂情况: s=\"aabbcc\", t=\"abc\"",
			s:        "aabbcc",
			t:        "abc",
			expected: 8,
		},
	}

	methods := map[string]func(string, string) int{
		"二维DP":   numDistinct1,
		"滚动数组优化": numDistinct2,
		"递归+记忆化": numDistinct3,
		"优化的DP":  numDistinct4,
	}

	for methodName, methodFunc := range methods {
		fmt.Printf("方法：%s\n", methodName)
		pass := 0
		for i, tc := range testCases {
			got := methodFunc(tc.s, tc.t)
			ok := got == tc.expected
			status := "✅"
			if !ok {
				status = "❌"
			}
			fmt.Printf("  测试%d(%s): %s\n", i+1, tc.name, status)
			if !ok {
				fmt.Printf("    输出: %d\n    期望: %d\n", got, tc.expected)
			} else {
				pass++
			}
		}
		fmt.Printf("  通过: %d/%d\n\n", pass, len(testCases))
	}
}
