package main

import (
	"fmt"
	"reflect"
	"time"
)

// 方法一：边界收缩算法
// 最优解法，维护四个边界，每遍历完一层就收缩边界
func spiralOrder1(matrix [][]int) []int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return []int{}
	}

	m, n := len(matrix), len(matrix[0])
	result := make([]int, 0, m*n)

	top, bottom := 0, m-1
	left, right := 0, n-1

	for top <= bottom && left <= right {
		// 遍历上边界（从左到右）
		for i := left; i <= right; i++ {
			result = append(result, matrix[top][i])
		}
		top++

		// 遍历右边界（从上到下）
		for i := top; i <= bottom; i++ {
			result = append(result, matrix[i][right])
		}
		right--

		// 遍历下边界（从右到左）
		if top <= bottom {
			for i := right; i >= left; i-- {
				result = append(result, matrix[bottom][i])
			}
			bottom--
		}

		// 遍历左边界（从下到上）
		if left <= right {
			for i := bottom; i >= top; i-- {
				result = append(result, matrix[i][left])
			}
			left++
		}
	}

	return result
}

// 方法二：方向数组算法
// 使用方向数组控制遍历方向，使用visited数组标记已访问元素
func spiralOrder2(matrix [][]int) []int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return []int{}
	}

	m, n := len(matrix), len(matrix[0])
	result := make([]int, 0, m*n)
	visited := make([][]bool, m)
	for i := range visited {
		visited[i] = make([]bool, n)
	}

	// 方向数组：右、下、左、上
	dirs := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	row, col, dir := 0, 0, 0

	for i := 0; i < m*n; i++ {
		result = append(result, matrix[row][col])
		visited[row][col] = true

		// 计算下一个位置
		nextRow := row + dirs[dir][0]
		nextCol := col + dirs[dir][1]

		// 检查是否需要改变方向
		if nextRow < 0 || nextRow >= m || nextCol < 0 || nextCol >= n || visited[nextRow][nextCol] {
			dir = (dir + 1) % 4
			nextRow = row + dirs[dir][0]
			nextCol = col + dirs[dir][1]
		}

		row, col = nextRow, nextCol
	}

	return result
}

// 方法三：递归分层算法
// 将问题分解为外层和内层，递归处理每一层
func spiralOrder3(matrix [][]int) []int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return []int{}
	}

	return spiralHelper(matrix, 0, len(matrix)-1, 0, len(matrix[0])-1)
}

// 递归辅助函数
func spiralHelper(matrix [][]int, top, bottom, left, right int) []int {
	if top > bottom || left > right {
		return []int{}
	}

	result := []int{}

	// 遍历上边界
	for i := left; i <= right; i++ {
		result = append(result, matrix[top][i])
	}

	// 遍历右边界
	for i := top + 1; i <= bottom; i++ {
		result = append(result, matrix[i][right])
	}

	// 遍历下边界
	if top < bottom {
		for i := right - 1; i >= left; i-- {
			result = append(result, matrix[bottom][i])
		}
	}

	// 遍历左边界
	if left < right {
		for i := bottom - 1; i > top; i-- {
			result = append(result, matrix[i][left])
		}
	}

	// 递归处理内层
	inner := spiralHelper(matrix, top+1, bottom-1, left+1, right-1)
	result = append(result, inner...)

	return result
}

// 方法四：模拟遍历算法
// 直接模拟螺旋遍历的过程，使用visited数组标记已访问元素
func spiralOrder4(matrix [][]int) []int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return []int{}
	}

	m, n := len(matrix), len(matrix[0])
	result := make([]int, 0, m*n)
	visited := make([][]bool, m)
	for i := range visited {
		visited[i] = make([]bool, n)
	}

	row, col := 0, 0
	direction := 0 // 0: 右, 1: 下, 2: 左, 3: 上

	for len(result) < m*n {
		result = append(result, matrix[row][col])
		visited[row][col] = true

		// 根据方向计算下一个位置
		var nextRow, nextCol int
		switch direction {
		case 0: // 右
			nextRow, nextCol = row, col+1
		case 1: // 下
			nextRow, nextCol = row+1, col
		case 2: // 左
			nextRow, nextCol = row, col-1
		case 3: // 上
			nextRow, nextCol = row-1, col
		}

		// 检查是否需要改变方向
		if nextRow < 0 || nextRow >= m || nextCol < 0 || nextCol >= n || visited[nextRow][nextCol] {
			direction = (direction + 1) % 4
			switch direction {
			case 0:
				nextRow, nextCol = row, col+1
			case 1:
				nextRow, nextCol = row+1, col
			case 2:
				nextRow, nextCol = row, col-1
			case 3:
				nextRow, nextCol = row-1, col
			}
		}

		row, col = nextRow, nextCol
	}

	return result
}

