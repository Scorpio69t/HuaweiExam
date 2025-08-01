package main

import (
	"fmt"
	"sort"
	"strings"
	"time"
)

// 解法一：双重二分查找法（推荐解法）
// 时间复杂度：O(log n)，空间复杂度：O(1)
func searchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}

	// 查找左边界
	leftBound := findLeftBound(nums, target)
	if leftBound == -1 {
		return []int{-1, -1}
	}

	// 查找右边界
	rightBound := findRightBound(nums, target)

	return []int{leftBound, rightBound}
}

// 查找左边界：第一个大于等于target的位置
func findLeftBound(nums []int, target int) int {
	left, right := 0, len(nums)

	for left < right {
		mid := left + (right-left)/2
		if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid
		}
	}

	// 检查找到的位置是否有效
	if left < len(nums) && nums[left] == target {
		return left
	}
	return -1
}

// 查找右边界：最后一个小于等于target的位置
func findRightBound(nums []int, target int) int {
	left, right := 0, len(nums)

	for left < right {
		mid := left + (right-left)/2
		if nums[mid] <= target {
			left = mid + 1
		} else {
			right = mid
		}
	}

	// left-1是最后一个小于等于target的位置
	return left - 1
}

// 解法二：优化的双重二分查找（更清晰的边界处理）
// 时间复杂度：O(log n)，空间复杂度：O(1)
func searchRangeOptimized(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}

	// 使用统一的二分查找模板
	leftBound := binarySearchLeft(nums, target)
	rightBound := binarySearchRight(nums, target)

	if leftBound <= rightBound {
		return []int{leftBound, rightBound}
	}
	return []int{-1, -1}
}

// 二分查找左边界（左闭右闭区间）
func binarySearchLeft(nums []int, target int) int {
	left, right := 0, len(nums)-1

	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] >= target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	// 检查边界和目标值
	if left < len(nums) && nums[left] == target {
		return left
	}
	return len(nums) // 表示未找到
}

// 二分查找右边界（左闭右闭区间）
func binarySearchRight(nums []int, target int) int {
	left, right := 0, len(nums)-1

	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] <= target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	// 检查边界和目标值
	if right >= 0 && nums[right] == target {
		return right
	}
	return -1 // 表示未找到
}

// 解法三：单次二分查找法
// 时间复杂度：O(log n)，空间复杂度：O(1)
func searchRangeSingle(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}

	// 先用标准二分查找找到任意一个target位置
	pos := binarySearch(nums, target)
	if pos == -1 {
		return []int{-1, -1}
	}

	// 从找到的位置向两边扩展
	left, right := pos, pos

	// 向左扩展找到左边界
	for left > 0 && nums[left-1] == target {
		left--
	}

	// 向右扩展找到右边界
	for right < len(nums)-1 && nums[right+1] == target {
		right++
	}

	return []int{left, right}
}

// 标准二分查找
func binarySearch(nums []int, target int) int {
	left, right := 0, len(nums)-1

	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return -1
}

// 解法四：使用Go标准库
// 时间复杂度：O(log n)，空间复杂度：O(1)
func searchRangeStdLib(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}

	// 使用sort.SearchInts查找左边界
	leftBound := sort.SearchInts(nums, target)

	// 检查是否找到目标值
	if leftBound >= len(nums) || nums[leftBound] != target {
		return []int{-1, -1}
	}

	// 查找右边界：第一个大于target的位置减1
	rightBound := sort.SearchInts(nums, target+1) - 1

	return []int{leftBound, rightBound}
}

// 解法五：线性查找法（不满足时间复杂度要求，仅用于对比）
// 时间复杂度：O(n)，空间复杂度：O(1)
func searchRangeLinear(nums []int, target int) []int {
	left, right := -1, -1

	// 从左到右找第一个目标值
	for i := 0; i < len(nums); i++ {
		if nums[i] == target {
			left = i
			break
		}
	}

	if left == -1 {
		return []int{-1, -1}
	}

	// 从右到左找最后一个目标值
	for i := len(nums) - 1; i >= 0; i-- {
		if nums[i] == target {
			right = i
			break
		}
	}

	return []int{left, right}
}

// 解法六：递归二分查找
// 时间复杂度：O(log n)，空间复杂度：O(log n)
func searchRangeRecursive(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}

	left := findLeftBoundRecursive(nums, target, 0, len(nums)-1)
	if left == -1 {
		return []int{-1, -1}
	}

	right := findRightBoundRecursive(nums, target, 0, len(nums)-1)

	return []int{left, right}
}

