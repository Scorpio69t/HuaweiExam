package main

import (
	"fmt"
	"math"
	"time"
)

// 方法一：暴力枚举算法
// 最直观的解法，枚举所有可能的子数组并计算和
func maxSubArray1(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	maxSum := nums[0]

	// 枚举所有可能的子数组
	for i := 0; i < len(nums); i++ {
		for j := i; j < len(nums); j++ {
			sum := 0
			for k := i; k <= j; k++ {
				sum += nums[k]
			}
			if sum > maxSum {
				maxSum = sum
			}
		}
	}

	return maxSum
}

// 方法二：动态规划算法
// 使用DP状态表示以当前位置结尾的最大子数组和
func maxSubArray2(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	maxSum := nums[0]
	currentSum := nums[0]

	for i := 1; i < len(nums); i++ {
		if currentSum > 0 {
			currentSum += nums[i]
		} else {
			currentSum = nums[i]
		}

		if currentSum > maxSum {
			maxSum = currentSum
		}
	}

	return maxSum
}

// 方法三：分治算法
// 将问题分解为左半部分、右半部分和跨越中点的子数组
func maxSubArray3(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	return divideConquer(nums, 0, len(nums)-1)
}

// 分治算法的递归函数
func divideConquer(nums []int, left, right int) int {
	if left == right {
		return nums[left]
	}

	mid := (left + right) / 2

	// 左半部分的最大子数组和
	leftMax := divideConquer(nums, left, mid)

	// 右半部分的最大子数组和
	rightMax := divideConquer(nums, mid+1, right)

	// 跨越中点的最大子数组和
	crossMax := maxCrossingSum(nums, left, mid, right)

	return max(leftMax, max(rightMax, crossMax))
}

// 计算跨越中点的最大子数组和
func maxCrossingSum(nums []int, left, mid, right int) int {
	// 从中点向左扩展
	leftSum := math.MinInt32
	sum := 0
	for i := mid; i >= left; i-- {
		sum += nums[i]
		if sum > leftSum {
			leftSum = sum
		}
	}

	// 从中点向右扩展
	rightSum := math.MinInt32
	sum = 0
	for i := mid + 1; i <= right; i++ {
		sum += nums[i]
		if sum > rightSum {
			rightSum = sum
		}
	}

	return leftSum + rightSum
}

// 方法四：贪心算法
// 使用贪心策略：如果当前子数组和为负数就重新开始
func maxSubArray4(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	maxSum := nums[0]
	currentSum := nums[0]

	for i := 1; i < len(nums); i++ {
		// 如果当前子数组和为负数，重新开始
		if currentSum < 0 {
			currentSum = nums[i]
		} else {
			currentSum += nums[i]
		}

		// 更新全局最大值
		if currentSum > maxSum {
			maxSum = currentSum
		}
	}

	return maxSum
}

// 辅助函数：求两个数的最大值
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 辅助函数：创建测试用例
func createTestCases() []struct {
	nums []int
	name string
} {
	return []struct {
		nums []int
		name string
	}{
		{[]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}, "示例1: 混合正负数"},
		{[]int{1}, "示例2: 单个元素"},
		{[]int{5, 4, -1, 7, 8}, "示例3: 全正数"},
		{[]int{-1, -2, -3}, "测试1: 全负数"},
		{[]int{-1, 2, 3, -4, 5}, "测试2: 交替正负"},
		{[]int{1, 2, 3, 4, 5}, "测试3: 递增正数"},
		{[]int{-5, -4, -3, -2, -1}, "测试4: 递减负数"},
		{[]int{0, 0, 0, 0, 0}, "测试5: 全零"},
		{[]int{1, -1, 1, -1, 1}, "测试6: 交替1和-1"},
		{[]int{-2, -1, -3, -4, -1, -2, -1, -5, -4}, "测试7: 全负数长数组"},
	}
}

// 性能测试函数
func benchmarkAlgorithm(algorithm func([]int) int, nums []int, name string) {
	iterations := 1000
	start := time.Now()

	for i := 0; i < iterations; i++ {
		algorithm(nums)
	}

	duration := time.Since(start)
	avgTime := duration.Nanoseconds() / int64(iterations)

	fmt.Printf("%s: 平均执行时间 %d 纳秒\n", name, avgTime)
}

// 辅助函数：验证结果是否正确
func validateResult(nums []int, result int) bool {
	// 验证结果是否合理
	if len(nums) == 0 {
		return result == 0
	}

	// 结果应该大于等于数组中的最大值
	maxVal := nums[0]
	for _, num := range nums {
		if num > maxVal {
			maxVal = num
		}
	}

	return result >= maxVal
}

// 辅助函数：比较两个结果是否相同
func compareResults(result1, result2 int) bool {
	return result1 == result2
}

