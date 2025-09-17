package main

import (
	"fmt"
	"sort"
	"time"
)

// 方法一：基础回溯算法
// 最直观的回溯实现，递归搜索所有可能的组合
func combinationSum2_1(candidates []int, target int) [][]int {
	sort.Ints(candidates) // 排序以便去重
	var result [][]int
	var path []int

	backtrack1(candidates, target, 0, path, &result)
	return result
}

// 基础回溯的递归辅助函数
func backtrack1(candidates []int, target, start int, path []int, result *[][]int) {
	if target == 0 {
		// 找到解，复制路径
		temp := make([]int, len(path))
		copy(temp, path)
		*result = append(*result, temp)
		return
	}

	if target < 0 {
		return // 剪枝：目标和为负数，无解
	}

	for i := start; i < len(candidates); i++ {
		// 去重：跳过重复数字
		if i > start && candidates[i] == candidates[i-1] {
			continue
		}

		path = append(path, candidates[i])
		backtrack1(candidates, target-candidates[i], i+1, path, result) // 注意：i+1，不能重复使用
		path = path[:len(path)-1]                                       // 回溯
	}
}

// 方法二：排序去重算法
// 排序候选数组，去重处理，避免重复组合
func combinationSum2_2(candidates []int, target int) [][]int {
	sort.Ints(candidates) // 排序
	var result [][]int
	var path []int

	backtrack2(candidates, target, 0, path, &result)
	return result
}

// 排序去重的递归辅助函数
func backtrack2(candidates []int, target, start int, path []int, result *[][]int) {
	if target == 0 {
		temp := make([]int, len(path))
		copy(temp, path)
		*result = append(*result, temp)
		return
	}

	for i := start; i < len(candidates); i++ {
		// 去重：跳过重复数字
		if i > start && candidates[i] == candidates[i-1] {
			continue
		}

		if candidates[i] > target {
			break // 剪枝：后续数字都更大，不可能有解
		}

		path = append(path, candidates[i])
		backtrack2(candidates, target-candidates[i], i+1, path, result)
		path = path[:len(path)-1] // 回溯
	}
}

// 方法三：剪枝优化算法
// 添加更多剪枝条件，进一步优化性能
func combinationSum2_3(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	var result [][]int
	var path []int

	backtrack3(candidates, target, 0, path, &result)
	return result
}

// 剪枝优化的递归辅助函数
func backtrack3(candidates []int, target, start int, path []int, result *[][]int) {
	if target == 0 {
		temp := make([]int, len(path))
		copy(temp, path)
		*result = append(*result, temp)
		return
	}

	for i := start; i < len(candidates); i++ {
		// 去重：跳过重复数字
		if i > start && candidates[i] == candidates[i-1] {
			continue
		}

		// 剪枝：当前数字大于剩余目标
		if candidates[i] > target {
			break
		}

		// 剪枝：剩余数字总和小于目标
		if sum(candidates[i:]) < target {
			break
		}

		path = append(path, candidates[i])
		backtrack3(candidates, target-candidates[i], i+1, path, result)
		path = path[:len(path)-1] // 回溯
	}
}

// 计算数组元素总和
func sum(nums []int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}

// 方法四：位运算算法
// 使用位运算表示状态，适合小规模数据
func combinationSum2_4(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	var result [][]int
	n := len(candidates)

	// 使用位掩码表示选择状态
	for mask := 0; mask < (1 << n); mask++ {
		var path []int
		sum := 0

		for i := 0; i < n; i++ {
			if (mask>>i)&1 == 1 {
				path = append(path, candidates[i])
				sum += candidates[i]
			}
		}

		if sum == target {
			// 检查是否重复
			if !isDuplicate(result, path) {
				result = append(result, path)
			}
		}
	}

	return result
}

