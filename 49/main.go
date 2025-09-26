package main

import (
	"fmt"
	"sort"
	"strings"
	"time"
)

// 方法一：排序分组算法
// 最直观的解法，对每个字符串排序后作为键
func groupAnagrams1(strs []string) [][]string {
	groups := make(map[string][]string)

	for _, str := range strs {
		// 排序字符串
		sorted := sortString(str)
		groups[sorted] = append(groups[sorted], str)
	}

	// 转换为结果格式
	result := make([][]string, 0, len(groups))
	for _, group := range groups {
		result = append(result, group)
	}

	return result
}

// 排序字符串的辅助函数
func sortString(s string) string {
	runes := []rune(s)
	sort.Slice(runes, func(i, j int) bool {
		return runes[i] < runes[j]
	})
	return string(runes)
}

// 方法二：字符计数算法
// 效率最高的解法，统计字符出现次数作为键
func groupAnagrams2(strs []string) [][]string {
	groups := make(map[string][]string)

	for _, str := range strs {
		// 统计字符出现次数
		count := make([]int, 26)
		for _, char := range str {
			count[char-'a']++
		}

		// 使用计数作为键
		key := fmt.Sprintf("%v", count)
		groups[key] = append(groups[key], str)
	}

	// 转换为结果格式
	result := make([][]string, 0, len(groups))
	for _, group := range groups {
		result = append(result, group)
	}

	return result
}

// 方法三：哈希优化算法
// 使用字符计数的哈希值作为键，优化查找效率
func groupAnagrams3(strs []string) [][]string {
	groups := make(map[uint64][]string)

	for _, str := range strs {
		// 统计字符出现次数
		count := make([]int, 26)
		for _, char := range str {
			count[char-'a']++
		}

		// 计算字符计数的哈希值
		hash := calculateCountHash(count)
		groups[hash] = append(groups[hash], str)
	}

	// 转换为结果格式
	result := make([][]string, 0, len(groups))
	for _, group := range groups {
		result = append(result, group)
	}

	return result
}

// 计算字符计数哈希值的辅助函数
func calculateCountHash(count []int) uint64 {
	var hash uint64 = 0
	for i, c := range count {
		hash = hash*31 + uint64(i)*1000 + uint64(c)
	}
	return hash
}

// 方法四：位运算算法（修正版）
// 使用字符计数的字符串表示作为键，避免位运算的局限性
func groupAnagrams4(strs []string) [][]string {
	groups := make(map[string][]string)

	for _, str := range strs {
		// 统计字符出现次数
		count := make([]int, 26)
		for _, char := range str {
			count[char-'a']++
		}

		// 使用字符计数的字符串表示作为键
		key := buildKey(count)
		groups[key] = append(groups[key], str)
	}

	// 转换为结果格式
	result := make([][]string, 0, len(groups))
	for _, group := range groups {
		result = append(result, group)
	}

	return result
}

// 构建字符计数键的辅助函数
func buildKey(count []int) string {
	var key strings.Builder
	for i, c := range count {
		if c > 0 {
			key.WriteByte(byte('a' + i))
			key.WriteString(fmt.Sprintf("%d", c))
		}
	}
	return key.String()
}

// 辅助函数：创建测试用例
func createTestCases() []struct {
	strs []string
	name string
} {
	return []struct {
		strs []string
		name string
	}{
		{[]string{"eat", "tea", "tan", "ate", "nat", "bat"}, "示例1: 基本异位词分组"},
		{[]string{""}, "示例2: 空字符串"},
		{[]string{"a"}, "示例3: 单个字符"},
		{[]string{"a", "a"}, "测试1: 相同字符串"},
		{[]string{"abc", "bca", "cab"}, "测试2: 三个异位词"},
		{[]string{"listen", "silent", "enlist"}, "测试3: 长字符串异位词"},
		{[]string{"rat", "tar", "art", "tar"}, "测试4: 包含重复的异位词"},
		{[]string{"a", "b", "c"}, "测试5: 无异位词"},
		{[]string{"a", "aa", "aaa"}, "测试6: 不同长度的字符串"},
		{[]string{"eat", "tea", "tan", "ate", "nat", "bat", "tab", "act", "cat"}, "测试7: 复杂分组"},
	}
}

// 性能测试函数
func benchmarkAlgorithm(algorithm func([]string) [][]string, strs []string, name string) {
	iterations := 100
	start := time.Now()

	for i := 0; i < iterations; i++ {
		algorithm(strs)
	}

	duration := time.Since(start)
	avgTime := duration.Nanoseconds() / int64(iterations)

	fmt.Printf("%s: 平均执行时间 %d 纳秒\n", name, avgTime)
}

// 辅助函数：验证结果是否正确
func validateResult(strs []string, result [][]string) bool {
	// 验证所有字符串都被分组
	totalStrings := 0
	for _, group := range result {
		totalStrings += len(group)
	}
	if totalStrings != len(strs) {
		return false
	}

	// 验证每个分组内的字符串都是异位词
	for _, group := range result {
		if len(group) > 1 {
			// 检查组内所有字符串是否都是异位词
			base := group[0]
			for i := 1; i < len(group); i++ {
				if !areAnagrams(base, group[i]) {
					return false
				}
			}
		}
	}

	// 验证不同分组之间不是异位词
	for i := 0; i < len(result); i++ {
		for j := i + 1; j < len(result); j++ {
			if areAnagrams(result[i][0], result[j][0]) {
				return false
			}
		}
	}

	return true
}

// 检查两个字符串是否为异位词
func areAnagrams(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}

	count1 := make([]int, 26)
	count2 := make([]int, 26)

	for _, char := range s1 {
		count1[char-'a']++
	}
	for _, char := range s2 {
		count2[char-'a']++
	}

	for i := 0; i < 26; i++ {
		if count1[i] != count2[i] {
			return false
		}
	}

	return true
}

