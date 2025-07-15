package main

import (
	"fmt"
	"time"
)

// 方法1：暴力法 - 时间复杂度O(n³)
func lengthOfLongestSubstringBruteForce(s string) int {
	n := len(s)
	if n == 0 {
		return 0
	}

	maxLength := 1

	// 枚举所有子串
	for i := 0; i < n; i++ {
		for j := i + 1; j <= n; j++ {
			if isUnique(s[i:j]) {
				if j-i > maxLength {
					maxLength = j - i
				}
			}
		}
	}

	return maxLength
}

// 检查字符串是否包含重复字符
func isUnique(s string) bool {
	charSet := make(map[byte]bool)
	for i := 0; i < len(s); i++ {
		if charSet[s[i]] {
			return false
		}
		charSet[s[i]] = true
	}
	return true
}

// 方法2：滑动窗口法 - 时间复杂度O(n)
func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}

	charMap := make(map[byte]int)
	left := 0
	maxLength := 0

	for right := 0; right < len(s); right++ {
		// 如果当前字符在窗口中已存在
		if pos, exists := charMap[s[right]]; exists && pos >= left {
			left = pos + 1
		}

		charMap[s[right]] = right

		// 更新最大长度
		if right-left+1 > maxLength {
			maxLength = right - left + 1
		}
	}

	return maxLength
}

// 方法3：优化的滑动窗口法（使用数组代替哈希表）
func lengthOfLongestSubstringOptimized(s string) int {
	if len(s) == 0 {
		return 0
	}

	// 使用数组存储字符最后出现的位置
	lastIndex := make([]int, 128) // ASCII字符集
	for i := range lastIndex {
		lastIndex[i] = -1
	}

	left := 0
	maxLength := 0

	for right := 0; right < len(s); right++ {
		char := s[right]

		// 如果当前字符在窗口中已存在
		if lastIndex[char] >= left {
			left = lastIndex[char] + 1
		}

		lastIndex[char] = right

		// 更新最大长度
		if right-left+1 > maxLength {
			maxLength = right - left + 1
		}
	}

	return maxLength
}

// 方法4：滑动窗口法（带详细过程展示）
func lengthOfLongestSubstringWithTrace(s string) int {
	if len(s) == 0 {
		return 0
	}

	fmt.Printf("输入字符串: %s\n", s)
	fmt.Println("滑动窗口过程:")
	fmt.Println("步骤\t左指针\t右指针\t当前窗口\t\t长度\t最大长度")

	charMap := make(map[byte]int)
	left := 0
	maxLength := 0
	step := 0

	for right := 0; right < len(s); right++ {
		step++

		// 如果当前字符在窗口中已存在
		if pos, exists := charMap[s[right]]; exists && pos >= left {
			left = pos + 1
		}

		charMap[s[right]] = right

		// 更新最大长度
		currentLength := right - left + 1
		if currentLength > maxLength {
			maxLength = currentLength
		}

		// 打印当前状态
		window := s[left : right+1]
		fmt.Printf("%d\t%d\t%d\t%s\t\t%d\t%d\n",
			step, left, right, window, currentLength, maxLength)
	}

	return maxLength
}

// 性能测试函数
func performanceTest() {
	testCases := []string{
		"abcabcbb",
		"bbbbb",
		"pwwkew",
		"",
		"abcdefghijklmnopqrstuvwxyz",
		"abcabcabcabcabc",
		"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaa",
	}

	fmt.Println("\n=== 性能测试 ===")
	fmt.Println("测试用例\t\t\t暴力法耗时\t滑动窗口耗时\t优化版耗时")

	for _, testCase := range testCases {
		// 暴力法测试
		start := time.Now()
		result1 := lengthOfLongestSubstringBruteForce(testCase)
		time1 := time.Since(start)

		// 滑动窗口法测试
		start = time.Now()
		result2 := lengthOfLongestSubstring(testCase)
		time2 := time.Since(start)

		// 优化版测试
		start = time.Now()
		result3 := lengthOfLongestSubstringOptimized(testCase)
		time3 := time.Since(start)

		// 验证结果一致性
		if result1 != result2 || result2 != result3 {
			fmt.Printf("错误：结果不一致！%d, %d, %d\n", result1, result2, result3)
		}

		displayCase := testCase
		if len(displayCase) > 15 {
			displayCase = displayCase[:15] + "..."
		}

		fmt.Printf("%-20s\t%v\t\t%v\t\t%v\n",
			displayCase, time1, time2, time3)
	}
}

func main() {
	// 基本测试用例
	testCases := []string{
		"abcabcbb",
		"bbbbb",
		"pwwkew",
		"",
		"au",
		"dvdf",
	}

	fmt.Println("=== 基本测试 ===")
	for _, testCase := range testCases {
		result := lengthOfLongestSubstring(testCase)
		fmt.Printf("输入: \"%s\" -> 输出: %d\n", testCase, result)
	}

	fmt.Println("\n=== 详细过程展示 ===")
	lengthOfLongestSubstringWithTrace("abcabcbb")

	// 性能测试
	performanceTest()

	fmt.Println("\n=== 算法复杂度分析 ===")
	fmt.Println("暴力法：")
	fmt.Println("  时间复杂度：O(n³)")
	fmt.Println("  空间复杂度：O(min(m,n)) - m为字符集大小")
	fmt.Println("\n滑动窗口法：")
	fmt.Println("  时间复杂度：O(n)")
	fmt.Println("  空间复杂度：O(min(m,n)) - m为字符集大小")
	fmt.Println("\n优化版滑动窗口：")
	fmt.Println("  时间复杂度：O(n)")
	fmt.Println("  空间复杂度：O(m) - 固定大小数组")
}
