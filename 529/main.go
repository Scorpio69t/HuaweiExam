package main

import (
	"fmt"
	"strings"
	"time"
)

// 八个方向的偏移量：上下左右和四个对角线
var directions = [][]int{
	{-1, -1}, {-1, 0}, {-1, 1},
	{0, -1}, {0, 1},
	{1, -1}, {1, 0}, {1, 1},
}

// 解法一：DFS递归（推荐解法）
// 时间复杂度：O(m×n)，空间复杂度：O(m×n)
func updateBoard(board [][]byte, click []int) [][]byte {
	if len(board) == 0 || len(board[0]) == 0 {
		return board
	}

	// 创建副本避免修改原数组
	result := copyBoard(board)
	row, col := click[0], click[1]

	// 如果点击的是地雷，游戏结束
	if result[row][col] == 'M' {
		result[row][col] = 'X'
		return result
	}

	// 开始DFS
	dfs(result, row, col)
	return result
}

// DFS递归函数
func dfs(board [][]byte, i, j int) {
	// 边界检查和状态检查
	if !inBounds(i, j, len(board), len(board[0])) || board[i][j] != 'E' {
		return
	}

	// 计算周围地雷数量
	mineCount := countMines(board, i, j)

	if mineCount > 0 {
		// 有相邻地雷，标记数字并停止扩展
		board[i][j] = byte('0' + mineCount)
	} else {
		// 没有相邻地雷，标记为'B'并继续扩展
		board[i][j] = 'B'

		// 递归处理8个相邻方向
		for _, dir := range directions {
			ni, nj := i+dir[0], j+dir[1]
			dfs(board, ni, nj)
		}
	}
}

// 解法二：DFS迭代（避免栈溢出）
// 时间复杂度：O(m×n)，空间复杂度：O(m×n)
func updateBoardIterative(board [][]byte, click []int) [][]byte {
	if len(board) == 0 || len(board[0]) == 0 {
		return board
	}

	result := copyBoard(board)
	row, col := click[0], click[1]

	// 如果点击的是地雷
	if result[row][col] == 'M' {
		result[row][col] = 'X'
		return result
	}

	// 使用栈进行迭代DFS
	stack := [][]int{{row, col}}

	for len(stack) > 0 {
		// 弹出栈顶元素
		pos := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		i, j := pos[0], pos[1]

		// 检查边界和状态
		if !inBounds(i, j, len(result), len(result[0])) || result[i][j] != 'E' {
			continue
		}

		// 计算相邻地雷数
		mineCount := countMines(result, i, j)

		if mineCount > 0 {
			result[i][j] = byte('0' + mineCount)
		} else {
			result[i][j] = 'B'

			// 将8个相邻位置加入栈
			for _, dir := range directions {
				ni, nj := i+dir[0], j+dir[1]
				stack = append(stack, []int{ni, nj})
			}
		}
	}

	return result
}

// 解法三：BFS（广度优先搜索）
// 时间复杂度：O(m×n)，空间复杂度：O(m×n)
func updateBoardBFS(board [][]byte, click []int) [][]byte {
	if len(board) == 0 || len(board[0]) == 0 {
		return board
	}

	result := copyBoard(board)
	row, col := click[0], click[1]

	// 如果点击的是地雷
	if result[row][col] == 'M' {
		result[row][col] = 'X'
		return result
	}

	// BFS队列
	queue := [][]int{{row, col}}
	visited := make(map[string]bool)

	for len(queue) > 0 {
		pos := queue[0]
		queue = queue[1:]
		i, j := pos[0], pos[1]

		key := fmt.Sprintf("%d,%d", i, j)
		if visited[key] {
			continue
		}
		visited[key] = true

		// 检查边界和状态
		if !inBounds(i, j, len(result), len(result[0])) || result[i][j] != 'E' {
			continue
		}

		// 计算相邻地雷数
		mineCount := countMines(result, i, j)

		if mineCount > 0 {
			result[i][j] = byte('0' + mineCount)
		} else {
			result[i][j] = 'B'

			// 将相邻的'E'格子加入队列
			for _, dir := range directions {
				ni, nj := i+dir[0], j+dir[1]
				if inBounds(ni, nj, len(result), len(result[0])) && result[ni][nj] == 'E' {
					nextKey := fmt.Sprintf("%d,%d", ni, nj)
					if !visited[nextKey] {
						queue = append(queue, []int{ni, nj})
					}
				}
			}
		}
	}

	return result
}

