package main

import (
	"fmt"
)

// nextPermutation 原地计算下一个排列（字典序更大），若不存在则重排为最小（升序）
// 时间 O(n)，空间 O(1)
func nextPermutation(nums []int) {
	n := len(nums)
	if n <= 1 {
		return
	}

	// 1) 从右向左找到首个下降点 i: nums[i] < nums[i+1]
	i := n - 2
	for i >= 0 && nums[i] >= nums[i+1] {
		i--
	}

	if i >= 0 {
		// 2) 从右向左找到第一个大于 nums[i] 的 j
		j := n - 1
		for j > i && nums[j] <= nums[i] {
			j--
		}
		// 3) 交换 i, j
		nums[i], nums[j] = nums[j], nums[i]
	}

	// 4) 反转 i+1..n-1，使尾部最小化（升序）
	reverse(nums, i+1, n-1)
}

func reverse(a []int, l, r int) {
	for l < r {
		a[l], a[r] = a[r], a[l]
		l++
		r--
	}
}

func clone(src []int) []int {
	dst := make([]int, len(src))
	copy(dst, src)
	return dst
}

func main() {
	fmt.Println("下一个排列 测试")
	fmt.Println("================")

	cases := [][]int{
		{1, 2, 3},
		{3, 2, 1},
		{1, 1, 5},
		{1, 3, 2},
		{2, 2, 0, 4, 3, 1},
	}

	for idx, c := range cases {
		arr := clone(c)
		nextPermutation(arr)
		fmt.Printf("用例%d: in=%v out=%v\n", idx+1, c, arr)
	}
}
