package main

import (
	"fmt"
	"strings"
	"time"
)

// 解法一：动态规划（推荐解法）
// 时间复杂度：O(n²×m)，空间复杂度：O(n)
func wordBreak(s string, wordDict []string) bool {
	n := len(s)
	if n == 0 {
		return true
	}

	// 将字典转为哈希集合，提高查找效率
	wordSet := make(map[string]bool)
	for _, word := range wordDict {
		wordSet[word] = true
	}

	// dp[i] 表示 s[0:i] 能否被拆分
	dp := make([]bool, n+1)
	dp[0] = true // 空字符串可以被拆分

	// 遍历每个位置
	for i := 1; i <= n; i++ {
		// 尝试所有可能的分割点
		for j := 0; j < i; j++ {
			// 如果 s[0:j] 可以拆分，且 s[j:i] 在字典中
			if dp[j] && wordSet[s[j:i]] {
				dp[i] = true
				break // 找到一种拆分方式即可
			}
		}
	}

	return dp[n]
}

// 解法二：DFS + 记忆化递归
// 时间复杂度：O(n²×m)，空间复杂度：O(n)
func wordBreakDFS(s string, wordDict []string) bool {
	wordSet := make(map[string]bool)
	for _, word := range wordDict {
		wordSet[word] = true
	}

	memo := make(map[int]bool)

	var dfs func(start int) bool
	dfs = func(start int) bool {
		// 到达字符串末尾，说明完全匹配
		if start == len(s) {
			return true
		}

		// 检查记忆化结果
		if val, exists := memo[start]; exists {
			return val
		}

		// 尝试所有可能的子串
		for end := start + 1; end <= len(s); end++ {
			substr := s[start:end]

			// 如果当前子串在字典中，且剩余部分可以拆分
			if wordSet[substr] && dfs(end) {
				memo[start] = true
				return true
			}
		}

		memo[start] = false
		return false
	}

	return dfs(0)
}

// 解法三：BFS（广度优先搜索）
// 时间复杂度：O(n²×m)，空间复杂度：O(n)
func wordBreakBFS(s string, wordDict []string) bool {
	if len(s) == 0 {
		return true
	}
	if len(wordDict) == 0 {
		return false
	}

	wordSet := make(map[string]bool)
	for _, word := range wordDict {
		wordSet[word] = true
	}

	n := len(s)
	visited := make([]bool, n)
	queue := []int{0}

	for len(queue) > 0 {
		start := queue[0]
		queue = queue[1:]

		// 如果已经访问过这个位置，跳过
		if visited[start] {
			continue
		}

		visited[start] = true

		// 尝试从当前位置开始的所有子串
		for end := start + 1; end <= n; end++ {
			substr := s[start:end]

			if wordSet[substr] {
				if end == n {
					return true // 到达末尾，找到完整拆分
				}

				if !visited[end] {
					queue = append(queue, end)
				}
			}
		}
	}

	return false
}

// 解法四：Trie + 动态规划（字典较大时的优化）
// 时间复杂度：O(n²×k)，空间复杂度：O(总字符数)
type TrieNode struct {
	children map[rune]*TrieNode
	isEnd    bool
}

type Trie struct {
	root *TrieNode
}

func newTrie() *Trie {
	return &Trie{
		root: &TrieNode{
			children: make(map[rune]*TrieNode),
			isEnd:    false,
		},
	}
}

func (t *Trie) insert(word string) {
	node := t.root
	for _, char := range word {
		if _, exists := node.children[char]; !exists {
			node.children[char] = &TrieNode{
				children: make(map[rune]*TrieNode),
				isEnd:    false,
			}
		}
		node = node.children[char]
	}
	node.isEnd = true
}

func wordBreakTrie(s string, wordDict []string) bool {
	if len(s) == 0 {
		return true
	}
	if len(wordDict) == 0 {
		return false
	}

	// 构建Trie树
	trie := newTrie()
	for _, word := range wordDict {
		trie.insert(word)
	}

	n := len(s)
	dp := make([]bool, n+1)
	dp[0] = true

	for i := 1; i <= n; i++ {
		node := trie.root

		// 从位置i-1向前查找
		for j := i - 1; j >= 0; j-- {
			char := rune(s[j])

			// 如果Trie中没有这个字符的路径，停止查找
			if _, exists := node.children[char]; !exists {
				break
			}

			node = node.children[char]

			// 如果当前节点是单词结尾，且前面的部分可以拆分
			if node.isEnd && dp[j] {
				dp[i] = true
				break
			}
		}
	}

	return dp[n]
}

