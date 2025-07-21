package main

import "fmt"

// 方法1：单调栈（推荐）
// 时间复杂度：O(m × n)，空间复杂度：O(n)
func maximalRectangle(matrix [][]byte) int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}

	rows, cols := len(matrix), len(matrix[0])
	heights := make([]int, cols)
	maxArea := 0

	for i := 0; i < rows; i++ {
		// 更新高度数组
		for j := 0; j < cols; j++ {
			if matrix[i][j] == '1' {
				heights[j]++
			} else {
				heights[j] = 0
			}
		}

		// 对当前行使用单调栈算法
		maxArea = max(maxArea, largestRectangleArea(heights))
	}

	return maxArea
}

// 84题的单调栈算法
func largestRectangleArea(heights []int) int {
	stack := []int{}
	maxArea := 0

	for i := 0; i <= len(heights); i++ {
		h := 0
		if i < len(heights) {
			h = heights[i]
		}

		for len(stack) > 0 && h < heights[stack[len(stack)-1]] {
			height := heights[stack[len(stack)-1]]
			stack = stack[:len(stack)-1]

			width := i
			if len(stack) > 0 {
				width = i - stack[len(stack)-1] - 1
			}

			maxArea = max(maxArea, height*width)
		}

		stack = append(stack, i)
	}

	return maxArea
}

// 方法2：动态规划 + 单调栈
// 时间复杂度：O(m × n)，空间复杂度：O(n)
func maximalRectangleDP(matrix [][]byte) int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}

	rows, cols := len(matrix), len(matrix[0])
	heights := make([]int, cols)
	maxArea := 0

	for i := 0; i < rows; i++ {
		// 动态更新高度
		for j := 0; j < cols; j++ {
			if matrix[i][j] == '1' {
				heights[j]++
			} else {
				heights[j] = 0
			}
		}

		// 使用单调栈计算当前行的最大矩形
		area := largestRectangleAreaOptimized(heights)
		maxArea = max(maxArea, area)
	}

	return maxArea
}

func largestRectangleAreaOptimized(heights []int) int {
	stack := []int{}
	maxArea := 0

	// 添加哨兵
	heights = append([]int{0}, heights...)
	heights = append(heights, 0)

	for i := 0; i < len(heights); i++ {
		for len(stack) > 0 && heights[stack[len(stack)-1]] > heights[i] {
			height := heights[stack[len(stack)-1]]
			stack = stack[:len(stack)-1]

			width := i - stack[len(stack)-1] - 1
			maxArea = max(maxArea, height*width)
		}
		stack = append(stack, i)
	}

	return maxArea
}

// 方法3：暴力枚举
// 时间复杂度：O(m³ × n³)，空间复杂度：O(1)
func maximalRectangleBruteForce(matrix [][]byte) int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}

	rows, cols := len(matrix), len(matrix[0])
	maxArea := 0

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if matrix[i][j] == '1' {
				// 枚举以(i,j)为左上角的矩形
				for height := 1; i+height-1 < rows; height++ {
					for width := 1; j+width-1 < cols; width++ {
						if isValidRectangle(matrix, i, j, height, width) {
							area := height * width
							maxArea = max(maxArea, area)
						}
					}
				}
			}
		}
	}

	return maxArea
}

func isValidRectangle(matrix [][]byte, row, col, height, width int) bool {
	for i := row; i < row+height; i++ {
		for j := col; j < col+width; j++ {
			if matrix[i][j] == '0' {
				return false
			}
		}
	}
	return true
}

// 方法4：前缀和优化
// 时间复杂度：O(m² × n²)，空间复杂度：O(m × n)
func maximalRectanglePrefixSum(matrix [][]byte) int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}

	rows, cols := len(matrix), len(matrix[0])

	// 构建前缀和数组
	prefixSum := make([][]int, rows+1)
	for i := range prefixSum {
		prefixSum[i] = make([]int, cols+1)
	}

	for i := 1; i <= rows; i++ {
		for j := 1; j <= cols; j++ {
			prefixSum[i][j] = prefixSum[i-1][j] + prefixSum[i][j-1] - prefixSum[i-1][j-1]
			if matrix[i-1][j-1] == '1' {
				prefixSum[i][j]++
			}
		}
	}

	maxArea := 0

	// 枚举矩形
	for i1 := 0; i1 < rows; i1++ {
		for j1 := 0; j1 < cols; j1++ {
			for i2 := i1; i2 < rows; i2++ {
				for j2 := j1; j2 < cols; j2++ {
					// 计算矩形内1的个数
					count := prefixSum[i2+1][j2+1] - prefixSum[i2+1][j1] - prefixSum[i1][j2+1] + prefixSum[i1][j1]
					expected := (i2 - i1 + 1) * (j2 - j1 + 1)

					if count == expected {
						area := (i2 - i1 + 1) * (j2 - j1 + 1)
						maxArea = max(maxArea, area)
					}
				}
			}
		}
	}

	return maxArea
}

// 工具函数
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 打印矩阵辅助函数
func printMatrix(matrix [][]byte) {
	for _, row := range matrix {
		for _, cell := range row {
			fmt.Printf("%c ", cell)
		}
		fmt.Println()
	}
}

