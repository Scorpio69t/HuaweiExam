package main

import (
	"fmt"
)

// romanToInt 罗马数字转整数 - 从左到右遍历法
// 时间复杂度: O(n)，其中n是罗马数字字符串的长度
// 空间复杂度: O(1)
func romanToInt(s string) int {
	// 定义罗马数字到整数的映射
	romanMap := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	result := 0
	prevValue := 0

	// 从右到左遍历，处理减法规则
	for i := len(s) - 1; i >= 0; i-- {
		currentValue := romanMap[s[i]]

		// 如果当前值小于前一个值，需要减去当前值
		if currentValue < prevValue {
			result -= currentValue
		} else {
			result += currentValue
		}

		prevValue = currentValue
	}

	return result
}

// romanToIntLeftToRight 从左到右遍历法 - 另一种实现
// 时间复杂度: O(n)
// 空间复杂度: O(1)
func romanToIntLeftToRight(s string) int {
	// 定义罗马数字到整数的映射
	romanMap := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	result := 0

	// 从左到右遍历
	for i := 0; i < len(s); i++ {
		currentValue := romanMap[s[i]]

		// 检查是否需要减法
		if i+1 < len(s) && currentValue < romanMap[s[i+1]] {
			result -= currentValue
		} else {
			result += currentValue
		}
	}

	return result
}

// romanToIntOptimized 优化版本 - 使用数组映射
// 时间复杂度: O(n)
// 空间复杂度: O(1)
func romanToIntOptimized(s string) int {
	// 使用数组映射，避免map查找开销
	values := [128]int{} // ASCII字符映射
	values['I'] = 1
	values['V'] = 5
	values['X'] = 10
	values['L'] = 50
	values['C'] = 100
	values['D'] = 500
	values['M'] = 1000

	result := 0
	prevValue := 0

	// 从右到左遍历
	for i := len(s) - 1; i >= 0; i-- {
		currentValue := values[s[i]]

		if currentValue < prevValue {
			result -= currentValue
		} else {
			result += currentValue
		}

		prevValue = currentValue
	}

	return result
}

// romanToIntRecursive 递归方法 - 分治思想
// 时间复杂度: O(n)
// 空间复杂度: O(n)，递归调用栈
func romanToIntRecursive(s string) int {
	if len(s) == 0 {
		return 0
	}

	if len(s) == 1 {
		return getRomanValue(s[0])
	}

	// 检查前两个字符是否构成减法组合
	if len(s) >= 2 {
		first := getRomanValue(s[0])
		second := getRomanValue(s[1])

		if first < second {
			// 减法情况，如IV、IX、XL、XC、CD、CM
			return second - first + romanToIntRecursive(s[2:])
		}
	}

	// 普通情况，直接相加
	return getRomanValue(s[0]) + romanToIntRecursive(s[1:])
}

// getRomanValue 获取单个罗马数字的值
func getRomanValue(c byte) int {
	switch c {
	case 'I':
		return 1
	case 'V':
		return 5
	case 'X':
		return 10
	case 'L':
		return 50
	case 'C':
		return 100
	case 'D':
		return 500
	case 'M':
		return 1000
	default:
		return 0
	}
}

// romanToIntBitwise 位运算优化版本
// 时间复杂度: O(n)
// 空间复杂度: O(1)
func romanToIntBitwise(s string) int {
	result := 0
	prevValue := 0

	// 从右到左遍历，使用位运算优化
	for i := len(s) - 1; i >= 0; i-- {
		currentValue := getRomanValueBitwise(s[i])

		// 使用位运算判断大小关系
		if (currentValue & 0x7FFFFFFF) < (prevValue & 0x7FFFFFFF) {
			result -= currentValue
		} else {
			result += currentValue
		}

		prevValue = currentValue
	}

	return result
}

// getRomanValueBitwise 使用位运算获取罗马数字值
func getRomanValueBitwise(c byte) int {
	// 使用位运算优化switch语句
	value := 0
	switch c {
	case 'I':
		value = 1
	case 'V':
		value = 5
	case 'X':
		value = 10
	case 'L':
		value = 50
	case 'C':
		value = 100
	case 'D':
		value = 500
	case 'M':
		value = 1000
	}
	return value
}

