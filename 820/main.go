package main

import (
	"fmt"
	"sort"
	"strings"
)

// 字典树节点
type TrieNode struct {
	children map[byte]*TrieNode
	isEnd    bool
}

// 创建新的字典树节点
func newTrieNode() *TrieNode {
	return &TrieNode{
		children: make(map[byte]*TrieNode),
		isEnd:    false,
	}
}

// 方法一：反向字典树（推荐）
// 时间复杂度：O(n*m)，空间复杂度：O(n*m)
func minimumLengthEncoding(words []string) int {
	if len(words) == 0 {
		return 0
	}

	// 去重处理
	wordSet := make(map[string]bool)
	for _, word := range words {
		wordSet[word] = true
	}

	// 构建反向字典树
	root := newTrieNode()
	for word := range wordSet {
		insertReverse(root, word)
	}

	// 计算编码长度
	return calculateLength(root, 0)
}

// 反向插入单词到字典树
func insertReverse(root *TrieNode, word string) {
	node := root
	// 从末尾开始插入
	for i := len(word) - 1; i >= 0; i-- {
		if node.children[word[i]] == nil {
			node.children[word[i]] = newTrieNode()
		}
		node = node.children[word[i]]
	}
	node.isEnd = true
}

// 计算字典树编码长度
func calculateLength(node *TrieNode, depth int) int {
	if len(node.children) == 0 {
		// 叶子节点，表示一个根单词
		return depth + 1 // +1 for '#'
	}

	length := 0
	for _, child := range node.children {
		length += calculateLength(child, depth+1)
	}
	return length
}

// 方法二：哈希表暴力检查
// 时间复杂度：O(n²*m)，空间复杂度：O(n*m)
func minimumLengthEncodingBruteForce(words []string) int {
	if len(words) == 0 {
		return 0
	}

	// 去重处理
	wordSet := make(map[string]bool)
	for _, word := range words {
		wordSet[word] = true
	}

	uniqueWords := make([]string, 0, len(wordSet))
	for word := range wordSet {
		uniqueWords = append(uniqueWords, word)
	}

	// 检查每个单词是否为其他单词的后缀
	needed := make(map[string]bool)
	for _, word := range uniqueWords {
		needed[word] = true
	}

	for _, word := range uniqueWords {
		for _, other := range uniqueWords {
			if word != other && isSuffix(word, other) {
				needed[word] = false
				break
			}
		}
	}

	// 计算总长度
	length := 0
	for word, need := range needed {
		if need {
			length += len(word) + 1 // +1 for '#'
		}
	}

	return length
}

// 检查word是否为target的后缀
func isSuffix(word, target string) bool {
	if len(word) > len(target) {
		return false
	}
	return target[len(target)-len(word):] == word
}

// 方法三：排序后检查
// 时间复杂度：O(n*m*log n)，空间复杂度：O(n*m)
func minimumLengthEncodingSorted(words []string) int {
	if len(words) == 0 {
		return 0
	}

	// 去重处理
	wordSet := make(map[string]bool)
	for _, word := range words {
		wordSet[word] = true
	}

	uniqueWords := make([]string, 0, len(wordSet))
	for word := range wordSet {
		uniqueWords = append(uniqueWords, word)
	}

	// 按长度降序排序
	sort.Slice(uniqueWords, func(i, j int) bool {
		return len(uniqueWords[i]) > len(uniqueWords[j])
	})

	// 检查每个单词是否为前面单词的后缀
	needed := make([]bool, len(uniqueWords))
	for i := range needed {
		needed[i] = true
	}

	for i := 0; i < len(uniqueWords); i++ {
		if !needed[i] {
			continue
		}
		for j := i + 1; j < len(uniqueWords); j++ {
			if isSuffix(uniqueWords[j], uniqueWords[i]) {
				needed[j] = false
			}
		}
	}

	// 计算总长度
	length := 0
	for i, word := range uniqueWords {
		if needed[i] {
			length += len(word) + 1 // +1 for '#'
		}
	}

	return length
}

// 方法四：集合操作
// 时间复杂度：O(n²*m)，空间复杂度：O(n*m)
func minimumLengthEncodingSet(words []string) int {
	if len(words) == 0 {
		return 0
	}

	// 创建单词集合
	wordSet := make(map[string]bool)
	for _, word := range words {
		wordSet[word] = true
	}

	// 移除所有后缀单词
	for word := range wordSet {
		for i := 1; i < len(word); i++ {
			suffix := word[i:]
			if wordSet[suffix] {
				delete(wordSet, suffix)
			}
		}
	}

	// 计算剩余单词的编码长度
	length := 0
	for word := range wordSet {
		length += len(word) + 1 // +1 for '#'
	}

	return length
}

// 辅助函数：创建go.mod
func createGoMod() {
	// 这里只是为了演示，实际中go.mod已经存在
}

func runTests() {
	type testCase struct {
		words    []string
		expected int
		desc     string
	}

	tests := []testCase{
		{[]string{"time", "me", "bell"}, 10, "示例1"},
		{[]string{"t"}, 2, "示例2"},
		{[]string{}, 0, "空数组"},
		{[]string{"abc", "abc"}, 4, "重复单词"},
		{[]string{"time", "me", "time"}, 5, "包含关系"},
		{[]string{"abc", "def"}, 8, "无关系"},
		{[]string{"time", "me", "e"}, 5, "多层包含"},
		{[]string{"abcdefg"}, 8, "长单词"},
		{[]string{"a", "a", "a"}, 2, "所有相同"},
		{[]string{"time", "atime", "btime"}, 12, "复杂场景"},
	}

	fmt.Println("=== 820. 单词的压缩编码 - 测试 ===")
	for i, tc := range tests {
		r1 := minimumLengthEncoding(tc.words)
		r2 := minimumLengthEncodingBruteForce(tc.words)
		r3 := minimumLengthEncodingSorted(tc.words)
		r4 := minimumLengthEncodingSet(tc.words)

		ok := (r1 == tc.expected) && (r2 == tc.expected) && (r3 == tc.expected) && (r4 == tc.expected)
		status := "✅"
		if !ok {
			status = "❌"
		}

		fmt.Printf("用例 %d: %s\n", i+1, tc.desc)
		fmt.Printf("输入: %v\n", tc.words)
		fmt.Printf("期望: %d\n", tc.expected)
		fmt.Printf("反向字典树: %d, 哈希表暴力: %d, 排序检查: %d, 集合操作: %d\n", r1, r2, r3, r4)
		fmt.Printf("结果: %s\n", status)
		fmt.Println(strings.Repeat("-", 60))
	}
}

func main() {
	runTests()
}
