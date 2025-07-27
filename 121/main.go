package main

import "fmt"

// 方法一：一次遍历（推荐）
// 时间复杂度：O(n)，空间复杂度：O(1)
func maxProfit1(prices []int) int {
	if len(prices) < 2 {
		return 0
	}

	minPrice := prices[0]
	maxProfit := 0

	for i := 1; i < len(prices); i++ {
		// 计算当前利润
		currentProfit := prices[i] - minPrice
		if currentProfit > maxProfit {
			maxProfit = currentProfit
		}

		// 更新最小价格
		if prices[i] < minPrice {
			minPrice = prices[i]
		}
	}

	return maxProfit
}

// 方法二：暴力解法
// 时间复杂度：O(n²)，空间复杂度：O(1)
func maxProfit2(prices []int) int {
	maxProfit := 0
	n := len(prices)

	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			profit := prices[j] - prices[i]
			if profit > maxProfit {
				maxProfit = profit
			}
		}
	}

	return maxProfit
}

// 方法三：动态规划
// 时间复杂度：O(n)，空间复杂度：O(n)
func maxProfit3(prices []int) int {
	if len(prices) < 2 {
		return 0
	}

	n := len(prices)
	dp := make([]int, n)
	dp[0] = 0
	minPrice := prices[0]

	for i := 1; i < n; i++ {
		// 当前价格与最小价格的差值
		currentProfit := prices[i] - minPrice
		// 取前一天的最大利润和当前利润的较大值
		dp[i] = max(dp[i-1], currentProfit)
		// 更新最小价格
		if prices[i] < minPrice {
			minPrice = prices[i]
		}
	}

	return dp[n-1]
}

// 方法四：分治法
// 时间复杂度：O(n log n)，空间复杂度：O(log n)
func maxProfit4(prices []int) int {
	if len(prices) < 2 {
		return 0
	}
	return maxProfitDivideAndConquer(prices, 0, len(prices)-1)
}

func maxProfitDivideAndConquer(prices []int, left, right int) int {
	if left >= right {
		return 0
	}

	mid := left + (right-left)/2

	// 递归计算左半部分的最大利润
	leftProfit := maxProfitDivideAndConquer(prices, left, mid)
	// 递归计算右半部分的最大利润
	rightProfit := maxProfitDivideAndConquer(prices, mid+1, right)

	// 计算跨越中点的最大利润
	crossProfit := maxProfitCross(prices, left, mid, right)

	return max(max(leftProfit, rightProfit), crossProfit)
}

func maxProfitCross(prices []int, left, mid, right int) int {
	// 在左半部分找到最小价格
	minPrice := prices[left]
	for i := left; i <= mid; i++ {
		if prices[i] < minPrice {
			minPrice = prices[i]
		}
	}

	// 在右半部分找到最大价格
	maxPrice := prices[mid+1]
	for i := mid + 1; i <= right; i++ {
		if prices[i] > maxPrice {
			maxPrice = prices[i]
		}
	}

	return maxPrice - minPrice
}

// 方法五：优化的动态规划（空间优化）
// 时间复杂度：O(n)，空间复杂度：O(1)
func maxProfit5(prices []int) int {
	if len(prices) < 2 {
		return 0
	}

	minPrice := prices[0]
	maxProfit := 0

	for i := 1; i < len(prices); i++ {
		// 更新最大利润
		maxProfit = max(maxProfit, prices[i]-minPrice)
		// 更新最小价格
		minPrice = min(minPrice, prices[i])
	}

	return maxProfit
}

