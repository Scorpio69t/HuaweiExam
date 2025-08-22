package main

import (
	"fmt"
	"sort"
)

// 解法1：回溯（剪枝）
func generateParenthesisBacktrack(n int) []string {
	var ans []string
	path := make([]byte, 0, 2*n)
	var dfs func(l, r int)
	dfs = func(l, r int) {
		if l > n || r > l { // 剪枝
			return
		}
		if l == n && r == n {
			ans = append(ans, string(path))
			return
		}
		if l < n {
			path = append(path, '(')
			dfs(l+1, r)
			path = path[:len(path)-1]
		}
		if r < l {
			path = append(path, ')')
			dfs(l, r+1)
			path = path[:len(path)-1]
		}
	}
	dfs(0, 0)
	return ans
}

// 解法2：生成式动态规划（卡塔兰结构拼接）
func generateParenthesisDP(n int) []string {
	dp := make([][]string, n+1)
	dp[0] = []string{""}
	for i := 1; i <= n; i++ {
		cur := make([]string, 0)
		for left := 0; left <= i-1; left++ {
			right := i - 1 - left
			for _, a := range dp[left] {
				for _, b := range dp[right] {
					cur = append(cur, "("+a+")"+b)
				}
			}
		}
		sort.Strings(cur) // 便于稳定输出
		dp[i] = cur
	}
	return dp[n]
}

// 辅助：卡塔兰数计数（用于校验数量）
func catalan(n int) int {
	if n <= 1 {
		return 1
	}
	// C(n) = C(2n, n) / (n+1)
	num := 1
	den := 1
	for i := 1; i <= n; i++ {
		num *= (n + i) // 从 n+1 乘到 2n
		den *= i       // 从 1 乘到 n
	}
	return num / den / (n + 1)
}

func main() {
	fmt.Println("=== 括号生成 算法测试 ===")
	for _, n := range []int{1, 2, 3, 4} {
		a := generateParenthesisBacktrack(n)
		b := generateParenthesisDP(n)
		fmt.Printf("n=%d: 回溯=%d, DP=%d, Catalan=%d\n", n, len(a), len(b), catalan(n))
	}

	fmt.Println("\n示例(n=3):")
	for _, s := range generateParenthesisBacktrack(3) {
		fmt.Println(s)
	}
}
