package main

import (
	"fmt"
	"sort"
	"strings"
	"time"
)

// ========== 方法1: 字典树方法（推荐） ==========

// TrieNode 字典树节点
type TrieNode struct {
	children [26]*TrieNode // 26个小写字母
	isRoot   bool          // 是否为词根结尾
}

// newTrieNode 创建新的字典树节点
func newTrieNode() *TrieNode {
	return &TrieNode{
		children: [26]*TrieNode{},
		isRoot:   false,
	}
}

// Trie 字典树
type Trie struct {
	root *TrieNode
}

// newTrie 创建新的字典树
func newTrie() *Trie {
	return &Trie{root: newTrieNode()}
}

// insert 插入词根到字典树
func (t *Trie) insert(root string) {
	node := t.root
	for i := 0; i < len(root); i++ {
		index := root[i] - 'a'
		if node.children[index] == nil {
			node.children[index] = newTrieNode()
		}
		node = node.children[index]
	}
	node.isRoot = true
}

// findRoot 查找单词的最短词根
func (t *Trie) findRoot(word string) string {
	node := t.root
	for i := 0; i < len(word); i++ {
		index := word[i] - 'a'
		if node.children[index] == nil {
			break
		}
		node = node.children[index]
		// 找到词根，立即返回（保证最短）
		if node.isRoot {
			return word[:i+1]
		}
	}
	return word // 没找到词根，返回原单词
}

func replaceWords1(dictionary []string, sentence string) string {
	// 构建字典树
	trie := newTrie()
	for _, root := range dictionary {
		trie.insert(root)
	}

	// 分割句子并替换
	words := strings.Split(sentence, " ")
	for i, word := range words {
		words[i] = trie.findRoot(word)
	}

	return strings.Join(words, " ")
}

// ========== 方法2: 哈希表方法 ==========
func replaceWords2(dictionary []string, sentence string) string {
	// 将词根存储到哈希表
	rootSet := make(map[string]bool)
	for _, root := range dictionary {
		rootSet[root] = true
	}

	words := strings.Split(sentence, " ")
	for i, word := range words {
		// 逐个检查前缀
		for j := 1; j <= len(word); j++ {
			prefix := word[:j]
			if rootSet[prefix] {
				words[i] = prefix
				break
			}
		}
	}

	return strings.Join(words, " ")
}

// ========== 方法3: 排序前缀匹配 ==========
func replaceWords3(dictionary []string, sentence string) string {
	// 按长度排序，确保最短的词根优先匹配
	sort.Slice(dictionary, func(i, j int) bool {
		return len(dictionary[i]) < len(dictionary[j])
	})

	words := strings.Split(sentence, " ")
	for i, word := range words {
		for _, root := range dictionary {
			if len(root) <= len(word) && strings.HasPrefix(word, root) {
				words[i] = root
				break // 找到最短的词根就停止
			}
		}
	}

	return strings.Join(words, " ")
}

// ========== 方法4: 优化字典树（提前终止） ==========

// TrieNode2 优化的字典树节点
type TrieNode2 struct {
	children map[byte]*TrieNode2
	isRoot   bool
	rootWord string // 直接存储词根
}

// newTrieNode2 创建优化的字典树节点
func newTrieNode2() *TrieNode2 {
	return &TrieNode2{
		children: make(map[byte]*TrieNode2),
		isRoot:   false,
		rootWord: "",
	}
}

// Trie2 优化的字典树
type Trie2 struct {
	root *TrieNode2
}

// newTrie2 创建优化的字典树
func newTrie2() *Trie2 {
	return &Trie2{root: newTrieNode2()}
}

// insert 插入词根（如果已存在更短的词根，不插入）
func (t *Trie2) insert(root string) {
	node := t.root
	for i := 0; i < len(root); i++ {
		char := root[i]

		// 如果当前路径上已经有词根，且更短，则不需要插入
		if node.isRoot {
			return
		}

		if node.children[char] == nil {
			node.children[char] = newTrieNode2()
		}
		node = node.children[char]
	}

	// 标记为词根，并清空其子节点（因为我们只需要最短的）
	node.isRoot = true
	node.rootWord = root
	node.children = make(map[byte]*TrieNode2)
}

