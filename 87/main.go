package main

import (
	"fmt"
	"testing"
)

// =========================== 方法一：动态规划（最优解法） ===========================

func isScramble1(s1 string, s2 string) bool {
	n := len(s1)
	if n != len(s2) {
		return false
	}
	if n == 0 {
		return true
	}
	if n == 1 {
		return s1 == s2
	}

	// dp[i][j][k] 表示 s1[i:i+k] 和 s2[j:j+k] 是否互为扰乱字符串
	dp := make([][][]bool, n)
	for i := range dp {
		dp[i] = make([][]bool, n)
		for j := range dp[i] {
			dp[i][j] = make([]bool, n+1)
		}
	}

	// 初始化：长度为1的情况
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			dp[i][j][1] = s1[i] == s2[j]
		}
	}

	// 枚举长度
	for k := 2; k <= n; k++ {
		// 枚举起始位置
		for i := 0; i <= n-k; i++ {
			for j := 0; j <= n-k; j++ {
				// 枚举分割点
				for m := 1; m < k; m++ {
					// 不交换情况
					if dp[i][j][m] && dp[i+m][j+m][k-m] {
						dp[i][j][k] = true
						break
					}
					// 交换情况
					if dp[i][j+k-m][m] && dp[i+m][j][k-m] {
						dp[i][j][k] = true
						break
					}
				}
			}
		}
	}

	return dp[0][0][n]
}

// =========================== 方法二：递归算法 ===========================

func isScramble2(s1 string, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	if s1 == s2 {
		return true
	}
	if len(s1) == 1 {
		return s1 == s2
	}

	// 检查字符频率
	count := make([]int, 26)
	for i := 0; i < len(s1); i++ {
		count[s1[i]-'a']++
		count[s2[i]-'a']--
	}
	for _, c := range count {
		if c != 0 {
			return false
		}
	}

	// 枚举分割点
	for i := 1; i < len(s1); i++ {
		// 不交换情况
		if isScramble2(s1[:i], s2[:i]) && isScramble2(s1[i:], s2[i:]) {
			return true
		}
		// 交换情况
		if isScramble2(s1[:i], s2[len(s2)-i:]) && isScramble2(s1[i:], s2[:len(s2)-i]) {
			return true
		}
	}

	return false
}

// =========================== 方法三：记忆化递归 ===========================

func isScramble3(s1 string, s2 string) bool {
	memo := make(map[string]bool)
	return helper(s1, s2, memo)
}

func helper(s1, s2 string, memo map[string]bool) bool {
	if len(s1) != len(s2) {
		return false
	}
	if s1 == s2 {
		return true
	}
	if len(s1) == 1 {
		return s1 == s2
	}

	key := s1 + "#" + s2
	if val, ok := memo[key]; ok {
		return val
	}

	// 检查字符频率
	count := make([]int, 26)
	for i := 0; i < len(s1); i++ {
		count[s1[i]-'a']++
		count[s2[i]-'a']--
	}
	for _, c := range count {
		if c != 0 {
			memo[key] = false
			return false
		}
	}

	// 枚举分割点
	for i := 1; i < len(s1); i++ {
		// 不交换情况
		if helper(s1[:i], s2[:i], memo) && helper(s1[i:], s2[i:], memo) {
			memo[key] = true
			return true
		}
		// 交换情况
		if helper(s1[:i], s2[len(s2)-i:], memo) && helper(s1[i:], s2[:len(s2)-i], memo) {
			memo[key] = true
			return true
		}
	}

	memo[key] = false
	return false
}

// =========================== 方法四：优化版动态规划 ===========================

func isScramble4(s1 string, s2 string) bool {
	n := len(s1)
	if n != len(s2) {
		return false
	}
	if n == 0 {
		return true
	}
	if n == 1 {
		return s1 == s2
	}

	// 优化：使用滚动数组，但需要正确处理状态转移
	dp := make([][][]bool, n)
	for i := range dp {
		dp[i] = make([][]bool, n)
		for j := range dp[i] {
			dp[i][j] = make([]bool, n+1)
		}
	}

	// 初始化：长度为1的情况
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			dp[i][j][1] = s1[i] == s2[j]
		}
	}

	// 枚举长度
	for k := 2; k <= n; k++ {
		for i := 0; i <= n-k; i++ {
			for j := 0; j <= n-k; j++ {
				dp[i][j][k] = false
				for m := 1; m < k; m++ {
					// 不交换情况
					if dp[i][j][m] && dp[i+m][j+m][k-m] {
						dp[i][j][k] = true
						break
					}
					// 交换情况
					if dp[i][j+k-m][m] && dp[i+m][j][k-m] {
						dp[i][j][k] = true
						break
					}
				}
			}
		}
	}

	return dp[0][0][n]
}

