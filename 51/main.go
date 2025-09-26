package main

import (
	"fmt"
	"math/bits"
	"sort"
	"strings"
	"time"
)

// 方法一：递归回溯算法
// 最直观的解法，使用递归回溯生成所有方案
func solveNQueens1(n int) [][]string {
	var result [][]string
	board := make([][]bool, n)
	for i := range board {
		board[i] = make([]bool, n)
	}

	backtrack(board, 0, &result)
	return result
}

// 递归回溯的辅助函数
func backtrack(board [][]bool, row int, result *[][]string) {
	n := len(board)
	if row == n {
		// 添加方案到结果集
		solution := make([]string, n)
		for i := 0; i < n; i++ {
			var rowStr strings.Builder
			for j := 0; j < n; j++ {
				if board[i][j] {
					rowStr.WriteByte('Q')
				} else {
					rowStr.WriteByte('.')
				}
			}
			solution[i] = rowStr.String()
		}
		*result = append(*result, solution)
		return
	}

	for col := 0; col < n; col++ {
		if isValid(board, row, col) {
			board[row][col] = true
			backtrack(board, row+1, result)
			board[row][col] = false
		}
	}
}

// 检查位置是否安全
func isValid(board [][]bool, row, col int) bool {
	n := len(board)

	// 检查列
	for i := 0; i < row; i++ {
		if board[i][col] {
			return false
		}
	}

	// 检查主对角线
	for i, j := row-1, col-1; i >= 0 && j >= 0; i, j = i-1, j-1 {
		if board[i][j] {
			return false
		}
	}

	// 检查副对角线
	for i, j := row-1, col+1; i >= 0 && j < n; i, j = i-1, j+1 {
		if board[i][j] {
			return false
		}
	}

	return true
}

// 方法二：位运算算法
// 使用位运算优化约束检查，效率最高
func solveNQueens2(n int) [][]string {
	var result [][]string
	var queens []int

	backtrackBitwise(n, 0, 0, 0, 0, &queens, &result)
	return result
}

// 位运算回溯的辅助函数
func backtrackBitwise(n, row, cols, diag1, diag2 int, queens *[]int, result *[][]string) {
	if row == n {
		// 添加方案到结果集
		solution := make([]string, n)
		for i := 0; i < n; i++ {
			var rowStr strings.Builder
			for j := 0; j < n; j++ {
				if (*queens)[i] == j {
					rowStr.WriteByte('Q')
				} else {
					rowStr.WriteByte('.')
				}
			}
			solution[i] = rowStr.String()
		}
		*result = append(*result, solution)
		return
	}

	available := ((1 << n) - 1) & (^(cols | diag1 | diag2))
	for available != 0 {
		pos := available & (-available)
		col := bits.TrailingZeros(uint(pos))

		*queens = append(*queens, col)
		backtrackBitwise(n, row+1, cols|pos, (diag1|pos)<<1, (diag2|pos)>>1, queens, result)
		*queens = (*queens)[:len(*queens)-1]

		available &= available - 1
	}
}

// 方法三：迭代回溯算法
// 使用栈模拟递归，避免栈溢出
func solveNQueens3(n int) [][]string {
	var result [][]string

	stack := []struct {
		board [][]bool
		row   int
	}{{make([][]bool, n), 0}}

	for i := range stack[0].board {
		stack[0].board[i] = make([]bool, n)
	}

	for len(stack) > 0 {
		current := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if current.row == n {
			// 添加方案到结果集
			solution := make([]string, n)
			for i := 0; i < n; i++ {
				var rowStr strings.Builder
				for j := 0; j < n; j++ {
					if current.board[i][j] {
						rowStr.WriteByte('Q')
					} else {
						rowStr.WriteByte('.')
					}
				}
				solution[i] = rowStr.String()
			}
			result = append(result, solution)
			continue
		}

		for col := 0; col < n; col++ {
			if isValid(current.board, current.row, col) {
				newBoard := make([][]bool, n)
				for i := range newBoard {
					newBoard[i] = make([]bool, n)
					copy(newBoard[i], current.board[i])
				}
				newBoard[current.row][col] = true

				stack = append(stack, struct {
					board [][]bool
					row   int
				}{newBoard, current.row + 1})
			}
		}
	}

	return result
}

// 方法四：启发式搜索算法
// 使用启发式函数优化搜索顺序
func solveNQueens4(n int) [][]string {
	var result [][]string
	board := make([][]bool, n)
	for i := range board {
		board[i] = make([]bool, n)
	}

	backtrackHeuristic(board, 0, &result)
	return result
}

