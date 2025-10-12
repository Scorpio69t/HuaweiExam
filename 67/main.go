package main

import (
	"fmt"
	"strconv"
	"strings"
)

// =========================== 方法一：双指针+进位（最优解法） ===========================

// addBinary 双指针+进位+StringBuilder
// 时间复杂度：O(max(m,n))，m和n分别为两个字符串的长度
// 空间复杂度：O(1)，不计结果字符串
func addBinary(a string, b string) string {
	i, j := len(a)-1, len(b)-1
	carry := 0
	var result strings.Builder

	for i >= 0 || j >= 0 || carry > 0 {
		sum := carry

		if i >= 0 {
			sum += int(a[i] - '0')
			i--
		}

		if j >= 0 {
			sum += int(b[j] - '0')
			j--
		}

		result.WriteByte(byte('0' + sum%2))
		carry = sum / 2
	}

	// 反转结果
	res := result.String()
	return reverse(res)
}

// reverse 反转字符串
func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// =========================== 方法二：直接前插（简单但效率低） ===========================

// addBinary2 直接前插构建结果
// 时间复杂度：O(n²)，字符串前插每次O(n)
// 空间复杂度：O(1)，不计结果字符串
func addBinary2(a string, b string) string {
	i, j := len(a)-1, len(b)-1
	carry := 0
	result := ""

	for i >= 0 || j >= 0 || carry > 0 {
		sum := carry

		if i >= 0 {
			sum += int(a[i] - '0')
			i--
		}

		if j >= 0 {
			sum += int(b[j] - '0')
			j--
		}

		// 直接在前面插入
		result = string('0'+byte(sum%2)) + result
		carry = sum / 2
	}

	return result
}

// =========================== 方法三：递归实现 ===========================

// addBinary3 递归解法
// 时间复杂度：O(max(m,n))
// 空间复杂度：O(max(m,n))，递归栈空间
func addBinary3(a string, b string) string {
	return addHelper(a, b, len(a)-1, len(b)-1, 0)
}

// addHelper 递归辅助函数
func addHelper(a, b string, i, j, carry int) string {
	// 递归终止条件
	if i < 0 && j < 0 && carry == 0 {
		return ""
	}

	sum := carry

	if i >= 0 {
		sum += int(a[i] - '0')
	}

	if j >= 0 {
		sum += int(b[j] - '0')
	}

	// 递归处理前面的位
	prefix := addHelper(a, b, i-1, j-1, sum/2)

	return prefix + string('0'+byte(sum%2))
}

// =========================== 方法四：位运算（仅适用于短字符串） ===========================

// addBinary4 位运算解法
// 时间复杂度：O(max(m,n))
// 空间复杂度：O(1)
// 注意：仅适用于长度<=63的字符串
func addBinary4(a string, b string) string {
	// 转换为整数
	numA := binaryToInt(a)
	numB := binaryToInt(b)

	// 使用位运算实现加法
	for numB != 0 {
		sum := numA ^ numB          // 不带进位的和
		carry := (numA & numB) << 1 // 进位
		numA = sum
		numB = carry
	}

	return intToBinary(numA)
}

// binaryToInt 二进制字符串转整数
func binaryToInt(s string) int64 {
	var result int64
	for _, ch := range s {
		result = result*2 + int64(ch-'0')
	}
	return result
}

// intToBinary 整数转二进制字符串
func intToBinary(n int64) string {
	if n == 0 {
		return "0"
	}
	var result string
	for n > 0 {
		result = string('0'+byte(n%2)) + result
		n /= 2
	}
	return result
}

// =========================== 测试代码 ===========================

