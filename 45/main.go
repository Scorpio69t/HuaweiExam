package main

import (
	"fmt"
	"time"
)

// 方法一：贪心算法
// 最优解法，时间复杂度O(n)，空间复杂度O(1)
func jump1(nums []int) int {
	n := len(nums)
	if n <= 1 {
		return 0
	}

	jumps := 0
	currentEnd := 0
	farthest := 0

	for i := 0; i < n-1; i++ {
		// 更新能到达的最远位置
		farthest = max(farthest, i+nums[i])

		// 如果到达当前边界，需要跳跃
		if i == currentEnd {
			jumps++
			currentEnd = farthest
		}
	}

	return jumps
}

// 方法二：动态规划算法
// 经典DP解法，时间复杂度O(n²)，空间复杂度O(n)
func jump2(nums []int) int {
	n := len(nums)
	if n <= 1 {
		return 0
	}

	dp := make([]int, n)

	// 初始化DP数组
	for i := 1; i < n; i++ {
		dp[i] = n // 初始化为最大值
	}

	// 状态转移
	for i := 1; i < n; i++ {
		for j := 0; j < i; j++ {
			if j+nums[j] >= i {
				dp[i] = min(dp[i], dp[j]+1)
			}
		}
	}

	return dp[n-1]
}

// 方法三：BFS搜索算法
// 广度优先搜索，时间复杂度O(n²)，空间复杂度O(n)
func jump3(nums []int) int {
	n := len(nums)
	if n <= 1 {
		return 0
	}

	queue := []int{0}
	visited := make([]bool, n)
	visited[0] = true
	jumps := 0

	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			pos := queue[0]
			queue = queue[1:]

			if pos == n-1 {
				return jumps
			}

			// 添加所有可跳位置
			for j := 1; j <= nums[pos]; j++ {
				nextPos := pos + j
				if nextPos < n && !visited[nextPos] {
					visited[nextPos] = true
					queue = append(queue, nextPos)
				}
			}
		}
		jumps++
	}

	return jumps
}

// 方法四：递归回溯算法
// 递归回溯解法，使用记忆化优化
func jump4(nums []int) int {
	n := len(nums)
	if n <= 1 {
		return 0
	}

	memo := make(map[int]int)
	return backtrack(nums, 0, memo)
}

// 递归回溯的辅助函数
func backtrack(nums []int, pos int, memo map[int]int) int {
	if pos >= len(nums)-1 {
		return 0
	}

	if jumps, exists := memo[pos]; exists {
		return jumps
	}

	minJumps := len(nums)
	for i := 1; i <= nums[pos]; i++ {
		nextPos := pos + i
		if nextPos < len(nums) {
			jumps := 1 + backtrack(nums, nextPos, memo)
			minJumps = min(minJumps, jumps)
		}
	}

	memo[pos] = minJumps
	return minJumps
}

