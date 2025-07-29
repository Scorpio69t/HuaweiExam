package main

import (
	"fmt"
	"strings"
	"time"
)

// 解法一：栈匹配法（推荐解法）
// 时间复杂度：O(n)，空间复杂度：O(n)
func isValid(s string) bool {
	// 奇数长度的字符串不可能有效
	if len(s)%2 == 1 {
		return false
	}

	// 括号映射表：右括号 -> 左括号
	pairs := map[byte]byte{
		')': '(',
		']': '[',
		'}': '{',
	}

	// 使用切片模拟栈
	stack := []byte{}

	for i := 0; i < len(s); i++ {
		char := s[i]

		// 如果是右括号
		if pair, exists := pairs[char]; exists {
			// 栈为空或栈顶元素不匹配
			if len(stack) == 0 || stack[len(stack)-1] != pair {
				return false
			}
			// 弹出栈顶元素（匹配成功）
			stack = stack[:len(stack)-1]
		} else {
			// 如果是左括号，压入栈中
			stack = append(stack, char)
		}
	}

	// 栈为空表示所有括号都匹配
	return len(stack) == 0
}

// 解法二：替换法
// 时间复杂度：O(n²)，空间复杂度：O(n)
func isValidReplace(s string) bool {
	// 不断替换匹配的括号对
	for len(s) > 0 {
		oldLen := len(s)
		// 替换所有可能的括号对
		s = replaceAll(s, "()", "")
		s = replaceAll(s, "[]", "")
		s = replaceAll(s, "{}", "")

		// 如果没有任何替换发生，说明无法继续匹配
		if len(s) == oldLen {
			break
		}
	}

	return len(s) == 0
}

// 简单的字符串替换函数
func replaceAll(s, old, new string) string {
	result := ""
	i := 0
	for i < len(s) {
		if i <= len(s)-len(old) && s[i:i+len(old)] == old {
			result += new
			i += len(old)
		} else {
			result += string(s[i])
			i++
		}
	}
	return result
}

// 解法三：计数器法（仅适用于单一类型括号）
// 时间复杂度：O(n)，空间复杂度：O(1)
func isValidSimple(s string) bool {
	// 仅处理圆括号的情况
	count := 0
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			count++
		} else if s[i] == ')' {
			count--
			// 右括号过多
			if count < 0 {
				return false
			}
		}
	}
	return count == 0
}

// 解法四：递归法
// 时间复杂度：O(n)，空间复杂度：O(n)
func isValidRecursive(s string) bool {
	return isValidHelper(s, 0, []byte{})
}

func isValidHelper(s string, index int, stack []byte) bool {
	// 递归终止条件
	if index == len(s) {
		return len(stack) == 0
	}

	char := s[index]
	pairs := map[byte]byte{
		')': '(',
		']': '[',
		'}': '{',
	}

	// 如果是右括号
	if pair, exists := pairs[char]; exists {
		// 栈为空或栈顶不匹配
		if len(stack) == 0 || stack[len(stack)-1] != pair {
			return false
		}
		// 弹出栈顶，继续递归
		return isValidHelper(s, index+1, stack[:len(stack)-1])
	} else {
		// 如果是左括号，压入栈中继续递归
		newStack := make([]byte, len(stack)+1)
		copy(newStack, stack)
		newStack[len(stack)] = char
		return isValidHelper(s, index+1, newStack)
	}
}