// 解法四：优化DFS（减少重复计算）
// 时间复杂度：O(m×n)，空间复杂度：O(m×n)
func updateBoardOptimized(board [][]byte, click []int) [][]byte {
	if len(board) == 0 || len(board[0]) == 0 {
		return board
	}

	result := copyBoard(board)
	row, col := click[0], click[1]

	// 如果点击的是地雷
	if result[row][col] == 'M' {
		result[row][col] = 'X'
		return result
	}

	// 预计算所有位置的地雷数（可选优化）
	dfsOptimized(result, row, col, make(map[string]bool))
	return result
}

// 优化的DFS，使用visited避免重复访问
func dfsOptimized(board [][]byte, i, j int, visited map[string]bool) {
	key := fmt.Sprintf("%d,%d", i, j)
	if visited[key] || !inBounds(i, j, len(board), len(board[0])) || board[i][j] != 'E' {
		return
	}

	visited[key] = true
	mineCount := countMines(board, i, j)

	if mineCount > 0 {
		board[i][j] = byte('0' + mineCount)
	} else {
		board[i][j] = 'B'

		for _, dir := range directions {
			ni, nj := i+dir[0], j+dir[1]
			dfsOptimized(board, ni, nj, visited)
		}
	}
}

// 辅助函数：检查坐标是否在边界内
func inBounds(i, j, m, n int) bool {
	return i >= 0 && i < m && j >= 0 && j < n
}

// 辅助函数：计算周围地雷数量
func countMines(board [][]byte, i, j int) int {
	count := 0
	for _, dir := range directions {
		ni, nj := i+dir[0], j+dir[1]
		if inBounds(ni, nj, len(board), len(board[0])) && board[ni][nj] == 'M' {
			count++
		}
	}
	return count
}

// 辅助函数：复制棋盘
func copyBoard(board [][]byte) [][]byte {
	result := make([][]byte, len(board))
	for i := range board {
		result[i] = make([]byte, len(board[i]))
		copy(result[i], board[i])
	}
	return result
}

// 辅助函数：打印棋盘
func printBoard(board [][]byte) {
	for _, row := range board {
		fmt.Print("[")
		for j, cell := range row {
			if j > 0 {
				fmt.Print(",")
			}
			fmt.Printf("\"%c\"", cell)
		}
		fmt.Print("]")
		fmt.Println()
	}
}

