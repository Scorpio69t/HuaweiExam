package main

import "fmt"

// 方法一：滑动窗口（推荐）
func equalSubstring(s string, t string, maxCost int) int {
	n := len(s)
	left, right := 0, 0
	currentCost := 0
	maxLength := 0

	for right < n {
		// 计算当前字符的开销
		cost := abs(int(s[right]) - int(t[right]))

		// 扩展窗口
		currentCost += cost
		right++

		// 如果开销超过预算，收缩窗口
		for currentCost > maxCost {
			currentCost -= abs(int(s[left]) - int(t[left]))
			left++
		}

		// 更新最大长度
		maxLength = max(maxLength, right-left)
	}

	return maxLength
}

// 方法二：前缀和 + 二分查找
func equalSubstring2(s string, t string, maxCost int) int {
	n := len(s)

	// 计算前缀和
	prefixSum := make([]int, n+1)
	for i := 0; i < n; i++ {
		cost := abs(int(s[i]) - int(t[i]))
		prefixSum[i+1] = prefixSum[i] + cost
	}

	maxLength := 0
	// 对于每个起始位置，二分查找最远的结束位置
	for i := 0; i < n; i++ {
		// 二分查找最远的j，使得prefixSum[j+1] - prefixSum[i] <= maxCost
		left, right := i, n-1
		for left <= right {
			mid := left + (right-left)/2
			cost := prefixSum[mid+1] - prefixSum[i]
			if cost <= maxCost {
				maxLength = max(maxLength, mid-i+1)
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}

	return maxLength
}

// 方法三：暴力解法（用于验证）
func equalSubstring3(s string, t string, maxCost int) int {
	n := len(s)
	maxLength := 0

	for i := 0; i < n; i++ {
		currentCost := 0
		for j := i; j < n; j++ {
			cost := abs(int(s[j]) - int(t[j]))
			currentCost += cost
			if currentCost <= maxCost {
				maxLength = max(maxLength, j-i+1)
			} else {
				break
			}
		}
	}

	return maxLength
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 测试函数
func testEqualSubstring() {
	testCases := []struct {
		s, t     string
		maxCost  int
		expected int
	}{
		{"abcd", "bcdf", 3, 3},
		{"abcd", "cdef", 3, 1},
		{"abcd", "acde", 0, 1},
		{"krrgw", "zjxss", 19, 2},
		{"abcd", "abcd", 0, 4},
		{"a", "b", 1, 1},
		{"a", "b", 0, 0},
		{"abcdef", "bcdefg", 5, 5},
		{"xyz", "abc", 10, 3}, // x->a: 23, y->b: 23, z->c: 23, 总开销=69 > 10，所以只能转换1个字符
		{"xyz", "abc", 2, 0},  // 任何字符转换开销都超过2，所以结果为0
	}

	fmt.Println("=== 测试滑动窗口算法 ===")
	for i, tc := range testCases {
		result := equalSubstring(tc.s, tc.t, tc.maxCost)
		if result == tc.expected {
			fmt.Printf("✅ 测试用例 %d 通过: s=%s, t=%s, maxCost=%d, 结果=%d\n",
				i+1, tc.s, tc.t, tc.maxCost, result)
		} else {
			fmt.Printf("❌ 测试用例 %d 失败: s=%s, t=%s, maxCost=%d, 期望=%d, 实际=%d\n",
				i+1, tc.s, tc.t, tc.maxCost, tc.expected, result)
		}
	}

	fmt.Println("\n=== 测试前缀和+二分查找算法 ===")
	for i, tc := range testCases {
		result := equalSubstring2(tc.s, tc.t, tc.maxCost)
		if result == tc.expected {
			fmt.Printf("✅ 测试用例 %d 通过: s=%s, t=%s, maxCost=%d, 结果=%d\n",
				i+1, tc.s, tc.t, tc.maxCost, result)
		} else {
			fmt.Printf("❌ 测试用例 %d 失败: s=%s, t=%s, maxCost=%d, 期望=%d, 实际=%d\n",
				i+1, tc.s, tc.t, tc.maxCost, tc.expected, result)
		}
	}

	fmt.Println("\n=== 测试暴力解法 ===")
	for i, tc := range testCases {
		result := equalSubstring3(tc.s, tc.t, tc.maxCost)
		if result == tc.expected {
			fmt.Printf("✅ 测试用例 %d 通过: s=%s, t=%s, maxCost=%d, 结果=%d\n",
				i+1, tc.s, tc.t, tc.maxCost, result)
		} else {
			fmt.Printf("❌ 测试用例 %d 失败: s=%s, t=%s, maxCost=%d, 期望=%d, 实际=%d\n",
				i+1, tc.s, tc.t, tc.maxCost, tc.expected, result)
		}
	}
}

// 性能对比测试
func benchmarkTest() {
	s := "abcdefghijklmnopqrstuvwxyz"
	t := "bcdefghijklmnopqrstuvwxyza"
	maxCost := 100

	fmt.Println("\n=== 性能对比测试 ===")
	fmt.Printf("输入: s=%s, t=%s, maxCost=%d\n", s, t, maxCost)

	// 测试滑动窗口
	result1 := equalSubstring(s, t, maxCost)
	fmt.Printf("滑动窗口结果: %d\n", result1)

	// 测试前缀和+二分查找
	result2 := equalSubstring2(s, t, maxCost)
	fmt.Printf("前缀和+二分查找结果: %d\n", result2)

	// 测试暴力解法
	result3 := equalSubstring3(s, t, maxCost)
	fmt.Printf("暴力解法结果: %d\n", result3)

	if result1 == result2 && result2 == result3 {
		fmt.Println("✅ 所有算法结果一致")
	} else {
		fmt.Println("❌ 算法结果不一致")
	}
}

// 调试函数：手动验证测试用例
func debugTestCase() {
	fmt.Println("\n=== 调试测试用例 ===")

	// 测试用例9: "xyz" -> "abc", maxCost=10
	s1, t1 := "xyz", "abc"
	maxCost1 := 10

	fmt.Printf("测试用例9: s=%s, t=%s, maxCost=%d\n", s1, t1, maxCost1)
	for i := 0; i < len(s1); i++ {
		cost := abs(int(s1[i]) - int(t1[i]))
		fmt.Printf("  %c -> %c: |%d - %d| = %d\n", s1[i], t1[i], int(s1[i]), int(t1[i]), cost)
	}

	// 计算所有可能的子字符串开销
	fmt.Println("  所有可能的子字符串开销:")
	for i := 0; i < len(s1); i++ {
		for j := i; j < len(s1); j++ {
			totalCost := 0
			for k := i; k <= j; k++ {
				totalCost += abs(int(s1[k]) - int(t1[k]))
			}
			fmt.Printf("    [%d,%d]: %s -> %s, 开销=%d\n", i, j, s1[i:j+1], t1[i:j+1], totalCost)
		}
	}

	result1 := equalSubstring(s1, t1, maxCost1)
	fmt.Printf("  结果: %d\n", result1)

	// 测试用例10: "xyz" -> "abc", maxCost=2
	maxCost2 := 2
	fmt.Printf("\n测试用例10: s=%s, t=%s, maxCost=%d\n", s1, t1, maxCost2)

	// 计算所有可能的子字符串开销
	fmt.Println("  所有可能的子字符串开销:")
	for i := 0; i < len(s1); i++ {
		for j := i; j < len(s1); j++ {
			totalCost := 0
			for k := i; k <= j; k++ {
				totalCost += abs(int(s1[k]) - int(t1[k]))
			}
			fmt.Printf("    [%d,%d]: %s -> %s, 开销=%d\n", i, j, s1[i:j+1], t1[i:j+1], totalCost)
		}
	}

	result2 := equalSubstring(s1, t1, maxCost2)
	fmt.Printf("  结果: %d\n", result2)
}

func main() {
	fmt.Println("1208. 尽可能使字符串相等")
	fmt.Println("========================")

	// 运行调试函数
	debugTestCase()

	// 运行测试用例
	testEqualSubstring()

	// 运行性能对比
	benchmarkTest()

	// 交互式测试
	fmt.Println("\n=== 交互式测试 ===")
	fmt.Println("请输入测试用例 (格式: s t maxCost，例如: abcd bcdf 3)")

	var s, t string
	var maxCost int

	fmt.Print("请输入 s: ")
	fmt.Scanln(&s)
	fmt.Print("请输入 t: ")
	fmt.Scanln(&t)
	fmt.Print("请输入 maxCost: ")
	fmt.Scanln(&maxCost)

	if len(s) != len(t) {
		fmt.Println("❌ 错误：字符串长度必须相同")
		return
	}

	result := equalSubstring(s, t, maxCost)
	fmt.Printf("结果: %d\n", result)
}
