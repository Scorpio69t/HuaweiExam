package main

import (
	"fmt"
	"strings"
)

// intToRoman 整数转罗马数字 - 贪心算法
// 时间复杂度: O(1)，因为罗马数字有固定的最大值
// 空间复杂度: O(1)
func intToRoman(num int) string {
	// 定义罗马数字的符号和对应的值，按从大到小排序
	values := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	symbols := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}

	var result strings.Builder

	// 贪心算法：每次选择最大的可能值
	for i := 0; i < len(values) && num > 0; i++ {
		// 当当前值小于等于剩余数字时，重复添加对应的符号
		for num >= values[i] {
			result.WriteString(symbols[i])
			num -= values[i]
		}
	}

	return result.String()
}

// intToRomanOptimized 优化版本 - 使用数组映射
// 时间复杂度: O(1)
// 空间复杂度: O(1)
func intToRomanOptimized(num int) string {
	// 使用数组映射，避免循环查找
	thousands := []string{"", "M", "MM", "MMM"}
	hundreds := []string{"", "C", "CC", "CCC", "CD", "D", "DC", "DCC", "DCCC", "CM"}
	tens := []string{"", "X", "XX", "XXX", "XL", "L", "LX", "LXX", "LXXX", "XC"}
	ones := []string{"", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX"}

	return thousands[num/1000] + hundreds[(num%1000)/100] + tens[(num%100)/10] + ones[num%10]
}

// intToRomanRecursive 递归方法 - 分治思想
// 时间复杂度: O(log num)
// 空间复杂度: O(log num)
func intToRomanRecursive(num int) string {
	if num == 0 {
		return ""
	}

	// 定义罗马数字映射
	romanMap := map[int]string{
		1:    "I",
		4:    "IV",
		5:    "V",
		9:    "IX",
		10:   "X",
		40:   "XL",
		50:   "L",
		90:   "XC",
		100:  "C",
		400:  "CD",
		500:  "D",
		900:  "CM",
		1000: "M",
	}

	// 找到小于等于num的最大罗马数字值
	maxValue := 0
	for value := range romanMap {
		if value <= num && value > maxValue {
			maxValue = value
		}
	}

	// 递归处理剩余部分
	return romanMap[maxValue] + intToRomanRecursive(num-maxValue)
}

// intToRomanBitwise 位运算优化版本
// 时间复杂度: O(1)
// 空间复杂度: O(1)
func intToRomanBitwise(num int) string {
	var result strings.Builder

	// 处理千位
	for num >= 1000 {
		result.WriteString("M")
		num -= 1000
	}

	// 处理百位
	if num >= 900 {
		result.WriteString("CM")
		num -= 900
	} else if num >= 500 {
		result.WriteString("D")
		num -= 500
		for num >= 100 {
			result.WriteString("C")
			num -= 100
		}
	} else if num >= 400 {
		result.WriteString("CD")
		num -= 400
	} else {
		for num >= 100 {
			result.WriteString("C")
			num -= 100
		}
	}

	// 处理十位
	if num >= 90 {
		result.WriteString("XC")
		num -= 90
	} else if num >= 50 {
		result.WriteString("L")
		num -= 50
		for num >= 10 {
			result.WriteString("X")
			num -= 10
		}
	} else if num >= 40 {
		result.WriteString("XL")
		num -= 40
	} else {
		for num >= 10 {
			result.WriteString("X")
			num -= 10
		}
	}

	// 处理个位
	if num >= 9 {
		result.WriteString("IX")
		num -= 9
	} else if num >= 5 {
		result.WriteString("V")
		num -= 5
		for num >= 1 {
			result.WriteString("I")
			num -= 1
		}
	} else if num >= 4 {
		result.WriteString("IV")
		num -= 4
	} else {
		for num >= 1 {
			result.WriteString("I")
			num -= 1
		}
	}

	return result.String()
}

// intToRomanTable 查表法 - 最直观的实现
// 时间复杂度: O(1)
// 空间复杂度: O(1)
func intToRomanTable(num int) string {
	// 预定义所有可能的罗马数字组合
	romanTable := []struct {
		value  int
		symbol string
	}{
		{1000, "M"},
		{900, "CM"},
		{500, "D"},
		{400, "CD"},
		{100, "C"},
		{90, "XC"},
		{50, "L"},
		{40, "XL"},
		{10, "X"},
		{9, "IX"},
		{5, "V"},
		{4, "IV"},
		{1, "I"},
	}

	var result strings.Builder

	// 遍历表，构建结果
	for _, item := range romanTable {
		for num >= item.value {
			result.WriteString(item.symbol)
			num -= item.value
		}
	}

	return result.String()
}

func main() {
	// 测试用例1
	num1 := 3749
	result1 := intToRoman(num1)
	fmt.Printf("示例1: num = %d\n", num1)
	fmt.Printf("输出: \"%s\"\n", result1)
	fmt.Printf("期望: \"MMMDCCXLIX\"\n")
	fmt.Printf("结果: %t\n", result1 == "MMMDCCXLIX")
	fmt.Println()

	// 测试用例2
	num2 := 58
	result2 := intToRoman(num2)
	fmt.Printf("示例2: num = %d\n", num2)
	fmt.Printf("输出: \"%s\"\n", result2)
	fmt.Printf("期望: \"LVIII\"\n")
	fmt.Printf("结果: %t\n", result2 == "LVIII")
	fmt.Println()

	// 测试用例3
	num3 := 1994
	result3 := intToRoman(num3)
	fmt.Printf("示例3: num = %d\n", num3)
	fmt.Printf("输出: \"%s\"\n", result3)
	fmt.Printf("期望: \"MCMXCIV\"\n")
	fmt.Printf("结果: %t\n", result3 == "MCMXCIV")
	fmt.Println()

	// 额外测试用例
	num4 := 3999
	result4 := intToRoman(num4)
	fmt.Printf("额外测试: num = %d\n", num4)
	fmt.Printf("输出: \"%s\"\n", result4)
	fmt.Printf("期望: \"MMMCMXCIX\"\n")
	fmt.Printf("结果: %t\n", result4 == "MMMCMXCIX")
	fmt.Println()

	num5 := 1
	result5 := intToRoman(num5)
	fmt.Printf("额外测试: num = %d\n", num5)
	fmt.Printf("输出: \"%s\"\n", result5)
	fmt.Printf("期望: \"I\"\n")
	fmt.Printf("结果: %t\n", result5 == "I")
	fmt.Println()

	// 测试优化版本
	fmt.Println("=== 优化版本测试 ===")
	result1Opt := intToRomanOptimized(num1)
	result2Opt := intToRomanOptimized(num2)
	fmt.Printf("优化版本示例1: %s\n", result1Opt)
	fmt.Printf("优化版本示例2: %s\n", result2Opt)
	fmt.Printf("结果一致: %t\n", result1Opt == result1 && result2Opt == result2)
	fmt.Println()

	// 测试递归版本
	fmt.Println("=== 递归版本测试 ===")
	result1Rec := intToRomanRecursive(num1)
	result2Rec := intToRomanRecursive(num2)
	fmt.Printf("递归版本示例1: %s\n", result1Rec)
	fmt.Printf("递归版本示例2: %s\n", result2Rec)
	fmt.Printf("结果一致: %t\n", result1Rec == result1 && result2Rec == result2)
	fmt.Println()

	// 测试位运算版本
	fmt.Println("=== 位运算版本测试 ===")
	result1Bit := intToRomanBitwise(num1)
	result2Bit := intToRomanBitwise(num2)
	fmt.Printf("位运算版本示例1: %s\n", result1Bit)
	fmt.Printf("位运算版本示例2: %s\n", result2Bit)
	fmt.Printf("结果一致: %t\n", result1Bit == result1 && result2Bit == result2)
	fmt.Println()

	// 测试查表版本
	fmt.Println("=== 查表版本测试 ===")
	result1Tab := intToRomanTable(num1)
	result2Tab := intToRomanTable(num2)
	fmt.Printf("查表版本示例1: %s\n", result1Tab)
	fmt.Printf("查表版本示例2: %s\n", result2Tab)
	fmt.Printf("结果一致: %t\n", result1Tab == result1 && result2Tab == result2)
	fmt.Println()

	// 边界值测试
	fmt.Println("=== 边界值测试 ===")
	boundaryTests := []int{
		1,    // 最小值
		4,    // 减法形式
		9,    // 减法形式
		40,   // 减法形式
		90,   // 减法形式
		400,  // 减法形式
		900,  // 减法形式
		1000, // 基本符号
		3999, // 最大值
	}

	for _, test := range boundaryTests {
		result := intToRoman(test)
		fmt.Printf("num = %d, roman = %s\n", test, result)
	}
}
