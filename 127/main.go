package main

import (
	"fmt"
	"strings"
	"time"
)

// 解法一：单向BFS（经典解法）
// 时间复杂度：O(N²×M)，空间复杂度：O(N×M)
func ladderLength(beginWord string, endWord string, wordList []string) int {
	// 特殊情况：起点等于终点
	if beginWord == endWord {
		return 1
	}

	// 将wordList转换为set以便快速查找
	wordSet := make(map[string]bool)
	for _, word := range wordList {
		wordSet[word] = true
	}

	// 如果endWord不在wordList中，无法到达
	if !wordSet[endWord] {
		return 0
	}

	// BFS队列和访问记录
	queue := []string{beginWord}
	visited := make(map[string]bool)
	visited[beginWord] = true
	level := 1

	for len(queue) > 0 {
		size := len(queue)

		// 处理当前层的所有单词
		for i := 0; i < size; i++ {
			current := queue[i]

			// 生成所有可能的邻居单词
			neighbors := getNeighbors(current, wordSet)
			for _, neighbor := range neighbors {
				if neighbor == endWord {
					return level + 1
				}

				if !visited[neighbor] {
					visited[neighbor] = true
					queue = append(queue, neighbor)
				}
			}
		}

		// 更新队列为下一层
		queue = queue[size:]
		level++
	}

	return 0
}

// 生成所有有效的邻居单词
func getNeighbors(word string, wordSet map[string]bool) []string {
	var neighbors []string
	chars := []rune(word)

	for i := 0; i < len(chars); i++ {
		original := chars[i]

		// 尝试26个字母
		for c := 'a'; c <= 'z'; c++ {
			if c == original {
				continue
			}

			chars[i] = c
			newWord := string(chars)

			if wordSet[newWord] {
				neighbors = append(neighbors, newWord)
			}
		}

		chars[i] = original // 恢复原字符
	}

	return neighbors
}

// 解法二：双向BFS（优化解法）
// 时间复杂度：O(N×M)，空间复杂度：O(N×M)
func ladderLengthBidirectional(beginWord string, endWord string, wordList []string) int {
	// 特殊情况：起点等于终点
	if beginWord == endWord {
		return 1
	}

	wordSet := make(map[string]bool)
	for _, word := range wordList {
		wordSet[word] = true
	}

	if !wordSet[endWord] {
		return 0
	}

	// 双向搜索集合
	beginSet := make(map[string]bool)
	endSet := make(map[string]bool)
	beginSet[beginWord] = true
	endSet[endWord] = true

	visited := make(map[string]bool)
	level := 1

	for len(beginSet) > 0 && len(endSet) > 0 {
		// 优化：始终扩展较小的集合
		if len(beginSet) > len(endSet) {
			beginSet, endSet = endSet, beginSet
		}

		nextSet := make(map[string]bool)

		for word := range beginSet {
			neighbors := getNeighbors(word, wordSet)

			for _, neighbor := range neighbors {
				// 如果在对方集合中找到，说明路径连通
				if endSet[neighbor] {
					return level + 1
				}

				// 如果未访问过，加入下一层
				if !visited[neighbor] {
					visited[neighbor] = true
					nextSet[neighbor] = true
				}
			}
		}

		beginSet = nextSet
		level++
	}

	return 0
}

// 解法三：使用模式匹配优化的BFS
// 时间复杂度：O(N×M)，空间复杂度：O(N×M)
func ladderLengthPattern(beginWord string, endWord string, wordList []string) int {
	if beginWord == endWord {
		return 1
	}

	// 构建模式到单词的映射
	patterns := make(map[string][]string)
	for _, word := range wordList {
		for i := 0; i < len(word); i++ {
			pattern := word[:i] + "*" + word[i+1:]
			patterns[pattern] = append(patterns[pattern], word)
		}
	}

	// 也要为beginWord建立模式
	for i := 0; i < len(beginWord); i++ {
		pattern := beginWord[:i] + "*" + beginWord[i+1:]
		patterns[pattern] = append(patterns[pattern], beginWord)
	}

	// BFS
	queue := []string{beginWord}
	visited := make(map[string]bool)
	visited[beginWord] = true
	level := 1

	for len(queue) > 0 {
		size := len(queue)

		for i := 0; i < size; i++ {
			current := queue[i]

			// 通过模式找到所有邻居
			for j := 0; j < len(current); j++ {
				pattern := current[:j] + "*" + current[j+1:]

				for _, neighbor := range patterns[pattern] {
					if neighbor == endWord {
						return level + 1
					}

					if !visited[neighbor] && neighbor != current {
						visited[neighbor] = true
						queue = append(queue, neighbor)
					}
				}
			}
		}

		queue = queue[size:]
		level++
	}

	return 0
}

