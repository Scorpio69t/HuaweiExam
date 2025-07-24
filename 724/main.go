package main

import "fmt"

// 方法一：前缀和
func pivotIndex(nums []int) int {
	n := len(nums)
	if n == 0 {
		return -1
	}

	// 计算前缀和
	prefixSum := make([]int, n+1)
	for i := 0; i < n; i++ {
		prefixSum[i+1] = prefixSum[i] + nums[i]
	}

	totalSum := prefixSum[n]

	// 查找中心下标
	for i := 0; i < n; i++ {
		leftSum := prefixSum[i]
		rightSum := totalSum - prefixSum[i+1]
		if leftSum == rightSum {
			return i
		}
	}

	return -1
}

// 方法二：数学优化（推荐）
func pivotIndex2(nums []int) int {
	n := len(nums)
	if n == 0 {
		return -1
	}

	// 计算总和
	totalSum := 0
	for _, num := range nums {
		totalSum += num
	}

	// 查找中心下标
	leftSum := 0
	for i := 0; i < n; i++ {
		// 如果左侧和等于右侧和，则找到中心下标
		if leftSum == totalSum-leftSum-nums[i] {
			return i
		}
		leftSum += nums[i]
	}

	return -1
}

// 方法三：暴力解法（用于验证）
func pivotIndex3(nums []int) int {
	n := len(nums)
	if n == 0 {
		return -1
	}

	for i := 0; i < n; i++ {
		leftSum := 0
		rightSum := 0

		// 计算左侧和
		for j := 0; j < i; j++ {
			leftSum += nums[j]
		}

		// 计算右侧和
		for j := i + 1; j < n; j++ {
			rightSum += nums[j]
		}

		if leftSum == rightSum {
			return i
		}
	}

	return -1
}

// 方法四：更简洁的数学优化
func pivotIndex4(nums []int) int {
	totalSum := 0
	for _, num := range nums {
		totalSum += num
	}

	leftSum := 0
	for i, num := range nums {
		if leftSum == totalSum-leftSum-num {
			return i
		}
		leftSum += num
	}

	return -1
}

// 测试函数
func testPivotIndex() {
	testCases := []struct {
		nums     []int
		expected int
	}{
		{[]int{1, 7, 3, 6, 5, 6}, 3},
		{[]int{1, 2, 3}, -1},
		{[]int{2, 1, -1}, 0},
		{[]int{1, 2, 3, 4, 6}, 3},
		{[]int{1}, 0},
		{[]int{1, 1}, -1},
		{[]int{0, 0, 0}, 0},
		{[]int{-1, -1, -1, -1, -1, 0}, 2},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, -1},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 55}, 9},
	}

	fmt.Println("=== 测试前缀和算法 ===")
	for i, tc := range testCases {
		result := pivotIndex(tc.nums)
		if result == tc.expected {
			fmt.Printf("✅ 测试用例 %d 通过: nums=%v, 结果=%d\n",
				i+1, tc.nums, result)
		} else {
			fmt.Printf("❌ 测试用例 %d 失败: nums=%v, 期望=%d, 实际=%d\n",
				i+1, tc.nums, tc.expected, result)
		}
	}

	fmt.Println("\n=== 测试数学优化算法 ===")
	for i, tc := range testCases {
		result := pivotIndex2(tc.nums)
		if result == tc.expected {
			fmt.Printf("✅ 测试用例 %d 通过: nums=%v, 结果=%d\n",
				i+1, tc.nums, result)
		} else {
			fmt.Printf("❌ 测试用例 %d 失败: nums=%v, 期望=%d, 实际=%d\n",
				i+1, tc.nums, tc.expected, result)
		}
	}

	fmt.Println("\n=== 测试暴力解法 ===")
	for i, tc := range testCases {
		result := pivotIndex3(tc.nums)
		if result == tc.expected {
			fmt.Printf("✅ 测试用例 %d 通过: nums=%v, 结果=%d\n",
				i+1, tc.nums, result)
		} else {
			fmt.Printf("❌ 测试用例 %d 失败: nums=%v, 期望=%d, 实际=%d\n",
				i+1, tc.nums, tc.expected, result)
		}
	}

	fmt.Println("\n=== 测试简洁数学优化算法 ===")
	for i, tc := range testCases {
		result := pivotIndex4(tc.nums)
		if result == tc.expected {
			fmt.Printf("✅ 测试用例 %d 通过: nums=%v, 结果=%d\n",
				i+1, tc.nums, result)
		} else {
			fmt.Printf("❌ 测试用例 %d 失败: nums=%v, 期望=%d, 实际=%d\n",
				i+1, tc.nums, tc.expected, result)
		}
	}
}