// 辅助函数：比较两个结果是否相同
func compareResults(result1, result2 [][]string) bool {
	if len(result1) != len(result2) {
		return false
	}

	// 将结果转换为可比较的格式
	normalize := func(result [][]string) []string {
		var normalized []string
		for _, group := range result {
			sort.Strings(group)
			normalized = append(normalized, strings.Join(group, ","))
		}
		sort.Strings(normalized)
		return normalized
	}

	norm1 := normalize(result1)
	norm2 := normalize(result2)

	for i := range norm1 {
		if norm1[i] != norm2[i] {
			return false
		}
	}

	return true
}

// 辅助函数：打印分组结果
func printGroupResult(strs []string, result [][]string, title string) {
	fmt.Printf("%s: strs=%v -> %d 个分组\n", title, strs, len(result))
	if len(result) <= 3 {
		fmt.Printf("  分组结果: %v\n", result)
	}
}

func main() {
	fmt.Println("=== 49. 字母异位词分组 ===")
	fmt.Println()

	// 创建测试用例
	testCases := createTestCases()
	algorithms := []struct {
		name string
		fn   func([]string) [][]string
	}{
		{"排序分组算法", groupAnagrams1},
		{"字符计数算法", groupAnagrams2},
		{"哈希优化算法", groupAnagrams3},
		{"位运算算法", groupAnagrams4},
	}

	// 运行测试
	fmt.Println("=== 算法正确性测试 ===")
	for _, testCase := range testCases {
		fmt.Printf("测试: %s\n", testCase.name)

		results := make([][][]string, len(algorithms))
		for i, algo := range algorithms {
			results[i] = algo.fn(testCase.strs)
		}

		// 验证所有算法结果一致
		allEqual := true
		for i := 1; i < len(results); i++ {
			if !compareResults(results[i], results[0]) {
				allEqual = false
				break
			}
		}

		// 验证结果是否正确
		allValid := true
		for _, result := range results {
			if !validateResult(testCase.strs, result) {
				allValid = false
				break
			}
		}

		if allEqual && allValid {
			fmt.Printf("  ✅ 所有算法结果一致且正确: %d 个分组\n", len(results[0]))
			if len(testCase.strs) <= 6 {
				printGroupResult(testCase.strs, results[0], "  分组结果")
			}
		} else {
			fmt.Printf("  ❌ 算法结果不一致或错误\n")
			for i, algo := range algorithms {
				fmt.Printf("    %s: %d 个分组\n", algo.name, len(results[i]))
			}
		}
		fmt.Println()
	}

	// 性能测试
	fmt.Println("=== 性能测试 ===")
	performanceStrs := []string{"eat", "tea", "tan", "ate", "nat", "bat", "tab", "act", "cat", "tac"}

	fmt.Printf("测试数据: %v\n", performanceStrs)
	fmt.Println()

	for _, algo := range algorithms {
		benchmarkAlgorithm(algo.fn, performanceStrs, algo.name)
	}
	fmt.Println()

	// 算法分析
	fmt.Println("=== 算法分析 ===")
	fmt.Println("字母异位词分组问题的特点:")
	fmt.Println("1. 需要将字母异位词分组在一起")
	fmt.Println("2. 字母异位词由相同字母组成但顺序不同")
	fmt.Println("3. 需要高效的键生成方法")
	fmt.Println("4. 字符计数算法是最优解法")
	fmt.Println()

	// 复杂度分析
	fmt.Println("=== 复杂度分析 ===")
	fmt.Println("时间复杂度:")
	fmt.Println("- 排序分组: O(nklogk)，需要对每个字符串排序")
	fmt.Println("- 字符计数: O(nk)，只需要遍历每个字符一次")
	fmt.Println("- 哈希优化: O(nk)，哈希计算和查找")
	fmt.Println("- 位运算: O(nk)，位运算计算")
	fmt.Println()
	fmt.Println("空间复杂度:")
	fmt.Println("- 所有算法: O(nk)，需要存储所有字符串和分组结果")
	fmt.Println()

	// 算法总结
	fmt.Println("=== 算法总结 ===")
	fmt.Println("1. 排序分组算法：最直观的解法，逻辑清晰")
	fmt.Println("2. 字符计数算法：效率最高的解法，避免排序")
	fmt.Println("3. 哈希优化算法：使用哈希表优化查找")
	fmt.Println("4. 位运算算法：使用位运算优化空间")
	fmt.Println()
	fmt.Println("推荐使用：字符计数算法（方法二），效率最高")
	fmt.Println()

	// 应用场景
	fmt.Println("=== 应用场景 ===")
	fmt.Println("- 文本分析：分析文本中的词汇模式")
	fmt.Println("- 数据挖掘：发现数据中的相似模式")
	fmt.Println("- 搜索引擎：处理搜索查询的变体")
	fmt.Println("- 算法竞赛：字符串处理的经典应用")
	fmt.Println("- 自然语言处理：词汇相似性分析")
	fmt.Println()

	// 优化技巧总结
	fmt.Println("=== 优化技巧总结 ===")
	fmt.Println("1. 排序分组：最直观的解法，逻辑清晰")
	fmt.Println("2. 字符计数：效率最高的解法，避免排序")
	fmt.Println("3. 哈希优化：使用哈希表优化查找")
	fmt.Println("4. 位运算：使用位运算优化空间")
	fmt.Println("5. 键的生成：理解不同键生成方法的优劣")
	fmt.Println("6. 分组策略：学会高效的分组算法")
}
