package main

import (
	"fmt"
)

// =========================== 方法一：双指针逆序遍历（最优解法） ===========================

func merge(nums1 []int, m int, nums2 []int, n int) {
	// 三个指针：i指向nums1有效元素末尾，j指向nums2末尾，k指向nums1末尾
	i, j, k := m-1, n-1, m+n-1

	// 从后往前填充nums1
	for k >= 0 {
		if j < 0 {
			// nums2已经全部处理完毕，停止
			break
		}
		if i >= 0 && nums1[i] > nums2[j] {
			nums1[k] = nums1[i]
			i--
		} else {
			nums1[k] = nums2[j]
			j--
		}
		k--
	}
}

// =========================== 方法二：双指针正序遍历（需要额外空间） ===========================

func merge2(nums1 []int, m int, nums2 []int, n int) {
	// 复制nums1的有效元素
	temp := make([]int, m)
	copy(temp, nums1[:m])

	i, j, k := 0, 0, 0

	// 从前往后填充nums1
	for i < m && j < n {
		if temp[i] <= nums2[j] {
			nums1[k] = temp[i]
			i++
		} else {
			nums1[k] = nums2[j]
			j++
		}
		k++
	}

	// 复制剩余元素
	for i < m {
		nums1[k] = temp[i]
		i++
		k++
	}
	for j < n {
		nums1[k] = nums2[j]
		j++
		k++
	}
}

// =========================== 方法三：简化版双指针 ===========================

func merge3(nums1 []int, m int, nums2 []int, n int) {
	i, j, k := m-1, n-1, m+n-1

	for j >= 0 {
		if i >= 0 && nums1[i] > nums2[j] {
			nums1[k] = nums1[i]
			i--
		} else {
			nums1[k] = nums2[j]
			j--
		}
		k--
	}
}

// =========================== 方法四：优化版（减少比较次数） ===========================

func merge4(nums1 []int, m int, nums2 []int, n int) {
	i, j, k := m-1, n-1, m+n-1

	for i >= 0 && j >= 0 {
		if nums1[i] > nums2[j] {
			nums1[k] = nums1[i]
			i--
		} else {
			nums1[k] = nums2[j]
			j--
		}
		k--
	}

	// 复制nums2的剩余元素
	for j >= 0 {
		nums1[k] = nums2[j]
		j--
		k--
	}
}

// =========================== 测试代码 ===========================

func main() {
	fmt.Println("=== LeetCode 88: 合并两个有序数组 ===\n")

	testCases := []struct {
		name     string
		nums1    []int
		m        int
		nums2    []int
		n        int
		expected []int
	}{
		{
			name:     "Test1: Basic case",
			nums1:    []int{1, 2, 3, 0, 0, 0},
			m:        3,
			nums2:    []int{2, 5, 6},
			n:        3,
			expected: []int{1, 2, 2, 3, 5, 6},
		},
		{
			name:     "Test2: nums2 is empty",
			nums1:    []int{1},
			m:        1,
			nums2:    []int{},
			n:        0,
			expected: []int{1},
		},
		{
			name:     "Test3: nums1 is empty",
			nums1:    []int{0},
			m:        0,
			nums2:    []int{1},
			n:        1,
			expected: []int{1},
		},
		{
			name:     "Test4: nums1 all less than nums2",
			nums1:    []int{1, 2, 3, 0, 0, 0},
			m:        3,
			nums2:    []int{4, 5, 6},
			n:        3,
			expected: []int{1, 2, 3, 4, 5, 6},
		},
		{
			name:     "Test5: nums2 all less than nums1",
			nums1:    []int{4, 5, 6, 0, 0, 0},
			m:        3,
			nums2:    []int{1, 2, 3},
			n:        3,
			expected: []int{1, 2, 3, 4, 5, 6},
		},
		{
			name:     "Test6: Single element in both",
			nums1:    []int{2, 0},
			m:        1,
			nums2:    []int{1},
			n:        1,
			expected: []int{1, 2},
		},
		{
			name:     "Test7: Empty arrays",
			nums1:    []int{0},
			m:        0,
			nums2:    []int{},
			n:        0,
			expected: []int{0},
		},
		{
			name:     "Test8: All same elements",
			nums1:    []int{1, 1, 1, 0, 0, 0},
			m:        3,
			nums2:    []int{1, 1, 1},
			n:        3,
			expected: []int{1, 1, 1, 1, 1, 1},
		},
	}

	methods := map[string]func([]int, int, []int, int){
		"双指针逆序遍历（最优解法）":   merge,
		"双指针正序遍历（需要额外空间）": merge2,
		"简化版双指针":          merge3,
		"优化版（减少比较次数）":     merge4,
	}

	for name, method := range methods {
		fmt.Printf("方法%s：%s\n", name, name)
		passCount := 0
		for i, tt := range testCases {
			// 复制输入数组，避免修改影响后续测试
			nums1Copy := make([]int, len(tt.nums1))
			copy(nums1Copy, tt.nums1)

			method(nums1Copy, tt.m, tt.nums2, tt.n)

			status := "✅"
			if !equal(nums1Copy, tt.expected) {
				status = "❌"
			} else {
				passCount++
			}
			fmt.Printf("  测试%d: %s\n", i+1, status)
			if status == "❌" {
				fmt.Printf("    输入: nums1=%v, m=%d, nums2=%v, n=%d\n", tt.nums1, tt.m, tt.nums2, tt.n)
				fmt.Printf("    输出: %v\n", nums1Copy)
				fmt.Printf("    期望: %v\n", tt.expected)
			}
		}
		fmt.Printf("  通过: %d/%d\n\n", passCount, len(testCases))
	}
}

func equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
