package main

import "fmt"

// 方法一：前缀和 + 哈希表解法（推荐）
// 时间复杂度：O(n)，空间复杂度：O(n)
func subarraySum(nums []int, k int) int {
	count := 0
	prefixSum := 0
	prefixMap := make(map[int]int)
	prefixMap[0] = 1 // 初始化，空数组的前缀和为0

	for _, num := range nums {
		prefixSum += num
		// 查找前缀和为 prefixSum - k 的个数
		if val, exists := prefixMap[prefixSum-k]; exists {
			count += val
		}
		// 将当前前缀和加入哈希表
		prefixMap[prefixSum]++
	}

	return count
}

// 方法二：暴力枚举解法
// 时间复杂度：O(n²)，空间复杂度：O(1)
func subarraySumBruteForce(nums []int, k int) int {
	count := 0
	n := len(nums)

	for i := 0; i < n; i++ {
		sum := 0
		for j := i; j < n; j++ {
			sum += nums[j]
			if sum == k {
				count++
			}
		}
	}

	return count
}

// 方法三：滑动窗口解法（仅适用于非负数数组）
// 时间复杂度：O(n)，空间复杂度：O(1)
func subarraySumSlidingWindow(nums []int, k int) int {
	// 检查是否所有元素都是非负数
	for _, num := range nums {
		if num < 0 {
			return -1 // 表示不适用此方法
		}
	}

	count := 0
	left, sum := 0, 0

	for right := 0; right < len(nums); right++ {
		sum += nums[right]

		// 当窗口和大于k时，收缩左边界
		for sum > k && left <= right {
			sum -= nums[left]
			left++
		}

		// 如果窗口和等于k，增加计数
		if sum == k {
			count++
		}
	}

	return count
}

// 方法四：优化的前缀和解法（更清晰的实现）
func subarraySumOptimized(nums []int, k int) int {
	count := 0
	prefixSum := 0
	prefixMap := make(map[int]int)
	prefixMap[0] = 1

	for i := 0; i < len(nums); i++ {
		prefixSum += nums[i]
		target := prefixSum - k

		// 查找目标前缀和的个数
		if freq, exists := prefixMap[target]; exists {
			count += freq
		}

		// 更新当前前缀和的频率
		prefixMap[prefixSum]++
	}

	return count
}

func main() {
	fmt.Println("=== 560. 和为 K 的子数组 ===")

	// 测试用例1
	nums1 := []int{1, 1, 1}
	k1 := 2
	fmt.Printf("测试用例1: nums=%v, k=%d\n", nums1, k1)
	fmt.Printf("前缀和+哈希表解法结果: %d\n", subarraySum(nums1, k1))
	fmt.Printf("暴力枚举解法结果: %d\n", subarraySumBruteForce(nums1, k1))
	fmt.Printf("优化前缀和解法结果: %d\n", subarraySumOptimized(nums1, k1))
	fmt.Println()

	// 测试用例2
	nums2 := []int{1, 2, 3}
	k2 := 3
	fmt.Printf("测试用例2: nums=%v, k=%d\n", nums2, k2)
	fmt.Printf("前缀和+哈希表解法结果: %d\n", subarraySum(nums2, k2))
	fmt.Printf("暴力枚举解法结果: %d\n", subarraySumBruteForce(nums2, k2))
	fmt.Printf("优化前缀和解法结果: %d\n", subarraySumOptimized(nums2, k2))
	fmt.Println()

	// 测试用例3（包含负数）
	nums3 := []int{1, -1, 0}
	k3 := 0
	fmt.Printf("测试用例3: nums=%v, k=%d\n", nums3, k3)
	fmt.Printf("前缀和+哈希表解法结果: %d\n", subarraySum(nums3, k3))
	fmt.Printf("暴力枚举解法结果: %d\n", subarraySumBruteForce(nums3, k3))
	fmt.Printf("优化前缀和解法结果: %d\n", subarraySumOptimized(nums3, k3))
	fmt.Printf("滑动窗口解法结果: %d (不适用于负数)\n", subarraySumSlidingWindow(nums3, k3))
	fmt.Println()

	// 边界测试用例
	testCases := []struct {
		nums []int
		k    int
		desc string
	}{
		{[]int{1}, 1, "单个元素"},
		{[]int{1, 2, 3, 4, 5}, 9, "连续数组"},
		{[]int{0, 0, 0, 0, 0}, 0, "全零数组"},
		{[]int{1, 2, 3, 4, 5}, 15, "整个数组"},
		{[]int{-1, -1, 1}, 0, "包含负数"},
		{[]int{1, 2, 3, 4, 5}, 100, "目标值过大"},
	}

	for _, tc := range testCases {
		result := subarraySum(tc.nums, tc.k)
		fmt.Printf("%s: nums=%v, k=%d, 结果=%d\n", tc.desc, tc.nums, tc.k, result)
	}

	fmt.Println("\n=== 性能测试 ===")
	// 大数组性能测试
	largeNums := make([]int, 1000)
	for i := range largeNums {
		largeNums[i] = i % 10 // 生成0-9的循环数组
	}
	largeK := 15

	fmt.Printf("大数组测试: 长度=%d, k=%d\n", len(largeNums), largeK)
	result := subarraySum(largeNums, largeK)
	fmt.Printf("前缀和+哈希表解法结果: %d\n", result)

	// 算法正确性验证
	fmt.Println("\n=== 算法正确性验证 ===")
	smallNums := []int{1, 2, 3, 4, 5}
	smallK := 7
	fmt.Printf("验证数组: nums=%v, k=%d\n", smallNums, smallK)
	fmt.Printf("前缀和解法: %d\n", subarraySum(smallNums, smallK))
	fmt.Printf("暴力解法: %d\n", subarraySumBruteForce(smallNums, smallK))
	fmt.Printf("优化解法: %d\n", subarraySumOptimized(smallNums, smallK))
	
	// 验证所有解法结果一致
	if subarraySum(smallNums, smallK) == subarraySumBruteForce(smallNums, smallK) &&
		subarraySum(smallNums, smallK) == subarraySumOptimized(smallNums, smallK) {
		fmt.Println("✅ 所有解法结果一致，算法正确！")
	} else {
		fmt.Println("❌ 解法结果不一致，需要检查！")
	}
}
