package main

import "fmt"

// 方法1：单调栈（推荐）
// 时间复杂度：O(n)，空间复杂度：O(n)
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

// 方法2：预处理 + 枚举
// 时间复杂度：O(n)，空间复杂度：O(n)
func largestRectangleAreaPreprocess(heights []int) int {
	n := len(heights)
	if n == 0 {
		return 0
	}

	left := make([]int, n)
	right := make([]int, n)

	// 预处理左边界
	stack := []int{}
	for i := 0; i < n; i++ {
		for len(stack) > 0 && heights[stack[len(stack)-1]] >= heights[i] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) == 0 {
			left[i] = -1
		} else {
			left[i] = stack[len(stack)-1]
		}
		stack = append(stack, i)
	}

	// 预处理右边界
	stack = []int{}
	for i := n - 1; i >= 0; i-- {
		for len(stack) > 0 && heights[stack[len(stack)-1]] >= heights[i] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) == 0 {
			right[i] = n
		} else {
			right[i] = stack[len(stack)-1]
		}
		stack = append(stack, i)
	}

	// 计算最大面积
	maxArea := 0
	for i := 0; i < n; i++ {
		width := right[i] - left[i] - 1
		area := heights[i] * width
		maxArea = max(maxArea, area)
	}

	return maxArea
}

// 方法3：暴力枚举
// 时间复杂度：O(n²)，空间复杂度：O(1)
func largestRectangleAreaBruteForce(heights []int) int {
	maxArea := 0
	n := len(heights)

	for i := 0; i < n; i++ {
		minHeight := heights[i]
		for j := i; j < n; j++ {
			minHeight = min(minHeight, heights[j])
			width := j - i + 1
			area := minHeight * width
			maxArea = max(maxArea, area)
		}
	}

	return maxArea
}

// 方法4：分治法
// 时间复杂度：O(n log n) 平均，O(n²) 最坏，空间复杂度：O(log n)
func largestRectangleAreaDivideConquer(heights []int) int {
	return divideConquer(heights, 0, len(heights)-1)
}

func divideConquer(heights []int, left, right int) int {
	if left > right {
		return 0
	}

	if left == right {
		return heights[left]
	}

	// 找到最小值位置
	minIdx := left
	for i := left + 1; i <= right; i++ {
		if heights[i] < heights[minIdx] {
			minIdx = i
		}
	}

	// 跨越最小值的矩形面积
	crossArea := heights[minIdx] * (right - left + 1)

	// 递归求解左右子问题
	leftArea := divideConquer(heights, left, minIdx-1)
	rightArea := divideConquer(heights, minIdx+1, right)

	return max(crossArea, max(leftArea, rightArea))
}

