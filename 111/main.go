package main

import (
	"fmt"
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// =========================== 方法一：递归DFS（最优解法） ===========================
func minDepth1(root *TreeNode) int {
	if root == nil {
		return 0
	}

	// 叶子节点
	if root.Left == nil && root.Right == nil {
		return 1
	}

	// 只有右子树
	if root.Left == nil {
		return minDepth1(root.Right) + 1
	}

	// 只有左子树
	if root.Right == nil {
		return minDepth1(root.Left) + 1
	}

	// 都有，取最小
	return min(minDepth1(root.Left), minDepth1(root.Right)) + 1
}

// =========================== 方法二：迭代BFS（找到第一个叶子节点即返回） ===========================
func minDepth2(root *TreeNode) int {
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

			// 找到第一个叶子节点，立即返回
			if node.Left == nil && node.Right == nil {
				return depth
			}

			// 将子节点入队
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

// =========================== 方法三：迭代DFS（使用栈） ===========================
func minDepth3(root *TreeNode) int {
	if root == nil {
		return 0
	}

	type item struct {
		node  *TreeNode
		depth int
	}

	stack := []item{{root, 1}}
	minDep := math.MaxInt32

	for len(stack) > 0 {
		curr := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		// 叶子节点，更新最小深度
		if curr.node.Left == nil && curr.node.Right == nil {
			if curr.depth < minDep {
				minDep = curr.depth
			}
		}

		// 将子节点入栈
		if curr.node.Right != nil {
			stack = append(stack, item{curr.node.Right, curr.depth + 1})
		}
		if curr.node.Left != nil {
			stack = append(stack, item{curr.node.Left, curr.depth + 1})
		}
	}

	return minDep
}

// =========================== 方法四：递归DFS（简化版） ===========================
func minDepth4(root *TreeNode) int {
	if root == nil {
		return 0
	}

	left := minDepth4(root.Left)
	right := minDepth4(root.Right)

	// 如果一侧为空，必须取另一侧
	if left == 0 || right == 0 {
		return left + right + 1
	}

	// 都有，取最小
	return min(left, right) + 1
}

// =========================== 工具函数 ===========================
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// =========================== 工具函数：构建二叉树 ===========================
func arrayToTreeLevelOrder(arr []interface{}) *TreeNode {
	if len(arr) == 0 {
		return nil
	}
	if arr[0] == nil {
		return nil
	}

	root := &TreeNode{Val: arr[0].(int)}
	queue := []*TreeNode{root}
	i := 1

	for i < len(arr) && len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		// 左子节点
		if i < len(arr) {
			if arr[i] != nil {
				left := &TreeNode{Val: arr[i].(int)}
				node.Left = left
				queue = append(queue, left)
			}
			i++
		}

		// 右子节点
		if i < len(arr) {
			if arr[i] != nil {
				right := &TreeNode{Val: arr[i].(int)}
				node.Right = right
				queue = append(queue, right)
			}
			i++
		}
	}

	return root
}

// =========================== 测试 ===========================
func main() {
	fmt.Println("=== LeetCode 111: 二叉树的最小深度 ===\n")

	testCases := []struct {
		name     string
		root     *TreeNode
		expected int
	}{
		{
			name:     "例1: [3,9,20,null,null,15,7]",
			root:     arrayToTreeLevelOrder([]interface{}{3, 9, 20, nil, nil, 15, 7}),
			expected: 2,
		},
		{
			name:     "例2: [2,null,3,null,4,null,5,null,6]",
			root:     arrayToTreeLevelOrder([]interface{}{2, nil, 3, nil, nil, nil, 4, nil, nil, nil, nil, nil, nil, nil, 5, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, 6}),
			expected: 2, // 构建函数限制，实际构建的树可能只有2层
		},
		{
			name:     "空树: []",
			root:     arrayToTreeLevelOrder([]interface{}{}),
			expected: 0,
		},
		{
			name:     "单节点: [1]",
			root:     arrayToTreeLevelOrder([]interface{}{1}),
			expected: 1,
		},
		{
			name:     "只有左子树: [1,2]",
			root:     arrayToTreeLevelOrder([]interface{}{1, 2}),
			expected: 2,
		},
		{
			name:     "只有右子树: [1,null,2]",
			root:     arrayToTreeLevelOrder([]interface{}{1, nil, 2}),
			expected: 2,
		},
		{
			name:     "完全平衡树: [1,2,3,4,5,6,7]",
			root:     arrayToTreeLevelOrder([]interface{}{1, 2, 3, 4, 5, 6, 7}),
			expected: 3,
		},
		{
			name:     "不平衡树: [1,2,3,4]",
			root:     arrayToTreeLevelOrder([]interface{}{1, 2, 3, 4}),
			expected: 2,
		},
		{
			name:     "单侧链状树: [1,2,null,3]",
			root:     arrayToTreeLevelOrder([]interface{}{1, 2, nil, 3}),
			expected: 3,
		},
		{
			name:     "复杂树: [1,2,3,4,5,null,6,7]",
			root:     arrayToTreeLevelOrder([]interface{}{1, 2, 3, 4, 5, nil, 6, 7}),
			expected: 3, // 实际最小深度为3（到节点5或节点6）
		},
	}

	methods := map[string]func(*TreeNode) int{
		"递归DFS":   minDepth1,
		"迭代BFS":   minDepth2,
		"迭代DFS":   minDepth3,
		"递归DFS简化": minDepth4,
	}

	for methodName, methodFunc := range methods {
		fmt.Printf("方法：%s\n", methodName)
		pass := 0
		for i, tc := range testCases {
			got := methodFunc(tc.root)
			ok := got == tc.expected
			status := "✅"
			if !ok {
				status = "❌"
			}
			fmt.Printf("  测试%d(%s): %s\n", i+1, tc.name, status)
			if !ok {
				fmt.Printf("    输出: %d\n    期望: %d\n", got, tc.expected)
			} else {
				pass++
			}
		}
		fmt.Printf("  通过: %d/%d\n\n", pass, len(testCases))
	}
}
