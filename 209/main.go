package main

import "fmt"

// 方法1：滑动窗口（推荐）
// 时间复杂度：O(n)，空间复杂度：O(1)
func minSubArrayLen(target int, nums []int) int {
	n := len(nums)
	left, right := 0, 0
	sum := 0
	minLen := n + 1 // 初始化为不可能的长度

	for right < n {
		// 扩大窗口
		sum += nums[right]
		right++

		// 当窗口和满足条件时，尝试缩小窗口
		for sum >= target {
			// 更新最小长度
			if right-left < minLen {
				minLen = right - left
			}
			// 缩小窗口
			sum -= nums[left]
			left++
		}
	}

	// 如果没有找到符合条件的子数组
	if minLen == n+1 {
		return 0
	}
	return minLen
}

// 方法2：前缀和 + 二分查找
// 时间复杂度：O(n log n)，空间复杂度：O(n)
func minSubArrayLenBinarySearch(target int, nums []int) int {
	n := len(nums)

	// 计算前缀和
	prefixSum := make([]int, n+1)
	for i := 0; i < n; i++ {
		prefixSum[i+1] = prefixSum[i] + nums[i]
	}

	minLen := n + 1

	// 对每个位置二分查找
	for i := 0; i < n; i++ {
		// 二分查找满足条件的最远位置
		left, right := i+1, n
		for left <= right {
			mid := left + (right-left)/2
			if prefixSum[mid]-prefixSum[i] >= target {
				// 找到满足条件的位置，尝试更小的
				if mid-i < minLen {
					minLen = mid - i
				}
				right = mid - 1
			} else {
				left = mid + 1
			}
		}
	}

	if minLen == n+1 {
		return 0
	}
	return minLen
}

// 方法3：暴力枚举
// 时间复杂度：O(n²)，空间复杂度：O(1)
func minSubArrayLenBruteForce(target int, nums []int) int {
	n := len(nums)
	minLen := n + 1

	for i := 0; i < n; i++ {
		sum := 0
		for j := i; j < n; j++ {
			sum += nums[j]
			if sum >= target {
				if j-i+1 < minLen {
					minLen = j - i + 1
				}
				break // 找到满足条件的子数组，可以提前结束
			}
		}
	}

	if minLen == n+1 {
		return 0
	}
	return minLen
}

