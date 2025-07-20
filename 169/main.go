package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

// 方法一：摩尔投票法（推荐）
// 时间复杂度：O(n)，空间复杂度：O(1)
func majorityElement(nums []int) int {
	candidate := 0
	count := 0

	// 投票阶段
	for _, num := range nums {
		if count == 0 {
			candidate = num
		}
		if num == candidate {
			count++
		} else {
			count--
		}
	}

	return candidate
}

// 方法二：哈希表统计法
// 时间复杂度：O(n)，空间复杂度：O(n)
func majorityElementHashMap(nums []int) int {
	n := len(nums)
	countMap := make(map[int]int)

	// 统计每个元素出现次数
	for _, num := range nums {
		countMap[num]++
		// 一旦发现某个元素次数超过n/2，立即返回
		if countMap[num] > n/2 {
			return num
		}
	}

	// 理论上不会执行到这里（题目保证多数元素存在）
	return nums[0]
}

// 方法三：排序法
// 时间复杂度：O(n log n)，空间复杂度：O(1)
func majorityElementSort(nums []int) int {
	sort.Ints(nums)
	// 排序后，中位数必然是多数元素
	return nums[len(nums)/2]
}

// 方法四：分治法
// 时间复杂度：O(n log n)，空间复杂度：O(log n)
func majorityElementDivideConquer(nums []int) int {
	return majorityElementRec(nums, 0, len(nums)-1)
}

func majorityElementRec(nums []int, left, right int) int {
	// 基础情况
	if left == right {
		return nums[left]
	}

	// 分治递归
	mid := left + (right-left)/2
	leftMajority := majorityElementRec(nums, left, mid)
	rightMajority := majorityElementRec(nums, mid+1, right)

	// 如果左右部分的多数元素相同，直接返回
	if leftMajority == rightMajority {
		return leftMajority
	}

	// 统计两个候选元素在整个区间的出现次数
	leftCount := countInRange(nums, leftMajority, left, right)
	rightCount := countInRange(nums, rightMajority, left, right)

	if leftCount > rightCount {
		return leftMajority
	}
	return rightMajority
}

func countInRange(nums []int, target, left, right int) int {
	count := 0
	for i := left; i <= right; i++ {
		if nums[i] == target {
			count++
		}
	}
	return count
}

// 方法五：随机化算法
// 期望时间复杂度：O(n)，空间复杂度：O(1)
func majorityElementRandomized(nums []int) int {
	n := len(nums)
	rand.Seed(time.Now().UnixNano())

	for {
		// 随机选择一个元素
		candidate := nums[rand.Intn(n)]

		// 统计该元素出现次数
		count := 0
		for _, num := range nums {
			if num == candidate {
				count++
			}
		}

		// 如果是多数元素，返回
		if count > n/2 {
			return candidate
		}
	}
}

// 方法六：摩尔投票法的详细步骤版本（用于演示）
func majorityElementMooreVotingDetailed(nums []int) int {
	fmt.Println("摩尔投票过程演示：")
	candidate := 0
	count := 0

	for i, num := range nums {
		fmt.Printf("步骤%d: 当前元素=%d, ", i+1, num)

		if count == 0 {
			candidate = num
			count = 1
			fmt.Printf("计数器为0，候选者更新为%d，计数器=1\n", candidate)
		} else if num == candidate {
			count++
			fmt.Printf("相同元素，计数器+1，当前计数器=%d\n", count)
		} else {
			count--
			fmt.Printf("不同元素，计数器-1，当前计数器=%d\n", count)
		}
	}

	fmt.Printf("最终候选者: %d\n", candidate)
	return candidate
}

// 辅助函数：验证结果是否正确
func verifyResult(nums []int, result int) bool {
	count := 0
	for _, num := range nums {
		if num == result {
			count++
		}
	}
	return count > len(nums)/2
}

// 辅助函数：性能测试
func measureTime(name string, fn func() int) int {
	start := time.Now()
	result := fn()
	duration := time.Since(start)
	fmt.Printf("%s: 结果=%d, 耗时=%v\n", name, result, duration)
	return result
}

// 辅助函数：复制切片
func copySlice(nums []int) []int {
	copied := make([]int, len(nums))
	copy(copied, nums)
	return copied
}