// 解法五：优化的动态规划（长度过滤）
// 时间复杂度：O(n×maxLen×m)，空间复杂度：O(n)
func wordBreakOptimized(s string, wordDict []string) bool {
	if len(s) == 0 {
		return true
	}
	if len(wordDict) == 0 {
		return false
	}

	wordSet := make(map[string]bool)
	minLen, maxLen := len(wordDict[0]), len(wordDict[0])

	// 构建字典集合并记录最短和最长单词长度
	for _, word := range wordDict {
		wordSet[word] = true
		if len(word) < minLen {
			minLen = len(word)
		}
		if len(word) > maxLen {
			maxLen = len(word)
		}
	}

	n := len(s)
	dp := make([]bool, n+1)
	dp[0] = true

	for i := minLen; i <= n; i++ {
		// 只检查可能的单词长度范围
		for length := minLen; length <= maxLen && i-length >= 0; length++ {
			if dp[i-length] && wordSet[s[i-length:i]] {
				dp[i] = true
				break
			}
		}
	}

	return dp[n]
}

// 辅助函数：打印拆分方案（调试用）
func findAllWordBreaks(s string, wordDict []string) [][]string {
	wordSet := make(map[string]bool)
	for _, word := range wordDict {
		wordSet[word] = true
	}

	var result [][]string
	var path []string

	var backtrack func(start int)
	backtrack = func(start int) {
		if start == len(s) {
			// 复制当前路径
			solution := make([]string, len(path))
			copy(solution, path)
			result = append(result, solution)
			return
		}

		for end := start + 1; end <= len(s); end++ {
			substr := s[start:end]
			if wordSet[substr] {
				path = append(path, substr)
				backtrack(end)
				path = path[:len(path)-1] // 回溯
			}
		}
	}

	backtrack(0)
	return result
}

// 测试函数
func testWordBreak() {
	testCases := []struct {
		s        string
		wordDict []string
		expected bool
		desc     string
	}{
		{
			"leetcode",
			[]string{"leet", "code"},
			true,
			"示例1：基本拆分",
		},
		{
			"applepenapple",
			[]string{"apple", "pen"},
			true,
			"示例2：重复使用单词",
		},
		{
			"catsandog",
			[]string{"cats", "dog", "sand", "and", "cat"},
			false,
			"示例3：无法完全拆分",
		},
		{
			"",
			[]string{"a"},
			true,
			"空字符串",
		},
		{
			"a",
			[]string{"a"},
			true,
			"单字符匹配",
		},
		{
			"a",
			[]string{"b"},
			false,
			"单字符不匹配",
		},
		{
			"aaaaaaa",
			[]string{"aaaa", "aaa"},
			true,
			"多种拆分方式",
		},
		{
			"aaaaaaa",
			[]string{"aaaa", "aa"},
			false,
			"无法完整拆分",
		},
		{
			"abcd",
			[]string{"a", "abc", "b", "cd"},
			true,
			"多种单词选择",
		},
		{
			"cars",
			[]string{"car", "ca", "rs"},
			true,
			"前缀重叠",
		},
		{
			"aaaaa",
			[]string{"aa", "aaa"},
			true,
			"长度组合匹配",
		},
		{
			"goalspecial",
			[]string{"go", "goal", "goals", "special"},
			true,
			"复杂拆分",
		},
	}

	fmt.Println("=== 单词拆分测试 ===")
	fmt.Println()

	for i, tc := range testCases {
		// 测试不同算法
		result1 := wordBreak(tc.s, tc.wordDict)
		result2 := wordBreakDFS(tc.s, tc.wordDict)
		result3 := wordBreakBFS(tc.s, tc.wordDict)
		result4 := wordBreakTrie(tc.s, tc.wordDict)
		result5 := wordBreakOptimized(tc.s, tc.wordDict)

		status := "✅"
		if result1 != tc.expected {
			status = "❌"
		}

		fmt.Printf("测试 %d: %s\n", i+1, tc.desc)
		fmt.Printf("输入: s=\"%s\", wordDict=%v\n", tc.s, tc.wordDict)
		fmt.Printf("期望: %t, 实际: %t\n", tc.expected, result1)

		// 验证所有算法结果一致
		consistent := result1 == result2 && result2 == result3 &&
			result3 == result4 && result4 == result5
		fmt.Printf("算法一致性: %t\n", consistent)

		// 如果可以拆分，展示拆分方案
		if result1 && len(tc.s) <= 20 {
			solutions := findAllWordBreaks(tc.s, tc.wordDict)
			fmt.Printf("拆分方案: %v\n", solutions)
		}

		fmt.Printf("结果: %s\n", status)
		fmt.Println(strings.Repeat("-", 50))
	}
}

// 性能测试
func benchmarkWordBreak() {
	fmt.Println()
	fmt.Println("=== 性能测试 ===")
	fmt.Println()

	// 构造测试数据
	testData := []struct {
		s        string
		wordDict []string
		desc     string
	}{
		{
			generateString(50, "abc"),
			[]string{"a", "ab", "abc", "bc", "c"},
			"短字符串，小字典",
		},
		{
			generateString(100, "abcdef"),
			generateWordDict(50, 3),
			"中等字符串，中等字典",
		},
		{
			generateString(200, "abcdefghij"),
			generateWordDict(100, 5),
			"长字符串，大字典",
		},
		{
			strings.Repeat("a", 100),
			[]string{"a", "aa", "aaa"},
			"最坏情况：重复字符",
		},
	}

	algorithms := []struct {
		name string
		fn   func(string, []string) bool
	}{
		{"动态规划", wordBreak},
		{"DFS+记忆化", wordBreakDFS},
		{"BFS", wordBreakBFS},
		{"Trie+DP", wordBreakTrie},
		{"优化DP", wordBreakOptimized},
	}

	for _, data := range testData {
		fmt.Printf("%s:\n", data.desc)
		fmt.Printf("  字符串长度: %d, 字典大小: %d\n", len(data.s), len(data.wordDict))

		for _, algo := range algorithms {
			start := time.Now()
			result := algo.fn(data.s, data.wordDict)
			duration := time.Since(start)

			fmt.Printf("  %s: %t, 耗时 %v\n", algo.name, result, duration)
		}
		fmt.Println()
	}
}

