package main

import (
	"fmt"
	"strings"
)

// letterCombinations 电话号码的字母组合 - 回溯法
// 时间复杂度: O(4^n)，其中n是digits的长度，每个数字最多对应4个字母
// 空间复杂度: O(n)，递归调用栈深度
func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}

	// 数字到字母的映射
	digitMap := map[byte]string{
		'2': "abc",
		'3': "def",
		'4': "ghi",
		'5': "jkl",
		'6': "mno",
		'7': "pqrs",
		'8': "tuv",
		'9': "wxyz",
	}

	var result []string
	var backtrack func(index int, current string)

	backtrack = func(index int, current string) {
		// 如果已经处理完所有数字，添加当前组合到结果
		if index == len(digits) {
			result = append(result, current)
			return
		}

		// 获取当前数字对应的字母
		letters := digitMap[digits[index]]

		// 尝试每个字母
		for _, letter := range letters {
			backtrack(index+1, current+string(letter))
		}
	}

	backtrack(0, "")
	return result
}

// letterCombinationsIterative 迭代法 - 使用队列
// 时间复杂度: O(4^n)
// 空间复杂度: O(4^n)，存储所有可能的组合
func letterCombinationsIterative(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}

	// 数字到字母的映射
	digitMap := map[byte]string{
		'2': "abc",
		'3': "def",
		'4': "ghi",
		'5': "jkl",
		'6': "mno",
		'7': "pqrs",
		'8': "tuv",
		'9': "wxyz",
	}

	// 使用队列存储当前所有可能的组合
	queue := []string{""}

	// 逐个处理每个数字
	for i := 0; i < len(digits); i++ {
		letters := digitMap[digits[i]]
		levelSize := len(queue)

		// 处理当前层的所有组合
		for j := 0; j < levelSize; j++ {
			current := queue[0]
			queue = queue[1:]

			// 为当前组合添加每个可能的字母
			for _, letter := range letters {
				queue = append(queue, current+string(letter))
			}
		}
	}

	return queue
}

// letterCombinationsOptimized 优化版本 - 使用strings.Builder
// 时间复杂度: O(4^n)
// 空间复杂度: O(n)
func letterCombinationsOptimized(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}

	// 使用数组映射，避免map查找开销
	digitMap := [][]byte{
		{}, {}, // 0, 1 不对应字母
		{'a', 'b', 'c'},      // 2
		{'d', 'e', 'f'},      // 3
		{'g', 'h', 'i'},      // 4
		{'j', 'k', 'l'},      // 5
		{'m', 'n', 'o'},      // 6
		{'p', 'q', 'r', 's'}, // 7
		{'t', 'u', 'v'},      // 8
		{'w', 'x', 'y', 'z'}, // 9
	}

	var result []string
	var backtrack func(index int, current *strings.Builder)

	backtrack = func(index int, current *strings.Builder) {
		if index == len(digits) {
			result = append(result, current.String())
			return
		}

		digit := digits[index] - '0'
		letters := digitMap[digit]

		for _, letter := range letters {
			current.WriteByte(letter)
			backtrack(index+1, current)
			// 回溯：移除最后添加的字符
			currentStr := current.String()
			current.Reset()
			current.WriteString(currentStr[:len(currentStr)-1])
		}
	}

	backtrack(0, &strings.Builder{})
	return result
}

// letterCombinationsBFS BFS方法 - 层序遍历
// 时间复杂度: O(4^n)
// 空间复杂度: O(4^n)
func letterCombinationsBFS(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}

	digitMap := map[byte]string{
		'2': "abc",
		'3': "def",
		'4': "ghi",
		'5': "jkl",
		'6': "mno",
		'7': "pqrs",
		'8': "tuv",
		'9': "wxyz",
	}

	// BFS队列
	queue := []string{""}

	for i := 0; i < len(digits); i++ {
		letters := digitMap[digits[i]]
		levelSize := len(queue)

		// 处理当前层的所有组合
		for j := 0; j < levelSize; j++ {
			current := queue[0]
			queue = queue[1:]

			// 为当前组合添加每个可能的字母
			for _, letter := range letters {
				queue = append(queue, current+string(letter))
			}
		}
	}

	return queue
}

// letterCombinationsRecursive 纯递归方法 - 分治思想
// 时间复杂度: O(4^n)
// 空间复杂度: O(n)
func letterCombinationsRecursive(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}

	if len(digits) == 1 {
		return getLetters(digits[0])
	}

	// 分治：处理第一个数字，递归处理剩余数字
	firstLetters := getLetters(digits[0])
	remainingCombinations := letterCombinationsRecursive(digits[1:])

	var result []string
	for _, letter := range firstLetters {
		for _, combination := range remainingCombinations {
			result = append(result, string(letter)+combination)
		}
	}

	return result
}