// 辅助函数：创建测试用例
func createTestCases() []struct {
	matrix   [][]int
	expected []int
	name     string
} {
	return []struct {
		matrix   [][]int
		expected []int
		name     string
	}{
		{
			[][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
			[]int{1, 2, 3, 6, 9, 8, 7, 4, 5},
			"示例1: 3×3矩阵",
		},
		{
			[][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}},
			[]int{1, 2, 3, 4, 8, 12, 11, 10, 9, 5, 6, 7},
			"示例2: 3×4矩阵",
		},
		{
			[][]int{{1}},
			[]int{1},
			"测试1: 单个元素",
		},
		{
			[][]int{{1, 2, 3, 4}},
			[]int{1, 2, 3, 4},
			"测试2: 单行矩阵",
		},
		{
			[][]int{{1}, {2}, {3}, {4}},
			[]int{1, 2, 3, 4},
			"测试3: 单列矩阵",
		},
		{
			[][]int{{1, 2}, {3, 4}},
			[]int{1, 2, 4, 3},
			"测试4: 2×2矩阵",
		},
		{
			[][]int{{1, 2, 3}, {4, 5, 6}},
			[]int{1, 2, 3, 6, 5, 4},
			"测试5: 2×3矩阵",
		},
		{
			[][]int{{1, 2}, {3, 4}, {5, 6}},
			[]int{1, 2, 4, 6, 5, 3},
			"测试6: 3×2矩阵",
		},
		{
			[][]int{{1, 2, 3, 4, 5}, {6, 7, 8, 9, 10}, {11, 12, 13, 14, 15}},
			[]int{1, 2, 3, 4, 5, 10, 15, 14, 13, 12, 11, 6, 7, 8, 9},
			"测试7: 3×5矩阵",
		},
		{
			[][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}, {10, 11, 12}},
			[]int{1, 2, 3, 6, 9, 12, 11, 10, 7, 4, 5, 8},
			"测试8: 4×3矩阵",
		},
	}
}

// 性能测试函数
func benchmarkAlgorithm(algorithm func([][]int) []int, matrix [][]int, name string) {
	iterations := 1000
	start := time.Now()

	for i := 0; i < iterations; i++ {
		algorithm(matrix)
	}

	duration := time.Since(start)
	avgTime := duration.Nanoseconds() / int64(iterations)

	fmt.Printf("%s: 平均执行时间 %d 纳秒\n", name, avgTime)
}

// 辅助函数：打印矩阵
func printMatrix(matrix [][]int) {
	fmt.Print("[")
	for i, row := range matrix {
		if i > 0 {
			fmt.Print(",")
		}
		fmt.Print("[")
		for j, val := range row {
			if j > 0 {
				fmt.Print(",")
			}
			fmt.Print(val)
		}
		fmt.Print("]")
	}
	fmt.Print("]")
}

// 辅助函数：打印结果
func printResult(result []int) {
	fmt.Print("[")
	for i, val := range result {
		if i > 0 {
			fmt.Print(",")
		}
		fmt.Print(val)
	}
	fmt.Print("]")
}

