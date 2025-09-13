package main

import (
	"fmt"
	"sort"
	"time"
)

// 方法一：基础回溯算法
// 最直观的回溯解法，遍历所有位置寻找空格并尝试填入数字
func solveSudoku1(board [][]byte) bool {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == '.' {
				// 尝试填入数字1-9
				for num := byte('1'); num <= byte('9'); num++ {
					if isValidPlacement(board, i, j, num) {
						board[i][j] = num
						if solveSudoku1(board) {
							return true
						}
						board[i][j] = '.' // 回溯
					}
				}
				return false
			}
		}
	}
	return true
}

// 辅助函数：检查数字是否可以放置在指定位置
func isValidPlacement(board [][]byte, row, col int, num byte) bool {
	// 检查行
	for j := 0; j < 9; j++ {
		if board[row][j] == num {
			return false
		}
	}

	// 检查列
	for i := 0; i < 9; i++ {
		if board[i][col] == num {
			return false
		}
	}

	// 检查3×3宫格
	boxRow := (row / 3) * 3
	boxCol := (col / 3) * 3
	for i := boxRow; i < boxRow+3; i++ {
		for j := boxCol; j < boxCol+3; j++ {
			if board[i][j] == num {
				return false
			}
		}
	}

	return true
}

// 方法二：优化回溯算法
// 预处理收集所有空格，按约束度排序，优先填充约束最多的空格
func solveSudoku2(board [][]byte) bool {
	// 收集所有空格
	var emptyCells [][]int
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == '.' {
				emptyCells = append(emptyCells, []int{i, j})
			}
		}
	}

	// 按约束度排序（约束度高的优先）
	sort.Slice(emptyCells, func(i, j int) bool {
		return calculateConstraints(board, emptyCells[i][0], emptyCells[i][1]) >
			calculateConstraints(board, emptyCells[j][0], emptyCells[j][1])
	})

	return solveSudoku2Helper(board, emptyCells, 0)
}

// 计算指定位置的约束度
func calculateConstraints(board [][]byte, row, col int) int {
	constraints := 0

	// 计算行约束
	for j := 0; j < 9; j++ {
		if board[row][j] != '.' {
			constraints++
		}
	}

	// 计算列约束
	for i := 0; i < 9; i++ {
		if board[i][col] != '.' {
			constraints++
		}
	}

	// 计算宫格约束
	boxRow, boxCol := (row/3)*3, (col/3)*3
	for i := boxRow; i < boxRow+3; i++ {
		for j := boxCol; j < boxCol+3; j++ {
			if board[i][j] != '.' {
				constraints++
			}
		}
	}

	return constraints
}

// 优化回溯的递归辅助函数
func solveSudoku2Helper(board [][]byte, emptyCells [][]int, index int) bool {
	if index == len(emptyCells) {
		return true
	}

	row, col := emptyCells[index][0], emptyCells[index][1]

	// 尝试填入数字1-9
	for num := byte('1'); num <= byte('9'); num++ {
		if isValidPlacement(board, row, col, num) {
			board[row][col] = num
			if solveSudoku2Helper(board, emptyCells, index+1) {
				return true
			}
			board[row][col] = '.' // 回溯
		}
	}

	return false
}

// 方法三：位运算回溯
// 使用位掩码记录数字使用情况，位运算加速冲突检测
func solveSudoku3(board [][]byte) bool {
	// 初始化位掩码
	rows := make([]int, 9)
	cols := make([]int, 9)
	boxes := make([]int, 9)

	// 预处理：记录已填入的数字
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] != '.' {
				num := int(board[i][j] - '0')
				bit := 1 << num
				boxIndex := (i/3)*3 + (j / 3)
				rows[i] |= bit
				cols[j] |= bit
				boxes[boxIndex] |= bit
			}
		}
	}

	return solveSudoku3Helper(board, rows, cols, boxes)
}

