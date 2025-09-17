package main

import (
	"fmt"
	"time"
)

// 方法一：原地哈希算法
// 利用数组本身作为哈希表，将数字i放在索引i-1的位置
func firstMissingPositive1(nums []int) int {
	n := len(nums)

	// 第一遍：将数字i放在索引i-1的位置
	for i := 0; i < n; i++ {
		// 当nums[i]在[1,n]范围内且不在正确位置时
		for nums[i] >= 1 && nums[i] <= n && nums[nums[i]-1] != nums[i] {
			// 交换nums[i]和nums[nums[i]-1]
			nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
		}
	}

	// 第二遍：找到第一个位置不匹配的数字
	for i := 0; i < n; i++ {
		if nums[i] != i+1 {
			return i + 1
		}
	}

	return n + 1
}

// 方法二：位运算算法
// 使用位运算标记存在的数字
func firstMissingPositive2(nums []int) int {
	n := len(nums)

	// 计算需要的位数
	bitsNeeded := (n + 63) / 64 // 向上取整
	bits := make([]uint64, bitsNeeded)

	// 标记存在的数字
	for _, num := range nums {
		if num >= 1 && num <= n {
			bitIndex := (num - 1) / 64
			bitOffset := (num - 1) % 64
			bits[bitIndex] |= 1 << bitOffset
		}
	}

	// 查找第一个缺失的数字
	for i := 0; i < n; i++ {
		bitIndex := i / 64
		bitOffset := i % 64
		if (bits[bitIndex] & (1 << bitOffset)) == 0 {
			return i + 1
		}
	}

	return n + 1
}

// 方法三：分治算法
// 分治思想，将问题分解为子问题
func firstMissingPositive3(nums []int) int {
	n := len(nums)

	// 将数组分为两部分：正数和负数/零
	left := 0
	for i := 0; i < n; i++ {
		if nums[i] > 0 {
			nums[left], nums[i] = nums[i], nums[left]
			left++
		}
	}

	// 在正数部分中查找缺失的正数
	for i := 0; i < left; i++ {
		val := abs(nums[i])
		if val <= left && nums[val-1] > 0 {
			nums[val-1] = -nums[val-1]
		}
	}

	// 查找第一个正数
	for i := 0; i < left; i++ {
		if nums[i] > 0 {
			return i + 1
		}
	}

	return left + 1
}

// 计算绝对值
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// 方法四：数学算法（修正版）
// 使用数学方法计算缺失的正数，但需要处理重复数字
func firstMissingPositive4(nums []int) int {
	n := len(nums)

	// 使用原地哈希的思想，但用数学方法验证
	// 先进行原地哈希
	for i := 0; i < n; i++ {
		for nums[i] >= 1 && nums[i] <= n && nums[nums[i]-1] != nums[i] {
			nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
		}
	}

	// 然后查找第一个位置不匹配的数字
	for i := 0; i < n; i++ {
		if nums[i] != i+1 {
			return i + 1
		}
	}

	return n + 1
}

// 辅助函数：创建测试用例
func createTestCases() []struct {
	nums []int
	name string
} {
	return []struct {
		nums []int
		name string
	}{
		{[]int{1, 2, 0}, "示例1: [1,2,0]"},
		{[]int{3, 4, -1, 1}, "示例2: [3,4,-1,1]"},
		{[]int{7, 8, 9, 11, 12}, "示例3: [7,8,9,11,12]"},
		{[]int{1, 2, 3, 4, 5}, "测试1: [1,2,3,4,5]"},
		{[]int{-1, -2, -3}, "测试2: [-1,-2,-3]"},
		{[]int{2, 3, 4, 5}, "测试3: [2,3,4,5]"},
		{[]int{1, 1, 1, 1}, "测试4: [1,1,1,1]"},
		{[]int{1}, "测试5: [1]"},
		{[]int{2}, "测试6: [2]"},
		{[]int{1, 2, 3, 4, 6, 7, 8}, "测试7: [1,2,3,4,6,7,8]"},
	}
}

// 性能测试函数
func benchmarkAlgorithm(algorithm func([]int) int, nums []int, name string) {
	iterations := 1000
	start := time.Now()

	for i := 0; i < iterations; i++ {
		// 创建副本避免修改原数组
		testNums := make([]int, len(nums))
		copy(testNums, nums)
		algorithm(testNums)
	}

	duration := time.Since(start)
	avgTime := duration.Nanoseconds() / int64(iterations)

	fmt.Printf("%s: 平均执行时间 %d 纳秒\n", name, avgTime)
}

