package main

import (
	"fmt"
	"strings"
)

// 方法一：横向扫描（推荐）
// 时间复杂度：O(S)，空间复杂度：O(1)
func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	prefix := strs[0]
	for i := 1; i < len(strs); i++ {
		prefix = commonPrefix(prefix, strs[i])
		if prefix == "" {
			return ""
		}
	}

	return prefix
}

// 辅助函数：求两个字符串的公共前缀
func commonPrefix(str1, str2 string) string {
	length := min(len(str1), len(str2))
	index := 0
	for index < length && str1[index] == str2[index] {
		index++
	}
	return str1[:index]
}

// 方法二：纵向扫描
// 时间复杂度：O(S)，空间复杂度：O(1)
func longestCommonPrefixVertical(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	// 找到最短字符串的长度
	minLength := len(strs[0])
	for _, str := range strs {
		if len(str) < minLength {
			minLength = len(str)
		}
	}

	// 逐列比较
	for i := 0; i < minLength; i++ {
		char := strs[0][i]
		for j := 1; j < len(strs); j++ {
			if strs[j][i] != char {
				return strs[0][:i]
			}
		}
	}

	return strs[0][:minLength]
}

// 方法三：分治算法
// 时间复杂度：O(S)，空间复杂度：O(m log n)
func longestCommonPrefixDivide(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	if len(strs) == 1 {
		return strs[0]
	}

	return longestCommonPrefixDivideHelper(strs, 0, len(strs)-1)
}

func longestCommonPrefixDivideHelper(strs []string, start, end int) string {
	if start == end {
		return strs[start]
	}

	mid := (start + end) / 2
	leftPrefix := longestCommonPrefixDivideHelper(strs, start, mid)
	rightPrefix := longestCommonPrefixDivideHelper(strs, mid+1, end)

	return commonPrefix(leftPrefix, rightPrefix)
}

// 方法四：二分查找
// 时间复杂度：O(S × log m)，空间复杂度：O(1)
func longestCommonPrefixBinary(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	// 找到最短字符串的长度
	minLength := len(strs[0])
	for _, str := range strs {
		if len(str) < minLength {
			minLength = len(str)
		}
	}

	// 二分查找
	left, right := 0, minLength
	for left < right {
		mid := (left + right + 1) / 2
		if isCommonPrefix(strs, mid) {
			left = mid
		} else {
			right = mid - 1
		}
	}

	return strs[0][:left]
}

// 辅助函数：检查前length个字符是否为公共前缀
func isCommonPrefix(strs []string, length int) bool {
	if length == 0 {
		return true
	}

	prefix := strs[0][:length]
	for i := 1; i < len(strs); i++ {
		if !strings.HasPrefix(strs[i], prefix) {
			return false
		}
	}
	return true
}

// 方法五：优化的横向扫描
func longestCommonPrefixOptimized(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	// 找到最短字符串
	shortest := strs[0]
	for _, str := range strs {
		if len(str) < len(shortest) {
			shortest = str
		}
	}

	// 以最短字符串为基准
	for i := 0; i < len(shortest); i++ {
		char := shortest[i]
		for j := 0; j < len(strs); j++ {
			if strs[j][i] != char {
				return shortest[:i]
			}
		}
	}

	return shortest
}

// 方法六：Trie树解法（适用于大量字符串）
func longestCommonPrefixTrie(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	// 构建Trie树
	trie := NewTrieNode()
	for _, str := range strs {
		trie.Insert(str)
	}

	// 查找最长公共前缀
	return trie.FindLongestCommonPrefix()
}

// Trie树节点
type TrieNode struct {
	children map[byte]*TrieNode
	isEnd    bool
	count    int // 记录有多少个字符串经过这个节点
}

func NewTrieNode() *TrieNode {
	return &TrieNode{
		children: make(map[byte]*TrieNode),
		isEnd:    false,
		count:    0,
	}
}