func main() {
	fmt.Println("=== 54. 螺旋矩阵 ===")
	fmt.Println()

	// 创建测试用例
	testCases := createTestCases()
	algorithms := []struct {
		name string
		fn   func([][]int) []int
	}{
		{"边界收缩算法", spiralOrder1},
		{"方向数组算法", spiralOrder2},
		{"递归分层算法", spiralOrder3},
		{"模拟遍历算法", spiralOrder4},
	}

	// 运行测试
	fmt.Println("=== 算法正确性测试 ===")
	for _, testCase := range testCases {
		fmt.Printf("测试: %s\n", testCase.name)

		results := make([][]int, len(algorithms))
		for i, algo := range algorithms {
			results[i] = algo.fn(testCase.matrix)
		}

		// 验证所有算法结果一致
		allEqual := true
		for i := 1; i < len(results); i++ {
			if !reflect.DeepEqual(results[i], results[0]) {
				allEqual = false
				break
			}
		}

		// 验证结果是否正确
		allValid := reflect.DeepEqual(results[0], testCase.expected)

		if allEqual && allValid {
			fmt.Printf("  ✅ 所有算法结果一致且正确\n")
			fmt.Print("  输入矩阵: ")
			printMatrix(testCase.matrix)
			fmt.Println()
			fmt.Print("  输出结果: ")
			printResult(results[0])
			fmt.Println()
		} else {
			fmt.Printf("  ❌ 算法结果不一致或错误\n")
			fmt.Print("  输入矩阵: ")
			printMatrix(testCase.matrix)
			fmt.Println()
			fmt.Print("  预期结果: ")
			printResult(testCase.expected)
			fmt.Println()
			for i, algo := range algorithms {
				fmt.Printf("    %s: ", algo.name)
				printResult(results[i])
				fmt.Println()
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
	}

	fmt.Printf("测试数据: %d×%d矩阵\n", len(performanceMatrix), len(performanceMatrix[0]))
	fmt.Println()

	for _, algo := range algorithms {
		benchmarkAlgorithm(algo.fn, performanceMatrix, algo.name)
	}
	fmt.Println()

	// 算法分析
	fmt.Println("=== 算法分析 ===")
	fmt.Println("螺旋矩阵问题的特点:")
	fmt.Println("1. 需要按照顺时针螺旋顺序遍历矩阵")
	fmt.Println("2. 遍历顺序为：右→下→左→上")
	fmt.Println("3. 需要维护边界或使用方向数组")
	fmt.Println("4. 边界收缩法是最优解法")
	fmt.Println()

	// 复杂度分析
	fmt.Println("=== 复杂度分析 ===")
	fmt.Println("时间复杂度:")
	fmt.Println("- 边界收缩: O(m×n)，需要遍历所有元素一次")
	fmt.Println("- 方向数组: O(m×n)，需要遍历所有元素一次")
	fmt.Println("- 递归分层: O(m×n)，需要遍历所有元素一次")
	fmt.Println("- 模拟遍历: O(m×n)，需要遍历所有元素一次")
	fmt.Println()
	fmt.Println("空间复杂度:")
	fmt.Println("- 边界收缩: O(1)，只使用常数额外空间")
	fmt.Println("- 方向数组: O(m×n)，需要visited数组")
	fmt.Println("- 递归分层: O(min(m,n))，递归栈的深度")
	fmt.Println("- 模拟遍历: O(m×n)，需要visited数组")
	fmt.Println()

	// 算法总结
	fmt.Println("=== 算法总结 ===")
	fmt.Println("1. 边界收缩算法：最优解法，空间复杂度O(1)")
	fmt.Println("2. 方向数组算法：逻辑清晰，但需要额外空间")
	fmt.Println("3. 递归分层算法：递归实现，空间开销较大")
	fmt.Println("4. 模拟遍历算法：最直观，但空间开销大")
	fmt.Println()
	fmt.Println("推荐使用：边界收缩算法（方法一），效率最高")
	fmt.Println()

	// 应用场景
	fmt.Println("=== 应用场景 ===")
	fmt.Println("- 算法竞赛：矩阵遍历的经典应用")
	fmt.Println("- 图像处理：螺旋扫描图像")
	fmt.Println("- 数据可视化：螺旋布局")
	fmt.Println("- 游戏开发：地图遍历")
	fmt.Println("- 打印输出：螺旋打印")
	fmt.Println()

	// 优化技巧总结
	fmt.Println("=== 优化技巧总结 ===")
	fmt.Println("1. 边界收缩：掌握四个边界的更新规则")
	fmt.Println("2. 方向控制：理解方向数组的使用方法")
	fmt.Println("3. 递归思想：学会将问题分解为子问题")
	fmt.Println("4. 边界处理：注意各种边界情况")
	fmt.Println("5. 算法选择：根据问题特点选择合适的算法")
	fmt.Println("6. 优化策略：学会时间和空间优化技巧")
}
