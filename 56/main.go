package main

import (
	"fmt"
	"reflect"
	"sort"
	"time"
)

// 方法一：排序合并算法
// 最优解法，按起始位置排序后合并重叠区间
func merge1(intervals [][]int) [][]int {
	if len(intervals) <= 1 {
		return intervals
	}

	// 按起始位置排序
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	result := [][]int{intervals[0]}

	for i := 1; i < len(intervals); i++ {
		last := result[len(result)-1]
		curr := intervals[i]

		// 判断是否重叠
		if curr[0] <= last[1] {
			// 合并区间，更新结束位置
			last[1] = max(last[1], curr[1])
		} else {
			// 添加新区间
			result = append(result, curr)
		}
	}

	return result
}

// 方法二：栈优化算法
// 使用栈辅助合并区间
func merge2(intervals [][]int) [][]int {
	if len(intervals) <= 1 {
		return intervals
	}

	// 按起始位置排序
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	stack := [][]int{intervals[0]}

	for i := 1; i < len(intervals); i++ {
		top := stack[len(stack)-1]
		curr := intervals[i]

		if curr[0] <= top[1] {
			// 合并区间
			stack[len(stack)-1][1] = max(top[1], curr[1])
		} else {
			// 添加新区间
			stack = append(stack, curr)
		}
	}

	return stack
}

// 方法三：并查集算法
// 使用并查集存储重叠关系
type UnionFind struct {
	parent []int
}

func NewUnionFind(n int) *UnionFind {
	parent := make([]int, n)
	for i := range parent {
		parent[i] = i
	}
	return &UnionFind{parent}
}

func (uf *UnionFind) Find(x int) int {
	if uf.parent[x] != x {
		uf.parent[x] = uf.Find(uf.parent[x])
	}
	return uf.parent[x]
}

func (uf *UnionFind) Union(x, y int) {
	uf.parent[uf.Find(x)] = uf.Find(y)
}

func merge3(intervals [][]int) [][]int {
	n := len(intervals)
	if n <= 1 {
		return intervals
	}

	uf := NewUnionFind(n)

	// 检查所有区间对是否重叠
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if isOverlap(intervals[i], intervals[j]) {
				uf.Union(i, j)
			}
		}
	}

	// 合并同一组的区间
	groups := make(map[int][][]int)
	for i := 0; i < n; i++ {
		root := uf.Find(i)
		groups[root] = append(groups[root], intervals[i])
	}

	result := [][]int{}
	for _, group := range groups {
		merged := mergeGroup(group)
		result = append(result, merged)
	}

	// 排序结果
	sort.Slice(result, func(i, j int) bool {
		return result[i][0] < result[j][0]
	})

	return result
}

func isOverlap(a, b []int) bool {
	return !(a[1] < b[0] || b[1] < a[0])
}

func mergeGroup(intervals [][]int) []int {
	minStart := intervals[0][0]
	maxEnd := intervals[0][1]

	for _, interval := range intervals {
		if interval[0] < minStart {
			minStart = interval[0]
		}
		if interval[1] > maxEnd {
			maxEnd = interval[1]
		}
	}

	return []int{minStart, maxEnd}
}

// 方法四：扫描线算法
// 使用事件扫描线合并区间
func merge4(intervals [][]int) [][]int {
	if len(intervals) <= 1 {
		return intervals
	}

	// 创建事件列表
	type Event struct {
		pos int
		typ int // 1: 起始, -1: 结束
		idx int // 结束事件的索引
	}

	events := []Event{}
	for i, interval := range intervals {
		events = append(events, Event{interval[0], 1, i})
		events = append(events, Event{interval[1], -1, i})
	}

	// 按位置排序，位置相同时起始事件优先
	sort.Slice(events, func(i, j int) bool {
		if events[i].pos == events[j].pos {
			return events[i].typ > events[j].typ
		}
		return events[i].pos < events[j].pos
	})

	result := [][]int{}
	count := 0
	start := 0

	for _, event := range events {
		if count == 0 {
			start = event.pos
		}

		count += event.typ

		if count == 0 {
			result = append(result, []int{start, event.pos})
		}
	}

	return result
}

// 辅助函数：求两个数的最大值
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 辅助函数：创建测试用例
func createTestCases() []struct {
	intervals [][]int
	expected  [][]int
	name      string
} {
	return []struct {
		intervals [][]int
		expected  [][]int
		name      string
	}{
		{
			[][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}},
			[][]int{{1, 6}, {8, 10}, {15, 18}},
			"示例1: 基础区间合并",
		},
		{
			[][]int{{1, 4}, {4, 5}},
			[][]int{{1, 5}},
			"示例2: 相邻区间",
		},
		{
			[][]int{{4, 7}, {1, 4}},
			[][]int{{1, 7}},
			"示例3: 无序区间",
		},
		{
			[][]int{{1, 3}},
			[][]int{{1, 3}},
			"测试1: 单个区间",
		},
		{
			[][]int{{1, 10}, {2, 5}, {3, 7}},
			[][]int{{1, 10}},
			"测试2: 完全重叠",
		},
		{
			[][]int{{1, 2}, {3, 4}, {5, 6}},
			[][]int{{1, 2}, {3, 4}, {5, 6}},
			"测试3: 完全不重叠",
		},
		{
			[][]int{{1, 2}, {2, 3}, {3, 4}},
			[][]int{{1, 4}},
			"测试4: 连续相邻",
		},
		{
			[][]int{{1, 4}, {0, 4}},
			[][]int{{0, 4}},
			"测试5: 包含关系",
		},
		{
			[][]int{{1, 4}, {2, 3}},
			[][]int{{1, 4}},
			"测试6: 内部包含",
		},
		{
			[][]int{{1, 4}, {0, 0}},
			[][]int{{0, 0}, {1, 4}},
			"测试7: 单点区间",
		},
	}
}

