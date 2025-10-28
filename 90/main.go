package main

import (
	"fmt"
	"sort"
)

// =========================== 方法一：回溯+去重（最优解法） ===========================

func subsetsWithDup(nums []int) [][]int {
	// 排序，使相同元素相邻
	sort.Ints(nums)

	var result [][]int
	var path []int

	var backtrack func(int)
	backtrack = func(start int) {
		// 添加当前路径到结果
		temp := make([]int, len(path))
		copy(temp, path)
		result = append(result, temp)

		// 遍历剩余元素
		for i := start; i < len(nums); i++ {
			// 去重：跳过重复元素
			if i > start && nums[i] == nums[i-1] {
				continue
			}

			// 选择当前元素
			path = append(path, nums[i])
			backtrack(i + 1)
			path = path[:len(path)-1] // 撤销选择
		}
	}

	backtrack(0)
	return result
}

// =========================== 方法二：位运算+去重 ===========================

func subsetsWithDup2(nums []int) [][]int {
	sort.Ints(nums)
	var result [][]int
	seen := make(map[string]bool)

	n := len(nums)
	for mask := 0; mask < (1 << n); mask++ {
		var subset []int
		for i := 0; i < n; i++ {
			if (mask>>i)&1 == 1 {
				subset = append(subset, nums[i])
			}
		}

		// 生成子集的字符串表示
		key := fmt.Sprintf("%v", subset)
		if !seen[key] {
			seen[key] = true
			result = append(result, subset)
		}
	}

	return result
}

// =========================== 方法三：递归+去重 ===========================

func subsetsWithDup3(nums []int) [][]int {
	sort.Ints(nums)
	var result [][]int

	var dfs func(int, []int)
	dfs = func(start int, path []int) {
		// 添加当前路径
		temp := make([]int, len(path))
		copy(temp, path)
		result = append(result, temp)

		for i := start; i < len(nums); i++ {
			if i > start && nums[i] == nums[i-1] {
				continue
			}

			path = append(path, nums[i])
			dfs(i+1, path)
			path = path[:len(path)-1]
		}
	}

	dfs(0, []int{})
	return result
}

// =========================== 方法四：迭代+去重（优化版） ===========================

func subsetsWithDup4(nums []int) [][]int {
	sort.Ints(nums)
	result := [][]int{{}}
	prevSize := 0

	for i := 0; i < len(nums); i++ {
		size := len(result)
		start := 0

		// 如果当前元素与前一个元素相同，只从上一轮新增的子集开始
		if i > 0 && nums[i] == nums[i-1] {
			start = prevSize
		}

		prevSize = size

		for j := start; j < size; j++ {
			newSubset := make([]int, len(result[j]))
			copy(newSubset, result[j])
			newSubset = append(newSubset, nums[i])
			result = append(result, newSubset)
		}
	}

	return result
}

// =========================== 测试代码 ===========================

func main() {
	fmt.Println("=== LeetCode 90: 子集 II ===\n")

	testCases := []struct {
		name     string
		nums     []int
		expected [][]int
	}{
		{
			name:     "Test1: Basic case",
			nums:     []int{1, 2, 2},
			expected: [][]int{{}, {1}, {1, 2}, {1, 2, 2}, {2}, {2, 2}},
		},
		{
			name:     "Test2: Single element",
			nums:     []int{0},
			expected: [][]int{{}, {0}},
		},
		{
			name:     "Test3: All same elements",
			nums:     []int{1, 1, 1},
			expected: [][]int{{}, {1}, {1, 1}, {1, 1, 1}},
		},
		{
			name:     "Test4: No duplicates",
			nums:     []int{1, 2, 3},
			expected: [][]int{{}, {1}, {1, 2}, {1, 2, 3}, {1, 3}, {2}, {2, 3}, {3}},
		},
		{
			name:     "Test5: Empty array",
			nums:     []int{},
			expected: [][]int{{}},
		},
		{
			name:     "Test6: Two duplicates",
			nums:     []int{4, 4, 4, 1, 4},
			expected: [][]int{{}, {1}, {1, 4}, {1, 4, 4}, {1, 4, 4, 4}, {1, 4, 4, 4, 4}, {4}, {4, 4}, {4, 4, 4}, {4, 4, 4, 4}},
		},
	}

	methods := map[string]func([]int) [][]int{
		"回溯+去重（最优解法）": subsetsWithDup,
		"位运算+去重":      subsetsWithDup2,
		"递归+去重":       subsetsWithDup3,
		"迭代+去重":       subsetsWithDup4,
	}

	for name, method := range methods {
		fmt.Printf("方法%s：%s\n", name, name)
		passCount := 0
		for i, tt := range testCases {
			got := method(tt.nums)

			// 验证结果是否正确
			valid := isValidSubsets(got, tt.expected)
			status := "✅"
			if !valid {
				status = "❌"
			} else {
				passCount++
			}
			fmt.Printf("  测试%d: %s\n", i+1, status)
			if status == "❌" {
				fmt.Printf("    输入: %v\n", tt.nums)
				fmt.Printf("    输出: %v\n", got)
				fmt.Printf("    期望: %v\n", tt.expected)
			}
		}
		fmt.Printf("  通过: %d/%d\n\n", passCount, len(testCases))
	}
}

// 验证子集是否正确
func isValidSubsets(got, expected [][]int) bool {
	if len(got) != len(expected) {
		return false
	}

	// 将结果转换为集合进行比较
	gotSet := make(map[string]bool)
	for _, subset := range got {
		key := fmt.Sprintf("%v", subset)
		gotSet[key] = true
	}

	expectedSet := make(map[string]bool)
	for _, subset := range expected {
		key := fmt.Sprintf("%v", subset)
		expectedSet[key] = true
	}

	// 比较两个集合
	for key := range gotSet {
		if !expectedSet[key] {
			return false
		}
	}

	for key := range expectedSet {
		if !gotSet[key] {
			return false
		}
	}

	return true
}
