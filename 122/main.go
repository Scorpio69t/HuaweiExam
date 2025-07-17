package main

import (
	"fmt"
	"time"
)

// 方法1：贪心算法 - 抓住每一次上涨机会
// 时间复杂度：O(n)，空间复杂度：O(1)
func maxProfitGreedy(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}
	
	maxProfit := 0
	// 只要第二天比第一天价格高，就在第一天买入，第二天卖出
	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[i-1] {
			maxProfit += prices[i] - prices[i-1]
		}
	}
	
	return maxProfit
}

// 方法2：动态规划
// 时间复杂度：O(n)，空间复杂度：O(n)
func maxProfitDP(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}
	
	n := len(prices)
	// dp[i][0] 表示第i天不持有股票的最大利润
	// dp[i][1] 表示第i天持有股票的最大利润
	dp := make([][2]int, n)
	
	// 初始状态
	dp[0][0] = 0          // 第0天不持有股票，利润为0
	dp[0][1] = -prices[0] // 第0天买入股票，利润为-prices[0]
	
	for i := 1; i < n; i++ {
		// 第i天不持有股票：可能是前一天就不持有，或者前一天持有今天卖出
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i])
		// 第i天持有股票：可能是前一天就持有，或者前一天不持有今天买入
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]-prices[i])
	}
	
	// 最后一天不持有股票肯定比持有股票利润更大
	return dp[n-1][0]
}

// 方法3：动态规划（空间优化）
// 时间复杂度：O(n)，空间复杂度：O(1)
func maxProfitDPOptimized(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}
	
	// 只需要记录当前的状态，不需要整个数组
	hold := -prices[0] // 持有股票的最大利润
	sold := 0          // 不持有股票的最大利润
	
	for i := 1; i < len(prices); i++ {
		newSold := max(sold, hold+prices[i]) // 今天卖出
		newHold := max(hold, sold-prices[i]) // 今天买入
		
		sold = newSold
		hold = newHold
	}
	
	return sold
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	testCases := [][]int{
		{7, 1, 5, 3, 6, 4}, // 预期输出: 7
		{1, 2, 3, 4, 5},    // 预期输出: 4
		{7, 6, 4, 3, 1},    // 预期输出: 0
		{1, 2, 1, 2, 1},    // 预期输出: 2
		{5},                // 预期输出: 0
		{},                 // 预期输出: 0
	}
	
	fmt.Println("=== 买卖股票的最佳时机 II - 多种解法测试 ===")
	fmt.Println()
	
	methods := []struct {
		name string
		fn   func([]int) int
	}{
		{"贪心算法", maxProfitGreedy},
		{"动态规划", maxProfitDP},
		{"动态规划(优化)", maxProfitDPOptimized},
	}
	
	for i, testCase := range testCases {
		fmt.Printf("测试用例 %d: %v\n", i+1, testCase)
		
		for _, method := range methods {
			start := time.Now()
			result := method.fn(testCase)
			duration := time.Since(start)
			fmt.Printf("  %s: %d (用时: %v)\n", method.name, result, duration)
		}
		fmt.Println()
	}
	
	// 算法思路演示
	fmt.Println("=== 算法思路演示 ===")
	demonstrateGreedyApproach([]int{7, 1, 5, 3, 6, 4})
}

func demonstrateGreedyApproach(prices []int) {
	fmt.Printf("\n贪心算法思路演示:\n")
	fmt.Printf("价格数组: %v\n", prices)
	
	maxProfit := 0
	transactions := []string{}
	
	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[i-1] {
			profit := prices[i] - prices[i-1]
			maxProfit += profit
			transaction := fmt.Sprintf("第%d天买入(%d) -> 第%d天卖出(%d), 利润: %d", 
				i, prices[i-1], i+1, prices[i], profit)
			transactions = append(transactions, transaction)
		}
	}
	
	fmt.Println("交易记录:")
	for _, transaction := range transactions {
		fmt.Printf("  %s\n", transaction)
	}
	fmt.Printf("总利润: %d\n", maxProfit)
} 