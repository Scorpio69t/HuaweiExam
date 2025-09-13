package main

import (
	"fmt"
	"time"
)

// 方法一：三次遍历验证法
// 分别验证行、列、3×3宫格
func isValidSudoku1(board [][]byte) bool {
	// 验证每一行
	for i := 0; i < 9; i++ {
		seen := make(map[byte]bool)
		for j := 0; j < 9; j++ {
			if board[i][j] != '.' {
				if seen[board[i][j]] {
					return false
				}
				seen[board[i][j]] = true
			}
		}
	}

	// 验证每一列
	for j := 0; j < 9; j++ {
		seen := make(map[byte]bool)
		for i := 0; i < 9; i++ {
			if board[i][j] != '.' {
				if seen[board[i][j]] {
					return false
				}
				seen[board[i][j]] = true
			}
		}
	}

	// 验证每个3×3宫格
	for box := 0; box < 9; box++ {
		seen := make(map[byte]bool)
		startRow := (box / 3) * 3
		startCol := (box % 3) * 3
		for i := startRow; i < startRow+3; i++ {
			for j := startCol; j < startCol+3; j++ {
				if board[i][j] != '.' {
					if seen[board[i][j]] {
						return false
					}
					seen[board[i][j]] = true
				}
			}
		}
	}

	return true
}

// 方法二：一次遍历优化法
// 使用三个哈希表同时记录行、列、宫格的数字使用情况
func isValidSudoku2(board [][]byte) bool {
	// 使用数组代替哈希表，提高访问速度
	rows := [9][10]bool{}  // 行标记数组
	cols := [9][10]bool{}  // 列标记数组
	boxes := [9][10]bool{} // 宫格标记数组

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == '.' {
				continue
			}
			num := int(board[i][j] - '0')
			boxIndex := (i/3)*3 + (j / 3)

			// 检查是否已经存在冲突
			if rows[i][num] || cols[j][num] || boxes[boxIndex][num] {
				return false
			}

			// 标记数字已使用
			rows[i][num] = true
			cols[j][num] = true
			boxes[boxIndex][num] = true
		}
	}
	return true
}

// 方法三：位运算优化法
// 使用位掩码压缩空间使用，最高效的解法
func isValidSudoku3(board [][]byte) bool {
	rows := [9]int{}  // 行位掩码
	cols := [9]int{}  // 列位掩码
	boxes := [9]int{} // 宫格位掩码

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == '.' {
				continue
			}
			num := int(board[i][j] - '0')
			bit := 1 << num
			boxIndex := (i/3)*3 + (j / 3)

			// 使用位运算检查冲突
			if (rows[i]&bit) != 0 || (cols[j]&bit) != 0 || (boxes[boxIndex]&bit) != 0 {
				return false
			}

			// 使用位运算更新掩码
			rows[i] |= bit
			cols[j] |= bit
			boxes[boxIndex] |= bit
		}
	}
	return true
}

// 方法四：集合验证法
// 使用Set数据结构，代码最简洁
func isValidSudoku4(board [][]byte) bool {
	// 创建三个集合分别记录行、列、宫格的数字
	rows := make([]map[byte]bool, 9)
	cols := make([]map[byte]bool, 9)
	boxes := make([]map[byte]bool, 9)

	// 初始化集合
	for i := 0; i < 9; i++ {
		rows[i] = make(map[byte]bool)
		cols[i] = make(map[byte]bool)
		boxes[i] = make(map[byte]bool)
	}

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == '.' {
				continue
			}
			num := board[i][j]
			boxIndex := (i/3)*3 + (j / 3)

			// 检查冲突
			if rows[i][num] || cols[j][num] || boxes[boxIndex][num] {
				return false
			}

			// 添加到集合
			rows[i][num] = true
			cols[j][num] = true
			boxes[boxIndex][num] = true
		}
	}
	return true
}

// 辅助函数：打印数独板
func printBoard(board [][]byte) {
	fmt.Println("数独板:")
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			fmt.Printf("%c ", board[i][j])
			if j == 2 || j == 5 {
				fmt.Print("| ")
			}
		}
		fmt.Println()
		if i == 2 || i == 5 {
			fmt.Println("------+-------+------")
		}
	}
	fmt.Println()
}

// 辅助函数：创建测试用例
func createTestCases() [][][]byte {
	testCases := make([][][]byte, 0)

	// 测试用例1：有效数独
	validSudoku := [][]byte{
		{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	}
	testCases = append(testCases, validSudoku)

	// 测试用例2：行冲突
	rowConflict := [][]byte{
		{'8', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	}
	testCases = append(testCases, rowConflict)

	// 测试用例3：列冲突
	colConflict := [][]byte{
		{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'5', '.', '.', '.', '8', '.', '.', '7', '9'}, // 第一列有两个5
	}
	testCases = append(testCases, colConflict)

	// 测试用例4：宫格冲突
	boxConflict := [][]byte{
		{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	}
	// 在左上角3×3宫格添加冲突
	boxConflict[0][0] = '8' // 与boxConflict[0][1]='8'冲突
	testCases = append(testCases, boxConflict)

	// 测试用例5：空数独（所有位置都是'.'）
	emptySudoku := make([][]byte, 9)
	for i := 0; i < 9; i++ {
		emptySudoku[i] = make([]byte, 9)
		for j := 0; j < 9; j++ {
			emptySudoku[i][j] = '.'
		}
	}
	testCases = append(testCases, emptySudoku)

	return testCases
}

// 性能测试函数
func benchmarkAlgorithm(algorithm func([][]byte) bool, board [][]byte, name string) {
	iterations := 10000
	start := time.Now()

	for i := 0; i < iterations; i++ {
		algorithm(board)
	}

	duration := time.Since(start)
	avgTime := duration.Nanoseconds() / int64(iterations)

	fmt.Printf("%s: 平均执行时间 %d 纳秒\n", name, avgTime)
}

func main() {
	fmt.Println("=== 36. 有效的数独 ===")
	fmt.Println()

	// 创建测试用例
	testCases := createTestCases()
	testNames := []string{
		"有效数独",
		"行冲突数独",
		"列冲突数独",
		"宫格冲突数独",
		"空数独",
	}

	// 测试所有算法
	algorithms := []struct {
		name string
		fn   func([][]byte) bool
	}{
		{"三次遍历验证法", isValidSudoku1},
		{"一次遍历优化法", isValidSudoku2},
		{"位运算优化法", isValidSudoku3},
		{"集合验证法", isValidSudoku4},
	}

	// 运行测试
	for i, testCase := range testCases {
		fmt.Printf("--- 测试用例 %d: %s ---\n", i+1, testNames[i])
		printBoard(testCase)

		for _, algo := range algorithms {
			result := algo.fn(testCase)
			fmt.Printf("%s: %t\n", algo.name, result)
		}
		fmt.Println()
	}

	// 性能测试
	fmt.Println("=== 性能测试 ===")
	performanceBoard := testCases[0] // 使用有效数独进行性能测试

	for _, algo := range algorithms {
		benchmarkAlgorithm(algo.fn, performanceBoard, algo.name)
	}

	fmt.Println()
	fmt.Println("=== 算法总结 ===")
	fmt.Println("1. 三次遍历验证法：最直观易懂，适合理解算法逻辑")
	fmt.Println("2. 一次遍历优化法：平衡了性能和代码可读性")
	fmt.Println("3. 位运算优化法：性能最佳，空间使用最少")
	fmt.Println("4. 集合验证法：代码最简洁，但性能略低")
	fmt.Println()
	fmt.Println("推荐使用：位运算优化法（方法三），在保证性能的同时代码简洁")
}