func (t *TrieNode) Insert(word string) {
	node := t
	for i := 0; i < len(word); i++ {
		char := word[i]
		if node.children[char] == nil {
			node.children[char] = NewTrieNode()
		}
		node = node.children[char]
		node.count++
	}
	node.isEnd = true
}

func (t *TrieNode) FindLongestCommonPrefix() string {
	if len(t.children) != 1 {
		return ""
	}

	var result strings.Builder
	node := t
	totalStrings := 0

	// 计算总字符串数
	for _, child := range node.children {
		totalStrings = child.count
		break
	}

	for len(node.children) == 1 && !node.isEnd {
		var char byte
		var child *TrieNode
		for c, ch := range node.children {
			char = c
			child = ch
			break
		}
		if child.count < totalStrings {
			break
		}
		result.WriteByte(char)
		node = child
	}

	return result.String()
}

// 辅助函数：求最小值
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// 辅助函数：验证结果是否正确
func validateResult(strs []string, result string) bool {
	if len(strs) == 0 {
		return result == ""
	}
	if len(strs) == 1 {
		return result == strs[0]
	}

	// 检查result是否为所有字符串的前缀
	for _, str := range strs {
		if !strings.HasPrefix(str, result) {
			return false
		}
	}

	// 检查是否为最长前缀
	for i := len(result) + 1; i <= len(strs[0]); i++ {
		if i > len(strs[0]) {
			break
		}
		extendedPrefix := strs[0][:i]
		allHavePrefix := true
		for _, str := range strs {
			if !strings.HasPrefix(str, extendedPrefix) {
				allHavePrefix = false
				break
			}
		}
		if allHavePrefix {
			return false // 存在更长的前缀
		}
	}

	return true
}

// 辅助函数：比较两个结果是否相等
func compareResults(result1, result2 string) bool {
	return result1 == result2
}

