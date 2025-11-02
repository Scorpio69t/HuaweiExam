package main

import (
	"fmt"
	"strings"
	"time"
)

// TreeNode 二叉树节点定义
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// ========== 方法1: 递归分治（经典解法） ==========
func generateTrees1(n int) []*TreeNode {
	if n == 0 {
		return []*TreeNode{}
	}
	return generateTreesHelper1(1, n)
}

func generateTreesHelper1(start, end int) []*TreeNode {
	// 终止条件：区间无效
	if start > end {
		return []*TreeNode{nil}
	}

	allTrees := []*TreeNode{}

	// 枚举每个数字作为根节点
	for i := start; i <= end; i++ {
		// 递归生成所有可能的左子树
		leftTrees := generateTreesHelper1(start, i-1)

		// 递归生成所有可能的右子树
		rightTrees := generateTreesHelper1(i+1, end)

		// 笛卡尔积：组合所有可能的左右子树
		for _, left := range leftTrees {
			for _, right := range rightTrees {
				// 创建当前根节点
				root := &TreeNode{Val: i}
				root.Left = left
				root.Right = right
				allTrees = append(allTrees, root)
			}
		}
	}

	return allTrees
}

// ========== 方法2: 记忆化递归（优化） ==========
func generateTrees2(n int) []*TreeNode {
	if n == 0 {
		return []*TreeNode{}
	}
	memo := make(map[[2]int][]*TreeNode)
	return generateTreesHelper2(1, n, memo)
}

func generateTreesHelper2(start, end int, memo map[[2]int][]*TreeNode) []*TreeNode {
	if start > end {
		return []*TreeNode{nil}
	}

	// 检查缓存
	key := [2]int{start, end}
	if cached, exists := memo[key]; exists {
		return cached
	}

	allTrees := []*TreeNode{}

	for i := start; i <= end; i++ {
		leftTrees := generateTreesHelper2(start, i-1, memo)
		rightTrees := generateTreesHelper2(i+1, end, memo)

		for _, left := range leftTrees {
			for _, right := range rightTrees {
				root := &TreeNode{Val: i}
				root.Left = left
				root.Right = right
				allTrees = append(allTrees, root)
			}
		}
	}

	// 缓存结果
	memo[key] = allTrees
	return allTrees
}

// ========== 方法3: 动态规划（自底向上） ==========
func generateTrees3(n int) []*TreeNode {
	if n == 0 {
		return []*TreeNode{}
	}

	// dp[len][start] 表示从start开始，长度为len的所有BST
	dp := make(map[int]map[int][]*TreeNode)

	// 初始化长度为0的情况
	dp[0] = make(map[int][]*TreeNode)
	for i := 1; i <= n+1; i++ {
		dp[0][i] = []*TreeNode{nil}
	}

	// 按长度递推
	for length := 1; length <= n; length++ {
		dp[length] = make(map[int][]*TreeNode)

		for start := 1; start <= n-length+1; start++ {
			end := start + length - 1
			allTrees := []*TreeNode{}

			for i := start; i <= end; i++ {
				leftLen := i - start
				rightLen := end - i

				leftTrees := dp[leftLen][start]
				rightTrees := dp[rightLen][i+1]

				for _, left := range leftTrees {
					for _, right := range rightTrees {
						root := &TreeNode{Val: i}
						root.Left = left
						root.Right = right
						allTrees = append(allTrees, root)
					}
				}
			}

			dp[length][start] = allTrees
		}
	}

	return dp[n][1]
}

// ========== 方法4: 克隆偏移法 ==========
func generateTrees4(n int) []*TreeNode {
	if n == 0 {
		return []*TreeNode{}
	}
	return generateTreesClone(1, n)
}

func generateTreesClone(start, end int) []*TreeNode {
	if start > end {
		return []*TreeNode{nil}
	}

	if start == end {
		return []*TreeNode{{Val: start}}
	}

	allTrees := []*TreeNode{}

	for i := start; i <= end; i++ {
		leftTrees := generateTreesClone(start, i-1)
		rightTrees := generateTreesClone(i+1, end)

		for _, left := range leftTrees {
			for _, right := range rightTrees {
				root := &TreeNode{Val: i}
				root.Left = cloneTree(left)
				root.Right = cloneTree(right)
				allTrees = append(allTrees, root)
			}
		}
	}

	return allTrees
}

