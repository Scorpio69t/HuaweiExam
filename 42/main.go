package main

import (
	"fmt"
	"time"
)

// 方法1：暴力法 - 时间复杂度O(n²)
func trapBruteForce(height []int) int {
	n := len(height)
	if n <= 2 {
		return 0
	}

	totalWater := 0

	// 对每个位置计算能接的雨水
	for i := 1; i < n-1; i++ {
		// 找左边最高的柱子
		leftMax := 0
		for j := 0; j < i; j++ {
			if height[j] > leftMax {
				leftMax = height[j]
			}
		}

		// 找右边最高的柱子
		rightMax := 0
		for j := i + 1; j < n; j++ {
			if height[j] > rightMax {
				rightMax = height[j]
			}
		}

		// 计算当前位置能接的雨水
		minHeight := min(leftMax, rightMax)
		if minHeight > height[i] {
			totalWater += minHeight - height[i]
		}
	}

	return totalWater
}

// 方法2：动态规划 - 时间复杂度O(n)，空间复杂度O(n)
func trapDP(height []int) int {
	n := len(height)
	if n <= 2 {
		return 0
	}

	// 预计算左边最高和右边最高
	leftMax := make([]int, n)
	rightMax := make([]int, n)

	// 计算每个位置左边的最高柱子
	leftMax[0] = height[0]
	for i := 1; i < n; i++ {
		leftMax[i] = max(leftMax[i-1], height[i])
	}

	// 计算每个位置右边的最高柱子
	rightMax[n-1] = height[n-1]
	for i := n - 2; i >= 0; i-- {
		rightMax[i] = max(rightMax[i+1], height[i])
	}

	totalWater := 0
	for i := 1; i < n-1; i++ {
		minHeight := min(leftMax[i], rightMax[i])
		if minHeight > height[i] {
			totalWater += minHeight - height[i]
		}
	}

	return totalWater
}

// 方法3：双指针 - 时间复杂度O(n)，空间复杂度O(1)
func trapTwoPointers(height []int) int {
	n := len(height)
	if n <= 2 {
		return 0
	}

	left, right := 0, n-1
	leftMax, rightMax := 0, 0
	totalWater := 0

	for left < right {
		if height[left] < height[right] {
			if height[left] >= leftMax {
				leftMax = height[left]
			} else {
				totalWater += leftMax - height[left]
			}
			left++
		} else {
			if height[right] >= rightMax {
				rightMax = height[right]
			} else {
				totalWater += rightMax - height[right]
			}
			right--
		}
	}

	return totalWater
}

// 方法4：单调栈 - 时间复杂度O(n)，空间复杂度O(n)
func trapMonotonicStack(height []int) int {
	n := len(height)
	if n <= 2 {
		return 0
	}

	stack := make([]int, 0) // 存储下标
	totalWater := 0

	for i := 0; i < n; i++ {
		// 当前柱子高度大于栈顶柱子时，可以形成凹槽
		for len(stack) > 0 && height[i] > height[stack[len(stack)-1]] {
			// 弹出栈顶（凹槽底部）
			bottom := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			// 如果栈为空，无法形成凹槽
			if len(stack) == 0 {
				break
			}

			// 计算凹槽的宽度和高度
			left := stack[len(stack)-1]
			width := i - left - 1
			minHeight := min(height[left], height[i])
			waterHeight := minHeight - height[bottom]

			totalWater += width * waterHeight
		}

		stack = append(stack, i)
	}

	return totalWater
}

// 方法5：带详细过程展示的双指针法
func trapWithTrace(height []int) int {
	n := len(height)
	if n <= 2 {
		return 0
	}

	fmt.Printf("输入数组: %v\n", height)
	fmt.Println("双指针过程:")
	fmt.Println("步骤\t左指针\t右指针\t左最高\t右最高\t当前水量\t总水量")

	left, right := 0, n-1
	leftMax, rightMax := 0, 0
	totalWater := 0
	step := 0

	for left < right {
		step++
		currentWater := 0

		if height[left] < height[right] {
			if height[left] >= leftMax {
				leftMax = height[left]
			} else {
				currentWater = leftMax - height[left]
				totalWater += currentWater
			}
			fmt.Printf("%d\t%d\t%d\t%d\t%d\t%d\t\t%d\n",
				step, left, right, leftMax, rightMax, currentWater, totalWater)
			left++
		} else {
			if height[right] >= rightMax {
				rightMax = height[right]
			} else {
				currentWater = rightMax - height[right]
				totalWater += currentWater
			}
			fmt.Printf("%d\t%d\t%d\t%d\t%d\t%d\t\t%d\n",
				step, left, right, leftMax, rightMax, currentWater, totalWater)
			right--
		}
	}

	return totalWater
}

