package main

import "fmt"

// 方法一：二分查找解法（推荐）
// 时间复杂度：O(log n)，空间复杂度：O(1)
func search(nums []int, target int) int {
	left, right := 0, len(nums)-1

	for left <= right {
		mid := left + (right-left)/2

		if nums[mid] == target {
			return mid
		}

		// 判断左半部分是否有序
		if nums[left] <= nums[mid] {
			// 左半部分有序
			if nums[left] <= target && target < nums[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else {
			// 右半部分有序
			if nums[mid] < target && target <= nums[right] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}

	return -1
}

// 方法二：先找旋转点，再二分查找
// 时间复杂度：O(log n)，空间复杂度：O(1)
func searchFindPivot(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}

	// 找到旋转点（最小值的位置）
	pivot := findPivot(nums)
	
	// 根据target与边界值的比较，确定搜索区间
	if pivot == 0 {
		// 数组没有旋转，直接二分查找
		return binarySearch(nums, target, 0, len(nums)-1)
	}

	// 在左半部分搜索
	if nums[0] <= target && target <= nums[pivot-1] {
		return binarySearch(nums, target, 0, pivot-1)
	}
	// 在右半部分搜索
	return binarySearch(nums, target, pivot, len(nums)-1)
}

// 找到旋转点（最小值的位置）
func findPivot(nums []int) int {
	left, right := 0, len(nums)-1

	for left < right {
		mid := left + (right-left)/2

		if nums[mid] > nums[right] {
			left = mid + 1
		} else {
			right = mid
		}
	}

	return left
}

// 标准二分查找
func binarySearch(nums []int, target, left, right int) int {
	for left <= right {
		mid := left + (right-left)/2

		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return -1
}

// 方法三：线性搜索（用于验证和调试）
// 时间复杂度：O(n)，空间复杂度：O(1)
func searchLinear(nums []int, target int) int {
	for i, num := range nums {
		if num == target {
			return i
		}
	}
	return -1
}

// 方法四：优化的二分查找解法
// 时间复杂度：O(log n)，空间复杂度：O(1)
func searchOptimized(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}

	left, right := 0, len(nums)-1

	for left <= right {
		mid := left + (right-left)/2

		if nums[mid] == target {
			return mid
		}

		// 处理重复元素的情况（虽然题目保证元素互不相同）
		if nums[left] == nums[mid] && nums[mid] == nums[right] {
			left++
			right--
			continue
		}

		// 判断左半部分是否有序
		if nums[left] <= nums[mid] {
			// 左半部分有序
			if nums[left] <= target && target < nums[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else {
			// 右半部分有序
			if nums[mid] < target && target <= nums[right] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}

	return -1
}

// 辅助函数：打印数组和搜索过程
func printSearchProcess(nums []int, target int) {
	fmt.Printf("数组: %v, 目标值: %d\n", nums, target)
	
	// 找到旋转点
	pivot := findPivot(nums)
	fmt.Printf("旋转点位置: %d (值: %d)\n", pivot, nums[pivot])
	
	// 显示有序区间
	if pivot == 0 {
		fmt.Println("数组没有旋转，完全有序")
	} else {
		fmt.Printf("左半部分有序区间: [%d, %d]\n", 0, pivot-1)
		fmt.Printf("右半部分有序区间: [%d, %d]\n", pivot, len(nums)-1)
	}
	
	result := search(nums, target)
	fmt.Printf("搜索结果: %d\n", result)
	fmt.Println()
}

func main() {
	fmt.Println("=== 33. 搜索旋转排序数组 ===")

	// 测试用例1
	nums1 := []int{4, 5, 6, 7, 0, 1, 2}
	target1 := 0
	fmt.Printf("测试用例1: nums=%v, target=%d\n", nums1, target1)
	fmt.Printf("二分查找解法结果: %d\n", search(nums1, target1))
	fmt.Printf("找旋转点+二分解法结果: %d\n", searchFindPivot(nums1, target1))
	fmt.Printf("线性搜索解法结果: %d\n", searchLinear(nums1, target1))
	fmt.Printf("优化二分查找解法结果: %d\n", searchOptimized(nums1, target1))
	fmt.Println()

	// 测试用例2
	nums2 := []int{4, 5, 6, 7, 0, 1, 2}
	target2 := 3
	fmt.Printf("测试用例2: nums=%v, target=%d\n", nums2, target2)
	fmt.Printf("二分查找解法结果: %d\n", search(nums2, target2))
	fmt.Printf("找旋转点+二分解法结果: %d\n", searchFindPivot(nums2, target2))
	fmt.Printf("线性搜索解法结果: %d\n", searchLinear(nums2, target2))
	fmt.Printf("优化二分查找解法结果: %d\n", searchOptimized(nums2, target2))
	fmt.Println()

	// 边界测试用例
	testCases := []struct {
		nums   []int
		target int
		desc   string
	}{
		{[]int{1}, 0, "单个元素，目标不存在"},
		{[]int{1}, 1, "单个元素，目标存在"},
		{[]int{3, 1}, 1, "两个元素"},
		{[]int{3, 1}, 3, "两个元素，目标在开头"},
		{[]int{1, 3}, 3, "两个元素，未旋转"},
		{[]int{1, 3}, 1, "两个元素，未旋转，目标在开头"},
		{[]int{1, 2, 3, 4, 5}, 3, "完全有序数组"},
		{[]int{5, 1, 2, 3, 4}, 1, "旋转一个位置"},
		{[]int{2, 3, 4, 5, 1}, 1, "旋转到末尾"},
	}

	for _, tc := range testCases {
		result := search(tc.nums, tc.target)
		fmt.Printf("%s: nums=%v, target=%d, 结果=%d\n", tc.desc, tc.nums, tc.target, result)
	}

	// 详细搜索过程演示
	fmt.Println("\n=== 详细搜索过程演示 ===")
	demoCases := []struct {
		nums   []int
		target int
	}{
		{[]int{4, 5, 6, 7, 0, 1, 2}, 0},
		{[]int{4, 5, 6, 7, 0, 1, 2}, 6},
		{[]int{3, 1}, 1},
		{[]int{1, 2, 3, 4, 5}, 3},
	}

	for _, demo := range demoCases {
		printSearchProcess(demo.nums, demo.target)
	}

	// 算法正确性验证
	fmt.Println("=== 算法正确性验证 ===")
	verifyNums := []int{4, 5, 6, 7, 0, 1, 2}
	verifyTarget := 5
	
	fmt.Printf("验证数组: %v, 目标值: %d\n", verifyNums, verifyTarget)
	fmt.Printf("二分查找解法: %d\n", search(verifyNums, verifyTarget))
	fmt.Printf("找旋转点+二分解法: %d\n", searchFindPivot(verifyNums, verifyTarget))
	fmt.Printf("线性搜索解法: %d\n", searchLinear(verifyNums, verifyTarget))
	fmt.Printf("优化二分查找解法: %d\n", searchOptimized(verifyNums, verifyTarget))

	// 验证所有解法结果一致
	if search(verifyNums, verifyTarget) == searchFindPivot(verifyNums, verifyTarget) &&
		search(verifyNums, verifyTarget) == searchLinear(verifyNums, verifyTarget) &&
		search(verifyNums, verifyTarget) == searchOptimized(verifyNums, verifyTarget) {
		fmt.Println("✅ 所有解法结果一致，算法正确！")
	} else {
		fmt.Println("❌ 解法结果不一致，需要检查！")
	}

	// 性能测试
	fmt.Println("\n=== 性能测试 ===")
	largeNums := make([]int, 10000)
	for i := range largeNums {
		largeNums[i] = i
	}
	// 旋转数组
	rotated := append(largeNums[5000:], largeNums[:5000]...)
	largeTarget := 7500

	fmt.Printf("大数组测试: 长度=%d, 目标值=%d\n", len(rotated), largeTarget)
	result := search(rotated, largeTarget)
	fmt.Printf("二分查找解法结果: %d\n", result)
}