func main() {
	fmt.Println("=== LeetCode 67: 二进制求和 ===\n")

	// 测试用例
	testCases := []struct {
		a      string
		b      string
		expect string
	}{
		{"11", "1", "100"},        // 示例1: 3 + 1 = 4
		{"1010", "1011", "10101"}, // 示例2: 10 + 11 = 21
		{"0", "0", "0"},           // 边界: 0 + 0
		{"1", "1", "10"},          // 边界: 1 + 1 = 2
		{"1111", "1", "10000"},    // 不同长度: 15 + 1 = 16
		{"1111", "1111", "11110"}, // 相同长度: 15 + 15 = 30
		{"10100000100100110110010000010101111011011001101110111111111101000000101111001110001111100001101", "110101001011101110001111100110001010100001101011101010000011011011001011101111001100000011011110011", "110111101100010011000101110110100000011101000101011001000011011000001100011110011010010011000000000"}, // 长字符串
		{"0", "1010", "1010"},       // 一个为0
		{"100", "110010", "110110"}, // 不同长度
		{"1", "111", "1000"},        // 差距大
	}

	fmt.Println("方法一：双指针+进位+StringBuilder")
	runTests(testCases, addBinary)

	fmt.Println("\n方法二：直接前插")
	runTests(testCases, addBinary2)

	fmt.Println("\n方法三：递归实现")
	runTests(testCases, addBinary3)

	fmt.Println("\n方法四：位运算（仅短字符串）")
	// 只测试短字符串（长度<=63）
	shortCases := testCases[:len(testCases)-1] // 排除超长字符串
	runTests(shortCases, addBinary4)

	// 性能对比
	fmt.Println("\n=== 性能对比 ===")
	performanceTest()

	// 验证二进制计算正确性
	fmt.Println("\n=== 验证计算正确性 ===")
	verifyResults(testCases)
}

// runTests 运行测试用例
func runTests(testCases []struct {
	a      string
	b      string
	expect string
}, fn func(string, string) string) {
	passCount := 0
	for i, tc := range testCases {
		result := fn(tc.a, tc.b)
		status := "✅"
		if result != tc.expect {
			status = "❌"
		} else {
			passCount++
		}

		// 显示简化版的测试信息
		aDisplay := tc.a
		bDisplay := tc.b
		resultDisplay := result
		if len(tc.a) > 20 {
			aDisplay = tc.a[:17] + "..."
		}
		if len(tc.b) > 20 {
			bDisplay = tc.b[:17] + "..."
		}
		if len(result) > 20 {
			resultDisplay = result[:17] + "..."
		}

		fmt.Printf("  测试%d: %s\n", i+1, status)
		fmt.Printf("    输入: a=%s, b=%s\n", aDisplay, bDisplay)
		fmt.Printf("    输出: %s\n", resultDisplay)
		if result != tc.expect {
			expectDisplay := tc.expect
			if len(tc.expect) > 20 {
				expectDisplay = tc.expect[:17] + "..."
			}
			fmt.Printf("    期望: %s\n", expectDisplay)
		}
	}
	fmt.Printf("  通过: %d/%d\n", passCount, len(testCases))
}

// performanceTest 性能测试
func performanceTest() {
	// 生成长字符串
	a := strings.Repeat("1", 1000)
	b := strings.Repeat("1", 1000)

	fmt.Println("  测试数据：两个1000位的全1二进制字符串")

	// 测试方法一
	result1 := addBinary(a, b)
	fmt.Printf("  方法一（双指针+StringBuilder）: 结果长度=%d\n", len(result1))

	// 测试方法二（可能较慢）
	result2 := addBinary2(a, b)
	fmt.Printf("  方法二（直接前插）: 结果长度=%d\n", len(result2))

	// 测试方法三
	result3 := addBinary3(a, b)
	fmt.Printf("  方法三（递归）: 结果长度=%d\n", len(result3))

	fmt.Println("  注：方法四（位运算）不适用于长字符串，会溢出")
}

// verifyResults 验证结果正确性
func verifyResults(testCases []struct {
	a      string
	b      string
	expect string
}) {
	for i, tc := range testCases {
		// 跳过超长字符串（无法转换为int64）
		if len(tc.a) > 60 || len(tc.b) > 60 {
			continue
		}

		// 转换为十进制验证
		numA, _ := strconv.ParseInt(tc.a, 2, 64)
		numB, _ := strconv.ParseInt(tc.b, 2, 64)
		expected, _ := strconv.ParseInt(tc.expect, 2, 64)
		sum := numA + numB

		status := "✅"
		if sum != expected {
			status = "❌"
		}

		fmt.Printf("  验证%d: %s %d + %d = %d (期望%d)\n",
			i+1, status, numA, numB, sum, expected)
	}
}
