package main

import (
	"fmt"
	"strings"
	"time"
)

// ========== 方法1: 二维动态规划（经典解法） ==========
func uniquePaths1(m int, n int) int {
	// 创建二维DP数组
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}

	// 初始化边界条件
	// 第一行：只能向右走，路径数都是1
	for j := 0; j < n; j++ {
		dp[0][j] = 1
	}
	// 第一列：只能向下走，路径数都是1
	for i := 0; i < m; i++ {
		dp[i][0] = 1
	}

	// 填充DP表格
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			// 当前位置路径数 = 上方路径数 + 左方路径数
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}

	return dp[m-1][n-1]
}

// ========== 方法2: 一维动态规划（空间优化） ==========
func uniquePaths2(m int, n int) int {
	// 只保存一行的状态
	dp := make([]int, n)

	// 初始化第一行为1
	for j := 0; j < n; j++ {
		dp[j] = 1
	}

	// 逐行更新
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			// dp[j] 现在表示上一行的值（上方）
			// dp[j-1] 表示当前行左侧的值（左方）
			dp[j] = dp[j] + dp[j-1]
		}
	}

	return dp[n-1]
}

// ========== 方法3: 数学公式法（组合数学） ==========
func uniquePaths3(m int, n int) int {
	// 边界检查
	if m == 1 || n == 1 {
		return 1
	}

	// 总共需要走 (m-1) 步向下，(n-1) 步向右
	// 总步数 = (m-1) + (n-1) = m + n - 2
	// 问题转化为：在 (m+n-2) 个位置中选择 (m-1) 个位置向下
	// 即计算组合数 C(m+n-2, m-1)

	// 为了避免溢出，选择较小的参数
	smaller := m - 1
	if n-1 < smaller {
		smaller = n - 1
	}

	// 防止大数计算溢出，限制在合理范围
	if m+n-2 > 100 {
		return 0 // 超出计算范围，返回0
	}

	numerator := int64(1)   // 分子
	denominator := int64(1) // 分母

	for i := 0; i < smaller; i++ {
		numerator *= int64(m + n - 2 - i)
		denominator *= int64(i + 1)

		// 提前约分防止溢出
		if numerator%denominator == 0 {
			numerator /= denominator
			denominator = 1
		}
	}

	return int(numerator / denominator)
}

// ========== 方法4: 记忆化递归 ==========
func uniquePaths4(m int, n int) int {
	memo := make(map[[2]int]int)
	return uniquePathsHelper4(m-1, n-1, memo)
}

func uniquePathsHelper4(i, j int, memo map[[2]int]int) int {
	// 边界条件
	if i == 0 || j == 0 {
		return 1
	}

	// 检查缓存
	key := [2]int{i, j}
	if val, exists := memo[key]; exists {
		return val
	}

	// 递归计算
	result := uniquePathsHelper4(i-1, j, memo) + uniquePathsHelper4(i, j-1, memo)
	memo[key] = result
	return result
}

// ========== 方法5: 优化的数学公式法 ==========
func uniquePaths5(m int, n int) int {
	if m == 1 || n == 1 {
		return 1
	}

	// 确保 m <= n，减少计算量
	if m > n {
		m, n = n, m
	}

	// 防止大数溢出
	if m+n-2 > 100 {
		return 0
	}

	// 计算 C(m+n-2, m-1)
	result := int64(1)
	for i := 0; i < m-1; i++ {
		result = result * int64(n+i) / int64(i+1)

		// 检查溢出
		if result < 0 {
			return 0
		}
	}

	return int(result)
}

// ========== 工具函数 ==========

// 打印路径矩阵（显示每个位置的路径数）
func printPathMatrix(m, n int) {
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}

	// 计算路径数
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 || j == 0 {
				dp[i][j] = 1
			} else {
				dp[i][j] = dp[i-1][j] + dp[i][j-1]
			}
		}
	}

	// 打印矩阵
	fmt.Printf("%d×%d网格的路径数矩阵:\n", m, n)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			fmt.Printf("%4d", dp[i][j])
		}
		fmt.Println()
	}
	fmt.Println()
}

// 验证组合数公式
func verifyFormula(m, n int) {
	fmt.Printf("验证组合数公式 C(%d, %d):\n", m+n-2, m-1)
	fmt.Printf("总步数: %d, 向下步数: %d, 向右步数: %d\n", m+n-2, m-1, n-1)

	// 计算组合数
	result := 1
	for i := 0; i < m-1; i++ {
		result = result * (m + n - 2 - i) / (i + 1)
		fmt.Printf("第%d步: %d\n", i+1, result)
	}
	fmt.Printf("最终结果: %d\n\n", result)
}

// 生成大型测试用例
func generateLargeTest() (int, int) {
	// 生成接近上限的测试用例
	return 50, 50
}

