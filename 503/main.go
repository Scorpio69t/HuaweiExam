package main

import "fmt"

// 方法1：单调栈 + 数组复制（推荐）
// 时间复杂度：O(n)，空间复杂度：O(n)
func nextGreaterElements(nums []int) []int {
	n := len(nums)
	result := make([]int, n)
	for i := range result {
		result[i] = -1
	}

	// 创建2n长度的数组
	extended := make([]int, 2*n)
	copy(extended, nums)
	copy(extended[n:], nums)

	stack := []int{} // 存储索引

	for i := 0; i < 2*n; i++ {
		// 当栈非空且当前元素大于栈顶元素时
		for len(stack) > 0 && extended[i] > extended[stack[len(stack)-1]] {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			// 只更新原数组范围内的结果
			if top < n {
				result[top] = extended[i]
			}
		}

		// 将当前索引入栈
		stack = append(stack, i)
	}

	return result
}

// 方法2：单调栈 + 两遍遍历
// 时间复杂度：O(n)，空间复杂度：O(n)
func nextGreaterElementsTwoPass(nums []int) []int {
	n := len(nums)
	result := make([]int, n)
	for i := range result {
		result[i] = -1
	}

	stack := []int{} // 存储索引

	// 遍历两遍数组
	for i := 0; i < 2*n; i++ {
		// 获取实际索引
		actualIndex := i % n

		// 当栈非空且当前元素大于栈顶元素时
		for len(stack) > 0 && nums[actualIndex] > nums[stack[len(stack)-1]] {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			result[top] = nums[actualIndex]
		}

		// 只在第一遍遍历时入栈
		if i < n {
			stack = append(stack, i)
		}
	}

	return result
}

// 方法3：单调栈 + 索引映射
// 时间复杂度：O(n)，空间复杂度：O(n)
func nextGreaterElementsModulo(nums []int) []int {
	n := len(nums)
	result := make([]int, n)
	for i := range result {
		result[i] = -1
	}

	stack := []int{}

	for i := 0; i < 2*n; i++ {
		actualIndex := i % n

		for len(stack) > 0 && nums[actualIndex] > nums[stack[len(stack)-1]] {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			result[top] = nums[actualIndex]
		}

		// 只在第一遍遍历时入栈
		if i < n {
			stack = append(stack, i)
		}
	}

	return result
}

