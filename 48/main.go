package main

import (
	"fmt"
	"time"
)

// 方法一：转置+翻转算法
// 最直观的解法，先转置矩阵，再翻转每一行
func rotate1(matrix [][]int) {
	n := len(matrix)

	// 转置矩阵
	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}

	// 翻转每一行
	for i := 0; i < n; i++ {
		for j := 0; j < n/2; j++ {
			matrix[i][j], matrix[i][n-1-j] = matrix[i][n-1-j], matrix[i][j]
		}
	}
}

// 方法二：四角交换算法
// 效率最高的解法，直接交换四个相关位置的元素
func rotate2(matrix [][]int) {
	n := len(matrix)

	for i := 0; i < n/2; i++ {
		for j := i; j < n-1-i; j++ {
			// 保存左上角
			temp := matrix[i][j]

			// 右下角 → 左上角
			matrix[i][j] = matrix[n-1-j][i]

			// 左下角 → 右下角
			matrix[n-1-j][i] = matrix[n-1-i][n-1-j]

			// 右上角 → 左下角
			matrix[n-1-i][n-1-j] = matrix[j][n-1-i]

			// 保存的元素 → 右上角
			matrix[j][n-1-i] = temp
		}
	}
}

// 方法三：数学公式算法
// 使用数学公式计算旋转后的坐标
func rotate3(matrix [][]int) {
	n := len(matrix)

	for i := 0; i < n/2; i++ {
		for j := i; j < n-1-i; j++ {
			// 计算旋转后的坐标
			newI, newJ := j, n-1-i

			// 交换元素
			matrix[i][j], matrix[newI][newJ] = matrix[newI][newJ], matrix[i][j]
		}
	}
}

// 方法四：分块旋转算法
// 分块处理，适合大矩阵
func rotate4(matrix [][]int) {
	n := len(matrix)

	// 分块处理
	blockSize := 4
	for i := 0; i < n; i += blockSize {
		for j := 0; j < n; j += blockSize {
			rotateBlockHelper(matrix, i, j, min(i+blockSize, n), min(j+blockSize, n))
		}
	}
}

// 分块旋转的辅助函数
func rotateBlockHelper(matrix [][]int, startI, startJ, endI, endJ int) {
	// 处理当前块内的旋转
	for i := startI; i < endI; i++ {
		for j := startJ; j < endJ; j++ {
			// 计算旋转后的坐标
			newI, newJ := j, len(matrix)-1-i

			// 交换元素
			matrix[i][j], matrix[newI][newJ] = matrix[newI][newJ], matrix[i][j]
		}
	}
}

// 辅助函数：min函数
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// 辅助函数：创建测试用例
func createTestCases() []struct {
	matrix [][]int
	name   string
} {
	return []struct {
		matrix [][]int
		name   string
	}{
		{[][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}, "示例1: 3×3矩阵"},
		{[][]int{{5, 1, 9, 11}, {2, 4, 8, 10}, {13, 3, 6, 7}, {15, 14, 12, 16}}, "示例2: 4×4矩阵"},
		{[][]int{{1, 2}, {3, 4}}, "测试1: 2×2矩阵"},
		{[][]int{{1}}, "测试2: 1×1矩阵"},
		{[][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}, {13, 14, 15, 16}}, "测试3: 4×4矩阵"},
		{[][]int{{1, 2, 3, 4, 5}, {6, 7, 8, 9, 10}, {11, 12, 13, 14, 15}, {16, 17, 18, 19, 20}, {21, 22, 23, 24, 25}}, "测试4: 5×5矩阵"},
		{[][]int{{1, 2, 3, 4, 5, 6}, {7, 8, 9, 10, 11, 12}, {13, 14, 15, 16, 17, 18}, {19, 20, 21, 22, 23, 24}, {25, 26, 27, 28, 29, 30}, {31, 32, 33, 34, 35, 36}}, "测试5: 6×6矩阵"},
		{[][]int{{1, 2, 3, 4, 5, 6, 7}, {8, 9, 10, 11, 12, 13, 14}, {15, 16, 17, 18, 19, 20, 21}, {22, 23, 24, 25, 26, 27, 28}, {29, 30, 31, 32, 33, 34, 35}, {36, 37, 38, 39, 40, 41, 42}, {43, 44, 45, 46, 47, 48, 49}}, "测试6: 7×7矩阵"},
		{[][]int{{1, 2, 3, 4, 5, 6, 7, 8}, {9, 10, 11, 12, 13, 14, 15, 16}, {17, 18, 19, 20, 21, 22, 23, 24}, {25, 26, 27, 28, 29, 30, 31, 32}, {33, 34, 35, 36, 37, 38, 39, 40}, {41, 42, 43, 44, 45, 46, 47, 48}, {49, 50, 51, 52, 53, 54, 55, 56}, {57, 58, 59, 60, 61, 62, 63, 64}}, "测试7: 8×8矩阵"},
		{[][]int{{1, 2, 3, 4, 5, 6, 7, 8, 9}, {10, 11, 12, 13, 14, 15, 16, 17, 18}, {19, 20, 21, 22, 23, 24, 25, 26, 27}, {28, 29, 30, 31, 32, 33, 34, 35, 36}, {37, 38, 39, 40, 41, 42, 43, 44, 45}, {46, 47, 48, 49, 50, 51, 52, 53, 54}, {55, 56, 57, 58, 59, 60, 61, 62, 63}, {64, 65, 66, 67, 68, 69, 70, 71, 72}, {73, 74, 75, 76, 77, 78, 79, 80, 81}}, "测试8: 9×9矩阵"},
	}
}

