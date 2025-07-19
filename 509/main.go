package main

import (
	"fmt"
	"time"
)

// 方法一：朴素递归（效率最低，仅用于对比）
// 时间复杂度：O(2^n)，空间复杂度：O(n)
func fibRecursive(n int) int {
	if n <= 1 {
		return n
	}
	return fibRecursive(n-1) + fibRecursive(n-2)
}

// 方法二：备忘录递归（记忆化递归）- 重点推荐
// 时间复杂度：O(n)，空间复杂度：O(n)
func fibMemo(n int) int {
	// 创建备忘录，初始化为-1表示未计算
	memo := make(map[int]int)

	var helper func(int) int
	helper = func(num int) int {
		// 检查备忘录中是否已有结果
		if val, exists := memo[num]; exists {
			return val
		}

		// 基础情况
		if num <= 1 {
			memo[num] = num
			return num
		}

		// 递归计算并存入备忘录
		result := helper(num-1) + helper(num-2)
		memo[num] = result
		return result
	}

	return helper(n)
}

// 方法三：备忘录递归（使用数组实现）
// 时间复杂度：O(n)，空间复杂度：O(n)
func fibMemoArray(n int) int {
	if n <= 1 {
		return n
	}

	// 创建备忘录数组，-1表示未计算
	memo := make([]int, n+1)
	for i := range memo {
		memo[i] = -1
	}

	var helper func(int) int
	helper = func(num int) int {
		// 检查备忘录
		if memo[num] != -1 {
			return memo[num]
		}

		// 基础情况
		if num <= 1 {
			memo[num] = num
			return num
		}

		// 递归计算并存储
		memo[num] = helper(num-1) + helper(num-2)
		return memo[num]
	}

	return helper(n)
}

// 方法四：动态规划（自底向上）
// 时间复杂度：O(n)，空间复杂度：O(n)
func fibDP(n int) int {
	if n <= 1 {
		return n
	}

	// 创建DP数组
	dp := make([]int, n+1)
	dp[0] = 0
	dp[1] = 1

	// 自底向上填充
	for i := 2; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}

	return dp[n]
}

// 方法五：空间优化的动态规划
// 时间复杂度：O(n)，空间复杂度：O(1)
func fibOptimized(n int) int {
	if n <= 1 {
		return n
	}

	prev1, prev2 := 1, 0

	for i := 2; i <= n; i++ {
		current := prev1 + prev2
		prev2 = prev1
		prev1 = current
	}

	return prev1
}

// 方法六：矩阵快速幂（最高效）
// 时间复杂度：O(log n)，空间复杂度：O(1)
func fibMatrix(n int) int {
	if n <= 1 {
		return n
	}

	// 矩阵 [[1,1],[1,0]]
	matrix := [][]int{{1, 1}, {1, 0}}
	result := matrixPower(matrix, n-1)

	return result[0][0]
}

// 矩阵快速幂辅助函数
func matrixPower(matrix [][]int, n int) [][]int {
	if n == 1 {
		return matrix
	}

	if n%2 == 0 {
		half := matrixPower(matrix, n/2)
		return matrixMultiply(half, half)
	} else {
		return matrixMultiply(matrix, matrixPower(matrix, n-1))
	}
}

// 矩阵乘法
func matrixMultiply(a, b [][]int) [][]int {
	return [][]int{
		{a[0][0]*b[0][0] + a[0][1]*b[1][0], a[0][0]*b[0][1] + a[0][1]*b[1][1]},
		{a[1][0]*b[0][0] + a[1][1]*b[1][0], a[1][0]*b[0][1] + a[1][1]*b[1][1]},
	}
}

// 性能测试函数
func measureTime(name string, fn func() int) int {
	start := time.Now()
	result := fn()
	duration := time.Since(start)
	fmt.Printf("%s: 结果=%d, 耗时=%v\n", name, result, duration)
	return result
}