// 方法六：单调栈思想
// 时间复杂度：O(n)，空间复杂度：O(1)
func maxProfit6(prices []int) int {
	if len(prices) < 2 {
		return 0
	}

	maxProfit := 0
	minPrice := prices[0]

	for i := 1; i < len(prices); i++ {
		// 计算当前利润
		currentProfit := prices[i] - minPrice
		if currentProfit > maxProfit {
			maxProfit = currentProfit
		}
		// 更新最小价格
		if prices[i] < minPrice {
			minPrice = prices[i]
		}
	}

	return maxProfit
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
	fmt.Println("=== 121. 买卖股票的最佳时机 ===")

	// 测试用例1
	prices1 := []int{7, 1, 5, 3, 6, 4}
	fmt.Printf("测试用例1: prices=%v\n", prices1)
	fmt.Printf("方法一结果: %d\n", maxProfit1(prices1))
	fmt.Printf("方法二结果: %d\n", maxProfit2(prices1))
	fmt.Printf("方法三结果: %d\n", maxProfit3(prices1))
	fmt.Printf("方法四结果: %d\n", maxProfit4(prices1))
	fmt.Printf("方法五结果: %d\n", maxProfit5(prices1))
	fmt.Printf("方法六结果: %d\n", maxProfit6(prices1))
	fmt.Println()

	// 测试用例2
	prices2 := []int{7, 6, 4, 3, 1}
	fmt.Printf("测试用例2: prices=%v\n", prices2)
	fmt.Printf("方法一结果: %d\n", maxProfit1(prices2))
	fmt.Printf("方法二结果: %d\n", maxProfit2(prices2))
	fmt.Printf("方法三结果: %d\n", maxProfit3(prices2))
	fmt.Printf("方法四结果: %d\n", maxProfit4(prices2))
	fmt.Printf("方法五结果: %d\n", maxProfit5(prices2))
	fmt.Printf("方法六结果: %d\n", maxProfit6(prices2))
	fmt.Println()

	// 测试用例3
	prices3 := []int{1, 2, 3, 4, 5}
	fmt.Printf("测试用例3: prices=%v\n", prices3)
	fmt.Printf("方法一结果: %d\n", maxProfit1(prices3))
	fmt.Printf("方法二结果: %d\n", maxProfit2(prices3))
	fmt.Printf("方法三结果: %d\n", maxProfit3(prices3))
	fmt.Printf("方法四结果: %d\n", maxProfit4(prices3))
	fmt.Printf("方法五结果: %d\n", maxProfit5(prices3))
	fmt.Printf("方法六结果: %d\n", maxProfit6(prices3))
	fmt.Println()

	// 额外测试用例
	prices4 := []int{3, 2, 6, 5, 0, 3}
	fmt.Printf("额外测试: prices=%v\n", prices4)
	fmt.Printf("方法一结果: %d\n", maxProfit1(prices4))
	fmt.Printf("方法二结果: %d\n", maxProfit2(prices4))
	fmt.Printf("方法三结果: %d\n", maxProfit3(prices4))
	fmt.Printf("方法四结果: %d\n", maxProfit4(prices4))
	fmt.Printf("方法五结果: %d\n", maxProfit5(prices4))
	fmt.Printf("方法六结果: %d\n", maxProfit6(prices4))
	fmt.Println()

	// 边界测试用例
	prices5 := []int{1}
	fmt.Printf("边界测试: prices=%v\n", prices5)
	fmt.Printf("方法一结果: %d\n", maxProfit1(prices5))
	fmt.Printf("方法二结果: %d\n", maxProfit2(prices5))
	fmt.Printf("方法三结果: %d\n", maxProfit3(prices5))
	fmt.Printf("方法四结果: %d\n", maxProfit4(prices5))
	fmt.Printf("方法五结果: %d\n", maxProfit5(prices5))
	fmt.Printf("方法六结果: %d\n", maxProfit6(prices5))
	fmt.Println()

	// 复杂测试用例
	prices6 := []int{2, 4, 1, 7, 8, 3, 9, 5, 6}
	fmt.Printf("复杂测试: prices=%v\n", prices6)
	fmt.Printf("方法一结果: %d\n", maxProfit1(prices6))
	fmt.Printf("方法二结果: %d\n", maxProfit2(prices6))
	fmt.Printf("方法三结果: %d\n", maxProfit3(prices6))
	fmt.Printf("方法四结果: %d\n", maxProfit4(prices6))
	fmt.Printf("方法五结果: %d\n", maxProfit5(prices6))
	fmt.Printf("方法六结果: %d\n", maxProfit6(prices6))
	fmt.Println()

	// 单点测试用例
	prices7 := []int{5, 5, 5, 5, 5}
	fmt.Printf("单点测试: prices=%v\n", prices7)
	fmt.Printf("方法一结果: %d\n", maxProfit1(prices7))
	fmt.Printf("方法二结果: %d\n", maxProfit2(prices7))
	fmt.Printf("方法三结果: %d\n", maxProfit3(prices7))
	fmt.Printf("方法四结果: %d\n", maxProfit4(prices7))
	fmt.Printf("方法五结果: %d\n", maxProfit5(prices7))
	fmt.Printf("方法六结果: %d\n", maxProfit6(prices7))
	fmt.Println()

	// 大数测试用例
	prices8 := []int{10000, 9999, 9998, 9997, 9996}
	fmt.Printf("大数测试: prices=%v\n", prices8)
	fmt.Printf("方法一结果: %d\n", maxProfit1(prices8))
	fmt.Printf("方法二结果: %d\n", maxProfit2(prices8))
	fmt.Printf("方法三结果: %d\n", maxProfit3(prices8))
	fmt.Printf("方法四结果: %d\n", maxProfit4(prices8))
	fmt.Printf("方法五结果: %d\n", maxProfit5(prices8))
	fmt.Printf("方法六结果: %d\n", maxProfit6(prices8))
}