// findRoot 查找最短词根
func (t *Trie2) findRoot(word string) string {
	node := t.root
	for i := 0; i < len(word); i++ {
		char := word[i]
		if node.children[char] == nil {
			break
		}
		node = node.children[char]
		if node.isRoot {
			return node.rootWord
		}
	}
	return word
}

func replaceWords4(dictionary []string, sentence string) string {
	// 按长度排序，优先插入短的词根
	sort.Slice(dictionary, func(i, j int) bool {
		return len(dictionary[i]) < len(dictionary[j])
	})

	trie := newTrie2()
	for _, root := range dictionary {
		trie.insert(root)
	}

	words := strings.Split(sentence, " ")
	for i, word := range words {
		words[i] = trie.findRoot(word)
	}

	return strings.Join(words, " ")
}

// ========== 方法5: 字符串暴力匹配 ==========
func replaceWords5(dictionary []string, sentence string) string {
	words := strings.Split(sentence, " ")

	for i, word := range words {
		minRoot := word
		for _, root := range dictionary {
			if len(root) < len(minRoot) && strings.HasPrefix(word, root) {
				minRoot = root
			}
		}
		words[i] = minRoot
	}

	return strings.Join(words, " ")
}

// ========== 工具函数 ==========

// 打印字典树结构
func (t *Trie) printTrie() {
	fmt.Println("字典树结构:")
	t.printTrieHelper(t.root, "", "")
}

func (t *Trie) printTrieHelper(node *TrieNode, prefix, indent string) {
	if node.isRoot {
		fmt.Printf("%s%s [词根]\n", indent, prefix)
	}

	for i, child := range node.children {
		if child != nil {
			char := byte('a' + i)
			t.printTrieHelper(child, prefix+string(char), indent+"  ")
		}
	}
}

// 统计字典树节点数
func (t *Trie) countNodes() int {
	return t.countNodesHelper(t.root)
}

func (t *Trie) countNodesHelper(node *TrieNode) int {
	count := 1
	for _, child := range node.children {
		if child != nil {
			count += t.countNodesHelper(child)
		}
	}
	return count
}

// 获取所有词根
func (t *Trie) getAllRoots() []string {
	var roots []string
	t.getAllRootsHelper(t.root, "", &roots)
	return roots
}

func (t *Trie) getAllRootsHelper(node *TrieNode, prefix string, roots *[]string) {
	if node.isRoot {
		*roots = append(*roots, prefix)
	}

	for i, child := range node.children {
		if child != nil {
			char := byte('a' + i)
			t.getAllRootsHelper(child, prefix+string(char), roots)
		}
	}
}

// 生成测试数据
func generateTestCase(rootCount, sentenceLen int) ([]string, string) {
	// 生成词根
	roots := make([]string, rootCount)
	chars := "abcdefghijklmnopqrstuvwxyz"

	for i := 0; i < rootCount; i++ {
		length := 2 + i%4 // 长度2-5
		root := make([]byte, length)
		for j := 0; j < length; j++ {
			root[j] = chars[(i*3+j*7)%26]
		}
		roots[i] = string(root)
	}

	// 生成句子
	words := make([]string, sentenceLen)
	for i := 0; i < sentenceLen; i++ {
		if i%3 == 0 && i/3 < len(roots) {
			// 部分单词使用词根+后缀
			root := roots[i/3]
			suffix := "ing"
			words[i] = root + suffix
		} else {
			// 随机单词
			length := 3 + i%5
			word := make([]byte, length)
			for j := 0; j < length; j++ {
				word[j] = chars[(i*5+j*11)%26]
			}
			words[i] = string(word)
		}
	}

	return roots, strings.Join(words, " ")
}

