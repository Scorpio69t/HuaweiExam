package main

import (
	"fmt"
	"regexp"
	"strings"
	"time"
)

// 方法一：反向遍历算法（最优解法）
func lengthOfLastWord1(s string) int {
	length := 0
	i := len(s) - 1

	// 跳过尾部空格
	for i >= 0 && s[i] == ' ' {
		i--
	}

	// 计算单词长度
	for i >= 0 && s[i] != ' ' {
		length++
		i--
	}

	return length
}

// 方法二：字符串分割算法
func lengthOfLastWord2(s string) int {
	// 去除首尾空格并按空格分割
	words := strings.Fields(s)
	if len(words) == 0 {
		return 0
	}
	return len(words[len(words)-1])
}

// 方法三：双指针算法
func lengthOfLastWord3(s string) int {
	end := len(s) - 1

	// 找到最后一个单词的结束位置
	for end >= 0 && s[end] == ' ' {
		end--
	}

	if end < 0 {
		return 0
	}

	// 找到最后一个单词的开始位置
	start := end
	for start >= 0 && s[start] != ' ' {
		start--
	}

	return end - start
}

// 方法四：正则表达式算法
func lengthOfLastWord4(s string) int {
	// 使用正则表达式匹配最后一个单词
	re := regexp.MustCompile(`\w+`)
	matches := re.FindAllString(s, -1)
	if len(matches) == 0 {
		return 0
	}
	return len(matches[len(matches)-1])
}

// 测试用例
func createTestCases() []struct {
	input    string
	expected int
	name     string
} {
	return []struct {
		input    string
		expected int
		name     string
	}{
		{"Hello World", 5, "示例1: 基础测试"},
		{"   fly me   to   the moon  ", 4, "示例2: 多空格"},
		{"luffy is still joyboy", 6, "示例3: 无多余空格"},
		{"a", 1, "测试1: 单个字符"},
		{"a ", 1, "测试2: 单个字符+空格"},
		{" a", 1, "测试3: 空格+单个字符"},
		{"  hello  ", 5, "测试4: 前后空格"},
		{"test", 4, "测试5: 无空格"},
		{"one two three", 5, "测试6: 多个单词"},
		{"   ", 0, "测试7: 仅空格（理论上不会出现）"},
	}
}

// 性能测试
func benchmarkAlgorithm(algorithm func(string) int, input string, name string) {
	iterations := 10000
	start := time.Now()

	for i := 0; i < iterations; i++ {
		algorithm(input)
	}

	duration := time.Since(start)
	avgTime := duration.Nanoseconds() / int64(iterations)
	fmt.Printf("%s: 平均执行时间 %d 纳秒\n", name, avgTime)
}

func main() {
	fmt.Println("=== 58. 最后一个单词的长度 ===")
	fmt.Println()

	testCases := createTestCases()
	algorithms := []struct {
		name string
		fn   func(string) int
	}{
		{"反向遍历算法", lengthOfLastWord1},
		{"字符串分割算法", lengthOfLastWord2},
		{"双指针算法", lengthOfLastWord3},
		{"正则表达式算法", lengthOfLastWord4},
	}

	// 正确性测试
	fmt.Println("=== 算法正确性测试 ===")
	for _, testCase := range testCases {
		fmt.Printf("测试: %s\n", testCase.name)
		fmt.Printf("  输入: \"%s\"\n", testCase.input)

		results := make([]int, len(algorithms))
		for i, algo := range algorithms {
			results[i] = algo.fn(testCase.input)
		}

		allEqual := true
		for i := 1; i < len(results); i++ {
			if results[i] != results[0] {
				allEqual = false
				break
			}
		}

		allValid := results[0] == testCase.expected

		if allEqual && allValid {
			fmt.Printf("  ✅ 所有算法结果一致且正确: %d\n", results[0])
		} else {
			fmt.Printf("  ❌ 算法结果不一致或错误\n")
			fmt.Printf("  预期: %d\n", testCase.expected)
			for i, algo := range algorithms {
				fmt.Printf("  %s: %d\n", algo.name, results[i])
			}
		}
		fmt.Println()
	}

	// 性能测试
	fmt.Println("=== 性能测试 ===")
	perfInput := "   fly me   to   the moon  "

	for _, algo := range algorithms {
		benchmarkAlgorithm(algo.fn, perfInput, algo.name)
	}
	fmt.Println()

	// 算法分析
	fmt.Println("=== 算法分析 ===")
	fmt.Println("最后一个单词长度问题的特点:")
	fmt.Println("1. 从后向前遍历更高效")
	fmt.Println("2. 需要处理尾部空格")
	fmt.Println("3. 反向遍历法是最优解法")
	fmt.Println()

	fmt.Println("=== 复杂度分析 ===")
	fmt.Println("时间复杂度:")
	fmt.Println("- 反向遍历: O(n)，最坏情况遍历整个字符串")
	fmt.Println("- 字符串分割: O(n)，需要遍历整个字符串")
	fmt.Println("- 双指针: O(n)，需要找到单词边界")
	fmt.Println("- 正则表达式: O(n)，正则匹配需要遍历")
	fmt.Println()
	fmt.Println("空间复杂度:")
	fmt.Println("- 反向遍历: O(1)，只使用常数空间")
	fmt.Println("- 字符串分割: O(n)，需要存储分割后的单词")
	fmt.Println("- 双指针: O(1)，只使用常数空间")
	fmt.Println("- 正则表达式: O(n)，需要存储匹配结果")
	fmt.Println()

	fmt.Println("=== 优化技巧总结 ===")
	fmt.Println("1. 反向遍历：从后向前处理更高效")
	fmt.Println("2. 边界检查：注意索引越界")
	fmt.Println("3. 空格处理：先跳过尾部空格")
	fmt.Println("4. 一次遍历：O(n)时间复杂度")
	fmt.Println()

	fmt.Println("推荐使用：反向遍历算法（方法一），效率最高")
}
