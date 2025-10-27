package main

import (
	"fmt"
)

// =========================== 方法一：对称反射法（最优解法） ===========================

func grayCode1(n int) []int {
	if n == 1 {
		return []int{0, 1}
	}

	// 递归生成n-1位格雷码
	prev := grayCode1(n - 1)
	result := make([]int, len(prev)*2)

	// 前一半：直接添加0
	for i := 0; i < len(prev); i++ {
		result[i] = prev[i]
	}

	// 后一半：添加1并镜像反射
	for i := 0; i < len(prev); i++ {
		result[len(prev)+i] = prev[len(prev)-1-i] + (1 << (n - 1))
	}

	return result
}

// =========================== 方法二：位运算公式（最简洁） ===========================

func grayCode2(n int) []int {
	result := make([]int, 1<<n)
	for i := 0; i < 1<<n; i++ {
		result[i] = i ^ (i >> 1)
	}
	return result
}

// =========================== 方法三：迭代构造 ===========================

func grayCode3(n int) []int {
	result := []int{0}

	for i := 0; i < n; i++ {
		size := len(result)
		for j := size - 1; j >= 0; j-- {
			result = append(result, result[j]|(1<<i))
		}
	}

	return result
}

// =========================== 方法四：递归构造（DFS） ===========================

func grayCode4(n int) []int {
	if n == 1 {
		return []int{0, 1}
	}

	result := []int{0}
	mask := 1 << (n - 1)

	var dfs func(int, int)
	dfs = func(bits, pos int) {
		if bits == 0 {
			return
		}

		size := len(result)
		for i := size - 1; i >= 0; i-- {
			result = append(result, result[i]^mask)
		}

		dfs(bits-1, pos<<1)
	}

	dfs(n-1, 1)
	return result
}

// =========================== 测试代码 ===========================

func main() {
	fmt.Println("=== LeetCode 89: 格雷编码 ===\n")

	testCases := []struct {
		name     string
		n        int
		expected []int
	}{
		{
			name:     "Test1: n=1",
			n:        1,
			expected: []int{0, 1},
		},
		{
			name:     "Test2: n=2",
			n:        2,
			expected: []int{0, 1, 3, 2},
		},
		{
			name:     "Test3: n=3",
			n:        3,
			expected: []int{0, 1, 3, 2, 6, 7, 5, 4},
		},
		{
			name:     "Test4: n=4",
			n:        4,
			expected: []int{0, 1, 3, 2, 6, 7, 5, 4, 12, 13, 15, 14, 10, 11, 9, 8},
		},
	}

	methods := map[string]func(int) []int{
		"对称反射法（最优解法）": grayCode1,
		"位运算公式（最简洁）":  grayCode2,
		"迭代构造":        grayCode3,
		"递归构造（DFS）":   grayCode4,
	}

	for name, method := range methods {
		fmt.Printf("方法%s：%s\n", name, name)
		passCount := 0
		for i, tt := range testCases {
			got := method(tt.n)

			// 验证格雷码的有效性
			valid := isValidGrayCode(got)
			status := "✅"
			if !valid {
				status = "❌"
			} else {
				passCount++
			}
			fmt.Printf("  测试%d: %s\n", i+1, status)
			if status == "❌" {
				fmt.Printf("    输入: n=%d\n", tt.n)
				fmt.Printf("    输出: %v\n", got)
			}
		}
		fmt.Printf("  通过: %d/%d\n\n", passCount, len(testCases))
	}
}

// 验证格雷码的有效性
func isValidGrayCode(code []int) bool {
	if len(code) == 0 {
		return true
	}

	// 检查相邻元素是否只有一位不同
	for i := 0; i < len(code)-1; i++ {
		diff := code[i] ^ code[i+1]
		if diff == 0 || (diff&(diff-1)) != 0 {
			// 检查是否只有一个位不同
			ones := 0
			for diff > 0 {
				ones += diff & 1
				diff >>= 1
			}
			if ones != 1 {
				return false
			}
		}
	}

	// 检查首尾是否只有一位不同
	if len(code) > 1 {
		diff := code[0] ^ code[len(code)-1]
		ones := 0
		for diff > 0 {
			ones += diff & 1
			diff >>= 1
		}
		if ones != 1 {
			return false
		}
	}

	return true
}
