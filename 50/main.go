package main

import (
	"fmt"
	"math"
	"time"
)

// 方法一：递归快速幂算法
// 最直观的解法，使用递归实现快速幂
func myPow1(x float64, n int) float64 {
	// 处理边界条件
	if n == 0 {
		return 1
	}
	if n < 0 {
		return 1 / myPow1(x, -n)
	}

	// 递归计算
	if n%2 == 1 {
		return x * myPow1(x*x, n/2)
	}
	return myPow1(x*x, n/2)
}

// 方法二：迭代快速幂算法
// 空间效率最高的解法，避免递归栈溢出
func myPow2(x float64, n int) float64 {
	// 处理边界条件
	if n == 0 {
		return 1
	}
	if n < 0 {
		x = 1 / x
		n = -n
	}

	result := 1.0
	base := x

	for n > 0 {
		if n&1 == 1 {
			result *= base
		}
		base *= base
		n >>= 1
	}

	return result
}

// 方法三：分治算法
// 经典分治思想，将问题分解为子问题
func myPow3(x float64, n int) float64 {
	// 处理边界条件
	if n == 0 {
		return 1
	}
	if n < 0 {
		return 1 / myPow3(x, -n)
	}

	// 分治计算
	half := myPow3(x, n/2)
	if n%2 == 0 {
		return half * half
	}
	return half * half * x
}

// 方法四：位运算算法
// 使用位运算优化，效率最高
func myPow4(x float64, n int) float64 {
	// 处理边界条件
	if n == 0 {
		return 1
	}
	if n < 0 {
		x = 1 / x
		n = -n
	}

	result := 1.0
	base := x

	for n > 0 {
		if n&1 == 1 {
			result *= base
		}
		base *= base
		n >>= 1
	}

	return result
}

// 辅助函数：创建测试用例
func createTestCases() []struct {
	x    float64
	n    int
	name string
} {
	return []struct {
		x    float64
		n    int
		name string
	}{
		{2.0, 10, "示例1: 2^10"},
		{2.1, 3, "示例2: 2.1^3"},
		{2.0, -2, "示例3: 2^(-2)"},
		{1.0, 0, "测试1: 1^0"},
		{0.0, 1, "测试2: 0^1"},
		{1.0, 1, "测试3: 1^1"},
		{2.0, 0, "测试4: 2^0"},
		{0.5, 2, "测试5: 0.5^2"},
		{2.0, -3, "测试6: 2^(-3)"},
		{1.5, 4, "测试7: 1.5^4"},
		{2.0, 20, "测试8: 2^20"},
		{0.1, 3, "测试9: 0.1^3"},
	}
}

// 性能测试函数
func benchmarkAlgorithm(algorithm func(float64, int) float64, x float64, n int, name string) {
	iterations := 1000
	start := time.Now()

	for i := 0; i < iterations; i++ {
		algorithm(x, n)
	}

	duration := time.Since(start)
	avgTime := duration.Nanoseconds() / int64(iterations)

	fmt.Printf("%s: 平均执行时间 %d 纳秒\n", name, avgTime)
}

// 辅助函数：验证结果是否正确
func validateResult(x float64, n int, result float64) bool {
	// 计算期望结果
	expected := math.Pow(x, float64(n))

	// 允许一定的浮点数误差
	epsilon := 1e-9
	return math.Abs(result-expected) < epsilon
}

// 辅助函数：比较两个浮点数是否相等
func isEqual(a, b float64) bool {
	epsilon := 1e-9
	return math.Abs(a-b) < epsilon
}

// 辅助函数：打印幂运算结果
func printPowResult(x float64, n int, result float64, title string) {
	fmt.Printf("%s: %.5f^%d = %.5f\n", title, x, n, result)
}

