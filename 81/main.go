package main

import (
	"fmt"
)

// =========================== 方法一：二分查找（最优解法） ===========================

func search(nums []int, target int) bool {
	if len(nums) == 0 {
		return false
	}

	left, right := 0, len(nums)-1

	for left <= right {
		mid := left + (right-left)/2

		if nums[mid] == target {
			return true
		}

		// 处理重复元素
		if nums[left] == nums[mid] && nums[mid] == nums[right] {
			left++
			right--
		} else if nums[left] <= nums[mid] {
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

	return false
}

// =========================== 方法二：线性搜索 ===========================

func search2(nums []int, target int) bool {
	for _, num := range nums {
		if num == target {
			return true
		}
	}
	return false
}

// =========================== 方法三：哈希表 ===========================

func search3(nums []int, target int) bool {
	numMap := make(map[int]bool)
	for _, num := range nums {
		numMap[num] = true
	}
	return numMap[target]
}

// =========================== 方法四：优化版二分查找 ===========================

func search4(nums []int, target int) bool {
	if len(nums) == 0 {
		return false
	}

	left, right := 0, len(nums)-1

	for left <= right {
		mid := left + (right-left)/2

		if nums[mid] == target {
			return true
		}

		// 处理重复元素
		if nums[left] == nums[mid] && nums[mid] == nums[right] {
			left++
			right--
			continue
		}

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

	return false
}

// =========================== 测试代码 ===========================

func main() {
	fmt.Println("=== LeetCode 81: 搜索旋转排序数组 II ===\n")

	testCases := []struct {
		nums   []int
		target int
		expect bool
	}{
		{
			[]int{2, 5, 6, 0, 0, 1, 2},
			0,
			true,
		},
		{
			[]int{2, 5, 6, 0, 0, 1, 2},
			3,
			false,
		},
		{
			[]int{1},
			1,
			true,
		},
		{
			[]int{},
			0,
			false,
		},
		{
			[]int{1, 1, 1, 1, 1, 1, 1},
			1,
			true,
		},
		{
			[]int{1, 1, 1, 1, 1, 1, 1},
			2,
			false,
		},
		{
			[]int{1, 3, 1, 1, 1},
			3,
			true,
		},
		{
			[]int{1, 3, 1, 1, 1},
			2,
			false,
		},
	}

	fmt.Println("方法一：二分查找（最优解法）")
	runTests(testCases, search)

	fmt.Println("\n方法二：线性搜索")
	runTests(testCases, search2)

	fmt.Println("\n方法三：哈希表")
	runTests(testCases, search3)

	fmt.Println("\n方法四：优化版二分查找")
	runTests(testCases, search4)
}

func runTests(testCases []struct {
	nums   []int
	target int
	expect bool
}, fn func([]int, int) bool) {
	passCount := 0
	for i, tc := range testCases {
		result := fn(tc.nums, tc.target)
		status := "✅"
		if result != tc.expect {
			status = "❌"
		} else {
			passCount++
		}
		fmt.Printf("  测试%d: %s\n", i+1, status)
		if status == "❌" {
			fmt.Printf("    输入: nums=%v, target=%d\n", tc.nums, tc.target)
			fmt.Printf("    输出: %t\n", result)
			fmt.Printf("    期望: %t\n", tc.expect)
		}
	}
	fmt.Printf("  通过: %d/%d\n", passCount, len(testCases))
}
