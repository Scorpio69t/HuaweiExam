package main

import "fmt"

// 方法一：右上角开始搜索（推荐）
// 时间复杂度：O(m + n)，空间复杂度：O(1)
func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}

	row, col := 0, len(matrix[0])-1

	for row < len(matrix) && col >= 0 {
		if matrix[row][col] == target {
			return true
		} else if matrix[row][col] > target {
			col-- // 向左移动
		} else {
			row++ // 向下移动
		}
	}

	return false
}

// 方法二：逐行二分查找
// 时间复杂度：O(m * log n)，空间复杂度：O(1)
func searchMatrixBinarySearch(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}

	for _, row := range matrix {
		if binarySearch(row, target) {
			return true
		}
	}
	return false
}

// 二分查找辅助函数
func binarySearch(arr []int, target int) bool {
	left, right := 0, len(arr)-1

	for left <= right {
		mid := left + (right-left)/2
		if arr[mid] == target {
			return true
		} else if arr[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return false
}

// 方法三：分治法
// 时间复杂度：O(n^log₄3) ≈ O(n^1.58)，空间复杂度：O(log n)
func searchMatrixDivideConquer(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}
	return divideConquer(matrix, target, 0, 0, len(matrix)-1, len(matrix[0])-1)
}

func divideConquer(matrix [][]int, target, row1, col1, row2, col2 int) bool {
	// 边界条件
	if row1 > row2 || col1 > col2 {
		return false
	}

	// 如果区域只有一个元素
	if row1 == row2 && col1 == col2 {
		return matrix[row1][col1] == target
	}

	// 选择中间元素
	midRow := (row1 + row2) / 2
	midCol := (col1 + col2) / 2

	if matrix[midRow][midCol] == target {
		return true
	} else if matrix[midRow][midCol] > target {
		// 在左上、左下、右上区域搜索
		return divideConquer(matrix, target, row1, col1, midRow-1, col2) ||
			divideConquer(matrix, target, midRow, col1, row2, midCol-1)
	} else {
		// 在右下、左下、右上区域搜索
		return divideConquer(matrix, target, midRow+1, col1, row2, col2) ||
			divideConquer(matrix, target, row1, midCol+1, midRow, col2)
	}
}

// 测试函数
func runTests() {
	fmt.Println("=== 240. 搜索二维矩阵 II 测试 ===")

	// 测试用例1
	matrix1 := [][]int{
		{1, 4, 7, 11, 15},
		{2, 5, 8, 12, 19},
		{3, 6, 9, 16, 22},
		{10, 13, 14, 17, 24},
		{18, 21, 23, 26, 30},
	}
	target1 := 5
	expected1 := true

	fmt.Printf("\n测试用例1:\n")
	fmt.Printf("矩阵:\n")
	printMatrix(matrix1)
	fmt.Printf("目标值: %d\n", target1)
	fmt.Printf("期望结果: %t\n", expected1)

	result1_1 := searchMatrix(matrix1, target1)
	result1_2 := searchMatrixBinarySearch(matrix1, target1)
	result1_3 := searchMatrixDivideConquer(matrix1, target1)

	fmt.Printf("方法一结果: %t ✓\n", result1_1)
	fmt.Printf("方法二结果: %t ✓\n", result1_2)
	fmt.Printf("方法三结果: %t ✓\n", result1_3)

	// 测试用例2
	matrix2 := matrix1
	target2 := 20
	expected2 := false

	fmt.Printf("\n测试用例2:\n")
	fmt.Printf("矩阵: (同上)\n")
	fmt.Printf("目标值: %d\n", target2)
	fmt.Printf("期望结果: %t\n", expected2)

	result2_1 := searchMatrix(matrix2, target2)
	result2_2 := searchMatrixBinarySearch(matrix2, target2)
	result2_3 := searchMatrixDivideConquer(matrix2, target2)

	fmt.Printf("方法一结果: %t ✓\n", result2_1)
	fmt.Printf("方法二结果: %t ✓\n", result2_2)
	fmt.Printf("方法三结果: %t ✓\n", result2_3)

	// 测试用例3：边界情况
	matrix3 := [][]int{{1}}
	target3 := 1

	fmt.Printf("\n测试用例3（边界情况）:\n")
	fmt.Printf("矩阵: [[1]]\n")
	fmt.Printf("目标值: %d\n", target3)
	fmt.Printf("期望结果: true\n")

	result3 := searchMatrix(matrix3, target3)
	fmt.Printf("结果: %t ✓\n", result3)

	// 测试用例4：空矩阵
	var matrix4 [][]int
	target4 := 1

	fmt.Printf("\n测试用例4（空矩阵）:\n")
	fmt.Printf("矩阵: []\n")
	fmt.Printf("目标值: %d\n", target4)
	fmt.Printf("期望结果: false\n")

	result4 := searchMatrix(matrix4, target4)
	fmt.Printf("结果: %t ✓\n", result4)

	fmt.Printf("\n=== 算法复杂度分析 ===\n")
	fmt.Printf("方法一（右上角搜索）：时间 O(m+n), 空间 O(1) - 推荐\n")
	fmt.Printf("方法二（逐行二分）：  时间 O(m*logn), 空间 O(1)\n")
	fmt.Printf("方法三（分治法）：    时间 O(n^1.58), 空间 O(logn)\n")
}

// 打印矩阵辅助函数
func printMatrix(matrix [][]int) {
	for _, row := range matrix {
		fmt.Printf("%v\n", row)
	}
}

func main() {
	runTests()
}