// 解法四：A*搜索算法
// 时间复杂度：O(N×M×log N)，空间复杂度：O(N×M)
func ladderLengthAStar(beginWord string, endWord string, wordList []string) int {
	wordSet := make(map[string]bool)
	for _, word := range wordList {
		wordSet[word] = true
	}

	if !wordSet[endWord] {
		return 0
	}

	// 启发式函数：计算两个单词的差异字符数
	heuristic := func(word1, word2 string) int {
		diff := 0
		for i := 0; i < len(word1); i++ {
			if word1[i] != word2[i] {
				diff++
			}
		}
		return diff
	}

	// 优先队列节点
	type Node struct {
		word string
		g    int // 从起点到当前节点的实际距离
		f    int // g + h（启发式距离）
	}

	// 简化的优先队列实现
	var pq []Node

	// 添加起始节点
	start := Node{
		word: beginWord,
		g:    0,
		f:    heuristic(beginWord, endWord),
	}
	pq = append(pq, start)

	visited := make(map[string]int)
	visited[beginWord] = 0

	for len(pq) > 0 {
		// 简单的优先队列：找f值最小的节点
		minIdx := 0
		for i := 1; i < len(pq); i++ {
			if pq[i].f < pq[minIdx].f {
				minIdx = i
			}
		}

		current := pq[minIdx]
		pq = append(pq[:minIdx], pq[minIdx+1:]...)

		if current.word == endWord {
			return current.g + 1
		}

		neighbors := getNeighbors(current.word, wordSet)
		for _, neighbor := range neighbors {
			newG := current.g + 1

			if prevG, exists := visited[neighbor]; !exists || newG < prevG {
				visited[neighbor] = newG
				node := Node{
					word: neighbor,
					g:    newG,
					f:    newG + heuristic(neighbor, endWord),
				}
				pq = append(pq, node)
			}
		}
	}

	return 0
}

// 解法五：DFS + 记忆化（递归回溯）
// 时间复杂度：O(N!)，空间复杂度：O(N×M + 深度)
func ladderLengthDFS(beginWord string, endWord string, wordList []string) int {
	wordSet := make(map[string]bool)
	for _, word := range wordList {
		wordSet[word] = true
	}

	if !wordSet[endWord] {
		return 0
	}

	memo := make(map[string]int)
	visited := make(map[string]bool)

	var dfs func(word string) int
	dfs = func(word string) int {
		if word == endWord {
			return 1
		}

		if val, exists := memo[word]; exists {
			return val
		}

		visited[word] = true
		minLen := 0

		neighbors := getNeighbors(word, wordSet)
		for _, neighbor := range neighbors {
			if !visited[neighbor] {
				length := dfs(neighbor)
				if length > 0 {
					if minLen == 0 || length+1 < minLen {
						minLen = length + 1
					}
				}
			}
		}

		visited[word] = false
		memo[word] = minLen
		return minLen
	}

	return dfs(beginWord)
}