// 性能测试函数
func benchmarkAlgorithm(algorithm func([][]int) [][]int, intervals [][]int, name string) {
	iterations := 1000
	start := time.Now()

	for i := 0; i < iterations; i++ {
		// 需要复制数组，因为某些算法会修改输入
		intervalsCopy := make([][]int, len(intervals))
		for j := range intervals {
			intervalsCopy[j] = make([]int, len(intervals[j]))
			copy(intervalsCopy[j], intervals[j])
		}
		algorithm(intervalsCopy)
	}

	duration := time.Since(start)
	avgTime := duration.Nanoseconds() / int64(iterations)

	fmt.Printf("%s: 平均执行时间 %d 纳秒\n", name, avgTime)
}

// 辅助函数：打印区间数组
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
	fmt.Println("=== 56. 合并区间 ===")
	fmt.Println()

	// 创建测试用例
	testCases := createTestCases()
	algorithms := []struct {
		name string
		fn   func([][]int) [][]int
	}{
		{"排序合并算法", merge1},
		{"栈优化算法", merge2},
		{"并查集算法", merge3},
		{"扫描线算法", merge4},
	}

	// 运行测试
	fmt.Println("=== 算法正确性测试 ===")
	for _, testCase := range testCases {
		fmt.Printf("测试: %s\n", testCase.name)

		results := make([][][]int, len(algorithms))
		for i, algo := range algorithms {
			// 复制输入数组
			intervalsCopy := make([][]int, len(testCase.intervals))
			for j := range testCase.intervals {
				intervalsCopy[j] = make([]int, len(testCase.intervals[j]))
				copy(intervalsCopy[j], testCase.intervals[j])
			}
			results[i] = algo.fn(intervalsCopy)
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
			fmt.Print("  输入区间: ")
			printIntervals(testCase.intervals)
			fmt.Println()
			fmt.Print("  输出结果: ")
			printIntervals(results[0])
			fmt.Println()
		} else {
			fmt.Printf("  ❌ 算法结果不一致或错误\n")
			fmt.Print("  输入区间: ")
			printIntervals(testCase.intervals)
			fmt.Println()
			fmt.Print("  预期结果: ")
			printIntervals(testCase.expected)
			fmt.Println()
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
	performanceIntervals := [][]int{
		{1, 3}, {2, 6}, {8, 10}, {15, 18}, {5, 7}, {9, 12}, {14, 16},
		{11, 13}, {4, 8}, {17, 20}, {19, 22}, {21, 25}, {23, 27},
	}

	fmt.Printf("测试数据: %d个区间\n", len(performanceIntervals))
	fmt.Println()

	for _, algo := range algorithms {
		benchmarkAlgorithm(algo.fn, performanceIntervals, algo.name)
	}
	fmt.Println()

	// 算法分析
	fmt.Println("=== 算法分析 ===")
	fmt.Println("合并区间问题的特点:")
	fmt.Println("1. 需要合并所有重叠的区间")
	fmt.Println("2. 区间重叠判断：当前起始 <= 前一个结束")
	fmt.Println("3. 排序是关键步骤")
	fmt.Println("4. 排序合并法是最优解法")
	fmt.Println()

	// 复杂度分析
	fmt.Println("=== 复杂度分析 ===")
	fmt.Println("时间复杂度:")
	fmt.Println("- 排序合并: O(n log n)，排序O(n log n)+遍历O(n)")
	fmt.Println("- 栈优化: O(n log n)，排序O(n log n)+遍历O(n)")
	fmt.Println("- 并查集: O(n²)，需要检查所有区间对")
	fmt.Println("- 扫描线: O(n log n)，排序O(n log n)+遍历O(n)")
	fmt.Println()
	fmt.Println("空间复杂度:")
	fmt.Println("- 排序合并: O(log n)，排序的递归栈空间")
	fmt.Println("- 栈优化: O(n)，需要栈存储区间")
	fmt.Println("- 并查集: O(n)，需要并查集结构")
	fmt.Println("- 扫描线: O(n)，需要事件列表")
	fmt.Println()

	// 算法总结
	fmt.Println("=== 算法总结 ===")
	fmt.Println("1. 排序合并算法：最优解法，逻辑清晰")
	fmt.Println("2. 栈优化算法：代码简洁，易于理解")
	fmt.Println("3. 并查集算法：适合动态添加区间")
	fmt.Println("4. 扫描线算法：适合复杂区间问题")
	fmt.Println()
	fmt.Println("推荐使用：排序合并算法（方法一），效率最高")
	fmt.Println()

	// 应用场景
	fmt.Println("=== 应用场景 ===")
	fmt.Println("- 算法竞赛：区间合并的经典应用")
	fmt.Println("- 日程安排：合并重叠的时间段")
	fmt.Println("- 资源分配：合并重叠的资源占用")
	fmt.Println("- 数据压缩：合并连续的数据段")
	fmt.Println("- 系统设计：合并重叠的请求时间窗口")
	fmt.Println()

	// 优化技巧总结
	fmt.Println("=== 优化技巧总结 ===")
	fmt.Println("1. 排序：掌握按起始位置排序的技巧")
	fmt.Println("2. 合并判断：理解重叠判断的条件")
	fmt.Println("3. 边界更新：学会更新合并后的边界")
	fmt.Println("4. 边界处理：注意各种边界情况")
	fmt.Println("5. 算法选择：根据问题特点选择合适的算法")
	fmt.Println("6. 优化策略：学会时间和空间优化技巧")
}
