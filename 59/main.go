package main

import (
	"fmt"
	"time"
)

// 方法一：边界收缩算法（最优解法）
func generateMatrix1(n int) [][]int {
	// 创建n×n矩阵
	matrix := make([][]int, n)
	for i := range matrix {
		matrix[i] = make([]int, n)
	}

	// 初始化边界
	top, bottom := 0, n-1
	left, right := 0, n-1
	num := 1

	// 螺旋填充
	for top <= bottom && left <= right {
		// 填充上边界（从左到右）
		for i := left; i <= right; i++ {
			matrix[top][i] = num
			num++
		}
		top++

		// 填充右边界（从上到下）
		for i := top; i <= bottom; i++ {
			matrix[i][right] = num
			num++
		}
		right--

		// 填充下边界（从右到左）
		if top <= bottom {
			for i := right; i >= left; i-- {
				matrix[bottom][i] = num
				num++
			}
			bottom--
		}

		// 填充左边界（从下到上）
		if left <= right {
			for i := bottom; i >= top; i-- {
				matrix[i][left] = num
				num++
			}
			left++
		}
	}

	return matrix
}

// 方法二：方向数组算法
func generateMatrix2(n int) [][]int {
	matrix := make([][]int, n)
	for i := range matrix {
		matrix[i] = make([]int, n)
	}

	// 方向数组：右、下、左、上
	dirs := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	dirIdx := 0

	row, col := 0, 0
	visited := make([][]bool, n)
	for i := range visited {
		visited[i] = make([]bool, n)
	}

	for num := 1; num <= n*n; num++ {
		matrix[row][col] = num
		visited[row][col] = true

		// 计算下一个位置
		nextRow := row + dirs[dirIdx][0]
		nextCol := col + dirs[dirIdx][1]

		// 检查是否需要转向
		if nextRow < 0 || nextRow >= n || nextCol < 0 || nextCol >= n || visited[nextRow][nextCol] {
			// 转向
			dirIdx = (dirIdx + 1) % 4
			nextRow = row + dirs[dirIdx][0]
			nextCol = col + dirs[dirIdx][1]
		}

		row, col = nextRow, nextCol
	}

	return matrix
}

// 方法三：递归分层算法
func generateMatrix3(n int) [][]int {
	matrix := make([][]int, n)
	for i := range matrix {
		matrix[i] = make([]int, n)
	}

	num := 1
	fillLayer(matrix, 0, n-1, 0, n-1, &num)
	return matrix
}

// 递归填充一层
func fillLayer(matrix [][]int, top, bottom, left, right int, num *int) {
	if top > bottom || left > right {
		return
	}

	// 填充上边界
	for i := left; i <= right; i++ {
		matrix[top][i] = *num
		(*num)++
	}

	// 填充右边界
	for i := top + 1; i <= bottom; i++ {
		matrix[i][right] = *num
		(*num)++
	}

	// 填充下边界（如果还有）
	if top < bottom {
		for i := right - 1; i >= left; i-- {
			matrix[bottom][i] = *num
			(*num)++
		}
	}

	// 填充左边界（如果还有）
	if left < right {
		for i := bottom - 1; i > top; i-- {
			matrix[i][left] = *num
			(*num)++
		}
	}

	// 递归处理内层
	fillLayer(matrix, top+1, bottom-1, left+1, right-1, num)
}

// 方法四：模拟填充算法
func generateMatrix4(n int) [][]int {
	matrix := make([][]int, n)
	for i := range matrix {
		matrix[i] = make([]int, n)
	}

	num := 1
	row, col := 0, 0
	direction := 0 // 0:右, 1:下, 2:左, 3:上

	for num <= n*n {
		matrix[row][col] = num
		num++

		// 尝试按当前方向前进
		var nextRow, nextCol int
		switch direction {
		case 0: // 右
			nextCol = col + 1
			nextRow = row
			if nextCol >= n || matrix[nextRow][nextCol] != 0 {
				direction = 1
				nextRow = row + 1
				nextCol = col
			}
		case 1: // 下
			nextRow = row + 1
			nextCol = col
			if nextRow >= n || matrix[nextRow][nextCol] != 0 {
				direction = 2
				nextRow = row
				nextCol = col - 1
			}
		case 2: // 左
			nextCol = col - 1
			nextRow = row
			if nextCol < 0 || matrix[nextRow][nextCol] != 0 {
				direction = 3
				nextRow = row - 1
				nextCol = col
			}
		case 3: // 上
			nextRow = row - 1
			nextCol = col
			if nextRow < 0 || matrix[nextRow][nextCol] != 0 {
				direction = 0
				nextRow = row
				nextCol = col + 1
			}
		}

		row, col = nextRow, nextCol
	}

	return matrix
}

// 辅助函数：打印矩阵
func printMatrix(matrix [][]int) {
	for _, row := range matrix {
		fmt.Println(row)
	}
}

// 辅助函数：比较两个矩阵是否相等
func matrixEqual(m1, m2 [][]int) bool {
	if len(m1) != len(m2) {
		return false
	}
	for i := range m1 {
		if len(m1[i]) != len(m2[i]) {
			return false
		}
		for j := range m1[i] {
			if m1[i][j] != m2[i][j] {
				return false
			}
		}
	}
	return true
}

