package main

import (
	"fmt"
	"strings"
)

// convert Z字形变换
// 时间复杂度: O(n)，其中n是字符串长度
// 空间复杂度: O(n)
func convert(s string, numRows int) string {
	if numRows == 1 || numRows >= len(s) {
		return s
	}

	// 创建numRows个字符串构建器
	rows := make([]strings.Builder, numRows)

	// 当前行索引和方向
	currentRow := 0
	direction := 1 // 1表示向下，-1表示向上

	// 遍历字符串，按Z字形填充
	for _, char := range s {
		// 将当前字符添加到当前行
		rows[currentRow].WriteByte(byte(char))

		// 更新行索引
		currentRow += direction

		// 如果到达边界，改变方向
		if currentRow == 0 || currentRow == numRows-1 {
			direction = -direction
		}
	}

	// 将所有行连接起来
	var result strings.Builder
	for _, row := range rows {
		result.WriteString(row.String())
	}

	return result.String()
}

// convertOptimized 优化版本，使用数学规律
// 时间复杂度: O(n)
// 空间复杂度: O(n)
func convertOptimized(s string, numRows int) string {
	if numRows == 1 || numRows >= len(s) {
		return s
	}

	var result strings.Builder
	n := len(s)

	// 第一行和最后一行的字符间隔
	cycleLen := 2*numRows - 2

	// 逐行构建结果
	for row := 0; row < numRows; row++ {
		for i := row; i < n; i += cycleLen {
			// 添加垂直方向的字符
			result.WriteByte(s[i])

			// 如果不是第一行和最后一行，还需要添加斜向的字符
			if row != 0 && row != numRows-1 {
				// 计算斜向字符的位置
				diagonalIndex := i + cycleLen - 2*row
				if diagonalIndex < n {
					result.WriteByte(s[diagonalIndex])
				}
			}
		}
	}

	return result.String()
}

// convertSimulation 模拟法：实际构建Z字形矩阵
// 时间复杂度: O(n)
// 空间复杂度: O(n*numRows)
func convertSimulation(s string, numRows int) string {
	if numRows == 1 || numRows >= len(s) {
		return s
	}

	// 创建Z字形矩阵
	matrix := make([][]byte, numRows)
	for i := range matrix {
		matrix[i] = make([]byte, 0)
	}

	currentRow := 0
	direction := 1

	// 填充矩阵
	for _, char := range s {
		matrix[currentRow] = append(matrix[currentRow], byte(char))
		currentRow += direction

		if currentRow == 0 || currentRow == numRows-1 {
			direction = -direction
		}
	}

	// 按行读取结果
	var result strings.Builder
	for _, row := range matrix {
		result.Write(row)
	}

	return result.String()
}

func main() {
	// 测试用例1
	s1 := "PAYPALISHIRING"
	numRows1 := 3
	result1 := convert(s1, numRows1)
	fmt.Printf("示例1: s = \"%s\", numRows = %d\n", s1, numRows1)
	fmt.Printf("输出: \"%s\"\n", result1)
	fmt.Printf("期望: \"PAHNAPLSIIGYIR\"\n")
	fmt.Printf("结果: %t\n", result1 == "PAHNAPLSIIGYIR")
	fmt.Println()

	// 测试用例2
	s2 := "PAYPALISHIRING"
	numRows2 := 4
	result2 := convert(s2, numRows2)
	fmt.Printf("示例2: s = \"%s\", numRows = %d\n", s2, numRows2)
	fmt.Printf("输出: \"%s\"\n", result2)
	fmt.Printf("期望: \"PINALSIGYAHRPI\"\n")
	fmt.Printf("结果: %t\n", result2 == "PINALSIGYAHRPI")
	fmt.Println()

	// 测试用例3
	s3 := "A"
	numRows3 := 1
	result3 := convert(s3, numRows3)
	fmt.Printf("示例3: s = \"%s\", numRows = %d\n", s3, numRows3)
	fmt.Printf("输出: \"%s\"\n", result3)
	fmt.Printf("期望: \"A\"\n")
	fmt.Printf("结果: %t\n", result3 == "A")
	fmt.Println()

	// 额外测试用例
	s4 := "AB"
	numRows4 := 1
	result4 := convert(s4, numRows4)
	fmt.Printf("额外测试: s = \"%s\", numRows = %d\n", s4, numRows4)
	fmt.Printf("输出: \"%s\"\n", result4)
	fmt.Printf("期望: \"AB\"\n")
	fmt.Printf("结果: %t\n", result4 == "AB")
	fmt.Println()

	// 测试优化版本
	fmt.Println("=== 优化版本测试 ===")
	result1Opt := convertOptimized(s1, numRows1)
	result2Opt := convertOptimized(s2, numRows2)
	fmt.Printf("优化版本示例1: %s\n", result1Opt)
	fmt.Printf("优化版本示例2: %s\n", result2Opt)
	fmt.Printf("结果一致: %t\n", result1Opt == result1 && result2Opt == result2)
	fmt.Println()

	// 测试模拟版本
	fmt.Println("=== 模拟版本测试 ===")
	result1Sim := convertSimulation(s1, numRows1)
	result2Sim := convertSimulation(s2, numRows2)
	fmt.Printf("模拟版本示例1: %s\n", result1Sim)
	fmt.Printf("模拟版本示例2: %s\n", result2Sim)
	fmt.Printf("结果一致: %t\n", result1Sim == result1 && result2Sim == result2)
}