func main() {
	fmt.Println("=== 14. 最长公共前缀 ===")

	// 测试用例1
	strs1 := []string{"flower", "flow", "flight"}
	fmt.Printf("测试用例1: %v\n", strs1)
	
	result1 := longestCommonPrefix(strs1)
	fmt.Printf("横向扫描解法结果: %s\n", result1)
	
	result1Vertical := longestCommonPrefixVertical(strs1)
	fmt.Printf("纵向扫描解法结果: %s\n", result1Vertical)
	
	result1Divide := longestCommonPrefixDivide(strs1)
	fmt.Printf("分治算法解法结果: %s\n", result1Divide)
	
	result1Binary := longestCommonPrefixBinary(strs1)
	fmt.Printf("二分查找解法结果: %s\n", result1Binary)
	
	result1Optimized := longestCommonPrefixOptimized(strs1)
	fmt.Printf("优化解法结果: %s\n", result1Optimized)
	
	// 验证结果
	if validateResult(strs1, result1) {
		fmt.Println("✅ 结果验证通过！")
	} else {
		fmt.Println("❌ 结果验证失败！")
	}
	fmt.Println()

	// 测试用例2
	strs2 := []string{"dog", "racecar", "car"}
	fmt.Printf("测试用例2: %v\n", strs2)
	
	result2 := longestCommonPrefix(strs2)
	fmt.Printf("横向扫描解法结果: %s\n", result2)
	
	// 验证结果
	if validateResult(strs2, result2) {
		fmt.Println("✅ 结果验证通过！")
	} else {
		fmt.Println("❌ 结果验证失败！")
	}
	fmt.Println()

	// 测试用例3
	strs3 := []string{"interspecies", "interstellar", "interstate"}
	fmt.Printf("测试用例3: %v\n", strs3)
	
	result3 := longestCommonPrefix(strs3)
	fmt.Printf("横向扫描解法结果: %s\n", result3)
	
	// 验证结果
	if validateResult(strs3, result3) {
		fmt.Println("✅ 结果验证通过！")
	} else {
		fmt.Println("❌ 结果验证失败！")
	}
	fmt.Println()

	// 边界测试用例
	testCases := []struct {
		strs []string
		desc string
	}{
		{[]string{"throne", "throne"}, "相同字符串"},
		{[]string{"a"}, "单个字符串"},
		{[]string{}, "空数组"},
		{[]string{"", "b", "c"}, "包含空字符串"},
		{[]string{"a", "b", "c"}, "无公共前缀"},
		{[]string{"abc", "abcde", "abcdef"}, "最短字符串为前缀"},
		{[]string{"hello", "hell", "he"}, "递减长度"},
		{[]string{"test", "testing", "tested"}, "相同前缀"},
	}

	for _, tc := range testCases {
		fmt.Printf("%s: %v\n", tc.desc, tc.strs)
		result := longestCommonPrefix(tc.strs)
		fmt.Printf("结果: %s\n", result)
		
		// 验证结果
		if validateResult(tc.strs, result) {
			fmt.Println("✅ 验证通过")
		} else {
			fmt.Println("❌ 验证失败")
		}
		fmt.Println()
	}

	// 算法正确性验证
	fmt.Println("=== 算法正确性验证 ===")
	verifyStrs := []string{"flower", "flow", "flight"}
	
	fmt.Printf("验证字符串数组: %v\n", verifyStrs)
	
	verifyResult1 := longestCommonPrefix(verifyStrs)
	verifyResult2 := longestCommonPrefixVertical(verifyStrs)
	verifyResult3 := longestCommonPrefixDivide(verifyStrs)
	verifyResult4 := longestCommonPrefixBinary(verifyStrs)
	verifyResult5 := longestCommonPrefixOptimized(verifyStrs)

	fmt.Printf("横向扫描解法: %s\n", verifyResult1)
	fmt.Printf("纵向扫描解法: %s\n", verifyResult2)
	fmt.Printf("分治算法解法: %s\n", verifyResult3)
	fmt.Printf("二分查找解法: %s\n", verifyResult4)
	fmt.Printf("优化解法: %s\n", verifyResult5)

	// 验证所有解法结果一致
	if compareResults(verifyResult1, verifyResult2) && 
	   compareResults(verifyResult2, verifyResult3) && 
	   compareResults(verifyResult3, verifyResult4) && 
	   compareResults(verifyResult4, verifyResult5) {
		fmt.Println("✅ 所有解法结果一致！")
		
		// 验证结果正确性
		if validateResult(verifyStrs, verifyResult1) {
			fmt.Println("✅ 结果验证通过！")
		} else {
			fmt.Println("❌ 结果验证失败！")
		}
	} else {
		fmt.Println("❌ 解法结果不一致，需要检查！")
	}

	// 性能测试
	fmt.Println("\n=== 性能测试 ===")
	
	// 创建大量字符串进行测试
	largeStrs := make([]string, 1000)
	for i := 0; i < 1000; i++ {
		largeStrs[i] = fmt.Sprintf("prefix_%d_suffix", i)
	}

	fmt.Printf("大数组测试: 长度=%d\n", len(largeStrs))
	result := longestCommonPrefix(largeStrs)
	fmt.Printf("横向扫描解法结果: %s\n", result)
	
	// 验证大数组结果
	if validateResult(largeStrs, result) {
		fmt.Println("✅ 大数组结果验证通过！")
	} else {
		fmt.Println("❌ 大数组结果验证失败！")
	}

	// 额外测试
	fmt.Println("\n=== 额外测试 ===")
	extraStrs := []string{"apple", "app", "application", "apply"}
	fmt.Printf("额外测试: %v\n", extraStrs)
	extraResult := longestCommonPrefix(extraStrs)
	fmt.Printf("结果: %s\n", extraResult)
	
	if validateResult(extraStrs, extraResult) {
		fmt.Println("✅ 额外测试结果验证通过！")
	} else {
		fmt.Println("❌ 额外测试结果验证失败！")
	}
}