// ========== 测试和性能评估 ==========
func main() {
	// 测试用例
	testCases := []struct {
		name       string
		dictionary []string
		sentence   string
		expected   string
	}{
		{
			name:       "示例1: 基础替换",
			dictionary: []string{"cat", "bat", "rat"},
			sentence:   "the cattle was rattled by the battery",
			expected:   "the cat was rat by the bat",
		},
		{
			name:       "示例2: 单字符词根",
			dictionary: []string{"a", "b", "c"},
			sentence:   "aadsfasf absbs bbab cadsfafs",
			expected:   "a a b c",
		},
		{
			name:       "测试3: 无匹配",
			dictionary: []string{"cat", "bat", "rat"},
			sentence:   "hello world programming",
			expected:   "hello world programming",
		},
		{
			name:       "测试4: 完全匹配",
			dictionary: []string{"hello", "world"},
			sentence:   "hello world",
			expected:   "hello world",
		},
		{
			name:       "测试5: 多个词根匹配",
			dictionary: []string{"a", "aa", "aaa"},
			sentence:   "aaaaaaaaa",
			expected:   "a",
		},
		{
			name:       "测试6: 长词根",
			dictionary: []string{"application", "app", "apple"},
			sentence:   "application development",
			expected:   "app development",
		},
		{
			name:       "测试7: 空词根",
			dictionary: []string{},
			sentence:   "hello world",
			expected:   "hello world",
		},
		{
			name:       "测试8: 重复词根",
			dictionary: []string{"cat", "cat", "dog"},
			sentence:   "the caterpillar and doggy",
			expected:   "the cat and dog",
		},
		{
			name:       "测试9: 前缀包含",
			dictionary: []string{"car", "card", "care"},
			sentence:   "careful careless cards",
			expected:   "car car car", // car是最短的词根，会被优先匹配
		},
		{
			name:       "测试10: 单词词根",
			dictionary: []string{"i", "love", "leetcode"},
			sentence:   "i love solving leetcode problems",
			expected:   "i love solving leetcode problems",
		},
	}

	// 算法方法
	methods := []struct {
		name string
		fn   func([]string, string) string
	}{
		{"字典树方法", replaceWords1},
		{"哈希表方法", replaceWords2},
		{"排序前缀匹配", replaceWords3},
		{"优化字典树", replaceWords4},
		{"暴力字符串匹配", replaceWords5},
	}

	fmt.Println("=== LeetCode 648. 单词替换 - 测试结果 ===")
	fmt.Println()

	// 运行测试
	for _, tc := range testCases {
		fmt.Printf("测试用例: %s\n", tc.name)
		fmt.Printf("词根字典: %v\n", tc.dictionary)
		fmt.Printf("原句子: %s\n", tc.sentence)

		allPassed := true
		var results []string
		var times []time.Duration

		for _, method := range methods {
			start := time.Now()
			result := method.fn(tc.dictionary, tc.sentence)
			elapsed := time.Since(start)

			results = append(results, result)
			times = append(times, elapsed)

			status := "✅"
			if result != tc.expected {
				status = "❌"
				allPassed = false
			}

			fmt.Printf("  %s: %s (耗时: %v)\n", method.name, status, elapsed)
			if result != tc.expected {
				fmt.Printf("    预期: %s\n", tc.expected)
				fmt.Printf("    实际: %s\n", result)
			}
		}

		fmt.Printf("期望结果: %s\n", tc.expected)

		if allPassed {
			fmt.Println("✅ 所有方法均通过")
		} else {
			fmt.Println("❌ 存在失败的方法")
		}
		fmt.Println(strings.Repeat("-", 60))
	}

	// 字典树演示
	fmt.Println("\n=== 字典树结构演示 ===")
	demoTrie()

	// 性能对比测试
	fmt.Println("\n=== 性能对比测试 ===")
	performanceTest()

	// 算法特性总结
	fmt.Println("\n=== 算法特性总结 ===")
	fmt.Println("1. 字典树方法:")
	fmt.Println("   - 时间复杂度: O(N+M)")
	fmt.Println("   - 空间复杂度: O(L)")
	fmt.Println("   - 特点: 最优解法，前缀搜索高效")
	fmt.Println()
	fmt.Println("2. 哈希表方法:")
	fmt.Println("   - 时间复杂度: O(N×K)")
	fmt.Println("   - 空间复杂度: O(M×K)")
	fmt.Println("   - 特点: 实现简单，理解容易")
	fmt.Println()
	fmt.Println("3. 排序前缀匹配:")
	fmt.Println("   - 时间复杂度: O(M×logM+N×M×K)")
	fmt.Println("   - 空间复杂度: O(M×K)")
	fmt.Println("   - 特点: 排序保证最短优先")
	fmt.Println()
	fmt.Println("4. 优化字典树:")
	fmt.Println("   - 时间复杂度: O(N+M)")
	fmt.Println("   - 空间复杂度: O(L)")
	fmt.Println("   - 特点: 提前终止，空间最优")
	fmt.Println()
	fmt.Println("5. 暴力字符串匹配:")
	fmt.Println("   - 时间复杂度: O(N×M×K)")
	fmt.Println("   - 空间复杂度: O(1)")
	fmt.Println("   - 特点: 最直接但效率低")

	// 单词替换演示
	fmt.Println("\n=== 单词替换演示 ===")
	demoWordReplacement()
}

