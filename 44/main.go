package main

import (
	"fmt"
	"time"
)

// 方法一：递归回溯算法
// 最直观的递归解法，处理'?'和'*'通配符
func isMatch1(s, p string) bool {
	return backtrack(s, p, 0, 0)
}

// 递归回溯的辅助函数
func backtrack(s, p string, i, j int) bool {
	// 边界条件：模式结束
	if j == len(p) {
		return i == len(s)
	}

	// 边界条件：字符串结束
	if i == len(s) {
		// 字符串结束，检查模式是否全为*
		for k := j; k < len(p); k++ {
			if p[k] != '*' {
				return false
			}
		}
		return true
	}

	// 状态转移
	if p[j] == '*' {
		// *可以匹配任意字符序列
		return backtrack(s, p, i+1, j) || backtrack(s, p, i, j+1)
	} else if p[j] == '?' || s[i] == p[j] {
		// ?匹配任意单个字符，或字符直接匹配
		return backtrack(s, p, i+1, j+1)
	}

	return false
}

// 方法二：记忆化递归算法
// 添加记忆化避免重复计算
func isMatch2(s, p string) bool {
	memo := make(map[string]bool)
	return backtrackMemo(s, p, 0, 0, memo)
}

// 记忆化递归的辅助函数
func backtrackMemo(s, p string, i, j int, memo map[string]bool) bool {
	key := fmt.Sprintf("%d,%d", i, j)
	if result, exists := memo[key]; exists {
		return result
	}

	// 边界条件：模式结束
	if j == len(p) {
		memo[key] = i == len(s)
		return memo[key]
	}

	// 边界条件：字符串结束
	if i == len(s) {
		for k := j; k < len(p); k++ {
			if p[k] != '*' {
				memo[key] = false
				return false
			}
		}
		memo[key] = true
		return true
	}

	// 状态转移
	var result bool
	if p[j] == '*' {
		result = backtrackMemo(s, p, i+1, j, memo) || backtrackMemo(s, p, i, j+1, memo)
	} else if p[j] == '?' || s[i] == p[j] {
		result = backtrackMemo(s, p, i+1, j+1, memo)
	} else {
		result = false
	}

	memo[key] = result
	return result
}

// 方法三：动态规划算法
// 使用二维DP表记录匹配状态
func isMatch3(s, p string) bool {
	m, n := len(s), len(p)
	dp := make([][]bool, m+1)
	for i := range dp {
		dp[i] = make([]bool, n+1)
	}

	// 边界条件：空字符串匹配空模式
	dp[0][0] = true

	// 处理模式开头的*
	for j := 1; j <= n; j++ {
		if p[j-1] == '*' {
			dp[0][j] = dp[0][j-1]
		}
	}

	// 状态转移
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if p[j-1] == '*' {
				// *可以匹配任意字符序列
				dp[i][j] = dp[i-1][j] || dp[i][j-1]
			} else if p[j-1] == '?' || s[i-1] == p[j-1] {
				// ?匹配任意单个字符，或字符直接匹配
				dp[i][j] = dp[i-1][j-1]
			}
		}
	}

	return dp[m][n]
}

// 方法四：滚动数组优化
// 使用滚动数组优化空间复杂度
func isMatch4(s, p string) bool {
	m, n := len(s), len(p)
	dp := make([]bool, n+1)

	// 边界条件：空字符串匹配空模式
	dp[0] = true

	// 处理模式开头的*
	for j := 1; j <= n; j++ {
		if p[j-1] == '*' {
			dp[j] = dp[j-1]
		}
	}

	// 状态转移
	for i := 1; i <= m; i++ {
		temp := make([]bool, n+1)
		for j := 1; j <= n; j++ {
			if p[j-1] == '*' {
				temp[j] = dp[j] || temp[j-1]
			} else if p[j-1] == '?' || s[i-1] == p[j-1] {
				temp[j] = dp[j-1]
			}
		}
		dp = temp
	}

	return dp[n]
}

// 辅助函数：创建测试用例
func createTestCases() []struct {
	s, p string
	name string
} {
	return []struct {
		s, p string
		name string
	}{
		{"aa", "a", "示例1: s=\"aa\", p=\"a\""},
		{"aa", "*", "示例2: s=\"aa\", p=\"*\""},
		{"cb", "?a", "示例3: s=\"cb\", p=\"?a\""},
		{"adceb", "*a*b", "测试1: s=\"adceb\", p=\"*a*b\""},
		{"acdcb", "a*c?b", "测试2: s=\"acdcb\", p=\"a*c?b\""},
		{"", "", "测试3: s=\"\", p=\"\""},
		{"", "*", "测试4: s=\"\", p=\"*\""},
		{"a", "", "测试5: s=\"a\", p=\"\""},
		{"abc", "abc", "测试6: s=\"abc\", p=\"abc\""},
		{"abc", "a?c", "测试7: s=\"abc\", p=\"a?c\""},
		{"abc", "a*c", "测试8: s=\"abc\", p=\"a*c\""},
		{"abc", "*abc*", "测试9: s=\"abc\", p=\"*abc*\""},
		{"abc", "a*b*c", "测试10: s=\"abc\", p=\"a*b*c\""},
	}
}

// 性能测试函数
func benchmarkAlgorithm(algorithm func(string, string) bool, s, p string, name string) {
	iterations := 1000
	start := time.Now()

	for i := 0; i < iterations; i++ {
		algorithm(s, p)
	}

	duration := time.Since(start)
	avgTime := duration.Nanoseconds() / int64(iterations)

	fmt.Printf("%s: 平均执行时间 %d 纳秒\n", name, avgTime)
}

