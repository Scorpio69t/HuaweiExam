package main

import (
	"fmt"
)

// =========================== 方法一：回溯算法（最优解法） ===========================

func subsets(nums []int) [][]int {
	var result [][]int
	var path []int

	var backtrack func(index int)
	backtrack = func(index int) {
		// 记录当前状态（每个状态都是有效子集）
		temp := make([]int, len(path))
		copy(temp, path)
		result = append(result, temp)

		// 终止条件
		if index == len(nums) {
			return
		}

		// 选择当前元素
		for i := index; i < len(nums); i++ {
			path = append(path, nums[i])
			backtrack(i + 1)
			path = path[:len(path)-1] // 回溯
		}
	}

	backtrack(0)
	return result
}

// =========================== 方法二：迭代算法（非递归） ===========================

func subsets2(nums []int) [][]int {
	var result [][]int
	result = append(result, []int{}) // 空集

	for _, num := range nums {
		size := len(result)
		for i := 0; i < size; i++ {
			newSubset := make([]int, len(result[i]))
			copy(newSubset, result[i])
			newSubset = append(newSubset, num)
			result = append(result, newSubset)
		}
	}

	return result
}

// =========================== 方法三：位运算算法 ===========================

func subsets3(nums []int) [][]int {
	var result [][]int
	n := len(nums)

	// 生成所有可能的位掩码
	for mask := 0; mask < (1 << n); mask++ {
		var subset []int
		for i := 0; i < n; i++ {
			if mask&(1<<i) != 0 {
				subset = append(subset, nums[i])
			}
		}
		result = append(result, subset)
	}

	return result
}

// =========================== 方法四：递归枚举 ===========================

func subsets4(nums []int) [][]int {
	var result [][]int

	var generate func(index int, current []int)
	generate = func(index int, current []int) {
		if index == len(nums) {
			result = append(result, current)
			return
		}

		// 不选择当前元素
		generate(index+1, current)

		// 选择当前元素
		newCurrent := make([]int, len(current))
		copy(newCurrent, current)
		newCurrent = append(newCurrent, nums[index])
		generate(index+1, newCurrent)
	}

	generate(0, []int{})
	return result
}

// =========================== 测试代码 ===========================

func main() {
	fmt.Println("=== LeetCode 78: 子集 ===\n")

	testCases := []struct {
		nums   []int
		expect [][]int
	}{
		{
			[]int{1, 2, 3},
			[][]int{{}, {1}, {1, 2}, {1, 2, 3}, {1, 3}, {2}, {2, 3}, {3}},
		},
		{
			[]int{0},
			[][]int{{}, {0}},
		},
		{
			[]int{},
			[][]int{{}},
		},
		{
			[]int{1, 2},
			[][]int{{}, {1}, {1, 2}, {2}},
		},
		{
			[]int{1, 2, 3, 4},
			[][]int{{}, {1}, {1, 2}, {1, 2, 3}, {1, 2, 3, 4}, {1, 2, 4}, {1, 3}, {1, 3, 4}, {1, 4}, {2}, {2, 3}, {2, 3, 4}, {2, 4}, {3}, {3, 4}, {4}},
		},
	}

	fmt.Println("方法一：回溯算法（最优解法）")
	runTests(testCases, subsets)

	fmt.Println("\n方法二：迭代算法（非递归）")
	runTests(testCases, subsets2)

	fmt.Println("\n方法三：位运算算法")
	runTests(testCases, subsets3)

	fmt.Println("\n方法四：递归枚举")
	runTests(testCases, subsets4)
}

func runTests(testCases []struct {
	nums   []int
	expect [][]int
}, fn func([]int) [][]int) {
	passCount := 0
	for i, tc := range testCases {
		result := fn(tc.nums)
		status := "✅"
		if !compareSubsets(result, tc.expect) {
			status = "❌"
		} else {
			passCount++
		}
		fmt.Printf("  测试%d: %s\n", i+1, status)
		if status == "❌" {
			fmt.Printf("    输入: %v\n", tc.nums)
			fmt.Printf("    输出: %v\n", result)
			fmt.Printf("    期望: %v\n", tc.expect)
		}
	}
	fmt.Printf("  通过: %d/%d\n", passCount, len(testCases))
}

func compareSubsets(result, expect [][]int) bool {
	if len(result) != len(expect) {
		return false
	}

	// 创建结果映射
	resultMap := make(map[string]bool)
	for _, subset := range result {
		key := fmt.Sprintf("%v", subset)
		resultMap[key] = true
	}

	expectMap := make(map[string]bool)
	for _, subset := range expect {
		key := fmt.Sprintf("%v", subset)
		expectMap[key] = true
	}

	// 比较映射
	if len(resultMap) != len(expectMap) {
		return false
	}

	for key := range resultMap {
		if !expectMap[key] {
			return false
		}
	}

	return true
}
