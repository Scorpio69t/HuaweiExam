package main

import (
	"fmt"
	"strings"
)

// TreeNode 二叉树节点定义
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// =========================== 方法一：递归DFS（最优雅） ===========================

// maxDepth 递归DFS求最大深度
// 时间复杂度：O(n)，n为节点数，每个节点访问一次
// 空间复杂度：O(h)，h为树高度，递归调用栈深度
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	leftDepth := maxDepth(root.Left)
	rightDepth := maxDepth(root.Right)

	return max(leftDepth, rightDepth) + 1
}

// max 返回两个整数的最大值
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// =========================== 方法二：递归DFS（一行版） ===========================

// maxDepth2 极简一行递归
// 时间复杂度：O(n)
// 空间复杂度：O(h)
func maxDepth2(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(maxDepth2(root.Left), maxDepth2(root.Right)) + 1
}

// =========================== 方法三：迭代BFS（层序遍历） ===========================

// maxDepth3 迭代BFS层序遍历
// 时间复杂度：O(n)
// 空间复杂度：O(w)，w为树的最大宽度
func maxDepth3(root *TreeNode) int {
	if root == nil {
		return 0
	}

	queue := []*TreeNode{root}
	depth := 0

	for len(queue) > 0 {
		size := len(queue)
		depth++

		// 处理当前层的所有节点
		for i := 0; i < size; i++ {
			node := queue[0]
			queue = queue[1:]

			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
	}

	return depth
}

// =========================== 方法四：迭代DFS（栈实现） ===========================

// Pair 节点和深度的配对
type Pair struct {
	node  *TreeNode
	depth int
}

// maxDepth4 迭代DFS使用栈
// 时间复杂度：O(n)
// 空间复杂度：O(h)
func maxDepth4(root *TreeNode) int {
	if root == nil {
		return 0
	}

	stack := []Pair{{root, 1}}
	maxDepth := 0

	for len(stack) > 0 {
		// 出栈
		pair := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		node, depth := pair.node, pair.depth

		// 更新最大深度
		if depth > maxDepth {
			maxDepth = depth
		}

		// 右子节点先入栈（后处理）
		if node.Right != nil {
			stack = append(stack, Pair{node.Right, depth + 1})
		}
		// 左子节点后入栈（先处理）
		if node.Left != nil {
			stack = append(stack, Pair{node.Left, depth + 1})
		}
	}

	return maxDepth
}

// =========================== 方法五：后序遍历（模拟递归） ===========================

// maxDepth5 后序遍历模拟递归过程
// 时间复杂度：O(n)
// 空间复杂度：O(h)
func maxDepth5(root *TreeNode) int {
	if root == nil {
		return 0
	}

	type Frame struct {
		node  *TreeNode
		state int // 0: 未处理, 1: 处理完左子树, 2: 处理完右子树
		depth int
	}

	stack := []Frame{{root, 0, 0}}
	maxDepth := 0
	depths := make(map[*TreeNode]int) // 记录每个节点的深度

	for len(stack) > 0 {
		frame := &stack[len(stack)-1]

		switch frame.state {
		case 0: // 处理左子树
			frame.state = 1
			if frame.node.Left != nil {
				stack = append(stack, Frame{frame.node.Left, 0, 0})
			}

		case 1: // 处理右子树
			frame.state = 2
			if frame.node.Right != nil {
				stack = append(stack, Frame{frame.node.Right, 0, 0})
			}

		case 2: // 计算当前节点深度
			leftDepth := 0
			rightDepth := 0

			if frame.node.Left != nil {
				leftDepth = depths[frame.node.Left]
			}
			if frame.node.Right != nil {
				rightDepth = depths[frame.node.Right]
			}

			currentDepth := max(leftDepth, rightDepth) + 1
			depths[frame.node] = currentDepth

			if currentDepth > maxDepth {
				maxDepth = currentDepth
			}

			stack = stack[:len(stack)-1]
		}
	}

	return maxDepth
}

// =========================== 辅助函数 ===========================

// buildTree 从数组构建二叉树（层序遍历方式）
// -1 表示 nil 节点
func buildTree(arr []int) *TreeNode {
	if len(arr) == 0 || arr[0] == -1 {
		return nil
	}

	root := &TreeNode{Val: arr[0]}
	queue := []*TreeNode{root}
	i := 1

	for len(queue) > 0 && i < len(arr) {
		node := queue[0]
		queue = queue[1:]

		// 左子节点
		if i < len(arr) && arr[i] != -1 {
			node.Left = &TreeNode{Val: arr[i]}
			queue = append(queue, node.Left)
		}
		i++

		// 右子节点
		if i < len(arr) && arr[i] != -1 {
			node.Right = &TreeNode{Val: arr[i]}
			queue = append(queue, node.Right)
		}
		i++
	}

	return root
}

// printTree 打印树结构（层序遍历）
func printTree(root *TreeNode) {
	if root == nil {
		fmt.Println("[]")
		return
	}

	queue := []*TreeNode{root}
	var result []string

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		if node == nil {
			result = append(result, "null")
		} else {
			result = append(result, fmt.Sprintf("%d", node.Val))
			queue = append(queue, node.Left)
			queue = append(queue, node.Right)
		}
	}

	// 移除末尾的null
	for len(result) > 0 && result[len(result)-1] == "null" {
		result = result[:len(result)-1]
	}

	fmt.Print("[")
	for i, v := range result {
		if i > 0 {
			fmt.Print(",")
		}
		fmt.Print(v)
	}
	fmt.Println("]")
}

