package main

import (
	"fmt"
	"sort"
	"time"
)

// 方法一：递归回溯算法
// 最直观的递归解法，使用路径和已使用标记
func permute1(nums []int) [][]int {
	var result [][]int
	var path []int
	used := make([]bool, len(nums))

	backtrack(nums, &path, used, &result)
	return result
}

// 递归回溯的辅助函数
func backtrack(nums []int, path *[]int, used []bool, result *[][]int) {
	// 终止条件
	if len(*path) == len(nums) {
		// 复制当前路径
		temp := make([]int, len(*path))
		copy(temp, *path)
		*result = append(*result, temp)
		return
	}

	// 遍历所有元素
	for i := 0; i < len(nums); i++ {
		if !used[i] {
			// 选择
			used[i] = true
			*path = append(*path, nums[i])

			// 递归
			backtrack(nums, path, used, result)

			// 撤销选择
			*path = (*path)[:len(*path)-1]
			used[i] = false
		}
	}
}

// 方法二：交换回溯算法
// 原地交换，空间效率最高
func permute2(nums []int) [][]int {
	var result [][]int
	backtrackSwap(nums, 0, &result)
	return result
}

// 交换回溯的辅助函数
func backtrackSwap(nums []int, start int, result *[][]int) {
	// 终止条件
	if start == len(nums) {
		// 复制当前排列
		temp := make([]int, len(nums))
		copy(temp, nums)
		*result = append(*result, temp)
		return
	}

	// 遍历剩余元素
	for i := start; i < len(nums); i++ {
		// 交换
		nums[start], nums[i] = nums[i], nums[start]

		// 递归
		backtrackSwap(nums, start+1, result)

		// 交换回来
		nums[start], nums[i] = nums[i], nums[start]
	}
}

// 方法三：迭代回溯算法
// 使用栈模拟递归，避免栈溢出
func permute3(nums []int) [][]int {
	var result [][]int
	stack := []struct {
		path []int
		used []bool
	}{{[]int{}, make([]bool, len(nums))}}

	for len(stack) > 0 {
		current := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		// 检查终止条件
		if len(current.path) == len(nums) {
			result = append(result, current.path)
			continue
		}

		// 遍历所有元素
		for i := 0; i < len(nums); i++ {
			if !current.used[i] {
				// 创建新状态
				newPath := make([]int, len(current.path))
				copy(newPath, current.path)
				newPath = append(newPath, nums[i])

				newUsed := make([]bool, len(current.used))
				copy(newUsed, current.used)
				newUsed[i] = true

				stack = append(stack, struct {
					path []int
					used []bool
				}{newPath, newUsed})
			}
		}
	}

	return result
}

// 方法四：字典序算法
// 按字典序生成排列，顺序固定
func permute4(nums []int) [][]int {
	var result [][]int

	// 排序确保字典序
	sort.Ints(nums)

	for {
		// 添加当前排列
		temp := make([]int, len(nums))
		copy(temp, nums)
		result = append(result, temp)

		// 找到下一个排列
		if !nextPermutation(nums) {
			break
		}
	}

	return result
}

// 找到下一个排列
func nextPermutation(nums []int) bool {
	n := len(nums)
	i := n - 2

	// 找到第一个递减的位置
	for i >= 0 && nums[i] >= nums[i+1] {
		i--
	}

	if i < 0 {
		return false
	}

	// 找到第一个大于nums[i]的位置
	j := n - 1
	for nums[j] <= nums[i] {
		j--
	}

	// 交换
	nums[i], nums[j] = nums[j], nums[i]

	// 反转后面的部分
	reverse(nums[i+1:])

	return true
}

// 反转数组
func reverse(nums []int) {
	for i, j := 0, len(nums)-1; i < j; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}
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
		{[]int{1, 2, 3}, "示例1: [1,2,3]"},
		{[]int{0, 1}, "示例2: [0,1]"},
		{[]int{1}, "示例3: [1]"},
		{[]int{1, 2}, "测试1: [1,2]"},
		{[]int{1, 2, 3, 4}, "测试2: [1,2,3,4]"},
		{[]int{5, 4, 6}, "测试3: [5,4,6]"},
		{[]int{1, 2, 3, 4, 5}, "测试4: [1,2,3,4,5]"},
		{[]int{0, 1, 2}, "测试5: [0,1,2]"},
		{[]int{1, 2, 3, 4, 5, 6}, "测试6: [1,2,3,4,5,6]"},
		{[]int{3, 1, 2}, "测试7: [3,1,2]"},
	}
}

// 性能测试函数
func benchmarkAlgorithm(algorithm func([]int) [][]int, nums []int, name string) {
	iterations := 100
	start := time.Now()

	for i := 0; i < iterations; i++ {
		algorithm(nums)
	}

	duration := time.Since(start)
	avgTime := duration.Nanoseconds() / int64(iterations)

	fmt.Printf("%s: 平均执行时间 %d 纳秒\n", name, avgTime)
}

// 辅助函数：验证结果是否正确
func validateResult(nums []int, result [][]int) bool {
	// 验证结果数量
	expectedCount := factorial(len(nums))
	if len(result) != expectedCount {
		return false
	}

	// 验证每个排列是否有效
	for _, perm := range result {
		if !isValidPermutation(nums, perm) {
			return false
		}
	}

	// 验证是否有重复
	if hasDuplicates(result) {
		return false
	}

	return true
}

// 计算阶乘
func factorial(n int) int {
	if n <= 1 {
		return 1
	}
	return n * factorial(n-1)
}

