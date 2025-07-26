package main

import "fmt"

// 方法一：前缀和 + 哈希表（推荐）
// 时间复杂度：O(n)，空间复杂度：O(n)
func numberOfSubarrays1(nums []int, k int) int {
	count := make(map[int]int)
	count[0] = 1 // 前缀和为0的情况

	prefixSum := 0
	result := 0

	for _, num := range nums {
		if num%2 == 1 {
			prefixSum++
		}
		result += count[prefixSum-k]
		count[prefixSum]++
	}

	return result
}

// 方法二：滑动窗口
// 时间复杂度：O(n)，空间复杂度：O(1)
func numberOfSubarrays2(nums []int, k int) int {
	n := len(nums)
	result := 0
	oddCount := 0
	left := 0
	leftCount := 0 // 记录左边界移动的次数

	for right := 0; right < n; right++ {
		if nums[right]%2 == 1 {
			oddCount++
			leftCount = 0 // 重置左边界计数
		}

		// 当奇数个数等于k时
		if oddCount == k {
			// 移动左边界，统计所有可能的优美子数组
			for left < n && oddCount == k {
				if nums[left]%2 == 1 {
					oddCount--
				}
				left++
				leftCount++
			}
		}

		result += leftCount
	}

	return result
}

// 方法三：数学方法
// 时间复杂度：O(n)，空间复杂度：O(n)
func numberOfSubarrays3(nums []int, k int) int {
	// 记录所有奇数的位置
	oddPositions := []int{-1} // 添加虚拟位置-1
	for i, num := range nums {
		if num%2 == 1 {
			oddPositions = append(oddPositions, i)
		}
	}
	oddPositions = append(oddPositions, len(nums)) // 添加虚拟位置n

	if len(oddPositions)-2 < k {
		return 0
	}

	result := 0
	// 对于每k个连续的奇数
	for i := 1; i+k-1 < len(oddPositions)-1; i++ {
		left := oddPositions[i-1] + 1  // 左边界
		right := oddPositions[i+k] - 1 // 右边界
		result += (oddPositions[i] - left + 1) * (right - oddPositions[i+k-1] + 1)
	}

	return result
}

// 方法四：优化的滑动窗口
// 时间复杂度：O(n)，空间复杂度：O(1)
func numberOfSubarrays4(nums []int, k int) int {
	n := len(nums)
	result := 0
	oddCount := 0
	left := 0
	leftCount := 0 // 记录左边界移动的次数

	for right := 0; right < n; right++ {
		if nums[right]%2 == 1 {
			oddCount++
			leftCount = 0 // 重置左边界计数
		}

		// 当奇数个数等于k时
		if oddCount == k {
			// 移动左边界，统计所有可能的优美子数组
			for left < n && oddCount == k {
				if nums[left]%2 == 1 {
					oddCount--
				}
				left++
				leftCount++
			}
		}

		result += leftCount
	}

	return result
}

func main() {
	fmt.Println("=== 1248. 统计「优美子数组」===")

	// 测试用例1
	nums1 := []int{1, 1, 2, 1, 1}
	k1 := 3
	fmt.Printf("测试用例1: nums=%v, k=%d\n", nums1, k1)
	fmt.Printf("方法一结果: %d\n", numberOfSubarrays1(nums1, k1))
	fmt.Printf("方法二结果: %d\n", numberOfSubarrays2(nums1, k1))
	fmt.Printf("方法三结果: %d\n", numberOfSubarrays3(nums1, k1))
	fmt.Printf("方法四结果: %d\n", numberOfSubarrays4(nums1, k1))
	fmt.Println()

	// 测试用例2
	nums2 := []int{2, 4, 6}
	k2 := 1
	fmt.Printf("测试用例2: nums=%v, k=%d\n", nums2, k2)
	fmt.Printf("方法一结果: %d\n", numberOfSubarrays1(nums2, k2))
	fmt.Printf("方法二结果: %d\n", numberOfSubarrays2(nums2, k2))
	fmt.Printf("方法三结果: %d\n", numberOfSubarrays3(nums2, k2))
	fmt.Printf("方法四结果: %d\n", numberOfSubarrays4(nums2, k2))
	fmt.Println()

	// 测试用例3
	nums3 := []int{2, 2, 2, 1, 2, 2, 1, 2, 2, 2}
	k3 := 2
	fmt.Printf("测试用例3: nums=%v, k=%d\n", nums3, k3)
	fmt.Printf("方法一结果: %d\n", numberOfSubarrays1(nums3, k3))
	fmt.Printf("方法二结果: %d\n", numberOfSubarrays2(nums3, k3))
	fmt.Printf("方法三结果: %d\n", numberOfSubarrays3(nums3, k3))
	fmt.Printf("方法四结果: %d\n", numberOfSubarrays4(nums3, k3))
	fmt.Println()

	// 额外测试用例
	nums4 := []int{1, 2, 3, 4, 5}
	k4 := 2
	fmt.Printf("额外测试: nums=%v, k=%d\n", nums4, k4)
	fmt.Printf("方法一结果: %d\n", numberOfSubarrays1(nums4, k4))
	fmt.Printf("方法二结果: %d\n", numberOfSubarrays2(nums4, k4))
	fmt.Printf("方法三结果: %d\n", numberOfSubarrays3(nums4, k4))
	fmt.Printf("方法四结果: %d\n", numberOfSubarrays4(nums4, k4))
	fmt.Println()

	// 边界测试用例
	nums5 := []int{1}
	k5 := 1
	fmt.Printf("边界测试: nums=%v, k=%d\n", nums5, k5)
	fmt.Printf("方法一结果: %d\n", numberOfSubarrays1(nums5, k5))
	fmt.Printf("方法二结果: %d\n", numberOfSubarrays2(nums5, k5))
	fmt.Printf("方法三结果: %d\n", numberOfSubarrays3(nums5, k5))
	fmt.Printf("方法四结果: %d\n", numberOfSubarrays4(nums5, k5))
	fmt.Println()

	// 复杂测试用例
	nums6 := []int{1, 1, 1, 1, 1}
	k6 := 3
	fmt.Printf("复杂测试: nums=%v, k=%d\n", nums6, k6)
	fmt.Printf("方法一结果: %d\n", numberOfSubarrays1(nums6, k6))
	fmt.Printf("方法二结果: %d\n", numberOfSubarrays2(nums6, k6))
	fmt.Printf("方法三结果: %d\n", numberOfSubarrays3(nums6, k6))
	fmt.Printf("方法四结果: %d\n", numberOfSubarrays4(nums6, k6))
	fmt.Println()

	// 全偶数测试用例
	nums7 := []int{2, 4, 6, 8, 10}
	k7 := 1
	fmt.Printf("全偶数测试: nums=%v, k=%d\n", nums7, k7)
	fmt.Printf("方法一结果: %d\n", numberOfSubarrays1(nums7, k7))
	fmt.Printf("方法二结果: %d\n", numberOfSubarrays2(nums7, k7))
	fmt.Printf("方法三结果: %d\n", numberOfSubarrays3(nums7, k7))
	fmt.Printf("方法四结果: %d\n", numberOfSubarrays4(nums7, k7))
}