// 深度克隆树
func cloneTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	newRoot := &TreeNode{Val: root.Val}
	newRoot.Left = cloneTree(root.Left)
	newRoot.Right = cloneTree(root.Right)
	return newRoot
}

// ========== 方法5: 优化的记忆化递归 ==========
type TreeList []*TreeNode

func generateTrees5(n int) []*TreeNode {
	if n == 0 {
		return []*TreeNode{}
	}

	memo := make(map[string][]*TreeNode)
	return generateTreesHelper5(1, n, memo)
}

func generateTreesHelper5(start, end int, memo map[string][]*TreeNode) []*TreeNode {
	if start > end {
		return []*TreeNode{nil}
	}

	// 使用字符串作为键
	key := fmt.Sprintf("%d-%d", start, end)
	if cached, exists := memo[key]; exists {
		return cached
	}

	allTrees := []*TreeNode{}

	for i := start; i <= end; i++ {
		leftTrees := generateTreesHelper5(start, i-1, memo)
		rightTrees := generateTreesHelper5(i+1, end, memo)

		for _, left := range leftTrees {
			for _, right := range rightTrees {
				root := &TreeNode{Val: i}
				root.Left = left
				root.Right = right
				allTrees = append(allTrees, root)
			}
		}
	}

	memo[key] = allTrees
	return allTrees
}

// ========== 工具函数 ==========

// 序列化树为数组（层序遍历）
func serializeTree(root *TreeNode) []interface{} {
	if root == nil {
		return []interface{}{}
	}

	result := []interface{}{}
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		if node == nil {
			result = append(result, nil)
		} else {
			result = append(result, node.Val)
			queue = append(queue, node.Left)
			queue = append(queue, node.Right)
		}
	}

	// 移除尾部的nil
	for len(result) > 0 && result[len(result)-1] == nil {
		result = result[:len(result)-1]
	}

	return result
}

// 打印树结构（中序遍历验证BST性质）
func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	result := []int{}
	result = append(result, inorderTraversal(root.Left)...)
	result = append(result, root.Val)
	result = append(result, inorderTraversal(root.Right)...)

	return result
}

// 验证是否为BST
func isValidBST(root *TreeNode) bool {
	return isValidBSTHelper(root, nil, nil)
}

func isValidBSTHelper(root *TreeNode, min, max *int) bool {
	if root == nil {
		return true
	}

	if min != nil && root.Val <= *min {
		return false
	}
	if max != nil && root.Val >= *max {
		return false
	}

	return isValidBSTHelper(root.Left, min, &root.Val) &&
		isValidBSTHelper(root.Right, &root.Val, max)
}

// 计算树的节点数
func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return 1 + countNodes(root.Left) + countNodes(root.Right)
}

// 计算树的高度
func treeHeight(root *TreeNode) int {
	if root == nil {
		return 0
	}

	leftHeight := treeHeight(root.Left)
	rightHeight := treeHeight(root.Right)

	if leftHeight > rightHeight {
		return leftHeight + 1
	}
	return rightHeight + 1
}

// 打印树的可视化结构
func printTree(root *TreeNode, prefix string, isLeft bool) {
	if root == nil {
		return
	}

	fmt.Print(prefix)
	if isLeft {
		fmt.Print("├── ")
	} else {
		fmt.Print("└── ")
	}
	fmt.Println(root.Val)

	if root.Left != nil || root.Right != nil {
		if root.Left != nil {
			printTree(root.Left, prefix+getPrefix(isLeft), true)
		} else {
			fmt.Println(prefix + getPrefix(isLeft) + "├── nil")
		}

		if root.Right != nil {
			printTree(root.Right, prefix+getPrefix(isLeft), false)
		} else {
			fmt.Println(prefix + getPrefix(isLeft) + "└── nil")
		}
	}
}

func getPrefix(isLeft bool) string {
	if isLeft {
		return "│   "
	}
	return "    "
}

// 计算卡塔兰数
func catalanNumber(n int) int {
	if n <= 1 {
		return 1
	}

	catalan := make([]int, n+1)
	catalan[0], catalan[1] = 1, 1

	for i := 2; i <= n; i++ {
		for j := 0; j < i; j++ {
			catalan[i] += catalan[j] * catalan[i-1-j]
		}
	}

	return catalan[n]
}