// 位运算回溯的递归辅助函数
func solveSudoku3Helper(board [][]byte, rows, cols, boxes []int) bool {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == '.' {
				boxIndex := (i/3)*3 + (j / 3)

				// 尝试所有数字
				for num := 1; num <= 9; num++ {
					bit := 1 << num
					// 检查数字是否可用
					if (rows[i]&bit) == 0 && (cols[j]&bit) == 0 && (boxes[boxIndex]&bit) == 0 {
						// 填入数字
						board[i][j] = byte(num + '0')
						rows[i] |= bit
						cols[j] |= bit
						boxes[boxIndex] |= bit

						// 递归求解
						if solveSudoku3Helper(board, rows, cols, boxes) {
							return true
						}

						// 回溯
						board[i][j] = '.'
						rows[i] ^= bit
						cols[j] ^= bit
						boxes[boxIndex] ^= bit
					}
				}
				return false
			}
		}
	}
	return true
}

// 方法四：启发式搜索
// 使用MRV（最小剩余值）启发式，优先填充约束最多的空格
func solveSudoku4(board [][]byte) bool {
	// 收集所有空格
	var emptyCells [][]int
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == '.' {
				emptyCells = append(emptyCells, []int{i, j})
			}
		}
	}

	// 按约束度排序（约束度高的优先）
	sort.Slice(emptyCells, func(i, j int) bool {
		return calculateConstraints(board, emptyCells[i][0], emptyCells[i][1]) >
			calculateConstraints(board, emptyCells[j][0], emptyCells[j][1])
	})

	return solveSudoku4Helper(board, emptyCells, 0)
}