// 运行所有测试用例
func runTests() {
	fmt.Println("=== 169. 多数元素 测试用例 ===")

	// 测试用例1：示例1
	fmt.Println("\n--- 测试用例1：示例1 ---")
	nums1 := []int{3, 2, 3}
	fmt.Printf("输入: %v\n", nums1)

	result1_1 := majorityElement(copySlice(nums1))
	result1_2 := majorityElementHashMap(copySlice(nums1))
	result1_3 := majorityElementSort(copySlice(nums1))
	result1_4 := majorityElementDivideConquer(copySlice(nums1))
	result1_5 := majorityElementRandomized(copySlice(nums1))

	fmt.Printf("摩尔投票法: %d\n", result1_1)
	fmt.Printf("哈希表统计法: %d\n", result1_2)
	fmt.Printf("排序法: %d\n", result1_3)
	fmt.Printf("分治法: %d\n", result1_4)
	fmt.Printf("随机化算法: %d\n", result1_5)

	// 验证结果
	expected1 := 3
	if result1_1 == expected1 && result1_2 == expected1 &&
		result1_3 == expected1 && result1_4 == expected1 && result1_5 == expected1 {
		fmt.Println("✅ 所有方法结果正确且一致")
	} else {
		fmt.Println("❌ 结果不一致！")
	}

	// 测试用例2：示例2
	fmt.Println("\n--- 测试用例2：示例2 ---")
	nums2 := []int{2, 2, 1, 1, 1, 2, 2}
	fmt.Printf("输入: %v\n", nums2)

	// 详细演示摩尔投票过程
	result2 := majorityElementMooreVotingDetailed(copySlice(nums2))

	// 测试其他方法
	result2_2 := majorityElementHashMap(copySlice(nums2))
	result2_3 := majorityElementSort(copySlice(nums2))
	result2_4 := majorityElementDivideConquer(copySlice(nums2))

	fmt.Printf("哈希表统计法: %d\n", result2_2)
	fmt.Printf("排序法: %d\n", result2_3)
	fmt.Printf("分治法: %d\n", result2_4)

	// 验证结果
	expected2 := 2
	if result2 == expected2 && result2_2 == expected2 &&
		result2_3 == expected2 && result2_4 == expected2 {
		fmt.Println("✅ 所有方法结果正确且一致")
	} else {
		fmt.Println("❌ 结果不一致！")
	}

	// 测试用例3：单元素数组
	fmt.Println("\n--- 测试用例3：单元素数组 ---")
	nums3 := []int{1}
	fmt.Printf("输入: %v\n", nums3)
	result3 := majorityElement(nums3)
	fmt.Printf("结果: %d (预期: 1)\n", result3)

	// 测试用例4：所有元素相同
	fmt.Println("\n--- 测试用例4：所有元素相同 ---")
	nums4 := []int{5, 5, 5, 5, 5}
	fmt.Printf("输入: %v\n", nums4)
	result4 := majorityElement(nums4)
	fmt.Printf("结果: %d (预期: 5)\n", result4)

	// 测试用例5：包含负数
	fmt.Println("\n--- 测试用例5：包含负数 ---")
	nums5 := []int{-1, -1, -1, 2, 2}
	fmt.Printf("输入: %v\n", nums5)
	result5 := majorityElement(nums5)
	fmt.Printf("结果: %d (预期: -1)\n", result5)

	// 测试用例6：刚好过半
	fmt.Println("\n--- 测试用例6：刚好过半 ---")
	nums6 := []int{1, 1, 1, 2, 2}
	fmt.Printf("输入: %v\n", nums6)
	result6 := majorityElement(nums6)
	fmt.Printf("结果: %d (预期: 1)\n", result6)

	// 性能测试
	fmt.Println("\n--- 性能对比测试 ---")
	performanceTest()
}

