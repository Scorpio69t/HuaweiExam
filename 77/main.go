package main

import (
	"fmt"
	"math/bits"
)

// =========================== 方法一：回溯算法（最优解法） ===========================

func combine(n int, k int) [][]int {
	var result [][]int
	var path []int

	var backtrack func(start int)
	backtrack = func(start int) {
		// 终止条件
		if len(path) == k {
			// 复制当前路径
			temp := make([]int, len(path))
			copy(temp, path)
			result = append(result, temp)
			return
		}

		// 剪枝：剩余数字不足
		if n-start+1 < k-len(path) {
			return
		}

		// 选择数字
		for i := start; i <= n; i++ {
			path = append(path, i)
			backtrack(i + 1)
			path = path[:len(path)-1] // 回溯
		}
	}

	backtrack(1)
	return result
}

// =========================== 方法二：迭代算法（非递归） ===========================

func combine2(n int, k int) [][]int {
	var result [][]int
	var stack [][]int

	// 初始化栈
	for i := 1; i <= n; i++ {
		stack = append(stack, []int{i})
	}

	for len(stack) > 0 {
		current := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if len(current) == k {
			result = append(result, current)
		} else {
			// 添加下一个数字
			start := current[len(current)-1] + 1
			for i := start; i <= n; i++ {
				newPath := make([]int, len(current))
				copy(newPath, current)
				newPath = append(newPath, i)
				stack = append(stack, newPath)
			}
		}
	}

	return result
}

// =========================== 方法三：位运算算法 ===========================

func combine3(n int, k int) [][]int {
	var result [][]int

	// 生成所有可能的位掩码
	for mask := 0; mask < (1 << n); mask++ {
		if bits.OnesCount(uint(mask)) == k {
			var path []int
			for i := 0; i < n; i++ {
				if mask&(1<<i) != 0 {
					path = append(path, i+1)
				}
			}
			result = append(result, path)
		}
	}

	return result
}

// =========================== 方法四：数学公式优化 ===========================

func combine4(n int, k int) [][]int {
	// 特殊情况处理
	if k == 0 {
		return [][]int{{}}
	}
	if k == n {
		var result []int
		for i := 1; i <= n; i++ {
			result = append(result, i)
		}
		return [][]int{result}
	}
	if k == 1 {
		var result [][]int
		for i := 1; i <= n; i++ {
			result = append(result, []int{i})
		}
		return result
	}

	// 使用回溯算法
	return combineBacktrack(n, k)
}

func combineBacktrack(n int, k int) [][]int {
	var result [][]int
	var path []int

	var backtrack func(start int)
	backtrack = func(start int) {
		if len(path) == k {
			temp := make([]int, len(path))
			copy(temp, path)
			result = append(result, temp)
			return
		}

		if n-start+1 < k-len(path) {
			return
		}

		for i := start; i <= n; i++ {
			path = append(path, i)
			backtrack(i + 1)
			path = path[:len(path)-1]
		}
	}

	backtrack(1)
	return result
}

// =========================== 测试代码 ===========================

func main() {
	fmt.Println("=== LeetCode 77: 组合 ===\n")

	testCases := []struct {
		n      int
		k      int
		expect [][]int
	}{
		{
			4, 2,
			[][]int{{1, 2}, {1, 3}, {1, 4}, {2, 3}, {2, 4}, {3, 4}},
		},
		{
			1, 1,
			[][]int{{1}},
		},
		{
			3, 3,
			[][]int{{1, 2, 3}},
		},
		{
			4, 1,
			[][]int{{1}, {2}, {3}, {4}},
		},
		{
			3, 2,
			[][]int{{1, 2}, {1, 3}, {2, 3}},
		},
		{
			5, 3,
			[][]int{{1, 2, 3}, {1, 2, 4}, {1, 2, 5}, {1, 3, 4}, {1, 3, 5}, {1, 4, 5}, {2, 3, 4}, {2, 3, 5}, {2, 4, 5}, {3, 4, 5}},
		},
	}

	fmt.Println("方法一：回溯算法（最优解法）")
	runTests(testCases, combine)

	fmt.Println("\n方法二：迭代算法（非递归）")
	runTests(testCases, combine2)

	fmt.Println("\n方法三：位运算算法")
	runTests(testCases, combine3)

	fmt.Println("\n方法四：数学公式优化")
	runTests(testCases, combine4)
}

func runTests(testCases []struct {
	n      int
	k      int
	expect [][]int
}, fn func(int, int) [][]int) {
	passCount := 0
	for i, tc := range testCases {
		result := fn(tc.n, tc.k)
		status := "✅"
		if !compareResults(result, tc.expect) {
			status = "❌"
		} else {
			passCount++
		}
		fmt.Printf("  测试%d: %s\n", i+1, status)
		if status == "❌" {
			fmt.Printf("    输入: n=%d, k=%d\n", tc.n, tc.k)
			fmt.Printf("    输出: %v\n", result)
			fmt.Printf("    期望: %v\n", tc.expect)
		}
	}
	fmt.Printf("  通过: %d/%d\n", passCount, len(testCases))
}

func compareResults(result, expect [][]int) bool {
	if len(result) != len(expect) {
		return false
	}

	// 创建结果映射
	resultMap := make(map[string]bool)
	for _, r := range result {
		key := fmt.Sprintf("%v", r)
		resultMap[key] = true
	}

	expectMap := make(map[string]bool)
	for _, e := range expect {
		key := fmt.Sprintf("%v", e)
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
