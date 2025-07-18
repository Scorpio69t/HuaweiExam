package main

import (
	"fmt"
	"math"
	"time"
)

// 方法1：动态规划（标准解法）
func climbStairs(n int) int {
	if n <= 2 {
		return n
	}

	dp := make([]int, n+1)
	dp[1] = 1
	dp[2] = 2

	for i := 3; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}

	return dp[n]
}

// 方法2：空间优化动态规划（最优解）
func climbStairsOptimized(n int) int {
	if n <= 2 {
		return n
	}

	prev2 := 1 // f(1)
	prev1 := 2 // f(2)

	for i := 3; i <= n; i++ {
		curr := prev1 + prev2
		prev2 = prev1
		prev1 = curr
	}

	return prev1
}

// 方法3：递归 + 记忆化
func climbStairsMemo(n int) int {
	memo := make(map[int]int)
	return climbStairsMemoHelper(n, memo)
}

func climbStairsMemoHelper(n int, memo map[int]int) int {
	if n <= 2 {
		return n
	}

	if val, exists := memo[n]; exists {
		return val
	}

	result := climbStairsMemoHelper(n-1, memo) + climbStairsMemoHelper(n-2, memo)
	memo[n] = result
	return result
}

// 方法4：数学公式（斐波那契通项公式）
func climbStairsFormula(n int) int {
	if n <= 2 {
		return n
	}

	sqrt5 := math.Sqrt(5)
	phi := (1 + sqrt5) / 2 // 黄金比例
	psi := (1 - sqrt5) / 2 // 共轭黄金比例

	// 斐波那契通项公式：F(n) = (φ^n - ψ^n) / √5
	// 这里是 F(n+1)，因为我们的序列是 f(1)=1, f(2)=2
	result := (math.Pow(phi, float64(n+1)) - math.Pow(psi, float64(n+1))) / sqrt5

	return int(math.Round(result))
}

// 方法5：矩阵快速幂
func climbStairsMatrix(n int) int {
	if n <= 2 {
		return n
	}

	// 标准斐波那契矩阵：[[1,1],[1,0]]
	// F(n) = [[1,1],[1,0]]^(n-1) 的第一行第一列
	// 爬楼梯问题：climbStairs(n) = F(n+1)
	
	base := [][]int{{1, 1}, {1, 0}}
	result := matrixPower(base, n)
	
	// result[0][0] = F(n+1), result[0][1] = F(n)
	return result[0][0]
}

func matrixPower(matrix [][]int, n int) [][]int {
	size := len(matrix)
	result := make([][]int, size)
	for i := range result {
		result[i] = make([]int, size)
		result[i][i] = 1 // 单位矩阵
	}

	base := make([][]int, size)
	for i := range base {
		base[i] = make([]int, size)
		copy(base[i], matrix[i])
	}

	for n > 0 {
		if n&1 == 1 {
			result = matrixMultiply(result, base)
		}
		base = matrixMultiply(base, base)
		n >>= 1
	}

	return result
}

func matrixMultiply(a, b [][]int) [][]int {
	size := len(a)
	result := make([][]int, size)
	for i := range result {
		result[i] = make([]int, size)
		for j := 0; j < size; j++ {
			for k := 0; k < size; k++ {
				result[i][j] += a[i][k] * b[k][j]
			}
		}
	}
	return result
}

