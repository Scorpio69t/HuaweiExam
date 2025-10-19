package main

import (
	"fmt"
)

// =========================== 方法一：双指针算法（最优解法） ===========================

func removeDuplicates(nums []int) int {
	if len(nums) <= 2 {
		return len(nums)
	}

	slow := 0
	count := 1

	for fast := 1; fast < len(nums); fast++ {
		if nums[fast] == nums[fast-1] {
			count++
		} else {
			count = 1
		}

		if count <= 2 {
			slow++
			nums[slow] = nums[fast]
		}
	}

	return slow + 1
}

// =========================== 方法二：计数法 ===========================

func removeDuplicates2(nums []int) int {
	if len(nums) <= 2 {
		return len(nums)
	}

	slow := 0
	count := 1

	for fast := 1; fast < len(nums); fast++ {
		if nums[fast] == nums[fast-1] {
			count++
		} else {
			count = 1
		}

		if count <= 2 {
			slow++
			nums[slow] = nums[fast]
		}
	}

	return slow + 1
}

// =========================== 方法三：快慢指针 ===========================

func removeDuplicates3(nums []int) int {
	if len(nums) <= 2 {
		return len(nums)
	}

	slow := 2

	for fast := 2; fast < len(nums); fast++ {
		if nums[fast] != nums[slow-2] {
			nums[slow] = nums[fast]
			slow++
		}
	}

	return slow
}

// =========================== 方法四：暴力法 ===========================

func removeDuplicates4(nums []int) int {
	if len(nums) <= 2 {
		return len(nums)
	}

	length := len(nums)
	for i := 2; i < length; i++ {
		if nums[i] == nums[i-1] && nums[i] == nums[i-2] {
			// 删除当前元素
			for j := i; j < length-1; j++ {
				nums[j] = nums[j+1]
			}
			length--
			i-- // 重新检查当前位置
		}
	}

	return length
}

// =========================== 测试代码 ===========================

func main() {
	fmt.Println("=== LeetCode 80: 删除有序数组中的重复项 II ===\n")

	testCases := []struct {
		nums   []int
		expect int
		result []int
	}{
		{
			[]int{1, 1, 1, 2, 2, 3},
			5,
			[]int{1, 1, 2, 2, 3},
		},
		{
			[]int{0, 0, 1, 1, 1, 1, 2, 3, 3},
			7,
			[]int{0, 0, 1, 1, 2, 3, 3},
		},
		{
			[]int{1, 2, 3, 4, 5},
			5,
			[]int{1, 2, 3, 4, 5},
		},
		{
			[]int{1, 1, 1, 1, 1},
			2,
			[]int{1, 1},
		},
		{
			[]int{},
			0,
			[]int{},
		},
		{
			[]int{1},
			1,
			[]int{1},
		},
		{
			[]int{1, 1},
			2,
			[]int{1, 1},
		},
	}

	fmt.Println("方法一：双指针算法（最优解法）")
	runTests(testCases, removeDuplicates)

	fmt.Println("\n方法二：计数法")
	runTests(testCases, removeDuplicates2)

	fmt.Println("\n方法三：快慢指针")
	runTests(testCases, removeDuplicates3)

	fmt.Println("\n方法四：暴力法")
	runTests(testCases, removeDuplicates4)
}

func runTests(testCases []struct {
	nums   []int
	expect int
	result []int
}, fn func([]int) int) {
	passCount := 0
	for i, tc := range testCases {
		// 创建nums的副本，避免原地修改影响其他测试
		numsCopy := make([]int, len(tc.nums))
		copy(numsCopy, tc.nums)

		result := fn(numsCopy)
		status := "✅"
		if result != tc.expect {
			status = "❌"
		} else {
			// 检查结果数组
			if !compareArrays(numsCopy[:result], tc.result) {
				status = "❌"
			} else {
				passCount++
			}
		}
		fmt.Printf("  测试%d: %s\n", i+1, status)
		if status == "❌" {
			fmt.Printf("    输入: %v\n", tc.nums)
			fmt.Printf("    输出: %d, %v\n", result, numsCopy[:result])
			fmt.Printf("    期望: %d, %v\n", tc.expect, tc.result)
		}
	}
	fmt.Printf("  通过: %d/%d\n", passCount, len(testCases))
}

func compareArrays(a, b []int) bool {
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