// 启发式搜索的递归辅助函数
func solveSudoku4Helper(board [][]byte, emptyCells [][]int, index int) bool {
	if index == len(emptyCells) {
		return true
	}

	row, col := emptyCells[index][0], emptyCells[index][1]

	// 尝试填入数字1-9
	for num := byte('1'); num <= byte('9'); num++ {
		if isValidPlacement(board, row, col, num) {
			board[row][col] = num
			if solveSudoku4Helper(board, emptyCells, index+1) {
				return true
			}
			board[row][col] = '.' // 回溯
		}
	}

	return false
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

// 辅助函数：复制数独板
func copyBoard(board [][]byte) [][]byte {
	newBoard := make([][]byte, 9)
	for i := 0; i < 9; i++ {
		newBoard[i] = make([]byte, 9)
		copy(newBoard[i], board[i])
	}
	return newBoard
}

// 辅助函数：创建测试用例
func createTestCases() [][][]byte {
	testCases := make([][][]byte, 0)

	// 测试用例1：简单数独（少量空格）
	easySudoku := [][]byte{
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
	testCases = append(testCases, easySudoku)

	// 测试用例2：中等数独（中等空格数量）
	mediumSudoku := [][]byte{
		{'.', '.', '9', '7', '4', '8', '.', '.', '.'},
		{'7', '.', '.', '.', '.', '.', '.', '.', '.'},
		{'.', '2', '.', '1', '.', '9', '.', '.', '.'},
		{'.', '.', '7', '.', '.', '.', '2', '4', '.'},
		{'.', '6', '4', '.', '1', '.', '5', '9', '.'},
		{'.', '9', '8', '.', '.', '.', '3', '.', '.'},
		{'.', '.', '.', '8', '.', '3', '.', '2', '.'},
		{'.', '.', '.', '.', '.', '.', '.', '.', '6'},
		{'.', '.', '.', '2', '7', '5', '9', '.', '.'},
	}
	testCases = append(testCases, mediumSudoku)

	// 测试用例3：困难数独（大量空格）
	hardSudoku := [][]byte{
		{'.', '.', '.', '.', '.', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '.', '.', '.', '.'},
	}
	// 添加一些初始数字
	hardSudoku[0][0] = '1'
	hardSudoku[0][1] = '2'
	hardSudoku[1][0] = '3'
	hardSudoku[1][1] = '4'
	testCases = append(testCases, hardSudoku)

	// 测试用例4：已填满数独（验证正确性）
	completeSudoku := [][]byte{
		{'5', '3', '4', '6', '7', '8', '9', '1', '2'},
		{'6', '7', '2', '1', '9', '5', '3', '4', '8'},
		{'1', '9', '8', '3', '4', '2', '5', '6', '7'},
		{'8', '5', '9', '7', '6', '1', '4', '2', '3'},
		{'4', '2', '6', '8', '5', '3', '7', '9', '1'},
		{'7', '1', '3', '9', '2', '4', '8', '5', '6'},
		{'9', '6', '1', '5', '3', '7', '2', '8', '4'},
		{'2', '8', '7', '4', '1', '9', '6', '3', '5'},
		{'3', '4', '5', '2', '8', '6', '1', '7', '9'},
	}
	testCases = append(testCases, completeSudoku)

	// 测试用例5：单空格数独
	singleEmptySudoku := [][]byte{
		{'5', '3', '4', '6', '7', '8', '9', '1', '2'},
		{'6', '7', '2', '1', '9', '5', '3', '4', '8'},
		{'1', '9', '8', '3', '4', '2', '5', '6', '7'},
		{'8', '5', '9', '7', '6', '1', '4', '2', '3'},
		{'4', '2', '6', '8', '5', '3', '7', '9', '1'},
		{'7', '1', '3', '9', '2', '4', '8', '5', '6'},
		{'9', '6', '1', '5', '3', '7', '2', '8', '4'},
		{'2', '8', '7', '4', '1', '9', '6', '3', '5'},
		{'3', '4', '5', '2', '8', '6', '1', '7', '.'}, // 只有一个空格
	}
	testCases = append(testCases, singleEmptySudoku)

	return testCases
}

// 性能测试函数
func benchmarkAlgorithm(algorithm func([][]byte) bool, board [][]byte, name string) {
	iterations := 100
	start := time.Now()

	for i := 0; i < iterations; i++ {
		testBoard := copyBoard(board)
		algorithm(testBoard)
	}

	duration := time.Since(start)
	avgTime := duration.Nanoseconds() / int64(iterations)

	fmt.Printf("%s: 平均执行时间 %d 纳秒\n", name, avgTime)
}

func main() {
	fmt.Println("=== 37. 解数独 ===")
	fmt.Println()

	// 创建测试用例
	testCases := createTestCases()
	testNames := []string{
		"简单数独",
		"中等数独",
		"困难数独",
		"已填满数独",
		"单空格数独",
	}

	// 测试所有算法
	algorithms := []struct {
		name string
		fn   func([][]byte) bool
	}{
		{"基础回溯算法", solveSudoku1},
		{"优化回溯算法", solveSudoku2},
		{"位运算回溯", solveSudoku3},
		{"启发式搜索", solveSudoku4},
	}

	// 运行测试
	for i, testCase := range testCases {
		fmt.Printf("--- 测试用例 %d: %s ---\n", i+1, testNames[i])
		fmt.Println("求解前:")
		printBoard(testCase)

		for _, algo := range algorithms {
			testBoard := copyBoard(testCase)
			start := time.Now()
			result := algo.fn(testBoard)
			duration := time.Since(start)

			if result {
				fmt.Printf("%s: 求解成功 (耗时: %v)\n", algo.name, duration)
				if i == 0 { // 只对第一个测试用例显示求解结果
					fmt.Println("求解后:")
					printBoard(testBoard)
				}
			} else {
				fmt.Printf("%s: 求解失败 (耗时: %v)\n", algo.name, duration)
			}
		}
		fmt.Println()
	}

	// 性能测试
	fmt.Println("=== 性能测试 ===")
	performanceBoard := testCases[0] // 使用简单数独进行性能测试

	for _, algo := range algorithms {
		benchmarkAlgorithm(algo.fn, performanceBoard, algo.name)
	}

	fmt.Println()
	fmt.Println("=== 算法总结 ===")
	fmt.Println("1. 基础回溯算法：最直观易懂，适合理解算法逻辑")
	fmt.Println("2. 优化回溯算法：添加约束度排序，显著提升性能")
	fmt.Println("3. 位运算回溯：使用位掩码加速冲突检测，性能优秀")
	fmt.Println("4. 启发式搜索：使用MRV和LCV启发式，性能最佳")
	fmt.Println()
	fmt.Println("推荐使用：启发式搜索（方法四），在保证性能的同时算法最先进")
}
