package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"unicode"
)

// myAtoi 字符串转换整数 - 状态机方法
// 时间复杂度: O(n)，其中n是字符串长度
// 空间复杂度: O(1)
func myAtoi(s string) int {
	if len(s) == 0 {
		return 0
	}

	// 状态机状态
	const (
		START = iota // 开始状态
		SIGN         // 符号状态
		DIGIT        // 数字状态
		END          // 结束状态
	)

	state := START
	sign := 1
	result := 0

	for i := 0; i < len(s) && state != END; i++ {
		char := s[i]

		switch state {
		case START:
			if char == ' ' {
				// 继续跳过空格
				continue
			} else if char == '+' || char == '-' {
				// 遇到符号
				if char == '-' {
					sign = -1
				}
				state = SIGN
			} else if unicode.IsDigit(rune(char)) {
				// 遇到数字
				result = int(char - '0')
				state = DIGIT
			} else {
				// 遇到非数字非空格字符
				state = END
			}

		case SIGN:
			if unicode.IsDigit(rune(char)) {
				// 符号后遇到数字
				result = int(char - '0')
				state = DIGIT
			} else {
				// 符号后遇到非数字
				state = END
			}

		case DIGIT:
			if unicode.IsDigit(rune(char)) {
				// 继续读取数字
				digit := int(char - '0')

				// 检查溢出
				if result > math.MaxInt32/10 || (result == math.MaxInt32/10 && digit > 7) {
					if sign == 1 {
						return math.MaxInt32
					} else {
						return math.MinInt32
					}
				}

				result = result*10 + digit
			} else {
				// 遇到非数字字符，结束
				state = END
			}
		}
	}

	return sign * result
}

// myAtoiOptimized 优化版本 - 简化逻辑
// 时间复杂度: O(n)
// 空间复杂度: O(1)
func myAtoiOptimized(s string) int {
	if len(s) == 0 {
		return 0
	}

	// 跳过前导空格
	i := 0
	for i < len(s) && s[i] == ' ' {
		i++
	}

	if i == len(s) {
		return 0
	}

	// 处理符号
	sign := 1
	if s[i] == '+' || s[i] == '-' {
		if s[i] == '-' {
			sign = -1
		}
		i++
	}

	// 读取数字
	result := 0
	for i < len(s) && unicode.IsDigit(rune(s[i])) {
		digit := int(s[i] - '0')

		// 检查溢出
		if result > math.MaxInt32/10 || (result == math.MaxInt32/10 && digit > 7) {
			if sign == 1 {
				return math.MaxInt32
			} else {
				return math.MinInt32
			}
		}

		result = result*10 + digit
		i++
	}

	return sign * result
}

// myAtoiRegex 正则表达式思路 - 使用字符串处理
// 时间复杂度: O(n)
// 空间复杂度: O(n)
func myAtoiRegex(s string) int {
	// 去除前导空格
	s = strings.TrimSpace(s)

	if len(s) == 0 {
		return 0
	}

	// 检查符号
	sign := 1
	if s[0] == '+' || s[0] == '-' {
		if s[0] == '-' {
			sign = -1
		}
		s = s[1:]
	}

	if len(s) == 0 {
		return 0
	}

	// 提取数字部分
	var digits strings.Builder
	for _, char := range s {
		if unicode.IsDigit(char) {
			digits.WriteRune(char)
		} else {
			break
		}
	}

	if digits.Len() == 0 {
		return 0
	}

	// 转换为整数
	result, err := strconv.Atoi(digits.String())
	if err != nil {
		return 0
	}

	// 应用符号
	result = sign * result

	// 检查溢出
	if result > math.MaxInt32 {
		return math.MaxInt32
	}
	if result < math.MinInt32 {
		return math.MinInt32
	}

	return result
}

// myAtoiBitwise 位运算优化版本
// 时间复杂度: O(n)
// 空间复杂度: O(1)
func myAtoiBitwise(s string) int {
	if len(s) == 0 {
		return 0
	}

	// 跳过前导空格
	i := 0
	for i < len(s) && s[i] == ' ' {
		i++
	}

	if i == len(s) {
		return 0
	}

	// 处理符号
	sign := 1
	if s[i] == '+' || s[i] == '-' {
		if s[i] == '-' {
			sign = -1
		}
		i++
	}

	// 读取数字
	result := 0
	for i < len(s) && unicode.IsDigit(rune(s[i])) {
		digit := int(s[i] - '0')

		// 使用位运算优化溢出检查
		if result > (math.MaxInt32-digit)/10 {
			if sign == 1 {
				return math.MaxInt32
			} else {
				return math.MinInt32
			}
		}

		result = result*10 + digit
		i++
	}

	return sign * result
}

