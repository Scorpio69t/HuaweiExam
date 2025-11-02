package main

import (
	"fmt"
)

// ==================== 方法一：二维动态规划 ====================
// 时间复杂度：O(m*n)，空间复杂度：O(m*n)
// 经典的二维 DP 解法，状态定义清晰，易于理解
func isInterleave1(s1 string, s2 string, s3 string) bool {
	m, n := len(s1), len(s2)
	// 长度验证：必须满足 len(s1) + len(s2) == len(s3)
	if m+n != len(s3) {
		return false
	}

	// dp[i][j] 表示 s1 的前 i 个字符和 s2 的前 j 个字符
	// 能否交错组成 s3 的前 i+j 个字符
	dp := make([][]bool, m+1)
	for i := range dp {
		dp[i] = make([]bool, n+1)
	}

	// 初始化：空字符串可以交错组成空字符串
	dp[0][0] = true

	// 初始化第一列：只使用 s1 的字符
	for i := 1; i <= m; i++ {
		dp[i][0] = dp[i-1][0] && s1[i-1] == s3[i-1]
	}

	// 初始化第一行：只使用 s2 的字符
	for j := 1; j <= n; j++ {
		dp[0][j] = dp[0][j-1] && s2[j-1] == s3[j-1]
	}

	// 状态转移
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			// 从 s1 取字符：s1[i-1] 匹配 s3[i+j-1]
			// 从 s2 取字符：s2[j-1] 匹配 s3[i+j-1]
			dp[i][j] = (dp[i-1][j] && s1[i-1] == s3[i+j-1]) ||
				(dp[i][j-1] && s2[j-1] == s3[i+j-1])
		}
	}

	return dp[m][n]
}

// ==================== 方法二：一维动态规划（空间优化）====================
// 时间复杂度：O(m*n)，空间复杂度：O(n)
// 滚动数组优化，满足进阶要求的 O(s2.length) 空间
func isInterleave2(s1 string, s2 string, s3 string) bool {
	m, n := len(s1), len(s2)
	if m+n != len(s3) {
		return false
	}

	// dp[j] 表示 s1 的前 i 个字符和 s2 的前 j 个字符
	// 能否交错组成 s3 的前 i+j 个字符
	dp := make([]bool, n+1)
	dp[0] = true

	// 初始化第一行：只使用 s2 的字符
	for j := 1; j <= n; j++ {
		dp[j] = dp[j-1] && s2[j-1] == s3[j-1]
	}

	// 逐行更新
	for i := 1; i <= m; i++ {
		// 更新第一列：只使用 s1 的字符
		dp[0] = dp[0] && s1[i-1] == s3[i-1]

		for j := 1; j <= n; j++ {
			// dp[j] 保留了上一行的值（相当于 dp[i-1][j]）
			// dp[j-1] 是当前行左侧的值（相当于 dp[i][j-1]）
			dp[j] = (dp[j] && s1[i-1] == s3[i+j-1]) ||
				(dp[j-1] && s2[j-1] == s3[i+j-1])
		}
	}

	return dp[n]
}

// ==================== 方法三：记忆化递归（DFS + Memo）====================
// 时间复杂度：O(m*n)，空间复杂度：O(m*n)
// 自顶向下的递归解法，使用记忆化避免重复计算
func isInterleave3(s1 string, s2 string, s3 string) bool {
	m, n := len(s1), len(s2)
	if m+n != len(s3) {
		return false
	}

	// memo[i][j] 缓存 dfs(i, j) 的结果
	// 0: 未计算, 1: true, -1: false
	memo := make([][]int, m+1)
	for i := range memo {
		memo[i] = make([]int, n+1)
	}

	var dfs func(i, j int) bool
	dfs = func(i, j int) bool {
		// 递归终止：所有字符都匹配完成
		if i == m && j == n {
			return true
		}

		// 如果已经计算过，直接返回缓存结果
		if memo[i][j] != 0 {
			return memo[i][j] == 1
		}

		result := false
		// 尝试从 s1 取字符
		if i < m && s1[i] == s3[i+j] {
			result = dfs(i+1, j)
		}
		// 尝试从 s2 取字符
		if !result && j < n && s2[j] == s3[i+j] {
			result = dfs(i, j+1)
		}

		// 缓存结果
		if result {
			memo[i][j] = 1
		} else {
			memo[i][j] = -1
		}

		return result
	}

	return dfs(0, 0)
}