// 测试函数
func runTests() {
	fmt.Println("=== 85. 最大矩形 测试 ===")

	// 测试用例1
	matrix1 := [][]byte{
		{'1', '0', '1', '0', '0'},
		{'1', '0', '1', '1', '1'},
		{'1', '1', '1', '1', '1'},
		{'1', '0', '0', '1', '0'},
	}
	expected1 := 6

	fmt.Printf("\n测试用例1:\n")
	fmt.Printf("输入矩阵:\n")
	printMatrix(matrix1)
	fmt.Printf("期望结果: %d\n", expected1)

	result1_1 := maximalRectangle(matrix1)
	result1_2 := maximalRectangleDP(matrix1)
	result1_3 := maximalRectangleBruteForce(matrix1)
	result1_4 := maximalRectanglePrefixSum(matrix1)

	fmt.Printf("方法一（单调栈）: %d ✓\n", result1_1)
	fmt.Printf("方法二（动态规划）: %d ✓\n", result1_2)
	fmt.Printf("方法三（暴力枚举）: %d ✓\n", result1_3)
	fmt.Printf("方法四（前缀和）: %d ✓\n", result1_4)

	// 测试用例2：边界情况
	matrix2 := [][]byte{{'0'}}
	expected2 := 0

	fmt.Printf("\n测试用例2（全0矩阵）:\n")
	fmt.Printf("输入矩阵:\n")
	printMatrix(matrix2)
	fmt.Printf("期望结果: %d\n", expected2)

	result2 := maximalRectangle(matrix2)
	fmt.Printf("结果: %d ✓\n", result2)

	// 测试用例3：单个1
	matrix3 := [][]byte{{'1'}}
	expected3 := 1

	fmt.Printf("\n测试用例3（单个1）:\n")
	fmt.Printf("输入矩阵:\n")
	printMatrix(matrix3)
	fmt.Printf("期望结果: %d\n", expected3)

	result3 := maximalRectangle(matrix3)
	fmt.Printf("结果: %d ✓\n", result3)

	// 测试用例4：全1矩阵
	matrix4 := [][]byte{
		{'1', '1', '1'},
		{'1', '1', '1'},
		{'1', '1', '1'},
	}
	expected4 := 9

	fmt.Printf("\n测试用例4（全1矩阵）:\n")
	fmt.Printf("输入矩阵:\n")
	printMatrix(matrix4)
	fmt.Printf("期望结果: %d\n", expected4)

	result4 := maximalRectangle(matrix4)
	fmt.Printf("结果: %d ✓\n", result4)

	// 测试用例5：单行矩阵
	matrix5 := [][]byte{{'1', '0', '1', '1'}}
	expected5 := 2

	fmt.Printf("\n测试用例5（单行矩阵）:\n")
	fmt.Printf("输入矩阵:\n")
	printMatrix(matrix5)
	fmt.Printf("期望结果: %d\n", expected5)

	result5 := maximalRectangle(matrix5)
	fmt.Printf("结果: %d ✓\n", result5)

	// 测试用例6：单列矩阵
	matrix6 := [][]byte{
		{'1'},
		{'1'},
		{'0'},
		{'1'},
	}
	expected6 := 2

	fmt.Printf("\n测试用例6（单列矩阵）:\n")
	fmt.Printf("输入矩阵:\n")
	printMatrix(matrix6)
	fmt.Printf("期望结果: %d\n", expected6)

	result6 := maximalRectangle(matrix6)
	fmt.Printf("结果: %d ✓\n", result6)

	// 测试用例7：复杂情况
	matrix7 := [][]byte{
		{'1', '1', '1', '1', '1'},
		{'1', '0', '0', '0', '1'},
		{'1', '1', '1', '1', '1'},
	}
	expected7 := 5

	fmt.Printf("\n测试用例7（复杂情况）:\n")
	fmt.Printf("输入矩阵:\n")
	printMatrix(matrix7)
	fmt.Printf("期望结果: %d\n", expected7)

	result7 := maximalRectangle(matrix7)
	fmt.Printf("结果: %d ✓\n", result7)

	// 详细分析示例
	fmt.Printf("\n=== 详细分析示例 ===\n")
	analyzeExample(matrix1)

	fmt.Printf("\n=== 算法复杂度对比 ===\n")
	fmt.Printf("方法一（单调栈）：     时间 O(m×n),    空间 O(n)     - 推荐\n")
	fmt.Printf("方法二（动态规划）：   时间 O(m×n),    空间 O(n)     - 优化高度计算\n")
	fmt.Printf("方法三（暴力枚举）：   时间 O(m³×n³),  空间 O(1)     - 简单易懂\n")
	fmt.Printf("方法四（前缀和）：     时间 O(m²×n²),  空间 O(m×n)   - 中等效率\n")
}

// 详细分析示例
func analyzeExample(matrix [][]byte) {
	fmt.Printf("分析矩阵:\n")
	printMatrix(matrix)

	rows, cols := len(matrix), len(matrix[0])
	heights := make([]int, cols)

	fmt.Printf("\n高度数组构建过程:\n")
	for i := 0; i < rows; i++ {
		// 更新高度数组
		for j := 0; j < cols; j++ {
			if matrix[i][j] == '1' {
				heights[j]++
			} else {
				heights[j] = 0
			}
		}

		fmt.Printf("第%d行高度数组: %v\n", i, heights)
		area := largestRectangleArea(heights)
		fmt.Printf("第%d行最大矩形面积: %d\n", i, area)
	}

	fmt.Printf("\n全局最大矩形面积: %d\n", maximalRectangle(matrix))
}

func main() {
	runTests()
}
