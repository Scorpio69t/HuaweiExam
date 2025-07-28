package main

import "fmt"

// 方法一：中心扩展法（推荐）
// 时间复杂度：O(n²)，空间复杂度：O(1)
func longestPalindrome1(s string) string {
	if len(s) < 2 {
		return s
	}

	start, maxLen := 0, 1

	for i := 0; i < len(s); i++ {
		// 奇数长度回文
		len1 := expandAroundCenter(s, i, i)
		// 偶数长度回文
		len2 := expandAroundCenter(s, i, i+1)

		maxLenCur := max(len1, len2)
		if maxLenCur > maxLen {
			start = i - (maxLenCur-1)/2
			maxLen = maxLenCur
		}
	}

	return s[start : start+maxLen]
}

func expandAroundCenter(s string, left, right int) int {
	for left >= 0 && right < len(s) && s[left] == s[right] {
		left--
		right++
	}
	return right - left - 1
}

// 方法二：动态规划
// 时间复杂度：O(n²)，空间复杂度：O(n²)
func longestPalindrome2(s string) string {
	if len(s) < 2 {
		return s
	}

	n := len(s)
	dp := make([][]bool, n)
	for i := range dp {
		dp[i] = make([]bool, n)
	}

	start, maxLen := 0, 1

	// 初始化长度为1的回文
	for i := 0; i < n; i++ {
		dp[i][i] = true
	}

	// 初始化长度为2的回文
	for i := 0; i < n-1; i++ {
		if s[i] == s[i+1] {
			dp[i][i+1] = true
			start = i
			maxLen = 2
		}
	}

	// 按长度递增填充dp数组
	for length := 3; length <= n; length++ {
		for i := 0; i <= n-length; i++ {
			j := i + length - 1
			if s[i] == s[j] && dp[i+1][j-1] {
				dp[i][j] = true
				if length > maxLen {
					start = i
					maxLen = length
				}
			}
		}
	}

	return s[start : start+maxLen]
}

// 方法三：Manacher算法
// 时间复杂度：O(n)，空间复杂度：O(n)
func longestPalindrome3(s string) string {
	if len(s) < 2 {
		return s
	}

	// 预处理字符串，插入分隔符
	t := "#"
	for _, char := range s {
		t += string(char) + "#"
	}

	n := len(t)
	p := make([]int, n)
	center, right := 0, 0

	// 计算回文半径
	for i := 0; i < n; i++ {
		if i < right {
			mirror := 2*center - i
			p[i] = min(right-i, p[mirror])
		}

		// 尝试扩展
		left := i - (p[i] + 1)
		right_bound := i + (p[i] + 1)
		for left >= 0 && right_bound < n && t[left] == t[right_bound] {
			p[i]++
			left--
			right_bound++
		}

		// 更新中心点和右边界
		if i+p[i] > right {
			center = i
			right = i + p[i]
		}
	}

	// 找到最大回文半径
	maxRadius := 0
	maxCenter := 0
	for i := 0; i < n; i++ {
		if p[i] > maxRadius {
			maxRadius = p[i]
			maxCenter = i
		}
	}

	// 转换回原字符串的索引
	start := (maxCenter - maxRadius) / 2
	end := (maxCenter + maxRadius) / 2

	return s[start:end]
}

// 方法四：暴力解法
// 时间复杂度：O(n³)，空间复杂度：O(1)
func longestPalindrome4(s string) string {
	if len(s) < 2 {
		return s
	}

	start, maxLen := 0, 1

	for i := 0; i < len(s); i++ {
		for j := i + 1; j < len(s); j++ {
			if isPalindrome(s, i, j) && j-i+1 > maxLen {
				start = i
				maxLen = j - i + 1
			}
		}
	}

	return s[start : start+maxLen]
}