// 递归查找左边界
func findLeftBoundRecursive(nums []int, target, left, right int) int {
	if left > right {
		return -1
	}

	mid := left + (right-left)/2

	if nums[mid] == target {
		// 检查是否是最左边的
		if mid == 0 || nums[mid-1] != target {
			return mid
		}
		return findLeftBoundRecursive(nums, target, left, mid-1)
	} else if nums[mid] < target {
		return findLeftBoundRecursive(nums, target, mid+1, right)
	} else {
		return findLeftBoundRecursive(nums, target, left, mid-1)
	}
}

// 递归查找右边界
func findRightBoundRecursive(nums []int, target, left, right int) int {
	if left > right {
		return -1
	}

	mid := left + (right-left)/2

	if nums[mid] == target {
		// 检查是否是最右边的
		if mid == len(nums)-1 || nums[mid+1] != target {
			return mid
		}
		return findRightBoundRecursive(nums, target, mid+1, right)
	} else if nums[mid] < target {
		return findRightBoundRecursive(nums, target, mid+1, right)
	} else {
		return findRightBoundRecursive(nums, target, left, mid-1)
	}
}

// 测试函数
func testSearchRange() {
	testCases := []struct {
		nums     []int
		target   int
		expected []int
		desc     string
	}{
		{[]int{5, 7, 7, 8, 8, 10}, 8, []int{3, 4}, "示例1：目标存在多个"},
		{[]int{5, 7, 7, 8, 8, 10}, 6, []int{-1, -1}, "示例2：目标不存在"},
		{[]int{}, 0, []int{-1, -1}, "示例3：空数组"},
		{[]int{1}, 1, []int{0, 0}, "单元素匹配"},
		{[]int{1}, 2, []int{-1, -1}, "单元素不匹配"},
		{[]int{1, 1, 1, 1, 1}, 1, []int{0, 4}, "全部相同元素"},
		{[]int{1, 2, 3, 4, 5}, 1, []int{0, 0}, "目标在首位"},
		{[]int{1, 2, 3, 4, 5}, 5, []int{4, 4}, "目标在末位"},
		{[]int{1, 2, 3, 4, 5}, 3, []int{2, 2}, "目标在中间单个"},
		{[]int{1, 2, 2, 2, 3}, 2, []int{1, 3}, "目标在中间多个"},
		{[]int{1, 3, 5, 7, 9}, 4, []int{-1, -1}, "目标在间隙"},
		{[]int{1, 1, 2, 2, 3, 3}, 2, []int{2, 3}, "连续重复"},
		{[]int{-1, 0, 3, 5, 9, 12}, 9, []int{4, 4}, "包含负数"},
		{[]int{-3, -1, 0, 3, 5}, -1, []int{1, 1}, "负数目标"},
		{[]int{0, 0, 0, 1, 1, 1}, 0, []int{0, 2}, "零值连续"},
	}

	fmt.Println("=== 查找元素范围测试 ===\n")

	for i, tc := range testCases {
		// 测试主要解法
		result1 := searchRange(tc.nums, tc.target)
		result2 := searchRangeOptimized(tc.nums, tc.target)
		result3 := searchRangeStdLib(tc.nums, tc.target)

		status := "✅"
		if !equalSlices(result1, tc.expected) {
			status = "❌"
		}

		fmt.Printf("测试 %d: %s\n", i+1, tc.desc)
		fmt.Printf("输入: nums=%v, target=%d\n", tc.nums, tc.target)
		fmt.Printf("期望: %v\n", tc.expected)
		fmt.Printf("双重二分: %v\n", result1)
		fmt.Printf("优化二分: %v\n", result2)
		fmt.Printf("标准库法: %v\n", result3)
		fmt.Printf("结果: %s\n", status)
		fmt.Println(strings.Repeat("-", 40))
	}
}