// 验证矩阵是否正确（包含1到n²且按螺旋顺序）
func validateMatrix(matrix [][]int) bool {
	n := len(matrix)
	if n == 0 {
		return false
	}

	// 检查是否包含所有数字
	seen := make(map[int]bool)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] < 1 || matrix[i][j] > n*n {
				return false
			}
			if seen[matrix[i][j]] {
				return false
			}
			seen[matrix[i][j]] = true
		}
	}

	return len(seen) == n*n
}

// 测试用例
func createTestCases() []struct {
	n    int
	name string
} {
	return []struct {
		n    int
		name string
	}{
		{1, "n=1（最小值）"},
		{2, "n=2（偶数）"},
		{3, "n=3（奇数）"},
		{4, "n=4（偶数）"},
		{5, "n=5（奇数）"},
		{10, "n=10（较大值）"},
	}
}

// 性能测试
func benchmarkAlgorithm(algorithm func(int) [][]int, n int, name string) {
	iterations := 1000
	start := time.Now()

	for i := 0; i < iterations; i++ {
		algorithm(n)
	}

	duration := time.Since(start)
	avgTime := duration.Nanoseconds() / int64(iterations)
	fmt.Printf("%s (n=%d): 平均执行时间 %d 纳秒\n", name, n, avgTime)
}

func main() {
	fmt.Println("=== 59. 螺旋矩阵 II ===")
	fmt.Println()

	testCases := createTestCases()
	algorithms := []struct {
		name string
		fn   func(int) [][]int
	}{
		{"边界收缩算法", generateMatrix1},
		{"方向数组算法", generateMatrix2},
		{"递归分层算法", generateMatrix3},
		{"模拟填充算法", generateMatrix4},
	}

	// 正确性测试
	fmt.Println("=== 算法正确性测试 ===")
	for _, testCase := range testCases {
		fmt.Printf("测试: %s\n", testCase.name)

		results := make([][][]int, len(algorithms))
		for i, algo := range algorithms {
			results[i] = algo.fn(testCase.n)
		}

		// 检查所有算法结果是否一致
		allEqual := true
		for i := 1; i < len(results); i++ {
			if !matrixEqual(results[i], results[0]) {
				allEqual = false
				break
			}
		}

		// 验证结果是否正确
		isValid := validateMatrix(results[0])

		if allEqual && isValid {
			fmt.Printf("  ✅ 所有算法结果一致且正确\n")
			if testCase.n <= 5 {
				fmt.Printf("  矩阵结果:\n")
				for _, row := range results[0] {
					fmt.Printf("  %v\n", row)
				}
			}
		} else {
			fmt.Printf("  ❌ 算法结果不一致或错误\n")
			for i, algo := range algorithms {
				valid := validateMatrix(results[i])
				fmt.Printf("  %s: 有效=%v\n", algo.name, valid)
				if testCase.n <= 5 {
					printMatrix(results[i])
				}
			}
		}
		fmt.Println()
	}

	// 性能测试
	fmt.Println("=== 性能测试 ===")
	perfTestSizes := []int{5, 10}

	for _, n := range perfTestSizes {
		fmt.Printf("测试规模 n=%d:\n", n)
		for _, algo := range algorithms {
			benchmarkAlgorithm(algo.fn, n, algo.name)
		}
		fmt.Println()
	}

	// 算法分析
	fmt.Println("=== 算法分析 ===")
	fmt.Println("螺旋矩阵II问题的特点:")
	fmt.Println("1. 按顺时针螺旋顺序填充1到n²")
	fmt.Println("2. 需要维护四个边界或方向状态")
	fmt.Println("3. 边界收缩法是最优解法")
	fmt.Println()

	fmt.Println("=== 复杂度分析 ===")
	fmt.Println("时间复杂度:")
	fmt.Println("- 边界收缩: O(n²)，遍历所有元素一次")
	fmt.Println("- 方向数组: O(n²)，遍历所有元素一次")
	fmt.Println("- 递归分层: O(n²)，递归处理每一层")
	fmt.Println("- 模拟填充: O(n²)，遍历所有元素一次")
	fmt.Println()
	fmt.Println("空间复杂度:")
	fmt.Println("- 边界收缩: O(1)，只使用常数额外空间")
	fmt.Println("- 方向数组: O(n²)，需要visited数组")
	fmt.Println("- 递归分层: O(n)，递归调用栈深度")
	fmt.Println("- 模拟填充: O(1)，只使用常数额外空间")
	fmt.Println()

	fmt.Println("=== 优化技巧总结 ===")
	fmt.Println("1. 边界管理：准确维护四个边界的变化")
	fmt.Println("2. 填充顺序：严格按照右→下→左→上")
	fmt.Println("3. 边界检查：避免数组越界")
	fmt.Println("4. 数字递增：每填充一个位置num++")
	fmt.Println()

	fmt.Println("推荐使用：边界收缩算法（方法一），空间效率最高")
}
