package main

import (
	"fmt"
	"time"
)

// 方法1：动态规划（标准解法）
func rob(nums []int) int {
	n := len(nums)
	if n == 1 {
		return nums[0]
	}
	if n == 2 {
		return max(nums[0], nums[1])
	}

	// 情况1：偷第一间，不偷最后一间 [0, n-2]
	case1 := robLinear(nums[:n-1])

	// 情况2：不偷第一间，可偷最后一间 [1, n-1]
	case2 := robLinear(nums[1:])

	return max(case1, case2)
}

func robLinear(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	if n == 1 {
		return nums[0]
	}

	dp := make([]int, n)
	dp[0] = nums[0]
	dp[1] = max(nums[0], nums[1])

	for i := 2; i < n; i++ {
		dp[i] = max(dp[i-1], dp[i-2]+nums[i])
	}

	return dp[n-1]
}

// 方法2：空间优化动态规划（最优解）
func robOptimized(nums []int) int {
	n := len(nums)
	if n == 1 {
		return nums[0]
	}
	if n == 2 {
		return max(nums[0], nums[1])
	}

	case1 := robLinearOptimized(nums[:n-1])
	case2 := robLinearOptimized(nums[1:])

	return max(case1, case2)
}

func robLinearOptimized(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	if n == 1 {
		return nums[0]
	}

	prev2 := nums[0]
	prev1 := max(nums[0], nums[1])

	for i := 2; i < n; i++ {
		curr := max(prev1, prev2+nums[i])
		prev2 = prev1
		prev1 = curr
	}

	return prev1
}

// 方法3：记忆化递归
func robMemo(nums []int) int {
	n := len(nums)
	if n == 1 {
		return nums[0]
	}
	if n == 2 {
		return max(nums[0], nums[1])
	}

	// 情况1：偷第一间，不偷最后一间
	memo1 := make(map[int]int)
	case1 := robMemoHelper(nums[:n-1], 0, memo1)

	// 情况2：不偷第一间，可偷最后一间
	memo2 := make(map[int]int)
	case2 := robMemoHelper(nums[1:], 0, memo2)

	return max(case1, case2)
}

func robMemoHelper(nums []int, index int, memo map[int]int) int {
	if index >= len(nums) {
		return 0
	}
	if val, exists := memo[index]; exists {
		return val
	}

	// 偷当前房屋
	robCurrent := nums[index] + robMemoHelper(nums, index+2, memo)
	// 不偷当前房屋
	notRob := robMemoHelper(nums, index+1, memo)

	result := max(robCurrent, notRob)
	memo[index] = result
	return result
}

// 方法4：状态机模拟
func robStateMachine(nums []int) int {
	n := len(nums)
	if n == 1 {
		return nums[0]
	}
	if n == 2 {
		return max(nums[0], nums[1])
	}

	case1 := robStateMachineHelper(nums[:n-1])
	case2 := robStateMachineHelper(nums[1:])

	return max(case1, case2)
}

func robStateMachineHelper(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	if n == 1 {
		return nums[0]
	}

	// 状态：[当前位置是否偷窃][累计金额]
	// true: 偷了当前房屋
	// false: 没偷当前房屋

	robbed := nums[0]    // 偷了第0间房屋的最大金额
	notRobbed := 0       // 没偷第0间房屋的最大金额

	for i := 1; i < n; i++ {
		newRobbed := notRobbed + nums[i]           // 偷当前房屋
		newNotRobbed := max(robbed, notRobbed)     // 不偷当前房屋

		robbed = newRobbed
		notRobbed = newNotRobbed
	}

	return max(robbed, notRobbed)
}

