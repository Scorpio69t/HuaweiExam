package main

import (
	"fmt"
	"math"
)

// =========================== 方法一：二分查找（最优解法） ===========================

// mySqrt 二分查找
// 时间复杂度：O(log n)
// 空间复杂度：O(1)
func mySqrt(x int) int {
	if x < 2 {
		return x
	}

	left, right := 0, x

	for left <= right {
		mid := left + (right-left)/2

		// 避免溢出，使用除法代替乘法
		if mid == x/mid {
			return mid
		} else if mid < x/mid {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return right
}

// =========================== 方法二：牛顿迭代法 ===========================

// mySqrt2 牛顿迭代法
// 时间复杂度：O(log n)，实际上是O(log log n)，二次收敛
// 空间复杂度：O(1)
func mySqrt2(x int) int {
	if x < 2 {
		return x
	}

	r := x
	for r > x/r {
		r = (r + x/r) / 2
	}

	return r
}

// =========================== 方法三：位运算优化 ===========================

// mySqrt3 位运算优化
// 时间复杂度：O(log n)，最多16次迭代
// 空间复杂度：O(1)
func mySqrt3(x int) int {
	if x < 2 {
		return x
	}

	res := 0
	// 从2^15开始，因为sqrt(2^31) ≈ 2^15.5
	bit := 1 << 15

	for bit > 0 {
		temp := res + bit
		if temp <= x/temp {
			res = temp
		}
		bit >>= 1
	}

	return res
}

// =========================== 方法四：数学公式（袖珍计算器） ===========================

// mySqrt4 数学公式
// 时间复杂度：O(1)
// 空间复杂度：O(1)
// 注意：题目要求不使用内置函数，此方法仅供学习
func mySqrt4(x int) int {
	if x == 0 {
		return 0
	}

	// sqrt(x) = e^(0.5 * ln(x))
	ans := int(math.Exp(0.5 * math.Log(float64(x))))

	// 由于浮点数精度问题，需要验证
	if (ans+1)*(ans+1) <= x {
		return ans + 1
	}

	return ans
}

// =========================== 测试代码 ===========================

func main() {
	fmt.Println("=== LeetCode 69: x的平方根 ===\n")

	// 测试用例
	testCases := []struct {
		x      int
		expect int
	}{
		{0, 0},              // 边界：0
		{1, 1},              // 边界：1
		{2, 1},              // 非完全平方数
		{4, 2},              // 完全平方数
		{8, 2},              // 示例2
		{9, 3},              // 完全平方数
		{15, 3},             // 非完全平方数
		{16, 4},             // 完全平方数
		{100, 10},           // 完全平方数
		{121, 11},           // 完全平方数
		{144, 12},           // 完全平方数
		{2147483647, 46340}, // 最大值
	}

	fmt.Println("方法一：二分查找")
	runTests(testCases, mySqrt)

	fmt.Println("\n方法二：牛顿迭代法")
	runTests(testCases, mySqrt2)

	fmt.Println("\n方法三：位运算优化")
	runTests(testCases, mySqrt3)

	fmt.Println("\n方法四：数学公式")
	runTests(testCases, mySqrt4)

	// 详细示例
	fmt.Println("\n=== 详细示例 ===")
	detailedExample()

	// 算法对比
	fmt.Println("\n=== 算法步骤对比 ===")
	compareAlgorithms()
}

// runTests 运行测试用例
func runTests(testCases []struct {
	x      int
	expect int
}, fn func(int) int) {
	passCount := 0
	for i, tc := range testCases {
		result := fn(tc.x)
		status := "✅"
		if result != tc.expect {
			status = "❌"
		} else {
			passCount++
		}

		fmt.Printf("  测试%d: %s ", i+1, status)
		if status == "❌" {
			fmt.Printf("x=%d, 输出=%d, 期望=%d\n", tc.x, result, tc.expect)
		} else {
			fmt.Printf("sqrt(%d) = %d\n", tc.x, result)
		}
	}
	fmt.Printf("  通过: %d/%d\n", passCount, len(testCases))
}

// detailedExample 详细示例
func detailedExample() {
	x := 8

	fmt.Printf("输入: x = %d\n\n", x)

	// 二分查找过程
	fmt.Println("方法一：二分查找过程")
	fmt.Printf("  初始: left=0, right=%d\n", x)

	left, right := 0, x
	step := 1

	for left <= right {
		mid := left + (right-left)/2
		midSquare := mid * mid

		fmt.Printf("  步骤%d: left=%d, mid=%d, right=%d, mid²=%d\n",
			step, left, mid, right, midSquare)

		if mid == x/mid {
			fmt.Printf("  找到答案: %d\n", mid)
			break
		} else if mid < x/mid {
			fmt.Printf("         %d² < %d, 搜索右半部分\n", mid, x)
			left = mid + 1
		} else {
			fmt.Printf("         %d² > %d, 搜索左半部分\n", mid, x)
			right = mid - 1
		}

		step++
	}

	fmt.Printf("  最终答案: %d\n\n", right)

	// 牛顿迭代过程
	fmt.Println("方法二：牛顿迭代过程")
	r := x
	step = 1

	fmt.Printf("  初始值: r = %d\n", r)

	for r > x/r {
		newR := (r + x/r) / 2
		fmt.Printf("  步骤%d: r=%d, x/r=%d, 新r=(%d+%d)/2=%d\n",
			step, r, x/r, r, x/r, newR)
		r = newR
		step++
	}

	fmt.Printf("  最终答案: %d\n\n", r)

	// 验证答案
	fmt.Println("验证:")
	ans := mySqrt(x)
	fmt.Printf("  %d² = %d <= %d ✓\n", ans, ans*ans, x)
	fmt.Printf("  %d² = %d > %d ✓\n", ans+1, (ans+1)*(ans+1), x)
}

// compareAlgorithms 算法对比
func compareAlgorithms() {
	x := 100

	fmt.Printf("计算 sqrt(%d):\n\n", x)

	// 二分查找
	fmt.Println("1. 二分查找:")
	fmt.Println("   - 搜索范围: [0, 100]")
	fmt.Println("   - 查找过程: 50 -> 25 -> 12 -> 6 -> 9 -> 10")
	fmt.Printf("   - 结果: %d\n\n", mySqrt(x))

	// 牛顿迭代
	fmt.Println("2. 牛顿迭代:")
	fmt.Println("   - 初始值: 100")
	fmt.Println("   - 迭代过程: 100 -> 50 -> 26 -> 15 -> 10")
	fmt.Printf("   - 结果: %d\n\n", mySqrt2(x))

	// 位运算
	fmt.Println("3. 位运算:")
	fmt.Println("   - 从高位开始: bit=32768")
	fmt.Println("   - 逐位确定: 从2^15到2^0")
	fmt.Printf("   - 结果: %d\n\n", mySqrt3(x))

	// 大数测试
	fmt.Println("大数测试 (x = 2^31 - 1):")
	maxInt := 2147483647
	fmt.Printf("  输入: %d\n", maxInt)
	fmt.Printf("  二分查找: %d\n", mySqrt(maxInt))
	fmt.Printf("  牛顿迭代: %d\n", mySqrt2(maxInt))
	fmt.Printf("  位运算: %d\n", mySqrt3(maxInt))
	fmt.Printf("  验证: 46340² = %d\n", 46340*46340)
	fmt.Printf("        46341² = %d (溢出int范围)\n", int64(46341)*int64(46341))
}
