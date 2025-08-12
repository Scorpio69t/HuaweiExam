package main

import (
	"fmt"
	"sort"
)

// 方法一：哈希表解法（推荐）
// 时间复杂度：O(n)，空间复杂度：O(n)
func twoSum(nums []int, target int) []int {
	numMap := make(map[int]int)

	for i, num := range nums {
		complement := target - num
		if j, exists := numMap[complement]; exists {
			return []int{j, i}
		}
		numMap[num] = i
	}

	return []int{}
}

// 方法二：暴力枚举解法
// 时间复杂度：O(n²)，空间复杂度：O(1)
func twoSumBruteForce(nums []int, target int) []int {
	n := len(nums)
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return []int{}
}

// 方法三：排序 + 双指针解法
// 时间复杂度：O(n log n)，空间复杂度：O(n)
func twoSumTwoPointers(nums []int, target int) []int {
	n := len(nums)
	
	// 创建索引数组，用于保存原始索引
	indices := make([]int, n)
	for i := range indices {
		indices[i] = i
	}
	
	// 根据nums值对索引数组排序
	sort.Slice(indices, func(i, j int) bool {
		return nums[indices[i]] < nums[indices[j]]
	})
	
	// 双指针查找
	left, right := 0, n-1
	for left < right {
		sum := nums[indices[left]] + nums[indices[right]]
		if sum == target {
			return []int{indices[left], indices[right]}
		} else if sum < target {
			left++
		} else {
			right--
		}
	}
	
	return []int{}
}

// 方法四：优化的哈希表解法（处理重复元素）
func twoSumOptimized(nums []int, target int) []int {
	numMap := make(map[int]int)
	
	for i, num := range nums {
		complement := target - num
		
		// 检查补数是否存在
		if j, exists := numMap[complement]; exists {
			// 确保不使用同一个元素
			if j != i {
				return []int{j, i}
			}
		}
		
		// 将当前数字加入哈希表
		numMap[num] = i
	}
	
	return []int{}
}

// 方法五：一次遍历的哈希表解法
func twoSumOnePass(nums []int, target int) []int {
	numMap := make(map[int]int)
	
	for i, num := range nums {
		complement := target - num
		
		// 先查找补数，再添加当前数字
		if j, exists := numMap[complement]; exists {
			return []int{j, i}
		}
		
		numMap[num] = i
	}
	
	return []int{}
}

// 辅助函数：验证结果是否正确
func validateResult(nums []int, target int, result []int) bool {
	if len(result) != 2 {
		return false
	}
	
	i, j := result[0], result[1]
	if i < 0 || i >= len(nums) || j < 0 || j >= len(nums) || i == j {
		return false
	}
	
	return nums[i]+nums[j] == target
}

// 辅助函数：打印数组和结果
func printResult(nums []int, target int, result []int, method string) {
	fmt.Printf("%s: nums=%v, target=%d\n", method, nums, target)
	if len(result) == 2 {
		fmt.Printf("结果: [%d,%d], 验证: nums[%d]+nums[%d]=%d+%d=%d\n", 
			result[0], result[1], result[0], result[1], 
			nums[result[0]], nums[result[1]], nums[result[0]]+nums[result[1]])
	} else {
		fmt.Printf("结果: %v (未找到)\n", result)
	}
	fmt.Println()
}

