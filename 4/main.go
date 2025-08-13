package main

import (
	"fmt"
	"math"
)

// findMedianSortedArrays 寻找两个正序数组的中位数
// 时间复杂度: O(log(min(m,n)))
// 空间复杂度: O(1)
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	// 确保nums1是较短的数组，这样可以减少二分查找的范围
	if len(nums1) > len(nums2) {
		nums1, nums2 = nums2, nums1
	}

	m, n := len(nums1), len(nums2)
	left, right := 0, m

	// 二分查找
	for left <= right {
		// 前一部分包含 nums1[0 .. i-1] 和 nums2[0 .. j-1]
		// 后一部分包含 nums1[i .. m-1] 和 nums2[j .. n-1]
		i := (left + right) / 2
		j := (m+n+1)/2 - i

		// nums1[i-1], nums1[i], nums2[j-1], nums2[j] 分别表示
		// nums1 和 nums2 中前一部分的最大值和后一部分的最小值

		// 当 i = 0 时，nums1[i-1] 不存在，我们将其设为负无穷
		// 当 i = m 时，nums1[i] 不存在，我们将其设为正无穷
		nums1LeftMax := math.MinInt32
		if i > 0 {
			nums1LeftMax = nums1[i-1]
		}
		nums1RightMin := math.MaxInt32
		if i < m {
			nums1RightMin = nums1[i]
		}

		// 当 j = 0 时，nums2[j-1] 不存在，我们将其设为负无穷
		// 当 j = n 时，nums2[j] 不存在，我们将其设为正无穷
		nums2LeftMax := math.MinInt32
		if j > 0 {
			nums2LeftMax = nums2[j-1]
		}
		nums2RightMin := math.MaxInt32
		if j < n {
			nums2RightMin = nums2[j]
		}

		// 前一部分的最大值应该小于等于后一部分的最小值
		if nums1LeftMax <= nums2RightMin && nums2LeftMax <= nums1RightMin {
			// 找到了正确的分割点
			if (m+n)%2 == 1 {
				// 奇数个元素，中位数是前一部分的最大值
				return float64(max(nums1LeftMax, nums2LeftMax))
			} else {
				// 偶数个元素，中位数是前一部分的最大值和后一部分的最小值的平均值
				return float64(max(nums1LeftMax, nums2LeftMax)+min(nums1RightMin, nums2RightMin)) / 2.0
			}
		} else if nums1LeftMax > nums2RightMin {
			// nums1 的前一部分太大，需要减小
			right = i - 1
		} else {
			// nums1 的前一部分太小，需要增大
			left = i + 1
		}
	}

	// 正常情况下不会到达这里
	return 0.0
}

// 辅助函数
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	// 测试用例1
	nums1 := []int{1, 3}
	nums2 := []int{2}
	result1 := findMedianSortedArrays(nums1, nums2)
	fmt.Printf("示例1: nums1 = %v, nums2 = %v\n", nums1, nums2)
	fmt.Printf("输出: %.5f\n", result1)
	fmt.Println("期望: 2.00000")
	fmt.Println()

	// 测试用例2
	nums3 := []int{1, 2}
	nums4 := []int{3, 4}
	result2 := findMedianSortedArrays(nums3, nums4)
	fmt.Printf("示例2: nums1 = %v, nums2 = %v\n", nums3, nums4)
	fmt.Printf("输出: %.5f\n", result2)
	fmt.Println("期望: 2.50000")
	fmt.Println()

	// 额外测试用例
	nums5 := []int{0, 0}
	nums6 := []int{0, 0}
	result3 := findMedianSortedArrays(nums5, nums6)
	fmt.Printf("额外测试: nums1 = %v, nums2 = %v\n", nums5, nums6)
	fmt.Printf("输出: %.5f\n", result3)
	fmt.Println("期望: 0.00000")
	fmt.Println()

	nums7 := []int{}
	nums8 := []int{1}
	result4 := findMedianSortedArrays(nums7, nums8)
	fmt.Printf("额外测试: nums1 = %v, nums2 = %v\n", nums7, nums8)
	fmt.Printf("输出: %.5f\n", result4)
	fmt.Println("期望: 1.00000")
}