// 辅助函数：打印子数组和结果
func printSubArrayResult(nums []int, result int, title string) {
	fmt.Printf("%s: 数组%v -> 最大子数组和 = %d\n", title, nums, result)
}

func main() {
	fmt.Println("=== 53. 最大子数组和 ===")
	fmt.Println()

	// 创建测试用例
	testCases := createTestCases()
	algorithms := []struct {
		name string
		fn   func([]int) int
	}{
		{"暴力枚举算法", maxSubArray1},
		{"动态规划算法", maxSubArray2},
		{"分治算法", maxSubArray3},
		{"贪心算法", maxSubArray4},
	}

	// 运行测试
	fmt.Println("=== 算法正确性测试 ===")
	for _, testCase := range testCases {
		fmt.Printf("测试: %s\n", testCase.name)

		results := make([]int, len(algorithms))
		for i, algo := range algorithms {
			results[i] = algo.fn(testCase.nums)
		}

		// 验证所有算法结果一致
		allEqual := true
		for i := 1; i < len(results); i++ {
			if !compareResults(results[i], results[0]) {
				allEqual = false
				break
			}
		}

		// 验证结果是否正确
		allValid := true
		for _, result := range results {
			if !validateResult(testCase.nums, result) {
				allValid = false
				break
			}
		}

		if allEqual && allValid {
			fmt.Printf("  ✅ 所有算法结果一致且正确: %d\n", results[0])
			if len(testCase.nums) <= 10 {
				printSubArrayResult(testCase.nums, results[0], "  最大子数组和")
			}
		} else {
			fmt.Printf("  ❌ 算法结果不一致或错误\n")
			for i, algo := range algorithms {
				fmt.Printf("    %s: %d\n", algo.name, results[i])
			}
		}
		fmt.Println()
	}

	// 性能测试
	fmt.Println("=== 性能测试 ===")
	performanceNums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4, 6, -3, 2, 1, -4, 5, -2, 3, 1, -1, 2}

	fmt.Printf("测试数据: 长度为%d的数组\n", len(performanceNums))
	fmt.Println()

	for _, algo := range algorithms {
		benchmarkAlgorithm(algo.fn, performanceNums, algo.name)
	}
	fmt.Println()

	// 算法分析
	fmt.Println("=== 算法分析 ===")
	fmt.Println("最大子数组和问题的特点:")
	fmt.Println("1. 需要找到连续子数组的最大和")
	fmt.Println("2. 子数组至少包含一个元素")
	fmt.Println("3. 可以使用动态规划、分治、贪心等多种方法")
	fmt.Println("4. 动态规划和贪心算法是最优解法")
	fmt.Println()

	// 复杂度分析
	fmt.Println("=== 复杂度分析 ===")
	fmt.Println("时间复杂度:")
	fmt.Println("- 暴力枚举: O(n³)，需要三重循环")
	fmt.Println("- 动态规划: O(n)，只需要一次遍历")
	fmt.Println("- 分治算法: O(n log n)，递归深度为log n")
	fmt.Println("- 贪心算法: O(n)，只需要一次遍历")
	fmt.Println()
	fmt.Println("空间复杂度:")
	fmt.Println("- 暴力枚举: O(1)，只使用常数空间")
	fmt.Println("- 动态规划: O(1)，空间优化后只使用常数空间")
	fmt.Println("- 分治算法: O(log n)，递归栈的深度")
	fmt.Println("- 贪心算法: O(1)，只使用常数空间")
	fmt.Println()

	// 算法总结
	fmt.Println("=== 算法总结 ===")
	fmt.Println("1. 暴力枚举算法：最直观的解法，但效率低")
	fmt.Println("2. 动态规划算法：最优解法，效率最高")
	fmt.Println("3. 分治算法：分治思想，递归实现")
	fmt.Println("4. 贪心算法：贪心策略，与DP本质相同")
	fmt.Println()
	fmt.Println("推荐使用：动态规划算法（方法二），效率最高")
	fmt.Println()

	// 应用场景
	fmt.Println("=== 应用场景 ===")
	fmt.Println("- 算法竞赛：动态规划的经典应用")
	fmt.Println("- 数据分析：寻找最大收益区间")
	fmt.Println("- 金融分析：股票价格分析")
	fmt.Println("- 信号处理：寻找最大信号强度")
	fmt.Println("- 机器学习：特征选择")
	fmt.Println()

	// 优化技巧总结
	fmt.Println("=== 优化技巧总结 ===")
	fmt.Println("1. 动态规划：掌握状态定义和状态转移")
	fmt.Println("2. 贪心策略：理解贪心选择的性质")
	fmt.Println("3. 分治算法：学会将问题分解为子问题")
	fmt.Println("4. 边界处理：注意各种边界情况")
	fmt.Println("5. 算法选择：根据问题特点选择合适的算法")
	fmt.Println("6. 优化策略：学会时间和空间优化技巧")
}
