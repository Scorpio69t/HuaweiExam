package main

import (
	"fmt"
)

// searchInsertClassic 经典二分：寻找第一个 >= target 的位置
// 返回插入位置（若存在相等即返回其索引）
// 时间 O(log n)，空间 O(1)
func searchInsertClassic(nums []int, target int) int {
	left, right := 0, len(nums)
	for left < right {
		mid := left + (right-left)/2
		if nums[mid] >= target {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}

// searchInsertRightBias 右偏二分：寻找最后一个 < target 的位置，再 +1
// 等价于 lower_bound 的另一种写法
func searchInsertRightBias(nums []int, target int) int {
	left, right := 0, len(nums)-1
	ans := len(nums)
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] >= target {
			ans = mid
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return ans
}

// searchInsertLinear 基线法（仅用于对拍/小数据），O(n)
func searchInsertLinear(nums []int, target int) int {
	for i := 0; i < len(nums); i++ {
		if nums[i] >= target {
			return i
		}
	}
	return len(nums)
}

type method struct {
	name string
	fn   func([]int, int) int
}

func main() {
	tests := []struct {
		nums   []int
		target int
		expect int
	}{
		{[]int{1, 3, 5, 6}, 5, 2},
		{[]int{1, 3, 5, 6}, 2, 1},
		{[]int{1, 3, 5, 6}, 7, 4},
		{[]int{1, 3, 5, 6}, 0, 0},
		{[]int{1}, 0, 0},
		{[]int{1}, 1, 0},
		{[]int{1}, 2, 1},
		{[]int{}, 3, 0},
	}

	methods := []method{
		{"经典二分", searchInsertClassic},
		{"右偏二分", searchInsertRightBias},
		{"线性基线", searchInsertLinear},
	}

	fmt.Println("35. 搜索插入位置 - 多解法对比")
	for _, tc := range tests {
		fmt.Printf("nums=%v target=%d\n", tc.nums, tc.target)
		best := -1
		for _, m := range methods {
			got := m.fn(tc.nums, tc.target)
			if best == -1 {
				best = got
			}
			status := "✅"
			if got != tc.expect {
				status = "❌"
			}
			fmt.Printf("  %-6s => %d %s\n", m.name, got, status)
		}
		fmt.Printf("  期望 => %d\n", tc.expect)
		fmt.Println("------------------------------")
	}
}