// 性能对比测试
func performanceTest() {
	// 创建大规模测试数据
	n := 100000
	testNums := make([]int, n)

	// 生成多数元素为1的数组
	majorityCount := n/2 + 1
	for i := 0; i < majorityCount; i++ {
		testNums[i] = 1
	}
	for i := majorityCount; i < n; i++ {
		testNums[i] = i%100 + 2 // 避免与多数元素1重复
	}

	// 打乱数组
	rand.Seed(time.Now().UnixNano())
	for i := len(testNums) - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		testNums[i], testNums[j] = testNums[j], testNums[i]
	}

	fmt.Printf("测试数据：%d个元素，多数元素为1\n", n)

	// 测试各种方法的性能
	result1 := measureTime("摩尔投票法", func() int {
		return majorityElement(copySlice(testNums))
	})

	result2 := measureTime("哈希表统计法", func() int {
		return majorityElementHashMap(copySlice(testNums))
	})

	result3 := measureTime("排序法", func() int {
		return majorityElementSort(copySlice(testNums))
	})

	result4 := measureTime("分治法", func() int {
		return majorityElementDivideConquer(copySlice(testNums))
	})

	// 随机化算法可能比较慢，只在小数据上测试
	smallNums := testNums[:1000]
	result5 := measureTime("随机化算法(1000元素)", func() int {
		return majorityElementRandomized(copySlice(smallNums))
	})

	// 验证结果一致性
	if result1 == 1 && result2 == 1 && result3 == 1 && result4 == 1 && result5 == 1 {
		fmt.Println("✅ 所有方法结果一致")
	} else {
		fmt.Printf("❌ 结果不一致！摩尔:%d, 哈希:%d, 排序:%d, 分治:%d, 随机:%d\n",
			result1, result2, result3, result4, result5)
	}
}

// 演示分治过程
func demonstrateDivideConquer() {
	fmt.Println("\n--- 分治过程演示 ---")

	nums := []int{2, 2, 1, 1, 1, 2, 2}
	fmt.Printf("原始数组: %v\n", nums)

	fmt.Println("\n分治递归过程：")
	result := demonstrateDivideConquerRec(nums, 0, len(nums)-1, 0)
	fmt.Printf("最终结果: %d\n", result)
}

func demonstrateDivideConquerRec(nums []int, left, right, depth int) int {
	indent := ""
	for i := 0; i < depth; i++ {
		indent += "  "
	}

	fmt.Printf("%s处理区间[%d,%d]: %v\n", indent, left, right, nums[left:right+1])

	if left == right {
		fmt.Printf("%s基础情况，返回: %d\n", indent, nums[left])
		return nums[left]
	}

	mid := left + (right-left)/2
	fmt.Printf("%s分割为[%d,%d]和[%d,%d]\n", indent, left, mid, mid+1, right)

	leftMajority := demonstrateDivideConquerRec(nums, left, mid, depth+1)
	rightMajority := demonstrateDivideConquerRec(nums, mid+1, right, depth+1)

	fmt.Printf("%s左部分多数元素: %d, 右部分多数元素: %d\n", indent, leftMajority, rightMajority)

	if leftMajority == rightMajority {
		fmt.Printf("%s左右相同，返回: %d\n", indent, leftMajority)
		return leftMajority
	}

	leftCount := countInRange(nums, leftMajority, left, right)
	rightCount := countInRange(nums, rightMajority, left, right)

	fmt.Printf("%s统计次数 - %d出现%d次, %d出现%d次\n",
		indent, leftMajority, leftCount, rightMajority, rightCount)

	if leftCount > rightCount {
		fmt.Printf("%s返回: %d\n", indent, leftMajority)
		return leftMajority
	}
	fmt.Printf("%s返回: %d\n", indent, rightMajority)
	return rightMajority
}

// 算法正确性验证
func verifyAlgorithmCorrectness() {
	fmt.Println("\n--- 算法正确性验证 ---")

	testCases := []struct {
		name     string
		nums     []int
		expected int
	}{
		{"基本情况1", []int{3, 2, 3}, 3},
		{"基本情况2", []int{2, 2, 1, 1, 1, 2, 2}, 2},
		{"单元素", []int{1}, 1},
		{"全相同", []int{1, 1, 1, 1}, 1},
		{"负数", []int{-1, -1, 0}, -1},
		{"大数", []int{1000000000, 1000000000, 1000000001}, 1000000000},
		{"刚好过半", []int{1, 1, 1, 2, 2}, 1},
	}

	allPassed := true
	for _, tc := range testCases {
		result := majorityElement(tc.nums)
		if result == tc.expected && verifyResult(tc.nums, result) {
			fmt.Printf("✅ %s: 通过 (结果=%d)\n", tc.name, result)
		} else {
			fmt.Printf("❌ %s: 失败 (结果=%d, 预期=%d)\n", tc.name, result, tc.expected)
			allPassed = false
		}
	}

	if allPassed {
		fmt.Println("✅ 所有测试用例通过")
	} else {
		fmt.Println("❌ 部分测试用例失败")
	}
}

func main() {
	runTests()
	demonstrateDivideConquer()
	verifyAlgorithmCorrectness()
}