// 性能测试函数
func benchmarkAlgorithm(algorithm func([][]int), matrix [][]int, name string) {
	iterations := 100
	start := time.Now()

	for i := 0; i < iterations; i++ {
		// 复制矩阵，避免修改原矩阵
		tempMatrix := make([][]int, len(matrix))
		for j := range matrix {
			tempMatrix[j] = make([]int, len(matrix[j]))
			copy(tempMatrix[j], matrix[j])
		}
		algorithm(tempMatrix)
	}

	duration := time.Since(start)
	avgTime := duration.Nanoseconds() / int64(iterations)

	fmt.Printf("%s: 平均执行时间 %d 纳秒\n", name, avgTime)
}

// 辅助函数：验证结果是否正确
func validateResult(original, rotated [][]int) bool {
	n := len(original)

	// 验证旋转后的矩阵是否正确
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			// 旋转后的元素应该在正确的位置
			expectedValue := original[n-1-j][i]
			if rotated[i][j] != expectedValue {
				return false
			}
		}
	}

	return true
}

// 辅助函数：复制矩阵
func copyMatrix(matrix [][]int) [][]int {
	result := make([][]int, len(matrix))
	for i := range matrix {
		result[i] = make([]int, len(matrix[i]))
		copy(result[i], matrix[i])
	}
	return result
}

// 辅助函数：打印矩阵
func printMatrix(matrix [][]int, title string) {
	fmt.Printf("%s:\n", title)
	for _, row := range matrix {
		fmt.Printf("  %v\n", row)
	}
}

// 辅助函数：比较矩阵
func compareMatrix(matrix1, matrix2 [][]int) bool {
	if len(matrix1) != len(matrix2) {
		return false
	}

	for i := range matrix1 {
		if len(matrix1[i]) != len(matrix2[i]) {
			return false
		}
		for j := range matrix1[i] {
			if matrix1[i][j] != matrix2[i][j] {
				return false
			}
		}
	}

	return true
}

