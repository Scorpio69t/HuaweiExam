package main

import (
	"fmt"
)

// 解法1: 动态规划 - 推荐
func numTreesDP(n int) int {
	// dp[i] 表示由i个节点组成的二叉搜索树的种数
	dp := make([]int, n+1)
	
	// 基础情况
	dp[0] = 1 // 空树
	dp[1] = 1 // 单节点树
	
	// 状态转移
	for i := 2; i <= n; i++ {
		for j := 1; j <= i; j++ {
			// 以j为根节点，左子树有j-1个节点，右子树有i-j个节点
			dp[i] += dp[j-1] * dp[i-j]
		}
	}
	
	return dp[n]
}

// 解法2: 记忆化递归
func numTreesMemo(n int) int {
	memo := make(map[int]int)
	return numTreesHelper(n, memo)
}

func numTreesHelper(n int, memo map[int]int) int {
	// 基础情况
	if n <= 1 {
		return 1
	}
	
	// 检查是否已计算过
	if val, exists := memo[n]; exists {
		return val
	}
	
	count := 0
	for i := 1; i <= n; i++ {
		// 以i为根节点
		left := numTreesHelper(i-1, memo)   // 左子树
		right := numTreesHelper(n-i, memo)  // 右子树
		count += left * right
	}
	
	memo[n] = count
	return count
}

// 解法3: 卡塔兰数公式
func numTreesCatalan(n int) int {
	// 第n个卡塔兰数：C(n) = (2n)! / ((n+1)! * n!)
	// 递推公式：C(n) = C(2n, n) / (n+1)
	
	if n <= 1 {
		return 1
	}
	
	// 计算组合数C(2n, n)
	numerator := 1
	denominator := 1
	
	for i := 0; i < n; i++ {
		numerator *= (2*n - i)
		denominator *= (i + 1)
	}
	
	return numerator / denominator / (n + 1)
}

// 解法4: 卡塔兰数优化递推
func numTreesCatalanOpt(n int) int {
	if n <= 1 {
		return 1
	}
	
	// 使用递推公式：C(n) = C(n-1) * 2 * (2n-1) / (n+1)
	catalan := 1
	for i := 2; i <= n; i++ {
		catalan = catalan * 2 * (2*i - 1) / (i + 1)
	}
	
	return catalan
}

// 解法5: 纯递归（用于理解，效率较低）
func numTreesRecursive(n int) int {
	if n <= 1 {
		return 1
	}
	
	count := 0
	for i := 1; i <= n; i++ {
		left := numTreesRecursive(i - 1)
		right := numTreesRecursive(n - i)
		count += left * right
	}
	
	return count
}

// 生成实际的BST结构（额外功能）
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 生成所有可能的BST
func generateTrees(n int) []*TreeNode {
	if n == 0 {
		return []*TreeNode{nil}
	}
	return generateTreesRange(1, n)
}

func generateTreesRange(start, end int) []*TreeNode {
	var result []*TreeNode
	
	if start > end {
		result = append(result, nil)
		return result
	}
	
	for i := start; i <= end; i++ {
		// 生成左右子树的所有可能组合
		leftTrees := generateTreesRange(start, i-1)
		rightTrees := generateTreesRange(i+1, end)
		
		// 组合左右子树
		for _, left := range leftTrees {
			for _, right := range rightTrees {
				root := &TreeNode{Val: i}
				root.Left = left
				root.Right = right
				result = append(result, root)
			}
		}
	}
	
	return result
}

// 可视化BST结构数量
func visualizeBSTCount(n int) {
	fmt.Printf("\n=== n=%d 的BST构造分析 ===\n", n)
	
	for i := 1; i <= n; i++ {
		left := numTreesDP(i - 1)
		right := numTreesDP(n - i)
		count := left * right
		
		fmt.Printf("以%d为根: 左子树(%d种) × 右子树(%d种) = %d种\n", 
			i, left, right, count)
	}
	
	total := numTreesDP(n)
	fmt.Printf("总计: %d种不同的BST\n", total)
}

// 性能测试
func performanceTest() {
	fmt.Println("\n=== 性能测试 (n=15) ===")
	n := 15
	
	// 测试动态规划
	fmt.Printf("动态规划: %d\n", numTreesDP(n))
	
	// 测试记忆化
	fmt.Printf("记忆化递归: %d\n", numTreesMemo(n))
	
	// 测试卡塔兰数
	fmt.Printf("卡塔兰数公式: %d\n", numTreesCatalan(n))
	
	// 测试优化递推
	fmt.Printf("卡塔兰数递推: %d\n", numTreesCatalanOpt(n))
	
	// 递归解法太慢，只测试小数据
	if n <= 10 {
		fmt.Printf("纯递归: %d\n", numTreesRecursive(n))
	}
}

// 测试函数
func runTests() {
	fmt.Println("=== 不同二叉搜索树算法测试 ===")
	
	testCases := []int{1, 2, 3, 4, 5, 6}
	
	fmt.Println("\n基本测试:")
	for _, n := range testCases {
		dp := numTreesDP(n)
		memo := numTreesMemo(n)
		catalan := numTreesCatalan(n)
		catalanOpt := numTreesCatalanOpt(n)
		
		fmt.Printf("n=%d: DP=%d, Memo=%d, Catalan=%d, CatalanOpt=%d\n", 
			n, dp, memo, catalan, catalanOpt)
		
		// 验证所有方法结果一致
		if dp != memo || memo != catalan || catalan != catalanOpt {
			fmt.Printf("❌ 结果不一致！\n")
		}
	}
	
	// 特殊测试
	fmt.Println("\n特殊测试:")
	fmt.Printf("n=1: %d (预期: 1)\n", numTreesDP(1))
	fmt.Printf("n=3: %d (预期: 5)\n", numTreesDP(3))
	
	// 边界测试
	fmt.Println("\n边界测试:")
	fmt.Printf("n=0: %d\n", numTreesDP(0))
	fmt.Printf("n=19: %d (题目上限)\n", numTreesDP(19))
	
	// 可视化分析
	visualizeBSTCount(3)
	
	// 性能测试
	performanceTest()
	
	// 实际生成BST（小数据）
	fmt.Println("\n=== 实际BST生成 (n=3) ===")
	trees := generateTrees(3)
	fmt.Printf("生成了 %d 种不同的BST结构\n", len(trees))
	
	fmt.Println("\n所有测试完成! ✅")
}

func main() {
	runTests()
}