func main() {
	// 测试用例1
	s1 := "III"
	result1 := romanToInt(s1)
	fmt.Printf("示例1: s = \"%s\"\n", s1)
	fmt.Printf("输出: %d\n", result1)
	fmt.Printf("期望: 3\n")
	fmt.Printf("结果: %t\n", result1 == 3)
	fmt.Println()

	// 测试用例2
	s2 := "IV"
	result2 := romanToInt(s2)
	fmt.Printf("示例2: s = \"%s\"\n", s2)
	fmt.Printf("输出: %d\n", result2)
	fmt.Printf("期望: 4\n")
	fmt.Printf("结果: %t\n", result2 == 4)
	fmt.Println()

	// 测试用例3
	s3 := "IX"
	result3 := romanToInt(s3)
	fmt.Printf("示例3: s = \"%s\"\n", s3)
	fmt.Printf("输出: %d\n", result3)
	fmt.Printf("期望: 9\n")
	fmt.Printf("结果: %t\n", result3 == 9)
	fmt.Println()

	// 测试用例4
	s4 := "LVIII"
	result4 := romanToInt(s4)
	fmt.Printf("示例4: s = \"%s\"\n", s4)
	fmt.Printf("输出: %d\n", result4)
	fmt.Printf("期望: 58\n")
	fmt.Printf("结果: %t\n", result4 == 58)
	fmt.Println()

	// 测试用例5
	s5 := "MCMXCIV"
	result5 := romanToInt(s5)
	fmt.Printf("示例5: s = \"%s\"\n", s5)
	fmt.Printf("输出: %d\n", result5)
	fmt.Printf("期望: 1994\n")
	fmt.Printf("结果: %t\n", result5 == 1994)
	fmt.Println()

	// 额外测试用例
	s6 := "MMMDCCXLIX"
	result6 := romanToInt(s6)
	fmt.Printf("额外测试: s = \"%s\"\n", s6)
	fmt.Printf("输出: %d\n", result6)
	fmt.Printf("期望: 3749\n")
	fmt.Printf("结果: %t\n", result6 == 3749)
	fmt.Println()

	// 测试从左到右版本
	fmt.Println("=== 从左到右版本测试 ===")
	result1LTR := romanToIntLeftToRight(s1)
	result2LTR := romanToIntLeftToRight(s2)
	fmt.Printf("从左到右版本示例1: %d\n", result1LTR)
	fmt.Printf("从左到右版本示例2: %d\n", result2LTR)
	fmt.Printf("结果一致: %t\n", result1LTR == result1 && result2LTR == result2)
	fmt.Println()

	// 测试优化版本
	fmt.Println("=== 优化版本测试 ===")
	result1Opt := romanToIntOptimized(s1)
	result2Opt := romanToIntOptimized(s2)
	fmt.Printf("优化版本示例1: %d\n", result1Opt)
	fmt.Printf("优化版本示例2: %d\n", result2Opt)
	fmt.Printf("结果一致: %t\n", result1Opt == result1 && result2Opt == result2)
	fmt.Println()

	// 测试递归版本
	fmt.Println("=== 递归版本测试 ===")
	result1Rec := romanToIntRecursive(s1)
	result2Rec := romanToIntRecursive(s2)
	fmt.Printf("递归版本示例1: %d\n", result1Rec)
	fmt.Printf("递归版本示例2: %d\n", result2Rec)
	fmt.Printf("结果一致: %t\n", result1Rec == result1 && result2Rec == result2)
	fmt.Println()

	// 测试位运算版本
	fmt.Println("=== 位运算版本测试 ===")
	result1Bit := romanToIntBitwise(s1)
	result2Bit := romanToIntBitwise(s2)
	fmt.Printf("位运算版本示例1: %d\n", result1Bit)
	fmt.Printf("位运算版本示例2: %d\n", result2Bit)
	fmt.Printf("结果一致: %t\n", result1Bit == result1 && result2Bit == result2)
	fmt.Println()

	// 边界值测试
	fmt.Println("=== 边界值测试 ===")
	boundaryTests := []string{
		"I",         // 最小值
		"MMMCMXCIX", // 最大值
		"IV",        // 减法规则
		"IX",        // 减法规则
		"XL",        // 减法规则
		"XC",        // 减法规则
		"CD",        // 减法规则
		"CM",        // 减法规则
		"XII",       // 普通加法
		"XXVII",     // 普通加法
	}

	for _, test := range boundaryTests {
		result := romanToInt(test)
		fmt.Printf("s = \"%s\", result = %d\n", test, result)
	}
}