// 可视化雨水图
func visualizeRainwater(height []int) {
	if len(height) == 0 {
		return
	}

	maxHeight := 0
	for _, h := range height {
		if h > maxHeight {
			maxHeight = h
		}
	}

	fmt.Println("\n=== 雨水可视化 ===")
	fmt.Printf("原始柱子: %v\n", height)

	// 计算每个位置的水位
	water := make([]int, len(height))
	trapResult := trapDP(height)

	// 使用动态规划计算水位
	n := len(height)
	leftMax := make([]int, n)
	rightMax := make([]int, n)

	leftMax[0] = height[0]
	for i := 1; i < n; i++ {
		leftMax[i] = max(leftMax[i-1], height[i])
	}

	rightMax[n-1] = height[n-1]
	for i := n - 2; i >= 0; i-- {
		rightMax[i] = max(rightMax[i+1], height[i])
	}

	for i := 0; i < n; i++ {
		water[i] = min(leftMax[i], rightMax[i])
	}

	// 从上到下打印
	for level := maxHeight; level > 0; level-- {
		for i := 0; i < len(height); i++ {
			if height[i] >= level {
				fmt.Print("█") // 柱子
			} else if water[i] >= level {
				fmt.Print("~") // 雨水
			} else {
				fmt.Print(" ") // 空气
			}
		}
		fmt.Println()
	}

	// 打印底部索引
	for i := 0; i < len(height); i++ {
		fmt.Printf("%d", i%10)
	}
	fmt.Printf("\n总雨水量: %d\n", trapResult)
}

// 性能测试函数
func performanceTest() {
	testCases := [][]int{
		{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1},
		{4, 2, 0, 3, 2, 5},
		{3, 0, 2, 0, 4},
		{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1, 0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1},
	}

	fmt.Println("\n=== 性能测试 ===")
	fmt.Println("测试用例\t\t\t暴力法耗时\t动态规划耗时\t双指针耗时\t单调栈耗时")

	for i, testCase := range testCases {
		// 暴力法测试
		start := time.Now()
		result1 := trapBruteForce(testCase)
		time1 := time.Since(start)

		// 动态规划测试
		start = time.Now()
		result2 := trapDP(testCase)
		time2 := time.Since(start)

		// 双指针测试
		start = time.Now()
		result3 := trapTwoPointers(testCase)
		time3 := time.Since(start)

		// 单调栈测试
		start = time.Now()
		result4 := trapMonotonicStack(testCase)
		time4 := time.Since(start)

		// 验证结果一致性
		if result1 != result2 || result2 != result3 || result3 != result4 {
			fmt.Printf("错误：结果不一致！%d, %d, %d, %d\n", result1, result2, result3, result4)
		}

		fmt.Printf("测试用例%d\t\t\t%v\t\t%v\t\t%v\t\t%v\n",
			i+1, time1, time2, time3, time4)
	}
}

// 辅助函数
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	// 基本测试用例
	testCases := [][]int{
		{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1},
		{4, 2, 0, 3, 2, 5},
		{3, 0, 2, 0, 4},
		{},
		{1},
		{1, 2},
	}

	fmt.Println("=== 基本测试 ===")
	for i, testCase := range testCases {
		result := trapTwoPointers(testCase)
		fmt.Printf("测试用例%d: %v -> 输出: %d\n", i+1, testCase, result)
	}

	fmt.Println("\n=== 详细过程展示 ===")
	trapWithTrace([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1})

	// 可视化展示
	visualizeRainwater([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1})

	// 性能测试
	performanceTest()

	fmt.Println("\n=== 算法复杂度分析 ===")
	fmt.Println("暴力法：")
	fmt.Println("  时间复杂度：O(n²)")
	fmt.Println("  空间复杂度：O(1)")
	fmt.Println("\n动态规划：")
	fmt.Println("  时间复杂度：O(n)")
	fmt.Println("  空间复杂度：O(n)")
	fmt.Println("\n双指针：")
	fmt.Println("  时间复杂度：O(n)")
	fmt.Println("  空间复杂度：O(1) - 最优解")
	fmt.Println("\n单调栈：")
	fmt.Println("  时间复杂度：O(n)")
	fmt.Println("  空间复杂度：O(n)")
}
