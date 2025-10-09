package main

import (
	"fmt"
	"reflect"
	"sort"
	"time"
)

// 方法一：线性扫描算法（最优解法）
func insert1(intervals [][]int, newInterval []int) [][]int {
	result := [][]int{}
	i := 0
	n := len(intervals)

	// 添加所有在新区间左侧的区间
	for i < n && intervals[i][1] < newInterval[0] {
		result = append(result, intervals[i])
		i++
	}

	// 合并所有与新区间重叠的区间
	for i < n && intervals[i][0] <= newInterval[1] {
		newInterval[0] = min(newInterval[0], intervals[i][0])
		newInterval[1] = max(newInterval[1], intervals[i][1])
		i++
	}
	result = append(result, newInterval)

	// 添加所有在新区间右侧的区间
	for i < n {
		result = append(result, intervals[i])
		i++
	}

	return result
}

// 方法二：二分查找算法
func insert2(intervals [][]int, newInterval []int) [][]int {
	if len(intervals) == 0 {
		return [][]int{newInterval}
	}

	result := [][]int{}

	// 找到插入位置
	left := 0
	for left < len(intervals) && intervals[left][1] < newInterval[0] {
		result = append(result, intervals[left])
		left++
	}

	// 合并重叠区间
	for left < len(intervals) && intervals[left][0] <= newInterval[1] {
		newInterval[0] = min(newInterval[0], intervals[left][0])
		newInterval[1] = max(newInterval[1], intervals[left][1])
		left++
	}
	result = append(result, newInterval)

	// 添加剩余区间
	for left < len(intervals) {
		result = append(result, intervals[left])
		left++
	}

	return result
}

// 方法三：合并排序算法
func insert3(intervals [][]int, newInterval []int) [][]int {
	// 将新区间加入列表
	intervals = append(intervals, newInterval)

	// 排序所有区间
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	// 合并重叠区间
	result := [][]int{intervals[0]}
	for i := 1; i < len(intervals); i++ {
		last := result[len(result)-1]
		curr := intervals[i]

		if curr[0] <= last[1] {
			last[1] = max(last[1], curr[1])
		} else {
			result = append(result, curr)
		}
	}

	return result
}