// 方法4：暴力枚举
// 时间复杂度：O(n²)，空间复杂度：O(1)
func nextGreaterElementsBruteForce(nums []int) []int {
	n := len(nums)
	result := make([]int, n)

	for i := 0; i < n; i++ {
		result[i] = -1

		// 从i+1开始查找
		for j := 1; j < n; j++ {
			index := (i + j) % n
			if nums[index] > nums[i] {
				result[i] = nums[index]
				break
			}
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
	fmt.Println("=== 503. 下一个更大元素 II 测试 ===")

	// 测试用例1
	nums1 := []int{1, 2, 1}
	expected1 := []int{2, -1, 2}

	fmt.Printf("\n测试用例1:\n")
	printArray(nums1, "输入数组")
	printArray(expected1, "期望结果")

	result1_1 := nextGreaterElements(nums1)
	result1_2 := nextGreaterElementsTwoPass(nums1)
	result1_3 := nextGreaterElementsModulo(nums1)
	result1_4 := nextGreaterElementsBruteForce(nums1)

	fmt.Printf("方法一（单调栈+复制）: ")
	printArray(result1_1, "")
	fmt.Printf("方法二（单调栈+两遍）: ")
	printArray(result1_2, "")
	fmt.Printf("方法三（单调栈+映射）: ")
	printArray(result1_3, "")
	fmt.Printf("方法四（暴力枚举）: ")
	printArray(result1_4, "")

	// 测试用例2
	nums2 := []int{1, 2, 3, 4, 3}
	expected2 := []int{2, 3, 4, -1, 4}

	fmt.Printf("\n测试用例2:\n")
	printArray(nums2, "输入数组")
	printArray(expected2, "期望结果")

	result2_1 := nextGreaterElements(nums2)
	result2_2 := nextGreaterElementsTwoPass(nums2)
	result2_3 := nextGreaterElementsModulo(nums2)
	result2_4 := nextGreaterElementsBruteForce(nums2)

	fmt.Printf("方法一（单调栈+复制）: ")
	printArray(result2_1, "")
	fmt.Printf("方法二（单调栈+两遍）: ")
	printArray(result2_2, "")
	fmt.Printf("方法三（单调栈+映射）: ")
	printArray(result2_3, "")
	fmt.Printf("方法四（暴力枚举）: ")
	printArray(result2_4, "")

	// 测试用例3：边界情况
	nums3 := []int{1}
	expected3 := []int{-1}

	fmt.Printf("\n测试用例3（单个元素）:\n")
	printArray(nums3, "输入数组")
	printArray(expected3, "期望结果")

	result3 := nextGreaterElements(nums3)
	fmt.Printf("结果: ")
	printArray(result3, "")

	// 测试用例4：相同元素
	nums4 := []int{1, 1, 1}
	expected4 := []int{-1, -1, -1}

	fmt.Printf("\n测试用例4（相同元素）:\n")
	printArray(nums4, "输入数组")
	printArray(expected4, "期望结果")

	result4 := nextGreaterElements(nums4)
	fmt.Printf("结果: ")
	printArray(result4, "")

	// 测试用例5：递减序列
	nums5 := []int{5, 4, 3, 2, 1}
	expected5 := []int{-1, 5, 5, 5, 5}

	fmt.Printf("\n测试用例5（递减序列）:\n")
	printArray(nums5, "输入数组")
	printArray(expected5, "期望结果")

	result5 := nextGreaterElements(nums5)
	fmt.Printf("结果: ")
	printArray(result5, "")

	// 测试用例6：递增序列
	nums6 := []int{1, 2, 3, 4, 5}
	expected6 := []int{2, 3, 4, 5, -1}

	fmt.Printf("\n测试用例6（递增序列）:\n")
	printArray(nums6, "输入数组")
	printArray(expected6, "期望结果")

	result6 := nextGreaterElements(nums6)
	fmt.Printf("结果: ")
	printArray(result6, "")

	// 测试用例7：复杂情况
	nums7 := []int{3, 8, 4, 1, 2}
	expected7 := []int{8, -1, 8, 2, 3}

	fmt.Printf("\n测试用例7（复杂情况）:\n")
	printArray(nums7, "输入数组")
	printArray(expected7, "期望结果")

	result7 := nextGreaterElements(nums7)
	fmt.Printf("结果: ")
	printArray(result7, "")

	// 测试用例8：重复元素
	nums8 := []int{1, 2, 1, 2, 1}
	expected8 := []int{2, -1, 2, -1, 2}

	fmt.Printf("\n测试用例8（重复元素）:\n")
	printArray(nums8, "输入数组")
	printArray(expected8, "期望结果")

	result8 := nextGreaterElements(nums8)
	fmt.Printf("结果: ")
	printArray(result8, "")

	// 测试用例9：全相同元素
	nums9 := []int{5, 5, 5, 5}
	expected9 := []int{-1, -1, -1, -1}

	fmt.Printf("\n测试用例9（全相同元素）:\n")
	printArray(nums9, "输入数组")
	printArray(expected9, "期望结果")

	result9 := nextGreaterElements(nums9)
	fmt.Printf("结果: ")
	printArray(result9, "")

	// 详细分析示例
	fmt.Printf("\n=== 详细分析示例 ===\n")
	analyzeExample(nums1)

	fmt.Printf("\n=== 算法复杂度对比 ===\n")
	fmt.Printf("方法一（单调栈+复制）：   时间 O(n),      空间 O(n)     - 推荐\n")
	fmt.Printf("方法二（单调栈+两遍）：   时间 O(n),      空间 O(n)     - 避免复制\n")
	fmt.Printf("方法三（单调栈+映射）：   时间 O(n),      空间 O(n)     - 代码简洁\n")
	fmt.Printf("方法四（暴力枚举）：     时间 O(n²),     空间 O(1)     - 简单易懂\n")
}

// 详细分析示例
func analyzeExample(nums []int) {
	fmt.Printf("分析循环数组: ")
	printArray(nums, "")
	fmt.Printf("索引位置: [")
	for i := range nums {
		if i > 0 {
			fmt.Printf(", ")
		}
		fmt.Printf("%d", i)
	}
	fmt.Printf("]\n")

	// 使用单调栈分析过程
	stack := []int{}
	result := make([]int, len(nums))
	for i := range result {
		result[i] = -1
	}
	fmt.Printf("\n单调栈处理过程（两遍遍历）:\n")

	for i := 0; i < 2*len(nums); i++ {
		actualIndex := i % len(nums)
		fmt.Printf("步骤 %d: 实际索引=%d, 当前元素=%d, 栈=%v ", i, actualIndex, nums[actualIndex], stack)

		// 当栈非空且当前元素大于栈顶元素时
		for len(stack) > 0 && nums[actualIndex] > nums[stack[len(stack)-1]] {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			result[top] = nums[actualIndex]
			fmt.Printf("弹出索引%d(元素%d), result[%d]=%d ", top, nums[top], top, result[top])
		}

		// 只在第一遍遍历时入栈
		if i < len(nums) {
			stack = append(stack, i)
			fmt.Printf("入栈索引%d, 新栈=%v", i, stack)
		}
		fmt.Printf("\n")
	}

	fmt.Printf("最终结果: ")
	printArray(result, "")
}

func main() {
	runTests()
}