// visualizeTree 可视化打印树结构
func visualizeTree(root *TreeNode, prefix string, isLeft bool) {
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
			newPrefix := prefix
			if isLeft {
				newPrefix += "│   "
			} else {
				newPrefix += "    "
			}
			visualizeTree(root.Left, newPrefix, true)
		}

		if root.Right != nil {
			newPrefix := prefix
			if isLeft {
				newPrefix += "│   "
			} else {
				newPrefix += "    "
			}
			visualizeTree(root.Right, newPrefix, false)
		}
	}
}

// =========================== 扩展功能 ===========================

// minDepth 计算最小深度
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	// 叶子节点
	if root.Left == nil && root.Right == nil {
		return 1
	}

	// 只有右子树
	if root.Left == nil {
		return minDepth(root.Right) + 1
	}

	// 只有左子树
	if root.Right == nil {
		return minDepth(root.Left) + 1
	}

	// 都有，取最小
	return min(minDepth(root.Left), minDepth(root.Right)) + 1
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// isBalanced 判断是否平衡二叉树
func isBalanced(root *TreeNode) bool {
	return checkBalance(root) != -1
}

func checkBalance(root *TreeNode) int {
	if root == nil {
		return 0
	}

	left := checkBalance(root.Left)
	if left == -1 {
		return -1
	}

	right := checkBalance(root.Right)
	if right == -1 {
		return -1
	}

	if abs(left-right) > 1 {
		return -1
	}

	return max(left, right) + 1
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// countNodes 计算节点总数
func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return countNodes(root.Left) + countNodes(root.Right) + 1
}

// maxDepthPath 返回达到最大深度的路径
func maxDepthPath(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	leftPath := maxDepthPath(root.Left)
	rightPath := maxDepthPath(root.Right)

	// 选择更深的路径
	if len(leftPath) > len(rightPath) {
		return append([]int{root.Val}, leftPath...)
	}
	return append([]int{root.Val}, rightPath...)
}

// =========================== 测试代码 ===========================

