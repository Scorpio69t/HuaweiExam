package main

import "fmt"

// =========================== 方法一：一维二分（最优解法） ===========================

func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}

	m, n := len(matrix), len(matrix[0])
	left, right := 0, m*n-1

	for left <= right {
		mid := left + (right-left)/2

		// 关键：一维索引转二维坐标
		row := mid / n
		col := mid % n
		midVal := matrix[row][col]

		if midVal == target {
			return true
		} else if midVal < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return false
}

// =========================== 方法二：两次二分 ===========================

func searchMatrix2(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}

	m, n := len(matrix), len(matrix[0])

	// 第一次二分：找目标行
	top, bottom := 0, m-1
	targetRow := -1

	for top <= bottom {
		mid := top + (bottom-top)/2

		if matrix[mid][0] <= target && target <= matrix[mid][n-1] {
			targetRow = mid
			break
		} else if matrix[mid][0] > target {
			bottom = mid - 1
		} else {
			top = mid + 1
		}
	}

	if targetRow == -1 {
		return false
	}

	// 第二次二分：在目标行中查找
	left, right := 0, n-1

	for left <= right {
		mid := left + (right-left)/2

		if matrix[targetRow][mid] == target {
			return true
		} else if matrix[targetRow][mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return false
}

// =========================== 方法三：从右上角搜索 ===========================

func searchMatrix3(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}

	m, n := len(matrix), len(matrix[0])
	row, col := 0, n-1 // 从右上角开始

	for row < m && col >= 0 {
		if matrix[row][col] == target {
			return true
		} else if matrix[row][col] > target {
			col-- // 当前值太大，向左移
		} else {
			row++ // 当前值太小，向下移
		}
	}

	return false
}

// =========================== 方法四：暴力搜索 ===========================

func searchMatrix4(matrix [][]int, target int) bool {
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j] == target {
				return true
			}
		}
	}
	return false
}

// =========================== 测试代码 ===========================

func main() {
	fmt.Println("=== LeetCode 74: 搜索二维矩阵 ===\n")

	testCases := []struct {
		matrix [][]int
		target int
		expect bool
	}{
		{
			[][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}},
			3,
			true,
		},
		{
			[][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}},
			13,
			false,
		},
		{
			[][]int{{5}},
			5,
			true,
		},
		{
			[][]int{{5}},
			3,
			false,
		},
		{
			[][]int{{1, 3, 5, 7}},
			1,
			true,
		},
		{
			[][]int{{1, 3, 5, 7}},
			7,
			true,
		},
		{
			[][]int{{1}, {3}, {5}, {7}},
			3,
			true,
		},
	}

	fmt.Println("方法一：一维二分（最优解法）")
	runTests(testCases, searchMatrix)

	fmt.Println("\n方法二：两次二分")
	runTests(testCases, searchMatrix2)

	fmt.Println("\n方法三：从右上角搜索")
	runTests(testCases, searchMatrix3)

	fmt.Println("\n方法四：暴力搜索")
	runTests(testCases, searchMatrix4)
}

func runTests(testCases []struct {
	matrix [][]int
	target int
	expect bool
}, fn func([][]int, int) bool) {
	passCount := 0
	for i, tc := range testCases {
		result := fn(tc.matrix, tc.target)
		status := "✅"
		if result != tc.expect {
			status = "❌"
		} else {
			passCount++
		}
		fmt.Printf("  测试%d: %s\n", i+1, status)
		if status == "❌" {
			fmt.Printf("    矩阵: %v, 目标: %d\n", tc.matrix, tc.target)
			fmt.Printf("    输出: %v\n", result)
			fmt.Printf("    期望: %v\n", tc.expect)
		}
	}
	fmt.Printf("  通过: %d/%d\n", passCount, len(testCases))
}
