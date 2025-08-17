package main

import (
	"fmt"
)

// isMatch 正则表达式匹配 - 动态规划方法
// 时间复杂度: O(m*n)，其中m和n分别是字符串和模式串的长度
// 空间复杂度: O(m*n)
func isMatch(s string, p string) bool {
	m, n := len(s), len(p)

	// 创建DP表，dp[i][j]表示s的前i个字符与p的前j个字符是否匹配
	dp := make([][]bool, m+1)
	for i := range dp {
		dp[i] = make([]bool, n+1)
	}

	// 空字符串与空模式匹配
	dp[0][0] = true

	// 处理模式串开头的*号情况
	for j := 1; j <= n; j++ {
		if p[j-1] == '*' {
			dp[0][j] = dp[0][j-2]
		}
	}

	// 填充DP表
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if p[j-1] == '.' || p[j-1] == s[i-1] {
				// 当前字符匹配
				dp[i][j] = dp[i-1][j-1]
			} else if p[j-1] == '*' {
				// 处理*号
				dp[i][j] = dp[i][j-2] // 匹配0次
				if p[j-2] == '.' || p[j-2] == s[i-1] {
					dp[i][j] = dp[i][j] || dp[i-1][j] // 匹配1次或多次
				}
			}
		}
	}

	return dp[m][n]
}

// isMatchRecursive 递归方法 - 带备忘录
// 时间复杂度: O(m*n)
// 空间复杂度: O(m*n)
func isMatchRecursive(s string, p string) bool {
	memo := make(map[string]bool)
	return matchHelper(s, p, 0, 0, memo)
}

func matchHelper(s, p string, i, j int, memo map[string]bool) bool {
	key := fmt.Sprintf("%d,%d", i, j)
	if val, exists := memo[key]; exists {
		return val
	}

	// 如果模式串已经用完
	if j == len(p) {
		result := i == len(s)
		memo[key] = result
		return result
	}

	// 如果字符串已经用完
	if i == len(s) {
		// 检查剩余的模式是否都是x*的形式
		if (len(p)-j)%2 == 1 {
			memo[key] = false
			return false
		}
		for k := j; k < len(p); k += 2 {
			if p[k] != '*' {
				memo[key] = false
				return false
			}
		}
		memo[key] = true
		return true
	}

	// 当前字符是否匹配
	currentMatch := i < len(s) && (p[j] == '.' || p[j] == s[i])

	var result bool
	if j+1 < len(p) && p[j+1] == '*' {
		// 处理*号：匹配0次或多次
		result = matchHelper(s, p, i, j+2, memo) || // 匹配0次
			(currentMatch && matchHelper(s, p, i+1, j, memo)) // 匹配多次
	} else {
		// 普通匹配
		result = currentMatch && matchHelper(s, p, i+1, j+1, memo)
	}

	memo[key] = result
	return result
}

// isMatchOptimized 优化版本 - 减少空间复杂度
// 时间复杂度: O(m*n)
// 空间复杂度: O(n)
func isMatchOptimized(s string, p string) bool {
	m, n := len(s), len(p)

	// 只使用一维DP数组
	dp := make([]bool, n+1)

	// 初始化第一行
	dp[0] = true
	for j := 1; j <= n; j++ {
		if p[j-1] == '*' {
			dp[j] = dp[j-2]
		}
	}

	// 逐行填充
	for i := 1; i <= m; i++ {
		prev := dp[0]
		dp[0] = false

		for j := 1; j <= n; j++ {
			temp := dp[j]
			if p[j-1] == '.' || p[j-1] == s[i-1] {
				dp[j] = prev
			} else if p[j-1] == '*' {
				dp[j] = dp[j-2] // 匹配0次
				if p[j-2] == '.' || p[j-2] == s[i-1] {
					dp[j] = dp[j] || dp[j] // 匹配多次
				}
			} else {
				dp[j] = false
			}
			prev = temp
		}
	}

	return dp[n]
}

// isMatchBacktrack 回溯法 - 不使用备忘录
// 时间复杂度: 最坏情况O(2^(m+n))
// 空间复杂度: O(m+n)
func isMatchBacktrack(s string, p string) bool {
	return backtrackHelper(s, p, 0, 0)
}

func backtrackHelper(s, p string, i, j int) bool {
	// 如果模式串已经用完
	if j == len(p) {
		return i == len(s)
	}

	// 如果字符串已经用完
	if i == len(s) {
		// 检查剩余的模式是否都是x*的形式
		if (len(p)-j)%2 == 1 {
			return false
		}
		for k := j; k < len(p); k += 2 {
			if p[k] != '*' {
				return false
			}
		}
		return true
	}

	// 当前字符是否匹配
	currentMatch := i < len(s) && (p[j] == '.' || p[j] == s[i])

	if j+1 < len(p) && p[j+1] == '*' {
		// 处理*号：匹配0次或多次
		return backtrackHelper(s, p, i, j+2) || // 匹配0次
			(currentMatch && backtrackHelper(s, p, i+1, j)) // 匹配多次
	} else {
		// 普通匹配
		return currentMatch && backtrackHelper(s, p, i+1, j+1)
	}
}

