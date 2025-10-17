package main

import (
	"fmt"
	"math"
)

// =========================== 方法一：滑动窗口（最优解法） ===========================

func minWindow(s string, t string) string {
	need := make(map[byte]int)
	window := make(map[byte]int)

	// 统计t中字符需求
	for i := 0; i < len(t); i++ {
		need[t[i]]++
	}

	left, right := 0, 0
	valid := 0 // 窗口中满足需求的字符种类数
	start, length := 0, math.MaxInt32

	for right < len(s) {
		// 扩展窗口
		c := s[right]
		right++

		if need[c] > 0 {
			window[c]++
			if window[c] == need[c] {
				valid++
			}
		}

		// 收缩窗口
		for valid == len(need) {
			// 更新最小窗口
			if right-left < length {
				start = left
				length = right - left
			}

			d := s[left]
			left++

			if need[d] > 0 {
				if window[d] == need[d] {
					valid--
				}
				window[d]--
			}
		}
	}

	if length == math.MaxInt32 {
		return ""
	}
	return s[start : start+length]
}

// =========================== 方法二：优化版滑动窗口 ===========================

func minWindow2(s string, t string) string {
	if len(s) < len(t) {
		return ""
	}

	need := make(map[byte]int)
	window := make(map[byte]int)
	for i := 0; i < len(t); i++ {
		need[t[i]]++
	}

	left, right := 0, 0
	valid := 0
	start, minLen := 0, len(s)+1

	for right < len(s) {
		// 扩展窗口
		c := s[right]
		right++

		if need[c] > 0 {
			window[c]++
			if window[c] == need[c] {
				valid++
			}
		}

		// 收缩窗口
		for valid == len(need) {
			if right-left < minLen {
				start = left
				minLen = right - left
			}

			d := s[left]
			left++

			if need[d] > 0 {
				if window[d] == need[d] {
					valid--
				}
				window[d]--
			}
		}
	}

	if minLen == len(s)+1 {
		return ""
	}
	return s[start : start+minLen]
}

// =========================== 方法三：双指针 + 数组（ASCII优化） ===========================

func minWindow3(s string, t string) string {
	if len(s) < len(t) {
		return ""
	}

	// 使用数组代替map，适用于ASCII字符
	need := make([]int, 128)
	window := make([]int, 128)

	for i := 0; i < len(t); i++ {
		need[t[i]]++
	}

	// 统计t中不同字符的数量
	needCount := 0
	for i := 0; i < 128; i++ {
		if need[i] > 0 {
			needCount++
		}
	}

	left, right := 0, 0
	valid := 0
	start, minLen := 0, len(s)+1

	for right < len(s) {
		c := s[right]
		right++

		if need[c] > 0 {
			window[c]++
			if window[c] == need[c] {
				valid++
			}
		}

		for valid == needCount {
			if right-left < minLen {
				start = left
				minLen = right - left
			}

			d := s[left]
			left++

			if need[d] > 0 {
				if window[d] == need[d] {
					valid--
				}
				window[d]--
			}
		}
	}

	if minLen == len(s)+1 {
		return ""
	}
	return s[start : start+minLen]
}

// =========================== 方法四：暴力枚举（对比用） ===========================

func minWindow4(s string, t string) string {
	minLen := len(s) + 1
	result := ""

	for i := 0; i < len(s); i++ {
		for j := i; j < len(s); j++ {
			if isValid(s[i:j+1], t) {
				if j-i+1 < minLen {
					minLen = j - i + 1
					result = s[i : j+1]
				}
			}
		}
	}

	return result
}

func isValid(sub, t string) bool {
	count := make(map[byte]int)
	for i := 0; i < len(sub); i++ {
		count[sub[i]]++
	}

	for i := 0; i < len(t); i++ {
		if count[t[i]] == 0 {
			return false
		}
		count[t[i]]--
	}

	return true
}

// =========================== 测试代码 ===========================

func main() {
	fmt.Println("=== LeetCode 76: 最小覆盖子串 ===\n")

	testCases := []struct {
		s      string
		t      string
		expect string
	}{
		{
			"ADOBECODEBANC",
			"ABC",
			"BANC",
		},
		{
			"a",
			"a",
			"a",
		},
		{
			"a",
			"aa",
			"",
		},
		{
			"aab",
			"aab",
			"aab",
		},
		{
			"ab",
			"b",
			"b",
		},
		{
			"bba",
			"ab",
			"ba",
		},
	}

	fmt.Println("方法一：滑动窗口（最优解法）")
	runTests(testCases, minWindow)

	fmt.Println("\n方法二：优化版滑动窗口")
	runTests(testCases, minWindow2)

	fmt.Println("\n方法三：双指针 + 数组（ASCII优化）")
	runTests(testCases, minWindow3)

	fmt.Println("\n方法四：暴力枚举（对比用）")
	runTests(testCases, minWindow4)
}

func runTests(testCases []struct {
	s      string
	t      string
	expect string
}, fn func(string, string) string) {
	passCount := 0
	for i, tc := range testCases {
		result := fn(tc.s, tc.t)
		status := "✅"
		if result != tc.expect {
			status = "❌"
		} else {
			passCount++
		}
		fmt.Printf("  测试%d: %s\n", i+1, status)
		if status == "❌" {
			fmt.Printf("    输入: s=\"%s\", t=\"%s\"\n", tc.s, tc.t)
			fmt.Printf("    输出: \"%s\"\n", result)
			fmt.Printf("    期望: \"%s\"\n", tc.expect)
		}
	}
	fmt.Printf("  通过: %d/%d\n", passCount, len(testCases))
}