// 比较两棵树是否结构相同
func isSameTree(p, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}
	if p.Val != q.Val {
		return false
	}
	return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}

// 检查树列表中是否有重复
func hasDuplicateTrees(trees []*TreeNode) bool {
	for i := 0; i < len(trees); i++ {
		for j := i + 1; j < len(trees); j++ {
			if isSameTree(trees[i], trees[j]) {
				return true
			}
		}
	}
	return false
}

// ========== 测试和性能评估 ==========
func main() {
	// 测试用例
	testCases := []struct {
		name     string
		n        int
		expected int // 期望的树的数量（卡塔兰数）
	}{
		{
			name:     "示例1: n=3",
			n:        3,
			expected: 5,
		},
		{
			name:     "示例2: n=1",
			n:        1,
			expected: 1,
		},
		{
			name:     "测试3: n=2",
			n:        2,
			expected: 2,
		},
		{
			name:     "测试4: n=4",
			n:        4,
			expected: 14,
		},
		{
			name:     "测试5: n=5",
			n:        5,
			expected: 42,
		},
		{
			name:     "测试6: n=6",
			n:        6,
			expected: 132,
		},
		{
			name:     "测试7: n=7",
			n:        7,
			expected: 429,
		},
	}

	// 算法方法
	methods := []struct {
		name string
		fn   func(int) []*TreeNode
	}{
		{"递归分治", generateTrees1},
		{"记忆化递归", generateTrees2},
		{"动态规划", generateTrees3},
		{"克隆偏移法", generateTrees4},
		{"优化记忆化", generateTrees5},
	}

	fmt.Println("=== LeetCode 95. 不同的二叉搜索树 II - 测试结果 ===")
	fmt.Println()

	// 运行测试
	for _, tc := range testCases {
		fmt.Printf("测试用例: %s\n", tc.name)
		fmt.Printf("输入: n = %d\n", tc.n)
		fmt.Printf("期望树的数量(第%d个卡塔兰数): %d\n", tc.n, tc.expected)

		allPassed := true
		var results [][]*TreeNode
		var times []time.Duration

		for _, method := range methods {
			start := time.Now()
			result := method.fn(tc.n)
			elapsed := time.Since(start)

			results = append(results, result)
			times = append(times, elapsed)

			status := "✅"
			valid := true

			// 验证数量
			if len(result) != tc.expected {
				status = "❌"
				valid = false
				allPassed = false
			}

			// 验证每棵树都是有效的BST
			if valid {
				for _, tree := range result {
					if !isValidBST(tree) {
						status = "❌"
						valid = false
						allPassed = false
						break
					}
				}
			}

			// 检查是否有重复
			if valid && hasDuplicateTrees(result) {
				status = "⚠️"
				fmt.Printf("  %s: %s (数量: %d, 耗时: %v) - 有重复树\n",
					method.name, status, len(result), elapsed)
			} else {
				fmt.Printf("  %s: %s (数量: %d, 耗时: %v)\n",
					method.name, status, len(result), elapsed)
			}
		}

		if allPassed {
			fmt.Println("✅ 所有方法均通过")
		} else {
			fmt.Println("❌ 存在失败的方法")
		}
		fmt.Println(strings.Repeat("-", 60))
	}

	// BST结构演示
	fmt.Println("\n=== BST结构演示 (n=3) ===")
	demoBST()

	// 卡塔兰数演示
	fmt.Println("\n=== 卡塔兰数序列 ===")
	demoCatalan()

	// 性能对比测试
	fmt.Println("\n=== 性能对比测试 ===")
	performanceTest()

	// 算法特性总结
	fmt.Println("\n=== 算法特性总结 ===")
	fmt.Println("1. 递归分治:")
	fmt.Println("   - 时间复杂度: O(Cn × n)")
	fmt.Println("   - 空间复杂度: O(Cn)")
	fmt.Println("   - 特点: 最直观，易于理解")
	fmt.Println()
	fmt.Println("2. 记忆化递归:")
	fmt.Println("   - 时间复杂度: O(Cn)")
	fmt.Println("   - 空间复杂度: O(n² × Cn)")
	fmt.Println("   - 特点: 最优解法，避免重复计算")
	fmt.Println()
	fmt.Println("3. 动态规划:")
	fmt.Println("   - 时间复杂度: O(Cn × n)")
	fmt.Println("   - 空间复杂度: O(n² × Cn)")
	fmt.Println("   - 特点: 自底向上，迭代实现")
	fmt.Println()
	fmt.Println("4. 克隆偏移法:")
	fmt.Println("   - 时间复杂度: O(Cn × n)")
	fmt.Println("   - 空间复杂度: O(Cn × n)")
	fmt.Println("   - 特点: 深度克隆，避免引用问题")
	fmt.Println()
	fmt.Println("5. 优化记忆化:")
	fmt.Println("   - 时间复杂度: O(Cn)")
	fmt.Println("   - 空间复杂度: O(n² × Cn)")
	fmt.Println("   - 特点: 字符串键优化，更灵活")
	fmt.Println()
	fmt.Println("注: Cn为第n个卡塔兰数")

	// BST构造演示
	fmt.Println("\n=== BST构造过程演示 ===")
	demoBSTConstruction()
}