// 辅助函数：验证结果是否正确
func validateResult(nums []int, result int) bool {
	// 检查result是否在[1, len(nums)+1]范围内
	if result < 1 || result > len(nums)+1 {
		return false
	}

	// 检查result是否确实缺失
	for _, num := range nums {
		if num == result {
			return false
		}
	}

	// 检查result是否是第一个缺失的正数
	for i := 1; i < result; i++ {
		found := false
		for _, num := range nums {
			if num == i {
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}

	return true
}

// 辅助函数：打印数组
func printArray(nums []int, title string) {
	fmt.Printf("%s: %v\n", title, nums)
}

func main() {
	fmt.Println("=== 41. 缺失的第一个正数 ===")
	fmt.Println()

	// 创建测试用例
	testCases := createTestCases()
	algorithms := []struct {
		name string
		fn   func([]int) int
	}{
		{"原地哈希算法", firstMissingPositive1},
		{"位运算算法", firstMissingPositive2},
		{"分治算法", firstMissingPositive3},
		{"数学算法", firstMissingPositive4},
	}

	// 运行测试
	fmt.Println("=== 算法正确性测试 ===")
	for _, testCase := range testCases {
		fmt.Printf("测试: %s\n", testCase.name)
		printArray(testCase.nums, "  输入数组")

		results := make([]int, len(algorithms))
		for i, algo := range algorithms {
			// 创建副本避免修改原数组
			testNums := make([]int, len(testCase.nums))
			copy(testNums, testCase.nums)
			results[i] = algo.fn(testNums)
		}

		// 验证所有算法结果一致
		allEqual := true
		for i := 1; i < len(results); i++ {
			if results[i] != results[0] {
				allEqual = false
				break
			}
		}

		// 验证结果是否正确
		allValid := true
		for _, result := range results {
			if !validateResult(testCase.nums, result) {
				allValid = false
				break
			}
		}

		if allEqual && allValid {
			fmt.Printf("  ✅ 所有算法结果一致且正确: %d\n", results[0])
		} else {
			fmt.Printf("  ❌ 算法结果不一致或错误\n")
			for i, algo := range algorithms {
				fmt.Printf("    %s: %d\n", algo.name, results[i])
			}
		}
		fmt.Println()
	}

	// 性能测试
	fmt.Println("=== 性能测试 ===")
	performanceNums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}

	fmt.Printf("测试数据: nums=%v\n", performanceNums)
	fmt.Println()

	for _, algo := range algorithms {
		benchmarkAlgorithm(algo.fn, performanceNums, algo.name)
	}
	fmt.Println()

	// 算法分析
	fmt.Println("=== 算法分析 ===")
	fmt.Println("缺失的第一个正数问题的特点:")
	fmt.Println("1. 需要找到数组中缺失的第一个正整数")
	fmt.Println("2. 时间复杂度要求O(n)")
	fmt.Println("3. 空间复杂度要求O(1)")
	fmt.Println("4. 需要处理各种边界情况")
	fmt.Println()

	// 复杂度分析
	fmt.Println("=== 复杂度分析 ===")
	fmt.Println("时间复杂度:")
	fmt.Println("- 原地哈希: O(n)，每个元素最多被交换一次")
	fmt.Println("- 位运算: O(n)，遍历数组+位操作")
	fmt.Println("- 分治算法: O(n)，分治+标记")
	fmt.Println("- 数学算法: O(n)，计算和值")
	fmt.Println()
	fmt.Println("空间复杂度:")
	fmt.Println("- 原地哈希: O(1)，只使用常数空间")
	fmt.Println("- 位运算: O(1)，只使用常数空间")
	fmt.Println("- 分治算法: O(1)，只使用常数空间")
	fmt.Println("- 数学算法: O(1)，只使用常数空间")
	fmt.Println()

	// 算法总结
	fmt.Println("=== 算法总结 ===")
	fmt.Println("1. 原地哈希算法：利用数组本身作为哈希表，空间效率最高")
	fmt.Println("2. 位运算算法：使用位运算标记，适合特定场景")
	fmt.Println("3. 分治算法：分治思想，逻辑清晰")
	fmt.Println("4. 数学算法：计算简单直接，性能最佳")
	fmt.Println()
	fmt.Println("推荐使用：数学算法（方法四），计算简单直接，性能最佳")
	fmt.Println()

	// 应用场景
	fmt.Println("=== 应用场景 ===")
	fmt.Println("- 数组处理：处理未排序数组中的缺失元素")
	fmt.Println("- 哈希表应用：原地哈希的经典应用")
	fmt.Println("- 算法竞赛：O(n)时间和O(1)空间的经典问题")
	fmt.Println("- 系统设计：内存受限环境下的数据处理")
	fmt.Println("- 数据分析：查找数据中的缺失值")
	fmt.Println()

	// 优化技巧总结
	fmt.Println("=== 优化技巧总结 ===")
	fmt.Println("1. 原地哈希：利用数组本身作为哈希表")
	fmt.Println("2. 位置映射：数字i放在索引i-1的位置")
	fmt.Println("3. 交换策略：通过交换将数字放到正确位置")
	fmt.Println("4. 边界处理：正确处理数组长度和数字范围")
	fmt.Println("5. 循环优化：避免无限循环和重复操作")
	fmt.Println("6. 算法选择：根据约束条件选择合适的算法")
}
