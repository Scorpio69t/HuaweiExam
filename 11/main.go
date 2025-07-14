package main

import (
	"fmt"
	"strings"
	"time"
)

// 暴力解法 - O(n²) 时间复杂度
func maxAreaBruteForce(height []int) int {
	maxArea := 0
	n := len(height)
	
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			// 计算当前容器的面积
			width := j - i
			minHeight := min(height[i], height[j])
			area := width * minHeight
			
			if area > maxArea {
				maxArea = area
			}
		}
	}
	
	return maxArea
}

// 双指针优化解法 - O(n) 时间复杂度
func maxAreaTwoPointers(height []int) int {
	left, right := 0, len(height)-1
	maxArea := 0
	
	for left < right {
		// 计算当前容器面积
		width := right - left
		minHeight := min(height[left], height[right])
		area := width * minHeight
		
		// 更新最大面积
		if area > maxArea {
			maxArea = area
		}
		
		// 移动较矮的指针
		// 关键洞察：移动较高的指针只会让宽度减小，高度不增加，面积必然减小
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}
	
	return maxArea
}

// 双指针优化解法（带详细过程记录）
func maxAreaWithProcess(height []int) (int, []string) {
	left, right := 0, len(height)-1
	maxArea := 0
	process := []string{}
	
	for left < right {
		width := right - left
		minHeight := min(height[left], height[right])
		area := width * minHeight
		
		process = append(process, fmt.Sprintf("左指针=%d(高度=%d), 右指针=%d(高度=%d), 宽度=%d, 面积=%d", 
			left, height[left], right, height[right], width, area))
		
		if area > maxArea {
			maxArea = area
			process = append(process, fmt.Sprintf("  -> 更新最大面积: %d", maxArea))
		}
		
		if height[left] < height[right] {
			left++
			process = append(process, "  -> 移动左指针")
		} else {
			right--
			process = append(process, "  -> 移动右指针")
		}
	}
	
	return maxArea, process
}

// 辅助函数：求最小值
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// 性能测试函数
func performanceTest(height []int) {
	fmt.Println("=== 性能测试 ===")
	
	// 测试暴力解法
	start := time.Now()
	result1 := maxAreaBruteForce(height)
	duration1 := time.Since(start)
	
	// 测试双指针解法
	start = time.Now()
	result2 := maxAreaTwoPointers(height)
	duration2 := time.Since(start)
	
	fmt.Printf("暴力解法结果: %d, 耗时: %v\n", result1, duration1)
	fmt.Printf("双指针解法结果: %d, 耗时: %v\n", result2, duration2)
	fmt.Printf("性能提升: %.2fx\n", float64(duration1)/float64(duration2))
}

func main() {
	// 测试用例
	testCases := []struct {
		name   string
		height []int
		expected int
	}{
		{"示例1", []int{1, 8, 6, 2, 5, 4, 8, 3, 7}, 49},
		{"示例2", []int{1, 1}, 1},
		{"递增序列", []int{1, 2, 3, 4, 5}, 6},
		{"递减序列", []int{5, 4, 3, 2, 1}, 6},
		{"全相同", []int{3, 3, 3, 3}, 9},
		{"两个元素", []int{2, 1}, 1},
		{"大数据", []int{1, 8, 6, 2, 5, 4, 8, 3, 7, 10, 9, 2, 6}, 48},
	}
	
	fmt.Println("=== 盛最多水的容器 - 算法测试 ===\n")
	
	for _, tc := range testCases {
		fmt.Printf("测试: %s\n", tc.name)
		fmt.Printf("输入: %v\n", tc.height)
		
		// 双指针解法
		result := maxAreaTwoPointers(tc.height)
		fmt.Printf("双指针结果: %d\n", result)
		
		// 验证结果
		if result == tc.expected {
			fmt.Println("✅ 测试通过")
		} else {
			fmt.Printf("❌ 测试失败，期望: %d\n", tc.expected)
		}
		
		// 对于小数据集，显示详细过程
		if len(tc.height) <= 9 {
			fmt.Println("\n详细计算过程:")
			_, process := maxAreaWithProcess(tc.height)
			for _, step := range process {
				fmt.Println(step)
			}
		}
		
		fmt.Println(strings.Repeat("-", 50))
	}
	
	// 性能测试
	largeData := make([]int, 10000)
	for i := range largeData {
		largeData[i] = i % 100 + 1
	}
	performanceTest(largeData)
}