// ========== 测试和性能评估 ==========
func main() {
	// 测试用例
	testCases := []struct {
		name     string
		m        int
		n        int
		expected int
	}{
		{
			name:     "示例1: 3×7网格",
			m:        3,
			n:        7,
			expected: 28,
		},
		{
			name:     "示例2: 3×2网格",
			m:        3,
			n:        2,
			expected: 3,
		},
		{
			name:     "示例3: 7×3网格",
			m:        7,
			n:        3,
			expected: 28,
		},
		{
			name:     "示例4: 3×3网格",
			m:        3,
			n:        3,
			expected: 6,
		},
		{
			name:     "测试5: 1×1网格",
			m:        1,
			n:        1,
			expected: 1,
		},
		{
			name:     "测试6: 1×10网格",
			m:        1,
			n:        10,
			expected: 1,
		},
		{
			name:     "测试7: 10×1网格",
			m:        10,
			n:        1,
			expected: 1,
		},
		{
			name:     "测试8: 2×2网格",
			m:        2,
			n:        2,
			expected: 2,
		},
		{
			name:     "测试9: 4×4网格",
			m:        4,
			n:        4,
			expected: 20,
		},
		{
			name:     "测试10: 5×5网格",
			m:        5,
			n:        5,
			expected: 70,
		},
	}

	// 算法方法
	methods := []struct {
		name string
		fn   func(int, int) int
	}{
		{"二维动态规划", uniquePaths1},
		{"一维动态规划", uniquePaths2},
		{"数学公式法", uniquePaths3},
		{"记忆化递归", uniquePaths4},
		{"优化数学法", uniquePaths5},
	}

	fmt.Println("=== LeetCode 62. 不同路径 - 测试结果 ===")
	fmt.Println()

	// 运行测试
	for _, tc := range testCases {
		fmt.Printf("测试用例: %s\n", tc.name)
		fmt.Printf("网格大小: %d×%d\n", tc.m, tc.n)

		allPassed := true
		var results []int
		var times []time.Duration

		for _, method := range methods {
			start := time.Now()
			result := method.fn(tc.m, tc.n)
			elapsed := time.Since(start)

			results = append(results, result)
			times = append(times, elapsed)

			status := "✅"
			if result != tc.expected {
				status = "❌"
				allPassed = false
			}

			fmt.Printf("  %s: %s (结果: %d, 耗时: %v)\n", method.name, status, result, elapsed)
		}

		fmt.Printf("期望结果: %d\n", tc.expected)

		if allPassed {
			fmt.Println("✅ 所有方法均通过")
		} else {
			fmt.Println("❌ 存在失败的方法")
		}
		fmt.Println(strings.Repeat("-", 60))
	}

	// 路径矩阵演示
	fmt.Println("\n=== 路径矩阵演示 ===")
	printPathMatrix(4, 5)
	printPathMatrix(3, 3)

	// 组合数验证
	fmt.Println("=== 组合数公式验证 ===")
	verifyFormula(3, 7)
	verifyFormula(4, 4)

	// 性能对比测试
	fmt.Println("=== 性能对比测试 ===")
	performanceTest()

	// 算法特性总结
	fmt.Println("\n=== 算法特性总结 ===")
	fmt.Println("1. 二维动态规划:")
	fmt.Println("   - 时间复杂度: O(m×n)")
	fmt.Println("   - 空间复杂度: O(m×n)")
	fmt.Println("   - 特点: 最容易理解，直观清晰")
	fmt.Println()
	fmt.Println("2. 一维动态规划:")
	fmt.Println("   - 时间复杂度: O(m×n)")
	fmt.Println("   - 空间复杂度: O(min(m,n))")
	fmt.Println("   - 特点: 空间优化，滚动数组")
	fmt.Println()
	fmt.Println("3. 数学公式法:")
	fmt.Println("   - 时间复杂度: O(min(m,n))")
	fmt.Println("   - 空间复杂度: O(1)")
	fmt.Println("   - 特点: 最优解法，组合数学")
	fmt.Println()
	fmt.Println("4. 记忆化递归:")
	fmt.Println("   - 时间复杂度: O(m×n)")
	fmt.Println("   - 空间复杂度: O(m×n)")
	fmt.Println("   - 特点: 自顶向下，自然思维")
	fmt.Println()
	fmt.Println("5. 优化数学法:")
	fmt.Println("   - 时间复杂度: O(min(m,n))")
	fmt.Println("   - 空间复杂度: O(1)")
	fmt.Println("   - 特点: 防溢出优化，高精度")

	// 机器人路径演示
	fmt.Println("\n=== 机器人路径演示 ===")
	demoRobotPath()
}

// 性能测试
func performanceTest() {
	sizes := [][2]int{{10, 10}, {20, 20}, {30, 30}, {40, 40}}
	methods := []struct {
		name string
		fn   func(int, int) int
	}{
		{"二维DP", uniquePaths1},
		{"一维DP", uniquePaths2},
		{"数学公式", uniquePaths3},
		{"记忆化递归", uniquePaths4},
		{"优化数学", uniquePaths5},
	}

	for _, size := range sizes {
		m, n := size[0], size[1]
		fmt.Printf("性能测试 - 网格大小: %d×%d\n", m, n)

		for _, method := range methods {
			start := time.Now()
			result := method.fn(m, n)
			elapsed := time.Since(start)

			if result > 0 {
				fmt.Printf("  %s: 路径数=%d, 耗时=%v\n",
					method.name, result, elapsed)
			} else {
				fmt.Printf("  %s: 溢出或超限, 耗时=%v\n",
					method.name, elapsed)
			}
		}
		fmt.Println()
	}
}

// 机器人路径演示
func demoRobotPath() {
	m, n := 3, 4
	fmt.Printf("机器人在%d×%d网格中寻路:\n", m, n)
	fmt.Println("Start → → → Finish")
	fmt.Println("  ↓   ↓   ↓   ↓")
	fmt.Println("  ↓   ↓   ↓   ↓")
	fmt.Println()

	// 显示每个位置的路径数
	printPathMatrix(m, n)

	result := uniquePaths1(m, n)
	fmt.Printf("从左上角到右下角共有 %d 条不同路径\n", result)

	fmt.Println("示例路径:")
	fmt.Println("1. 右→右→右→下→下")
	fmt.Println("2. 右→右→下→右→下")
	fmt.Println("3. 右→下→右→右→下")
	fmt.Println("4. 下→右→右→右→下")
	fmt.Println("5. 下→下→右→右→右")
	fmt.Println("6. 右→下→右→下→右")
	fmt.Println("... 等等")
	fmt.Println("机器人寻路完成!")
}
