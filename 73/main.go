package main

import "fmt"

// =========================== 方法一：原地标记（O(1)空间，最优解法） ===========================

func setZeroes(matrix [][]int) {
	if len(matrix) == 0 {
		return
	}

	m, n := len(matrix), len(matrix[0])
	firstRowZero, firstColZero := false, false

	// 检查第一行是否有0
	for j := 0; j < n; j++ {
		if matrix[0][j] == 0 {
			firstRowZero = true
			break
		}
	}

	// 检查第一列是否有0
	for i := 0; i < m; i++ {
		if matrix[i][0] == 0 {
			firstColZero = true
			break
		}
	}

	// 用第一行和第一列标记
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if matrix[i][j] == 0 {
				matrix[i][0] = 0
				matrix[0][j] = 0
			}
		}
	}

	// 根据标记置零（从后向前）
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if matrix[i][0] == 0 || matrix[0][j] == 0 {
				matrix[i][j] = 0
			}
		}
	}

	// 处理第一行
	if firstRowZero {
		for j := 0; j < n; j++ {
			matrix[0][j] = 0
		}
	}

	// 处理第一列
	if firstColZero {
		for i := 0; i < m; i++ {
			matrix[i][0] = 0
		}
	}
}

// =========================== 方法二：额外数组（O(m+n)空间） ===========================

func setZeroes2(matrix [][]int) {
	if len(matrix) == 0 {
		return
	}

	m, n := len(matrix), len(matrix[0])
	rows := make([]bool, m)
	cols := make([]bool, n)

	// 记录需要置零的行和列
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == 0 {
				rows[i] = true
				cols[j] = true
			}
		}
	}

	// 置零
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if rows[i] || cols[j] {
				matrix[i][j] = 0
			}
		}
	}
}

// =========================== 方法三：使用集合标记 ===========================

func setZeroes3(matrix [][]int) {
	if len(matrix) == 0 {
		return
	}

	m, n := len(matrix), len(matrix[0])
	rowSet := make(map[int]bool)
	colSet := make(map[int]bool)

	// 记录需要置零的行和列
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == 0 {
				rowSet[i] = true
				colSet[j] = true
			}
		}
	}

	// 置零
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if rowSet[i] || colSet[j] {
				matrix[i][j] = 0
			}
		}
	}
}

// =========================== 方法四：两个变量优化（O(1)空间） ===========================

func setZeroes4(matrix [][]int) {
	if len(matrix) == 0 {
		return
	}

	m, n := len(matrix), len(matrix[0])
	col0 := false

	// 使用matrix[i][0]和matrix[0][j]作为标记
	for i := 0; i < m; i++ {
		if matrix[i][0] == 0 {
			col0 = true
		}
		for j := 1; j < n; j++ {
			if matrix[i][j] == 0 {
				matrix[i][0] = 0
				matrix[0][j] = 0
			}
		}
	}

	// 从后向前置零
	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 1; j-- {
			if matrix[i][0] == 0 || matrix[0][j] == 0 {
				matrix[i][j] = 0
			}
		}
		if col0 {
			matrix[i][0] = 0
		}
	}
}

// =========================== 测试代码 ===========================

func main() {
	fmt.Println("=== LeetCode 73: 矩阵置零 ===\n")

	testCases := []struct {
		matrix [][]int
		expect [][]int
	}{
		{
			[][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}},
			[][]int{{1, 0, 1}, {0, 0, 0}, {1, 0, 1}},
		},
		{
			[][]int{{0, 1, 2, 0}, {3, 4, 5, 2}, {1, 3, 1, 5}},
			[][]int{{0, 0, 0, 0}, {0, 4, 5, 0}, {0, 3, 1, 0}},
		},
		{
			[][]int{{1}},
			[][]int{{1}},
		},
		{
			[][]int{{0}},
			[][]int{{0}},
		},
	}

	fmt.Println("方法一：原地标记（O(1)空间）")
	runTests(testCases, setZeroes)

	fmt.Println("\n方法二：额外数组（O(m+n)空间）")
	runTests(testCases, setZeroes2)

	fmt.Println("\n方法三：使用集合")
	runTests(testCases, setZeroes3)

	fmt.Println("\n方法四：两个变量优化")
	runTests(testCases, setZeroes4)
}

func runTests(testCases []struct {
	matrix [][]int
	expect [][]int
}, fn func([][]int)) {
	passCount := 0
	for i, tc := range testCases {
		matrix := copyMatrix(tc.matrix)
		fn(matrix)
		status := "✅"
		if !equalMatrix(matrix, tc.expect) {
			status = "❌"
		} else {
			passCount++
		}
		fmt.Printf("  测试%d: %s\n", i+1, status)
		if status == "❌" {
			fmt.Printf("    输入: %v\n", tc.matrix)
			fmt.Printf("    输出: %v\n", matrix)
			fmt.Printf("    期望: %v\n", tc.expect)
		}
	}
	fmt.Printf("  通过: %d/%d\n", passCount, len(testCases))
}

func copyMatrix(matrix [][]int) [][]int {
	result := make([][]int, len(matrix))
	for i := range matrix {
		result[i] = make([]int, len(matrix[i]))
		copy(result[i], matrix[i])
	}
	return result
}

func equalMatrix(a, b [][]int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if len(a[i]) != len(b[i]) {
			return false
		}
		for j := range a[i] {
			if a[i][j] != b[i][j] {
				return false
			}
		}
	}
	return true
}
