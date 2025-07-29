package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

// 解法一：模拟乘法竖式（推荐解法）
// 时间复杂度：O(m×n)，空间复杂度：O(m+n)
func multiply(num1 string, num2 string) string {
	// 特殊情况：任一数字为0
	if num1 == "0" || num2 == "0" {
		return "0"
	}

	m, n := len(num1), len(num2)
	// 结果最多有 m+n 位
	result := make([]int, m+n)

	// 从右到左遍历两个数字
	for i := m - 1; i >= 0; i-- {
		digit1 := int(num1[i] - '0')
		for j := n - 1; j >= 0; j-- {
			digit2 := int(num2[j] - '0')

			// 计算乘积并加到对应位置
			product := digit1 * digit2
			posLow := i + j + 1 // 个位
			posHigh := i + j    // 十位

			// 累加到当前位
			sum := product + result[posLow]
			result[posLow] = sum % 10   // 保留个位
			result[posHigh] += sum / 10 // 进位到高位
		}
	}

	// 转换为字符串，跳过前导零
	var sb strings.Builder
	start := 0
	for start < len(result) && result[start] == 0 {
		start++
	}

	for i := start; i < len(result); i++ {
		sb.WriteByte(byte(result[i] + '0'))
	}

	return sb.String()
}

// 解法二：优化的模拟乘法（预处理优化）
// 时间复杂度：O(m×n)，空间复杂度：O(m+n)
func multiplyOptimized(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}

	m, n := len(num1), len(num2)
	result := make([]int, m+n)

	// 预转换字符为数字，减少重复转换
	digits1 := make([]int, m)
	digits2 := make([]int, n)

	for i := 0; i < m; i++ {
		digits1[i] = int(num1[i] - '0')
	}
	for i := 0; i < n; i++ {
		digits2[i] = int(num2[i] - '0')
	}

	// 执行乘法
	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			product := digits1[i] * digits2[j]
			posLow := i + j + 1
			posHigh := i + j

			sum := product + result[posLow]
			result[posLow] = sum % 10
			result[posHigh] += sum / 10
		}
	}

	// 构建结果字符串
	var sb strings.Builder
	sb.Grow(m + n) // 预分配容量

	start := 0
	for start < len(result) && result[start] == 0 {
		start++
	}

	for i := start; i < len(result); i++ {
		sb.WriteByte(byte(result[i] + '0'))
	}

	return sb.String()
}

// 解法三：字符串加法组合
// 时间复杂度：O(m×n²)，空间复杂度：O(m×n)
func multiplyByAddition(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}

	result := "0"

	// 逐位相乘并累加
	for i := len(num2) - 1; i >= 0; i-- {
		digit := int(num2[i] - '0')

		// 计算 num1 × digit
		temp := multiplyByDigit(num1, digit)

		// 添加相应的零（位数偏移）
		zeros := len(num2) - 1 - i
		for j := 0; j < zeros; j++ {
			temp += "0"
		}

		// 累加到结果
		result = addStrings(result, temp)
	}

	return result
}

// 辅助函数：数字字符串乘以单个数字
func multiplyByDigit(num string, digit int) string {
	if digit == 0 {
		return "0"
	}

	result := make([]int, len(num)+1)
	carry := 0

	for i := len(num) - 1; i >= 0; i-- {
		product := int(num[i]-'0')*digit + carry
		result[i+1] = product % 10
		carry = product / 10
	}
	result[0] = carry

	// 转换为字符串
	var sb strings.Builder
	start := 0
	if result[0] == 0 {
		start = 1
	}

	for i := start; i < len(result); i++ {
		sb.WriteByte(byte(result[i] + '0'))
	}

	return sb.String()
}

// 辅助函数：字符串加法
func addStrings(num1 string, num2 string) string {
	i, j := len(num1)-1, len(num2)-1
	carry := 0
	var result strings.Builder

	for i >= 0 || j >= 0 || carry > 0 {
		sum := carry
		if i >= 0 {
			sum += int(num1[i] - '0')
			i--
		}
		if j >= 0 {
			sum += int(num2[j] - '0')
			j--
		}

		result.WriteByte(byte(sum%10 + '0'))
		carry = sum / 10
	}

	// 反转结果
	return reverseString(result.String())
}