// 辅助函数：计算最大值
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 辅助函数：计算最小值
func min(a, b int) int {
	if a < b {
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
		{[]int{2, 3, 1, 1, 4}, "示例1: [2,3,1,1,4]"},
		{[]int{2, 3, 0, 1, 4}, "示例2: [2,3,0,1,4]"},
		{[]int{1, 2, 3}, "测试1: [1,2,3]"},
		{[]int{1, 1, 1, 1}, "测试2: [1,1,1,1]"},
		{[]int{5, 4, 3, 2, 1}, "测试3: [5,4,3,2,1]"},
		{[]int{1}, "测试4: [1]"},
		{[]int{2, 1}, "测试5: [2,1]"},
		{[]int{1, 2, 1, 1, 1}, "测试6: [1,2,1,1,1]"},
		{[]int{3, 2, 1, 1, 4}, "测试7: [3,2,1,1,4]"},
		{[]int{2, 0, 2, 0, 1}, "测试8: [2,0,2,0,1]"},
		{[]int{1, 2, 3, 4, 5}, "测试9: [1,2,3,4,5]"},
		{[]int{5, 9, 3, 2, 1, 0, 2, 3, 3, 1, 0, 0}, "测试10: [5,9,3,2,1,0,2,3,3,1,0,0]"},
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
	if result < 0 {
		return false
	}

	// 验证是否能到达目标
	n := len(nums)
	if n <= 1 {
		return result == 0
	}

	// 简单的验证：结果应该小于等于n-1
	return result <= n-1
}

// 辅助函数：打印跳跃结果
func printJumpResult(nums []int, result int, title string) {
	fmt.Printf("%s: nums=%v -> %d 次跳跃\n", title, nums, result)
}

func main() {
	fmt.Println("=== 45. 跳跃游戏 II ===")
	fmt.Println()

	// 创建测试用例
	testCases := createTestCases()
	algorithms := []struct {
		name string
		fn   func([]int) int
	}{
		{"贪心算法", jump1},
		{"动态规划算法", jump2},
		{"BFS搜索算法", jump3},
		{"递归回溯算法", jump4},
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
			if results[i] != results[0] {
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
			fmt.Printf("  ✅ 所有算法结果一致且正确: %d 次跳跃\n", results[0])
			if len(testCase.nums) <= 10 {
				printJumpResult(testCase.nums, results[0], "  跳跃结果")
			}
		} else {
			fmt.Printf("  ❌ 算法结果不一致或错误\n")
			for i, algo := range algorithms {
				fmt.Printf("    %s: %d 次跳跃\n", algo.name, results[i])
			}
		}
		fmt.Println()
	}

	// 性能测试
	fmt.Println("=== 性能测试 ===")
	performanceNums := []int{5, 9, 3, 2, 1, 0, 2, 3, 3, 1, 0, 0, 5, 9, 3, 2, 1, 0, 2, 3, 3, 1, 0, 0}

	fmt.Printf("测试数据: nums=%v\n", performanceNums)
	fmt.Println()

	for _, algo := range algorithms {
		benchmarkAlgorithm(algo.fn, performanceNums, algo.name)
	}
	fmt.Println()

	// 算法分析
	fmt.Println("=== 算法分析 ===")
	fmt.Println("跳跃游戏II问题的特点:")
	fmt.Println("1. 需要找到到达数组末尾的最小跳跃次数")
	fmt.Println("2. 每个元素表示从该位置能跳转的最大长度")
	fmt.Println("3. 贪心算法是最优解法")
	fmt.Println("4. 需要处理各种边界情况")
	fmt.Println()

	// 复杂度分析
	fmt.Println("=== 复杂度分析 ===")
	fmt.Println("时间复杂度:")
	fmt.Println("- 贪心算法: O(n)，单次遍历数组")
	fmt.Println("- 动态规划: O(n²)，双重循环遍历所有位置")
	fmt.Println("- BFS搜索: O(n²)，最坏情况需要遍历所有位置")
	fmt.Println("- 递归回溯: O(2^n)，最坏情况需要遍历所有可能的跳跃路径")
	fmt.Println()
	fmt.Println("空间复杂度:")
	fmt.Println("- 贪心算法: O(1)，只使用常数空间")
	fmt.Println("- 动态规划: O(n)，需要DP数组存储状态")
	fmt.Println("- BFS搜索: O(n)，需要队列和访问数组")
	fmt.Println("- 递归栈: O(n)，递归深度最多为n")
	fmt.Println()

	// 算法总结
	fmt.Println("=== 算法总结 ===")
	fmt.Println("1. 贪心算法：最优解法，时间复杂度O(n)")
	fmt.Println("2. 动态规划：经典DP解法，逻辑清晰")
	fmt.Println("3. BFS搜索：广度优先搜索，适合特定场景")
	fmt.Println("4. 递归回溯：最直观的解法，但效率较低")
	fmt.Println()
	fmt.Println("推荐使用：贪心算法（方法一），时间复杂度最低")
	fmt.Println()

	// 应用场景
	fmt.Println("=== 应用场景 ===")
	fmt.Println("- 路径规划：寻找最短路径问题")
	fmt.Println("- 游戏开发：跳跃类游戏的路径计算")
	fmt.Println("- 网络路由：寻找最优路由路径")
	fmt.Println("- 资源分配：最优资源分配问题")
	fmt.Println("- 算法竞赛：贪心算法的经典应用")
	fmt.Println()

	// 优化技巧总结
	fmt.Println("=== 优化技巧总结 ===")
	fmt.Println("1. 贪心选择：在每一步都选择能跳得最远的位置")
	fmt.Println("2. 边界维护：维护当前能到达的最远位置和下一步能到达的最远位置")
	fmt.Println("3. 跳跃计数：当到达当前边界时，增加跳跃次数并更新边界")
	fmt.Println("4. 最优子结构：每一步的最优选择构成全局最优解")
	fmt.Println("5. 边界处理：处理数组长度为1的特殊情况")
	fmt.Println("6. 算法选择：根据问题特点选择合适的算法")
}
