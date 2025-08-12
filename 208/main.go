package main

import (
	"fmt"
	"strings"
	"time"
)

// ========== 方法1: 标准字典树实现（推荐） ==========

// TrieNode 字典树节点
type TrieNode struct {
	children map[byte]*TrieNode // 字符到子节点的映射
	isEnd    bool               // 标记是否为单词结尾
}

// newTrieNode 创建新的字典树节点
func newTrieNode() *TrieNode {
	return &TrieNode{
		children: make(map[byte]*TrieNode),
		isEnd:    false,
	}
}

// Trie 字典树结构
type Trie struct {
	root *TrieNode
}

// Constructor 初始化字典树
func Constructor() Trie {
	return Trie{root: newTrieNode()}
}

// Insert 向字典树中插入单词
func (t *Trie) Insert(word string) {
	node := t.root
	for i := 0; i < len(word); i++ {
		char := word[i]
		if node.children[char] == nil {
			node.children[char] = newTrieNode()
		}
		node = node.children[char]
	}
	node.isEnd = true
}

// Search 搜索单词是否在字典树中
func (t *Trie) Search(word string) bool {
	node := t.searchPrefix(word)
	return node != nil && node.isEnd
}

// StartsWith 检查是否有以给定前缀开头的单词
func (t *Trie) StartsWith(prefix string) bool {
	return t.searchPrefix(prefix) != nil
}

// searchPrefix 搜索前缀，返回最后一个节点
func (t *Trie) searchPrefix(prefix string) *TrieNode {
	node := t.root
	for i := 0; i < len(prefix); i++ {
		char := prefix[i]
		if node.children[char] == nil {
			return nil
		}
		node = node.children[char]
	}
	return node
}

// ========== 方法2: 数组优化实现 ==========

// TrieNode2 使用数组优化的字典树节点
type TrieNode2 struct {
	children [26]*TrieNode2 // 固定26个小写字母
	isEnd    bool
}

// newTrieNode2 创建新的优化字典树节点
func newTrieNode2() *TrieNode2 {
	return &TrieNode2{
		children: [26]*TrieNode2{},
		isEnd:    false,
	}
}

// Trie2 数组优化字典树
type Trie2 struct {
	root *TrieNode2
}

// Constructor2 初始化数组优化字典树
func Constructor2() Trie2 {
	return Trie2{root: newTrieNode2()}
}

// Insert 向字典树中插入单词
func (t *Trie2) Insert(word string) {
	node := t.root
	for i := 0; i < len(word); i++ {
		index := word[i] - 'a'
		if node.children[index] == nil {
			node.children[index] = newTrieNode2()
		}
		node = node.children[index]
	}
	node.isEnd = true
}

// Search 搜索单词是否在字典树中
func (t *Trie2) Search(word string) bool {
	node := t.searchPrefix2(word)
	return node != nil && node.isEnd
}

// StartsWith 检查是否有以给定前缀开头的单词
func (t *Trie2) StartsWith(prefix string) bool {
	return t.searchPrefix2(prefix) != nil
}

// searchPrefix2 搜索前缀，返回最后一个节点
func (t *Trie2) searchPrefix2(prefix string) *TrieNode2 {
	node := t.root
	for i := 0; i < len(prefix); i++ {
		index := prefix[i] - 'a'
		if node.children[index] == nil {
			return nil
		}
		node = node.children[index]
	}
	return node
}

// ========== 方法3: 压缩字典树实现 ==========

// TrieNode3 压缩字典树节点
type TrieNode3 struct {
	children map[string]*TrieNode3 // 压缩路径存储
	isEnd    bool
	value    string // 存储压缩的字符串
}

// newTrieNode3 创建新的压缩字典树节点
func newTrieNode3() *TrieNode3 {
	return &TrieNode3{
		children: make(map[string]*TrieNode3),
		isEnd:    false,
		value:    "",
	}
}

// Trie3 压缩字典树
type Trie3 struct {
	root *TrieNode3
}

// Constructor3 初始化压缩字典树
func Constructor3() Trie3 {
	return Trie3{root: newTrieNode3()}
}

