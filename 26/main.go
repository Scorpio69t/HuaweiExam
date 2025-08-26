package main

import (
	"fmt"
)

// removeDuplicates 双指针法：删除有序数组中的重复项
// 时间复杂度: O(n)；空间复杂度: O(1)
func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	// 使用双指针，slow指向下一个唯一元素应该放置的位置
	slow := 0
	for fast := 1; fast < len(nums); fast++ {
		// 当快指针遇到不同的元素时，将其放到慢指针位置
		if nums[fast] != nums[slow] {
			slow++
			nums[slow] = nums[fast]
		}
	}

	return slow + 1
}

// removeDuplicatesOptimized 优化双指针法：减少不必要的赋值
func removeDuplicatesOptimized(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	slow := 0
	for fast := 1; fast < len(nums); fast++ {
		if nums[fast] != nums[slow] {
			slow++
			// 只有当slow != fast时才需要赋值，避免自我赋值
			if slow != fast {
				nums[slow] = nums[fast]
			}
		}
	}

	return slow + 1
}

// removeDuplicatesCount 计数法：统计重复元素数量
func removeDuplicatesCount(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	duplicates := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] {
			duplicates++
		} else {
			// 将非重复元素移动到正确位置
			nums[i-duplicates] = nums[i]
		}
	}

	return len(nums) - duplicates
}

// removeDuplicatesRecursive 递归法：使用递归思想
func removeDuplicatesRecursive(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}

	// 递归处理子数组
	uniqueCount := removeDuplicatesRecursive(nums[:len(nums)-1])

	// 检查最后一个元素是否与前面的唯一元素重复
	if uniqueCount == 0 || nums[len(nums)-1] != nums[uniqueCount-1] {
		nums[uniqueCount] = nums[len(nums)-1]
		return uniqueCount + 1
	}

	return uniqueCount
}

// removeDuplicatesInPlace 原地操作法：直接在原数组上操作
func removeDuplicatesInPlace(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	writeIndex := 0
	for readIndex := 0; readIndex < len(nums); readIndex++ {
		// 如果是第一个元素，或者与前一个元素不同，则写入
		if readIndex == 0 || nums[readIndex] != nums[readIndex-1] {
			nums[writeIndex] = nums[readIndex]
			writeIndex++
		}
	}

	return writeIndex
}

// 辅助函数：打印数组前k个元素
func printArray(nums []int, k int) {
	fmt.Printf("[")
	for i := 0; i < k; i++ {
		if i > 0 {
			fmt.Printf(",")
		}
		fmt.Printf("%d", nums[i])
	}
	fmt.Printf("]\n")
}

func main() {
	// 测试用例1: [1,1,2] -> [1,2], 长度2
	nums1 := []int{1, 1, 2}
	fmt.Println("测试用例1:")
	fmt.Printf("原数组: %v\n", nums1)
	k1 := removeDuplicates(nums1)
	fmt.Printf("结果: 长度=%d, 数组=", k1)
	printArray(nums1, k1)

	// 测试用例2: [0,0,1,1,1,2,2,3,3,4] -> [0,1,2,3,4], 长度5
	nums2 := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	fmt.Println("\n测试用例2:")
	fmt.Printf("原数组: %v\n", nums2)
	k2 := removeDuplicates(nums2)
	fmt.Printf("结果: 长度=%d, 数组=", k2)
	printArray(nums2, k2)

	// 测试边界情况：空数组
	nums3 := []int{}
	fmt.Println("\n边界测试 - 空数组:")
	fmt.Printf("原数组: %v\n", nums3)
	k3 := removeDuplicates(nums3)
	fmt.Printf("结果: 长度=%d\n", k3)

	// 测试边界情况：单元素数组
	nums4 := []int{5}
	fmt.Println("\n边界测试 - 单元素数组:")
	fmt.Printf("原数组: %v\n", nums4)
	k4 := removeDuplicates(nums4)
	fmt.Printf("结果: 长度=%d, 数组=", k4)
	printArray(nums4, k4)

	// 测试边界情况：全相同元素
	nums5 := []int{3, 3, 3, 3, 3}
	fmt.Println("\n边界测试 - 全相同元素:")
	fmt.Printf("原数组: %v\n", nums5)
	k5 := removeDuplicates(nums5)
	fmt.Printf("结果: 长度=%d, 数组=", k5)
	printArray(nums5, k5)

	// 测试边界情况：无重复元素
	nums6 := []int{1, 2, 3, 4, 5}
	fmt.Println("\n边界测试 - 无重复元素:")
	fmt.Printf("原数组: %v\n", nums6)
	k6 := removeDuplicates(nums6)
	fmt.Printf("结果: 长度=%d, 数组=", k6)
	printArray(nums6, k6)

	// 测试不同算法的结果一致性
	nums7 := []int{1, 1, 2, 2, 3, 3, 4, 4, 5}
	fmt.Println("\n算法一致性测试:")
	fmt.Printf("原数组: %v\n", nums7)

	// 复制数组用于不同算法测试
	nums7Copy1 := make([]int, len(nums7))
	nums7Copy2 := make([]int, len(nums7))
	nums7Copy3 := make([]int, len(nums7))
	copy(nums7Copy1, nums7)
	copy(nums7Copy2, nums7)
	copy(nums7Copy3, nums7)

	k7_1 := removeDuplicates(nums7)
	k7_2 := removeDuplicatesOptimized(nums7Copy1)
	k7_3 := removeDuplicatesCount(nums7Copy2)
	k7_4 := removeDuplicatesRecursive(nums7Copy3)

	fmt.Printf("双指针法: 长度=%d\n", k7_1)
	fmt.Printf("优化双指针: 长度=%d\n", k7_2)
	fmt.Printf("计数法: 长度=%d\n", k7_3)
	fmt.Printf("递归法: 长度=%d\n", k7_4)
}