func main() {
	fmt.Println("=== 48. 旋转图像 ===")
	fmt.Println()

	// 创建测试用例
	testCases := createTestCases()
	algorithms := []struct {
		name string
		fn   func([][]int)
	}{
		{"转置+翻转算法", rotate1},
		{"四角交换算法", rotate2},
		{"数学公式算法", rotate3},
		{"分块旋转算法", rotate4},
	}

	// 运行测试
	fmt.Println("=== 算法正确性测试 ===")
	for _, testCase := range testCases {
		fmt.Printf("测试: %s\n", testCase.name)

		results := make([][][]int, len(algorithms))
		for i, algo := range algorithms {
			// 复制原始矩阵
			tempMatrix := copyMatrix(testCase.matrix)
			algo.fn(tempMatrix)
			results[i] = tempMatrix
		}

		// 验证所有算法结果一致
		allEqual := true
		for i := 1; i < len(results); i++ {
			if !compareMatrix(results[i], results[0]) {
				allEqual = false
				break
			}
		}

		// 验证结果是否正确
		allValid := true
		for _, result := range results {
			if !validateResult(testCase.matrix, result) {
				allValid = false
				break
			}
		}

		if allEqual && allValid {
			fmt.Printf("  ✅ 所有算法结果一致且正确\n")
			if len(testCase.matrix) <= 3 {
				printMatrix(results[0], "  旋转结果")
			}
		} else {
			fmt.Printf("  ❌ 算法结果不一致或错误\n")
			for i, algo := range algorithms {
				fmt.Printf("    %s: 结果长度 %d×%d\n", algo.name, len(results[i]), len(results[i][0]))
			}
		}
		fmt.Println()
	}

	// 性能测试
	fmt.Println("=== 性能测试 ===")
	performanceMatrix := [][]int{
		{1, 2, 3, 4, 5},
		{6, 7, 8, 9, 10},
		{11, 12, 13, 14, 15},
		{16, 17, 18, 19, 20},
		{21, 22, 23, 24, 25},
	}

	fmt.Printf("测试数据: 5×5矩阵\n")
	fmt.Println()

	for _, algo := range algorithms {
		benchmarkAlgorithm(algo.fn, performanceMatrix, algo.name)
	}
	fmt.Println()

	// 算法分析
	fmt.Println("=== 算法分析 ===")
	fmt.Println("旋转图像问题的特点:")
	fmt.Println("1. 需要将n×n矩阵顺时针旋转90度")
	fmt.Println("2. 必须在原地旋转，不能使用额外空间")
	fmt.Println("3. 需要理解旋转的数学原理")
	fmt.Println("4. 四角交换算法是最优解法")
	fmt.Println()

	// 复杂度分析
	fmt.Println("=== 复杂度分析 ===")
	fmt.Println("时间复杂度:")
	fmt.Println("- 转置+翻转: O(n²)，需要遍历整个矩阵两次")
	fmt.Println("- 四角交换: O(n²)，需要遍历整个矩阵一次")
	fmt.Println("- 数学公式: O(n²)，需要遍历整个矩阵一次")
	fmt.Println("- 分块旋转: O(n²)，需要遍历整个矩阵一次")
	fmt.Println()
	fmt.Println("空间复杂度:")
	fmt.Println("- 所有算法: O(1)，只使用常数空间，原地操作")
	fmt.Println()

	// 算法总结
	fmt.Println("=== 算法总结 ===")
	fmt.Println("1. 转置+翻转算法：最直观的解法，逻辑清晰")
	fmt.Println("2. 四角交换算法：效率最高的解法，直接交换")
	fmt.Println("3. 数学公式算法：使用数学变换，代码简洁")
	fmt.Println("4. 分块旋转算法：适合大矩阵，分块处理")
	fmt.Println()
	fmt.Println("推荐使用：四角交换算法（方法二），效率最高")
	fmt.Println()

	// 应用场景
	fmt.Println("=== 应用场景 ===")
	fmt.Println("- 图像处理：旋转图像文件")
	fmt.Println("- 游戏开发：旋转游戏对象")
	fmt.Println("- 数据分析：旋转数据矩阵")
	fmt.Println("- 算法竞赛：矩阵操作的经典应用")
	fmt.Println("- 计算机图形学：2D变换的基础操作")
	fmt.Println()

	// 优化技巧总结
	fmt.Println("=== 优化技巧总结 ===")
	fmt.Println("1. 转置+翻转：最直观的解法，逻辑清晰")
	fmt.Println("2. 四角交换：效率最高的解法，直接交换")
	fmt.Println("3. 数学公式：使用数学变换，代码简洁")
	fmt.Println("4. 分块旋转：适合大矩阵，分块处理")
	fmt.Println("5. 坐标映射：理解旋转的数学原理")
	fmt.Println("6. 原地操作：避免使用额外空间")
}