// 启发式搜索回溯的辅助函数
func backtrackHeuristic(board [][]bool, row int, result *[][]string) {
	n := len(board)
	if row == n {
		// 添加方案到结果集
		solution := make([]string, n)
		for i := 0; i < n; i++ {
			var rowStr strings.Builder
			for j := 0; j < n; j++ {
				if board[i][j] {
					rowStr.WriteByte('Q')
				} else {
					rowStr.WriteByte('.')
				}
			}
			solution[i] = rowStr.String()
		}
		*result = append(*result, solution)
		return
	}

	// 使用启发式函数选择列
	cols := getHeuristicCols(board, row)
	for _, col := range cols {
		if isValid(board, row, col) {
			board[row][col] = true
			backtrackHeuristic(board, row+1, result)
			board[row][col] = false
		}
	}
}

// 获取启发式列
func getHeuristicCols(board [][]bool, row int) []int {
	n := len(board)
	var cols []int

	for col := 0; col < n; col++ {
		if isValid(board, row, col) {
			cols = append(cols, col)
		}
	}

	// 按启发式函数排序
	sort.Slice(cols, func(i, j int) bool {
		return getHeuristicValue(board, row, cols[i]) < getHeuristicValue(board, row, cols[j])
	})

	return cols
}

// 计算启发式值
func getHeuristicValue(board [][]bool, row, col int) int {
	// 简单的启发式函数：选择约束最少的列
	n := len(board)
	count := 0

	for i := row + 1; i < n; i++ {
		for j := 0; j < n; j++ {
			if isValid(board, i, j) {
				count++
			}
		}
	}

	return count
}

// 辅助函数：创建测试用例
func createTestCases() []struct {
	n    int
	name string
} {
	return []struct {
		n    int
		name string
	}{
		{1, "测试1: n=1"},
		{2, "测试2: n=2"},
		{3, "测试3: n=3"},
		{4, "示例1: n=4"},
		{5, "测试4: n=5"},
		{6, "测试5: n=6"},
		{7, "测试6: n=7"},
		{8, "测试7: n=8"},
		{9, "测试8: n=9"},
	}
}

// 性能测试函数
func benchmarkAlgorithm(algorithm func(int) [][]string, n int, name string) {
	iterations := 10
	start := time.Now()

	for i := 0; i < iterations; i++ {
		algorithm(n)
	}

	duration := time.Since(start)
	avgTime := duration.Nanoseconds() / int64(iterations)

	fmt.Printf("%s: 平均执行时间 %d 纳秒\n", name, avgTime)
}

// 辅助函数：验证结果是否正确
func validateResult(n int, result [][]string) bool {
	// 验证每个方案是否有效
	for _, solution := range result {
		if !isValidSolution(solution, n) {
			return false
		}
	}

	// 验证方案数量是否正确
	expectedCount := getExpectedCount(n)
	if len(result) != expectedCount {
		return false
	}

	return true
}

// 验证单个方案是否有效
func isValidSolution(solution []string, n int) bool {
	if len(solution) != n {
		return false
	}

	// 检查每行是否只有一个皇后
	for i := 0; i < n; i++ {
		if len(solution[i]) != n {
			return false
		}
		queenCount := 0
		for j := 0; j < n; j++ {
			if solution[i][j] == 'Q' {
				queenCount++
			}
		}
		if queenCount != 1 {
			return false
		}
	}

	// 检查皇后之间是否相互攻击
	queens := make([][2]int, 0, n)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if solution[i][j] == 'Q' {
				queens = append(queens, [2]int{i, j})
			}
		}
	}

	for i := 0; i < len(queens); i++ {
		for j := i + 1; j < len(queens); j++ {
			if queens[i][0] == queens[j][0] || // 同一行
				queens[i][1] == queens[j][1] || // 同一列
				abs(queens[i][0]-queens[j][0]) == abs(queens[i][1]-queens[j][1]) { // 同一对角线
				return false
			}
		}
	}

	return true
}

// 计算绝对值
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// 获取期望的解的数量
func getExpectedCount(n int) int {
	switch n {
	case 1:
		return 1
	case 2, 3:
		return 0
	case 4:
		return 2
	case 5:
		return 10
	case 6:
		return 4
	case 7:
		return 40
	case 8:
		return 92
	case 9:
		return 352
	default:
		return 0
	}
}

// 辅助函数：比较两个结果是否相同
func compareResults(result1, result2 [][]string) bool {
	if len(result1) != len(result2) {
		return false
	}

	// 将结果转换为可比较的格式
	normalize := func(result [][]string) []string {
		var normalized []string
		for _, solution := range result {
			normalized = append(normalized, strings.Join(solution, "|"))
		}
		sort.Strings(normalized)
		return normalized
	}

	norm1 := normalize(result1)
	norm2 := normalize(result2)

	for i := range norm1 {
		if norm1[i] != norm2[i] {
			return false
		}
	}

	return true
}

// 辅助函数：打印棋盘结果
func printBoardResult(n int, result [][]string, title string) {
	fmt.Printf("%s: n=%d -> %d 个解\n", title, n, len(result))
	if len(result) <= 2 {
		for i, solution := range result {
			fmt.Printf("  解 %d:\n", i+1)
			for _, row := range solution {
				fmt.Printf("    %s\n", row)
			}
		}
	}
}

