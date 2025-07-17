package main

import (
	"fmt"
	"time"
)

// 方法1：动态规划 - 五种状态
// 时间复杂度：O(n)，空间复杂度：O(1)
func maxProfitDP5States(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}
	
	// 定义五种状态
	// 0: 没有操作
	// 1: 第一次买入
	// 2: 第一次卖出
	// 3: 第二次买入
	// 4: 第二次卖出
	buy1 := -prices[0]  // 第一次买入
	sell1 := 0          // 第一次卖出
	buy2 := -prices[0]  // 第二次买入
	sell2 := 0          // 第二次卖出
	
	for i := 1; i < len(prices); i++ {
		// 第二次卖出：要么保持原状，要么在第二次买入基础上卖出
		sell2 = max(sell2, buy2+prices[i])
		// 第二次买入：要么保持原状，要么在第一次卖出基础上买入
		buy2 = max(buy2, sell1-prices[i])
		// 第一次卖出：要么保持原状，要么在第一次买入基础上卖出
		sell1 = max(sell1, buy1+prices[i])
		// 第一次买入：要么保持原状，要么今天买入
		buy1 = max(buy1, -prices[i])
	}
	
	return sell2
}

// 方法2：动态规划 - 通用k笔交易解法（k=2）
// 时间复杂度：O(n)，空间复杂度：O(1)
func maxProfitDPGeneral(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}
	
	k := 2 // 最多2笔交易
	
	// buy[i] 表示第i笔交易买入后的最大利润
	// sell[i] 表示第i笔交易卖出后的最大利润
	buy := make([]int, k+1)
	sell := make([]int, k+1)
	
	// 初始化
	for i := 1; i <= k; i++ {
		buy[i] = -prices[0]
		sell[i] = 0
	}
	
	for i := 1; i < len(prices); i++ {
		for j := k; j >= 1; j-- {
			// 第j笔交易卖出：保持原状 vs 在第j笔买入基础上卖出
			sell[j] = max(sell[j], buy[j]+prices[i])
			// 第j笔交易买入：保持原状 vs 在第j-1笔卖出基础上买入
			buy[j] = max(buy[j], sell[j-1]-prices[i])
		}
	}
	
	return sell[k]
}

// 方法3：动态规划 - 完整状态表
// 时间复杂度：O(n)，空间复杂度：O(n)
func maxProfitDPTable(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}
	
	n := len(prices)
	// dp[i][j][k] 表示第i天，已完成j笔交易，持股状态为k的最大利润
	// k=0表示不持股，k=1表示持股
	dp := make([][][2]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([][2]int, 3) // 0,1,2笔交易
	}
	
	// 初始化第0天
	dp[0][0][0] = 0          // 第0天，0笔交易，不持股
	dp[0][0][1] = -prices[0] // 第0天，0笔交易，持股（买入）
	dp[0][1][0] = 0          // 第0天，1笔交易，不持股（不可能）
	dp[0][1][1] = -prices[0] // 第0天，1笔交易，持股（不可能）
	dp[0][2][0] = 0          // 第0天，2笔交易，不持股（不可能）
	dp[0][2][1] = -prices[0] // 第0天，2笔交易，持股（不可能）
	
	for i := 1; i < n; i++ {
		for j := 0; j <= 2; j++ {
			// 不持股：昨天就不持股 or 今天卖出
			if j > 0 {
				dp[i][j][0] = max(dp[i-1][j][0], dp[i-1][j-1][1]+prices[i])
			} else {
				dp[i][j][0] = dp[i-1][j][0]
			}
			
			// 持股：昨天就持股 or 今天买入
			dp[i][j][1] = max(dp[i-1][j][1], dp[i-1][j][0]-prices[i])
		}
	}
	
	return max(dp[n-1][0][0], max(dp[n-1][1][0], dp[n-1][2][0]))
}

// 方法4：分割数组 - 枚举分割点
// 时间复杂度：O(n^2)，空间复杂度：O(n)
func maxProfitSplit(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}
	
	n := len(prices)
	
	// 计算从左到右的最大利润（一笔交易）
	leftProfit := make([]int, n)
	minPrice := prices[0]
	for i := 1; i < n; i++ {
		minPrice = min(minPrice, prices[i])
		leftProfit[i] = max(leftProfit[i-1], prices[i]-minPrice)
	}
	
	// 计算从右到左的最大利润（一笔交易）
	rightProfit := make([]int, n)
	maxPrice := prices[n-1]
	for i := n - 2; i >= 0; i-- {
		maxPrice = max(maxPrice, prices[i])
		rightProfit[i] = max(rightProfit[i+1], maxPrice-prices[i])
	}
	
	// 枚举分割点，计算最大利润
	maxProfit := 0
	for i := 0; i < n; i++ {
		profit := leftProfit[i]
		if i + 1 < n {
			profit += rightProfit[i+1]
		}
		maxProfit = max(maxProfit, profit)
	}
	
	return maxProfit
}