// 运行所有测试用例
func runTests() {
	fmt.Println("=== 509. 斐波那契数 测试用例 ===")

	// 基础功能测试
	fmt.Println("\n--- 基础功能测试 ---")
	testCases := []int{0, 1, 2, 3, 4, 5, 10, 15, 20}

	for _, n := range testCases {
		fmt.Printf("\nF(%d):\n", n)

		// 所有方法都测试（小数值时）
		if n <= 20 {
			result1 := fibMemo(n)
			result2 := fibMemoArray(n)
			result3 := fibDP(n)
			result4 := fibOptimized(n)
			result5 := fibMatrix(n)

			fmt.Printf("  备忘录递归(map): %d\n", result1)
			fmt.Printf("  备忘录递归(array): %d\n", result2)
			fmt.Printf("  动态规划: %d\n", result3)
			fmt.Printf("  空间优化DP: %d\n", result4)
			fmt.Printf("  矩阵快速幂: %d\n", result5)

			// 验证结果一致性
			if result1 == result2 && result2 == result3 && result3 == result4 && result4 == result5 {
				fmt.Printf("  ✅ 所有方法结果一致\n")
			} else {
				fmt.Printf("  ❌ 结果不一致！\n")
			}
		}
	}

	// 性能对比测试
	fmt.Println("\n--- 性能对比测试 ---")

	// 小规模性能测试（包含朴素递归）
	fmt.Printf("\n小规模测试 F(25):\n")
	measureTime("朴素递归", func() int { return fibRecursive(25) })
	measureTime("备忘录递归(map)", func() int { return fibMemo(25) })
	measureTime("备忘录递归(array)", func() int { return fibMemoArray(25) })
	measureTime("动态规划", func() int { return fibDP(25) })
	measureTime("空间优化DP", func() int { return fibOptimized(25) })
	measureTime("矩阵快速幂", func() int { return fibMatrix(25) })

	// 中等规模性能测试（排除朴素递归）
	fmt.Printf("\n中等规模测试 F(40):\n")
	measureTime("备忘录递归(map)", func() int { return fibMemo(40) })
	measureTime("备忘录递归(array)", func() int { return fibMemoArray(40) })
	measureTime("动态规划", func() int { return fibDP(40) })
	measureTime("空间优化DP", func() int { return fibOptimized(40) })
	measureTime("矩阵快速幂", func() int { return fibMatrix(40) })

	// 大规模性能测试
	fmt.Printf("\n大规模测试 F(45):\n")
	measureTime("备忘录递归(map)", func() int { return fibMemo(45) })
	measureTime("备忘录递归(array)", func() int { return fibMemoArray(45) })
	measureTime("动态规划", func() int { return fibDP(45) })
	measureTime("空间优化DP", func() int { return fibOptimized(45) })
	measureTime("矩阵快速幂", func() int { return fibMatrix(45) })

	// 备忘录效果演示
	fmt.Println("\n--- 备忘录效果演示 ---")
	demonstrateMemoization()

	// 边界情况测试
	fmt.Println("\n--- 边界情况测试 ---")
	fmt.Printf("F(0) = %d (预期: 0)\n", fibOptimized(0))
	fmt.Printf("F(1) = %d (预期: 1)\n", fibOptimized(1))
	fmt.Printf("F(30) = %d\n", fibOptimized(30))
}

// 演示备忘录的效果
func demonstrateMemoization() {
	fmt.Println("备忘录递归 vs 朴素递归效果对比：")

	n := 30
	fmt.Printf("\n计算 F(%d):\n", n)

	// 朴素递归（会很慢）
	start := time.Now()
	result1 := fibRecursive(n)
	duration1 := time.Since(start)

	// 备忘录递归（很快）
	start = time.Now()
	result2 := fibMemo(n)
	duration2 := time.Since(start)

	fmt.Printf("朴素递归: 结果=%d, 耗时=%v\n", result1, duration1)
	fmt.Printf("备忘录递归: 结果=%d, 耗时=%v\n", result2, duration2)
	fmt.Printf("性能提升: %.2fx\n", float64(duration1.Nanoseconds())/float64(duration2.Nanoseconds()))
}

// 展示备忘录的工作原理
func demonstrateMemoWorkflow() {
	fmt.Println("\n--- 备忘录工作原理演示 ---")

	// 带调试信息的备忘录递归
	memo := make(map[int]int)
	callCount := 0

	var fibWithDebug func(int) int
	fibWithDebug = func(n int) int {
		callCount++
		fmt.Printf("调用 F(%d)", n)

		if val, exists := memo[n]; exists {
			fmt.Printf(" -> 从备忘录返回: %d\n", val)
			return val
		}

		if n <= 1 {
			fmt.Printf(" -> 基础情况: %d\n", n)
			memo[n] = n
			return n
		}

		fmt.Printf(" -> 需要计算\n")
		result := fibWithDebug(n-1) + fibWithDebug(n-2)
		memo[n] = result
		fmt.Printf("F(%d) = %d (已存入备忘录)\n", n, result)
		return result
	}

	fmt.Println("计算 F(5) 的过程：")
	result := fibWithDebug(5)
	fmt.Printf("\n最终结果: F(5) = %d\n", result)
	fmt.Printf("总函数调用次数: %d\n", callCount)
	fmt.Printf("备忘录内容: %v\n", memo)
}

func main() {
	runTests()
	demonstrateMemoWorkflow()
}