func main() {
	fmt.Println("=== 50. Pow(x, n) ===")
	fmt.Println()

	// 创建测试用例
	testCases := createTestCases()
	algorithms := []struct {
		name string
		fn   func(float64, int) float64
	}{
		{"递归快速幂算法", myPow1},
		{"迭代快速幂算法", myPow2},
		{"分治算法", myPow3},
		{"位运算算法", myPow4},
	}

	// 运行测试
	fmt.Println("=== 算法正确性测试 ===")
	for _, testCase := range testCases {
		fmt.Printf("测试: %s\n", testCase.name)

		results := make([]float64, len(algorithms))
		for i, algo := range algorithms {
			results[i] = algo.fn(testCase.x, testCase.n)
		}

		// 验证所有算法结果一致
		allEqual := true
		for i := 1; i < len(results); i++ {
			if !isEqual(results[i], results[0]) {
				allEqual = false
				break
			}
		}

		// 验证结果是否正确
		allValid := true
		for _, result := range results {
			if !validateResult(testCase.x, testCase.n, result) {
				allValid = false
				break
			}
		}

		if allEqual && allValid {
			fmt.Printf("  ✅ 所有算法结果一致且正确: %.5f\n", results[0])
			if testCase.n <= 10 {
				printPowResult(testCase.x, testCase.n, results[0], "  计算结果")
			}
		} else {
			fmt.Printf("  ❌ 算法结果不一致或错误\n")
			for i, algo := range algorithms {
				fmt.Printf("    %s: %.5f\n", algo.name, results[i])
			}
		}
		fmt.Println()
	}

	// 性能测试
	fmt.Println("=== 性能测试 ===")
	performanceX := 2.0
	performanceN := 1000

	fmt.Printf("测试数据: %.1f^%d\n", performanceX, performanceN)
	fmt.Println()

	for _, algo := range algorithms {
		benchmarkAlgorithm(algo.fn, performanceX, performanceN, algo.name)
	}
	fmt.Println()

	// 算法分析
	fmt.Println("=== 算法分析 ===")
	fmt.Println("Pow(x, n)问题的特点:")
	fmt.Println("1. 需要计算x的n次幂")
	fmt.Println("2. 指数n可能为负数")
	fmt.Println("3. 需要高效的算法避免超时")
	fmt.Println("4. 快速幂算法是最优解法")
	fmt.Println()

	// 复杂度分析
	fmt.Println("=== 复杂度分析 ===")
	fmt.Println("时间复杂度:")
	fmt.Println("- 递归快速幂: O(logn)，每次递归将指数减半")
	fmt.Println("- 迭代快速幂: O(logn)，循环次数等于指数的二进制位数")
	fmt.Println("- 分治算法: O(logn)，分治深度为logn")
	fmt.Println("- 位运算: O(logn)，遍历指数的每一位")
	fmt.Println()
	fmt.Println("空间复杂度:")
	fmt.Println("- 递归栈: O(logn)，递归深度最多为logn")
	fmt.Println("- 迭代算法: O(1)，只使用常数空间")
	fmt.Println("- 分治算法: O(logn)，递归调用栈")
	fmt.Println("- 位运算: O(1)，只使用常数空间")
	fmt.Println()

	// 算法总结
	fmt.Println("=== 算法总结 ===")
	fmt.Println("1. 递归快速幂算法：最直观的解法，逻辑清晰")
	fmt.Println("2. 迭代快速幂算法：空间效率最高，避免栈溢出")
	fmt.Println("3. 分治算法：经典分治思想，易于理解")
	fmt.Println("4. 位运算算法：使用位运算优化，效率最高")
	fmt.Println()
	fmt.Println("推荐使用：位运算算法（方法四），效率最高")
	fmt.Println()

	// 应用场景
	fmt.Println("=== 应用场景 ===")
	fmt.Println("- 数学计算：计算幂运算")
	fmt.Println("- 密码学：RSA加密算法中的模幂运算")
	fmt.Println("- 算法竞赛：快速幂的经典应用")
	fmt.Println("- 科学计算：数值计算中的幂运算")
	fmt.Println("- 图形学：3D变换中的幂运算")
	fmt.Println()

	// 优化技巧总结
	fmt.Println("=== 优化技巧总结 ===")
	fmt.Println("1. 快速幂：掌握快速幂的核心思想")
	fmt.Println("2. 分治算法：理解分治在幂运算中的应用")
	fmt.Println("3. 位运算：学会使用位运算优化")
	fmt.Println("4. 边界处理：正确处理各种边界情况")
	fmt.Println("5. 算法选择：根据问题特点选择合适的算法")
	fmt.Println("6. 优化策略：学会时间和空间优化技巧")
}
