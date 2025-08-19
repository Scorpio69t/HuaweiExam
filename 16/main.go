package main

import (
	"fmt"
	"sort"
)

// threeSumClosest 最接近的三数之和 - 排序+双指针法
// 时间复杂度: O(n²)，其中n是数组长度
// 空间复杂度: O(1)
func threeSumClosest(nums []int, target int) int {
	// 先对数组排序
	sort.Ints(nums)

	closestSum := nums[0] + nums[1] + nums[2]
	minDiff := abs(closestSum - target)

	// 固定第一个数，使用双指针寻找另外两个数
	for i := 0; i < len(nums)-2; i++ {
		// 跳过重复元素
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		left := i + 1
		right := len(nums) - 1

		for left < right {
			sum := nums[i] + nums[left] + nums[right]
			diff := abs(sum - target)

			// 更新最接近的和
			if diff < minDiff {
				minDiff = diff
				closestSum = sum
			}

			// 根据sum与target的关系移动指针
			if sum < target {
				left++
			} else if sum > target {
				right--
			} else {
				// 找到完全相等的，直接返回
				return sum
			}
		}
	}

	return closestSum
}

// threeSumClosestOptimized 优化版本 - 提前剪枝
// 时间复杂度: O(n²)
// 空间复杂度: O(1)
func threeSumClosestOptimized(nums []int, target int) int {
	sort.Ints(nums)

	closestSum := nums[0] + nums[1] + nums[2]
	minDiff := abs(closestSum - target)

	for i := 0; i < len(nums)-2; i++ {
		// 跳过重复元素
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		left := i + 1
		right := len(nums) - 1

		// 提前剪枝：如果当前最小值已经为0，直接返回
		if minDiff == 0 {
			return closestSum
		}

		for left < right {
			sum := nums[i] + nums[left] + nums[right]
			diff := abs(sum - target)

			if diff < minDiff {
				minDiff = diff
				closestSum = sum
			}

			if sum < target {
				left++
				// 跳过重复元素
				for left < right && nums[left] == nums[left-1] {
					left++
				}
			} else if sum > target {
				right--
				// 跳过重复元素
				for left < right && nums[right] == nums[right+1] {
					right--
				}
			} else {
				return sum
			}
		}
	}

	return closestSum
}

// threeSumClosestBinarySearch 二分查找版本
// 时间复杂度: O(n² log n)
// 空间复杂度: O(1)
func threeSumClosestBinarySearch(nums []int, target int) int {
	sort.Ints(nums)

	closestSum := nums[0] + nums[1] + nums[2]
	minDiff := abs(closestSum - target)

	for i := 0; i < len(nums)-2; i++ {
		for j := i + 1; j < len(nums)-1; j++ {
			// 使用二分查找寻找第三个数
			remaining := target - nums[i] - nums[j]
			left := j + 1
			right := len(nums) - 1

			// 二分查找最接近remaining的数
			for left <= right {
				mid := left + (right-left)/2
				sum := nums[i] + nums[j] + nums[mid]
				diff := abs(sum - target)

				if diff < minDiff {
					minDiff = diff
					closestSum = sum
				}

				if nums[mid] < remaining {
					left = mid + 1
				} else if nums[mid] > remaining {
					right = mid - 1
				} else {
					return sum
				}
			}
		}
	}

	return closestSum
}

// threeSumClosestBruteForce 暴力解法 - 三重循环
// 时间复杂度: O(n³)
// 空间复杂度: O(1)
func threeSumClosestBruteForce(nums []int, target int) int {
	closestSum := nums[0] + nums[1] + nums[2]
	minDiff := abs(closestSum - target)

	for i := 0; i < len(nums)-2; i++ {
		for j := i + 1; j < len(nums)-1; j++ {
			for k := j + 1; k < len(nums); k++ {
				sum := nums[i] + nums[j] + nums[k]
				diff := abs(sum - target)

				if diff < minDiff {
					minDiff = diff
					closestSum = sum
				}
			}
		}
	}

	return closestSum
}

// threeSumClosestRecursive 递归方法 - 回溯思想
// 时间复杂度: O(n³)
// 空间复杂度: O(n)，递归调用栈
func threeSumClosestRecursive(nums []int, target int) int {
	closestSum := nums[0] + nums[1] + nums[2]
	minDiff := abs(closestSum - target)

	var backtrack func(index, count, currentSum int)
	backtrack = func(index, count, currentSum int) {
		if count == 3 {
			diff := abs(currentSum - target)
			if diff < minDiff {
				minDiff = diff
				closestSum = currentSum
			}
			return
		}

		if index >= len(nums) {
			return
		}

		// 选择当前数
		backtrack(index+1, count+1, currentSum+nums[index])
		// 不选择当前数
		backtrack(index+1, count, currentSum)
	}

	backtrack(0, 0, 0)
	return closestSum
}

