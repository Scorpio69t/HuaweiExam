package main

import (
	"fmt"
	"strconv"
	"strings"
)

// 方法一：回溯算法解法（推荐）
// 时间复杂度：O(3^4)，空间复杂度：O(n)
func restoreIpAddresses(s string) []string {
	var result []string
	var segments []string

	backtrack(s, 0, segments, &result)
	return result
}

func backtrack(s string, start int, segments []string, result *[]string) {
	// 如果已经有4个段，检查是否用完所有字符
	if len(segments) == 4 {
		if start == len(s) {
			*result = append(*result, strings.Join(segments, "."))
		}
		return
	}

	// 尝试取1-3位数字
	for i := 1; i <= 3 && start+i <= len(s); i++ {
		segment := s[start : start+i]

		// 验证段的有效性
		if isValidSegment(segment) {
			segments = append(segments, segment)
			backtrack(s, start+i, segments, result)
			segments = segments[:len(segments)-1] // 回溯
		}
	}
}

func isValidSegment(segment string) bool {
	// 长度检查
	if len(segment) == 0 || len(segment) > 3 {
		return false
	}

	// 前导0检查
	if len(segment) > 1 && segment[0] == '0' {
		return false
	}

	// 数值范围检查
	num, err := strconv.Atoi(segment)
	if err != nil {
		return false
	}

	return num >= 0 && num <= 255
}

// 方法二：暴力枚举解法
// 时间复杂度：O(3^4)，空间复杂度：O(n)
func restoreIpAddressesBruteForce(s string) []string {
	var result []string
	n := len(s)

	// 枚举所有可能的分割点
	for i := 1; i <= 3 && i < n-2; i++ {
		for j := i + 1; j <= i+3 && j < n-1; j++ {
			for k := j + 1; k <= j+3 && k < n; k++ {
				// 分割成4个段
				seg1 := s[:i]
				seg2 := s[i:j]
				seg3 := s[j:k]
				seg4 := s[k:]

				// 检查所有段是否有效
				if isValidSegment(seg1) && isValidSegment(seg2) && 
				   isValidSegment(seg3) && isValidSegment(seg4) {
					result = append(result, seg1+"."+seg2+"."+seg3+"."+seg4)
				}
			}
		}
	}

	return result
}

// 方法三：优化的回溯算法
// 时间复杂度：O(3^4)，空间复杂度：O(n)
func restoreIpAddressesOptimized(s string) []string {
	var result []string
	var segments []string

	backtrackOptimized(s, 0, segments, &result)
	return result
}

func backtrackOptimized(s string, start int, segments []string, result *[]string) {
	// 剪枝：如果剩余字符太多或太少，直接返回
	remaining := len(s) - start
	needed := 4 - len(segments)
	
	if remaining < needed || remaining > needed*3 {
		return
	}

	// 如果已经有4个段，检查是否用完所有字符
	if len(segments) == 4 {
		if start == len(s) {
			*result = append(*result, strings.Join(segments, "."))
		}
		return
	}

	// 尝试取1-3位数字
	for i := 1; i <= 3 && start+i <= len(s); i++ {
		segment := s[start : start+i]

		// 验证段的有效性
		if isValidSegmentOptimized(segment) {
			segments = append(segments, segment)
			backtrackOptimized(s, start+i, segments, result)
			segments = segments[:len(segments)-1] // 回溯
		}
	}
}

func isValidSegmentOptimized(segment string) bool {
	// 长度检查
	if len(segment) == 0 || len(segment) > 3 {
		return false
	}

	// 前导0检查
	if len(segment) > 1 && segment[0] == '0' {
		return false
	}

	// 数值范围检查（优化版本）
	if len(segment) == 1 {
		return true // 单个数字总是有效的
	}
	if len(segment) == 2 {
		return segment[0] != '0' // 两位数不能以0开头
	}
	if len(segment) == 3 {
		if segment[0] == '0' {
			return false
		}
		// 快速检查是否超过255
		if segment[0] > '2' {
			return false
		}
		if segment[0] == '2' {
			if segment[1] > '5' {
				return false
			}
			if segment[1] == '5' && segment[2] > '5' {
				return false
			}
		}
		return true
	}

	return false
}

// 方法四：递归解法（更直观的实现）
func restoreIpAddressesRecursive(s string) []string {
	var result []string
	dfs(s, 0, 0, "", &result)
	return result
}