// 方法4：滑动窗口优化版本
// 时间复杂度：O(n)，空间复杂度：O(1)
func minSubArrayLenOptimized(target int, nums []int) int {
	n := len(nums)
	left, sum := 0, 0
	minLen := n + 1

	for right := 0; right < n; right++ {
		sum += nums[right]

		// 当窗口和满足条件时，尝试缩小窗口
		for sum >= target {
			minLen = min(minLen, right-left+1)
			sum -= nums[left]
			left++
		}
	}

	if minLen == n+1 {
		return 0
	}
	return minLen
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// 工具函数：打印数组
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
	fmt.Println("=== 209. 长度最小的子数组 测试 ===")

	// 测试用例1
	target1 := 7
	nums1 := []int{2, 3, 1, 2, 4, 3}
	expected1 := 2

	fmt.Printf("\n测试用例1:\n")
	fmt.Printf("target: %d\n", target1)
	printArray(nums1, "输入数组")
	fmt.Printf("期望结果: %d\n", expected1)

	result1_1 := minSubArrayLen(target1, nums1)
	result1_2 := minSubArrayLenBinarySearch(target1, nums1)
	result1_3 := minSubArrayLenBruteForce(target1, nums1)
	result1_4 := minSubArrayLenOptimized(target1, nums1)

	fmt.Printf("方法一（滑动窗口）: %d\n", result1_1)
	fmt.Printf("方法二（前缀和+二分）: %d\n", result1_2)
	fmt.Printf("方法三（暴力枚举）: %d\n", result1_3)
	fmt.Printf("方法四（滑动窗口优化）: %d\n", result1_4)

	// 测试用例2
	target2 := 4
	nums2 := []int{1, 4, 4}
	expected2 := 1

	fmt.Printf("\n测试用例2:\n")
	fmt.Printf("target: %d\n", target2)
	printArray(nums2, "输入数组")
	fmt.Printf("期望结果: %d\n", expected2)

	result2_1 := minSubArrayLen(target2, nums2)
	result2_2 := minSubArrayLenBinarySearch(target2, nums2)
	result2_3 := minSubArrayLenBruteForce(target2, nums2)
	result2_4 := minSubArrayLenOptimized(target2, nums2)

	fmt.Printf("方法一（滑动窗口）: %d\n", result2_1)
	fmt.Printf("方法二（前缀和+二分）: %d\n", result2_2)
	fmt.Printf("方法三（暴力枚举）: %d\n", result2_3)
	fmt.Printf("方法四（滑动窗口优化）: %d\n", result2_4)

	// 测试用例3：无解情况
	target3 := 11
	nums3 := []int{1, 1, 1, 1, 1, 1, 1, 1}
	expected3 := 0

	fmt.Printf("\n测试用例3（无解情况）:\n")
	fmt.Printf("target: %d\n", target3)
	printArray(nums3, "输入数组")
	fmt.Printf("期望结果: %d\n", expected3)

	result3 := minSubArrayLen(target3, nums3)
	fmt.Printf("结果: %d\n", result3)

	// 测试用例4：边界情况
	target4 := 1
	nums4 := []int{1}
	expected4 := 1

	fmt.Printf("\n测试用例4（单个元素）:\n")
	fmt.Printf("target: %d\n", target4)
	printArray(nums4, "输入数组")
	fmt.Printf("期望结果: %d\n", expected4)

	result4 := minSubArrayLen(target4, nums4)
	fmt.Printf("结果: %d\n", result4)

	// 测试用例5：需要多个元素
	target5 := 10
	nums5 := []int{1, 2, 3, 4, 5}
	expected5 := 3

	fmt.Printf("\n测试用例5（需要多个元素）:\n")
	fmt.Printf("target: %d\n", target5)
	printArray(nums5, "输入数组")
	fmt.Printf("期望结果: %d\n", expected5)

	result5 := minSubArrayLen(target5, nums5)
	fmt.Printf("结果: %d\n", result5)

	// 测试用例6：极值测试
	target6 := 100
	nums6 := []int{1, 2, 3, 4, 5}
	expected6 := 0

	fmt.Printf("\n测试用例6（极值测试）:\n")
	fmt.Printf("target: %d\n", target6)
	printArray(nums6, "输入数组")
	fmt.Printf("期望结果: %d\n", expected6)

	result6 := minSubArrayLen(target6, nums6)
	fmt.Printf("结果: %d\n", result6)

	// 测试用例7：单个元素满足条件
	target7 := 5
	nums7 := []int{5}
	expected7 := 1

	fmt.Printf("\n测试用例7（单个元素满足条件）:\n")
	fmt.Printf("target: %d\n", target7)
	printArray(nums7, "输入数组")
	fmt.Printf("期望结果: %d\n", expected7)

	result7 := minSubArrayLen(target7, nums7)
	fmt.Printf("结果: %d\n", result7)

	// 测试用例8：复杂情况
	target8 := 15
	nums8 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	expected8 := 2

	fmt.Printf("\n测试用例8（复杂情况）:\n")
	fmt.Printf("target: %d\n", target8)
	printArray(nums8, "输入数组")
	fmt.Printf("期望结果: %d\n", expected8)

	result8 := minSubArrayLen(target8, nums8)
	fmt.Printf("结果: %d\n", result8)

	// 测试用例9：重复元素
	target9 := 6
	nums9 := []int{2, 2, 2, 2, 2}
	expected9 := 3

	fmt.Printf("\n测试用例9（重复元素）:\n")
	fmt.Printf("target: %d\n", target9)
	printArray(nums9, "输入数组")
	fmt.Printf("期望结果: %d\n", expected9)

	result9 := minSubArrayLen(target9, nums9)
	fmt.Printf("结果: %d\n", result9)

	// 详细分析示例
	fmt.Printf("\n=== 详细分析示例 ===\n")
	analyzeExample(target1, nums1)

	fmt.Printf("\n=== 算法复杂度对比 ===\n")
	fmt.Printf("方法一（滑动窗口）：       时间 O(n),      空间 O(1)     - 推荐\n")
	fmt.Printf("方法二（前缀和+二分）：     时间 O(n log n), 空间 O(n)     - 满足进阶要求\n")
	fmt.Printf("方法三（暴力枚举）：       时间 O(n²),     空间 O(1)     - 简单易懂\n")
	fmt.Printf("方法四（滑动窗口优化）：   时间 O(n),      空间 O(1)     - 代码简洁\n")
}

// 详细分析示例
func analyzeExample(target int, nums []int) {
	fmt.Printf("分析滑动窗口过程: target=%d, nums=", target)
	printArray(nums, "")
	fmt.Printf("索引位置: [")
	for i := range nums {
		if i > 0 {
			fmt.Printf(", ")
		}
		fmt.Printf("%d", i)
	}
	fmt.Printf("]\n")

	// 使用滑动窗口分析过程
	left, right := 0, 0
	sum := 0
	minLen := len(nums) + 1
	fmt.Printf("\n滑动窗口处理过程:\n")

	for right < len(nums) {
		// 扩大窗口
		sum += nums[right]
		fmt.Printf("步骤 %d: right=%d, 加入元素%d, sum=%d, 窗口[%d,%d] ", right+1, right, nums[right], sum, left, right)

		// 当窗口和满足条件时，尝试缩小窗口
		for sum >= target {
			if right-left+1 < minLen {
				minLen = right - left + 1
				fmt.Printf("更新minLen=%d ", minLen)
			}
			// 缩小窗口
			sum -= nums[left]
			fmt.Printf("缩小窗口, 移除元素%d, sum=%d, 窗口[%d,%d]", nums[left], sum, left+1, right)
			left++
		}
		fmt.Printf("\n")
		right++
	}

	fmt.Printf("最终结果: %d\n", minLen)
	if minLen == len(nums)+1 {
		fmt.Printf("没有找到满足条件的子数组\n")
	} else {
		fmt.Printf("最小长度子数组长度为: %d\n", minLen)
	}
}

func main() {
	runTests()
}
