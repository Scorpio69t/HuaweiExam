package main

import "fmt"

// 方法一：滑动窗口解法（推荐）
// 时间复杂度：O(n)，空间复杂度：O(1)
func longestOnes(nums []int, k int) int {
	left, zeros := 0, 0
	maxLen := 0

	for right := 0; right < len(nums); right++ {
		if nums[right] == 0 {
			zeros++
		}

		// 当0的个数超过k时，收缩左边界
		for zeros > k {
			if nums[left] == 0 {
				zeros--
			}
			left++
		}

		maxLen = max(maxLen, right-left+1)
	}

	return maxLen
}

// 方法二：前缀和 + 二分查找
// 时间复杂度：O(n log n)，空间复杂度：O(n)
func longestOnesBinarySearch(nums []int, k int) int {
	n := len(nums)
	prefix := make([]int, n+1)
	
	// 计算前缀和（0的个数）
	for i := 0; i < n; i++ {
		prefix[i+1] = prefix[i] + (1 - nums[i]) // 1-nums[i]：0变为1，1变为0
	}

	maxLen := 0
	for i := 0; i < n; i++ {
		// 二分查找最远的j，使得区间[i,j]中0的个数不超过k
		left, right := i, n-1
		for left <= right {
			mid := left + (right-left)/2
			zeros := prefix[mid+1] - prefix[i]
			if zeros <= k {
				maxLen = max(maxLen, mid-i+1)
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}

	return maxLen
}

// 方法三：动态规划
// 时间复杂度：O(nk)，空间复杂度：O(nk)
func longestOnesDP(nums []int, k int) int {
	n := len(nums)
	// dp[i][j]表示以位置i结尾，使用j次翻转机会的最长连续1长度
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, k+1)
	}

	// 初始化第一行
	if nums[0] == 1 {
		dp[0][0] = 1
	} else if k > 0 {
		dp[0][1] = 1
	}

	maxLen := 0
	for j := 0; j <= k; j++ {
		maxLen = max(maxLen, dp[0][j])
	}

	// 状态转移
	for i := 1; i < n; i++ {
		for j := 0; j <= k; j++ {
			if nums[i] == 1 {
				// 当前位置是1，可以直接接上
				dp[i][j] = dp[i-1][j] + 1
			} else if j > 0 {
				// 当前位置是0，需要消耗一次翻转机会
				dp[i][j] = dp[i-1][j-1] + 1
			}
			maxLen = max(maxLen, dp[i][j])
		}
	}

	return maxLen
}

// 辅助函数：返回两个数的最大值
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println("=== 1004. 最大连续1的个数 III ===")

	// 测试用例1
	nums1 := []int{1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0}
	k1 := 2
	result1 := longestOnes(nums1, k1)
	fmt.Printf("测试用例1: nums=%v, k=%d\n", nums1, k1)
	fmt.Printf("滑动窗口解法结果: %d\n", result1)
	fmt.Printf("二分查找解法结果: %d\n", longestOnesBinarySearch(nums1, k1))
	fmt.Printf("动态规划解法结果: %d\n", longestOnesDP(nums1, k1))
	fmt.Println()

	// 测试用例2
	nums2 := []int{0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 1, 1, 0, 0, 0, 1, 1, 1, 1}
	k2 := 3
	result2 := longestOnes(nums2, k2)
	fmt.Printf("测试用例2: nums=%v, k=%d\n", nums2, k2)
	fmt.Printf("滑动窗口解法结果: %d\n", result2)
	fmt.Printf("二分查找解法结果: %d\n", longestOnesBinarySearch(nums2, k2))
	fmt.Printf("动态规划解法结果: %d\n", longestOnesDP(nums2, k2))
	fmt.Println()

	// 边界测试用例
	testCases := []struct {
		nums []int
		k    int
		desc string
	}{
		{[]int{1, 1, 1, 1, 1}, 0, "全1数组，k=0"},
		{[]int{0, 0, 0, 0, 0}, 2, "全0数组，k=2"},
		{[]int{1, 0, 1, 0, 1}, 1, "交替数组，k=1"},
		{[]int{0}, 1, "单个0，k=1"},
		{[]int{1}, 0, "单个1，k=0"},
	}

	for _, tc := range testCases {
		result := longestOnes(tc.nums, tc.k)
		fmt.Printf("%s: nums=%v, k=%d, 结果=%d\n", tc.desc, tc.nums, tc.k, result)
	}

	fmt.Println("\n=== 性能测试 ===")
	// 大数组性能测试
	largeNums := make([]int, 10000)
	for i := range largeNums {
		if i%3 == 0 {
			largeNums[i] = 0
		} else {
			largeNums[i] = 1
		}
	}
	largeK := 100

	fmt.Printf("大数组测试: 长度=%d, k=%d\n", len(largeNums), largeK)
	result := longestOnes(largeNums, largeK)
	fmt.Printf("滑动窗口解法结果: %d\n", result)
}