func dfs(s string, index, count int, current string, result *[]string) {
	// 如果已经处理完所有字符
	if index == len(s) {
		if count == 4 {
			*result = append(*result, current[:len(current)-1]) // 移除最后的点
		}
		return
	}

	// 如果已经有4个段但还有字符剩余
	if count >= 4 {
		return
	}

	// 尝试取1-3位数字
	for i := 1; i <= 3 && index+i <= len(s); i++ {
		segment := s[index : index+i]
		
		if isValidSegment(segment) {
			newCurrent := current + segment + "."
			dfs(s, index+i, count+1, newCurrent, result)
		}
	}
}

// 辅助函数：打印IP地址列表
func printIpAddresses(ips []string, desc string) {
	fmt.Printf("%s: %v\n", desc, ips)
	fmt.Printf("数量: %d\n", len(ips))
	fmt.Println()
}

func main() {
	fmt.Println("=== 93. 复原 IP 地址 ===")

	// 测试用例1
	s1 := "25525511135"
	fmt.Printf("测试用例1: s=%s\n", s1)
	fmt.Printf("回溯算法解法结果: %v\n", restoreIpAddresses(s1))
	fmt.Printf("暴力枚举解法结果: %v\n", restoreIpAddressesBruteForce(s1))
	fmt.Printf("优化回溯解法结果: %v\n", restoreIpAddressesOptimized(s1))
	fmt.Printf("递归解法结果: %v\n", restoreIpAddressesRecursive(s1))
	fmt.Println()

	// 测试用例2
	s2 := "0000"
	fmt.Printf("测试用例2: s=%s\n", s2)
	fmt.Printf("回溯算法解法结果: %v\n", restoreIpAddresses(s2))
	fmt.Printf("暴力枚举解法结果: %v\n", restoreIpAddressesBruteForce(s2))
	fmt.Printf("优化回溯解法结果: %v\n", restoreIpAddressesOptimized(s2))
	fmt.Printf("递归解法结果: %v\n", restoreIpAddressesRecursive(s2))
	fmt.Println()

	// 测试用例3
	s3 := "101023"
	fmt.Printf("测试用例3: s=%s\n", s3)
	fmt.Printf("回溯算法解法结果: %v\n", restoreIpAddresses(s3))
	fmt.Printf("暴力枚举解法结果: %v\n", restoreIpAddressesBruteForce(s3))
	fmt.Printf("优化回溯解法结果: %v\n", restoreIpAddressesOptimized(s3))
	fmt.Printf("递归解法结果: %v\n", restoreIpAddressesRecursive(s3))
	fmt.Println()

	// 边界测试用例
	testCases := []struct {
		s    string
		desc string
	}{
		{"1111", "简单四位数"},
		{"123", "字符串太短"},
		{"12345678901234567890", "字符串太长"},
		{"010010", "包含前导0"},
		{"255255255255", "最大值"},
		{"0000", "全零"},
		{"1234", "四位数"},
		{"12345", "五位数"},
	}

	for _, tc := range testCases {
		result := restoreIpAddresses(tc.s)
		fmt.Printf("%s: s=%s, 结果=%v, 数量=%d\n", tc.desc, tc.s, result, len(result))
	}

	// 算法正确性验证
	fmt.Println("\n=== 算法正确性验证 ===")
	verifyS := "25525511135"
	fmt.Printf("验证字符串: %s\n", verifyS)
	
	result1 := restoreIpAddresses(verifyS)
	result2 := restoreIpAddressesBruteForce(verifyS)
	result3 := restoreIpAddressesOptimized(verifyS)
	result4 := restoreIpAddressesRecursive(verifyS)

	fmt.Printf("回溯算法: %v\n", result1)
	fmt.Printf("暴力枚举: %v\n", result2)
	fmt.Printf("优化回溯: %v\n", result3)
	fmt.Printf("递归解法: %v\n", result4)

	// 验证所有解法结果一致
	if len(result1) == len(result2) && len(result1) == len(result3) && len(result1) == len(result4) {
		fmt.Println("✅ 所有解法结果数量一致，算法正确！")
	} else {
		fmt.Println("❌ 解法结果数量不一致，需要检查！")
	}

	// 性能测试
	fmt.Println("\n=== 性能测试 ===")
	performanceS := "123456789"
	fmt.Printf("性能测试字符串: %s\n", performanceS)
	
	result := restoreIpAddresses(performanceS)
	fmt.Printf("回溯算法结果数量: %d\n", len(result))
	fmt.Printf("前5个结果: %v\n", result[:min(5, len(result))])
}

// 辅助函数：返回两个数的最小值
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