// getLetters 获取单个数字对应的字母
func getLetters(digit byte) []string {
	switch digit {
	case '2':
		return []string{"a", "b", "c"}
	case '3':
		return []string{"d", "e", "f"}
	case '4':
		return []string{"g", "h", "i"}
	case '5':
		return []string{"j", "k", "l"}
	case '6':
		return []string{"m", "n", "o"}
	case '7':
		return []string{"p", "q", "r", "s"}
	case '8':
		return []string{"t", "u", "v"}
	case '9':
		return []string{"w", "x", "y", "z"}
	default:
		return []string{}
	}
}

func main() {
	// 测试用例1
	digits1 := "23"
	result1 := letterCombinations(digits1)
	fmt.Printf("示例1: digits = \"%s\"\n", digits1)
	fmt.Printf("输出: %v\n", result1)
	fmt.Printf("期望: [ad ae af bd be bf cd ce cf]\n")
	fmt.Printf("结果正确: %t\n", len(result1) == 9)
	fmt.Println()

	// 测试用例2
	digits2 := ""
	result2 := letterCombinations(digits2)
	fmt.Printf("示例2: digits = \"%s\"\n", digits2)
	fmt.Printf("输出: %v\n", result2)
	fmt.Printf("期望: []\n")
	fmt.Printf("结果正确: %t\n", len(result2) == 0)
	fmt.Println()

	// 测试用例3
	digits3 := "2"
	result3 := letterCombinations(digits3)
	fmt.Printf("示例3: digits = \"%s\"\n", digits3)
	fmt.Printf("输出: %v\n", result3)
	fmt.Printf("期望: [a b c]\n")
	fmt.Printf("结果正确: %t\n", len(result3) == 3)
	fmt.Println()

	// 额外测试用例
	digits4 := "234"
	result4 := letterCombinations(digits4)
	fmt.Printf("额外测试: digits = \"%s\"\n", digits4)
	fmt.Printf("输出数量: %d\n", len(result4))
	fmt.Printf("期望数量: 27 (3×3×3)\n")
	fmt.Printf("结果正确: %t\n", len(result4) == 27)
	fmt.Println()

	// 测试迭代版本
	fmt.Println("=== 迭代版本测试 ===")
	result1Iter := letterCombinationsIterative(digits1)
	result2Iter := letterCombinationsIterative(digits2)
	fmt.Printf("迭代版本示例1: %v\n", result1Iter)
	fmt.Printf("迭代版本示例2: %v\n", result2Iter)
	fmt.Printf("结果一致: %t\n", len(result1Iter) == len(result1) && len(result2Iter) == len(result2))
	fmt.Println()

	// 测试优化版本
	fmt.Println("=== 优化版本测试 ===")
	result1Opt := letterCombinationsOptimized(digits1)
	result2Opt := letterCombinationsOptimized(digits2)
	fmt.Printf("优化版本示例1: %v\n", result1Opt)
	fmt.Printf("优化版本示例2: %v\n", result2Opt)
	fmt.Printf("结果一致: %t\n", len(result1Opt) == len(result1) && len(result2Opt) == len(result2))
	fmt.Println()

	// 测试BFS版本
	fmt.Println("=== BFS版本测试 ===")
	result1BFS := letterCombinationsBFS(digits1)
	result2BFS := letterCombinationsBFS(digits2)
	fmt.Printf("BFS版本示例1: %v\n", result1BFS)
	fmt.Printf("BFS版本示例2: %v\n", result2BFS)
	fmt.Printf("结果一致: %t\n", len(result1BFS) == len(result1) && len(result2BFS) == len(result2))
	fmt.Println()

	// 测试递归版本
	fmt.Println("=== 递归版本测试 ===")
	result1Rec := letterCombinationsRecursive(digits1)
	result2Rec := letterCombinationsRecursive(digits2)
	fmt.Printf("递归版本示例1: %v\n", result1Rec)
	fmt.Printf("递归版本示例2: %v\n", result2Rec)
	fmt.Printf("结果一致: %t\n", len(result1Rec) == len(result1) && len(result2Rec) == len(result2))
	fmt.Println()

	// 边界值测试
	fmt.Println("=== 边界值测试 ===")
	boundaryTests := []string{
		"",     // 空字符串
		"2",    // 单个数字
		"99",   // 两个相同数字
		"2345", // 四个不同数字
		"7777", // 四个相同数字
	}

	for _, test := range boundaryTests {
		result := letterCombinations(test)
		expectedCount := 1
		for _, digit := range test {
			switch digit {
			case '7', '9':
				expectedCount *= 4
			default:
				expectedCount *= 3
			}
		}
		fmt.Printf("digits = \"%s\", result count = %d, expected = %d\n", test, len(result), expectedCount)
	}
}
