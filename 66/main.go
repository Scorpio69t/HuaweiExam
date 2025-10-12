package main

import (
	"fmt"
	"time"
)

// 方法一：逆序遍历算法（最优解法）
func plusOne1(digits []int) []int {
	n := len(digits)

	// 从末尾开始处理
	for i := n - 1; i >= 0; i-- {
		// 如果当前位不是9，直接+1返回
		if digits[i] < 9 {
			digits[i]++
			return digits
		}
		// 当前位是9，变成0，继续向前进位
		digits[i] = 0
	}

	// 所有位都是9，需要在开头插入1
	result := make([]int, n+1)
	result[0] = 1
	// 其余位默认为0，不需要赋值
	return result
}

// 方法二：使用进位标志算法
func plusOne2(digits []int) []int {
	carry := 1
	n := len(digits)

	for i := n - 1; i >= 0 && carry > 0; i-- {
		sum := digits[i] + carry
		digits[i] = sum % 10
		carry = sum / 10
	}

	// 如果还有进位，在开头插入
	if carry > 0 {
		result := make([]int, n+1)
		result[0] = carry
		copy(result[1:], digits)
		return result
	}

	return digits
}

// 方法三：递归实现算法
func plusOne3(digits []int) []int {
	return addHelper(digits, len(digits)-1, 1)
}

func addHelper(digits []int, index int, carry int) []int {
	// 递归终止条件
	if index < 0 {
		if carry > 0 {
			// 需要在开头插入进位
			result := make([]int, len(digits)+1)
			result[0] = carry
			copy(result[1:], digits)
			return result
		}
		return digits
	}

	sum := digits[index] + carry
	digits[index] = sum % 10
	newCarry := sum / 10

	// 递归处理前一位
	return addHelper(digits, index-1, newCarry)
}

// 方法四：提前判断优化算法
func plusOne4(digits []int) []int {
	n := len(digits)

	// 快速路径：末位不是9
	if digits[n-1] < 9 {
		digits[n-1]++
		return digits
	}

	// 检查是否全是9
	allNine := true
	for _, d := range digits {
		if d != 9 {
			allNine = false
			break
		}
	}

	if allNine {
		// 全是9，直接构建结果
		result := make([]int, n+1)
		result[0] = 1
		return result
	}

	// 正常进位处理
	for i := n - 1; i >= 0; i-- {
		if digits[i] < 9 {
			digits[i]++
			return digits
		}
		digits[i] = 0
	}

	return digits
}

// 辅助函数：比较两个数组是否相等
func arrayEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

// 辅助函数：复制数组
func copyArray(arr []int) []int {
	result := make([]int, len(arr))
	copy(result, arr)
	return result
}

// 辅助函数：打印数组
func printArray(arr []int) string {
	if len(arr) == 0 {
		return "[]"
	}
	result := "["
	for i, v := range arr {
		if i > 0 {
			result += ","
		}
		result += fmt.Sprintf("%d", v)
	}
	result += "]"
	return result
}

// 测试用例
func createTestCases() []struct {
	input    []int
	expected []int
	name     string
} {
	return []struct {
		input    []int
		expected []int
		name     string
	}{
		{[]int{1, 2, 3}, []int{1, 2, 4}, "示例1: 无进位"},
		{[]int{4, 3, 2, 1}, []int{4, 3, 2, 2}, "示例2: 无进位"},
		{[]int{9}, []int{1, 0}, "示例3: 单个9进位"},
		{[]int{0}, []int{1}, "边界1: 最小值"},
		{[]int{1, 2, 9}, []int{1, 3, 0}, "测试1: 单次进位"},
		{[]int{1, 9, 9}, []int{2, 0, 0}, "测试2: 连续进位"},
		{[]int{9, 9, 9}, []int{1, 0, 0, 0}, "测试3: 全部进位"},
		{[]int{9, 9}, []int{1, 0, 0}, "测试4: 两位全9"},
		{[]int{9, 0, 0}, []int{9, 0, 1}, "测试5: 开头是9"},
		{[]int{1, 0, 0}, []int{1, 0, 1}, "测试6: 末尾是0"},
		{[]int{8, 9, 9, 9}, []int{9, 0, 0, 0}, "测试7: 部分进位"},
		{[]int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}, []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 1}, "测试8: 长数组"},
	}
}

// 性能测试
func benchmarkAlgorithm(algorithm func([]int) []int, input []int, name string) {
	iterations := 10000
	start := time.Now()

	for i := 0; i < iterations; i++ {
		// 每次都复制输入，避免修改原数组
		inputCopy := copyArray(input)
		algorithm(inputCopy)
	}

	duration := time.Since(start)
	avgTime := duration.Nanoseconds() / int64(iterations)
	fmt.Printf("%s: 平均执行时间 %d 纳秒\n", name, avgTime)
}

