package main

import "fmt"

// =========================== 方法一：三指针（荷兰国旗，最优） ===========================

func sortColors(nums []int) {
	left, cur, right := 0, 0, len(nums)-1

	for cur <= right {
		if nums[cur] == 0 {
			// 0应该在左边，与left交换
			nums[left], nums[cur] = nums[cur], nums[left]
			left++
			cur++
		} else if nums[cur] == 1 {
			// 1在中间，cur继续前进
			cur++
		} else {
			// 2应该在右边，与right交换
			nums[cur], nums[right] = nums[right], nums[cur]
			right--
			// 注意：cur不动，因为交换来的值还未检查
		}
	}
}

// =========================== 方法二：计数排序 ===========================

func sortColors2(nums []int) {
	count0, count1, count2 := 0, 0, 0

	// 第一次遍历：统计
	for _, num := range nums {
		if num == 0 {
			count0++
		} else if num == 1 {
			count1++
		} else {
			count2++
		}
	}

	// 第二次遍历：重建
	i := 0
	for count0 > 0 {
		nums[i] = 0
		i++
		count0--
	}
	for count1 > 0 {
		nums[i] = 1
		i++
		count1--
	}
	for count2 > 0 {
		nums[i] = 2
		i++
		count2--
	}
}

// =========================== 方法三：双指针 ===========================

func sortColors3(nums []int) {
	n := len(nums)
	left, right := 0, n-1

	for i := 0; i <= right; {
		if nums[i] == 0 {
			nums[i], nums[left] = nums[left], nums[i]
			left++
			i++
		} else if nums[i] == 2 {
			nums[i], nums[right] = nums[right], nums[i]
			right--
		} else {
			i++
		}
	}
}

// =========================== 方法四：快速排序变体 ===========================

func sortColors4(nums []int) {
	if len(nums) <= 1 {
		return
	}
	quickSort(nums, 0, len(nums)-1)
}

func quickSort(nums []int, left, right int) {
	if left >= right {
		return
	}

	// 三路快排，pivot=1
	lt, gt, i := left, right, left
	pivot := 1

	for i <= gt {
		if nums[i] < pivot {
			nums[lt], nums[i] = nums[i], nums[lt]
			lt++
			i++
		} else if nums[i] > pivot {
			nums[i], nums[gt] = nums[gt], nums[i]
			gt--
		} else {
			i++
		}
	}
}

// =========================== 测试代码 ===========================

func main() {
	fmt.Println("=== LeetCode 75: 颜色分类（荷兰国旗问题） ===\n")

	testCases := []struct {
		nums   []int
		expect []int
	}{
		{
			[]int{2, 0, 2, 1, 1, 0},
			[]int{0, 0, 1, 1, 2, 2},
		},
		{
			[]int{2, 0, 1},
			[]int{0, 1, 2},
		},
		{
			[]int{0},
			[]int{0},
		},
		{
			[]int{1, 0},
			[]int{0, 1},
		},
		{
			[]int{1, 1, 1},
			[]int{1, 1, 1},
		},
		{
			[]int{2, 1, 0},
			[]int{0, 1, 2},
		},
		{
			[]int{0, 0, 2, 2},
			[]int{0, 0, 2, 2},
		},
	}

	fmt.Println("方法一：三指针（荷兰国旗）")
	runTests(testCases, sortColors)

	fmt.Println("\n方法二：计数排序")
	runTests(testCases, sortColors2)

	fmt.Println("\n方法三：双指针")
	runTests(testCases, sortColors3)

	fmt.Println("\n方法四：快速排序变体")
	runTests(testCases, sortColors4)
}

func runTests(testCases []struct {
	nums   []int
	expect []int
}, fn func([]int)) {
	passCount := 0
	for i, tc := range testCases {
		nums := make([]int, len(tc.nums))
		copy(nums, tc.nums)
		fn(nums)

		status := "✅"
		if !equal(nums, tc.expect) {
			status = "❌"
		} else {
			passCount++
		}
		fmt.Printf("  测试%d: %s\n", i+1, status)
		if status == "❌" {
			fmt.Printf("    输入: %v\n", tc.nums)
			fmt.Printf("    输出: %v\n", nums)
			fmt.Printf("    期望: %v\n", tc.expect)
		}
	}
	fmt.Printf("  通过: %d/%d\n", passCount, len(testCases))
}

func equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