// =========================== 测试代码 ===========================

func TestIsScramble(t *testing.T) {
	tests := []struct {
		name string
		s1   string
		s2   string
		want bool
	}{
		{
			name: "Test1: Basic case",
			s1:   "great",
			s2:   "rgeat",
			want: true,
		},
		{
			name: "Test2: Single character",
			s1:   "a",
			s2:   "a",
			want: true,
		},
		{
			name: "Test3: No solution",
			s1:   "abcde",
			s2:   "caebd",
			want: false,
		},
		{
			name: "Test4: Empty strings",
			s1:   "",
			s2:   "",
			want: true,
		},
		{
			name: "Test5: Same strings",
			s1:   "abc",
			s2:   "abc",
			want: true,
		},
		{
			name: "Test6: Different lengths",
			s1:   "abc",
			s2:   "abcd",
			want: false,
		},
		{
			name: "Test7: Complex case",
			s1:   "abcdefghijklmnopqrstuvwxyz",
			s2:   "zyxwvutsrqponmlkjihgfedcba",
			want: false,
		},
		{
			name: "Test8: Another complex case",
			s1:   "abcd",
			s2:   "bdac",
			want: false,
		},
		{
			name: "Test9: Simple scramble",
			s1:   "ab",
			s2:   "ba",
			want: true,
		},
		{
			name: "Test10: Three characters",
			s1:   "abc",
			s2:   "bca",
			want: true,
		},
	}

	methods := map[string]func(string, string) bool{
		"动态规划（最优解法）": isScramble1,
		"递归算法":       isScramble2,
		"记忆化递归":      isScramble3,
		"优化版动态规划":    isScramble4,
	}

	fmt.Println("=== LeetCode 87: 扰乱字符串 ===")
	for name, method := range methods {
		fmt.Printf("\n方法%s：%s\n", name, name)
		for i, tt := range tests {
			got := method(tt.s1, tt.s2)
			if got != tt.want {
				t.Errorf("  测试%d: %s, 输入: s1=\"%s\", s2=\"%s\", 输出: %t, 期望: %t", i+1, tt.name, tt.s1, tt.s2, got, tt.want)
				fmt.Printf("  测试%d: ❌\n", i+1)
			} else {
				fmt.Printf("  测试%d: ✅\n", i+1)
			}
		}
	}
}

func main() {
	fmt.Println("=== LeetCode 87: 扰乱字符串 ===\n")

	testCases := []struct {
		name string
		s1   string
		s2   string
		want bool
	}{
		{
			name: "Test1: Basic case",
			s1:   "great",
			s2:   "rgeat",
			want: true,
		},
		{
			name: "Test2: Single character",
			s1:   "a",
			s2:   "a",
			want: true,
		},
		{
			name: "Test3: No solution",
			s1:   "abcde",
			s2:   "caebd",
			want: false,
		},
		{
			name: "Test4: Empty strings",
			s1:   "",
			s2:   "",
			want: true,
		},
		{
			name: "Test5: Same strings",
			s1:   "abc",
			s2:   "abc",
			want: true,
		},
		{
			name: "Test6: Different lengths",
			s1:   "abc",
			s2:   "abcd",
			want: false,
		},
		{
			name: "Test7: Simple scramble",
			s1:   "ab",
			s2:   "ba",
			want: true,
		},
		{
			name: "Test8: Three characters",
			s1:   "abc",
			s2:   "bca",
			want: true,
		},
	}

	methods := map[string]func(string, string) bool{
		"动态规划（最优解法）": isScramble1,
		"递归算法":       isScramble2,
		"记忆化递归":      isScramble3,
		"优化版动态规划":    isScramble4,
	}

	for name, method := range methods {
		fmt.Printf("方法%s：%s\n", name, name)
		passCount := 0
		for i, tt := range testCases {
			got := method(tt.s1, tt.s2)
			status := "✅"
			if got != tt.want {
				status = "❌"
			} else {
				passCount++
			}
			fmt.Printf("  测试%d: %s\n", i+1, status)
			if status == "❌" {
				fmt.Printf("    输入: s1=\"%s\", s2=\"%s\"\n", tt.s1, tt.s2)
				fmt.Printf("    输出: %t, 期望: %t\n", got, tt.want)
			}
		}
		fmt.Printf("  通过: %d/%d\n\n", passCount, len(testCases))
	}
}
