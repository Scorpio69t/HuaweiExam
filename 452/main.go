package main

import (
	"fmt"
	"sort"
	"strings"
)

// 方法一：按结束坐标贪心选择（推荐）
// 时间复杂度：O(n log n)，空间复杂度：O(log n)
func findMinArrowShots(points [][]int) int {
	if len(points) == 0 {
		return 0
	}

	// 按结束坐标升序排序
	sort.Slice(points, func(i, j int) bool {
		return points[i][1] < points[j][1]
	})

	arrows := 1             // 至少需要一支箭
	lastEnd := points[0][1] // 第一支箭的位置

	// 从第二个气球开始遍历
	for i := 1; i < len(points); i++ {
		// 如果当前气球的起点大于上一支箭的位置，需要新箭
		if points[i][0] > lastEnd {
			arrows++
			lastEnd = points[i][1]
		}
	}

	return arrows
}

// 方法二：按起点排序等价解法
// 时间复杂度：O(n log n)，空间复杂度：O(log n)
func findMinArrowShotsByStart(points [][]int) int {
	if len(points) == 0 {
		return 0
	}

	// 按起点升序排序
	sort.Slice(points, func(i, j int) bool {
		return points[i][0] < points[j][0]
	})

	arrows := 1
	end := points[0][1] // 当前箭能覆盖的最远位置

	for i := 1; i < len(points); i++ {
		// 如果当前气球的起点大于当前箭能覆盖的最远位置，需要新箭
		if points[i][0] > end {
			arrows++
			end = points[i][1]
		} else {
			// 当前气球可以被当前箭覆盖，更新箭能覆盖的最远位置
			if points[i][1] < end {
				end = points[i][1]
			}
		}
	}

	return arrows
}

// 方法三：优化版本（清晰的贪心实现）
// 时间复杂度：O(n log n)，空间复杂度：O(log n)
func findMinArrowShotsOptimized(points [][]int) int {
	if len(points) == 0 {
		return 0
	}

	// 按结束坐标排序
	sort.Slice(points, func(i, j int) bool {
		return points[i][1] < points[j][1]
	})

	arrows := 1
	arrowPos := points[0][1]

	for i := 1; i < len(points); i++ {
		// 当前气球无法被上一支箭击中
		if points[i][0] > arrowPos {
			arrows++
			arrowPos = points[i][1]
		}
	}

	return arrows
}

// 辅助函数：打印气球数组
func printPoints(points [][]int) string {
	if len(points) == 0 {
		return "[]"
	}
	result := "["
	for i, point := range points {
		if i > 0 {
			result += ","
		}
		result += fmt.Sprintf("[%d,%d]", point[0], point[1])
	}
	result += "]"
	return result
}

func runTests() {
	type testCase struct {
		points   [][]int
		expected int
		desc     string
	}

	tests := []testCase{
		{[][]int{{10, 16}, {2, 8}, {1, 6}, {7, 12}}, 2, "示例1"},
		{[][]int{{1, 2}, {3, 4}, {5, 6}, {7, 8}}, 4, "示例2"},
		{[][]int{{1, 2}, {2, 3}, {3, 4}, {4, 5}}, 2, "示例3"},
		{[][]int{}, 0, "空数组"},
		{[][]int{{1, 2}}, 1, "单气球"},
		{[][]int{{1, 3}, {2, 4}, {3, 5}}, 1, "全重叠"},
		{[][]int{{1, 2}, {3, 4}, {5, 6}}, 3, "无重叠"},
		{[][]int{{1, 100}, {2, 50}, {3, 75}}, 1, "大范围气球"},
		{[][]int{{-5, -1}, {-3, 2}, {0, 5}}, 2, "负数坐标"},
		{[][]int{{1, 3}, {2, 4}, {3, 5}, {6, 8}}, 2, "部分重叠"},
		{[][]int{{1, 2}, {2, 3}, {3, 4}, {4, 5}, {5, 6}}, 3, "连续重叠"},
		{[][]int{{1, 2}, {1, 3}, {1, 4}, {1, 5}}, 1, "同起点"},
	}

	fmt.Println("=== 452. 用最少数量的箭引爆气球 - 测试 ===")
	for i, tc := range tests {
		r1 := findMinArrowShots(clonePoints(tc.points))
		r2 := findMinArrowShotsByStart(clonePoints(tc.points))
		r3 := findMinArrowShotsOptimized(clonePoints(tc.points))

		ok := (r1 == tc.expected) && (r2 == tc.expected) && (r3 == tc.expected)
		status := "✅"
		if !ok {
			status = "❌"
		}

		fmt.Printf("用例 %d: %s\n", i+1, tc.desc)
		fmt.Printf("输入: %s\n", printPoints(tc.points))
		fmt.Printf("期望: %d\n", tc.expected)
		fmt.Printf("按结束贪心: %d, 按起点贪心: %d, 优化版本: %d\n", r1, r2, r3)
		fmt.Printf("结果: %s\n", status)
		fmt.Println(strings.Repeat("-", 60))
	}
}

func clonePoints(src [][]int) [][]int {
	dst := make([][]int, len(src))
	for i := range src {
		dst[i] = make([]int, len(src[i]))
		copy(dst[i], src[i])
	}
	return dst
}

func main() {
	runTests()
}
