package main

import (
	"fmt"
)

// =========================== 方法一：动态规划（最优解法） ===========================

func numDecodings(s string) int {
	n := len(s)
	if n == 0 || s[0] == '0' {
		return 0
	}

	dp := make([]int, n+1)
	dp[0] = 1
	dp[1] = 1

	for i := 2; i <= n; i++ {
		// 单字符解码
		if s[i-1] != '0' {
			dp[i] += dp[i-1]
		}

		// 双字符解码
		if s[i-2] == '1' || (s[i-2] == '2' && s[i-1] <= '6') {
			dp[i] += dp[i-2]
		}
	}

	return dp[n]
}

// =========================== 方法二：滚动数组优化 ===========================

func numDecodings2(s string) int {
	n := len(s)
	if n == 0 || s[0] == '0' {
		return 0
	}

	prev2 := 1 // dp[i-2]
	prev1 := 1 // dp[i-1]

	for i := 2; i <= n; i++ {
		curr := 0

		// 单字符解码
		if s[i-1] != '0' {
			curr += prev1
		}

		// 双字符解码
		if s[i-2] == '1' || (s[i-2] == '2' && s[i-1] <= '6') {
			curr += prev2
		}

		prev2 = prev1
		prev1 = curr
	}

	return prev1
}

// =========================== 方法三：递归+记忆化 ===========================

func numDecodings3(s string) int {
	if len(s) == 0 {
		return 0
	}
	memo := make(map[int]int)
	return helper(s, 0, memo)
}

func helper(s string, index int, memo map[int]int) int {
	if index == len(s) {
		return 1
	}

	if s[index] == '0' {
		return 0
	}

	if val, ok := memo[index]; ok {
		return val
	}

	result := helper(s, index+1, memo)

	if index+1 < len(s) {
		twoDigit := int(s[index]-'0')*10 + int(s[index+1]-'0')
		if twoDigit <= 26 {
			result += helper(s, index+2, memo)
		}
	}

	memo[index] = result
	return result
}

// =========================== 方法四：迭代DP（简化版） ===========================

func numDecodings4(s string) int {
	n := len(s)
	if n == 0 || s[0] == '0' {
		return 0
	}

	dp := make([]int, n+1)
	dp[0] = 1
	dp[1] = 1

	for i := 2; i <= n; i++ {
		dp[i] = 0

		// 单字符解码
		if s[i-1] != '0' {
			dp[i] += dp[i-1]
		}

		// 双字符解码
		if s[i-2] == '1' || (s[i-2] == '2' && s[i-1] <= '6') {
			dp[i] += dp[i-2]
		}
	}

	return dp[n]
}

// =========================== 测试代码 ===========================

func main() {
	fmt.Println("=== LeetCode 91: 解码方法 ===\n")

	testCases := []struct {
		name     string
		s        string
		expected int
	}{
		{
			name:     "Test1: Basic case",
			s:        "226",
			expected: 3,
		},
		{
			name:     "Test2: Simple case",
			s:        "12",
			expected: 2,
		},
		{
			name:     "Test3: Leading zero",
			s:        "06",
			expected: 0,
		},
		{
			name:     "Test4: Single zero",
			s:        "0",
			expected: 0,
		},
		{
			name:     "Test5: Complex case",
			s:        "11106",
			expected: 2,
		},
		{
			name:     "Test6: All ones",
			s:        "111",
			expected: 3,
		},
		{
			name:     "Test7: Large number",
			s:        "27",
			expected: 1,
		},
		{
			name:     "Test8: Empty string",
			s:        "",
			expected: 0,
		},
		{
			name:     "Test9: Single digit",
			s:        "1",
			expected: 1,
		},
		{
			name:     "Test10: Two zeros",
			s:        "00",
			expected: 0,
		},
	}

	methods := map[string]func(string) int{
		"动态规划（最优解法）": numDecodings,
		"滚动数组优化":     numDecodings2,
		"递归+记忆化":     numDecodings3,
		"迭代DP（简化版）":  numDecodings4,
	}

	for name, method := range methods {
		fmt.Printf("方法：%s\n", name)
		passCount := 0
		for i, tt := range testCases {
			got := method(tt.s)

			// 验证结果是否正确
			valid := got == tt.expected
			status := "✅"
			if !valid {
				status = "❌"
			} else {
				passCount++
			}
			fmt.Printf("  测试%d: %s\n", i+1, status)
			if status == "❌" {
				fmt.Printf("    输入: %s\n", tt.s)
				fmt.Printf("    输出: %d\n", got)
				fmt.Printf("    期望: %d\n", tt.expected)
			}
		}
		fmt.Printf("  通过: %d/%d\n\n", passCount, len(testCases))
	}
}
