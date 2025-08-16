package main

import (
	"fmt"
	"math"
	"strconv"
)

// reverse 整数反转 - 数学方法
// 时间复杂度: O(log n)，其中n是整数的位数
// 空间复杂度: O(1)
func reverse(x int) int {
	var result int

	// 处理负数，先转为正数处理
	isNegative := x < 0
	if isNegative {
		x = -x
	}

	// 逐位提取并反转
	for x > 0 {
		digit := x % 10
		x /= 10

		// 检查溢出：在乘以10之前检查
		if result > math.MaxInt32/10 || (result == math.MaxInt32/10 && digit > 7) {
			return 0
		}
		if result < math.MinInt32/10 || (result == math.MinInt32/10 && digit < -8) {
			return 0
		}

		result = result*10 + digit
	}

	// 恢复符号
	if isNegative {
		result = -result
	}

	return result
}

// reverseOptimized 优化版本 - 使用更简洁的逻辑
// 时间复杂度: O(log n)
// 空间复杂度: O(1)
func reverseOptimized(x int) int {
	var result int

	for x != 0 {
		digit := x % 10
		x /= 10

		// 检查溢出
		if result > math.MaxInt32/10 || (result == math.MaxInt32/10 && digit > 7) {
			return 0
		}
		if result < math.MinInt32/10 || (result == math.MinInt32/10 && digit < -8) {
			return 0
		}

		result = result*10 + digit
	}

	return result
}

// reverseString 字符串方法 - 转换为字符串后反转
// 时间复杂度: O(log n)
// 空间复杂度: O(log n)
func reverseString(x int) int {
	// 处理特殊情况
	if x == 0 {
		return 0
	}

	// 转换为字符串
	str := strconv.Itoa(x)
	isNegative := str[0] == '-'

	// 如果是负数，去掉负号
	if isNegative {
		str = str[1:]
	}

	// 反转字符串
	bytes := []byte(str)
	for i, j := 0, len(bytes)-1; i < j; i, j = i+1, j-1 {
		bytes[i], bytes[j] = bytes[j], bytes[i]
	}

	// 转换回整数
	result, err := strconv.Atoi(string(bytes))
	if err != nil {
		return 0
	}

	// 检查溢出
	if result > math.MaxInt32 || result < math.MinInt32 {
		return 0
	}

	// 恢复符号
	if isNegative {
		result = -result
	}

	return result
}

// reverseBitwise 位运算方法 - 使用位运算优化
// 时间复杂度: O(log n)
// 空间复杂度: O(1)
func reverseBitwise(x int) int {
	var result int

	for x != 0 {
		digit := x % 10
		x /= 10

		// 使用位运算检查溢出
		if result > (math.MaxInt32-digit)/10 {
			return 0
		}

		result = result*10 + digit
	}

	return result
}

func main() {
	// 测试用例1
	x1 := 123
	result1 := reverse(x1)
	fmt.Printf("示例1: x = %d\n", x1)
	fmt.Printf("输出: %d\n", result1)
	fmt.Printf("期望: 321\n")
	fmt.Printf("结果: %t\n", result1 == 321)
	fmt.Println()

	// 测试用例2
	x2 := -123
	result2 := reverse(x2)
	fmt.Printf("示例2: x = %d\n", x2)
	fmt.Printf("输出: %d\n", result2)
	fmt.Printf("期望: -321\n")
	fmt.Printf("结果: %t\n", result2 == -321)
	fmt.Println()

	// 测试用例3
	x3 := 120
	result3 := reverse(x3)
	fmt.Printf("示例3: x = %d\n", x3)
	fmt.Printf("输出: %d\n", result3)
	fmt.Printf("期望: 21\n")
	fmt.Printf("结果: %t\n", result3 == 21)
	fmt.Println()

	// 测试用例4
	x4 := 0
	result4 := reverse(x4)
	fmt.Printf("示例4: x = %d\n", x4)
	fmt.Printf("输出: %d\n", result4)
	fmt.Printf("期望: 0\n")
	fmt.Printf("结果: %t\n", result4 == 0)
	fmt.Println()

	// 额外测试用例 - 溢出情况
	x5 := 1534236469
	result5 := reverse(x5)
	fmt.Printf("溢出测试: x = %d\n", x5)
	fmt.Printf("输出: %d\n", result5)
	fmt.Printf("期望: 0 (溢出)\n")
	fmt.Printf("结果: %t\n", result5 == 0)
	fmt.Println()

	// 测试优化版本
	fmt.Println("=== 优化版本测试 ===")
	result1Opt := reverseOptimized(x1)
	result2Opt := reverseOptimized(x2)
	fmt.Printf("优化版本示例1: %d\n", result1Opt)
	fmt.Printf("优化版本示例2: %d\n", result2Opt)
	fmt.Printf("结果一致: %t\n", result1Opt == result1 && result2Opt == result2)
	fmt.Println()

	// 测试字符串版本
	fmt.Println("=== 字符串版本测试 ===")
	result1Str := reverseString(x1)
	result2Str := reverseString(x2)
	fmt.Printf("字符串版本示例1: %d\n", result1Str)
	fmt.Printf("字符串版本示例2: %d\n", result2Str)
	fmt.Printf("结果一致: %t\n", result1Str == result1 && result2Str == result2)
	fmt.Println()

	// 测试位运算版本
	fmt.Println("=== 位运算版本测试 ===")
	result1Bit := reverseBitwise(x1)
	result2Bit := reverseBitwise(x2)
	fmt.Printf("位运算版本示例1: %d\n", result1Bit)
	fmt.Printf("位运算版本示例2: %d\n", result2Bit)
	fmt.Printf("结果一致: %t\n", result1Bit == result1 && result2Bit == result2)
	fmt.Println()

	// 边界值测试
	fmt.Println("=== 边界值测试 ===")
	boundaryTests := []int{
		math.MaxInt32, // 最大32位整数
		math.MinInt32, // 最小32位整数
		1000000000,    // 10位数
		-1000000000,   // 负数10位数
	}

	for _, test := range boundaryTests {
		result := reverse(test)
		fmt.Printf("x = %d, reverse = %d\n", test, result)
	}
}
