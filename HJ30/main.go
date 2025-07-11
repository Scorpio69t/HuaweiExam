package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	var s, t string
	fmt.Scan(&s, &t)

	result := processString(s, t)
	fmt.Println(result)
}

func processString(s, t string) string {
	// 合并阶段
	u := s + t

	// 分离奇偶位字符（按1开始的下标）
	var oddChars, evenChars []rune
	for i, char := range u {
		if (i+1)%2 == 1 { // 奇数位（下标从1开始）
			oddChars = append(oddChars, char)
		} else { // 偶数位
			evenChars = append(evenChars, char)
		}
	}

	// 按ASCII码排序
	sort.Slice(oddChars, func(i, j int) bool {
		return oddChars[i] < oddChars[j]
	})
	sort.Slice(evenChars, func(i, j int) bool {
		return evenChars[i] < evenChars[j]
	})

	// 重新组合：按奇偶位交替排列
	var uPrime strings.Builder
	maxLen := len(oddChars)
	if len(evenChars) > maxLen {
		maxLen = len(evenChars)
	}

	for i := 0; i < maxLen; i++ {
		if i < len(oddChars) {
			uPrime.WriteRune(oddChars[i])
		}
		if i < len(evenChars) {
			uPrime.WriteRune(evenChars[i])
		}
	}

	// 调整阶段
	var result strings.Builder
	for _, char := range uPrime.String() {
		if isHexChar(char) {
			// 转换为十进制
			decValue := hexToDecimal(char)

			// 转换为4位二进制
			binary := fmt.Sprintf("%04b", decValue)

			// 翻转二进制
			reversed := reverseString(binary)

			// 转换回十六进制并转为大写
			reversedDec, _ := strconv.ParseInt(reversed, 2, 64)
			hexResult := strings.ToUpper(fmt.Sprintf("%X", reversedDec))
			result.WriteString(hexResult)
		} else {
			// 非十六进制字符直接添加
			result.WriteRune(char)
		}
	}

	return result.String()
}

func isHexChar(char rune) bool {
	return (char >= '0' && char <= '9') ||
		(char >= 'a' && char <= 'f') ||
		(char >= 'A' && char <= 'F')
}

func hexToDecimal(char rune) int {
	if char >= '0' && char <= '9' {
		return int(char - '0')
	}
	if char >= 'a' && char <= 'f' {
		return int(char - 'a' + 10)
	}
	if char >= 'A' && char <= 'F' {
		return int(char - 'A' + 10)
	}
	return 0
}

func reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
