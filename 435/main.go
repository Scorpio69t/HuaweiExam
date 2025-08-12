package main

import (
	"fmt"
	"sort"
	"strings"
)

// 解法一：按区间结束时间升序选择最多不重叠区间
// 需要移除的最小数量 = 总数 - 可选择的不重叠区间最大数量
// 时间复杂度：O(n log n)，空间复杂度：O(1)（排序除外）
func eraseOverlapIntervalsByEnd(intervals [][]int) int {
	n := len(intervals)
	if n <= 1 {
		return 0
	}

	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][1] == intervals[j][1] {
			return intervals[i][0] < intervals[j][0]
		}
		return intervals[i][1] < intervals[j][1]
	})

	countNonOverlap := 1
	prevEnd := intervals[0][1]
	for i := 1; i < n; i++ {
		if intervals[i][0] >= prevEnd {
			countNonOverlap++
			prevEnd = intervals[i][1]
		}
	}
	return n - countNonOverlap
}

// 解法二：按起点升序遍历，遇到重叠时移除“结束更晚”的那个
// 直接统计“移除次数”。与解法一等价，常见贪心写法。
// 时间复杂度：O(n log n)，空间复杂度：O(1)（排序除外）
func eraseOverlapIntervalsByRemoving(intervals [][]int) int {
	n := len(intervals)
	if n <= 1 {
		return 0
	}

	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] == intervals[j][0] {
			return intervals[i][1] < intervals[j][1]
		}
		return intervals[i][0] < intervals[j][0]
	})

	removals := 0
	prevEnd := intervals[0][1]
	for i := 1; i < n; i++ {
		// 有重叠：当前区间起点 < 上一个保留区间的结束
		if intervals[i][0] < prevEnd {
			removals++
			// 移除结束更晚的那个，相当于把 prevEnd 缩到更小的结束点
			if intervals[i][1] < prevEnd {
				prevEnd = intervals[i][1]
			}
		} else {
			// 无重叠，更新 prevEnd
			prevEnd = intervals[i][1]
		}
	}
	return removals
}

func runTests() {
	type testCase struct {
		intervals [][]int
		expected  int
		desc      string
	}

	tests := []testCase{
		{intervals: [][]int{{1, 2}, {2, 3}, {3, 4}, {1, 3}}, expected: 1, desc: "示例1"},
		{intervals: [][]int{{1, 2}, {1, 2}, {1, 2}}, expected: 2, desc: "示例2"},
		{intervals: [][]int{{1, 2}, {2, 3}}, expected: 0, desc: "示例3"},
		{intervals: [][]int{{1, 100}, {11, 22}, {1, 11}, {2, 12}}, expected: 2, desc: "覆盖大区间"},
		{intervals: [][]int{{0, 1}}, expected: 0, desc: "单区间"},
		{intervals: [][]int{}, expected: 0, desc: "空数组"},
		{intervals: [][]int{{1, 3}, {2, 4}, {3, 5}, {6, 7}}, expected: 1, desc: "部分重叠"},
		{intervals: [][]int{{-5, -1}, {-2, 2}, {2, 3}, {3, 4}}, expected: 1, desc: "含负数与接触"},
	}

	fmt.Println("=== 435. 无重叠区间 - 测试 ===")
	for i, tc := range tests {
		r1 := eraseOverlapIntervalsByEnd(cloneIntervals(tc.intervals))
		r2 := eraseOverlapIntervalsByRemoving(cloneIntervals(tc.intervals))
		ok := (r1 == tc.expected) && (r2 == tc.expected)
		status := "✅"
		if !ok {
			status = "❌"
		}
		fmt.Printf("用例 %d: %s\n", i+1, tc.desc)
		fmt.Printf("输入: %v\n期望: %d\n按结束贪心: %d, 按起点移除: %d\n", tc.intervals, tc.expected, r1, r2)
		fmt.Printf("结果: %s\n", status)
		fmt.Println(strings.Repeat("-", 40))
	}
}

func cloneIntervals(src [][]int) [][]int {
	dst := make([][]int, len(src))
	for i := range src {
		if src[i] == nil {
			continue
		}
		dst[i] = make([]int, len(src[i]))
		copy(dst[i], src[i])
	}
	return dst
}

func main() {
	runTests()
}