// Insert 向压缩字典树中插入单词
func (t *Trie3) Insert(word string) {
	node := t.root
	i := 0

	for i < len(word) {
		found := false
		for key, child := range node.children {
			if i < len(word) && len(key) > 0 && word[i] == key[0] {
				// 找到匹配的前缀
				commonLen := 0
				for commonLen < len(key) && i+commonLen < len(word) &&
					key[commonLen] == word[i+commonLen] {
					commonLen++
				}

				if commonLen == len(key) {
					// 完全匹配现有边
					node = child
					i += commonLen
					found = true
					break
				} else {
					// 部分匹配，需要分裂节点
					// 创建新的中间节点
					newNode := newTrieNode3()
					newNode.children[key[commonLen:]] = child

					// 更新当前边
					node.children[key[:commonLen]] = newNode
					delete(node.children, key)

					node = newNode
					i += commonLen
					found = true
					break
				}
			}
		}

		if !found {
			// 没有找到匹配的边，创建新边
			newNode := newTrieNode3()
			node.children[word[i:]] = newNode
			node = newNode
			break
		}
	}

	node.isEnd = true
}

// Search 搜索单词是否在压缩字典树中
func (t *Trie3) Search(word string) bool {
	node := t.searchPrefix3(word)
	return node != nil && node.isEnd
}

// StartsWith 检查是否有以给定前缀开头的单词
func (t *Trie3) StartsWith(prefix string) bool {
	return t.searchPrefix3(prefix) != nil
}

// searchPrefix3 搜索前缀
func (t *Trie3) searchPrefix3(prefix string) *TrieNode3 {
	node := t.root
	i := 0

	for i < len(prefix) {
		found := false
		for key, child := range node.children {
			if i < len(prefix) && len(key) > 0 && prefix[i] == key[0] {
				// 检查匹配长度
				matchLen := 0
				for matchLen < len(key) && i+matchLen < len(prefix) &&
					key[matchLen] == prefix[i+matchLen] {
					matchLen++
				}

				if i+matchLen == len(prefix) {
					// 前缀完全匹配
					return child
				} else if matchLen == len(key) {
					// 继续往下搜索
					node = child
					i += matchLen
					found = true
					break
				} else {
					// 部分匹配但前缀更长，不存在
					return nil
				}
			}
		}

		if !found {
			return nil
		}
	}

	return node
}

// ========== 方法4: 带统计功能的字典树 ==========

// TrieNode4 带统计的字典树节点
type TrieNode4 struct {
	children  map[byte]*TrieNode4
	isEnd     bool
	count     int // 记录经过此节点的单词数量
	wordCount int // 记录以此节点结尾的单词数量
}

// newTrieNode4 创建新的统计字典树节点
func newTrieNode4() *TrieNode4 {
	return &TrieNode4{
		children:  make(map[byte]*TrieNode4),
		isEnd:     false,
		count:     0,
		wordCount: 0,
	}
}

// Trie4 带统计功能的字典树
type Trie4 struct {
	root *TrieNode4
}

// Constructor4 初始化统计字典树
func Constructor4() Trie4 {
	return Trie4{root: newTrieNode4()}
}

// Insert 向字典树中插入单词
func (t *Trie4) Insert(word string) {
	node := t.root
	for i := 0; i < len(word); i++ {
		char := word[i]
		if node.children[char] == nil {
			node.children[char] = newTrieNode4()
		}
		node = node.children[char]
		node.count++
	}
	if !node.isEnd {
		node.isEnd = true
		node.wordCount++
	}
}

// Search 搜索单词是否在字典树中
func (t *Trie4) Search(word string) bool {
	node := t.searchPrefix4(word)
	return node != nil && node.isEnd
}

// StartsWith 检查是否有以给定前缀开头的单词
func (t *Trie4) StartsWith(prefix string) bool {
	return t.searchPrefix4(prefix) != nil
}

// CountWordsStartingWith 统计以给定前缀开头的单词数量
func (t *Trie4) CountWordsStartingWith(prefix string) int {
	node := t.searchPrefix4(prefix)
	if node == nil {
		return 0
	}
	return node.count
}

// searchPrefix4 搜索前缀
func (t *Trie4) searchPrefix4(prefix string) *TrieNode4 {
	node := t.root
	for i := 0; i < len(prefix); i++ {
		char := prefix[i]
		if node.children[char] == nil {
			return nil
		}
		node = node.children[char]
	}
	return node
}

// ========== 工具函数 ==========