// 辅助函数：比较两个棋盘是否相等
func equalBoards(a, b [][]byte) bool {
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

// 扫雷游戏模拟器
type Minesweeper struct {
	board    [][]byte
	original [][]byte
	rows     int
	cols     int
	gameOver bool
	moves    int
}

// 创建新的扫雷游戏
func newMinesweeper(board [][]byte) *Minesweeper {
	ms := &Minesweeper{
		board:    copyBoard(board),
		original: copyBoard(board),
		rows:     len(board),
		cols:     len(board[0]),
		gameOver: false,
		moves:    0,
	}
	return ms
}

// 执行点击操作
func (ms *Minesweeper) click(row, col int) bool {
	if ms.gameOver || !inBounds(row, col, ms.rows, ms.cols) {
		return false
	}

	// 只能点击未挖出的格子
	if ms.board[row][col] != 'E' && ms.board[row][col] != 'M' {
		return false
	}

	ms.moves++

	// 点击地雷
	if ms.board[row][col] == 'M' {
		ms.board[row][col] = 'X'
		ms.gameOver = true
		ms.revealAllMines()
		return false
	}

	// 正常点击
	dfs(ms.board, row, col)
	return true
}

// 显示所有地雷
func (ms *Minesweeper) revealAllMines() {
	for i := 0; i < ms.rows; i++ {
		for j := 0; j < ms.cols; j++ {
			if ms.original[i][j] == 'M' {
				ms.board[i][j] = 'M'
			}
		}
	}
}

// 检查是否获胜
func (ms *Minesweeper) isWin() bool {
	for i := 0; i < ms.rows; i++ {
		for j := 0; j < ms.cols; j++ {
			// 如果有未挖出的非地雷格子，游戏未结束
			if ms.original[i][j] != 'M' && ms.board[i][j] == 'E' {
				return false
			}
		}
	}
	return true
}

// 获取游戏状态
func (ms *Minesweeper) getStatus() string {
	if ms.gameOver {
		return "Game Over"
	}
	if ms.isWin() {
		return "You Win!"
	}
	return "Playing"
}

// 测试函数
func testMinesweeper() {
	testCases := []struct {
		board    [][]byte
		click    []int
		expected [][]byte
		desc     string
	}{
		{
			[][]byte{
				{'E', 'E', 'E', 'E', 'E'},
				{'E', 'E', 'M', 'E', 'E'},
				{'E', 'E', 'E', 'E', 'E'},
				{'E', 'E', 'E', 'E', 'E'},
			},
			[]int{3, 0},
			[][]byte{
				{'B', '1', 'E', '1', 'B'},
				{'B', '1', 'M', '1', 'B'},
				{'B', '1', '1', '1', 'B'},
				{'B', 'B', 'B', 'B', 'B'},
			},
			"示例1：点击空白区域自动展开",
		},
		{
			[][]byte{
				{'B', '1', 'E', '1', 'B'},
				{'B', '1', 'M', '1', 'B'},
				{'B', '1', '1', '1', 'B'},
				{'B', 'B', 'B', 'B', 'B'},
			},
			[]int{1, 2},
			[][]byte{
				{'B', '1', 'E', '1', 'B'},
				{'B', '1', 'X', '1', 'B'},
				{'B', '1', '1', '1', 'B'},
				{'B', 'B', 'B', 'B', 'B'},
			},
			"示例2：点击地雷游戏结束",
		},
		{
			[][]byte{
				{'E', 'E', 'E'},
				{'E', 'M', 'E'},
				{'E', 'E', 'E'},
			},
			[]int{0, 0},
			[][]byte{
				{'1', 'E', 'E'},
				{'E', 'M', 'E'},
				{'E', 'E', 'E'},
			},
			"中心地雷：点击角落",
		},
		{
			[][]byte{
				{'M', 'E'},
				{'E', 'E'},
			},
			[]int{1, 1},
			[][]byte{
				{'M', 'E'},
				{'E', '1'},
			},
			"小棋盘测试",
		},
		{
			[][]byte{
				{'E', 'E', 'E'},
				{'E', 'E', 'E'},
				{'E', 'E', 'E'},
			},
			[]int{1, 1},
			[][]byte{
				{'B', 'B', 'B'},
				{'B', 'B', 'B'},
				{'B', 'B', 'B'},
			},
			"无地雷棋盘：全部展开",
		},
		{
			[][]byte{
				{'M', 'M'},
				{'M', 'M'},
			},
			[]int{0, 0},
			[][]byte{
				{'X', 'M'},
				{'M', 'M'},
			},
			"全地雷棋盘：点击任意位置",
		},
		{
			[][]byte{
				{'E', 'M', 'E'},
				{'M', 'E', 'M'},
				{'E', 'M', 'E'},
			},
			[]int{1, 1},
			[][]byte{
				{'E', 'M', 'E'},
				{'M', '4', 'M'},
				{'E', 'M', 'E'},
			},
			"中心surrounded by mines",
		},
	}

	fmt.Println("=== 扫雷游戏测试 ===")
	fmt.Println()

	for i, tc := range testCases {
		// 测试不同算法
		result1 := updateBoard(tc.board, tc.click)
		result2 := updateBoardIterative(tc.board, tc.click)
		result3 := updateBoardBFS(tc.board, tc.click)
		result4 := updateBoardOptimized(tc.board, tc.click)

		status := "✅"
		if !equalBoards(result1, tc.expected) {
			status = "❌"
		}

		fmt.Printf("测试 %d: %s\n", i+1, tc.desc)
		fmt.Print("输入棋盘: ")
		printBoard(tc.board)
		fmt.Printf("点击位置: [%d,%d]\n", tc.click[0], tc.click[1])
		fmt.Print("期望结果: ")
		printBoard(tc.expected)
		fmt.Print("DFS递归: ")
		printBoard(result1)

		// 验证算法一致性
		consistent := equalBoards(result1, result2) &&
			equalBoards(result2, result3) &&
			equalBoards(result3, result4)
		fmt.Printf("算法一致性: %t\n", consistent)
		fmt.Printf("结果: %s\n", status)
		fmt.Println(strings.Repeat("-", 50))
	}
}

// 性能测试
func benchmarkMinesweeper() {
	fmt.Println()
	fmt.Println("=== 性能测试 ===")
	fmt.Println()

	// 构造测试数据
	testData := []struct {
		board [][]byte
		click []int
		desc  string
	}{
		{
			generateBoard(10, 10, 0.1),
			[]int{0, 0},
			"10×10棋盘，10%地雷",
		},
		{
			generateBoard(20, 20, 0.15),
			[]int{10, 10},
			"20×20棋盘，15%地雷",
		},
		{
			generateBoard(30, 30, 0.2),
			[]int{15, 15},
			"30×30棋盘，20%地雷",
		},
		{
			generateAllEmpty(50, 50),
			[]int{25, 25},
			"50×50全空棋盘（最坏情况）",
		},
	}

	algorithms := []struct {
		name string
		fn   func([][]byte, []int) [][]byte
	}{
		{"DFS递归", updateBoard},
		{"DFS迭代", updateBoardIterative},
		{"BFS", updateBoardBFS},
		{"优化DFS", updateBoardOptimized},
	}

	for _, data := range testData {
		fmt.Printf("%s:\n", data.desc)

		for _, algo := range algorithms {
			start := time.Now()
			result := algo.fn(data.board, data.click)
			duration := time.Since(start)

			revealed := countRevealed(result)
			fmt.Printf("  %s: 展开%d格，耗时 %v\n", algo.name, revealed, duration)
		}
		fmt.Println()
	}
}

// 生成测试棋盘
func generateBoard(rows, cols int, mineRatio float64) [][]byte {
	board := make([][]byte, rows)
	for i := range board {
		board[i] = make([]byte, cols)
		for j := range board[i] {
			// 根据比例随机生成地雷
			if float64((i*cols+j)%100)/100 < mineRatio {
				board[i][j] = 'M'
			} else {
				board[i][j] = 'E'
			}
		}
	}
	return board
}

// 生成全空棋盘
func generateAllEmpty(rows, cols int) [][]byte {
	board := make([][]byte, rows)
	for i := range board {
		board[i] = make([]byte, cols)
		for j := range board[i] {
			board[i][j] = 'E'
		}
	}
	return board
}

// 统计已揭露格子数
func countRevealed(board [][]byte) int {
	count := 0
	for i := range board {
		for j := range board[i] {
			if board[i][j] != 'E' && board[i][j] != 'M' {
				count++
			}
		}
	}
	return count
}

// 交互式扫雷游戏演示
func demonstrateGame() {
	fmt.Println()
	fmt.Println("=== 扫雷游戏过程演示 ===")

	// 创建一个简单的扫雷游戏
	board := [][]byte{
		{'E', 'E', 'E', 'E'},
		{'E', 'M', 'E', 'E'},
		{'E', 'E', 'E', 'M'},
		{'E', 'E', 'E', 'E'},
	}

	game := newMinesweeper(board)

	fmt.Println("初始棋盘:")
	printBoard(game.original)

	// 模拟几步操作
	moves := [][]int{
		{0, 0}, // 安全位置
		{3, 3}, // 角落位置
		{1, 0}, // 地雷附近
	}

	for i, move := range moves {
		fmt.Printf("\n第%d步: 点击位置 [%d,%d]\n", i+1, move[0], move[1])

		success := game.click(move[0], move[1])
		fmt.Printf("操作结果: %s\n", game.getStatus())

		fmt.Println("棋盘状态:")
		printBoard(game.board)

		if !success {
			fmt.Println("游戏结束!")
			break
		}
	}
}

// 算法比较演示
func demonstrateAlgorithms() {
	fmt.Println()
	fmt.Println("=== 算法实现对比 ===")

	board := [][]byte{
		{'E', 'E', 'M'},
		{'E', 'E', 'E'},
		{'M', 'E', 'E'},
	}

	click := []int{1, 1}

	fmt.Println("测试棋盘:")
	printBoard(board)
	fmt.Printf("点击位置: [%d,%d]\n", click[0], click[1])

	algorithms := []struct {
		name string
		fn   func([][]byte, []int) [][]byte
		desc string
	}{
		{"DFS递归", updateBoard, "简洁直观，可能栈溢出"},
		{"DFS迭代", updateBoardIterative, "避免栈溢出，手动管理栈"},
		{"BFS", updateBoardBFS, "层次遍历，适合可视化"},
		{"优化DFS", updateBoardOptimized, "减少重复访问"},
	}

	for _, algo := range algorithms {
		fmt.Printf("\n%s (%s):\n", algo.name, algo.desc)
		result := algo.fn(board, click)
		printBoard(result)
	}
}

func main() {
	fmt.Println("529. 扫雷游戏 - 多种解法实现")
	fmt.Println("==============================")

	// 基础功能测试
	testMinesweeper()

	// 性能对比测试
	benchmarkMinesweeper()

	// 游戏过程演示
	demonstrateGame()

	// 算法对比演示
	demonstrateAlgorithms()

	// 展示算法特点
	fmt.Println()
	fmt.Println("=== 算法特点分析 ===")
	fmt.Println("1. DFS递归：代码简洁，逻辑清晰，适合小规模棋盘")
	fmt.Println("2. DFS迭代：避免栈溢出，适合大规模棋盘")
	fmt.Println("3. BFS：层次遍历，便于可视化展开过程")
	fmt.Println("4. 优化DFS：减少重复访问，提升性能")

	fmt.Println()
	fmt.Println("=== 扫雷游戏技巧 ===")
	fmt.Println("• 八方向遍历：预定义方向数组简化代码")
	fmt.Println("• 边界检查：统一的边界检查函数")
	fmt.Println("• 状态管理：明确每种字符的含义和转换")
	fmt.Println("• 递归控制：合理选择递归深度和终止条件")
	fmt.Println("• 游戏逻辑：严格按照扫雷规则实现")
	fmt.Println("• 性能优化：根据棋盘特点选择最优算法")
}
