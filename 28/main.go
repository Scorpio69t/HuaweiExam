package main

import (
	"fmt"
	"strings"
)

// strStr 暴力匹配法：逐个位置尝试匹配
func strStr(haystack string, needle string) int {
	if needle == "" {
		return 0
	}
	if len(haystack) < len(needle) {
		return -1
	}

	n := len(haystack)
	m := len(needle)

	// 遍历所有可能的起始位置
	for i := 0; i <= n-m; i++ {
		// 检查从位置i开始的子串是否匹配
		if haystack[i:i+m] == needle {
			return i
		}
	}

	return -1
}

// strStrKMP KMP算法：利用部分匹配表优化
func strStrKMP(haystack string, needle string) int {
	if needle == "" {
		return 0
	}
	if len(haystack) < len(needle) {
		return -1
	}

	// 构建部分匹配表（next数组）
	next := buildNext(needle)

	n := len(haystack)
	m := len(needle)
	i, j := 0, 0

	for i < n && j < m {
		if j == -1 || haystack[i] == needle[j] {
			i++
			j++
		} else {
			j = next[j]
		}
	}

	if j == m {
		return i - m
	}
	return -1
}

// buildNext 构建KMP算法的next数组
func buildNext(pattern string) []int {
	m := len(pattern)
	next := make([]int, m)
	next[0] = -1

	i, j := 0, -1
	for i < m-1 {
		if j == -1 || pattern[i] == pattern[j] {
			i++
			j++
			if pattern[i] != pattern[j] {
				next[i] = j
			} else {
				next[i] = next[j]
			}
		} else {
			j = next[j]
		}
	}

	return next
}

// strStrBoyerMoore Boyer-Moore算法：坏字符规则（使用最后出现位置表）
func strStrBoyerMoore(haystack string, needle string) int {
	if needle == "" {
		return 0
	}
	if len(haystack) < len(needle) {
		return -1
	}

	// 构建坏字符最后出现位置表：不存在则为 -1
	last := buildBadCharTable(needle)

	n := len(haystack)
	m := len(needle)
	i := 0 // 窗口起始位置

	for i <= n-m {
		j := m - 1
		// 从右向左比较
		for j >= 0 && haystack[i+j] == needle[j] {
			j--
		}
		if j < 0 {
			return i
		}
		// 计算坏字符位移：max(1, j - last[text[i+j]])
		bc := haystack[i+j]
		shift := j - last[int(bc)]
		if shift < 1 {
			shift = 1
		}
		i += shift
	}

	return -1
}

// buildBadCharTable 构建“最后出现位置表”；字符未出现记为 -1
func buildBadCharTable(pattern string) []int {
	table := make([]int, 256)
	for i := range table {
		table[i] = -1
	}
	for i := 0; i < len(pattern); i++ {
		table[int(pattern[i])] = i
	}
	return table
}

// strStrRabinKarp Rabin-Karp算法：使用哈希函数
func strStrRabinKarp(haystack string, needle string) int {
	if needle == "" {
		return 0
	}
	if len(haystack) < len(needle) {
		return -1
	}

	n := len(haystack)
	m := len(needle)

	// 计算needle的哈希值
	needleHash := hash(needle)
	windowHash := hash(haystack[:m])

	// 如果第一个窗口匹配，直接返回
	if windowHash == needleHash && haystack[:m] == needle {
		return 0
	}

	// 滚动哈希
	power := 1
	for i := 0; i < m-1; i++ {
		power = (power * 26) % 1000000007
	}

	for i := 1; i <= n-m; i++ {
		// 移除最左边的字符，添加最右边的字符
		windowHash = (windowHash - int(haystack[i-1])*power) % 1000000007
		windowHash = (windowHash*26 + int(haystack[i+m-1])) % 1000000007

		if windowHash < 0 {
			windowHash += 1000000007
		}

		// 哈希值匹配时，进行字符串比较
		if windowHash == needleHash && haystack[i:i+m] == needle {
			return i
		}
	}

	return -1
}

// hash 计算字符串的哈希值
func hash(s string) int {
	h := 0
	for _, c := range s {
		h = (h*26 + int(c)) % 1000000007
	}
	return h
}

// strStrSunday Sunday算法：利用好后缀规则
func strStrSunday(haystack string, needle string) int {
	if needle == "" {
		return 0
	}
	if len(haystack) < len(needle) {
		return -1
	}

	// 构建Sunday算法的移动表
	shift := buildSundayShift(needle)

	n := len(haystack)
	m := len(needle)
	i := 0

	for i <= n-m {
		// 检查当前窗口是否匹配
		if haystack[i:i+m] == needle {
			return i
		}

		// 如果已经到达末尾，无法继续移动
		if i+m >= n {
			break
		}

		// 使用Sunday规则移动
		nextChar := haystack[i+m]
		if s, exists := shift[nextChar]; exists {
			i += s
		} else {
			i += m + 1
		}
	}

	return -1
}

// buildSundayShift 构建Sunday算法的移动表
func buildSundayShift(pattern string) map[byte]int {
	shift := make(map[byte]int)
	m := len(pattern)

	for i := 0; i < m; i++ {
		shift[pattern[i]] = m - i
	}

	return shift
}

// strStrOptimized 优化的暴力匹配：提前退出
func strStrOptimized(haystack string, needle string) int {
	if needle == "" {
		return 0
	}
	if len(haystack) < len(needle) {
		return -1
	}

	n := len(haystack)
	m := len(needle)

	// 遍历所有可能的起始位置
	for i := 0; i <= n-m; i++ {
		// 检查从位置i开始的子串是否匹配
		matched := true
		for j := 0; j < m; j++ {
			if haystack[i+j] != needle[j] {
				matched = false
				break
			}
		}
		if matched {
			return i
		}
	}

	return -1
}