// 打印字典树结构
func (t *Trie) PrintTrie() {
	fmt.Println("字典树结构:")
	t.printTrieHelper(t.root, "", "")
}

func (t *Trie) printTrieHelper(node *TrieNode, prefix, indent string) {
	if node.isEnd {
		fmt.Printf("%s%s [单词结尾]\n", indent, prefix)
	}

	for char, child := range node.children {
		fmt.Printf("%s%s%c\n", indent, prefix, char)
		t.printTrieHelper(child, prefix+string(char), indent+"  ")
	}
}

// 获取字典树中所有单词
func (t *Trie) GetAllWords() []string {
	var words []string
	t.getAllWordsHelper(t.root, "", &words)
	return words
}

func (t *Trie) getAllWordsHelper(node *TrieNode, prefix string, words *[]string) {
	if node.isEnd {
		*words = append(*words, prefix)
	}

	for char, child := range node.children {
		t.getAllWordsHelper(child, prefix+string(char), words)
	}
}

// 统计字典树节点数量
func (t *Trie) CountNodes() int {
	return t.countNodesHelper(t.root)
}

func (t *Trie) countNodesHelper(node *TrieNode) int {
	count := 1
	for _, child := range node.children {
		count += t.countNodesHelper(child)
	}
	return count
}

// ========== 测试和性能评估 ==========
func main() {
	// 测试用例
	testCases := []struct {
		name       string
		operations []string
		params     [][]string
		expected   []interface{}
	}{
		{
			name:       "示例1: 基础操作",
			operations: []string{"Trie", "insert", "search", "search", "startsWith", "insert", "search"},
			params:     [][]string{{}, {"apple"}, {"apple"}, {"app"}, {"app"}, {"app"}, {"app"}},
			expected:   []interface{}{nil, nil, true, false, true, nil, true},
		},
		{
			name:       "测试2: 空字符串",
			operations: []string{"Trie", "insert", "search", "startsWith"},
			params:     [][]string{{}, {""}, {""}, {""}},
			expected:   []interface{}{nil, nil, true, true},
		},
		{
			name:       "测试3: 重复插入",
			operations: []string{"Trie", "insert", "insert", "search"},
			params:     [][]string{{}, {"hello"}, {"hello"}, {"hello"}},
			expected:   []interface{}{nil, nil, nil, true},
		},
		{
			name:       "测试4: 不存在的单词",
			operations: []string{"Trie", "insert", "search", "search"},
			params:     [][]string{{}, {"world"}, {"word"}, {"worlds"}},
			expected:   []interface{}{nil, nil, false, false},
		},
		{
			name:       "测试5: 共享前缀",
			operations: []string{"Trie", "insert", "insert", "insert", "search", "search", "search", "startsWith"},
			params:     [][]string{{}, {"car"}, {"card"}, {"care"}, {"car"}, {"card"}, {"care"}, {"car"}},
			expected:   []interface{}{nil, nil, nil, nil, true, true, true, true},
		},
		{
			name:       "测试6: 长单词",
			operations: []string{"Trie", "insert", "search", "startsWith"},
			params:     [][]string{{}, {"supercalifragilisticexpialidocious"}, {"supercalifragilisticexpialidocious"}, {"super"}},
			expected:   []interface{}{nil, nil, true, true},
		},
		{
			name:       "测试7: 单字符",
			operations: []string{"Trie", "insert", "search", "startsWith"},
			params:     [][]string{{}, {"a"}, {"a"}, {"a"}},
			expected:   []interface{}{nil, nil, true, true},
		},
		{
			name:       "测试8: 前缀不匹配",
			operations: []string{"Trie", "insert", "startsWith", "startsWith"},
			params:     [][]string{{}, {"apple"}, {"orange"}, {"app"}},
			expected:   []interface{}{nil, nil, false, true},
		},
	}

	// 算法方法
	methods := []struct {
		name       string
		construct  func() interface{}
		insert     func(interface{}, string)
		search     func(interface{}, string) bool
		startsWith func(interface{}, string) bool
	}{
		{
			name: "标准字典树",
			construct: func() interface{} {
				trie := Constructor()
				return &trie
			},
			insert: func(t interface{}, word string) {
				t.(*Trie).Insert(word)
			},
			search: func(t interface{}, word string) bool {
				return t.(*Trie).Search(word)
			},
			startsWith: func(t interface{}, prefix string) bool {
				return t.(*Trie).StartsWith(prefix)
			},
		},
		{
			name: "数组优化",
			construct: func() interface{} {
				trie := Constructor2()
				return &trie
			},
			insert: func(t interface{}, word string) {
				t.(*Trie2).Insert(word)
			},
			search: func(t interface{}, word string) bool {
				return t.(*Trie2).Search(word)
			},
			startsWith: func(t interface{}, prefix string) bool {
				return t.(*Trie2).StartsWith(prefix)
			},
		},
		{
			name: "统计字典树",
			construct: func() interface{} {
				trie := Constructor4()
				return &trie
			},
			insert: func(t interface{}, word string) {
				t.(*Trie4).Insert(word)
			},
			search: func(t interface{}, word string) bool {
				return t.(*Trie4).Search(word)
			},
			startsWith: func(t interface{}, prefix string) bool {
				return t.(*Trie4).StartsWith(prefix)
			},
		},
	}

	fmt.Println("=== LeetCode 208. 实现 Trie (前缀树) - 测试结果 ===")
	fmt.Println()

	// 运行测试
	for _, tc := range testCases {
		fmt.Printf("测试用例: %s\n", tc.name)
		fmt.Printf("操作序列: %v\n", tc.operations)

		allPassed := true

		for _, method := range methods {
			fmt.Printf("  %s: ", method.name)

			var trie interface{}
			passed := true
			resultIndex := 0

			start := time.Now()

			for i, op := range tc.operations {
				switch op {
				case "Trie":
					trie = method.construct()
					resultIndex++
				case "insert":
					method.insert(trie, tc.params[i][0])
					resultIndex++
				case "search":
					result := method.search(trie, tc.params[i][0])
					expected := tc.expected[resultIndex].(bool)
					if result != expected {
						passed = false
						fmt.Printf("❌ search('%s') = %v, expected %v", tc.params[i][0], result, expected)
						break
					}
					resultIndex++
				case "startsWith":
					result := method.startsWith(trie, tc.params[i][0])
					expected := tc.expected[resultIndex].(bool)
					if result != expected {
						passed = false
						fmt.Printf("❌ startsWith('%s') = %v, expected %v", tc.params[i][0], result, expected)
						break
					}
					resultIndex++
				}
			}

			elapsed := time.Since(start)

			if passed {
				fmt.Printf("✅ (耗时: %v)\n", elapsed)
			} else {
				fmt.Printf(" (耗时: %v)\n", elapsed)
				allPassed = false
			}
		}

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
	fmt.Println("1. 标准字典树:")
	fmt.Println("   - 时间复杂度: O(m) per operation")
	fmt.Println("   - 空间复杂度: O(n*m)")
	fmt.Println("   - 特点: 通用性强，支持任意字符集")
	fmt.Println()
	fmt.Println("2. 数组优化:")
	fmt.Println("   - 时间复杂度: O(m) per operation")
	fmt.Println("   - 空间复杂度: O(n*m)")
	fmt.Println("   - 特点: 固定字符集下性能最优")
	fmt.Println()
	fmt.Println("3. 压缩字典树:")
	fmt.Println("   - 时间复杂度: O(m) per operation")
	fmt.Println("   - 空间复杂度: O(n*m)")
	fmt.Println("   - 特点: 空间效率最高，适合长单词")
	fmt.Println()
	fmt.Println("4. 统计字典树:")
	fmt.Println("   - 时间复杂度: O(m) per operation")
	fmt.Println("   - 空间复杂度: O(n*m)")
	fmt.Println("   - 特点: 支持统计功能，功能最全")

	// 应用场景演示
	fmt.Println("\n=== 应用场景演示 ===")
	applicationDemo()
}

// 字典树演示
func demoTrie() {
	fmt.Println("构建字典树，插入单词: [apple, app, application, apply]")
	trie := Constructor()

	words := []string{"apple", "app", "application", "apply"}
	for _, word := range words {
		trie.Insert(word)
		fmt.Printf("插入: %s\n", word)
	}

	fmt.Println("\n字典树中的所有单词:")
	allWords := trie.GetAllWords()
	for _, word := range allWords {
		fmt.Printf("  %s\n", word)
	}

	fmt.Printf("\n字典树节点总数: %d\n", trie.CountNodes())

	// 测试搜索
	fmt.Println("\n搜索测试:")
	testWords := []string{"app", "apple", "application", "orange"}
	for _, word := range testWords {
		result := trie.Search(word)
		fmt.Printf("搜索 '%s': %v\n", word, result)
	}

	// 测试前缀匹配
	fmt.Println("\n前缀匹配测试:")
	prefixes := []string{"app", "appl", "orange", "a"}
	for _, prefix := range prefixes {
		result := trie.StartsWith(prefix)
		fmt.Printf("前缀 '%s': %v\n", prefix, result)
	}
}

// 性能测试
func performanceTest() {
	wordSizes := []int{100, 500, 1000, 5000}
	methods := []struct {
		name      string
		construct func() interface{}
		insert    func(interface{}, string)
		search    func(interface{}, string) bool
	}{
		{
			name: "标准字典树",
			construct: func() interface{} {
				trie := Constructor()
				return &trie
			},
			insert: func(t interface{}, word string) {
				t.(*Trie).Insert(word)
			},
			search: func(t interface{}, word string) bool {
				return t.(*Trie).Search(word)
			},
		},
		{
			name: "数组优化",
			construct: func() interface{} {
				trie := Constructor2()
				return &trie
			},
			insert: func(t interface{}, word string) {
				t.(*Trie2).Insert(word)
			},
			search: func(t interface{}, word string) bool {
				return t.(*Trie2).Search(word)
			},
		},
		{
			name: "统计字典树",
			construct: func() interface{} {
				trie := Constructor4()
				return &trie
			},
			insert: func(t interface{}, word string) {
				t.(*Trie4).Insert(word)
			},
			search: func(t interface{}, word string) bool {
				return t.(*Trie4).Search(word)
			},
		},
	}

	for _, size := range wordSizes {
		fmt.Printf("性能测试 - 单词数量: %d\n", size)

		// 生成测试单词
		words := generateTestWords(size)

		for _, method := range methods {
			trie := method.construct()

			// 测试插入性能
			start := time.Now()
			for _, word := range words {
				method.insert(trie, word)
			}
			insertTime := time.Since(start)

			// 测试搜索性能
			start = time.Now()
			found := 0
			for _, word := range words {
				if method.search(trie, word) {
					found++
				}
			}
			searchTime := time.Since(start)

			fmt.Printf("  %s: 插入=%v, 搜索=%v, 找到=%d/%d\n",
				method.name, insertTime, searchTime, found, len(words))
		}
		fmt.Println()
	}
}

// 生成测试单词
func generateTestWords(count int) []string {
	words := make([]string, count)
	chars := "abcdefghijklmnopqrstuvwxyz"

	for i := 0; i < count; i++ {
		length := 3 + i%8 // 长度3-10
		word := make([]byte, length)
		for j := 0; j < length; j++ {
			word[j] = chars[(i*7+j*3)%26] // 伪随机字符
		}
		words[i] = string(word)
	}

	return words
}

// 应用场景演示
func applicationDemo() {
	fmt.Println("应用场景1: 自动补全")
	trie := Constructor()

	// 构建词典
	dictionary := []string{
		"apple", "application", "apply", "appreciate", "approach",
		"banana", "band", "bank", "basic", "battle",
		"cat", "car", "card", "care", "career",
	}

	for _, word := range dictionary {
		trie.Insert(word)
	}

	// 自动补全示例
	prefix := "app"
	fmt.Printf("输入前缀 '%s'，自动补全建议:\n", prefix)

	if trie.StartsWith(prefix) {
		allWords := trie.GetAllWords()
		suggestions := []string{}
		for _, word := range allWords {
			if len(word) >= len(prefix) && word[:len(prefix)] == prefix {
				suggestions = append(suggestions, word)
			}
		}

		for _, suggestion := range suggestions {
			fmt.Printf("  %s\n", suggestion)
		}
	} else {
		fmt.Println("  无匹配建议")
	}

	fmt.Println("\n应用场景2: 拼写检查")
	testWords := []string{"apple", "aple", "application", "aplicaton"}
	for _, word := range testWords {
		if trie.Search(word) {
			fmt.Printf("'%s': ✅ 拼写正确\n", word)
		} else {
			fmt.Printf("'%s': ❌ 拼写错误\n", word)
		}
	}

	fmt.Println("\n字典树应用完成!")
}
