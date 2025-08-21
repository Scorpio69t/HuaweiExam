package main

import (
	"fmt"
	"sort"
)

// fourSum 经典排序 + 双指针解法
// 时间复杂度: O(n^3)
// 空间复杂度: O(1)（不计结果集）
func fourSum(nums []int, target int) [][]int {
	res := make([][]int, 0)
	if len(nums) < 4 {
		return res
	}

	sort.Ints(nums)
	n := len(nums)
	t := int64(target)

	for i := 0; i < n-3; i++ {
		// 去重: 固定 i
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		// 剪枝: 最小可能和 > target 或 最大可能和 < target
		minSum := int64(nums[i]) + int64(nums[i+1]) + int64(nums[i+2]) + int64(nums[i+3])
		if minSum > t {
			break
		}
		maxSum := int64(nums[i]) + int64(nums[n-1]) + int64(nums[n-2]) + int64(nums[n-3])
		if maxSum < t {
			continue
		}

		for j := i + 1; j < n-2; j++ {
			// 去重: 固定 j
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}
			// 剪枝
			min2 := int64(nums[i]) + int64(nums[j]) + int64(nums[j+1]) + int64(nums[j+2])
			if min2 > t {
				break
			}
			max2 := int64(nums[i]) + int64(nums[j]) + int64(nums[n-1]) + int64(nums[n-2])
			if max2 < t {
				continue
			}

			left, right := j+1, n-1
			for left < right {
				sum := int64(nums[i]) + int64(nums[j]) + int64(nums[left]) + int64(nums[right])
				if sum == t {
					res = append(res, []int{nums[i], nums[j], nums[left], nums[right]})
					// 去重移动
					lv, rv := nums[left], nums[right]
					for left < right && nums[left] == lv {
						left++
					}
					for left < right && nums[right] == rv {
						right--
					}
				} else if sum < t {
					left++
				} else {
					right--
				}
			}
		}
	}
	return res
}

// fourSumKSum 通用 kSum 解法入口（调用 kSum with k=4）
func fourSumKSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	return kSum(nums, 4, 0, int64(target))
}

// kSum 通用递归解法
// nums 已排序，寻找从 start 开始 k 个数之和为 target 的所有组合
func kSum(nums []int, k int, start int, target int64) [][]int {
	res := make([][]int, 0)
	n := len(nums)
	if k == 2 {
		// 2Sum 双指针
		l, r := start, n-1
		for l < r {
			s := int64(nums[l]) + int64(nums[r])
			if s == target {
				res = append(res, []int{nums[l], nums[r]})
				lv, rv := nums[l], nums[r]
				for l < r && nums[l] == lv {
					l++
				}
				for l < r && nums[r] == rv {
					r--
				}
			} else if s < target {
				l++
			} else {
				r--
			}
		}
		return res
	}

	// 剪枝: 若最小可能和或最大可能和不满足，直接返回
	if start >= n {
		return res
	}
	minSum := int64(0)
	for i := 0; i < k; i++ {
		if start+i >= n {
			return res
		}
		minSum += int64(nums[start+i])
	}
	maxSum := int64(0)
	for i := 0; i < k; i++ {
		if n-1-i < start {
			return res
		}
		maxSum += int64(nums[n-1-i])
	}
	if minSum > target || maxSum < target {
		return res
	}

	for i := start; i <= n-k; i++ {
		if i > start && nums[i] == nums[i-1] {
			continue
		}
		// 递归找 (k-1)Sum
		partial := kSum(nums, k-1, i+1, target-int64(nums[i]))
		for _, comb := range partial {
			res = append(res, append([]int{nums[i]}, comb...))
		}
	}
	return res
}

func main() {
	// 示例 1
	nums1 := []int{1, 0, -1, 0, -2, 2}
	target1 := 0
	ans1 := fourSum(nums1, target1)
	fmt.Printf("示例1(双指针): nums=%v target=%d\n结果: %v\n\n", nums1, target1, ans1)

	// 示例 2
	nums2 := []int{2, 2, 2, 2, 2}
	target2 := 8
	ans2 := fourSum(nums2, target2)
	fmt.Printf("示例2(双指针): nums=%v target=%d\n结果: %v\n\n", nums2, target2, ans2)

	// 通用 kSum 版本对比
	ans1k := fourSumKSum(nums1, target1)
	ans2k := fourSumKSum(nums2, target2)
	fmt.Printf("示例1(kSum): %v\n示例2(kSum): %v\n\n", ans1k, ans2k)

	// 其他用例
	nums3 := []int{0, 0, 0, 0}
	target3 := 0
	fmt.Printf("全零用例: %v\n结果: %v\n\n", nums3, fourSum(nums3, target3))

	nums4 := []int{-3, -1, 0, 2, 4, 5}
	target4 := 2
	fmt.Printf("混合用例: %v target=%d\n结果: %v\n", nums4, target4, fourSum(nums4, target4))
}
