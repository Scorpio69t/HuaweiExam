package main

import (
	"fmt"
)

// 计算最长递增子序列（LIS）
func calculateLIS(heights []int) []int {
	n := len(heights)
	lis := make([]int, n)

	// 初始化：每个位置自身构成长度为1的递增序列
	for i := 0; i < n; i++ {
		lis[i] = 1
	}

	// 动态规划计算LIS
	for i := 1; i < n; i++ {
		for j := 0; j < i; j++ {
			// 如果heights[j] < heights[i]，可以构成更长的递增序列
			if heights[j] < heights[i] {
				if lis[j]+1 > lis[i] {
					lis[i] = lis[j] + 1
				}
			}
		}
	}

	return lis
}

// 计算最长递减子序列（LDS）
func calculateLDS(heights []int) []int {
	n := len(heights)
	lds := make([]int, n)

	// 初始化：每个位置自身构成长度为1的递减序列
	for i := 0; i < n; i++ {
		lds[i] = 1
	}

	// 动态规划计算LDS（从右向左）
	for i := n - 2; i >= 0; i-- {
		for j := i + 1; j < n; j++ {
			// 如果heights[i] > heights[j]，可以构成更长的递减序列
			if heights[i] > heights[j] {
				if lds[j]+1 > lds[i] {
					lds[i] = lds[j] + 1
				}
			}
		}
	}

	return lds
}

// 求解合唱队形问题
func solveChoirFormation(heights []int) int {
	n := len(heights)

	// 特殊情况：只有一个同学
	if n == 1 {
		return 0
	}

	// 计算LIS和LDS
	lis := calculateLIS(heights)
	lds := calculateLDS(heights)

	// 枚举每个位置作为峰值，找到最长的合唱队形
	maxLength := 0
	for i := 0; i < n; i++ {
		// 以位置i为峰值的合唱队形长度
		// LIS[i] + LDS[i] - 1（峰值被计算了两次，需要减1）
		choirLength := lis[i] + lds[i] - 1
		if choirLength > maxLength {
			maxLength = choirLength
		}
	}

	// 返回需要出列的同学数量
	return n - maxLength
}

func main() {
	var n int
	fmt.Scan(&n)

	heights := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&heights[i])
	}

	result := solveChoirFormation(heights)
	fmt.Println(result)
}