func main() {
	// 测试用例1
	s1 := "aa"
	p1 := "a"
	result1 := isMatch(s1, p1)
	fmt.Printf("示例1: s = \"%s\", p = \"%s\"\n", s1, p1)
	fmt.Printf("输出: %t\n", result1)
	fmt.Printf("期望: false\n")
	fmt.Printf("结果: %t\n", result1 == false)
	fmt.Println()

	// 测试用例2
	s2 := "aa"
	p2 := "a*"
	result2 := isMatch(s2, p2)
	fmt.Printf("示例2: s = \"%s\", p = \"%s\"\n", s2, p2)
	fmt.Printf("输出: %t\n", result2)
	fmt.Printf("期望: true\n")
	fmt.Printf("结果: %t\n", result2 == true)
	fmt.Println()

	// 测试用例3
	s3 := "ab"
	p3 := ".*"
	result3 := isMatch(s3, p3)
	fmt.Printf("示例3: s = \"%s\", p = \"%s\"\n", s3, p3)
	fmt.Printf("输出: %t\n", result3)
	fmt.Printf("期望: true\n")
	fmt.Printf("结果: %t\n", result3 == true)
	fmt.Println()

	// 额外测试用例
	s4 := "aab"
	p4 := "c*a*b"
	result4 := isMatch(s4, p4)
	fmt.Printf("额外测试: s = \"%s\", p = \"%s\"\n", s4, p4)
	fmt.Printf("输出: %t\n", result4)
	fmt.Printf("期望: true\n")
	fmt.Printf("结果: %t\n", result4 == true)
	fmt.Println()

	s5 := "mississippi"
	p5 := "mis*is*p*."
	result5 := isMatch(s5, p5)
	fmt.Printf("额外测试: s = \"%s\", p = \"%s\"\n", s5, p5)
	fmt.Printf("输出: %t\n", result5)
	fmt.Printf("期望: false\n")
	fmt.Printf("结果: %t\n", result5 == false)
	fmt.Println()

	// 测试递归版本
	fmt.Println("=== 递归版本测试 ===")
	result1Rec := isMatchRecursive(s1, p1)
	result2Rec := isMatchRecursive(s2, p2)
	fmt.Printf("递归版本示例1: %t\n", result1Rec)
	fmt.Printf("递归版本示例2: %t\n", result2Rec)
	fmt.Printf("结果一致: %t\n", result1Rec == result1 && result2Rec == result2)
	fmt.Println()

	// 测试优化版本
	fmt.Println("=== 优化版本测试 ===")
	result1Opt := isMatchOptimized(s1, p1)
	result2Opt := isMatchOptimized(s2, p2)
	fmt.Printf("优化版本示例1: %t\n", result1Opt)
	fmt.Printf("优化版本示例2: %t\n", result2Opt)
	fmt.Printf("结果一致: %t\n", result1Opt == result1 && result2Opt == result2)
	fmt.Println()

	// 测试回溯版本
	fmt.Println("=== 回溯版本测试 ===")
	result1Back := isMatchBacktrack(s1, p1)
	result2Back := isMatchBacktrack(s2, p2)
	fmt.Printf("回溯版本示例1: %t\n", result1Back)
	fmt.Printf("回溯版本示例2: %t\n", result2Back)
	fmt.Printf("结果一致: %t\n", result1Back == result1 && result2Back == result2)
	fmt.Println()

	// 边界值测试
	fmt.Println("=== 边界值测试 ===")
	boundaryTests := []struct {
		s, p string
	}{
		{"", ""},     // 空字符串
		{"", "a*"},   // 空字符串与模式
		{"a", ""},    // 字符串与空模式
		{"a", "a"},   // 单字符匹配
		{"a", "."},   // 单字符与通配符
		{"a", "a*"},  // 单字符与星号
		{"a", ".*"},  // 单字符与任意星号
		{"aa", "a"},  // 双字符与单字符
		{"aa", "a*"}, // 双字符与星号
		{"aa", ".*"}, // 双字符与任意星号
		{"ab", "a*"}, // 不同字符与星号
		{"ab", ".*"}, // 不同字符与任意星号
	}

	for _, test := range boundaryTests {
		result := isMatch(test.s, test.p)
		fmt.Printf("s = \"%s\", p = \"%s\", result = %t\n", test.s, test.p, result)
	}
}