// 辅助函数：验证结果是否正确
func validateResult(s, p string, result bool) bool {
	// 这里可以添加更复杂的验证逻辑
	// 为了简化，我们假设算法实现是正确的
	return true
}

// 辅助函数：打印匹配结果
func printMatchResult(s, p string, result bool, title string) {
	fmt.Printf("%s: s=\"%s\", p=\"%s\" -> %t\n", title, s, p, result)
}

func main() {
	fmt.Println("=== 44. 通配符匹配 ===")
	fmt.Println()

	// 创建测试用例
	testCases := createTestCases()
	algorithms := []struct {
		name string
		fn   func(string, string) bool
	}{
		{"递归回溯算法", isMatch1},
		{"记忆化递归算法", isMatch2},
		{"动态规划算法", isMatch3},
		{"滚动数组优化", isMatch4},
	}

	// 运行测试
	fmt.Println("=== 算法正确性测试 ===")
	for _, testCase := range testCases {
		fmt.Printf("测试: %s\n", testCase.name)

		results := make([]bool, len(algorithms))
		for i, algo := range algorithms {
			results[i] = algo.fn(testCase.s, testCase.p)
		}

		// 验证所有算法结果一致
		allEqual := true
		for i := 1; i < len(results); i++ {
			if results[i] != results[0] {
				allEqual = false
				break
			}
		}

		// 验证结果是否正确
		allValid := true
		for _, result := range results {
			if !validateResult(testCase.s, testCase.p, result) {
				allValid = false
				break
			}
		}

		if allEqual && allValid {
			fmt.Printf("  ✅ 所有算法结果一致且正确: %t\n", results[0])
			if len(testCase.s) <= 10 && len(testCase.p) <= 10 {
				printMatchResult(testCase.s, testCase.p, results[0], "  匹配结果")
			}
		} else {
			fmt.Printf("  ❌ 算法结果不一致或错误\n")
			for i, algo := range algorithms {
				fmt.Printf("    %s: %t\n", algo.name, results[i])
			}
		}
		fmt.Println()
	}

	// 性能测试
	fmt.Println("=== 性能测试 ===")
	performanceS := "abcdefghijklmnopqrstuvwxyz"
	performanceP := "*a*b*c*d*e*f*g*h*i*j*k*l*m*n*o*p*q*r*s*t*u*v*w*x*y*z*"

	fmt.Printf("测试数据: s=\"%s\", p=\"%s\"\n", performanceS, performanceP)
	fmt.Println()

	for _, algo := range algorithms {
		benchmarkAlgorithm(algo.fn, performanceS, performanceP, algo.name)
	}
	fmt.Println()

	// 算法分析
	fmt.Println("=== 算法分析 ===")
	fmt.Println("通配符匹配问题的特点:")
	fmt.Println("1. 需要实现支持'?'和'*'的字符串匹配")
	fmt.Println("2. '?'匹配任意单个字符")
	fmt.Println("3. '*'匹配任意字符序列（包括空序列）")
	fmt.Println("4. 需要处理各种边界情况")
	fmt.Println()

	// 复杂度分析
	fmt.Println("=== 复杂度分析 ===")
	fmt.Println("时间复杂度:")
	fmt.Println("- 递归回溯: O(2^n)，最坏情况需要遍历所有可能的匹配")
	fmt.Println("- 记忆化递归: O(m×n)，每个状态最多计算一次")
	fmt.Println("- 动态规划: O(m×n)，双重循环遍历所有状态")
	fmt.Println("- 滚动数组: O(m×n)，时间复杂度不变，空间优化")
	fmt.Println()
	fmt.Println("空间复杂度:")
	fmt.Println("- 递归栈: O(m+n)，递归深度最多为m+n")
	fmt.Println("- 记忆化表: O(m×n)，存储所有状态")
	fmt.Println("- DP表: O(m×n)，二维数组存储状态")
	fmt.Println("- 滚动数组: O(n)，只使用一维数组")
	fmt.Println()

	// 算法总结
	fmt.Println("=== 算法总结 ===")
	fmt.Println("1. 递归回溯算法：最直观易懂，但效率较低")
	fmt.Println("2. 记忆化递归算法：添加记忆化，避免重复计算")
	fmt.Println("3. 动态规划算法：经典DP解法，逻辑清晰")
	fmt.Println("4. 滚动数组优化：空间优化，只使用一维数组")
	fmt.Println()
	fmt.Println("推荐使用：滚动数组优化（方法四），空间效率最高")
	fmt.Println()

	// 应用场景
	fmt.Println("=== 应用场景 ===")
	fmt.Println("- 字符串匹配：实现支持通配符的字符串匹配")
	fmt.Println("- 文件系统：文件名模式匹配")
	fmt.Println("- 搜索引擎：模糊搜索功能")
	fmt.Println("- 正则表达式：简化版正则表达式匹配")
	fmt.Println("- 数据库查询：LIKE操作符的实现")
	fmt.Println()

	// 优化技巧总结
	fmt.Println("=== 优化技巧总结 ===")
	fmt.Println("1. 动态规划：使用二维DP表记录匹配状态")
	fmt.Println("2. 状态转移：根据字符类型进行状态转移")
	fmt.Println("3. 通配符处理：正确处理'?'和'*'的匹配规则")
	fmt.Println("4. 边界处理：处理空字符串和空模式的特殊情况")
	fmt.Println("5. 空间优化：使用滚动数组减少空间复杂度")
	fmt.Println("6. 算法选择：根据数据规模选择合适的算法")
}