// 工具函数
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// 测试函数
func runTests() {
	fmt.Println("=== 84. 柱状图中最大的矩形 测试 ===")

	// 测试用例1
	heights1 := []int{2, 1, 5, 6, 2, 3}
	expected1 := 10

	fmt.Printf("\n测试用例1:\n")
	fmt.Printf("输入: %v\n", heights1)
	fmt.Printf("期望结果: %d\n", expected1)

	result1_1 := largestRectangleArea(heights1)
	result1_2 := largestRectangleAreaPreprocess(heights1)
	result1_3 := largestRectangleAreaBruteForce(heights1)
	result1_4 := largestRectangleAreaDivideConquer(heights1)

	fmt.Printf("方法一（单调栈）: %d ✓\n", result1_1)
	fmt.Printf("方法二（预处理）: %d ✓\n", result1_2)
	fmt.Printf("方法三（暴力枚举）: %d ✓\n", result1_3)
	fmt.Printf("方法四（分治法）: %d ✓\n", result1_4)

	// 测试用例2
	heights2 := []int{2, 4}
	expected2 := 4

	fmt.Printf("\n测试用例2:\n")
	fmt.Printf("输入: %v\n", heights2)
	fmt.Printf("期望结果: %d\n", expected2)

	result2_1 := largestRectangleArea(heights2)
	result2_2 := largestRectangleAreaPreprocess(heights2)
	result2_3 := largestRectangleAreaBruteForce(heights2)
	result2_4 := largestRectangleAreaDivideConquer(heights2)

	fmt.Printf("方法一（单调栈）: %d ✓\n", result2_1)
	fmt.Printf("方法二（预处理）: %d ✓\n", result2_2)
	fmt.Printf("方法三（暴力枚举）: %d ✓\n", result2_3)
	fmt.Printf("方法四（分治法）: %d ✓\n", result2_4)

	// 测试用例3：单个柱子
	heights3 := []int{5}
	expected3 := 5

	fmt.Printf("\n测试用例3（单个柱子）:\n")
	fmt.Printf("输入: %v\n", heights3)
	fmt.Printf("期望结果: %d\n", expected3)

	result3 := largestRectangleArea(heights3)
	fmt.Printf("结果: %d ✓\n", result3)

	// 测试用例4：递增序列
	heights4 := []int{1, 2, 3, 4, 5}
	expected4 := 9 // 高度3，宽度3

	fmt.Printf("\n测试用例4（递增序列）:\n")
	fmt.Printf("输入: %v\n", heights4)
	fmt.Printf("期望结果: %d\n", expected4)

	result4 := largestRectangleArea(heights4)
	fmt.Printf("结果: %d ✓\n", result4)

	// 测试用例5：递减序列
	heights5 := []int{5, 4, 3, 2, 1}
	expected5 := 9 // 高度3，宽度3

	fmt.Printf("\n测试用例5（递减序列）:\n")
	fmt.Printf("输入: %v\n", heights5)
	fmt.Printf("期望结果: %d\n", expected5)

	result5 := largestRectangleArea(heights5)
	fmt.Printf("结果: %d ✓\n", result5)

	// 测试用例6：相同高度
	heights6 := []int{3, 3, 3, 3}
	expected6 := 12 // 高度3，宽度4

	fmt.Printf("\n测试用例6（相同高度）:\n")
	fmt.Printf("输入: %v\n", heights6)
	fmt.Printf("期望结果: %d\n", expected6)

	result6 := largestRectangleArea(heights6)
	fmt.Printf("结果: %d ✓\n", result6)

	// 测试用例7：山峰形状
	heights7 := []int{1, 2, 3, 2, 1}
	expected7 := 6 // 高度2，宽度3

	fmt.Printf("\n测试用例7（山峰形状）:\n")
	fmt.Printf("输入: %v\n", heights7)
	fmt.Printf("期望结果: %d\n", expected7)

	result7 := largestRectangleArea(heights7)
	fmt.Printf("结果: %d ✓\n", result7)

	// 测试用例8：包含0的情况
	heights8 := []int{2, 0, 2}
	expected8 := 2

	fmt.Printf("\n测试用例8（包含0）:\n")
	fmt.Printf("输入: %v\n", heights8)
	fmt.Printf("期望结果: %d\n", expected8)

	result8 := largestRectangleArea(heights8)
	fmt.Printf("结果: %d ✓\n", result8)

	// 复杂示例分析
	fmt.Printf("\n=== 详细分析示例 ===\n")
	analyzeExample([]int{2, 1, 5, 6, 2, 3})

	fmt.Printf("\n=== 算法复杂度对比 ===\n")
	fmt.Printf("方法一（单调栈）：   时间 O(n),      空间 O(n)     - 推荐\n")
	fmt.Printf("方法二（预处理）：   时间 O(n),      空间 O(n)     - 思路清晰\n")
	fmt.Printf("方法三（暴力枚举）： 时间 O(n²),     空间 O(1)     - 简单易懂\n")
	fmt.Printf("方法四（分治法）：   时间 O(n logn), 空间 O(logn)  - 理论优雅\n")
}

// 详细分析示例
func analyzeExample(heights []int) {
	fmt.Printf("分析柱状图: %v\n", heights)
	fmt.Printf("索引位置:   ")
	for i := range heights {
		fmt.Printf("%d ", i)
	}
	fmt.Println()

	// 使用单调栈分析过程
	stack := []int{}
	maxArea := 0
	fmt.Printf("\n单调栈处理过程:\n")

	for i := 0; i <= len(heights); i++ {
		h := 0
		if i < len(heights) {
			h = heights[i]
		}

		fmt.Printf("步骤 %d: ", i)
		if i < len(heights) {
			fmt.Printf("当前高度=%d ", h)
		} else {
			fmt.Printf("结束处理 ")
		}

		for len(stack) > 0 && h < heights[stack[len(stack)-1]] {
			height := heights[stack[len(stack)-1]]
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			width := i
			if len(stack) > 0 {
				width = i - stack[len(stack)-1] - 1
			}

			area := height * width
			maxArea = max(maxArea, area)
			fmt.Printf("弹出索引%d(高度%d), 宽度=%d, 面积=%d ", top, height, width, area)
		}

		stack = append(stack, i)
		fmt.Printf("栈状态: %v\n", stack)
	}

	fmt.Printf("最大矩形面积: %d\n", maxArea)
}

func main() {
	runTests()
}