// 测试函数
func testLadderLength() {
	testCases := []struct {
		beginWord string
		endWord   string
		wordList  []string
		expected  int
		desc      string
	}{
		{
			"hit", "cog",
			[]string{"hot", "dot", "dog", "lot", "log", "cog"},
			5, "示例1：标准转换路径",
		},
		{
			"hit", "cog",
			[]string{"hot", "dot", "dog", "lot", "log"},
			0, "示例2：目标不在字典中",
		},
		{
			"a", "c",
			[]string{"a", "b", "c"},
			2, "简单单字母转换",
		},
		{
			"hot", "dog",
			[]string{"hot", "dog", "dot"},
			3, "两步转换",
		},
		{
			"hot", "dog",
			[]string{"hot", "dog"},
			0, "无中间路径",
		},
		{
			"leet", "code",
			[]string{"lest", "leet", "lose", "code", "lode", "robe", "lost"},
			6, "复杂路径",
		},
		{
			"hit", "hit",
			[]string{"hit"},
			1, "起点等于终点",
		},
		{
			"qa", "sq",
			[]string{"si", "go", "se", "cm", "so", "ph", "mt", "db", "mb", "sb", "kr", "ln", "tm", "le", "av", "sm", "ar", "ci", "ca", "br", "ti", "ba", "to", "ra", "fa", "yo", "ow", "sn", "ya", "cr", "po", "fe", "ho", "ma", "re", "or", "rn", "au", "ur", "rh", "sr", "tc", "lt", "lo", "as", "fr", "nb", "yb", "if", "pb", "ge", "th", "pm", "rb", "sh", "co", "ga", "li", "ha", "hz", "no", "bi", "di", "hi", "qa", "pi", "os", "uh", "wm", "an", "me", "mo", "na", "la", "st", "er", "sc", "ne", "mn", "mi", "am", "ex", "pt", "io", "be", "fm", "ta", "tb", "ni", "mr", "pa", "he", "lr", "sq", "ye"},
			5, "大字典测试",
		},
	}

	fmt.Println("=== 单词接龙测试 ===")
	fmt.Println()

	for i, tc := range testCases {
		// 测试主要解法
		result1 := ladderLength(tc.beginWord, tc.endWord, tc.wordList)
		result2 := ladderLengthBidirectional(tc.beginWord, tc.endWord, tc.wordList)
		result3 := ladderLengthPattern(tc.beginWord, tc.endWord, tc.wordList)

		status := "✅"
		if result1 != tc.expected {
			status = "❌"
		}

		fmt.Printf("测试 %d: %s\n", i+1, tc.desc)
		fmt.Printf("输入: beginWord=\"%s\", endWord=\"%s\"\n", tc.beginWord, tc.endWord)
		fmt.Printf("字典: %v\n", tc.wordList)
		fmt.Printf("期望: %d\n", tc.expected)
		fmt.Printf("单向BFS: %d\n", result1)
		fmt.Printf("双向BFS: %d\n", result2)
		fmt.Printf("模式匹配: %d\n", result3)
		fmt.Printf("结果: %s\n", status)
		fmt.Println(strings.Repeat("-", 50))
	}
}

// 性能测试
func benchmarkLadderLength() {
	fmt.Println()
	fmt.Println("=== 性能测试 ===")
	fmt.Println()

	// 构造测试数据
	testData := []struct {
		beginWord string
		endWord   string
		wordList  []string
		desc      string
	}{
		{
			"hit", "cog",
			[]string{"hot", "dot", "dog", "lot", "log", "cog"},
			"小字典测试",
		},
		{
			"qa", "sq",
			generateLargeWordList(100, 2),
			"中等字典测试",
		},
		{
			"start", "enddd",
			generateLargeWordList(500, 5),
			"大字典测试",
		},
	}

	algorithms := []struct {
		name string
		fn   func(string, string, []string) int
	}{
		{"单向BFS", ladderLength},
		{"双向BFS", ladderLengthBidirectional},
		{"模式匹配", ladderLengthPattern},
		{"A*搜索", ladderLengthAStar},
	}

	for _, data := range testData {
		fmt.Printf("%s (字典大小: %d):\n", data.desc, len(data.wordList))

		for _, algo := range algorithms {
			start := time.Now()
			result := algo.fn(data.beginWord, data.endWord, data.wordList)
			duration := time.Since(start)

			fmt.Printf("  %s: 结果=%d, 耗时: %v\n", algo.name, result, duration)
		}
		fmt.Println()
	}
}

// 生成大规模测试词典
func generateLargeWordList(size, wordLen int) []string {
	words := make([]string, 0, size)

	// 生成一些相关的单词
	base := strings.Repeat("a", wordLen)
	words = append(words, base)

	for i := 1; i < size; i++ {
		word := []rune(base)

		// 随机改变1-2个字符
		changes := 1 + i%2
		for j := 0; j < changes; j++ {
			pos := (i + j) % wordLen
			word[pos] = rune('a' + (i+j)%26)
		}

		words = append(words, string(word))
	}

	return words
}