// 验证排列是否有效
func isValidPermutation(original, perm []int) bool {
	if len(original) != len(perm) {
		return false
	}

	// 检查是否包含所有元素
	originalMap := make(map[int]int)
	for _, num := range original {
		originalMap[num]++
	}

	permMap := make(map[int]int)
	for _, num := range perm {
		permMap[num]++
	}

	for num, count := range originalMap {
		if permMap[num] != count {
			return false
		}
	}

	return true
}

// 检查是否有重复
func hasDuplicates(result [][]int) bool {
	seen := make(map[string]bool)
	for _, perm := range result {
		key := fmt.Sprintf("%v", perm)
		if seen[key] {
			return true
		}
		seen[key] = true
	}
	return false
}

// 辅助函数：打印排列结果
func printPermutationResult(nums []int, result [][]int, title string) {
	fmt.Printf("%s: nums=%v -> %d 个排列\n", title, nums, len(result))
	if len(result) <= 6 {
		fmt.Printf("  排列结果: %v\n", result)
	}
}

func main() {
	fmt.Println("=== 46. 全排列 ===")
	fmt.Println()

	// 创建测试用例
	testCases := createTestCases()
	algorithms := []struct {
		name string
		fn   func([]int) [][]int
	}{
		{"递归回溯算法", permute1},
		{"交换回溯算法", permute2},
		{"迭代回溯算法", permute3},
		{"字典序算法", permute4},
	}

	// 运行测试
	fmt.Println("=== 算法正确性测试 ===")
	for _, testCase := range testCases {
		fmt.Printf("测试: %s\n", testCase.name)

		results := make([][][]int, len(algorithms))
		for i, algo := range algorithms {
			// 复制原始数组，因为有些算法会修改原数组
			tempNums := make([]int, len(testCase.nums))
			copy(tempNums, testCase.nums)
			results[i] = algo.fn(tempNums)
		}

		// 验证所有算法结果一致
		allEqual := true
		for i := 1; i < len(results); i++ {
			if len(results[i]) != len(results[0]) {
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
			fmt.Printf("  ✅ 所有算法结果一致且正确: %d 个排列\n", len(results[0]))
			if len(testCase.nums) <= 4 {
				printPermutationResult(testCase.nums, results[0], "  排列结果")
			}
		} else {
			fmt.Printf("  ❌ 算法结果不一致或错误\n")
			for i, algo := range algorithms {
				fmt.Printf("    %s: %d 个排列\n", algo.name, len(results[i]))
			}
		}
		fmt.Println()
	}

	// 性能测试
	fmt.Println("=== 性能测试 ===")
	performanceNums := []int{1, 2, 3, 4, 5}

	fmt.Printf("测试数据: nums=%v\n", performanceNums)
	fmt.Println()

	for _, algo := range algorithms {
		benchmarkAlgorithm(algo.fn, performanceNums, algo.name)
	}
	fmt.Println()

	// 算法分析
	fmt.Println("=== 算法分析 ===")
	fmt.Println("全排列问题的特点:")
	fmt.Println("1. 需要生成数组的所有可能排列")
	fmt.Println("2. 每个元素只能使用一次")
	fmt.Println("3. 回溯算法是最优解法")
	fmt.Println("4. 需要处理各种边界情况")
	fmt.Println()

	// 复杂度分析
	fmt.Println("=== 复杂度分析 ===")
	fmt.Println("时间复杂度:")
	fmt.Println("- 递归回溯: O(n!×n)，需要生成n!个排列，每个排列需要O(n)时间")
	fmt.Println("- 交换回溯: O(n!×n)，原地交换，但时间复杂度不变")
	fmt.Println("- 迭代回溯: O(n!×n)，使用栈模拟递归，时间复杂度相同")
	fmt.Println("- 字典序: O(n!×n)，按字典序生成，时间复杂度相同")
	fmt.Println()
	fmt.Println("空间复杂度:")
	fmt.Println("- 递归栈: O(n)，递归深度最多为n")
	fmt.Println("- 交换回溯: O(1)，只使用常数空间")
	fmt.Println("- 迭代栈: O(n)，栈的最大深度为n")
	fmt.Println("- 字典序: O(1)，只使用常数空间")
	fmt.Println()

	// 算法总结
	fmt.Println("=== 算法总结 ===")
	fmt.Println("1. 递归回溯算法：最直观的解法，逻辑清晰")
	fmt.Println("2. 交换回溯算法：空间效率最高，原地交换")
	fmt.Println("3. 迭代回溯算法：避免栈溢出，使用栈模拟递归")
	fmt.Println("4. 字典序算法：顺序固定，按字典序生成")
	fmt.Println()
	fmt.Println("推荐使用：交换回溯算法（方法二），空间效率最高")
	fmt.Println()

	// 应用场景
	fmt.Println("=== 应用场景 ===")
	fmt.Println("- 排列组合：生成所有可能的排列")
	fmt.Println("- 密码破解：尝试所有可能的密码组合")
	fmt.Println("- 游戏开发：生成所有可能的游戏状态")
	fmt.Println("- 数据分析：分析所有可能的数据组合")
	fmt.Println("- 算法竞赛：回溯算法的经典应用")
	fmt.Println()

	// 优化技巧总结
	fmt.Println("=== 优化技巧总结 ===")
	fmt.Println("1. 递归回溯：使用递归生成所有可能的排列")
	fmt.Println("2. 状态管理：维护当前路径和已使用的元素")
	fmt.Println("3. 选择与撤销：选择元素后递归，递归结束后撤销选择")
	fmt.Println("4. 终止条件：当路径长度等于数组长度时，找到一个完整排列")
	fmt.Println("5. 剪枝优化：避免重复选择同一元素")
	fmt.Println("6. 算法选择：根据问题特点选择合适的算法")
}