func main() {
	fmt.Println("=== 1. 两数之和 ===")

	// 测试用例1
	nums1 := []int{2, 7, 11, 15}
	target1 := 9
	fmt.Printf("测试用例1: nums=%v, target=%d\n", nums1, target1)
	
	result1 := twoSum(nums1, target1)
	fmt.Printf("哈希表解法结果: %v\n", result1)
	fmt.Printf("暴力枚举解法结果: %v\n", twoSumBruteForce(nums1, target1))
	fmt.Printf("双指针解法结果: %v\n", twoSumTwoPointers(nums1, target1))
	fmt.Printf("优化哈希表解法结果: %v\n", twoSumOptimized(nums1, target1))
	fmt.Printf("一次遍历解法结果: %v\n", twoSumOnePass(nums1, target1))
	fmt.Println()

	// 测试用例2
	nums2 := []int{3, 2, 4}
	target2 := 6
	fmt.Printf("测试用例2: nums=%v, target=%d\n", nums2, target2)
	
	result2 := twoSum(nums2, target2)
	fmt.Printf("哈希表解法结果: %v\n", result2)
	fmt.Printf("暴力枚举解法结果: %v\n", twoSumBruteForce(nums2, target2))
	fmt.Printf("双指针解法结果: %v\n", twoSumTwoPointers(nums2, target2))
	fmt.Printf("优化哈希表解法结果: %v\n", twoSumOptimized(nums2, target2))
	fmt.Printf("一次遍历解法结果: %v\n", twoSumOnePass(nums2, target2))
	fmt.Println()

	// 测试用例3
	nums3 := []int{3, 3}
	target3 := 6
	fmt.Printf("测试用例3: nums=%v, target=%d\n", nums3, target3)
	
	result3 := twoSum(nums3, target3)
	fmt.Printf("哈希表解法结果: %v\n", result3)
	fmt.Printf("暴力枚举解法结果: %v\n", twoSumBruteForce(nums3, target3))
	fmt.Printf("双指针解法结果: %v\n", twoSumTwoPointers(nums3, target3))
	fmt.Printf("优化哈希表解法结果: %v\n", twoSumOptimized(nums3, target3))
	fmt.Printf("一次遍历解法结果: %v\n", twoSumOnePass(nums3, target3))
	fmt.Println()

	// 边界测试用例
	testCases := []struct {
		nums   []int
		target int
		desc   string
	}{
		{[]int{1, 5, 8, 10, 13}, 18, "大数组测试"},
		{[]int{-1, -2, -3, -4, -5}, -8, "负数测试"},
		{[]int{0, 0}, 0, "零值测试"},
		{[]int{1, 2}, 3, "两元素测试"},
		{[]int{1, 2, 3, 4, 5}, 9, "连续数组测试"},
		{[]int{1000000000, -1000000000}, 0, "大数测试"},
	}

	for _, tc := range testCases {
		result := twoSum(tc.nums, tc.target)
		fmt.Printf("%s: nums=%v, target=%d, 结果=%v\n", tc.desc, tc.nums, tc.target, result)
	}

	// 算法正确性验证
	fmt.Println("\n=== 算法正确性验证 ===")
	verifyNums := []int{2, 7, 11, 15}
	verifyTarget := 9
	
	fmt.Printf("验证数组: %v, 目标值: %d\n", verifyNums, verifyTarget)
	verifyResult1 := twoSum(verifyNums, verifyTarget)
	verifyResult2 := twoSumBruteForce(verifyNums, verifyTarget)
	verifyResult3 := twoSumTwoPointers(verifyNums, verifyTarget)
	verifyResult4 := twoSumOptimized(verifyNums, verifyTarget)
	verifyResult5 := twoSumOnePass(verifyNums, verifyTarget)

	fmt.Printf("哈希表解法: %v\n", verifyResult1)
	fmt.Printf("暴力枚举解法: %v\n", verifyResult2)
	fmt.Printf("双指针解法: %v\n", verifyResult3)
	fmt.Printf("优化哈希表解法: %v\n", verifyResult4)
	fmt.Printf("一次遍历解法: %v\n", verifyResult5)

	// 验证所有解法结果一致
	if len(verifyResult1) == 2 && len(verifyResult2) == 2 && len(verifyResult3) == 2 && 
	   len(verifyResult4) == 2 && len(verifyResult5) == 2 {
		fmt.Println("✅ 所有解法都找到了答案！")
		
		// 验证结果正确性
		if validateResult(verifyNums, verifyTarget, verifyResult1) {
			fmt.Println("✅ 结果验证通过！")
		} else {
			fmt.Println("❌ 结果验证失败！")
		}
	} else {
		fmt.Println("❌ 解法结果不一致，需要检查！")
	}

	// 性能测试
	fmt.Println("\n=== 性能测试 ===")
	largeNums := make([]int, 10000)
	for i := range largeNums {
		largeNums[i] = i
	}
	largeTarget := 19998 // 9999 + 9999

	fmt.Printf("大数组测试: 长度=%d, 目标值=%d\n", len(largeNums), largeTarget)
	result := twoSum(largeNums, largeTarget)
	fmt.Printf("哈希表解法结果: %v\n", result)
	
	if len(result) == 2 {
		fmt.Printf("验证: nums[%d]+nums[%d]=%d+%d=%d\n", 
			result[0], result[1], largeNums[result[0]], largeNums[result[1]], 
			largeNums[result[0]]+largeNums[result[1]])
	}
}