// 方法5：状态机模拟
// 时间复杂度：O(n)，空间复杂度：O(1)
func maxProfitStateMachine(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}
	
	// 状态定义：[第几笔交易][买入/卖出状态]
	// hold1: 第一笔买入后的最大利润
	// sold1: 第一笔卖出后的最大利润
	// hold2: 第二笔买入后的最大利润
	// sold2: 第二笔卖出后的最大利润
	hold1 := -prices[0]
	sold1 := 0
	hold2 := -prices[0]
	sold2 := 0
	
	for i := 1; i < len(prices); i++ {
		// 更新顺序很重要：从后往前更新
		sold2 = max(sold2, hold2+prices[i]) // 第二次卖出
		hold2 = max(hold2, sold1-prices[i]) // 第二次买入
		sold1 = max(sold1, hold1+prices[i]) // 第一次卖出
		hold1 = max(hold1, -prices[i])      // 第一次买入
	}
	
	return sold2
}

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
	testCases := [][]int{
		{3, 3, 5, 0, 0, 3, 1, 4}, // 预期输出: 6
		{1, 2, 3, 4, 5},          // 预期输出: 4
		{7, 6, 4, 3, 1},          // 预期输出: 0
		{1},                      // 预期输出: 0
		{2, 1, 2, 0, 1},          // 预期输出: 2
		{1, 2, 4, 2, 5, 7, 2, 4, 9, 0}, // 预期输出: 13
	}
	
	fmt.Println("=== 买卖股票的最佳时机 III - 多种解法测试 ===")
	fmt.Println()
	
	methods := []struct {
		name string
		fn   func([]int) int
	}{
		{"五状态DP", maxProfitDP5States},
		{"通用k笔交易DP", maxProfitDPGeneral},
		{"完整状态表DP", maxProfitDPTable},
		{"分割数组", maxProfitSplit},
		{"状态机", maxProfitStateMachine},
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
	
	// 性能测试
	fmt.Println("=== 性能测试 ===")
	largeTestCase := make([]int, 1000)
	for i := 0; i < len(largeTestCase); i++ {
		largeTestCase[i] = (i*13 + 7) % 100 // 生成伪随机价格
	}
	
	fmt.Printf("大规模测试用例大小: %d\n", len(largeTestCase))
	for _, method := range methods {
		start := time.Now()
		result := method.fn(largeTestCase)
		duration := time.Since(start)
		fmt.Printf("%s: 结果=%d, 用时=%v\n", method.name, result, duration)
	}
	
	// 算法思路演示
	fmt.Println("\n=== 算法思路演示 ===")
	demonstrateStateTransition([]int{3, 3, 5, 0, 0, 3, 1, 4})
	demonstrateSplitApproach([]int{3, 3, 5, 0, 0, 3, 1, 4})
}

func demonstrateStateTransition(prices []int) {
	fmt.Printf("\n五状态动态规划演示:\n")
	fmt.Printf("价格数组: %v\n", prices)
	
	buy1 := -prices[0]
	sell1 := 0
	buy2 := -prices[0]
	sell2 := 0
	
	fmt.Printf("天数 | 价格 | 第1次买入 | 第1次卖出 | 第2次买入 | 第2次卖出\n")
	fmt.Printf("  0  |  %2d  |    %2d    |    %2d    |    %2d    |    %2d\n", 
		prices[0], buy1, sell1, buy2, sell2)
	
	for i := 1; i < len(prices); i++ {
		sell2 = max(sell2, buy2+prices[i])
		buy2 = max(buy2, sell1-prices[i])
		sell1 = max(sell1, buy1+prices[i])
		buy1 = max(buy1, -prices[i])
		
		fmt.Printf("  %d  |  %2d  |    %2d    |    %2d    |    %2d    |    %2d\n", 
			i, prices[i], buy1, sell1, buy2, sell2)
	}
	
	fmt.Printf("最大利润: %d\n", sell2)
}

func demonstrateSplitApproach(prices []int) {
	fmt.Printf("\n分割数组方法演示:\n")
	fmt.Printf("价格数组: %v\n", prices)
	
	n := len(prices)
	
	// 计算左侧最大利润
	leftProfit := make([]int, n)
	minPrice := prices[0]
	for i := 1; i < n; i++ {
		minPrice = min(minPrice, prices[i])
		leftProfit[i] = max(leftProfit[i-1], prices[i]-minPrice)
	}
	
	// 计算右侧最大利润
	rightProfit := make([]int, n)
	maxPrice := prices[n-1]
	for i := n - 2; i >= 0; i-- {
		maxPrice = max(maxPrice, prices[i])
		rightProfit[i] = max(rightProfit[i+1], maxPrice-prices[i])
	}
	
	fmt.Printf("位置 | 价格 | 左侧利润 | 右侧利润 | 总利润\n")
	maxTotalProfit := 0
	bestSplit := 0
	
	for i := 0; i < n; i++ {
		rightPart := 0
		if i + 1 < n {
			rightPart = rightProfit[i+1]
		}
		totalProfit := leftProfit[i] + rightPart
		
		if totalProfit > maxTotalProfit {
			maxTotalProfit = totalProfit
			bestSplit = i
		}
		
		fmt.Printf("  %d  |  %2d  |    %2d    |    %2d    |   %2d\n", 
			i, prices[i], leftProfit[i], rightPart, totalProfit)
	}
	
	fmt.Printf("最佳分割点: %d, 最大利润: %d\n", bestSplit, maxTotalProfit)
}
