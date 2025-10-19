package main

import (
	"fmt"
)

// =========================== 方法一：DFS回溯（最优解法） ===========================

func exist(board [][]byte, word string) bool {
	if len(board) == 0 || len(board[0]) == 0 {
		return false
	}

	m, n := len(board), len(board[0])

	// 方向数组：上、右、下、左
	dirs := [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}

	var dfs func(row, col, index int, visited [][]bool) bool
	dfs = func(row, col, index int, visited [][]bool) bool {
		// 边界检查
		if row < 0 || row >= m || col < 0 || col >= n {
			return false
		}

		// 已访问检查
		if visited[row][col] {
			return false
		}

		// 字符匹配检查
		if board[row][col] != word[index] {
			return false
		}

		// 到达单词末尾
		if index == len(word)-1 {
			return true
		}

		// 标记已访问
		visited[row][col] = true

		// 尝试四个方向
		for _, dir := range dirs {
			newRow, newCol := row+dir[0], col+dir[1]
			if dfs(newRow, newCol, index+1, visited) {
				return true
			}
		}

		// 撤销标记
		visited[row][col] = false
		return false
	}

	// 尝试所有起始位置
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			visited := make([][]bool, m)
			for k := range visited {
				visited[k] = make([]bool, n)
			}
			if dfs(i, j, 0, visited) {
				return true
			}
		}
	}

	return false
}

// =========================== 方法二：优化版DFS（原地标记） ===========================

func exist2(board [][]byte, word string) bool {
	if len(board) == 0 || len(board[0]) == 0 {
		return false
	}

	m, n := len(board), len(board[0])
	dirs := [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}

	var dfs func(row, col, index int) bool
	dfs = func(row, col, index int) bool {
		if row < 0 || row >= m || col < 0 || col >= n {
			return false
		}

		if board[row][col] != word[index] {
			return false
		}

		if index == len(word)-1 {
			return true
		}

		// 原地标记：将字符改为特殊值
		temp := board[row][col]
		board[row][col] = '#'

		for _, dir := range dirs {
			newRow, newCol := row+dir[0], col+dir[1]
			if dfs(newRow, newCol, index+1) {
				return true
			}
		}

		// 恢复原值
		board[row][col] = temp
		return false
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if dfs(i, j, 0) {
				return true
			}
		}
	}

	return false
}

// =========================== 方法三：BFS搜索 ===========================

func exist3(board [][]byte, word string) bool {
	if len(board) == 0 || len(board[0]) == 0 {
		return false
	}

	m, n := len(board), len(board[0])
	dirs := [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}

	type State struct {
		row, col, index int
		visited         [][]bool
	}

	queue := []State{}

	// 初始化队列
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == word[0] {
				visited := make([][]bool, m)
				for k := range visited {
					visited[k] = make([]bool, n)
				}
				visited[i][j] = true
				queue = append(queue, State{i, j, 0, visited})
			}
		}
	}

	for len(queue) > 0 {
		state := queue[0]
		queue = queue[1:]

		if state.index == len(word)-1 {
			return true
		}

		for _, dir := range dirs {
			newRow, newCol := state.row+dir[0], state.col+dir[1]
			if newRow >= 0 && newRow < m && newCol >= 0 && newCol < n {
				if !state.visited[newRow][newCol] && board[newRow][newCol] == word[state.index+1] {
					newVisited := make([][]bool, m)
					for i := range newVisited {
						newVisited[i] = make([]bool, n)
						copy(newVisited[i], state.visited[i])
					}
					newVisited[newRow][newCol] = true
					queue = append(queue, State{newRow, newCol, state.index + 1, newVisited})
				}
			}
		}
	}

	return false
}

// =========================== 方法四：递归枚举 ===========================

func exist4(board [][]byte, word string) bool {
	if len(board) == 0 || len(board[0]) == 0 {
		return false
	}

	m, n := len(board), len(board[0])
	dirs := [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}

	var search func(row, col, index int, visited [][]bool) bool
	search = func(row, col, index int, visited [][]bool) bool {
		// 边界检查
		if row < 0 || row >= m || col < 0 || col >= n {
			return false
		}

		// 已访问检查
		if visited[row][col] {
			return false
		}

		// 字符匹配检查
		if board[row][col] != word[index] {
			return false
		}

		// 到达单词末尾
		if index == len(word)-1 {
			return true
		}

		// 标记已访问
		visited[row][col] = true

		// 尝试四个方向
		for _, dir := range dirs {
			newRow, newCol := row+dir[0], col+dir[1]
			if search(newRow, newCol, index+1, visited) {
				return true
			}
		}

		// 撤销标记
		visited[row][col] = false
		return false
	}

	// 尝试所有起始位置
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			visited := make([][]bool, m)
			for k := range visited {
				visited[k] = make([]bool, n)
			}
			if search(i, j, 0, visited) {
				return true
			}
		}
	}

	return false
}

// =========================== 测试代码 ===========================

func main() {
	fmt.Println("=== LeetCode 79: 单词搜索 ===\n")

	testCases := []struct {
		board  [][]byte
		word   string
		expect bool
	}{
		{
			[][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}},
			"ABCCED",
			true,
		},
		{
			[][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}},
			"SEE",
			true,
		},
		{
			[][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}},
			"ABCB",
			false,
		},
		{
			[][]byte{{'A'}},
			"A",
			true,
		},
		{
			[][]byte{},
			"A",
			false,
		},
		{
			[][]byte{{'A', 'B'}, {'C', 'D'}},
			"ACDB",
			true,
		},
		{
			[][]byte{{'A', 'B'}, {'C', 'D'}},
			"ABCD",
			false,
		},
	}

	fmt.Println("方法一：DFS回溯（最优解法）")
	runTests(testCases, exist)

	fmt.Println("\n方法二：优化版DFS（原地标记）")
	runTests(testCases, exist2)

	fmt.Println("\n方法三：BFS搜索")
	runTests(testCases, exist3)

	fmt.Println("\n方法四：递归枚举")
	runTests(testCases, exist4)
}

func runTests(testCases []struct {
	board  [][]byte
	word   string
	expect bool
}, fn func([][]byte, string) bool) {
	passCount := 0
	for i, tc := range testCases {
		// 创建board的副本，避免原地修改影响其他测试
		boardCopy := make([][]byte, len(tc.board))
		for j := range tc.board {
			boardCopy[j] = make([]byte, len(tc.board[j]))
			copy(boardCopy[j], tc.board[j])
		}

		result := fn(boardCopy, tc.word)
		status := "✅"
		if result != tc.expect {
			status = "❌"
		} else {
			passCount++
		}
		fmt.Printf("  测试%d: %s\n", i+1, status)
		if status == "❌" {
			fmt.Printf("    输入: board=%v, word=\"%s\"\n", tc.board, tc.word)
			fmt.Printf("    输出: %t\n", result)
			fmt.Printf("    期望: %t\n", tc.expect)
		}
	}
	fmt.Printf("  通过: %d/%d\n", passCount, len(testCases))
}