// 辅助函数：反转字符串
func reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// 解法四：Karatsuba分治算法（适用于超大数）
// 时间复杂度：O(n^1.585)，空间复杂度：O(log n)
func multiplyKaratsuba(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}

	// 小数字直接计算
	if len(num1) <= 10 && len(num2) <= 10 {
		return multiply(num1, num2)
	}

	// 补齐到相同长度
	maxLen := max(len(num1), len(num2))
	if maxLen%2 == 1 {
		maxLen++
	}

	num1 = padLeft(num1, maxLen)
	num2 = padLeft(num2, maxLen)

	return karatsuba(num1, num2)
}

func karatsuba(x, y string) string {
	n := len(x)
	if n <= 10 {
		return multiply(x, y)
	}

	m := n / 2

	// 分割
	x1 := x[:n-m] // 高位
	x0 := x[n-m:] // 低位
	y1 := y[:n-m] // 高位
	y0 := y[n-m:] // 低位

	// 递归计算
	z0 := karatsuba(x0, y0)
	z2 := karatsuba(x1, y1)

	// 计算 (x1 + x0) * (y1 + y0)
	x1PlusX0 := addStrings(x1, x0)
	y1PlusY0 := addStrings(y1, y0)
	z1Temp := karatsuba(x1PlusX0, y1PlusY0)

	// z1 = z1Temp - z2 - z0
	z1 := subtractStrings(z1Temp, addStrings(z2, z0))

	// 组合结果: z2 * 10^(2m) + z1 * 10^m + z0
	result := addStrings(addStrings(
		multiplyByPowerOf10(z2, 2*m),
		multiplyByPowerOf10(z1, m)),
		z0)

	return removeLeadingZeros(result)
}

// 辅助函数：左填充零
func padLeft(s string, length int) string {
	if len(s) >= length {
		return s
	}
	return strings.Repeat("0", length-len(s)) + s
}

// 辅助函数：乘以10的幂
func multiplyByPowerOf10(num string, power int) string {
	if num == "0" || power == 0 {
		return num
	}
	return num + strings.Repeat("0", power)
}

// 辅助函数：字符串减法（假设num1 >= num2）
func subtractStrings(num1, num2 string) string {
	if compareStrings(num1, num2) < 0 {
		return "0" // 简化处理，实际应该处理负数
	}

	i, j := len(num1)-1, len(num2)-1
	borrow := 0
	var result strings.Builder

	for i >= 0 {
		diff := int(num1[i]-'0') - borrow
		if j >= 0 {
			diff -= int(num2[j] - '0')
			j--
		}

		if diff < 0 {
			diff += 10
			borrow = 1
		} else {
			borrow = 0
		}

		result.WriteByte(byte(diff + '0'))
		i--
	}

	return removeLeadingZeros(reverseString(result.String()))
}

// 辅助函数：比较两个数字字符串
func compareStrings(num1, num2 string) int {
	if len(num1) != len(num2) {
		return len(num1) - len(num2)
	}
	return strings.Compare(num1, num2)
}

// 辅助函数：移除前导零
func removeLeadingZeros(s string) string {
	i := 0
	for i < len(s) && s[i] == '0' {
		i++
	}
	if i == len(s) {
		return "0"
	}
	return s[i:]
}

// 辅助函数：取最大值
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 解法五：直接转换法（受整数范围限制，仅用于演示）
func multiplyDirect(num1 string, num2 string) string {
	// 仅适用于小数字，受int64范围限制
	if len(num1) > 18 || len(num2) > 18 {
		return multiply(num1, num2) // 回退到模拟算法
	}

	n1, _ := strconv.ParseInt(num1, 10, 64)
	n2, _ := strconv.ParseInt(num2, 10, 64)

	return strconv.FormatInt(n1*n2, 10)
}

// 测试函数
func testMultiply() {
	testCases := []struct {
		num1     string
		num2     string
		expected string
		desc     string
	}{
		{"2", "3", "6", "简单单位数相乘"},
		{"123", "456", "56088", "三位数相乘"},
		{"0", "123", "0", "零乘法"},
		{"999", "999", "998001", "最大单位数相乘"},
		{"12", "34", "408", "两位数相乘"},
		{"100", "200", "20000", "整百数相乘"},
		{"11", "11", "121", "相同数字相乘"},
		{"1", "123456789", "123456789", "乘以1"},
		{"999999999", "999999999", "999999998000000001", "九位数相乘"},
		{"12345", "67890", "838102050", "五位数相乘"},
		{"1000000", "2000000", "2000000000000", "大整数相乘"},
		{"98765", "43210", "4267635650", "随机数字"},
		{"777", "888", "689976", "重复数字"},
		{"505", "404", "204020", "包含零的数字"},
		{"99999", "99999", "9999800001", "五个9相乘"},
	}

	fmt.Println("=== 字符串相乘测试 ===\n")

	for i, tc := range testCases {
		// 测试主要解法
		result1 := multiply(tc.num1, tc.num2)
		result2 := multiplyOptimized(tc.num1, tc.num2)
		result3 := multiplyByAddition(tc.num1, tc.num2)

		status := "✅"
		if result1 != tc.expected {
			status = "❌"
		}

		fmt.Printf("测试 %d: %s\n", i+1, tc.desc)
		fmt.Printf("输入: \"%s\" × \"%s\"\n", tc.num1, tc.num2)
		fmt.Printf("期望: %s\n", tc.expected)
		fmt.Printf("模拟法: %s\n", result1)
		fmt.Printf("优化法: %s\n", result2)
		fmt.Printf("加法法: %s\n", result3)
		fmt.Printf("结果: %s\n", status)
		fmt.Println(strings.Repeat("-", 40))
	}
}