func isPalindrome(s string, left, right int) bool {
	for left < right {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}

// 方法五：优化的中心扩展法
// 时间复杂度：O(n²)，空间复杂度：O(1)
func longestPalindrome5(s string) string {
	if len(s) < 2 {
		return s
	}

	start, maxLen := 0, 1

	for i := 0; i < len(s); i++ {
		// 奇数长度回文
		len1 := expandAroundCenterOptimized(s, i, i)
		// 偶数长度回文
		len2 := expandAroundCenterOptimized(s, i, i+1)

		maxLenCur := max(len1, len2)
		if maxLenCur > maxLen {
			start = i - (maxLenCur-1)/2
			maxLen = maxLenCur
		}
	}

	return s[start : start+maxLen]
}

func expandAroundCenterOptimized(s string, left, right int) int {
	for left >= 0 && right < len(s) && s[left] == s[right] {
		left--
		right++
	}
	return right - left - 1
}

// 方法六：使用字符串哈希（Rabin-Karp思想）
// 时间复杂度：O(n²)，空间复杂度：O(1)
func longestPalindrome6(s string) string {
	if len(s) < 2 {
		return s
	}

	start, maxLen := 0, 1

	for i := 0; i < len(s); i++ {
		// 奇数长度回文
		len1 := expandAroundCenter(s, i, i)
		// 偶数长度回文
		len2 := expandAroundCenter(s, i, i+1)

		maxLenCur := max(len1, len2)
		if maxLenCur > maxLen {
			start = i - (maxLenCur-1)/2
			maxLen = maxLenCur
		}
	}

	return s[start : start+maxLen]
}

// 辅助函数
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	fmt.Println("=== 5. 最长回文子串 ===")

	// 测试用例1
	s1 := "babad"
	fmt.Printf("测试用例1: s=%s\n", s1)
	fmt.Printf("方法一结果: %s\n", longestPalindrome1(s1))
	fmt.Printf("方法二结果: %s\n", longestPalindrome2(s1))
	fmt.Printf("方法三结果: %s\n", longestPalindrome3(s1))
	fmt.Printf("方法四结果: %s\n", longestPalindrome4(s1))
	fmt.Printf("方法五结果: %s\n", longestPalindrome5(s1))
	fmt.Printf("方法六结果: %s\n", longestPalindrome6(s1))
	fmt.Println()

	// 测试用例2
	s2 := "cbbd"
	fmt.Printf("测试用例2: s=%s\n", s2)
	fmt.Printf("方法一结果: %s\n", longestPalindrome1(s2))
	fmt.Printf("方法二结果: %s\n", longestPalindrome2(s2))
	fmt.Printf("方法三结果: %s\n", longestPalindrome3(s2))
	fmt.Printf("方法四结果: %s\n", longestPalindrome4(s2))
	fmt.Printf("方法五结果: %s\n", longestPalindrome5(s2))
	fmt.Printf("方法六结果: %s\n", longestPalindrome6(s2))
	fmt.Println()

	// 测试用例3
	s3 := "a"
	fmt.Printf("测试用例3: s=%s\n", s3)
	fmt.Printf("方法一结果: %s\n", longestPalindrome1(s3))
	fmt.Printf("方法二结果: %s\n", longestPalindrome2(s3))
	fmt.Printf("方法三结果: %s\n", longestPalindrome3(s3))
	fmt.Printf("方法四结果: %s\n", longestPalindrome4(s3))
	fmt.Printf("方法五结果: %s\n", longestPalindrome5(s3))
	fmt.Printf("方法六结果: %s\n", longestPalindrome6(s3))
	fmt.Println()

	// 测试用例4
	s4 := "ac"
	fmt.Printf("测试用例4: s=%s\n", s4)
	fmt.Printf("方法一结果: %s\n", longestPalindrome1(s4))
	fmt.Printf("方法二结果: %s\n", longestPalindrome2(s4))
	fmt.Printf("方法三结果: %s\n", longestPalindrome3(s4))
	fmt.Printf("方法四结果: %s\n", longestPalindrome4(s4))
	fmt.Printf("方法五结果: %s\n", longestPalindrome5(s4))
	fmt.Printf("方法六结果: %s\n", longestPalindrome6(s4))
	fmt.Println()

	// 额外测试用例
	s5 := "racecar"
	fmt.Printf("额外测试: s=%s\n", s5)
	fmt.Printf("方法一结果: %s\n", longestPalindrome1(s5))
	fmt.Printf("方法二结果: %s\n", longestPalindrome2(s5))
	fmt.Printf("方法三结果: %s\n", longestPalindrome3(s5))
	fmt.Printf("方法四结果: %s\n", longestPalindrome4(s5))
	fmt.Printf("方法五结果: %s\n", longestPalindrome5(s5))
	fmt.Printf("方法六结果: %s\n", longestPalindrome6(s5))
	fmt.Println()

	// 边界测试用例
	s6 := ""
	fmt.Printf("边界测试: s=%s\n", s6)
	fmt.Printf("方法一结果: %s\n", longestPalindrome1(s6))
	fmt.Printf("方法二结果: %s\n", longestPalindrome2(s6))
	fmt.Printf("方法三结果: %s\n", longestPalindrome3(s6))
	fmt.Printf("方法四结果: %s\n", longestPalindrome4(s6))
	fmt.Printf("方法五结果: %s\n", longestPalindrome5(s6))
	fmt.Printf("方法六结果: %s\n", longestPalindrome6(s6))
	fmt.Println()

	// 复杂测试用例
	s7 := "abbaabba"
	fmt.Printf("复杂测试: s=%s\n", s7)
	fmt.Printf("方法一结果: %s\n", longestPalindrome1(s7))
	fmt.Printf("方法二结果: %s\n", longestPalindrome2(s7))
	fmt.Printf("方法三结果: %s\n", longestPalindrome3(s7))
	fmt.Printf("方法四结果: %s\n", longestPalindrome4(s7))
	fmt.Printf("方法五结果: %s\n", longestPalindrome5(s7))
	fmt.Printf("方法六结果: %s\n", longestPalindrome6(s7))
	fmt.Println()

	// 全相同字符测试用例
	s8 := "aaaa"
	fmt.Printf("全相同字符测试: s=%s\n", s8)
	fmt.Printf("方法一结果: %s\n", longestPalindrome1(s8))
	fmt.Printf("方法二结果: %s\n", longestPalindrome2(s8))
	fmt.Printf("方法三结果: %s\n", longestPalindrome3(s8))
	fmt.Printf("方法四结果: %s\n", longestPalindrome4(s8))
	fmt.Printf("方法五结果: %s\n", longestPalindrome5(s8))
	fmt.Printf("方法六结果: %s\n", longestPalindrome6(s8))
	fmt.Println()

	// 数字字符串测试用例
	s9 := "12321"
	fmt.Printf("数字字符串测试: s=%s\n", s9)
	fmt.Printf("方法一结果: %s\n", longestPalindrome1(s9))
	fmt.Printf("方法二结果: %s\n", longestPalindrome2(s9))
	fmt.Printf("方法三结果: %s\n", longestPalindrome3(s9))
	fmt.Printf("方法四结果: %s\n", longestPalindrome4(s9))
	fmt.Printf("方法五结果: %s\n", longestPalindrome5(s9))
	fmt.Printf("方法六结果: %s\n", longestPalindrome6(s9))
	fmt.Println()

	// 混合字符串测试用例
	s10 := "a1b2c3c2b1a"
	fmt.Printf("混合字符串测试: s=%s\n", s10)
	fmt.Printf("方法一结果: %s\n", longestPalindrome1(s10))
	fmt.Printf("方法二结果: %s\n", longestPalindrome2(s10))
	fmt.Printf("方法三结果: %s\n", longestPalindrome3(s10))
	fmt.Printf("方法四结果: %s\n", longestPalindrome4(s10))
	fmt.Printf("方法五结果: %s\n", longestPalindrome5(s10))
	fmt.Printf("方法六结果: %s\n", longestPalindrome6(s10))
}
