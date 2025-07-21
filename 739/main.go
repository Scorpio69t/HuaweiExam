package main

import "fmt"

// 方法1：单调栈（推荐）
// 时间复杂度：O(n)，空间复杂度：O(n)
func dailyTemperatures(temperatures []int) []int {
	n := len(temperatures)
	result := make([]int, n)
	stack := []int{} // 存储索引

	for i := 0; i < n; i++ {
		// 当栈非空且当前温度高于栈顶温度时
		for len(stack) > 0 && temperatures[i] > temperatures[stack[len(stack)-1]] {
			// 弹出栈顶索引
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			// 计算天数差
			result[top] = i - top
		}

		// 将当前索引入栈
		stack = append(stack, i)
	}

	return result
}

// 方法2：暴力枚举
// 时间复杂度：O(n²)，空间复杂度：O(1)
func dailyTemperaturesBruteForce(temperatures []int) []int {
	n := len(temperatures)
	result := make([]int, n)

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if temperatures[j] > temperatures[i] {
				result[i] = j - i
				break
			}
		}
	}

	return result
}

// 方法3：从右向左遍历
// 时间复杂度：O(n)，空间复杂度：O(n)
func dailyTemperaturesRightToLeft(temperatures []int) []int {
	n := len(temperatures)
	result := make([]int, n)
	stack := []int{}

	for i := n - 1; i >= 0; i-- {
		// 弹出栈中所有小于等于当前温度的元素
		for len(stack) > 0 && temperatures[stack[len(stack)-1]] <= temperatures[i] {
			stack = stack[:len(stack)-1]
		}

		// 如果栈非空，栈顶即为下一个更高温度的位置
		if len(stack) > 0 {
			result[i] = stack[len(stack)-1] - i
		} else {
			result[i] = 0
		}

		// 将当前索引入栈
		stack = append(stack, i)
	}

	return result
}

// 方法4：动态规划优化
// 时间复杂度：O(n)，空间复杂度：O(1)
func dailyTemperaturesDP(temperatures []int) []int {
	n := len(temperatures)
	result := make([]int, n)

	for i := n - 2; i >= 0; i-- {
		j := i + 1

		// 跳跃式查找
		for j < n && temperatures[j] <= temperatures[i] {
			if result[j] == 0 {
				// 如果j后面没有更高温度，i也没有
				j = n
				break
			}
			j += result[j]
		}

		if j < n {
			result[i] = j - i
		}
	}

	return result
}