// strStrBuiltin 使用Go内置的strings.Index函数
func strStrBuiltin(haystack string, needle string) int {
	return strings.Index(haystack, needle)
}

func main() {
	fmt.Println("字符串匹配问题测试")
	fmt.Println("==================")

	// 测试用例1
	fmt.Println("\n测试用例1:")
	haystack1 := "sadbutsad"
	needle1 := "sad"
	fmt.Printf("haystack: %s\n", haystack1)
	fmt.Printf("needle: %s\n", needle1)

	result1 := strStr(haystack1, needle1)
	fmt.Printf("暴力匹配结果: %d\n", result1)

	result1KMP := strStrKMP(haystack1, needle1)
	fmt.Printf("KMP算法结果: %d\n", result1KMP)

	result1BM := strStrBoyerMoore(haystack1, needle1)
	fmt.Printf("Boyer-Moore结果: %d\n", result1BM)

	result1RK := strStrRabinKarp(haystack1, needle1)
	fmt.Printf("Rabin-Karp结果: %d\n", result1RK)

	result1Sunday := strStrSunday(haystack1, needle1)
	fmt.Printf("Sunday算法结果: %d\n", result1Sunday)

	result1Opt := strStrOptimized(haystack1, needle1)
	fmt.Printf("优化暴力匹配结果: %d\n", result1Opt)

	result1Builtin := strStrBuiltin(haystack1, needle1)
	fmt.Printf("内置函数结果: %d\n", result1Builtin)

	fmt.Printf("期望结果: 0\n")

	// 测试用例2
	fmt.Println("\n测试用例2:")
	haystack2 := "leetcode"
	needle2 := "leeto"
	fmt.Printf("haystack: %s\n", haystack2)
	fmt.Printf("needle: %s\n", needle2)

	result2 := strStr(haystack2, needle2)
	fmt.Printf("暴力匹配结果: %d\n", result2)

	result2KMP := strStrKMP(haystack2, needle2)
	fmt.Printf("KMP算法结果: %d\n", result2KMP)

	result2BM := strStrBoyerMoore(haystack2, needle2)
	fmt.Printf("Boyer-Moore结果: %d\n", result2BM)

	result2RK := strStrRabinKarp(haystack2, needle2)
	fmt.Printf("Rabin-Karp结果: %d\n", result2RK)

	result2Sunday := strStrSunday(haystack2, needle2)
	fmt.Printf("Sunday算法结果: %d\n", result2Sunday)

	result2Opt := strStrOptimized(haystack2, needle2)
	fmt.Printf("优化暴力匹配结果: %d\n", result2Opt)

	result2Builtin := strStrBuiltin(haystack2, needle2)
	fmt.Printf("内置函数结果: %d\n", result2Builtin)

	fmt.Printf("期望结果: -1\n")

	// 测试用例3：空字符串
	fmt.Println("\n测试用例3 - 空字符串:")
	haystack3 := "hello"
	needle3 := ""
	fmt.Printf("haystack: %s\n", haystack3)
	fmt.Printf("needle: %s\n", needle3)

	result3 := strStr(haystack3, needle3)
	fmt.Printf("结果: %d\n", result3)
	fmt.Printf("期望结果: 0\n")

	// 测试用例4：完全匹配
	fmt.Println("\n测试用例4 - 完全匹配:")
	haystack4 := "abc"
	needle4 := "abc"
	fmt.Printf("haystack: %s\n", haystack4)
	fmt.Printf("needle: %s\n", needle4)

	result4 := strStr(haystack4, needle4)
	fmt.Printf("结果: %d\n", result4)
	fmt.Printf("期望结果: 0\n")

	// 测试用例5：在末尾匹配
	fmt.Println("\n测试用例5 - 在末尾匹配:")
	haystack5 := "hello world"
	needle5 := "world"
	fmt.Printf("haystack: %s\n", haystack5)
	fmt.Printf("needle: %s\n", needle5)

	result5 := strStr(haystack5, needle5)
	fmt.Printf("结果: %d\n", result5)
	fmt.Printf("期望结果: 6\n")

	// 测试用例6：重复模式
	fmt.Println("\n测试用例6 - 重复模式:")
	haystack6 := "aaaaa"
	needle6 := "aa"
	fmt.Printf("haystack: %s\n", haystack6)
	fmt.Printf("needle: %s\n", needle6)

	result6 := strStr(haystack6, needle6)
	fmt.Printf("结果: %d\n", result6)
	fmt.Printf("期望结果: 0\n")

	// 性能测试
	fmt.Println("\n性能测试:")
	largeHaystack := strings.Repeat("a", 10000) + "b" + strings.Repeat("a", 10000)
	largeNeedle := "b"

	fmt.Printf("大字符串长度: %d\n", len(largeHaystack))
	fmt.Printf("查找模式: %s\n", largeNeedle)

	// 测试各种算法的性能
	resultLarge := strStr(largeHaystack, largeNeedle)
	fmt.Printf("暴力匹配结果: %d\n", resultLarge)

	resultLargeKMP := strStrKMP(largeHaystack, largeNeedle)
	fmt.Printf("KMP算法结果: %d\n", resultLargeKMP)

	resultLargeBuiltin := strStrBuiltin(largeHaystack, largeNeedle)
	fmt.Printf("内置函数结果: %d\n", resultLargeBuiltin)
}