// 辅助函数：比较两个切片是否相等
func equalSlices(a, b []int) bool {
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

// 性能测试
func benchmarkSearchRange() {
	fmt.Println("\n=== 性能测试 ===\n")

	// 构造测试数据
	testData := []struct {
		nums   []int
		target int
		desc   string
	}{
		{generateSortedArray(1000, 5), 5, "1000元素数组"},
		{generateSortedArray(10000, 50), 50, "10000元素数组"},
		{generateSortedArray(100000, 500), 500, "100000元素数组"},
		{generateRepeatedArray(50000, 42), 42, "50000重复元素"},
	}

	algorithms := []struct {
		name string
		fn   func([]int, int) []int
	}{
		{"双重二分", searchRange},
		{"优化二分", searchRangeOptimized},
		{"标准库法", searchRangeStdLib},
		{"递归二分", searchRangeRecursive},
		{"线性查找", searchRangeLinear},
	}

	for _, data := range testData {
		fmt.Printf("%s:\n", data.desc)

		for _, algo := range algorithms {
			start := time.Now()
			result := algo.fn(data.nums, data.target)
			duration := time.Since(start)

			fmt.Printf("  %s: %v, 耗时: %v\n", algo.name, result, duration)
		}
		fmt.Println()
	}
}

// 生成有序数组（包含重复元素）
func generateSortedArray(size, targetCount int) []int {
	nums := make([]int, size)
	target := size / 2

	for i := 0; i < size; i++ {
		if i >= target && i < target+targetCount {
			nums[i] = target
		} else if i < target {
			nums[i] = i
		} else {
			nums[i] = i - targetCount + 1
		}
	}

	return nums
}

// 生成重复元素数组
func generateRepeatedArray(size, value int) []int {
	nums := make([]int, size)
	for i := range nums {
		if i < size/3 {
			nums[i] = value - 1
		} else if i < 2*size/3 {
			nums[i] = value
		} else {
			nums[i] = value + 1
		}
	}
	return nums
}

// 演示二分查找过程
func demonstrateBinarySearch() {
	fmt.Println("\n=== 二分查找过程演示 ===")
	nums := []int{5, 7, 7, 8, 8, 10}
	target := 8

	fmt.Printf("数组: %v, 目标: %d\n", nums, target)
	fmt.Println("\n查找左边界过程:")
	demonstrateLeftBound(nums, target)

	fmt.Println("\n查找右边界过程:")
	demonstrateRightBound(nums, target)

	result := searchRange(nums, target)
	fmt.Printf("\n最终结果: %v\n", result)
}

func demonstrateLeftBound(nums []int, target int) {
	left, right := 0, len(nums)
	step := 1

	fmt.Printf("初始: left=%d, right=%d\n", left, right)

	for left < right {
		mid := left + (right-left)/2
		fmt.Printf("步骤%d: left=%d, right=%d, mid=%d, nums[%d]=%d\n",
			step, left, right, mid, mid, nums[mid])

		if nums[mid] < target {
			left = mid + 1
			fmt.Printf("       nums[%d]=%d < %d, left=%d\n", mid, nums[mid], target, left)
		} else {
			right = mid
			fmt.Printf("       nums[%d]=%d >= %d, right=%d\n", mid, nums[mid], target, right)
		}
		step++
	}

	if left < len(nums) && nums[left] == target {
		fmt.Printf("找到左边界: %d\n", left)
	} else {
		fmt.Println("未找到目标值")
	}
}

func demonstrateRightBound(nums []int, target int) {
	left, right := 0, len(nums)
	step := 1

	fmt.Printf("初始: left=%d, right=%d\n", left, right)

	for left < right {
		mid := left + (right-left)/2
		fmt.Printf("步骤%d: left=%d, right=%d, mid=%d, nums[%d]=%d\n",
			step, left, right, mid, mid, nums[mid])

		if nums[mid] <= target {
			left = mid + 1
			fmt.Printf("       nums[%d]=%d <= %d, left=%d\n", mid, nums[mid], target, left)
		} else {
			right = mid
			fmt.Printf("       nums[%d]=%d > %d, right=%d\n", mid, nums[mid], target, right)
		}
		step++
	}

	rightBound := left - 1
	fmt.Printf("找到右边界: %d\n", rightBound)
}

func main() {
	fmt.Println("34. 在排序数组中查找元素的第一个和最后一个位置")
	fmt.Println("================================================")

	// 基础功能测试
	testSearchRange()

	// 性能对比测试
	benchmarkSearchRange()

	// 二分查找过程演示
	demonstrateBinarySearch()

	// 展示算法特点
	fmt.Println("\n=== 算法特点分析 ===")
	fmt.Println("1. 双重二分：经典解法，两次独立二分查找，清晰易懂")
	fmt.Println("2. 优化二分：统一模板，边界处理更加清晰")
	fmt.Println("3. 标准库法：利用内置函数，代码简洁")
	fmt.Println("4. 递归二分：递归实现，代码简洁但有栈溢出风险")
	fmt.Println("5. 线性查找：时间复杂度O(n)，不满足题目要求")

	fmt.Println("\n=== 关键技巧总结 ===")
	fmt.Println("• 左边界查找：找第一个>=target的位置")
	fmt.Println("• 右边界查找：找最后一个<=target的位置")
	fmt.Println("• 边界检查：确保找到的位置值等于target")
	fmt.Println("• 溢出防护：mid计算使用left+(right-left)/2")
	fmt.Println("• 模板统一：使用一致的二分查找边界处理")
}