// BST结构演示
func demoBST() {
	n := 3
	fmt.Printf("生成所有可能的BST (n=%d):\n\n", n)

	trees := generateTrees1(n)

	for i, tree := range trees {
		fmt.Printf("树 %d:\n", i+1)
		printTree(tree, "", false)

		inorder := inorderTraversal(tree)
		fmt.Printf("中序遍历: %v\n", inorder)
		fmt.Printf("节点数: %d, 高度: %d\n", countNodes(tree), treeHeight(tree))
		fmt.Printf("是否为BST: %v\n\n", isValidBST(tree))
	}

	fmt.Printf("共生成 %d 棵不同的BST\n", len(trees))
}

// 卡塔兰数演示
func demoCatalan() {
	fmt.Println("前10个卡塔兰数:")
	for i := 0; i <= 10; i++ {
		catalan := catalanNumber(i)
		fmt.Printf("C%d = %d", i, catalan)
		if i <= 8 {
			fmt.Printf(" (n=%d时BST数量)\n", i)
		} else {
			fmt.Println()
		}
	}

	fmt.Println("\n卡塔兰数应用场景:")
	fmt.Println("1. 不同的二叉搜索树数量")
	fmt.Println("2. 括号化方案数量")
	fmt.Println("3. 出栈序列数量")
	fmt.Println("4. 凸多边形三角剖分方案数")
	fmt.Println("5. 满二叉树的数量")
}

// 性能测试
func performanceTest() {
	testSizes := []int{3, 4, 5, 6, 7, 8}

	methods := []struct {
		name string
		fn   func(int) []*TreeNode
	}{
		{"递归分治", generateTrees1},
		{"记忆化递归", generateTrees2},
		{"动态规划", generateTrees3},
		{"优化记忆化", generateTrees5},
	}

	for _, n := range testSizes {
		expectedCount := catalanNumber(n)
		fmt.Printf("性能测试 - n=%d (期望%d棵树)\n", n, expectedCount)

		for _, method := range methods {
			start := time.Now()
			result := method.fn(n)
			elapsed := time.Since(start)

			fmt.Printf("  %s: 生成%d棵树, 耗时=%v\n",
				method.name, len(result), elapsed)
		}
		fmt.Println()
	}
}

// BST构造过程演示
func demoBSTConstruction() {
	n := 2
	fmt.Printf("详细演示 n=%d 的构造过程:\n\n", n)

	fmt.Println("步骤1: 选择根节点1")
	fmt.Println("  左子树: [] (空)")
	fmt.Println("  右子树: [2]")
	fmt.Println("  生成树:")
	tree1 := &TreeNode{Val: 1, Right: &TreeNode{Val: 2}}
	printTree(tree1, "    ", false)

	fmt.Println("\n步骤2: 选择根节点2")
	fmt.Println("  左子树: [1]")
	fmt.Println("  右子树: [] (空)")
	fmt.Println("  生成树:")
	tree2 := &TreeNode{Val: 2, Left: &TreeNode{Val: 1}}
	printTree(tree2, "    ", false)

	fmt.Printf("\n共生成 %d 棵不同的BST\n", catalanNumber(n))
	fmt.Println("\n构造过程演示完成!")
}