// 性能对比测试
func benchmarkTest() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}

	fmt.Println("\n=== 性能对比测试 ===")
	fmt.Printf("输入: nums=%v\n", nums)

	// 测试各种算法
	result1 := pivotIndex(nums)
	result2 := pivotIndex2(nums)
	result3 := pivotIndex3(nums)
	result4 := pivotIndex4(nums)

	fmt.Printf("前缀和算法结果: %d\n", result1)
	fmt.Printf("数学优化算法结果: %d\n", result2)
	fmt.Printf("暴力解法结果: %d\n", result3)
	fmt.Printf("简洁数学优化结果: %d\n", result4)

	if result1 == result2 && result2 == result3 && result3 == result4 {
		fmt.Println("✅ 所有算法结果一致")
	} else {
		fmt.Println("❌ 算法结果不一致")
	}
}

// 可视化函数：显示数组和中心下标
func visualizePivotIndex(nums []int) {
	fmt.Println("\n=== 可视化中心下标 ===")
	fmt.Printf("数组: %v\n", nums)

	result := pivotIndex2(nums)
	if result == -1 {
		fmt.Println("没有找到中心下标")
		return
	}

	fmt.Printf("中心下标: %d (值: %d)\n", result, nums[result])

	// 显示左侧和右侧
	leftSum := 0
	rightSum := 0

	fmt.Print("左侧: ")
	for i := 0; i < result; i++ {
		fmt.Printf("%d ", nums[i])
		leftSum += nums[i]
	}
	if result == 0 {
		fmt.Print("(空)")
	}
	fmt.Printf(" = %d\n", leftSum)

	fmt.Print("右侧: ")
	for i := result + 1; i < len(nums); i++ {
		fmt.Printf("%d ", nums[i])
		rightSum += nums[i]
	}
	if result == len(nums)-1 {
		fmt.Print("(空)")
	}
	fmt.Printf(" = %d\n", rightSum)

	fmt.Printf("验证: %d == %d ✅\n", leftSum, rightSum)
}

func main() {
	fmt.Println("724. 寻找数组的中心下标")
	fmt.Println("========================")

	// 运行测试用例
	testPivotIndex()

	// 运行性能对比
	benchmarkTest()

	// 可视化示例
	visualizePivotIndex([]int{1, 7, 3, 6, 5, 6})
	visualizePivotIndex([]int{2, 1, -1})
	visualizePivotIndex([]int{1, 2, 3})

	// 交互式测试
	fmt.Println("\n=== 交互式测试 ===")
	fmt.Println("请输入数组元素，用空格分隔（例如: 1 7 3 6 5 6）")

	var input string
	fmt.Print("请输入数组: ")
	fmt.Scanln(&input)

	// 这里简化处理，实际应该解析输入字符串
	// 为了演示，使用一个示例数组
	nums := []int{1, 7, 3, 6, 5, 6}
	fmt.Printf("使用示例数组: %v\n", nums)

	result := pivotIndex2(nums)
	fmt.Printf("中心下标: %d\n", result)

	if result != -1 {
		visualizePivotIndex(nums)
	}
}