func main() {
	fmt.Println("=== LeetCode 104: 二叉树的最大深度 ===\n")

	// 测试用例
	testCases := []struct {
		name   string
		arr    []int
		expect int
	}{
		{
			name:   "示例1: 完整二叉树",
			arr:    []int{3, 9, 20, -1, -1, 15, 7},
			expect: 3,
		},
		{
			name:   "示例2: 偏斜树",
			arr:    []int{1, -1, 2},
			expect: 2,
		},
		{
			name:   "边界: 空树",
			arr:    []int{},
			expect: 0,
		},
		{
			name:   "边界: 单节点",
			arr:    []int{0},
			expect: 1,
		},
		{
			name:   "完全二叉树",
			arr:    []int{1, 2, 3, 4, 5, 6, 7},
			expect: 3,
		},
		{
			name:   "左偏树",
			arr:    []int{1, 2, -1, 3, -1, -1, -1, 4},
			expect: 4,
		},
		{
			name:   "右偏树",
			arr:    []int{1, -1, 2, -1, -1, -1, 3, -1, -1, -1, -1, -1, -1, -1, 4},
			expect: 4,
		},
		{
			name:   "不平衡树",
			arr:    []int{1, 2, 3, 4, -1, -1, 5, 6, 7},
			expect: 4,
		},
	}

	methods := []struct {
		name string
		fn   func(*TreeNode) int
	}{
		{"方法一：递归DFS", maxDepth},
		{"方法二：递归DFS一行版", maxDepth2},
		{"方法三：迭代BFS", maxDepth3},
		{"方法四：迭代DFS栈", maxDepth4},
		{"方法五：后序遍历", maxDepth5},
	}

	// 对每种方法运行测试
	for _, method := range methods {
		fmt.Printf("\n%s\n", method.name)
		fmt.Println(strings.Repeat("=", 60))

		passCount := 0
		for i, tc := range testCases {
			root := buildTree(tc.arr)
			result := method.fn(root)

			status := "✅"
			if result != tc.expect {
				status = "❌"
			} else {
				passCount++
			}

			fmt.Printf("  测试%d: %s\n", i+1, status)
			fmt.Printf("    名称: %s\n", tc.name)
			fmt.Printf("    输入: ")
			printTree(root)
			fmt.Printf("    输出: %d\n", result)
			if result != tc.expect {
				fmt.Printf("    期望: %d\n", tc.expect)
			}

			// 为第一个示例打印树结构
			if i == 0 {
				fmt.Println("    树结构:")
				if root != nil {
					visualizeTree(root, "      ", false)
				}
			}
		}

		fmt.Printf("\n  通过: %d/%d\n", passCount, len(testCases))
	}

	// 扩展功能测试
	fmt.Println("\n\n=== 扩展功能测试 ===\n")
	testExtensions()

	// 性能对比
	fmt.Println("\n=== 性能对比 ===\n")
	performanceTest()
}

// testExtensions 测试扩展功能
func testExtensions() {
	fmt.Println("1. 最小深度测试")
	root1 := buildTree([]int{3, 9, 20, -1, -1, 15, 7})
	fmt.Printf("   树: ")
	printTree(root1)
	fmt.Printf("   最大深度: %d\n", maxDepth(root1))
	fmt.Printf("   最小深度: %d\n", minDepth(root1))

	fmt.Println("\n2. 平衡性检查")
	root2 := buildTree([]int{1, 2, 3, 4, 5, 6, 7})
	fmt.Printf("   树: ")
	printTree(root2)
	fmt.Printf("   是否平衡: %v\n", isBalanced(root2))

	root3 := buildTree([]int{1, 2, -1, 3, -1, -1, -1, 4})
	fmt.Printf("   树: ")
	printTree(root3)
	fmt.Printf("   是否平衡: %v\n", isBalanced(root3))

	fmt.Println("\n3. 节点计数")
	root4 := buildTree([]int{1, 2, 3, 4, 5, 6, 7})
	fmt.Printf("   树: ")
	printTree(root4)
	fmt.Printf("   节点总数: %d\n", countNodes(root4))

	fmt.Println("\n4. 最大深度路径")
	root5 := buildTree([]int{3, 9, 20, -1, -1, 15, 7})
	fmt.Printf("   树: ")
	printTree(root5)
	fmt.Printf("   最大深度路径: %v\n", maxDepthPath(root5))
}

// performanceTest 性能测试
func performanceTest() {
	// 构建深度为15的完全二叉树
	size := (1 << 15) - 1 // 2^15 - 1 = 32767个节点
	arr := make([]int, size)
	for i := 0; i < size; i++ {
		arr[i] = i + 1
	}

	fmt.Printf("测试数据：完全二叉树，节点数=%d，期望深度=15\n\n", size)

	root := buildTree(arr)

	fmt.Println("各方法性能测试:")
	result1 := maxDepth(root)
	fmt.Printf("  方法一（递归DFS）: 深度=%d\n", result1)

	result2 := maxDepth2(root)
	fmt.Printf("  方法二（一行递归）: 深度=%d\n", result2)

	result3 := maxDepth3(root)
	fmt.Printf("  方法三（迭代BFS）: 深度=%d\n", result3)

	result4 := maxDepth4(root)
	fmt.Printf("  方法四（迭代DFS）: 深度=%d\n", result4)

	result5 := maxDepth5(root)
	fmt.Printf("  方法五（后序遍历）: 深度=%d\n", result5)

	fmt.Println("\n说明：所有方法时间复杂度均为O(n)，空间复杂度因方法而异")
	fmt.Println("  - 递归方法：O(h)栈空间，h为树高度")
	fmt.Println("  - BFS方法：O(w)队列空间，w为树最大宽度")
	fmt.Println("  - DFS栈方法：O(h)栈空间")
}
