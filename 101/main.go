package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// =========================== 方法一：递归比较（最优解法） ===========================
func isSymmetric1(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return isMirror1(root.Left, root.Right)
}

func isMirror1(left *TreeNode, right *TreeNode) bool {
	// 两棵树都为空
	if left == nil && right == nil {
		return true
	}

	// 一棵树为空，另一棵不为空
	if left == nil || right == nil {
		return false
	}

	// 节点值不同
	if left.Val != right.Val {
		return false
	}

	// 递归比较镜像位置：左左对右右，左右对右左
	return isMirror1(left.Left, right.Right) && isMirror1(left.Right, right.Left)
}

// =========================== 方法二：BFS迭代比较 ===========================
func isSymmetric2(root *TreeNode) bool {
	if root == nil {
		return true
	}

	leftQueue := []*TreeNode{root.Left}
	rightQueue := []*TreeNode{root.Right}

	for len(leftQueue) > 0 {
		// 同时出队
		leftNode := leftQueue[0]
		rightNode := rightQueue[0]
		leftQueue = leftQueue[1:]
		rightQueue = rightQueue[1:]

		// 检查nil
		if leftNode == nil && rightNode == nil {
			continue
		}
		if leftNode == nil || rightNode == nil {
			return false
		}
		if leftNode.Val != rightNode.Val {
			return false
		}

		// 镜像入队：左左对右右，左右对右左
		leftQueue = append(leftQueue, leftNode.Left, leftNode.Right)
		rightQueue = append(rightQueue, rightNode.Right, rightNode.Left)
	}

	return true
}

// =========================== 方法三：DFS迭代比较 ===========================
func isSymmetric3(root *TreeNode) bool {
	if root == nil {
		return true
	}

	leftStack := []*TreeNode{root.Left}
	rightStack := []*TreeNode{root.Right}

	for len(leftStack) > 0 {
		// 同时出栈
		n1 := len(leftStack) - 1
		leftNode := leftStack[n1]
		rightNode := rightStack[n1]
		leftStack = leftStack[:n1]
		rightStack = rightStack[:n1]

		// 检查nil
		if leftNode == nil && rightNode == nil {
			continue
		}
		if leftNode == nil || rightNode == nil {
			return false
		}
		if leftNode.Val != rightNode.Val {
			return false
		}

		// 镜像入栈：左左对右右，左右对右左（注意顺序：先右后左，因为栈是后进先出）
		leftStack = append(leftStack, leftNode.Right, leftNode.Left)
		rightStack = append(rightStack, rightNode.Left, rightNode.Right)
	}

	return true
}

// =========================== 方法四：翻转右子树后比较 ===========================
func isSymmetric4(root *TreeNode) bool {
	if root == nil {
		return true
	}

	// 翻转右子树
	root.Right = invertTree4(root.Right)

	// 比较左子树和翻转后的右子树
	result := isSameTree4(root.Left, root.Right)

	// 恢复右子树（翻转回来）
	root.Right = invertTree4(root.Right)

	return result
}

func invertTree4(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	root.Left, root.Right = invertTree4(root.Right), invertTree4(root.Left)
	return root
}

func isSameTree4(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}
	return p.Val == q.Val && isSameTree4(p.Left, q.Left) && isSameTree4(p.Right, q.Right)
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
		if i < len(arr) && arr[i] != nil {
			left := &TreeNode{Val: arr[i].(int)}
			node.Left = left
			queue = append(queue, left)
		}
		i++

		// 右子节点
		if i < len(arr) && arr[i] != nil {
			right := &TreeNode{Val: arr[i].(int)}
			node.Right = right
			queue = append(queue, right)
		}
		i++
	}

	return root
}

// =========================== 测试 ===========================
func main() {
	fmt.Println("=== LeetCode 101: 对称二叉树 ===\n")

	testCases := []struct {
		name     string
		root     *TreeNode
		expected bool
	}{
		{
			name:     "例1: [1,2,2,3,4,4,3] - 对称",
			root:     arrayToTreeLevelOrder([]interface{}{1, 2, 2, 3, 4, 4, 3}),
			expected: true,
		},
		{
			name:     "例2: [1,2,2,null,3,null,3] - 不对称",
			root:     arrayToTreeLevelOrder([]interface{}{1, 2, 2, nil, 3, nil, 3}),
			expected: false,
		},
		{
			name:     "空树 - 对称",
			root:     arrayToTreeLevelOrder([]interface{}{}),
			expected: true,
		},
		{
			name:     "单节点: [1] - 对称",
			root:     arrayToTreeLevelOrder([]interface{}{1}),
			expected: true,
		},
		{
			name:     "完全对称: [1,2,2,3,4,4,3] - 对称",
			root:     arrayToTreeLevelOrder([]interface{}{1, 2, 2, 3, 4, 4, 3}),
			expected: true,
		},
		{
			name:     "值不对称: [1,2,2,3,4,5,3] - 不对称",
			root:     arrayToTreeLevelOrder([]interface{}{1, 2, 2, 3, 4, 5, 3}),
			expected: false,
		},
		{
			name:     "单层对称: [1,2,2] - 对称",
			root:     arrayToTreeLevelOrder([]interface{}{1, 2, 2}),
			expected: true,
		},
		{
			name:     "单层不对称: [1,2,3] - 不对称",
			root:     arrayToTreeLevelOrder([]interface{}{1, 2, 3}),
			expected: false,
		},
		{
			name:     "链状树: [1,2,null] - 不对称",
			root:     arrayToTreeLevelOrder([]interface{}{1, 2}),
			expected: false,
		},
		{
			name:     "复杂不对称: [1,2,2,3,null,3,null] - 不对称",
			root:     arrayToTreeLevelOrder([]interface{}{1, 2, 2, 3, nil, 3}),
			expected: false,
		},
	}

	methods := map[string]func(*TreeNode) bool{
		"递归比较":   isSymmetric1,
		"BFS迭代":  isSymmetric2,
		"DFS迭代":  isSymmetric3,
		"翻转比较": isSymmetric4,
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
				fmt.Printf("    输出: %v\n    期望: %v\n", got, tc.expected)
			} else {
				pass++
			}
		}
		fmt.Printf("  通过: %d/%d\n\n", pass, len(testCases))
	}
}