// 性能测试
func benchmarkMultiply() {
	fmt.Println("\n=== 性能测试 ===\n")

	// 构造测试数据
	testData := []struct {
		num1, num2 string
		desc       string
	}{
		{generateNumber(10), generateNumber(10), "10位×10位"},
		{generateNumber(50), generateNumber(50), "50位×50位"},
		{generateNumber(100), generateNumber(100), "100位×100位"},
		{generateNumber(200), generateNumber(200), "200位×200位"},
	}

	algorithms := []struct {
		name string
		fn   func(string, string) string
	}{
		{"模拟乘法", multiply},
		{"优化乘法", multiplyOptimized},
		{"Karatsuba", multiplyKaratsuba},
		{"加法组合", multiplyByAddition},
	}

	for _, data := range testData {
		fmt.Printf("%s:\n", data.desc)

		for _, algo := range algorithms {
			start := time.Now()
			result := algo.fn(data.num1, data.num2)
			duration := time.Since(start)

			fmt.Printf("  %s: 结果长度%d, 耗时: %v\n",
				algo.name, len(result), duration)
		}
		fmt.Println()
	}
}

// 生成指定长度的随机数字字符串
func generateNumber(length int) string {
	if length <= 0 {
		return "0"
	}

	var sb strings.Builder
	sb.Grow(length)

	// 首位不能为0
	sb.WriteByte(byte('1' + length%9))

	// 后续位可以是任意数字
	for i := 1; i < length; i++ {
		sb.WriteByte(byte('0' + (i*7+3)%10))
	}

	return sb.String()
}

// 演示手算验证
func demonstrateManualCalculation() {
	fmt.Println("\n=== 手算过程演示 ===")
	fmt.Println("计算 123 × 456:")
	fmt.Println()
	fmt.Println("    123")
	fmt.Println("  × 456")
	fmt.Println("  -----")
	fmt.Println("    738  (123 × 6)")
	fmt.Println("   6150  (123 × 50)")
	fmt.Println("  49200  (123 × 400)")
	fmt.Println("  -----")
	fmt.Println("  56088")
	fmt.Println()

	result := multiply("123", "456")
	fmt.Printf("算法结果: %s\n", result)
	fmt.Printf("验证: %s\n", map[bool]string{true: "✅ 正确", false: "❌ 错误"}[result == "56088"])
}

func main() {
	fmt.Println("43. 字符串相乘 - 多种解法实现")
	fmt.Println("========================================")

	// 基础功能测试
	testMultiply()

	// 性能对比测试
	benchmarkMultiply()

	// 手算演示
	demonstrateManualCalculation()

	// 展示算法特点
	fmt.Println("\n=== 算法特点分析 ===")
	fmt.Println("1. 模拟乘法：经典解法，时间O(m×n)，推荐日常使用")
	fmt.Println("2. 优化乘法：预处理优化，减少重复计算，性能略优")
	fmt.Println("3. Karatsuba：分治算法，时间O(n^1.585)，适合超大数")
	fmt.Println("4. 加法组合：思路直观，但效率较低O(m×n²)")
	fmt.Println("5. 直接转换：受整数范围限制，仅适用小数字")

	fmt.Println("\n=== 关键技巧总结 ===")
	fmt.Println("• 位置计算：num1[i]×num2[j] 结果放在 i+j+1 位置")
	fmt.Println("• 进位处理：先累加同位置乘积，再统一处理进位")
	fmt.Println("• 零值优化：任一数字为0时提前返回")
	fmt.Println("• 前导零处理：结果构造时跳过前导零")
	fmt.Println("• 容量预分配：StringBuilder预分配避免扩容")
}