// ==================== 方法四：BFS 搜索 ====================
// 时间复杂度：O(m*n)，空间复杂度：O(m*n)
// 将问题建模为图搜索，使用 BFS 找到目标状态
func isInterleave4(s1 string, s2 string, s3 string) bool {
	m, n := len(s1), len(s2)
	if m+n != len(s3) {
		return false
	}

	// 使用队列进行 BFS
	type State struct {
		i, j int
	}
	queue := []State{{0, 0}}

	// visited[i][j] 标记状态 (i, j) 是否访问过
	visited := make([][]bool, m+1)
	for i := range visited {
		visited[i] = make([]bool, n+1)
	}
	visited[0][0] = true

	for len(queue) > 0 {
		state := queue[0]
		queue = queue[1:]
		i, j := state.i, state.j

		// 到达目标状态
		if i == m && j == n {
			return true
		}

		// 尝试从 s1 取字符
		if i < m && s1[i] == s3[i+j] && !visited[i+1][j] {
			visited[i+1][j] = true
			queue = append(queue, State{i + 1, j})
		}

		// 尝试从 s2 取字符
		if j < n && s2[j] == s3[i+j] && !visited[i][j+1] {
			visited[i][j+1] = true
			queue = append(queue, State{i, j + 1})
		}
	}

	return false
}

// ==================== 方法五：DFS 回溯（暴力搜索）====================
// 时间复杂度：O(2^(m+n))，空间复杂度：O(m+n)
// 纯递归回溯，不使用记忆化，适合理解问题本质
func isInterleave5(s1 string, s2 string, s3 string) bool {
	m, n := len(s1), len(s2)
	if m+n != len(s3) {
		return false
	}

	var dfs func(i, j, k int) bool
	dfs = func(i, j, k int) bool {
		// 递归终止：s3 匹配完成
		if k == len(s3) {
			return i == m && j == n
		}

		// 尝试从 s1 取字符
		if i < m && s1[i] == s3[k] {
			if dfs(i+1, j, k+1) {
				return true
			}
		}

		// 尝试从 s2 取字符
		if j < n && s2[j] == s3[k] {
			if dfs(i, j+1, k+1) {
				return true
			}
		}

		return false
	}

	return dfs(0, 0, 0)
}

// ==================== 辅助函数：DP 表可视化 ====================
func printDPTable(s1, s2, s3 string) {
	m, n := len(s1), len(s2)
	if m+n != len(s3) {
		fmt.Println("长度不匹配，无法构建 DP 表")
		return
	}

	dp := make([][]bool, m+1)
	for i := range dp {
		dp[i] = make([]bool, n+1)
	}
	dp[0][0] = true

	for i := 1; i <= m; i++ {
		dp[i][0] = dp[i-1][0] && s1[i-1] == s3[i-1]
	}
	for j := 1; j <= n; j++ {
		dp[0][j] = dp[0][j-1] && s2[j-1] == s3[j-1]
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			dp[i][j] = (dp[i-1][j] && s1[i-1] == s3[i+j-1]) ||
				(dp[i][j-1] && s2[j-1] == s3[i+j-1])
		}
	}

	// 打印表头
	fmt.Print("     ε")
	for j := 0; j < n; j++ {
		fmt.Printf("  %c", s2[j])
	}
	fmt.Println()

	// 打印 DP 表
	for i := 0; i <= m; i++ {
		if i == 0 {
			fmt.Print("ε  ")
		} else {
			fmt.Printf("%c  ", s1[i-1])
		}
		for j := 0; j <= n; j++ {
			if dp[i][j] {
				fmt.Print("  T")
			} else {
				fmt.Print("  F")
			}
		}
		fmt.Println()
	}
}

// ==================== 辅助函数：路径追踪 ====================
func tracePath(s1, s2, s3 string) {
	m, n := len(s1), len(s2)
	if m+n != len(s3) {
		fmt.Println("无法交错组成")
		return
	}

	dp := make([][]bool, m+1)
	for i := range dp {
		dp[i] = make([]bool, n+1)
	}
	dp[0][0] = true

	for i := 1; i <= m; i++ {
		dp[i][0] = dp[i-1][0] && s1[i-1] == s3[i-1]
	}
	for j := 1; j <= n; j++ {
		dp[0][j] = dp[0][j-1] && s2[j-1] == s3[j-1]
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			dp[i][j] = (dp[i-1][j] && s1[i-1] == s3[i+j-1]) ||
				(dp[i][j-1] && s2[j-1] == s3[i+j-1])
		}
	}

	if !dp[m][n] {
		fmt.Println("无法交错组成")
		return
	}

	// 回溯路径
	path := make([]string, 0)
	i, j := m, n
	for i > 0 || j > 0 {
		if i > 0 && dp[i-1][j] && s1[i-1] == s3[i+j-1] {
			path = append([]string{fmt.Sprintf("s1[%d]='%c'", i-1, s1[i-1])}, path...)
			i--
		} else if j > 0 && dp[i][j-1] && s2[j-1] == s3[i+j-1] {
			path = append([]string{fmt.Sprintf("s2[%d]='%c'", j-1, s2[j-1])}, path...)
			j--
		}
	}

	fmt.Println("交错路径：")
	for idx, step := range path {
		fmt.Printf("步骤 %d: %s\n", idx+1, step)
	}
}