func main() {
	// 测试用例1
	s1 := "42"
	result1 := myAtoi(s1)
	fmt.Printf("示例1: s = \"%s\"\n", s1)
	fmt.Printf("输出: %d\n", result1)
	fmt.Printf("期望: 42\n")
	fmt.Printf("结果: %t\n", result1 == 42)
	fmt.Println()

	// 测试用例2
	s2 := " -042"
	result2 := myAtoi(s2)
	fmt.Printf("示例2: s = \"%s\"\n", s2)
	fmt.Printf("输出: %d\n", result2)
	fmt.Printf("期望: -42\n")
	fmt.Printf("结果: %t\n", result2 == -42)
	fmt.Println()

	// 测试用例3
	s3 := "1337c0d3"
	result3 := myAtoi(s3)
	fmt.Printf("示例3: s = \"%s\"\n", s3)
	fmt.Printf("输出: %d\n", result3)
	fmt.Printf("期望: 1337\n")
	fmt.Printf("结果: %t\n", result3 == 1337)
	fmt.Println()

	// 测试用例4
	s4 := "0-1"
	result4 := myAtoi(s4)
	fmt.Printf("示例4: s = \"%s\"\n", s4)
	fmt.Printf("输出: %d\n", result4)
	fmt.Printf("期望: 0\n")
	fmt.Printf("结果: %t\n", result4 == 0)
	fmt.Println()

	// 测试用例5
	s5 := "words and 987"
	result5 := myAtoi(s5)
	fmt.Printf("示例5: s = \"%s\"\n", s5)
	fmt.Printf("输出: %d\n", result5)
	fmt.Printf("期望: 0\n")
	fmt.Printf("结果: %t\n", result5 == 0)
	fmt.Println()

	// 额外测试用例 - 溢出情况
	s6 := "2147483648"
	result6 := myAtoi(s6)
	fmt.Printf("溢出测试: s = \"%s\"\n", s6)
	fmt.Printf("输出: %d\n", result6)
	fmt.Printf("期望: 2147483647 (MaxInt32)\n")
	fmt.Printf("结果: %t\n", result6 == math.MaxInt32)
	fmt.Println()

	// 测试优化版本
	fmt.Println("=== 优化版本测试 ===")
	result1Opt := myAtoiOptimized(s1)
	result2Opt := myAtoiOptimized(s2)
	fmt.Printf("优化版本示例1: %d\n", result1Opt)
	fmt.Printf("优化版本示例2: %d\n", result2Opt)
	fmt.Printf("结果一致: %t\n", result1Opt == result1 && result2Opt == result2)
	fmt.Println()

	// 测试正则版本
	fmt.Println("=== 正则版本测试 ===")
	result1Regex := myAtoiRegex(s1)
	result2Regex := myAtoiRegex(s2)
	fmt.Printf("正则版本示例1: %d\n", result1Regex)
	fmt.Printf("正则版本示例2: %d\n", result2Regex)
	fmt.Printf("结果一致: %t\n", result1Regex == result1 && result2Regex == result2)
	fmt.Println()

	// 测试位运算版本
	fmt.Println("=== 位运算版本测试 ===")
	result1Bit := myAtoiBitwise(s1)
	result2Bit := myAtoiBitwise(s2)
	fmt.Printf("位运算版本示例1: %d\n", result1Bit)
	fmt.Printf("位运算版本示例2: %d\n", result2Bit)
	fmt.Printf("结果一致: %t\n", result1Bit == result1 && result2Bit == result2)
	fmt.Println()

	// 边界值测试
	fmt.Println("=== 边界值测试 ===")
	boundaryTests := []string{
		"",            // 空字符串
		"   ",         // 只有空格
		"+",           // 只有加号
		"-",           // 只有减号
		"++123",       // 多个加号
		"--123",       // 多个减号
		"123abc",      // 数字后跟字母
		"abc123",      // 字母开头
		"  000123",    // 前导零
		"  -000123",   // 负号加前导零
		"2147483647",  // 最大32位整数
		"-2147483648", // 最小32位整数
		"2147483648",  // 超过最大值
		"-2147483649", // 小于最小值
	}

	for _, test := range boundaryTests {
		result := myAtoi(test)
		fmt.Printf("s = \"%s\", result = %d\n", test, result)
	}
}