// 生成测试字符串
func generateString(length int, charset string) string {
	result := make([]byte, length)
	for i := 0; i < length; i++ {
		result[i] = charset[i%len(charset)]
	}
	return string(result)
}

// 生成测试字典
func generateWordDict(count, maxLen int) []string {
	words := make([]string, count)
	charset := "abcdefghijklmnopqrstuvwxyz"

	for i := 0; i < count; i++ {
		length := (i % maxLen) + 1
		word := make([]byte, length)
		for j := 0; j < length; j++ {
			word[j] = charset[(i+j)%len(charset)]
		}
		words[i] = string(word)
	}

	return words
}

// 演示算法过程
func demonstrateAlgorithm() {
	fmt.Println()
	fmt.Println("=== 算法过程演示 ===")

	s := "leetcode"
	wordDict := []string{"leet", "code"}

	fmt.Printf("输入: s=\"%s\", wordDict=%v\n", s, wordDict)
	fmt.Println()

	// 演示动态规划过程
	fmt.Println("动态规划过程:")
	n := len(s)
	wordSet := make(map[string]bool)
	for _, word := range wordDict {
		wordSet[word] = true
	}

	dp := make([]bool, n+1)
	dp[0] = true

	fmt.Printf("初始状态: dp[0] = %t (空字符串)\n", dp[0])

	for i := 1; i <= n; i++ {
		fmt.Printf("\n检查位置 %d (字符串: \"%s\"):\n", i, s[:i])

		for j := 0; j < i; j++ {
			substr := s[j:i]
			inDict := wordSet[substr]
			fmt.Printf("  分割点 %d: dp[%d]=%t, \"%s\"在字典中=%t",
				j, j, dp[j], substr, inDict)

			if dp[j] && inDict {
				dp[i] = true
				fmt.Printf(" ✅ 找到拆分方式!")
				break
			} else {
				fmt.Printf(" ❌")
			}
			fmt.Println()
		}

		fmt.Printf("结果: dp[%d] = %t\n", i, dp[i])
	}

	fmt.Printf("\n最终结果: %t\n", dp[n])
}

// 展示拆分树
func showBreakTree() {
	fmt.Println()
	fmt.Println("=== 拆分决策树 ===")

	s := "applepenapple"
	wordDict := []string{"apple", "pen"}

	fmt.Printf("输入: \"%s\"\n", s)
	fmt.Printf("字典: %v\n", wordDict)
	fmt.Println()
	fmt.Println("拆分决策树:")
	fmt.Println("applepenapple")
	fmt.Println("├─ apple + penapple")
	fmt.Println("│  └─ pen + apple")
	fmt.Println("│     └─ apple (完成)")
	fmt.Println("└─ app (字典中没有)")
	fmt.Println()
	fmt.Println("成功路径: apple → pen → apple")

	solutions := findAllWordBreaks(s, wordDict)
	fmt.Printf("所有拆分方案: %v\n", solutions)
}

func main() {
	fmt.Println("139. 单词拆分 - 多种解法实现")
	fmt.Println("================================")

	// 基础功能测试
	testWordBreak()

	// 性能对比测试
	benchmarkWordBreak()

	// 算法过程演示
	demonstrateAlgorithm()

	// 拆分树展示
	showBreakTree()

	// 展示算法特点
	fmt.Println()
	fmt.Println("=== 算法特点分析 ===")
	fmt.Println("1. 动态规划：状态转移清晰，适合理解")
	fmt.Println("2. DFS+记忆化：递归思维，避免重复计算")
	fmt.Println("3. BFS：层次遍历，直观展示搜索过程")
	fmt.Println("4. Trie+DP：大字典时优化查找效率")
	fmt.Println("5. 优化DP：长度过滤，减少无效检查")

	fmt.Println()
	fmt.Println("=== 关键技巧总结 ===")
	fmt.Println("• 状态设计：dp[i]表示前i个字符是否可拆分")
	fmt.Println("• 转移方程：枚举所有分割点进行状态转移")
	fmt.Println("• 哈希优化：字典转HashMap提升查找效率")
	fmt.Println("• 记忆化：避免重复子问题的计算")
	fmt.Println("• 剪枝优化：长度过滤和早期终止")
	fmt.Println("• 数据结构：根据数据规模选择合适的结构")
}