// 测试函数
func testClimbingStairs() {
	fmt.Println("=== 70. 爬楼梯测试 ===")

	testCases := []struct {
		name     string
		n        int
		expected int
	}{
		// 基础测试用例
		{"示例1", 2, 2},
		{"示例2", 3, 3},
		{"示例3", 4, 5},
		{"示例4", 5, 8},

		// 边界测试用例
		{"最小值", 1, 1},
		{"小值测试", 6, 13},

		// 斐波那契验证
		{"斐波那契-10", 10, 89},
		{"斐波那契-15", 15, 987},
		{"斐波那契-20", 20, 10946},

		// 中等规模测试
		{"中等规模-25", 25, 121393},
		{"中等规模-30", 30, 1346269},

		// 大规模测试（接近题目限制）
		{"大规模-40", 40, 165580141},
		{"最大值-45", 45, 1836311903},
	}

	methods := []struct {
		name string
		fn   func(int) int
	}{
		{"标准DP", climbStairs},
		{"空间优化DP", climbStairsOptimized},
		{"记忆化递归", climbStairsMemo},
		{"数学公式", climbStairsFormula},
		{"矩阵快速幂", climbStairsMatrix},
	}

	for _, tc := range testCases {
		fmt.Printf("\n测试用例: %s (n=%d)\n", tc.name, tc.n)
		fmt.Printf("期望输出: %d\n", tc.expected)

		for _, method := range methods {
			start := time.Now()
			result := method.fn(tc.n)
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

	testSizes := []int{10, 20, 30, 40, 45}

	methods := []struct {
		name string
		fn   func(int) int
	}{
		{"标准DP", climbStairs},
		{"空间优化DP", climbStairsOptimized},
		{"记忆化递归", climbStairsMemo},
		{"数学公式", climbStairsFormula},
		{"矩阵快速幂", climbStairsMatrix},
	}

	for _, size := range testSizes {
		fmt.Printf("\n测试规模 n=%d:\n", size)

		for _, method := range methods {
			start := time.Now()
			result := method.fn(size)
			duration := time.Since(start)

			fmt.Printf("%s: 结果=%d, 耗时=%v\n",
				method.name, result, duration)
		}
	}
}

// 算法分析
func algorithmAnalysis() {
	fmt.Println("\n=== 算法分析 ===")

	fmt.Println("时间复杂度:")
	fmt.Println("  • 标准DP: O(n)")
	fmt.Println("  • 空间优化DP: O(n)")
	fmt.Println("  • 记忆化递归: O(n)")
	fmt.Println("  • 数学公式: O(1) ⭐ 最快")
	fmt.Println("  • 矩阵快速幂: O(log n)")

	fmt.Println("\n空间复杂度:")
	fmt.Println("  • 标准DP: O(n)")
	fmt.Println("  • 空间优化DP: O(1) ⭐ 最优实用")
	fmt.Println("  • 记忆化递归: O(n)")
	fmt.Println("  • 数学公式: O(1)")
	fmt.Println("  • 矩阵快速幂: O(1)")

	fmt.Println("\n核心思想:")
	fmt.Println("  1. 斐波那契数列本质：f(n) = f(n-1) + f(n-2)")
	fmt.Println("  2. 状态转移：到达第n阶只能从n-1或n-2阶到达")
	fmt.Println("  3. 空间优化：只需要保存前两个状态")
	fmt.Println("  4. 数学加速：利用黄金比例直接计算")

	fmt.Println("\n推荐使用:")
	fmt.Println("  • 面试/教学: 空间优化DP（平衡了效率和理解难度）")
	fmt.Println("  • 高性能场景: 数学公式（需要注意浮点精度）")
	fmt.Println("  • 大数场景: 矩阵快速幂（避免精度问题）")
}

// 斐波那契数列分析
func fibonacciAnalysis() {
	fmt.Println("\n=== 斐波那契数列分析 ===")

	fmt.Println("爬楼梯与斐波那契的关系:")
	fmt.Printf("%-10s %-15s %-15s\n", "n", "climbStairs(n)", "fibonacci(n+1)")
	fmt.Println(repeatString("-", 45))

	for i := 1; i <= 10; i++ {
		climb := climbStairsOptimized(i)
		fib := fibonacci(i + 1)
		fmt.Printf("%-10d %-15d %-15d\n", i, climb, fib)
	}

	fmt.Println("\n黄金比例验证:")
	for i := 5; i <= 15; i++ {
		fn := float64(climbStairsOptimized(i))
		fn1 := float64(climbStairsOptimized(i + 1))
		ratio := fn1 / fn
		phi := (1 + math.Sqrt(5)) / 2
		fmt.Printf("F(%d)/F(%d) = %.10f, φ = %.10f, 差值 = %.2e\n",
			i+1, i, ratio, phi, math.Abs(ratio-phi))
	}
}

// 辅助函数：计算斐波那契数列
func fibonacci(n int) int {
	if n <= 2 {
		return 1
	}
	a, b := 1, 1
	for i := 3; i <= n; i++ {
		a, b = b, a+b
	}
	return b
}

// 可视化演示
func visualDemo() {
	fmt.Println("\n=== 可视化演示 ===")

	n := 5
	fmt.Printf("示例: n = %d\n", n)

	fmt.Println("\n楼梯示意图:")
	for i := n; i >= 1; i-- {
		spaces := repeatString(" ", (n-i)*2)
		fmt.Printf("%s[%d]\n", spaces, i)
	}
	fmt.Println(repeatString(" ", n*2) + "[0] 起点")

	fmt.Println("\n状态转移过程:")
	fmt.Println("f(1) = 1 (一种方法: 爬1步)")
	fmt.Println("f(2) = 2 (两种方法: 1+1 或 2)")

	for i := 3; i <= n; i++ {
		prev1 := climbStairsOptimized(i - 1)
		prev2 := climbStairsOptimized(i - 2)
		curr := prev1 + prev2
		fmt.Printf("f(%d) = f(%d) + f(%d) = %d + %d = %d\n",
			i, i-1, i-2, prev1, prev2, curr)
	}

	fmt.Printf("\n最终答案: %d种方法\n", climbStairsOptimized(n))
}

// 辅助函数：重复字符串
func repeatString(s string, count int) string {
	if count <= 0 {
		return ""
	}
	result := ""
	for i := 0; i < count; i++ {
		result += s
	}
	return result
}

func main() {
	// 执行所有测试
	testClimbingStairs()
	performanceTest()
	algorithmAnalysis()
	fibonacciAnalysis()
	visualDemo()
}