// 工具函数：比较两个数组是否相等
func equal(a, b []int) bool {
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

// 打印数组辅助函数
func printArray(arr []int, name string) {
	fmt.Printf("%s: [", name)
	for i, val := range arr {
		if i > 0 {
			fmt.Printf(", ")
		}
		fmt.Printf("%d", val)
	}
	fmt.Printf("]\n")
}

// 测试函数
func runTests() {
	fmt.Println("=== 739. 每日温度 测试 ===")

	// 测试用例1
	temperatures1 := []int{73, 74, 75, 71, 69, 72, 76, 73}
	expected1 := []int{1, 1, 4, 2, 1, 1, 0, 0}

	fmt.Printf("\n测试用例1:\n")
	printArray(temperatures1, "输入温度")
	printArray(expected1, "期望结果")

	result1_1 := dailyTemperatures(temperatures1)
	result1_2 := dailyTemperaturesBruteForce(temperatures1)
	result1_3 := dailyTemperaturesRightToLeft(temperatures1)
	result1_4 := dailyTemperaturesDP(temperatures1)

	fmt.Printf("方法一（单调栈）: ")
	printArray(result1_1, "")
	fmt.Printf("方法二（暴力枚举）: ")
	printArray(result1_2, "")
	fmt.Printf("方法三（从右向左）: ")
	printArray(result1_3, "")
	fmt.Printf("方法四（动态规划）: ")
	printArray(result1_4, "")

	// 测试用例2
	temperatures2 := []int{30, 40, 50, 60}
	expected2 := []int{1, 1, 1, 0}

	fmt.Printf("\n测试用例2:\n")
	printArray(temperatures2, "输入温度")
	printArray(expected2, "期望结果")

	result2_1 := dailyTemperatures(temperatures2)
	result2_2 := dailyTemperaturesBruteForce(temperatures2)
	result2_3 := dailyTemperaturesRightToLeft(temperatures2)
	result2_4 := dailyTemperaturesDP(temperatures2)

	fmt.Printf("方法一（单调栈）: ")
	printArray(result2_1, "")
	fmt.Printf("方法二（暴力枚举）: ")
	printArray(result2_2, "")
	fmt.Printf("方法三（从右向左）: ")
	printArray(result2_3, "")
	fmt.Printf("方法四（动态规划）: ")
	printArray(result2_4, "")

	// 测试用例3
	temperatures3 := []int{30, 60, 90}
	expected3 := []int{1, 1, 0}

	fmt.Printf("\n测试用例3:\n")
	printArray(temperatures3, "输入温度")
	printArray(expected3, "期望结果")

	result3_1 := dailyTemperatures(temperatures3)
	result3_2 := dailyTemperaturesBruteForce(temperatures3)
	result3_3 := dailyTemperaturesRightToLeft(temperatures3)
	result3_4 := dailyTemperaturesDP(temperatures3)

	fmt.Printf("方法一（单调栈）: ")
	printArray(result3_1, "")
	fmt.Printf("方法二（暴力枚举）: ")
	printArray(result3_2, "")
	fmt.Printf("方法三（从右向左）: ")
	printArray(result3_3, "")
	fmt.Printf("方法四（动态规划）: ")
	printArray(result3_4, "")

	// 测试用例4：边界情况
	temperatures4 := []int{100}
	expected4 := []int{0}

	fmt.Printf("\n测试用例4（单个温度）:\n")
	printArray(temperatures4, "输入温度")
	printArray(expected4, "期望结果")

	result4 := dailyTemperatures(temperatures4)
	fmt.Printf("结果: ")
	printArray(result4, "")

	// 测试用例5：相同温度
	temperatures5 := []int{30, 30, 30}
	expected5 := []int{0, 0, 0}

	fmt.Printf("\n测试用例5（相同温度）:\n")
	printArray(temperatures5, "输入温度")
	printArray(expected5, "期望结果")

	result5 := dailyTemperatures(temperatures5)
	fmt.Printf("结果: ")
	printArray(result5, "")

	// 测试用例6：极值测试
	temperatures6 := []int{30, 100, 30, 100}
	expected6 := []int{1, 0, 1, 0}

	fmt.Printf("\n测试用例6（极值测试）:\n")
	printArray(temperatures6, "输入温度")
	printArray(expected6, "期望结果")

	result6 := dailyTemperatures(temperatures6)
	fmt.Printf("结果: ")
	printArray(result6, "")

	// 测试用例7：递增序列
	temperatures7 := []int{30, 40, 50, 60, 70}
	expected7 := []int{1, 1, 1, 1, 0}

	fmt.Printf("\n测试用例7（递增序列）:\n")
	printArray(temperatures7, "输入温度")
	printArray(expected7, "期望结果")

	result7 := dailyTemperatures(temperatures7)
	fmt.Printf("结果: ")
	printArray(result7, "")

	// 测试用例8：递减序列
	temperatures8 := []int{70, 60, 50, 40, 30}
	expected8 := []int{0, 0, 0, 0, 0}

	fmt.Printf("\n测试用例8（递减序列）:\n")
	printArray(temperatures8, "输入温度")
	printArray(expected8, "期望结果")

	result8 := dailyTemperatures(temperatures8)
	fmt.Printf("结果: ")
	printArray(result8, "")

	// 测试用例9：复杂情况
	temperatures9 := []int{34, 80, 80, 34, 34, 80, 80, 80, 80, 34}
	expected9 := []int{1, 0, 0, 2, 1, 0, 0, 0, 0, 0}

	fmt.Printf("\n测试用例9（复杂情况）:\n")
	printArray(temperatures9, "输入温度")
	printArray(expected9, "期望结果")

	result9 := dailyTemperatures(temperatures9)
	fmt.Printf("结果: ")
	printArray(result9, "")

	// 详细分析示例
	fmt.Printf("\n=== 详细分析示例 ===\n")
	analyzeExample(temperatures1)

	fmt.Printf("\n=== 算法复杂度对比 ===\n")
	fmt.Printf("方法一（单调栈）：     时间 O(n),      空间 O(n)     - 推荐\n")
	fmt.Printf("方法二（暴力枚举）：   时间 O(n²),     空间 O(1)     - 简单易懂\n")
	fmt.Printf("方法三（从右向左）：   时间 O(n),      空间 O(n)     - 等价算法\n")
	fmt.Printf("方法四（动态规划）：   时间 O(n),      空间 O(1)     - 空间最优\n")
}

// 详细分析示例
func analyzeExample(temperatures []int) {
	fmt.Printf("分析温度数组: ")
	printArray(temperatures, "")
	fmt.Printf("索引位置: [")
	for i := range temperatures {
		if i > 0 {
			fmt.Printf(", ")
		}
		fmt.Printf("%d", i)
	}
	fmt.Printf("]\n")

	// 使用单调栈分析过程
	stack := []int{}
	result := make([]int, len(temperatures))
	fmt.Printf("\n单调栈处理过程:\n")

	for i := 0; i < len(temperatures); i++ {
		fmt.Printf("步骤 %d: 当前温度=%d, 栈=%v ", i, temperatures[i], stack)

		// 当栈非空且当前温度高于栈顶温度时
		for len(stack) > 0 && temperatures[i] > temperatures[stack[len(stack)-1]] {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			result[top] = i - top
			fmt.Printf("弹出索引%d(温度%d), result[%d]=%d ", top, temperatures[top], top, result[top])
		}

		stack = append(stack, i)
		fmt.Printf("入栈索引%d, 新栈=%v\n", i, stack)
	}

	fmt.Printf("最终结果: ")
	printArray(result, "")
}

func main() {
	runTests()
}