// 检查路径是否重复
func isDuplicate(result [][]int, path []int) bool {
	for _, existing := range result {
		if len(existing) == len(path) {
			match := true
			for i := 0; i < len(path); i++ {
				if existing[i] != path[i] {
					match = false
					break
				}
			}
			if match {
				return true
			}
		}
	}
	return false
}

// 辅助函数：打印组合结果
func printCombinations(combinations [][]int, title string) {
	fmt.Printf("%s:\n", title)
	if len(combinations) == 0 {
		fmt.Println("  无解")
		return
	}

	for i, combination := range combinations {
		fmt.Printf("  组合%d: %v\n", i+1, combination)
	}
	fmt.Println()
}

// 辅助函数：验证组合是否正确
func validateCombination(candidates []int, combination []int, target int) bool {
	sum := 0
	for _, num := range combination {
		sum += num
		// 检查数字是否在候选数组中
		found := false
		for _, candidate := range candidates {
			if num == candidate {
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}
	return sum == target
}

// 辅助函数：创建测试用例
func createTestCases() []struct {
	candidates []int
	target     int
	name       string
} {
	return []struct {
		candidates []int
		target     int
		name       string
	}{
		{[]int{10, 1, 2, 7, 6, 1, 5}, 8, "示例1: [10,1,2,7,6,1,5], target=8"},
		{[]int{2, 5, 2, 1, 2}, 5, "示例2: [2,5,2,1,2], target=5"},
		{[]int{1, 1, 2, 5, 6, 7, 10}, 8, "测试1: [1,1,2,5,6,7,10], target=8"},
		{[]int{1, 2, 2, 2, 5}, 5, "测试2: [1,2,2,2,5], target=5"},
		{[]int{1, 2, 3}, 4, "测试3: [1,2,3], target=4"},
		{[]int{1, 1, 1}, 3, "测试4: [1,1,1], target=3"},
		{[]int{2, 3, 5}, 8, "测试5: [2,3,5], target=8"},
		{[]int{1, 2}, 4, "测试6: [1,2], target=4"},
	}
}

// 性能测试函数
func benchmarkAlgorithm(algorithm func([]int, int) [][]int, candidates []int, target int, name string) {
	iterations := 100
	start := time.Now()

	for i := 0; i < iterations; i++ {
		algorithm(candidates, target)
	}

	duration := time.Since(start)
	avgTime := duration.Nanoseconds() / int64(iterations)

	fmt.Printf("%s: 平均执行时间 %d 纳秒\n", name, avgTime)
}

// 辅助函数：比较两个结果是否相同
func compareResults(result1, result2 [][]int) bool {
	if len(result1) != len(result2) {
		return false
	}

	// 创建结果1的映射
	map1 := make(map[string]bool)
	for _, combination := range result1 {
		sort.Ints(combination)
		key := fmt.Sprintf("%v", combination)
		map1[key] = true
	}

	// 检查结果2中的每个组合是否在结果1中
	for _, combination := range result2 {
		sort.Ints(combination)
		key := fmt.Sprintf("%v", combination)
		if !map1[key] {
			return false
		}
	}

	return true
}

func main() {
	fmt.Println("=== 40. 组合总和 II ===")
	fmt.Println()

	// 创建测试用例
	testCases := createTestCases()
	algorithms := []struct {
		name string
		fn   func([]int, int) [][]int
	}{
		{"基础回溯算法", combinationSum2_1},
		{"排序去重算法", combinationSum2_2},
		{"剪枝优化算法", combinationSum2_3},
		{"位运算算法", combinationSum2_4},
	}

	// 运行测试
	fmt.Println("=== 算法正确性测试 ===")
	for _, testCase := range testCases {
		fmt.Printf("测试: %s\n", testCase.name)
		fmt.Printf("候选数组: %v, 目标: %d\n", testCase.candidates, testCase.target)

		results := make([][][]int, len(algorithms))
		for i, algo := range algorithms {
			results[i] = algo.fn(testCase.candidates, testCase.target)
		}

		// 验证所有算法结果一致
		allEqual := true
		for i := 1; i < len(results); i++ {
			if !compareResults(results[0], results[i]) {
				allEqual = false
				break
			}
		}

		if allEqual {
			fmt.Printf("  ✅ 所有算法结果一致，共找到 %d 个组合\n", len(results[0]))
			if len(results[0]) > 0 && len(results[0]) <= 5 {
				printCombinations(results[0], "  组合详情")
			}
		} else {
			fmt.Printf("  ❌ 算法结果不一致\n")
			for i, algo := range algorithms {
				fmt.Printf("    %s: %d 个组合\n", algo.name, len(results[i]))
			}
		}
		fmt.Println()
	}

	// 性能测试
	fmt.Println("=== 性能测试 ===")
	performanceCandidates := []int{1, 1, 2, 2, 3, 3, 4, 4, 5, 5}
	performanceTarget := 10

	fmt.Printf("测试数据: candidates=%v, target=%d\n", performanceCandidates, performanceTarget)
	fmt.Println()

	for _, algo := range algorithms {
		benchmarkAlgorithm(algo.fn, performanceCandidates, performanceTarget, algo.name)
	}
	fmt.Println()

	// 算法分析
	fmt.Println("=== 算法分析 ===")
	fmt.Println("组合总和II问题的特点:")
	fmt.Println("1. 每个数字只能使用一次")
	fmt.Println("2. 解集不能包含重复的组合")
	fmt.Println("3. 需要去重处理")
	fmt.Println("4. 可以通过排序和剪枝优化")
	fmt.Println()

	// 复杂度分析
	fmt.Println("=== 复杂度分析 ===")
	fmt.Println("时间复杂度:")
	fmt.Println("- 基础回溯: O(2^n)，最坏情况需要遍历所有可能的组合")
	fmt.Println("- 排序去重: O(n log n + 2^n)，排序开销+搜索开销")
	fmt.Println("- 剪枝优化: O(2^n)，但常数因子更小")
	fmt.Println("- 位运算: O(2^n)，位运算操作常数时间")
	fmt.Println()
	fmt.Println("空间复杂度:")
	fmt.Println("- 递归栈: O(n)，递归深度最多为n")
	fmt.Println("- 路径存储: O(n)，存储当前路径")
	fmt.Println("- 结果存储: O(2^n)，存储所有可能的组合")
	fmt.Println()

	// 算法总结
	fmt.Println("=== 算法总结 ===")
	fmt.Println("1. 基础回溯算法：最直观易懂，适合理解算法逻辑")
	fmt.Println("2. 排序去重算法：排序后去重，避免重复组合")
	fmt.Println("3. 剪枝优化算法：添加剪枝条件，性能最佳")
	fmt.Println("4. 位运算算法：适合小规模数据，但复杂度较高")
	fmt.Println()
	fmt.Println("推荐使用：剪枝优化算法（方法三），在保证性能的同时剪枝效果最好")
	fmt.Println()

	// 应用场景
	fmt.Println("=== 应用场景 ===")
	fmt.Println("- 组合优化：寻找满足条件的所有不重复组合")
	fmt.Println("- 背包问题：物品不能重复选择的背包问题")
	fmt.Println("- 数论问题：数字分解和组合问题")
	fmt.Println("- 算法竞赛：回溯算法的经典应用")
	fmt.Println("- 游戏开发：道具组合和技能搭配")
	fmt.Println()

	// 优化技巧总结
	fmt.Println("=== 优化技巧总结 ===")
	fmt.Println("1. 排序去重：排序后可以方便地去重")
	fmt.Println("2. 状态管理：合理管理递归状态和路径")
	fmt.Println("3. 去重处理：避免生成重复的组合")
	fmt.Println("4. 早期终止：发现无解立即返回")
	fmt.Println("5. 空间优化：及时释放不需要的内存")
	fmt.Println("6. 算法选择：根据数据规模选择合适的算法")
}