// 方法四：分段处理算法
func insert4(intervals [][]int, newInterval []int) [][]int {
	if len(intervals) == 0 {
		return [][]int{newInterval}
	}

	left := [][]int{}
	right := [][]int{}
	start, end := newInterval[0], newInterval[1]

	// 分段处理
	for _, interval := range intervals {
		if interval[1] < start {
			// 左侧区间
			left = append(left, interval)
		} else if interval[0] > end {
			// 右侧区间
			right = append(right, interval)
		} else {
			// 重叠区间，更新边界
			start = min(start, interval[0])
			end = max(end, interval[1])
		}
	}

	// 拼接结果
	result := append(left, []int{start, end})
	result = append(result, right...)

	return result
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

// 测试用例
func createTestCases() []struct {
	intervals   [][]int
	newInterval []int
	expected    [][]int
	name        string
} {
	return []struct {
		intervals   [][]int
		newInterval []int
		expected    [][]int
		name        string
	}{
		{
			[][]int{{1, 3}, {6, 9}},
			[]int{2, 5},
			[][]int{{1, 5}, {6, 9}},
			"示例1: 基础插入",
		},
		{
			[][]int{{1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16}},
			[]int{4, 8},
			[][]int{{1, 2}, {3, 10}, {12, 16}},
			"示例2: 多区间合并",
		},
		{
			[][]int{},
			[]int{5, 7},
			[][]int{{5, 7}},
			"测试1: 空列表",
		},
		{
			[][]int{{1, 5}},
			[]int{2, 3},
			[][]int{{1, 5}},
			"测试2: 完全包含",
		},
		{
			[][]int{{1, 5}},
			[]int{6, 8},
			[][]int{{1, 5}, {6, 8}},
			"测试3: 右侧插入",
		},
		{
			[][]int{{3, 5}, {6, 9}},
			[]int{0, 2},
			[][]int{{0, 2}, {3, 5}, {6, 9}},
			"测试4: 左侧插入",
		},
		{
			[][]int{{1, 3}, {6, 9}},
			[]int{0, 10},
			[][]int{{0, 10}},
			"测试5: 完全覆盖",
		},
		{
			[][]int{{1, 2}, {3, 4}, {5, 6}},
			[]int{0, 7},
			[][]int{{0, 7}},
			"测试6: 全部合并",
		},
	}
}

// 性能测试
func benchmarkAlgorithm(algorithm func([][]int, []int) [][]int, intervals [][]int, newInterval []int, name string) {
	iterations := 1000
	start := time.Now()

	for i := 0; i < iterations; i++ {
		intervalsCopy := make([][]int, len(intervals))
		for j := range intervals {
			intervalsCopy[j] = make([]int, len(intervals[j]))
			copy(intervalsCopy[j], intervals[j])
		}
		algorithm(intervalsCopy, newInterval)
	}

	duration := time.Since(start)
	avgTime := duration.Nanoseconds() / int64(iterations)
	fmt.Printf("%s: 平均执行时间 %d 纳秒\n", name, avgTime)
}

// 打印区间
func printIntervals(intervals [][]int) {
	fmt.Print("[")
	for i, interval := range intervals {
		if i > 0 {
			fmt.Print(",")
		}
		fmt.Printf("[%d,%d]", interval[0], interval[1])
	}
	fmt.Print("]")
}

func main() {
	fmt.Println("=== 57. 插入区间 ===")
	fmt.Println()

	testCases := createTestCases()
	algorithms := []struct {
		name string
		fn   func([][]int, []int) [][]int
	}{
		{"线性扫描算法", insert1},
		{"二分查找算法", insert2},
		{"合并排序算法", insert3},
		{"分段处理算法", insert4},
	}

	// 正确性测试
	fmt.Println("=== 算法正确性测试 ===")
	for _, testCase := range testCases {
		fmt.Printf("测试: %s\n", testCase.name)

		results := make([][][]int, len(algorithms))
		for i, algo := range algorithms {
			intervalsCopy := make([][]int, len(testCase.intervals))
			for j := range testCase.intervals {
				intervalsCopy[j] = make([]int, len(testCase.intervals[j]))
				copy(intervalsCopy[j], testCase.intervals[j])
			}
			results[i] = algo.fn(intervalsCopy, testCase.newInterval)
		}

		allEqual := true
		for i := 1; i < len(results); i++ {
			if !reflect.DeepEqual(results[i], results[0]) {
				allEqual = false
				break
			}
		}

		allValid := reflect.DeepEqual(results[0], testCase.expected)

		if allEqual && allValid {
			fmt.Printf("  ✅ 所有算法结果一致且正确\n")
			fmt.Print("  输入: ")
			printIntervals(testCase.intervals)
			fmt.Printf(", 新区间: [%d,%d]\n", testCase.newInterval[0], testCase.newInterval[1])
			fmt.Print("  输出: ")
			printIntervals(results[0])
			fmt.Println()
		} else {
			fmt.Printf("  ❌ 算法结果不一致或错误\n")
			for i, algo := range algorithms {
				fmt.Printf("    %s: ", algo.name)
				printIntervals(results[i])
				fmt.Println()
			}
		}
		fmt.Println()
	}

	// 性能测试
	fmt.Println("=== 性能测试 ===")
	perfIntervals := [][]int{{1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16}}
	perfNew := []int{4, 8}

	for _, algo := range algorithms {
		benchmarkAlgorithm(algo.fn, perfIntervals, perfNew, algo.name)
	}
	fmt.Println()

	// 算法分析
	fmt.Println("=== 算法分析 ===")
	fmt.Println("插入区间问题的特点:")
	fmt.Println("1. 输入区间已排序")
	fmt.Println("2. 需要插入新区间并合并")
	fmt.Println("3. 分三段处理：左、中、右")
	fmt.Println("4. 线性扫描法是最优解法")
	fmt.Println()

	fmt.Println("=== 复杂度分析 ===")
	fmt.Println("时间复杂度:")
	fmt.Println("- 线性扫描: O(n)，一次遍历")
	fmt.Println("- 二分查找: O(n)，总体仍需遍历")
	fmt.Println("- 合并排序: O(n log n)，需要排序")
	fmt.Println("- 分段处理: O(n)，三段处理")
	fmt.Println()

	fmt.Println("=== 优化技巧总结 ===")
	fmt.Println("1. 三段式处理：左侧、合并、右侧")
	fmt.Println("2. 边界判断：准确判断区间关系")
	fmt.Println("3. 利用排序：已排序特性优化")
	fmt.Println("4. 一次遍历：提高时间效率")
}