// ==================== 测试函数 ====================
func testIsInterleave(name string, fn func(string, string, string) bool) {
	fmt.Printf("\n========== %s ==========\n", name)

	testCases := []struct {
		s1, s2, s3 string
		expected   bool
	}{
		{"aabcc", "dbbca", "aadbbcbcac", true},
		{"aabcc", "dbbca", "aadbbbaccc", false},
		{"", "", "", true},
		{"abc", "", "abc", true},
		{"", "abc", "abc", true},
		{"ab", "cd", "abcdef", false},
		{"aaa", "aaa", "aaaaaa", true},
		{"a", "b", "ab", true},
		{"a", "b", "ba", true},
		{"aa", "ab", "aaba", true},
	}

	passed := 0
	for i, tc := range testCases {
		result := fn(tc.s1, tc.s2, tc.s3)
		status := "✓"
		if result != tc.expected {
			status = "✗"
		} else {
			passed++
		}
		fmt.Printf("测试用例 %d: %s (s1=\"%s\", s2=\"%s\", s3=\"%s\") => %v (期望: %v)\n",
			i+1, status, tc.s1, tc.s2, tc.s3, result, tc.expected)
	}
	fmt.Printf("通过率: %d/%d\n", passed, len(testCases))
}

// ==================== 性能测试 ====================
func benchmarkIsInterleave() {
	fmt.Println("\n========== 性能测试 ==========")

	// 构造较长的测试用例
	s1 := "bbbbbabbbbabaababaaaabbababbaaabbabbaaabaaaaababbbababbbbbabbbbababbabaabababbbaabababababbbaaababaa"
	s2 := "babaaaabbababbbabbbbaabaabbaabbbbaabaaabaababaaaabaaabbaaabaaaabaabaabbbbbbbbbbbabaaabbababbabbabaab"
	s3 := "babbbabbbaaabbababbbbababaabbabaabaaabbbbabbbaaabbbaaaaabbbbaabbaaabababbaaaaaabababbababaababbababbbababbbbaaaabaabbabbaaaaabbabbaaaabbbaabaaabaababaababbaaabbbbbabbbbaabbabaabbbbabaaabbababbabbabbab"

	methods := []struct {
		name string
		fn   func(string, string, string) bool
	}{
		{"二维 DP", isInterleave1},
		{"一维 DP", isInterleave2},
		{"记忆化递归", isInterleave3},
		{"BFS 搜索", isInterleave4},
		// 不测试 DFS 回溯，因为时间复杂度太高
	}

	for _, method := range methods {
		result := method.fn(s1, s2, s3)
		fmt.Printf("%s: %v\n", method.name, result)
	}
}

// ==================== 主函数 ====================
func main() {
	fmt.Println("LeetCode 97: 交错字符串")
	fmt.Println("==================================")

	// 测试所有方法
	testIsInterleave("方法一：二维动态规划", isInterleave1)
	testIsInterleave("方法二：一维动态规划", isInterleave2)
	testIsInterleave("方法三：记忆化递归", isInterleave3)
	testIsInterleave("方法四：BFS 搜索", isInterleave4)
	testIsInterleave("方法五：DFS 回溯", isInterleave5)

	// DP 表可视化
	fmt.Println("\n========== DP 表可视化 ==========")
	s1 := "aabcc"
	s2 := "dbbca"
	s3 := "aadbbcbcac"
	fmt.Printf("s1 = \"%s\"\n", s1)
	fmt.Printf("s2 = \"%s\"\n", s2)
	fmt.Printf("s3 = \"%s\"\n\n", s3)
	printDPTable(s1, s2, s3)

	// 路径追踪
	fmt.Println("\n========== 路径追踪 ==========")
	tracePath(s1, s2, s3)

	// 性能测试
	benchmarkIsInterleave()

	fmt.Println("\n========== 算法对比总结 ==========")
	fmt.Println("1. 二维 DP：时间 O(m*n)，空间 O(m*n)，经典解法，易理解")
	fmt.Println("2. 一维 DP：时间 O(m*n)，空间 O(n)，最优空间，符合进阶要求")
	fmt.Println("3. 记忆化递归：时间 O(m*n)，空间 O(m*n)，代码简洁，递归风格")
	fmt.Println("4. BFS 搜索：时间 O(m*n)，空间 O(m*n)，图搜索思路，直观")
	fmt.Println("5. DFS 回溯：时间 O(2^(m+n))，空间 O(m+n)，暴力搜索，不推荐")
	fmt.Println("\n推荐方案：一维 DP（空间最优）或二维 DP（易理解）")
}