// 解法五：优化栈法（预分配容量）
// 时间复杂度：O(n)，空间复杂度：O(n)
func isValidOptimized(s string) bool {
	n := len(s)
	if n%2 == 1 {
		return false
	}

	// 预分配栈容量，避免动态扩容
	stack := make([]byte, 0, n/2)

	for i := 0; i < n; i++ {
		char := s[i]
		switch char {
		case '(':
			stack = append(stack, char)
		case ')':
			if len(stack) == 0 || stack[len(stack)-1] != '(' {
				return false
			}
			stack = stack[:len(stack)-1]
		case '[':
			stack = append(stack, char)
		case ']':
			if len(stack) == 0 || stack[len(stack)-1] != '[' {
				return false
			}
			stack = stack[:len(stack)-1]
		case '{':
			stack = append(stack, char)
		case '}':
			if len(stack) == 0 || stack[len(stack)-1] != '{' {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}

	return len(stack) == 0
}

// 测试函数
func testIsValid() {
	testCases := []struct {
		input    string
		expected bool
		desc     string
	}{
		{"()", true, "简单圆括号"},
		{"()[]{}", true, "三种括号组合"},
		{"(]", false, "类型不匹配"},
		{"([])", true, "嵌套括号"},
		{"([)]", false, "交叉嵌套"},
		{"", true, "空字符串"},
		{"(", false, "单个左括号"},
		{")", false, "单个右括号"},
		{"((", false, "只有左括号"},
		{"))", false, "只有右括号"},
		{"(())", true, "双层嵌套"},
		{"({[]})", true, "多层复杂嵌套"},
		{"({[}])", false, "错误的嵌套顺序"},
		{"((()))", true, "深度嵌套"},
		{"{[]}", true, "不同类型嵌套"},
	}

	fmt.Println("=== 有效的括号测试 ===\n")

	for i, tc := range testCases {
		// 测试主要解法
		result1 := isValid(tc.input)
		result2 := isValidOptimized(tc.input)
		result3 := isValidRecursive(tc.input)

		status := "✅"
		if result1 != tc.expected {
			status = "❌"
		}

		fmt.Printf("测试 %d: %s\n", i+1, tc.desc)
		fmt.Printf("输入: \"%s\"\n", tc.input)
		fmt.Printf("期望: %t\n", tc.expected)
		fmt.Printf("栈法: %t\n", result1)
		fmt.Printf("优化: %t\n", result2)
		fmt.Printf("递归: %t\n", result3)
		fmt.Printf("结果: %s\n", status)
		fmt.Println(strings.Repeat("-", 30))
	}
}

// 性能测试
func benchmarkIsValid() {
	fmt.Println("\n=== 性能测试 ===\n")

	// 构造测试数据
	testData := []string{
		generateValidBrackets(100),  // 短字符串
		generateValidBrackets(1000), // 中等字符串
		generateValidBrackets(5000), // 长字符串
	}

	algorithms := []struct {
		name string
		fn   func(string) bool
	}{
		{"栈匹配法", isValid},
		{"优化栈法", isValidOptimized},
		{"递归法", isValidRecursive},
		{"替换法", isValidReplace},
	}

	for i, data := range testData {
		fmt.Printf("测试数据 %d (长度: %d):\n", i+1, len(data))

		for _, algo := range algorithms {
			start := time.Now()
			result := algo.fn(data)
			duration := time.Since(start)

			fmt.Printf("  %s: %t, 耗时: %v\n", algo.name, result, duration)
		}
		fmt.Println()
	}
}

// 生成有效的括号字符串用于测试
func generateValidBrackets(n int) string {
	if n%2 == 1 {
		n-- // 确保是偶数
	}

	result := make([]byte, 0, n)
	pairs := n / 2

	// 生成嵌套的括号
	for i := 0; i < pairs; i++ {
		bracketType := i % 3
		switch bracketType {
		case 0:
			result = append(result, '(')
		case 1:
			result = append(result, '[')
		case 2:
			result = append(result, '{')
		}
	}

	// 添加对应的右括号
	for i := pairs - 1; i >= 0; i-- {
		bracketType := i % 3
		switch bracketType {
		case 0:
			result = append(result, ')')
		case 1:
			result = append(result, ']')
		case 2:
			result = append(result, '}')
		}
	}

	return string(result)
}

func main() {
	fmt.Println("20. 有效的括号 - 多种解法实现")
	fmt.Println("========================================")

	// 基础功能测试
	testIsValid()

	// 性能对比测试
	benchmarkIsValid()

	// 展示算法特点
	fmt.Println("\n=== 算法特点分析 ===")
	fmt.Println("1. 栈匹配法：经典解法，时间O(n)，空间O(n)，推荐使用")
	fmt.Println("2. 优化栈法：预分配容量，避免动态扩容，性能略优")
	fmt.Println("3. 递归法：代码简洁，但可能栈溢出，适合教学")
	fmt.Println("4. 替换法：思路直观，但效率较低O(n²)，不推荐")
	fmt.Println("5. 计数器法：仅适用于单一类型括号，空间O(1)")

	fmt.Println("\n=== 关键技巧总结 ===")
	fmt.Println("• 使用map建立括号对应关系，提高查找效率")
	fmt.Println("• 奇数长度字符串可以提前返回false")
	fmt.Println("• 栈为空时遇到右括号立即返回false")
	fmt.Println("• 预分配栈容量可以避免动态扩容开销")
	fmt.Println("• 选择合适的数据结构是性能优化的关键")
}
