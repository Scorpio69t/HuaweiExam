package main

import (
	"fmt"
	"strconv"
)

// isPalindrome 回文数判断 - 数学方法
// 时间复杂度: O(log n)，其中n是整数的位数
// 空间复杂度: O(1)
func isPalindrome(x int) bool {
	// 负数不可能是回文数
	if x < 0 {
		return false
	}

	// 0是回文数
	if x == 0 {
		return true
	}

	// 如果数字以0结尾，只有0本身是回文数
	if x%10 == 0 {
		return false
	}

	// 反转后半部分数字
	reversed := 0
	for x > reversed {
		reversed = reversed*10 + x%10
		x /= 10
	}

	// 当数字长度为奇数时，需要去掉中间的数字
	// 例如：12321，反转后得到123，原数变成12
	// 当数字长度为偶数时，直接比较
	// 例如：1221，反转后得到12，原数变成12
	return x == reversed || x == reversed/10
}

// isPalindromeString 字符串方法 - 转换为字符串后判断
// 时间复杂度: O(log n)
// 空间复杂度: O(log n)
func isPalindromeString(x int) bool {
	if x < 0 {
		return false
	}

	// 转换为字符串
	str := strconv.Itoa(x)

	// 双指针判断回文
	left, right := 0, len(str)-1
	for left < right {
		if str[left] != str[right] {
			return false
		}
		left++
		right--
	}

	return true
}

// isPalindromeFullReverse 完全反转法 - 反转整个数字后比较
// 时间复杂度: O(log n)
// 空间复杂度: O(1)
func isPalindromeFullReverse(x int) bool {
	if x < 0 {
		return false
	}

	// 保存原数
	original := x

	// 完全反转
	reversed := 0
	for x > 0 {
		reversed = reversed*10 + x%10
		x /= 10
	}

	return original == reversed
}

// isPalindromeBitwise 位运算优化版本
// 时间复杂度: O(log n)
// 空间复杂度: O(1)
func isPalindromeBitwise(x int) bool {
	if x < 0 {
		return false
	}

	if x == 0 {
		return true
	}

	// 计算数字的位数
	divisor := 1
	temp := x
	for temp >= 10 {
		divisor *= 10
		temp /= 10
	}

	// 逐位比较首尾数字
	for x > 0 {
		firstDigit := x / divisor
		lastDigit := x % 10

		if firstDigit != lastDigit {
			return false
		}

		// 去掉首尾数字
		x = (x % divisor) / 10
		divisor /= 100
	}

	return true
}

// isPalindromeOptimized 优化版本 - 只反转一半
// 时间复杂度: O(log n)
// 空间复杂度: O(1)
func isPalindromeOptimized(x int) bool {
	// 负数不可能是回文数
	if x < 0 {
		return false
	}

	// 0是回文数
	if x == 0 {
		return true
	}

	// 如果数字以0结尾，只有0本身是回文数
	if x%10 == 0 {
		return false
	}

	// 反转后半部分数字
	reversed := 0
	for x > reversed {
		reversed = reversed*10 + x%10
		x /= 10
	}

	// 当数字长度为奇数时，需要去掉中间的数字
	// 当数字长度为偶数时，直接比较
	return x == reversed || x == reversed/10
}

func main() {
	// 测试用例1
	x1 := 121
	result1 := isPalindrome(x1)
	fmt.Printf("示例1: x = %d\n", x1)
	fmt.Printf("输出: %t\n", result1)
	fmt.Printf("期望: true\n")
	fmt.Printf("结果: %t\n", result1 == true)
	fmt.Println()

	// 测试用例2
	x2 := -121
	result2 := isPalindrome(x2)
	fmt.Printf("示例2: x = %d\n", x2)
	fmt.Printf("输出: %t\n", result2)
	fmt.Printf("期望: false\n")
	fmt.Printf("结果: %t\n", result2 == false)
	fmt.Println()

	// 测试用例3
	x3 := 10
	result3 := isPalindrome(x3)
	fmt.Printf("示例3: x = %d\n", x3)
	fmt.Printf("输出: %t\n", result3)
	fmt.Printf("期望: false\n")
	fmt.Printf("结果: %t\n", result3 == false)
	fmt.Println()

	// 额外测试用例
	x4 := 0
	result4 := isPalindrome(x4)
	fmt.Printf("额外测试: x = %d\n", x4)
	fmt.Printf("输出: %t\n", result4)
	fmt.Printf("期望: true\n")
	fmt.Printf("结果: %t\n", result4 == true)
	fmt.Println()

	x5 := 12321
	result5 := isPalindrome(x5)
	fmt.Printf("额外测试: x = %d\n", x5)
	fmt.Printf("输出: %t\n", result5)
	fmt.Printf("期望: true\n")
	fmt.Printf("结果: %t\n", result5 == true)
	fmt.Println()

	// 测试优化版本
	fmt.Println("=== 优化版本测试 ===")
	result1Opt := isPalindromeOptimized(x1)
	result2Opt := isPalindromeOptimized(x2)
	fmt.Printf("优化版本示例1: %t\n", result1Opt)
	fmt.Printf("优化版本示例2: %t\n", result2Opt)
	fmt.Printf("结果一致: %t\n", result1Opt == result1 && result2Opt == result2)
	fmt.Println()

	// 测试字符串版本
	fmt.Println("=== 字符串版本测试 ===")
	result1Str := isPalindromeString(x1)
	result2Str := isPalindromeString(x2)
	fmt.Printf("字符串版本示例1: %t\n", result1Str)
	fmt.Printf("字符串版本示例2: %t\n", result2Str)
	fmt.Printf("结果一致: %t\n", result1Str == result1 && result2Str == result2)
	fmt.Println()

	// 测试完全反转版本
	fmt.Println("=== 完全反转版本测试 ===")
	result1Full := isPalindromeFullReverse(x1)
	result2Full := isPalindromeFullReverse(x2)
	fmt.Printf("完全反转版本示例1: %t\n", result1Full)
	fmt.Printf("完全反转版本示例2: %t\n", result2Full)
	fmt.Printf("结果一致: %t\n", result1Full == result1 && result2Full == result2)
	fmt.Println()

	// 测试位运算版本
	fmt.Println("=== 位运算版本测试 ===")
	result1Bit := isPalindromeBitwise(x1)
	result2Bit := isPalindromeBitwise(x2)
	fmt.Printf("位运算版本示例1: %t\n", result1Bit)
	fmt.Printf("位运算版本示例2: %t\n", result2Bit)
	fmt.Printf("结果一致: %t\n", result1Bit == result1 && result2Bit == result2)
	fmt.Println()

	// 边界值测试
	fmt.Println("=== 边界值测试 ===")
	boundaryTests := []int{
		0,     // 0
		1,     // 单位数
		11,    // 两位数回文
		12,    // 两位数非回文
		121,   // 三位数回文
		123,   // 三位数非回文
		1221,  // 四位数回文
		1234,  // 四位数非回文
		12321, // 五位数回文
		12345, // 五位数非回文
		1001,  // 以0结尾的回文
		1000,  // 以0结尾的非回文
	}

	for _, test := range boundaryTests {
		result := isPalindrome(test)
		fmt.Printf("x = %d, isPalindrome = %t\n", test, result)
	}
}