// 字典树演示
func demoTrie() {
	fmt.Println("构建词根字典树: [cat, bat, rat]")
	dictionary := []string{"cat", "bat", "rat"}
	trie := newTrie()

	for _, root := range dictionary {
		trie.insert(root)
		fmt.Printf("插入词根: %s\n", root)
	}

	fmt.Printf("\n字典树节点总数: %d\n", trie.countNodes())
	fmt.Printf("存储的词根: %v\n", trie.getAllRoots())

	// 测试单词查找
	fmt.Println("\n单词查找测试:")
	testWords := []string{"cattle", "battery", "rattled", "hello"}
	for _, word := range testWords {
		root := trie.findRoot(word)
		if root != word {
			fmt.Printf("'%s' → '%s' (找到词根)\n", word, root)
		} else {
			fmt.Printf("'%s' → '%s' (无词根)\n", word, root)
		}
	}
}

// 性能测试
func performanceTest() {
	testSizes := []struct {
		roots       int
		sentenceLen int
	}{
		{10, 20},
		{50, 100},
		{100, 200},
		{200, 500},
	}

	methods := []struct {
		name string
		fn   func([]string, string) string
	}{
		{"字典树", replaceWords1},
		{"哈希表", replaceWords2},
		{"排序匹配", replaceWords3},
		{"优化字典树", replaceWords4},
	}

	for _, size := range testSizes {
		fmt.Printf("性能测试 - 词根数: %d, 句子长度: %d\n",
			size.roots, size.sentenceLen)

		dictionary, sentence := generateTestCase(size.roots, size.sentenceLen)

		for _, method := range methods {
			start := time.Now()
			result := method.fn(dictionary, sentence)
			elapsed := time.Since(start)

			wordCount := len(strings.Split(result, " "))
			fmt.Printf("  %s: 耗时=%v, 处理单词=%d\n",
				method.name, elapsed, wordCount)
		}
		fmt.Println()
	}
}

// 单词替换演示
func demoWordReplacement() {
	fmt.Println("单词替换场景演示:")

	// 场景1: 动物相关
	fmt.Println("\n场景1: 动物词汇替换")
	dictionary1 := []string{"cat", "dog", "bird"}
	sentence1 := "the cats and dogs are flying like birds"
	result1 := replaceWords1(dictionary1, sentence1)
	fmt.Printf("词根: %v\n", dictionary1)
	fmt.Printf("原句: %s\n", sentence1)
	fmt.Printf("替换: %s\n", result1)

	// 场景2: 技术词汇
	fmt.Println("\n场景2: 技术词汇替换")
	dictionary2 := []string{"program", "develop", "test", "debug"}
	sentence2 := "programming development testing debugging"
	result2 := replaceWords1(dictionary2, sentence2)
	fmt.Printf("词根: %v\n", dictionary2)
	fmt.Printf("原句: %s\n", sentence2)
	fmt.Printf("替换: %s\n", result2)

	// 场景3: 多层次词根
	fmt.Println("\n场景3: 多层次词根替换")
	dictionary3 := []string{"a", "ap", "app", "appl", "apple"}
	sentence3 := "application appreciate apple"
	result3 := replaceWords1(dictionary3, sentence3)
	fmt.Printf("词根: %v\n", dictionary3)
	fmt.Printf("原句: %s\n", sentence3)
	fmt.Printf("替换: %s\n", result3)

	fmt.Println("\n单词替换演示完成!")
}