// 演示BFS搜索过程
func demonstrateBFS() {
	fmt.Println("\n=== BFS搜索过程演示 ===")
	beginWord := "hit"
	endWord := "cog"
	wordList := []string{"hot", "dot", "dog", "lot", "log", "cog"}

	fmt.Printf("从 \"%s\" 到 \"%s\" 的转换过程:\n", beginWord, endWord)
	fmt.Printf("字典: %v\n\n", wordList)

	// 演示搜索层次
	wordSet := make(map[string]bool)
	for _, word := range wordList {
		wordSet[word] = true
	}

	queue := []string{beginWord}
	visited := make(map[string]bool)
	visited[beginWord] = true
	level := 1

	fmt.Printf("层次 %d: %v\n", level, queue)

	for len(queue) > 0 && level <= 5 {
		size := len(queue)
		nextLevel := []string{}

		for i := 0; i < size; i++ {
			current := queue[i]
			neighbors := getNeighbors(current, wordSet)

			for _, neighbor := range neighbors {
				if neighbor == endWord {
					fmt.Printf("层次 %d: 找到目标 %s！\n", level+1, neighbor)
					fmt.Printf("路径长度: %d\n", level+1)
					return
				}

				if !visited[neighbor] {
					visited[neighbor] = true
					nextLevel = append(nextLevel, neighbor)
				}
			}
		}

		if len(nextLevel) > 0 {
			queue = nextLevel
			level++
			fmt.Printf("层次 %d: %v\n", level, queue)
		} else {
			break
		}
	}
}

// 演示双向BFS
func demonstrateBidirectionalBFS() {
	fmt.Println("\n=== 双向BFS演示 ===")
	beginWord := "hit"
	endWord := "cog"
	wordList := []string{"hot", "dot", "dog", "lot", "log", "cog"}

	fmt.Printf("双向搜索从 \"%s\" 和 \"%s\" 同时开始\n", beginWord, endWord)

	wordSet := make(map[string]bool)
	for _, word := range wordList {
		wordSet[word] = true
	}

	beginSet := map[string]bool{beginWord: true}
	endSet := map[string]bool{endWord: true}
	visited := make(map[string]bool)
	level := 1

	fmt.Printf("层次 %d: 正向=%v, 反向=%v\n", level, getKeys(beginSet), getKeys(endSet))

	for len(beginSet) > 0 && len(endSet) > 0 && level <= 3 {
		if len(beginSet) > len(endSet) {
			beginSet, endSet = endSet, beginSet
			fmt.Println("交换搜索方向")
		}

		nextSet := make(map[string]bool)

		for word := range beginSet {
			neighbors := getNeighbors(word, wordSet)

			for _, neighbor := range neighbors {
				if endSet[neighbor] {
					fmt.Printf("层次 %d: 在 %s 处相遇！\n", level+1, neighbor)
					fmt.Printf("总路径长度: %d\n", level+1)
					return
				}

				if !visited[neighbor] {
					visited[neighbor] = true
					nextSet[neighbor] = true
				}
			}
		}

		beginSet = nextSet
		level++

		if len(beginSet) > 0 {
			fmt.Printf("层次 %d: 当前扩展=%v, 对方=%v\n", level, getKeys(beginSet), getKeys(endSet))
		}
	}
}

// 辅助函数：获取map的所有键
func getKeys(m map[string]bool) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

func main() {
	fmt.Println("127. 单词接龙 - 多种解法实现")
	fmt.Println("=====================================")

	// 基础功能测试
	testLadderLength()

	// 性能对比测试
	benchmarkLadderLength()

	// BFS过程演示
	demonstrateBFS()

	// 双向BFS演示
	demonstrateBidirectionalBFS()

	// 展示算法特点
	fmt.Println("\n=== 算法特点分析 ===")
	fmt.Println("1. 单向BFS：经典解法，层次遍历，简单直观")
	fmt.Println("2. 双向BFS：从两端搜索，搜索空间减半，性能最优")
	fmt.Println("3. 模式匹配：预处理优化邻居查找，减少重复计算")
	fmt.Println("4. A*搜索：启发式搜索，适合有明确目标的场景")
	fmt.Println("5. DFS回溯：递归实现，但时间复杂度较高")

	fmt.Println("\n=== 关键技巧总结 ===")
	fmt.Println("• 图的建模：将单词看作图中的节点，差一个字母的单词相连")
	fmt.Println("• BFS层次遍历：保证找到的第一个路径是最短的")
	fmt.Println("• 双向优化：从两端同时搜索，显著减少搜索空间")
	fmt.Println("• 访问控制：避免重复访问和环路问题")
	fmt.Println("• 集合操作：使用Set进行快速查找和去重")
}