// 工具函数
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 测试函数
func testRobberII() {
	fmt.Println("=== 213. 打家劫舍 II 测试 ===")
	
	testCases := []struct {
		name     string
		nums     []int
		expected int
	}{
		// 基础测试用例
		{"示例1", []int{2, 3, 2}, 3},
		{"示例2", []int{1, 2, 3, 1}, 4},
		{"示例3", []int{1, 2, 3}, 3},
		
		// 边界测试用例
		{"单个房屋", []int{5}, 5},
		{"两个房屋-递增", []int{1, 2}, 2},
		{"两个房屋-递减", []int{2, 1}, 2},
		
		// 极值测试
		{"全相同", []int{1, 1, 1, 1}, 2},
		{"大数值", []int{1000, 1, 1, 1000}, 1001}, // 环形约束下不能同时取第一间和最后一间
		{"全零", []int{0, 0, 0}, 0},
		
		// 特殊模式
		{"递增序列", []int{1, 2, 3, 4, 5}, 8}, // 环形约束下最优解为2+4+2或3+5，实际为8
		{"递减序列", []int{5, 4, 3, 2, 1}, 8},
		{"交替模式", []int{100, 1, 100, 1, 100}, 200}, // 环形约束下不能取所有100
		{"峰值模式", []int{1, 5, 1, 5, 1}, 10},
	}

	methods := []struct {
		name string
		fn   func([]int) int
	}{
		{"标准DP", rob},
		{"空间优化DP", robOptimized},
		{"记忆化递归", robMemo},
		{"状态机", robStateMachine},
	}

	for _, tc := range testCases {
		fmt.Printf("\n测试用例: %s\n", tc.name)
		fmt.Printf("输入: %v\n", tc.nums)
		fmt.Printf("期望输出: %d\n", tc.expected)
		
		for _, method := range methods {
			start := time.Now()
			result := method.fn(tc.nums)
			duration := time.Since(start)
			
			status := "✓"
			if result != tc.expected {
				status = "✗"
			}
			
			fmt.Printf("%s %s: %d (耗时: %v)\n", 
				status, method.name, result, duration)
		}
	}
}

// 性能测试
func performanceTest() {
	fmt.Println("\n=== 性能测试 ===")
	
	// 生成大规模测试数据
	largeNums := make([]int, 100)
	for i := 0; i < 100; i++ {
		largeNums[i] = i + 1
	}
	
	methods := []struct {
		name string
		fn   func([]int) int
	}{
		{"标准DP", rob},
		{"空间优化DP", robOptimized},
		{"记忆化递归", robMemo},
		{"状态机", robStateMachine},
	}
	
	fmt.Printf("测试数据规模: %d\n", len(largeNums))
	
	for _, method := range methods {
		start := time.Now()
		result := method.fn(largeNums)
		duration := time.Since(start)
		
		fmt.Printf("%s: 结果=%d, 耗时=%v\n", 
			method.name, result, duration)
	}
}

// 算法分析
func algorithmAnalysis() {
	fmt.Println("\n=== 算法分析 ===")
	
	fmt.Println("时间复杂度:")
	fmt.Println("  • 标准DP: O(n)")
	fmt.Println("  • 空间优化DP: O(n)")
	fmt.Println("  • 记忆化递归: O(n)")
	fmt.Println("  • 状态机: O(n)")
	
	fmt.Println("\n空间复杂度:")
	fmt.Println("  • 标准DP: O(n)")
	fmt.Println("  • 空间优化DP: O(1) ⭐ 最优")
	fmt.Println("  • 记忆化递归: O(n)")
	fmt.Println("  • 状态机: O(1)")
	
	fmt.Println("\n核心思想:")
	fmt.Println("  1. 将环形问题分解为两个线性子问题")
	fmt.Println("  2. 情况1: 包含第一间房屋，排除最后一间")
	fmt.Println("  3. 情况2: 排除第一间房屋，包含最后一间")
	fmt.Println("  4. 两种情况取最大值")
	
	fmt.Println("\n状态转移方程:")
	fmt.Println("  dp[i] = max(dp[i-1], dp[i-2] + nums[i])")
}

// 可视化演示
func visualDemo() {
	fmt.Println("\n=== 可视化演示 ===")
	
	nums := []int{2, 3, 2}
	fmt.Printf("示例: %v\n", nums)
	
	fmt.Println("\n房屋布局 (环形):")
	fmt.Println("    [2]")
	fmt.Println("   /   \\")
	fmt.Println(" [3] - [2]")
	
	fmt.Println("\n分解为两个子问题:")
	fmt.Println("情况1: [2, 3] (包含第一间，排除最后一间)")
	fmt.Printf("      线性DP结果: %d\n", robLinear(nums[:2]))
	
	fmt.Println("情况2: [3, 2] (排除第一间，包含最后一间)")
	fmt.Printf("      线性DP结果: %d\n", robLinear(nums[1:]))
	
	fmt.Printf("最终结果: max(%d, %d) = %d\n", 
		robLinear(nums[:2]), robLinear(nums[1:]), rob(nums))
}

func main() {
	// 执行所有测试
	testRobberII()
	performanceTest()
	algorithmAnalysis()
	visualDemo()
}
