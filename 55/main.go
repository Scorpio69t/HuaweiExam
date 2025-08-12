package main

import (
	"fmt"
	"strings"
	"time"
)

// 解法一：贪心（最优解，时间O(n)，空间O(1)）
// 维护最远可达下标 farthest，可以在遍历过程中动态扩展可达范围。
// 若某一时刻 i > farthest，说明当前位置不可达，直接返回 false。
// 最终只要 farthest >= n-1 即可到达终点。
func canJumpGreedy(nums []int) bool {
	farthest := 0
	for i := 0; i < len(nums); i++ {
		if i > farthest {
			return false
		}
		if i+nums[i] > farthest {
			farthest = i + nums[i]
		}
		if farthest >= len(nums)-1 {
			return true
		}
	}
	return true
}

// 解法二：动态规划（从后往前，时间O(n^2)，空间O(n)）
// dp[i] 表示从 i 出发是否能到达终点。枚举 i 能跳到的每个 j，看是否有 dp[j] 为真。
// 仅用于教学对比，实际大数据会超时；可加剪枝或二分优化思路，但贪心更优雅。
func canJumpDP(nums []int) bool {
	n := len(nums)
	dp := make([]bool, n)
	dp[n-1] = true
	for i := n - 2; i >= 0; i-- {
		maxJ := min(i+nums[i], n-1)
		for j := i + 1; j <= maxJ; j++ {
			if dp[j] {
				dp[i] = true
				break
			}
		}
	}
	return dp[0]
}

// 解法三：反向贪心（从右往左收缩目标，时间O(n)，空间O(1)）
// 维护一个 lastGood 作为“必须到达”的最右位置；如果 i + nums[i] >= lastGood，
// 则把 lastGood 更新为 i。最终判断 lastGood == 0。
func canJumpReverseGreedy(nums []int) bool {
	lastGood := len(nums) - 1
	for i := len(nums) - 2; i >= 0; i-- {
		if i+nums[i] >= lastGood {
			lastGood = i
		}
	}
	return lastGood == 0
}

// 测试与简单性能对比
func runTests() {
	type testCase struct {
		nums     []int
		expected bool
		desc     string
	}
	tests := []testCase{
		{[]int{2, 3, 1, 1, 4}, true, "示例1 可达"},
		{[]int{3, 2, 1, 0, 4}, false, "示例2 不可达"},
		{[]int{0}, true, "单元素 0"},
		{[]int{1, 0, 1, 0}, false, "被 0 阻断"},
		{[]int{2, 0, 0}, true, "边界跳过"},
		{[]int{1, 1, 1, 1}, true, "全 1"},
		{[]int{4, 0, 0, 0, 0}, true, "首位大跳"},
	}

	fmt.Println("=== 55. 跳跃游戏 - 测试 ===")
	for i, tc := range tests {
		g := canJumpGreedy(tc.nums)
		r := canJumpReverseGreedy(tc.nums)
		d := canJumpDP(tc.nums)
		ok := (g == tc.expected) && (r == tc.expected) && (d == tc.expected)
		status := "✅"
		if !ok {
			status = "❌"
		}
		fmt.Printf("用例 %d: %s\n", i+1, tc.desc)
		fmt.Printf("输入: %v\n期望: %v\n贪心: %v, 反向贪心: %v, DP: %v\n", tc.nums, tc.expected, g, r, d)
		fmt.Printf("结果: %s\n", status)
		fmt.Println(strings.Repeat("-", 40))
	}
}

func benchmark() {
	fmt.Println("\n=== 简单性能对比（粗略） ===")
	// 构造较大用例：长度 100000，前面给出足够跳力
	n := 100000
	nums := make([]int, n)
	for i := 0; i < n-1; i++ {
		nums[i] = 2
	}
	nums[n-1] = 0

	start := time.Now()
	_ = canJumpGreedy(nums)
	d1 := time.Since(start)

	start = time.Now()
	_ = canJumpReverseGreedy(nums)
	d2 := time.Since(start)

	// 警告：O(n^2) DP 在大输入上会非常慢，这里仅做一次小规模对比
	small := []int{2, 3, 1, 1, 4, 0, 0, 3, 1, 2, 0}
	start = time.Now()
	_ = canJumpDP(small)
	d3 := time.Since(start)

	fmt.Printf("贪心: %v\n", d1)
	fmt.Printf("反向贪心: %v\n", d2)
	fmt.Printf("DP(小规模): %v\n", d3)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	fmt.Println("55. 跳跃游戏")
	fmt.Println(strings.Repeat("=", 40))
	runTests()
	benchmark()
}