func main() {
	fmt.Println("=== 66. 加一 ===")
	fmt.Println()

	testCases := createTestCases()
	algorithms := []struct {
		name string
		fn   func([]int) []int
	}{
		{"逆序遍历算法", plusOne1},
		{"进位标志算法", plusOne2},
		{"递归实现算法", plusOne3},
		{"提前判断算法", plusOne4},
	}

	// 正确性测试
	fmt.Println("=== 算法正确性测试 ===")
	passCount := 0
	failCount := 0

	for _, testCase := range testCases {
		// 为每个算法准备独立的输入副本
		results := make([][]int, len(algorithms))
		for i, algo := range algorithms {
			inputCopy := copyArray(testCase.input)
			results[i] = algo.fn(inputCopy)
		}

		// 检查所有算法结果是否一致
		allEqual := true
		for i := 1; i < len(results); i++ {
			if !arrayEqual(results[i], results[0]) {
				allEqual = false
				break
			}
		}

		// 验证结果是否正确
		isValid := arrayEqual(results[0], testCase.expected)

		if allEqual && isValid {
			passCount++
			fmt.Printf("✅ %s: %s + 1 = %s\n",
				testCase.name, printArray(testCase.input), printArray(results[0]))
		} else {
			failCount++
			fmt.Printf("❌ %s: %s\n", testCase.name, printArray(testCase.input))
			fmt.Printf("   预期: %s\n", printArray(testCase.expected))
			for i, algo := range algorithms {
				fmt.Printf("   %s: %s\n", algo.name, printArray(results[i]))
			}
		}
	}

	fmt.Println()
	fmt.Printf("测试统计: 通过 %d/%d, 失败 %d/%d\n",
		passCount, len(testCases), failCount, len(testCases))
	fmt.Println()

	// 性能测试
	fmt.Println("=== 性能测试 ===")
	perfTests := [][]int{
		{1, 2, 3},
		{1, 2, 9},
		{9, 9, 9},
		{9, 9, 9, 9, 9, 9, 9, 9, 9, 9},
	}

	for _, test := range perfTests {
		fmt.Printf("测试输入: %s\n", printArray(test))
		for _, algo := range algorithms {
			benchmarkAlgorithm(algo.fn, test, algo.name)
		}
		fmt.Println()
	}

	// 算法分析
	fmt.Println("=== 算法分析 ===")
	fmt.Println("加一问题的特点:")
	fmt.Println("1. 模拟手工加法进位")
	fmt.Println("2. 从低位到高位处理")
	fmt.Println("3. 进位传播直到无进位")
	fmt.Println("4. 全9情况需要扩展数组")
	fmt.Println()

	fmt.Println("=== 进位示例 ===")
	fmt.Println("无进位: [1,2,3] + 1 = [1,2,4]")
	fmt.Println("单次进位: [1,2,9] + 1 = [1,3,0]")
	fmt.Println("连续进位: [1,9,9] + 1 = [2,0,0]")
	fmt.Println("全部进位: [9,9,9] + 1 = [1,0,0,0]")
	fmt.Println()

	fmt.Println("=== 复杂度分析 ===")
	fmt.Println("时间复杂度:")
	fmt.Println("- 最好情况: O(1)，末位不是9")
	fmt.Println("- 最坏情况: O(n)，全是9")
	fmt.Println("- 平均情况: O(1)，大多数不需要完整遍历")
	fmt.Println()
	fmt.Println("空间复杂度:")
	fmt.Println("- 逆序遍历: O(1)，原地修改")
	fmt.Println("- 进位标志: O(1)，常数空间")
	fmt.Println("- 递归实现: O(n)，递归栈")
	fmt.Println("- 提前判断: O(1)，原地修改")
	fmt.Println()

	fmt.Println("=== 为什么平均是O(1) ===")
	fmt.Println("末位不是9的概率: 90% → O(1)")
	fmt.Println("连续2位是9的概率: 1% → O(2)")
	fmt.Println("连续3位是9的概率: 0.1% → O(3)")
	fmt.Println("连续n位是9的概率: 极低")
	fmt.Println("加权平均: 接近O(1)")
	fmt.Println()

	fmt.Println("=== 优化技巧总结 ===")
	fmt.Println("1. 逆序处理：从低位到高位")
	fmt.Println("2. 提前返回：遇到非9直接返回")
	fmt.Println("3. 原地修改：节省空间")
	fmt.Println("4. 边界处理：全9时扩展数组")
	fmt.Println("5. 进位标志：清晰表达进位逻辑")
	fmt.Println()

	fmt.Println("=== 应用场景 ===")
	fmt.Println("1. 大整数运算：超过int范围的数字")
	fmt.Println("2. 计算器实现：任意精度加法")
	fmt.Println("3. 版本号递增：1.2.9 → 1.3.0")
	fmt.Println("4. 编号系统：自动递增编号")
	fmt.Println("5. 密码学：大素数运算基础")
	fmt.Println()

	fmt.Println("推荐使用：逆序遍历算法（方法一），最简洁高效")
}
