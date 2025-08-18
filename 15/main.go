package main

import (
	"fmt"
	"sort"
)

// 三数之和 - 排序 + 双指针（推荐）
// 时间复杂度：O(n^2)，空间复杂度：O(1) 额外空间（不计结果集）
func threeSum(nums []int) [][]int {
	res := [][]int{}
	n := len(nums)
	if n < 3 {
		return res
	}

	sort.Ints(nums)
	for i := 0; i < n-2; i++ {
		// 去重：固定数与前一个相同则跳过
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		// 剪枝：若最小的固定数已大于0，则不可能再找到和为0
		if nums[i] > 0 {
			break
		}

		l, r := i+1, n-1
		for l < r {
			s := nums[i] + nums[l] + nums[r]
			if s == 0 {
				res = append(res, []int{nums[i], nums[l], nums[r]})
				// 移动指针并跳过重复值
				l++
				for l < r && nums[l] == nums[l-1] {
					l++
				}
				r--
				for l < r && nums[r] == nums[r+1] {
					r--
				}
			} else if s < 0 {
				l++
			} else {
				r--
			}
		}
	}
	return res
}

// 备用思路：固定一个数 + 哈希集合查找 two-sum 的补数，同时控制去重
// 为演示而保留，复杂度同样 O(n^2)，实现上不如双指针直观
func threeSumWithHash(nums []int) [][]int {
	res := [][]int{}
	n := len(nums)
	if n < 3 {
		return res
	}
	sort.Ints(nums)
	for i := 0; i < n-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		target := -nums[i]
		seen := make(map[int]bool)
		// 为防重：记录已经加入结果的第二个数
		usedSecond := make(map[int]bool)
		for j := i + 1; j < n; j++ {
			need := target - nums[j]
			if seen[need] {
				// 三元组为 nums[i], need, nums[j]，此时 need <= nums[j]
				if !usedSecond[need] { // 控制去重
					res = append(res, []int{nums[i], need, nums[j]})
					usedSecond[need] = true
				}
			}
			seen[nums[j]] = true
		}
	}
	return res
}

// 简易校验：判断所有三元组和为 0，且三元组内部为非降序，整体无重复
func validateTriplets(ans [][]int) bool {
	seen := make(map[[3]int]bool)
	for _, t := range ans {
		if len(t) != 3 {
			return false
		}
		a, b, c := t[0], t[1], t[2]
		if a+b+c != 0 {
			return false
		}
		if !(a <= b && b <= c) {
			return false
		}
		key := [3]int{a, b, c}
		if seen[key] {
			return false
		}
		seen[key] = true
	}
	return true
}

func main() {
	fmt.Println("=== 15. 三数之和 ===")

	cases := []struct {
		nums []int
		name string
	}{
		{[]int{-1, 0, 1, 2, -1, -4}, "示例1"},
		{[]int{0, 1, 1}, "示例2"},
		{[]int{0, 0, 0}, "示例3"},
		{[]int{0, 0, 0, 0}, "多个零"},
		{[]int{-2, 0, 1, 1, 2}, "包含重复与正负"},
		{[]int{3, -2, 1, 0}, "无解"},
	}

	for _, c := range cases {
		fmt.Printf("%s: %v\n", c.name, c.nums)
		ans := threeSum(append([]int(nil), c.nums...))
		fmt.Printf("双指针解法结果: %v\n", ans)
		fmt.Printf("结果校验: %v\n\n", validateTriplets(ans))
	}

	// 可选：对比哈希解法输出（仅作为参考）
	extra := []int{-1, 0, 1, 2, -1, -4}
	fmt.Printf("对比（哈希解法）: %v\n", threeSumWithHash(append([]int(nil), extra...)))
}