func main() {
	fmt.Println("=== 51. N 皇后 ===")
	fmt.Println()

	// 创建测试用例
	testCases := createTestCases()
	algorithms := []struct {
		name string
		fn   func(int) [][]string
	}{
		{"递归回溯算法", solveNQueens1},
		{"位运算算法", solveNQueens2},
		{"迭代回溯算法", solveNQueens3},
		{"启发式搜索算法", solveNQueens4},
	}

	// 运行测试
	fmt.Println("=== 算法正确性测试 ===")
	for _, testCase := range testCases {
		fmt.Printf("测试: %s\n", testCase.name)

		results := make([][][]string, len(algorithms))
		for i, algo := range algorithms {
			results[i] = algo.fn(testCase.n)
		}

		// 验证所有算法结果一致
		allEqual := true
		for i := 1; i < len(results); i++ {
			if !compareResults(results[i], results[0]) {
				allEqual = false
				break
			}
		}

		// 验证结果是否正确
		allValid := true
		for _, result := range results {
			if !validateResult(testCase.n, result) {
				allValid = false
				break
			}
		}

		if allEqual && allValid {
			fmt.Printf("  ✅ 所有算法结果一致且正确: %d 个解\n", len(results[0]))
			if testCase.n <= 4 {
				printBoardResult(testCase.n, results[0], "  棋盘结果")
			}
		} else {
			fmt.Printf("  ❌ 算法结果不一致或错误\n")
			for i, algo := range algorithms {
				fmt.Printf("    %s: %d 个解\n", algo.name, len(results[i]))
			}
		}
		fmt.Println()
	}

	// 性能测试
	fmt.Println("=== 性能测试 ===")
	performanceN := 8

	fmt.Printf("测试数据: n=%d\n", performanceN)
	fmt.Println()

	for _, algo := range algorithms {
		benchmarkAlgorithm(algo.fn, performanceN, algo.name)
	}
	fmt.Println()

	// 算法分析
	fmt.Println("=== 算法分析 ===")
	fmt.Println("N皇后问题的特点:")
	fmt.Println("1. 需要将n个皇后放置在n×n的棋盘上")
	fmt.Println("2. 皇后之间不能相互攻击")
	fmt.Println("3. 需要找到所有可能的放置方案")
	fmt.Println("4. 回溯算法是最优解法")
	fmt.Println()

	// 复杂度分析
	fmt.Println("=== 复杂度分析 ===")
	fmt.Println("时间复杂度:")
	fmt.Println("- 递归回溯: O(n!)，需要尝试所有可能的放置方案")
	fmt.Println("- 位运算: O(n!)，使用位运算优化但时间复杂度不变")
	fmt.Println("- 迭代回溯: O(n!)，使用栈模拟递归，时间复杂度相同")
	fmt.Println("- 启发式搜索: O(n!)，使用启发式函数但时间复杂度不变")
	fmt.Println()
	fmt.Println("空间复杂度:")
	fmt.Println("- 递归栈: O(n)，递归深度最多为n")
	fmt.Println("- 位运算: O(1)，只使用常数空间")
	fmt.Println("- 迭代栈: O(n)，栈的最大深度为n")
	fmt.Println("- 启发式搜索: O(n)，需要存储启发式信息")
	fmt.Println()

	// 算法总结
	fmt.Println("=== 算法总结 ===")
	fmt.Println("1. 递归回溯算法：最直观的解法，逻辑清晰")
	fmt.Println("2. 位运算算法：使用位运算优化，效率最高")
	fmt.Println("3. 迭代回溯算法：使用栈模拟递归，避免栈溢出")
	fmt.Println("4. 启发式搜索算法：使用启发式函数优化搜索顺序")
	fmt.Println()
	fmt.Println("推荐使用：位运算算法（方法二），效率最高")
	fmt.Println()

	// 应用场景
	fmt.Println("=== 应用场景 ===")
	fmt.Println("- 算法竞赛：回溯算法的经典应用")
	fmt.Println("- 人工智能：约束满足问题")
	fmt.Println("- 游戏开发：棋盘游戏逻辑")
	fmt.Println("- 数学研究：组合数学问题")
	fmt.Println("- 教学演示：算法教学案例")
	fmt.Println()

	// 优化技巧总结
	fmt.Println("=== 优化技巧总结 ===")
	fmt.Println("1. 递归回溯：掌握递归回溯的核心思想")
	fmt.Println("2. 约束检查：学会高效地检查约束条件")
	fmt.Println("3. 位运算：学会使用位运算优化")
	fmt.Println("4. 剪枝优化：学会避免无效的搜索分支")
	fmt.Println("5. 算法选择：根据问题特点选择合适的算法")
	fmt.Println("6. 优化策略：学会时间和空间优化技巧")
}