// abs 计算绝对值
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	// 测试用例1
	nums1 := []int{-1, 2, 1, -4}
	target1 := 1
	result1 := threeSumClosest(nums1, target1)
	fmt.Printf("示例1: nums = %v, target = %d\n", nums1, target1)
	fmt.Printf("输出: %d\n", result1)
	fmt.Printf("期望: 2\n")
	fmt.Printf("结果: %t\n", result1 == 2)
	fmt.Println()

	// 测试用例2
	nums2 := []int{0, 0, 0}
	target2 := 1
	result2 := threeSumClosest(nums2, target2)
	fmt.Printf("示例2: nums = %v, target = %d\n", nums2, target2)
	fmt.Printf("输出: %d\n", result2)
	fmt.Printf("期望: 0\n")
	fmt.Printf("结果: %t\n", result2 == 0)
	fmt.Println()

	// 额外测试用例
	nums3 := []int{1, 1, 1, 0}
	target3 := -100
	result3 := threeSumClosest(nums3, target3)
	fmt.Printf("额外测试: nums = %v, target = %d\n", nums3, target3)
	fmt.Printf("输出: %d\n", result3)
	fmt.Printf("期望: 2\n")
	fmt.Printf("结果: %t\n", result3 == 2)
	fmt.Println()

	// 测试优化版本
	fmt.Println("=== 优化版本测试 ===")
	result1Opt := threeSumClosestOptimized(nums1, target1)
	result2Opt := threeSumClosestOptimized(nums2, target2)
	fmt.Printf("优化版本示例1: %d\n", result1Opt)
	fmt.Printf("优化版本示例2: %d\n", result2Opt)
	fmt.Printf("结果一致: %t\n", result1Opt == result1 && result2Opt == result2)
	fmt.Println()

	// 测试二分查找版本
	fmt.Println("=== 二分查找版本测试 ===")
	result1Bin := threeSumClosestBinarySearch(nums1, target1)
	result2Bin := threeSumClosestBinarySearch(nums2, target2)
	fmt.Printf("二分查找版本示例1: %d\n", result1Bin)
	fmt.Printf("二分查找版本示例2: %d\n", result2Bin)
	fmt.Printf("结果一致: %t\n", result1Bin == result1 && result2Bin == result2)
	fmt.Println()

	// 测试暴力解法
	fmt.Println("=== 暴力解法测试 ===")
	result1BF := threeSumClosestBruteForce(nums1, target1)
	result2BF := threeSumClosestBruteForce(nums2, target2)
	fmt.Printf("暴力解法示例1: %d\n", result1BF)
	fmt.Printf("暴力解法示例2: %d\n", result2BF)
	fmt.Printf("结果一致: %t\n", result1BF == result1 && result2BF == result2)
	fmt.Println()

	// 测试递归方法
	fmt.Println("=== 递归方法测试 ===")
	result1Rec := threeSumClosestRecursive(nums1, target1)
	result2Rec := threeSumClosestRecursive(nums2, target2)
	fmt.Printf("递归方法示例1: %d\n", result1Rec)
	fmt.Printf("递归方法示例2: %d\n", result2Rec)
	fmt.Printf("结果一致: %t\n", result1Rec == result1 && result2Rec == result2)
	fmt.Println()

	// 边界值测试
	fmt.Println("=== 边界值测试 ===")
	boundaryTests := []struct {
		nums   []int
		target int
	}{
		{[]int{1, 1, 1}, 3},                 // 最小值
		{[]int{1000, 1000, 1000}, 3000},     // 最大值
		{[]int{-1000, -1000, -1000}, -3000}, // 负值
		{[]int{0, 0, 0}, 0},                 // 零值
		{[]int{1, 2, 3}, 6},                 // 完全相等
		{[]int{1, 2, 3}, 5},                 // 接近但不相等
	}

	for i, test := range boundaryTests {
		result := threeSumClosest(test.nums, test.target)
		fmt.Printf("测试%d: nums = %v, target = %d, result = %d\n", i+1, test.nums, test.target, result)
	}
}
